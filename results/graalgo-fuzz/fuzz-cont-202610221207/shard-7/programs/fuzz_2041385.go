package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 uint32; F1 int; F2 int8 }
type GenS3 struct { F0 uint16; F1 uint16; F2 uint8 }

func main() {
    var v1 uint8 = 128
    v2 := []string{"hello"}
    var v3 uint32 = 0
    v4 := GenS1{F0: -2147483648, F1: 127}
    for fi5 := 0; fi5 < 5; fi5++ {
        if (((0.0 >= float64(float32(1.5))) && ((32767 != 32767) || (false && (!false)))) && (-2.25 == (0.1 * 0.0))) {
            _ = len(v2)
            v1 = v1
        } else {
            for fi6 := 0; fi6 < 2; fi6++ {
                v1 -= v1
                if ((fi6 ^ (^(fi6))) < (fi6 * fi6)) {
                    if (true && false) {
                        v7 := GenS2{F0: (v3 << ((0) & 63)), F1: fi5, F2: 0}
                        v2 = append(v2, "foo bar")
                        _ = v7
                    }
                    v3 += (v3 - 2)
                    switch "a" {
                    case "\n\t":
                        var v8 uint8 = (v1 % ((1) | 1))
                        v8 += v8
                        _ = v8
                    }
                }
                _ = fi6
            }
            var v9 float64 = 1.5
            _ = v9
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v10 int64 = 1
        _ = fi5
        _ = v10
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        var v12 int16 = -1
        v13 := GenS1{F0: -1, F1: (^(v1))}
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 uint8 = v1
    v15 := GenS2{F0: v3, F1: 42, F2: 0}
    var v16 int32 = 42
    v2 = append(v2, ("a" + "hello"))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (v3 > (v3 - (^(v3)))) {
        fmt.Println(v14, v16, v14)
    } else {
        switch ("foo bar" + "foo bar") {
        case "hello":
            var v17 int8 = -128
            _ = v17
        case "foo bar":
            if true {
                fmt.Println(v14, v1)
            } else {
                var v18 string = "a"
                v19 := GenS3{F0: 65535, F1: 1, F2: v14}
                _ = v18
                _ = v19
            }
        }
    }
    fmt.Println(v14, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
