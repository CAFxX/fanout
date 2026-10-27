package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 int }
type GenS2 struct { F0 int16; F1 uint16; F2 int }

func main() {
    var v1 int8 = -128
    var v2 uint = 4294967295
    v3 := []string{}
    v4 := []uint32{1}
    var v5 float32 = 0.0
    for fi6 := 0; fi6 < 8; fi6++ {
        if (417 == v2) {
            v4 = append(v4, 2147483648)
            fmt.Println(v2, v5, v5)
        } else {
            if (271 >= (^(fi6))) {
                v5 = 0.1
                v1 = v1
            }
            switch "" {
            case "a":
                switch ("" + "") {
                case "hello":
                    fmt.Println(v5)
                default:
                    v2 += (v2 >> (((v2 % (((v2 % (((v2 - 18446744073709551615)) | 1))) | 1))) & 63))
                }
            case "hello":
                v1 = (v1 ^ (v1 + 100))
            }
        }
        _ = fi6
    }
    if false {
        v1 = 42
        var v7 uint8 = 42
        for fi8 := 0; fi8 < 18; fi8++ {
            var v9 uint16 = 42
            _ = fi8
            _ = v9
        }
        _ = v7
    }
    var v10 uint64 = 1
    switch 128 {
    case 128:
        v11 := GenS1{F0: 0, F1: 54, F2: 127}
        _ = v11
    }
    var v12 string = "foo bar"
    var v13 bool = (v12 != "")
    v13 = true
    fmt.Println(v12, v13, v2)
    var v14 float32 = v5
    v3 = append(v3, "foo bar")
    if (v1 == (v1 & 0)) {
        for fi15 := 0; fi15 < 2; fi15++ {
            v16 := GenS2{F0: 1, F1: 0, F2: (fi15 + (fi15 * (fi15 + (^(fi15)))))}
            v13 = v13
            _ = fi15
            _ = v16
        }
    }
    for fi17 := 0; fi17 < 1; fi17++ {
        for fi18 := 0; fi18 < 12; fi18++ {
            v19 := GenS2{F0: -32768, F1: 0, F2: 85}
            v13 = v13
            v4 = append(v4, 704)
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v20 uint = v2
    v13 = (0 > v1)
    var v21 float64 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
