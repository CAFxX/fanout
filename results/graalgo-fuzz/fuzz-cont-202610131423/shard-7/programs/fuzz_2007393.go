package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 uint32; F1 uint64; F2 int }
type GenS3 struct { F0 int16 }

func main() {
    var v1 float32 = 1e10
    var v2 string = ("foo bar" + "")
    var v3 uint64 = 18446744073709551615
    var v4 int64 = 9223372036854775807
    var v5 int8 = -1
    var v6 int16 = 363
    v1 -= v1
    v3 = (v3 ^ v3)
    if false {
        v2 = "\n\t"
        var v7 bool = ((v3 * v3) <= (v3 + v3))
        if v7 {
            v8 := []int{0}
            _ = v8
        }
        _ = v7
    }
    if true {
        v2 = (v2 + v2)
        v3 -= v3
        var v9 int64 = v4
        _ = v9
    } else {
        if true {
            var v10 int64 = v4
            v1 = v1
            v4 = 9223372036854775807
            _ = v10
        }
        var v11 uint16 = 0
        _ = v11
    }
    var v12 uint32 = 747
    switch int(v4) {
    case 0:
        v13 := GenS1{F0: -371, F1: 32768}
        _ = v13
    case 2147483647:
        v14 := GenS2{F0: 2147483648, F1: (v3 ^ v3), F2: 235}
        _ = v14
    case 311:
        v2 = ""
    default:
        v6 = 0
    }
    fmt.Println(v12)
    if (v3 > v3) {
        switch v2 {
        case "hello":
            var v15 int = 941
            _ = v15
        case "\n\t":
            if true {
                v6 += v6
                switch -9223372036854775808 {
                case 334:
                    var v16 int16 = 0
                    _ = v16
                default:
                    v5 += v5
                }
            } else {
                for fi17 := 0; fi17 < 8; fi17++ {
                    if (((25 >= 65535) && (366 >= 271)) && false) {
                        v12 = (v12 % ((0) | 1))
                        v12 = (v12 | v12)
                        switch (((((v2 + v2) + "\n\t") + "") + ((("" + (v2 + v2)) + v2) + v2)) + "") {
                        case "foo bar":
                            for fi18 := 0; fi18 < 7; fi18++ {
                                v5 = v5
                                _ = fi18
                            }
                        }
                    }
                    v12 = 833
                    var v19 uint64 = (v3 / (((v3 & v3)) | 1))
                    _ = fi17
                    _ = v19
                }
                var v20 uint8 = 1
                _ = v20
            }
        default:
            var v21 uint = 65535
            _ = v21
        }
    } else {
        var v22 bool = (false || (v1 >= 0.0))
        _ = v22
    }
    for fi23 := 0; fi23 < 16; fi23++ {
        v6 = (v6 ^ v6)
        _ = fi23
    }
    v24 := []int{683, -1}
    fmt.Println(v1, v6, v6)
    v24 = append(v24, 2147483647)
    if (false && false) {
        switch (v2 + v2) {
        case "a":
            v25 := GenS3{F0: 32767}
            _ = v25
        default:
            if len(v24) > 0 {
                v24[0] = 2147483647
            }
        }
        _ = len(v24)
    } else {
        if false {
            fmt.Println(v4)
            var v26 bool = true
            fmt.Println(v26)
            _ = v26
        }
    }
    v1 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
