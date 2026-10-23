package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]uint32)
    var v3 uint16 = 255
    var v4 uint32 = 2147483647
    v5 := make(map[uint32]uint32)
    fmt.Println(v3, v3)
    v4 -= v4
    if (!false) {
        switch -9223372036854775808 {
        case -2147483648:
            v3 -= v3
        case 32767:
            if ((-2147483648 > 790) && false) {
                v3 -= (v3 - (v3 + v3))
                for fi6 := 0; fi6 < 1; fi6++ {
                    if false {
                        _ = len(v2)
                        switch (fi6 - (^(fi6))) {
                        case 127:
                            var v7 uint8 = 127
                            _ = v7
                        case 32768:
                            switch (fi6 + fi6) {
                            case -158:
                                var v8 int8 = 0
                                _ = v8
                            }
                        }
                    }
                    v4 -= v4
                    _ = fi6
                }
                fmt.Println(v4, v3)
            }
        case -1:
            var v9 uint8 = 102
            _ = v9
        }
    } else {
        fmt.Println(v3, v4)
    }
    for fi10 := 0; fi10 < 15; fi10++ {
        switch fi10 {
        case 0:
            for fi11 := 0; fi11 < 10; fi11++ {
                v12 := GenS1{F0: 2, F1: fi11}
                var v13 float64 = (((0.0 + (3.14159 - 1.5)) / float64((float32((1e10 * float64((-2.25 - 3.14159)))) + 0.0))) * ((1e10 + 1.5) - (1e10 / 0.1)))
                v13 = v13
                _ = fi11
                _ = v12
                _ = v13
            }
        }
        var v14 uint = 943
        _ = fi10
        _ = v14
    }
    switch "a" {
    case "foo bar":
        switch -1 {
        case 32767:
            if (!true) {
                if ((1.5 / (1.5 / 1.0)) == 0.1) {
                    for fi15 := 0; fi15 < 3; fi15++ {
                        var v16 int = (fi15 & fi15)
                        for fi17 := 0; fi17 < 3; fi17++ {
                            v18 := GenS1{F0: 95, F1: 127}
                            _ = fi17
                            _ = v18
                        }
                        _ = fi15
                        _ = v16
                    }
                    fmt.Println(v4, v3)
                }
            }
        case 9223372036854775807:
            delete(v2, 65535)
        case 128:
            if false {
                fmt.Println(v3, v3)
                var v19 float32 = 1.5
                _ = v19
            }
        }
    default:
        var v20 int16 = 1
        _ = v20
    }
    delete(v5, (v4 << ((255) & 63)))
    v3 += v3
    v4 = (v4 % (((^(v4))) | 1))
    v21 := v5[(v4 ^ 65535)]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
