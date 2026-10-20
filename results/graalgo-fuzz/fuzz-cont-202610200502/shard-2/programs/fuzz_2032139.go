package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32; F1 uint }

func main() {
    v1 := make(map[int]int)
    v2 := []int{42, 4294967295, 32767}
    var v3 uint16 = 256
    var v4 uint8 = 0
    var v5 uint16 = v3
    var v6 uint = 2
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := GenS1{F0: 560, F1: 965, F2: fi7}
        v5 = (v3 | v3)
        _ = fi7
        _ = v8
    }
    _ = len(v2)
    var v9 uint64 = uint64(v4)
    var v10 uint64 = v9
    for fi11 := 0; fi11 < 7; fi11++ {
        var v12 int16 = 0
        v6 -= (v6 + (^(v6)))
        v9 += 0
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 1; fi13++ {
        for fi14 := 0; fi14 < 5; fi14++ {
            v15 := v1[fi14]
            _ = fi14
            _ = v15
        }
        if false {
            for fi16 := 0; fi16 < 13; fi16++ {
                switch "hello" {
                case "hello":
                    v17 := GenS1{F0: 45, F1: -75, F2: 4294967295}
                    var v18 int8 = 48
                    _ = v17
                    _ = v18
                default:
                    var v19 int8 = 100
                    _ = v19
                }
                var v20 int8 = 42
                _ = fi16
                _ = v20
            }
            var v21 uint = v6
            fmt.Println(v6)
            _ = v21
        }
        var v22 uint = (v6 >> ((v6) & 63))
        _ = fi13
        _ = v22
    }
    var v23 float32 = 0.0
    v23 -= ((v23 * v23) / v23)
    v9 += 2
    delete(v1, -959)
    if (((v4 >> ((945) & 63)) < (v4 << ((18446744073709551615) & 63))) && false) {
        var v24 int = -9223372036854775808
        if false {
            for fi25 := 0; fi25 < 5; fi25++ {
                var v26 bool = (!true)
                _ = fi25
                _ = v26
            }
        }
        switch "" {
        case "":
            var v27 uint8 = 42
            _ = v27
        case "a":
            switch "" {
            case "hello":
                v28 := GenS2{F0: "\n\t"}
                _ = v28
            case "\n\t":
                v3 -= (v5 * (v5 - v3))
            default:
                if true {
                    v29 := GenS2{F0: ("a" + ("hello" + ""))}
                    var v30 uint8 = v4
                    _ = v29
                    _ = v30
                }
            }
        case "foo bar":
            v31 := GenS2{F0: "hello"}
            _ = v31
        }
        _ = v24
    } else {
        v1[2147483647] = -694
        var v32 int8 = 42
        delete(v1, -1)
        _ = v32
    }
    v33 := GenS3{F0: -205, F1: 2}
    v10 = (v9 | 673)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
