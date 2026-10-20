package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint16; F1 int8 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint = 1
    v3 := make(map[uint32]bool)
    var v4 int64 = -9223372036854775808
    if false {
        v5 := []int{-9223372036854775808, 4294967295, 4294967295, -19}
        _ = v5
    } else {
        if false {
            var v6 string = "\n\t"
            var v7 bool = false
            v6 = "foo bar"
            _ = v6
            _ = v7
        }
        v8 := GenS1{F0: (v2 >> ((255) & 63)), F1: -1}
        if (947 == 0) {
            switch ("" + ("hello" + (((("" + "hello") + "foo bar") + "\n\t") + "a"))) {
            case "":
                if ((1.5 - 1e10) <= 0.0) {
                    if (!true) {
                        v9 := v1[4294967295]
                        v10 := GenS2{F0: v4, F1: 2}
                        for fi11 := 0; fi11 < 9; fi11++ {
                            v12 := []int{-2147483648, (-(fi11)), fi11, (-(fi11))}
                            _ = fi11
                            _ = v12
                        }
                        _ = v10
                        _ = v9
                    } else {
                        v13 := GenS3{F0: 291}
                        v2 = 110
                        v14 := GenS4{F0: 2, F1: 127}
                        _ = v13
                        _ = v14
                    }
                }
            }
            for fi15 := 0; fi15 < 13; fi15++ {
                if (!(4294967295 > 42)) {
                    for fi16 := 0; fi16 < 11; fi16++ {
                        fmt.Println(fi15)
                        v1[0] = (42 > 645)
                        v1[42] = (-1 <= 42)
                        _ = fi16
                    }
                }
                if (1e10 == 0.0) {
                    var v17 int8 = -105
                    _ = v17
                }
                v2 = (v2 - 42)
                _ = fi15
            }
            v2 += 785
        }
        _ = v8
    }
    v18 := GenS2{F0: v4, F1: 2147483648}
    for fi19 := 0; fi19 < 18; fi19++ {
        v20 := GenS3{F0: (fi19 / ((127) | 1))}
        if true {
            switch (fi19 >> ((v2) & 63)) {
            case 2147483647:
                v4 -= (^(v4))
            default:
                fmt.Println(v4, fi19, v2)
            }
        } else {
            var v21 int = (fi19 & (fi19 >> (((v2 + (v2 ^ v2))) & 63)))
            _ = v21
        }
        var v22 uint = 0
        _ = fi19
        _ = v20
        _ = v22
    }
    v2 = (v2 >> (((^(v2))) & 63))
    v23 := v1[422]
    v24 := GenS1{F0: (v2 * (^(v2))), F1: -2147483648}
    var v25 int = 837
    fmt.Println(v25, v25)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
