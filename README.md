unidecode
=========

Unicode transliterator in Golang - Replaces non-GSM characters with their GSM approximations.

References:
- https://www.unicode.org/Public/MAPPINGS/ETSI/GSM0338.TXT
- https://www.etsi.org/deliver/etsi_ts/123000_123099/123038/05.00.00_60/ts_123038v050000p.pdf

Note: this transliterator does not include characters in the GSM extension table. NBSP is displayed as a space character.
