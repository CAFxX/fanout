package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 uint16; F2 bool }

func main() {
    var v1 uint = 2
    v2 := []uint32{2147483647}
    var v3 int = -340
    var v4 int32 = 2147483647
    var v5 float64 = 0.1
    var v6 int64 = 9223372036854775807
    for fi7 := 0; fi7 < 20; fi7++ {
        v6 = v6
        v6 = 9223372036854775807
        v3 = fi7
        _ = fi7
    }
    v1 -= (v1 - v1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch ("\n\t" + "") {
    case "foo bar":
        var v8 uint32 = 65535
        _ = v8
    }
    var v9 int32 = (v4 | (v4 % (((v4 * (^(v4)))) | 1)))
    var v10 int = -302
    if (!false) {
        var v11 int32 = (v9 | v4)
        _ = v11
    } else {
        switch "hello" {
        case "hello":
            var v12 uint8 = 127
            _ = v12
        case "\n\t":
            v6 = v6
        default:
            var v13 float64 = v5
            _ = v13
        }
        v14 := GenS1{F0: 1, F1: -32768}
        v4 -= (v9 / ((int32(v10)) | 1))
        _ = v14
    }
    var v15 float32 = (float32((3.14159 - v5)) / 1.0)
    v3 -= int(v1)
    v4 = v4
    var v16 uint32 = 65535
    for fi17 := 0; fi17 < 12; fi17++ {
        fmt.Println(v6, v3)
        fmt.Println(v3, v16, v6)
        _ = fi17
    }
    var v18 float64 = (v5 / float64(((v15 * v15) * (v15 - v15))))
    v5 -= v5
    if (((((v5 - v18) + (0.1 * (float64(0.1) / v5))) + 1e10) + v5) <= v5) {
        for fi19 := 0; fi19 < 20; fi19++ {
            if (true || (4294967295 == 2)) {
                var v20 int = 127
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                if len(v2) > 0 {
                    v2[0] = v16
                }
                _ = v20
            } else {
                for fi21 := 0; fi21 < 16; fi21++ {
                    var v22 uint = 65535
                    _ = fi21
                    _ = v22
                }
                var v23 int = -2147483648
                switch v10 {
                case 371:
                    v24 := GenS2{F0: int8(v16)}
                    _ = v24
                }
                _ = v23
            }
            for fi25 := 0; fi25 < 12; fi25++ {
                v2 = append(v2, v16)
                _ = fi25
            }
            fmt.Println(v4)
            _ = fi19
        }
        v26 := GenS3{F0: 107, F1: 0, F2: false}
        if true {
            if (18446744073709551615 != 962) {
                fmt.Println(v9)
                v1 += (v1 % ((v1) | 1))
            }
            for fi27 := 0; fi27 < 7; fi27++ {
                v1 = 0
                v28 := GenS2{F0: 0}
                _ = fi27
                _ = v28
            }
        } else {
            _ = len(v2)
        }
        _ = v26
    } else {
        var v29 float32 = v15
        _ = len(v2)
        var v30 int8 = 1
        _ = v29
        _ = v30
    }
    v6 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
