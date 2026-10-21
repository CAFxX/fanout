package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint8 }

func main() {
    v1 := []int{128}
    var v2 float32 = 0.1
    v3 := []uint32{482, 200, 701, 0}
    v2 -= v2
    if ((4294967295 != 0) || (!(127 <= 42))) {
        for fi4 := 0; fi4 < 6; fi4++ {
            _ = len(v3)
            v3 = append(v3, 2)
            _ = fi4
        }
    }
    v2 -= v2
    if (1 >= 65535) {
        var v5 int32 = 42
        var v6 string = ("\n\t" + "hello")
        if len(v3) > 0 {
            v3[0] = 65535
        }
        _ = v5
        _ = v6
    }
    fmt.Println(v2)
    v2 = -2.25
    v2 -= (3.14159 / 1.5)
    switch 1 {
    case 127:
        for fi7 := 0; fi7 < 1; fi7++ {
            if (42 < 71) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                fmt.Println(v2)
            } else {
                fmt.Println(v2, fi7)
            }
            _ = fi7
        }
    case -224:
        v2 = (v2 - ((3.14159 * -2.25) - (v2 / v2)))
    case 813:
        for fi8 := 0; fi8 < 1; fi8++ {
            if false {
                fmt.Println(fi8, v2)
                for fi9 := 0; fi9 < 19; fi9++ {
                    v10 := GenS1{F0: -1, F1: 273, F2: 127}
                    fmt.Println(fi8, fi8)
                    v2 = (float32(-2.25) + (v2 - -2.25))
                    v2 += float32(1e10)
                    var v11 int = -9223372036854775808
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                var v12 string = ""
                _ = v12
            } else {
                v2 = (v2 + v2)
            }
            var v13 int8 = 59
            _ = fi8
            _ = v13
        }
    }
    if true {
        v14 := GenS1{F0: -1, F1: 698, F2: 241}
        for fi15 := 0; fi15 < 14; fi15++ {
            var v16 uint16 = 630
            v16 = v16
            var v17 int64 = -1
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = len(v1)
        _ = v14
    } else {
        v2 -= (((0.1 + v2) * (0.1 + v2)) / float32(3.14159))
    }
    var v18 uint = 541
    v18 -= v18
    v18 -= (v18 + (v18 << ((42) & 63)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
}
