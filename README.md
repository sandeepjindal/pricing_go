### Resources

You will find attached

- A code stub in Go implementing an RPC endpoint



Take a look at the code stub and familiarise yourself with the structure.
This implements an RPC endpoint (`generate_pricing`) which is called with a single users details.

and translate it to a "factor". These factors get multiplied with the "Base Rate" price (found in the "Base Rates"
sheet) to form a total price the user should pay. The base rate changes depending on the duration of cover selected.
We typically give the user a price for every duration of cover we offer, letting the user pick the duration that makes
sense for them.
Take a look at the "Calculator" sheet and familiarise yourself with how the total price a user pays is calculated.


Update the `generate_pricing` endpoint to accept the following example input.

```json
{
    "date_of_birth": "1970-12-04",
    "insurance_group": 12,
    "license_held_since": "1988-08-01"
}
```

This represents details about a single user.


Update the `generate_pricing` endpoint to respond with the total price a user should pay for each duration of cover
found in the base rates sheet.

The output must be in JSON, the schema of the output is left to you.
