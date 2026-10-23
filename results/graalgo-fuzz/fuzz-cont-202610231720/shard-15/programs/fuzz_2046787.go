package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 uint; F1 bool }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int; F1 int64 }

func main() {
    v1 := make(map[string]string)
    var v2 int16 = 42
    v3 := make(map[string]bool)
    var v4 int64 = 9223372036854775807
    var v5 int16 = v2
    var v6 int64 = 1
    v7 := GenS1{F0: 141, F1: 4294967295}
    fmt.Println(v4, v2, v4)
    if ((42 < 255) || (true || (-313 > -1))) {
        var v8 int8 = -128
        for fi9 := 0; fi9 < 2; fi9++ {
            if false {
                var v10 uint16 = 0
                _ = v10
            } else {
                v11 := GenS2{F0: v8, F1: "foo bar"}
                v12 := GenS3{F0: 1, F1: false}
                v13 := GenS4{F0: v6}
                _ = v11
                _ = v12
                _ = v13
            }
            _ = fi9
        }
        for fi14 := 0; fi14 < 4; fi14++ {
            switch "" {
            case "\n\t":
                if ((false || false) || (!(!(!false)))) {
                    v8 -= v8
                }
            default:
                var v15 uint64 = 1
                _ = v15
            }
            fmt.Println(v8, v4, v4)
            if (42 < -1) {
                for fi16 := 0; fi16 < 17; fi16++ {
                    v2 = int16(v8)
                    fmt.Println(v5)
                    v5 = v2
                    _ = fi16
                }
                fmt.Println(v4)
                var v17 uint32 = 2147483647
                _ = v17
            }
            _ = fi14
        }
        _ = v8
    }
    v18 := GenS5{F0: 9223372036854775807, F1: (^(v4))}
    v6 -= 983
    fmt.Println(v4, v2, v4)
    if true {
        for fi19 := 0; fi19 < 15; fi19++ {
            var v20 uint = 549
            _ = fi19
            _ = v20
        }
    }
    v21 := GenS5{F0: 32767, F1: v6}
    v6 = v6
    var v22 uint8 = 130
    var v23 uint16 = 255
    v24 := GenS4{F0: (^(v4))}
    fmt.Println(v22, v5, v22)
    var v25 int64 = v4
    v2 -= v5
    var v26 uint8 = 127
    for fi27 := 0; fi27 < 18; fi27++ {
        var v28 uint32 = 5
        v22 += (v26 / ((v22) | 1))
        _ = fi27
        _ = v28
    }
    var v29 uint32 = 2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
