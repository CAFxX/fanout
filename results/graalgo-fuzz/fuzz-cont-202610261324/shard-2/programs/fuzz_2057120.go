package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 uint32; F1 int8; F2 int64 }
type GenS3 struct { F0 string; F1 int; F2 int16 }
func genFn0() uint64 {
    var v7 int16 = 1000
    var v8 float64 = ((3.14159 - 1e10) * 0.1)
    var v9 uint16 = 2
    _ = v7
    _ = v8
    _ = v9
    return 2
}
func genFn1(p0 int16, p1 uint64, p2 uint16) int {
    v10 := GenS1{F0: 2, F1: 0}
    var v11 int16 = p0
    var v12 string = "hello"
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    return -9223372036854775808
}

func main() {
    var v1 int64 = 0
    var v2 uint32 = 4294967295
    v3 := make(map[int]uint32)
    v4 := []string{}
    var v5 int32 = -1
    v6 := make(map[int]bool)
    if false {
        var v13 int64 = v1
        v14 := GenS1{F0: -1, F1: v2}
        v2 -= v2
        _ = v13
        _ = v14
    }
    v1 = v1
    switch int(v2) {
    case -2147483648:
        if len(v4) > 0 {
            v4[0] = ""
        }
    case 4294967295:
        v1 -= (v1 * -530)
    }
    for fi15 := 0; fi15 < 18; fi15++ {
        for fi16 := 0; fi16 < 19; fi16++ {
            var v17 int16 = 1000
            switch "a" {
            case "\n\t":
                v5 += (v5 | -1)
            }
            switch "" {
            case "\n\t":
                if false {
                    fmt.Println(fi15, fi16)
                    v5 = v5
                } else {
                    v1 = (v1 / (((v1 * (^(v1)))) | 1))
                }
            }
            switch (fi15 << ((661) & 63)) {
            case 32768:
                var v18 string = "foo bar"
                _ = v18
            case 0:
                var v19 int = int(v17)
                _ = v19
            case -828:
                if true {
                    fmt.Println(v2)
                    fmt.Println(v2, fi16, v5)
                }
                fmt.Println(fi15, v17)
            }
            switch (("a" + "") + "") {
            case "\n\t":
                v1 = (v1 + v1)
            case "a":
                v20 := GenS1{F0: 1, F1: v2}
                _ = v20
            }
            _ = fi16
            _ = v17
        }
        var v21 bool = (621 != (fi15 & fi15))
        v2 = (v2 << ((745) & 63))
        _ = fi15
        _ = v21
    }
    var v22 int32 = -769
    if ((!true) && false) {
        for fi23 := 0; fi23 < 16; fi23++ {
            var v24 float32 = 0.0
            _ = fi23
            _ = v24
        }
        v25 := GenS2{F0: (v2 + (v2 | (^(v2)))), F1: -1, F2: (v1 + (v1 | (v1 / ((v1) | 1))))}
        for fi26 := 0; fi26 < 12; fi26++ {
            v1 = v1
            v4 = append(v4, "a")
            _ = fi26
        }
        _ = v25
    }
    if (false && false) {
        var v27 uint = 165
        _ = v27
    } else {
        v28 := GenS3{F0: "\n\t", F1: -9223372036854775808, F2: 893}
        _ = v28
    }
    fmt.Println(v1, v5, v1)
    v5 += v5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
