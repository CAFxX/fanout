package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64 }
func genFn0(p0 string, p1 bool) int64 {
    var v7 int16 = 1000
    var v8 uint64 = 18446744073709551615
    v8 -= (v8 + (v8 + (v8 >> ((42) & 63))))
    for fi9 := 0; fi9 < 20; fi9++ {
        v8 = v8
        fmt.Println(p1)
        _ = fi9
    }
    v10 := GenS1{F0: 1}
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    return 1
}

func main() {
    v1 := []string{}
    v2 := make(map[uint32]int)
    var v3 uint16 = 65535
    var v4 int = 128
    var v5 int8 = 1
    var v6 uint8 = 1
    var v11 int8 = (v5 - v5)
    delete(v2, 65535)
    if false {
        v4 = 716
    } else {
        for fi12 := 0; fi12 < 17; fi12++ {
            var v13 int8 = (v5 | v5)
            var v14 uint64 = 0
            _ = fi12
            _ = v13
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 2; fi15++ {
        switch (^(fi15)) {
        case 9223372036854775807:
            var v16 int16 = -105
            _ = v16
        case 966:
            var v17 int = v4
            _ = v17
        }
        _ = fi15
    }
    v18 := GenS1{F0: v5}
    if (!true) {
        var v19 uint8 = (v6 << ((18446744073709551615) & 63))
        switch v4 {
        case 32768:
            v4 = v4
        case 330:
            if (!(!false)) {
                for fi20 := 0; fi20 < 14; fi20++ {
                    v19 = v6
                    _ = fi20
                }
                switch "\n\t" {
                case "a":
                    v21 := GenS1{F0: (v5 & -128)}
                    _ = v21
                case "\n\t":
                    var v22 int16 = 1000
                    _ = v22
                default:
                    for fi23 := 0; fi23 < 12; fi23++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        switch "\n\t" {
                        case "":
                            v11 -= v5
                        case "hello":
                            v1 = append(v1, ("\n\t" + ("hello" + "")))
                        }
                        _ = fi23
                    }
                }
                fmt.Println(v19, v19, v6)
            } else {
                v24 := GenS2{F0: -9223372036854775808}
                for fi25 := 0; fi25 < 18; fi25++ {
                    v11 = v5
                    for fi26 := 0; fi26 < 12; fi26++ {
                        fmt.Println(v4, v4, v3)
                        _ = fi26
                    }
                    _ = fi25
                }
                var v27 uint = 820
                _ = v24
                _ = v27
            }
        case -2147483648:
            v28 := GenS2{F0: 747}
            _ = v28
        default:
            v2[2147483647] = (v4 / ((v4) | 1))
        }
        _ = v19
    } else {
        v6 -= 230
        fmt.Println(v5, v4)
        switch (v4 >> ((255) & 63)) {
        case -351:
            var v29 int64 = -12
            _ = v29
        case 32768:
            for fi30 := 0; fi30 < 1; fi30++ {
                fmt.Println(v3, v11)
                for fi31 := 0; fi31 < 11; fi31++ {
                    _ = len(v2)
                    _ = fi31
                }
                var v32 int = v4
                _ = fi30
                _ = v32
            }
        case 568:
            v5 -= v11
        }
    }
    var v33 uint = 394
    v1 = append(v1, "")
    v5 = v5
    var v34 string = "\n\t"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v18
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
