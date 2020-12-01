[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]
![GitHub go.mod Go version][goversion-url]
![Go Report Card][goreport-url]
[![Coverage Status](https://coveralls.io/repos/github/michaelmcallister/coord2locality/badge.svg)](https://coveralls.io/github/michaelmcallister/coord2locality)
[![Build Status](https://travis-ci.org/michaelmcallister/coord2locality.svg?branch=main)](https://travis-ci.org/michaelmcallister/coord2locality)
<!-- PROJECT LOGO -->
<br />
<p align="center">
  <h3 align="center">coord2locality</h3>

  <p align="center">
    A simple Go library that can be used to find the nearest Lat/Long from the supplied one given a list of them (Offline!)
  </p>
</p>

## Usage

### As a CLI
You will need a JSON file containing the following fields:

- Postcode
-	Locality
-	State
-	Lat
-	Long

for Australian postcodes I highly recommend [this one](https://www.matthewproctor.com/australian_postcodes)

It's then a matter of supplying the coordinates of interest and the output will be the nearest locality from the supplied file.

See below for Parliament House in Canberra, ACT, Australia

```
 ./coord2locality -file australian_postcodes.json -lat -35.3082237 -long 149.1222036
{Postcode:2603 Locality:FORREST State:ACT Lat:-35.330382 Long:149.125739}
```

### As A Library
You will need to supply a JSON file that satisfies the Location struct as an argument to `New` and then called `CoordinatesToLocation` with the latitude and longitude of interest

```
  import (
    "fmt"
    "github.com/michaelmcallister/coord2locality"
  )
  func main() {
    c := coord2locality.New("postcodes.json")
    result, _ := c.CoordinatesToLocation(-35.3082237, 149.1222036)
    fmt.Printf("%+v\n", result)
  }
```

<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.

<!-- ACKNOWLEDGEMENTS -->
## Acknowledgements

* [Chris Veness](https://www.movable-type.co.uk/scripts/latlong.html) for his explanation of the haversine formula.
* [Matthew Proctor](https://www.matthewproctor.com/australian_postcodes) for his Australian postcode data.

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[license-shield]: https://img.shields.io/github/license/michaelmcallister/coord2locality.svg?style=flat-square
[license-url]: https://github.com/michaelmcallister/coord2locality/blob/master/LICENSE
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=flat-square&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/mpmcallister
[goversion-url]: https://img.shields.io/github/go-mod/go-version/michaelmcallister/coord2locality
[goreport-url]: https://goreportcard.com/badge/michaelmcallister/coord2locality