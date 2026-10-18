package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 int16; F1 int16; F2 uint8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16; F1 uint64 }

func main() {
    var v1 bool = (!true)
    var v2 int8 = 100
    v3 := make(map[int]uint32)
    v4 := GenS1{F0: 255, F1: 2}
    var v5 int32 = 0
    fmt.Println(v2)
    v6 := GenS1{F0: 203, F1: 261}
    v1 = (!((810 <= 659) && (v2 > v2)))
    switch "" {
    case "\n\t":
        fmt.Println(v5)
    }
    if (263 != 4294967295) {
        var v7 uint8 = 250
        _ = v7
    } else {
        switch -947 {
        case -9223372036854775808:
            v8 := GenS1{F0: 0, F1: 2}
            _ = v8
        }
        for fi9 := 0; fi9 < 16; fi9++ {
            v1 = v1
            v10 := []int{}
            switch ("" + "foo bar") {
            case "hello":
                for fi11 := 0; fi11 < 7; fi11++ {
                    v12 := GenS1{F0: 255, F1: 257}
                    _ = fi11
                    _ = v12
                }
            default:
                for fi13 := 0; fi13 < 5; fi13++ {
                    if ((fi9 | (fi13 & fi9)) != 127) {
                        var v14 int64 = 0
                        v2 = (v2 << ((2) & 63))
                        v2 -= v2
                        _ = v14
                    }
                    _ = fi13
                }
            }
            _ = fi9
            _ = v10
        }
        for fi15 := 0; fi15 < 16; fi15++ {
            v16 := GenS2{F0: -1, F1: 1000, F2: 2}
            v17 := GenS3{F0: 65535}
            _ = fi15
            _ = v16
            _ = v17
        }
    }
    if (!v1) {
        fmt.Println(v2, v5)
    } else {
        var v18 float64 = float64(1e10)
        _ = v18
    }
    switch (("hello" + "foo bar") + (("" + "\n\t") + "\n\t")) {
    case "":
        switch "\n\t" {
        case "":
            v1 = (v1 || (v1 || ((-2147483648 > -9223372036854775808) && ((943 < 2) || v1))))
        case "a":
            v19 := []int{32767, 42, -104, 32767}
            _ = v19
        default:
            for fi20 := 0; fi20 < 6; fi20++ {
                var v21 uint32 = 36
                _ = fi20
                _ = v21
            }
        }
    case "hello":
        v22 := GenS1{F0: 255, F1: 32767}
        _ = v22
    case "\n\t":
        delete(v3, -1)
    default:
        if (((v2 - v2) <= (^(v2))) && true) {
            v23 := GenS4{F0: -645, F1: 4294967295}
            for fi24 := 0; fi24 < 10; fi24++ {
                var v25 int16 = 42
                var v26 int8 = v2
                fmt.Println(v25, fi24)
                _ = fi24
                _ = v25
                _ = v26
            }
            _ = v23
        } else {
            var v27 string = ""
            v28 := v3[128]
            v1 = v1
            _ = v27
            _ = v28
        }
    }
    v29 := v3[2147483647]
    var v30 int8 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
