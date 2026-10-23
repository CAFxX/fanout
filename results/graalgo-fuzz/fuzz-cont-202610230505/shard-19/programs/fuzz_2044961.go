package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint8 = 0
    var v3 bool = true
    var v4 uint8 = 238
    v1[0] = 1
    v3 = (!(v2 >= v2))
    v4 += (v4 * (v4 << ((42) & 63)))
    if ((1e10 != ((((1e10 * 3.14159) / 1.0) + (3.14159 / (1.5 * -2.25))) / float64(1.5))) || (18446744073709551615 != 235)) {
        switch 128 {
        case 4294967295:
            v5 := v1[680]
            _ = v5
        case 32767:
            v6 := GenS1{F0: -39}
            _ = v6
        default:
            if v3 {
                for fi7 := 0; fi7 < 4; fi7++ {
                    fmt.Println(fi7, fi7, fi7)
                    if false {
                        v4 += (v2 << ((255) & 63))
                    } else {
                        v4 -= v2
                        for fi8 := 0; fi8 < 12; fi8++ {
                            var v9 int64 = -9223372036854775808
                            _ = fi8
                            _ = v9
                        }
                    }
                    _ = fi7
                }
            }
        }
        switch 0 {
        case 32767:
            for fi10 := 0; fi10 < 5; fi10++ {
                v11 := []int{(fi10 ^ 32767), fi10, 127}
                _ = fi10
                _ = v11
            }
        case 2147483647:
            var v12 string = "a"
            _ = v12
        case -9223372036854775808:
            _ = len(v1)
        default:
            v13 := v1[42]
            _ = v13
        }
    }
    switch ("foo bar" + ("\n\t" + ("foo bar" + ""))) {
    case "a":
        for fi14 := 0; fi14 < 2; fi14++ {
            var v15 uint8 = (v2 >> ((473) & 63))
            if (1 > 42) {
                fmt.Println(v2, v2)
            }
            _ = fi14
            _ = v15
        }
    case "hello":
        var v16 float64 = 1.5
        _ = v16
    case "\n\t":
        v17 := GenS1{F0: 1}
        _ = v17
    default:
        switch "\n\t" {
        case "a":
            var v18 float64 = 0.0
            _ = v18
        default:
            fmt.Println(v2, v4)
        }
    }
    var v19 uint = 42
    if v3 {
        v19 += 65535
        v4 = 42
        var v20 int64 = -1
        _ = v20
    } else {
        v4 = (^(v2))
        var v21 string = ("hello" + "a")
        var v22 int8 = -1
        _ = v21
        _ = v22
    }
    var v23 int32 = 2147483647
    if (false || (!v3)) {
        fmt.Println(v4, v19, v19)
        v4 -= v4
    } else {
        v2 = v4
        fmt.Println(v4, v4)
    }
    fmt.Println(v23, v3, v4)
    fmt.Println(v23, v2, v2)
    v23 -= v23
    var v24 uint16 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
