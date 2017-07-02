
http://www.imei.info/faq/

## GSM

### International Mobile Station Equipment Identity (IMEI)

The International Mobile Station Equipment Identity (IMEI) looks more like a serial number which distinctively identifies a mobile station internationally. This is allocated by the equipment manufacturer and registered by the network operator, who stores it in the Entrepreneurs-in-Residence (EIR). By means of IMEI, one recognizes obsolete, stolen, or non-functional equipment.

Following are the parts of IMEI:

* Type Approval Code (TAC) : 6 decimal places, centrally assigned.

* Final Assembly Code (FAC) : 6 decimal places, assigned by the manufacturer.

* Serial Number (SNR) : 6 decimal places, assigned by the manufacturer.

* Spare (SP) : 1 decimal place.

Thus, IMEI = TAC + FAC + SNR + SP - 
#### 19 decimal digits

* The IMEI is a permanent identity assigned by the Device Manufacturer
* Valid as long as the Device is in Use
( )Stored on the Device hardware and on the HSS (Home Subscriber Server)

It uniquely characterizes a mobile station and gives clues about the manufacturer and the date of manufacturing.

### International Mobile Subscriber Identity (IMSI)

Every registered user has an original International Mobile Subscriber Identity (IMSI) with a valid IMEI stored in their Subscriber Identity Module (SIM).

IMSI comprises of the following parts:

* Mobile Country Code (MCC) : 3 decimal places, internationally standardized.

* Mobile Network Code (MNC) : 2 decimal places, for unique identification of mobile network within the country.

* Mobile Subscriber Identification Number (MSIN) : Maximum 10 decimal places, identification number of the subscriber in the home mobile network.

The IMSI is a permanent identity assigned by the Service Provider
* It is valid as long as the Service is Active with the Service Provider
* It is stored on the USIM card and on the HSS (Home Subscriber Server)
* It globally and uniquely identifies a user on any 3GPP PLMN (Public Land Mobile Network)

All signaling and messaging in GSM and UMTS networks uses the IMSI as the primary identifier of a subscriber.

The IMSI is one of the pieces of information stored on a SIM card.

#### 15 decimal digits


## Mobile Subscriber ISDN Number (MSISDN)/MDN

The authentic telephone number of a mobile station is the Mobile Subscriber ISDN Number (MSISDN).
Based on the SIM, a mobile station can have many MSISDNs, as each subscriber is assigned with
a separate MSISDN to their SIM respectively.

Listed below is the structure followed by MSISDN categories,
as they are defined based on international ISDN number plan:

* Country Code (CC) : Up to 3 decimal places.

* National Destination Code (NDC) : Typically 2-3 decimal places.

* Subscriber Number (SN) : Maximum 10 decimal places.

#### 10-digit phone number of 2G or 3G device; 15-digit number of 4G device.

### ICCID = Integrated Circuit Card ID. 

This is the identifier of the actual SIM card itself - i.e. an identifier for the SIM chip. 
It is possible to change the information contained on a SIM (including the IMSI), 
but the identify of the SIM itself remains the same.

#### up to 22 digits long, including a single check digit calculated using the Luhn algorithm

### Electronic serial number (ESN) 

Created by the U.S. Federal Communications Commission (FCC) to uniquely identify mobile devices, from the days of AMPS in the United States starting in the early 1980s. The administrative role was taken over by the Telecommunications Industry Association in 1997 and is still maintained by them. ESNs are currently mainly used with CDMA phones (and were previously used by AMPS and TDMA phones), compared to International Mobile Equipment Identity (IMEI) numbers used by all GSM phones.[1]

The first 8 bits of the ESN was originally the manufacturer code, leaving 24 bits for the manufacturer to assign up to 16,777,215 codes to mobiles. To allow more than 256 manufacturers to be identified the manufacturer code was extended to 14 bits, leaving 18 bits for the manufacturer to assign up to 262,144 codes. Manufacturer code 0x80 is reserved from assignment and is used instead as an 8-bit prefix for pseudo-ESNs (pESN). The remaining 24 bits are the least significant bits of the SHA-1 hash of a mobile equipment identifier (MEID). Pseudo-ESNs are not guaranteed to be unique (the MEID is the unique identifier if the phone has a pseudo-ESN).

ESNs are often represented as either 11-digit decimal numbers or 8 digit hexadecimal numbers. For the decimal format the first three digits are the decimal representation of the first 8 bits (between 000 and 255 inclusive) and the next 8 digits are derived from the remaining 24 bits and will be between 00000000 and 16777215 inclusive. The decimal format of pseudo ESNs will therefore begin with 128. The decimal format separately displays 8 bit manufacturer codes in the first 3 digits, but 14 bit codes are not displayed as separate digits. The hexadecimal format displays an ESN as 8 digits and also does not separately display 14 bit manufacturer codes which occupy 3.5 hexadecimal digits.

#### Mostly for CDMA - migrated to MEID

### Mobile equipment identifier (MEID) 

Globally unique number identifying a physical piece of CDMA2000 mobile station equipment. 
The number format is defined by the 3GPP2 report S.R0048 
but in practical terms, it can be seen as an 
#### IMEI but with hexadecimal digits.

The decimal form is specified to be 18 digits grouped in a 5 5 4 4 pattern and is calculated by converting the manufacturer code portion (32 bits) to decimal and padding on the left with '0' digits to 10 digits and separately converting the serial number portion to decimal and padding on the left to 8 digits. A check-digit can be calculated from the 18 digit result using the standard base 10 Luhn algorithm and appended to the end. Note that to produce this form the MEID digits are treated as base 16 numbers even if all of them are in the range '0'-'9'.


#### max int64 is 9,223,372,036,854,775,807  - 18 digits (safe)

#### UUID 

string representations: 36 characters - 32 hex digits + 4 dashes.
##### 16 bytes
