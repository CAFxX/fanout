package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 int16 }
type GenS2 struct { F0 uint32; F1 int; F2 int8 }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint64) uint16 {
    if false {
        v7 := []int{9223372036854775807, -2147483648}
        var v8 uint = 42
        _ = v7
        _ = v8
    }
    var v9 bool = (true || (!false))
    var v10 uint = uint(p0)
    fmt.Println(p0)
    _ = p0
    _ = v10
    _ = v9
    return 551
}

func main() {
    v1 := []uint32{2, 2, 4294967295}
    var v2 uint64 = 0
    var v3 bool = (1 > -128)
    v4 := []int{-9223372036854775808, -545}
    v5 := []int{0, 0, 1, -9223372036854775808}
    var v6 int = -372
    v6 += (v6 ^ (v6 << ((4294967295) & 63)))
    var v11 int16 = 42
    for fi12 := 0; fi12 < 5; fi12++ {
        v3 = v3
        var v13 int16 = (v11 * -1)
        _ = fi12
        _ = v13
    }
    v14 := GenS1{F0: (false || v3), F1: (v6 & (^(v6))), F2: v11}
    switch (("\n\t" + "foo bar") + "") {
    case "a":
        for fi15 := 0; fi15 < 6; fi15++ {
            for fi16 := 0; fi16 < 1; fi16++ {
                v2 = 1
                v1 = append(v1, 2)
                _ = fi16
            }
            var v17 int8 = 42
            v18 := GenS1{F0: v3, F1: fi15, F2: v11}
            _ = fi15
            _ = v17
            _ = v18
        }
    }
    v4 = append(v4, (^(v6)))
    var v19 int16 = v11
    if v3 {
        v2 -= (v2 / ((v2) | 1))
        switch "" {
        case "foo bar":
            _ = len(v1)
        case "\n\t":
            if true {
                fmt.Println(v19, v19, v3)
                v20 := GenS1{F0: v3, F1: 32768, F2: v11}
                _ = v20
            } else {
                if (42 <= (v2 % (((v2 ^ 2)) | 1))) {
                    var v21 float32 = (1.5 / 1.0)
                    _ = v21
                } else {
                    var v22 string = "foo bar"
                    _ = v22
                }
                v3 = v3
            }
        case "a":
            v19 = int16(v2)
        default:
            v3 = false
        }
    } else {
        for fi23 := 0; fi23 < 15; fi23++ {
            for fi24 := 0; fi24 < 13; fi24++ {
                v2 = v2
                _ = fi24
            }
            v6 = v6
            var v25 int64 = 1
            _ = fi23
            _ = v25
        }
        var v26 int8 = 127
        switch "foo bar" {
        case "":
            v19 += v19
        case "hello":
            v27 := GenS2{F0: 323, F1: v6, F2: v26}
            _ = v27
        case "a":
            for fi28 := 0; fi28 < 5; fi28++ {
                v29 := GenS3{F0: 243}
                _ = fi28
                _ = v29
            }
        default:
            var v30 int64 = -9223372036854775808
            _ = v30
        }
        _ = v26
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
