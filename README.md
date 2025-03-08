# Magic Swap

**Magic Swap** is a blockchain that has the capability of mixed send and swap with two-way swap support.

## Get started

```bash
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

use this for a fresh state:

```bash
ignite chain serve --reset-once
```

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

#### Customization

The `config.yml` is modified in a way to support two-way token send and transfer:

```yml
- src_denom: stake
    dst_denom: stake1
    ratio: "0.50"
    min_supply_ratio_limit: "0.6667"
- src_denom: stake1
    dst_denom: stake
    ratio: "0.80"
    min_supply_ratio_limit: "0.3333"
```

This enables the blockchain to be able to support swapping from stake to stake1 and stake1 to stake.

- `ratio` is set as 50% for stake to stake1, and 33% for stake1 to stake which means, 50% of the stake token will be burnt and swapped to stake1, and 80% of the stake1 will be burnt and swapped to stake.
- `min_supply_ratio_limit` swapping from stake to stake1 is allowed only if the division of stake supply to stake1 supply should not be less than this parameter.

### CLI Application Logic Demonstration

The following steps helps demonstrate the swapping capabilities and validations.

Swap Even value `10000000stake` from alice. it will be slitted into two equal values as output. `5000000stake` will be swapped to `5000000stake1`, and `5000000stake` will be transferred to bob's account.

```bash
magicswapd tx swap partial-send $(magicswapd keys show bob -a)  10000000stake  --from alice -y
```

Swap Odd value `10000001stake` from alice. rounded 50% of this value will be burned `5000000stake1` and the rest of value will be sent `5000001stake` to bob's account.

```bash
magicswapd tx swap partial-send $(magicswapd keys show bob -a)  10000001stake  --from alice -y
```

To query all of the interaction by alice so far, the following query and output helps (addresses may differ in your environment):

```bash
magicswapd q swap list-partial-send-by-creator $(magicswapd keys show alice -a)
```

```yml
PartialSend:
- burntAmount:
    amount: "5000000"
    denom: stake
  creator: cosmos15csw4rtqzletlagr4f72nelxcue7e3686hhfwl
  recipient: cosmos1naw9g8h6pgpe332uda8hy49p0934rrn6jrnj7a
  sentAmount:
    amount: "5000000"
    denom: stake
  swappedAmount:
    amount: "5000000"
    denom: stake1
- burntAmount:
    amount: "5000000"
    denom: stake
  creator: cosmos15csw4rtqzletlagr4f72nelxcue7e3686hhfwl
  id: "1"
  recipient: cosmos1naw9g8h6pgpe332uda8hy49p0934rrn6jrnj7a
  sentAmount:
    amount: "5000001"
    denom: stake
  swappedAmount:
    amount: "5000000"
    denom: stake1
pagination:
  total: "2"
```

Checking the current balance of alice's account:

```bash
magicswapd q bank balances $(magicswapd keys show alice -a)
```

```yml
balances:
- amount: "79999999"
  denom: stake
pagination:
  total: "1"
```

```bash
magicswapd q bank balances $(magicswapd keys show bob -a)
```

```yml
balances:
- amount: "110000001"
  denom: stake
- amount: "10000000"
  denom: stake1
pagination:
  total: "2"
```

Now let's try the reverse swap from stake1 to stake, the following transaction will fail will not be committed to the finalized block. the reason is the `min_supply_ratio_limit` which is 33% that prevents stake1 to be transferred as long as the supply ratio is less than configured values in the module parameters.

```bash
magicswapd tx swap partial-send $(magicswapd keys show alice -a)  500stake1  --from bob -y
```

this is the raw log of the transaction error, the application logic prevents this transaction because the ratio of supply values is less than the parameter.

```text
raw_log: 'failed to execute message; message index: 0: swap rule validation failed:
  ratio 0.034481018091955299 is below the minimum supply ratio limit 0.333300000000000000:
```

Let's do more swap from alice to bob.

```bash
magicswapd tx swap partial-send $(magicswapd keys show bob -a)  60000000stake  --from alice -y
```

and more self-swap from bob's account.

```bash
magicswapd tx swap partial-send $(magicswapd keys show bob -a)  60000000stake  --from bob -y
```

now the same failed command works as expected and the transaction will be finalized in the blockchain.

```bash
magicswapd tx swap partial-send $(magicswapd keys show alice -a)  500stake1  --from bob -y
```

Now we can query the current state of the swaps from bob's account.

```bash
magicswapd q swap list-partial-send-by-creator $(magicswapd keys show bob -a)
```

```yml
- burntAmount:
    amount: "60000000"
    denom: stake
  creator: cosmos1naw9g8h6pgpe332uda8hy49p0934rrn6jrnj7a
  id: "3"
  recipient: cosmos1naw9g8h6pgpe332uda8hy49p0934rrn6jrnj7a
  sentAmount:
    amount: "0"
    denom: stake
  swappedAmount:
    amount: "60000000"
    denom: stake1
- burntAmount:
    amount: "400"
    denom: stake1
  creator: cosmos1naw9g8h6pgpe332uda8hy49p0934rrn6jrnj7a
  id: "4"
  recipient: cosmos15csw4rtqzletlagr4f72nelxcue7e3686hhfwl
  sentAmount:
    amount: "100"
    denom: stake1
  swappedAmount:
    amount: "400"
    denom: stake
```
