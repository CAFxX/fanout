package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 int64; F2 int64 }
type GenS3 struct { F0 int64 }

func main() {
    var v1 float64 = 0.0
    var v2 bool = false
    var v3 int16 = -260
    var v4 uint8 = 128
    v5 := []string{"", (("" + "\n\t") + "foo bar"), ("hello" + ("a" + ("\n\t" + ("hello" + "\n\t"))))}
    var v6 uint8 = (v4 / ((42) | 1))
    v4 = 39
    var v7 int64 = -221
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 uint16 = 178
        switch "" {
        case "":
            v4 -= 128
        }
        if v2 {
            switch (-(fi8)) {
            case -598:
                var v10 bool = false
                _ = v10
            case 4294967295:
                fmt.Println(v6, v4)
            default:
                v11 := GenS1{F0: (fi8 << ((18446744073709551615) & 63))}
                _ = v11
            }
        } else {
            v4 = (v4 - (v6 + (v4 + 255)))
            var v12 uint64 = 4294967295
            switch ("hello" + "") {
            case "a":
                if false {
                    v12 = v12
                    fmt.Println(v4)
                } else {
                    var v13 uint8 = 128
                    _ = v13
                }
            case "foo bar":
                var v14 int8 = 114
                _ = v14
            case "hello":
                var v15 int8 = 20
                _ = v15
            default:
                for fi16 := 0; fi16 < 10; fi16++ {
                    var v17 float64 = v1
                    if len(v5) > 0 {
                        v5[0] = "foo bar"
                    }
                    _ = fi16
                    _ = v17
                }
            }
            _ = v12
        }
        _ = fi8
        _ = v9
    }
    _ = len(v5)
    v6 = (^(v4))
    v18 := GenS2{F0: -1, F1: -205, F2: (v7 / (((v7 ^ (-(v7)))) | 1))}
    v19 := GenS2{F0: -492, F1: v7, F2: (v7 << ((1) & 63))}
    var v20 uint32 = 2
    v7 -= (v7 - v7)
    _ = len(v5)
    v2 = (255 > 4294967295)
    var v21 uint8 = v4
    for fi22 := 0; fi22 < 16; fi22++ {
        for fi23 := 0; fi23 < 17; fi23++ {
            if len(v5) > 0 {
                v5[0] = "hello"
            }
            var v24 int = (fi23 - 83)
            v25 := GenS3{F0: v7}
            _ = fi23
            _ = v24
            _ = v25
        }
        var v26 uint16 = 5
        _ = len(v5)
        _ = fi22
        _ = v26
    }
    var v27 int32 = 56
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
