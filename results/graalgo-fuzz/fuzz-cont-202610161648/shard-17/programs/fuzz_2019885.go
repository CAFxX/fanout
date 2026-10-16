package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 int64, p1 string, p2 bool) uint8 {
    p2 = true
    if p2 {
        if false {
            p2 = (!(!(255 < 65535)))
            switch "\n\t" {
            case "":
                p0 = 9223372036854775807
            case "a":
                fmt.Println(p2)
            default:
                fmt.Println(p1)
            }
        }
    } else {
        var v6 uint32 = 65535
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 int16 = 1
    v2 := []uint32{4294967295, 2147483648, 65535, 917}
    var v3 uint16 = 0
    var v4 int32 = -825
    var v5 int16 = v1
    v3 += (v3 * v3)
    _ = len(v2)
    v1 -= v5
    for fi7 := 0; fi7 < 16; fi7++ {
        for fi8 := 0; fi8 < 15; fi8++ {
            var v9 uint8 = 52
            v1 = v1
            _ = fi8
            _ = v9
        }
        if ((0 != 128) && (1.5 != float64(1.5))) {
            fmt.Println(v1)
            for fi10 := 0; fi10 < 2; fi10++ {
                var v11 int16 = v1
                _ = fi10
                _ = v11
            }
            var v12 int8 = 13
            _ = v12
        }
        var v13 uint16 = v3
        _ = fi7
        _ = v13
    }
    switch 0 {
    case -498:
        _ = len(v2)
    case 9223372036854775807:
        for fi14 := 0; fi14 < 16; fi14++ {
            var v15 int64 = -1
            for fi16 := 0; fi16 < 16; fi16++ {
                var v17 int16 = (v1 << ((922) & 63))
                var v18 int8 = 1
                switch fi16 {
                case 127:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                default:
                    var v19 int16 = 1000
                    _ = v19
                }
                _ = fi16
                _ = v17
                _ = v18
            }
            v2 = append(v2, 556)
            _ = fi14
            _ = v15
        }
    case 344:
        v4 = (v4 - v4)
    }
    v1 = v1
    v5 += 0
    switch 99 {
    case -143:
        var v20 uint64 = 239
        _ = v20
    case -246:
        var v21 int8 = -56
        _ = v21
    case 42:
        for fi22 := 0; fi22 < 11; fi22++ {
            v2 = append(v2, 93)
            var v23 int16 = -32768
            _ = fi22
            _ = v23
        }
    }
    v24 := GenS1{F0: 128}
    v5 = v1
    var v25 int32 = v4
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
