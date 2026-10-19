package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 bool }

func main() {
    var v1 string = "\n\t"
    var v2 int64 = -1
    var v3 uint8 = 128
    var v4 uint8 = (v3 % (((v3 << ((65535) & 63))) | 1))
    var v5 uint16 = 65535
    v6 := []uint32{65535, 65535, 965}
    if false {
        if true {
            var v7 string = v1
            _ = v7
        }
        if (v2 >= int64(v5)) {
            v5 = v5
            var v8 uint16 = (v5 >> ((0) & 63))
            v4 = (v3 << ((962) & 63))
            _ = v8
        }
    }
    v6 = append(v6, 2147483647)
    v2 = (v2 * 9223372036854775807)
    for fi9 := 0; fi9 < 3; fi9++ {
        v10 := GenS1{F0: 798, F1: (1 != 1), F2: false}
        var v11 uint32 = 244
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 uint32 = 767
    v1 = "\n\t"
    v2 = (v2 + int64(v5))
    fmt.Println(v12)
    v1 = v1
    fmt.Println(v3)
    v13 := GenS1{F0: v12, F1: false, F2: true}
    v12 = (v12 - (v12 >> ((58) & 63)))
    fmt.Println(v3)
    for fi14 := 0; fi14 < 7; fi14++ {
        v1 = v1
        if (198 >= 18446744073709551615) {
            var v15 bool = true
            switch (fi14 + fi14) {
            case -2147483648:
                v16 := GenS1{F0: 2147483648, F1: (v1 > v1), F2: false}
                _ = v16
            }
            for fi17 := 0; fi17 < 19; fi17++ {
                var v18 int64 = (v2 - v2)
                v6 = append(v6, (v12 & v12))
                v4 = (v4 * v4)
                _ = fi17
                _ = v18
            }
            _ = v15
        } else {
            var v19 float64 = ((-2.25 + 0.0) + 1e10)
            var v20 string = ((((v1 + v1) + "foo bar") + "foo bar") + (v1 + v1))
            fmt.Println(v19, fi14, v20)
            _ = v19
            _ = v20
        }
        _ = fi14
    }
    var v21 float64 = ((0.1 - (-2.25 * 3.14159)) * 1e10)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
