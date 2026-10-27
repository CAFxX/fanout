package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 int8 }
func genFn0() int16 {
    var v4 uint32 = 2147483647
    fmt.Println(v4)
    _ = v4
    return 807
}

func main() {
    v1 := []int{}
    var v2 uint32 = 42
    var v3 int8 = -76
    v5 := GenS1{F0: 1}
    v6 := GenS1{F0: 2}
    if false {
        v7 := GenS1{F0: 4294967295}
        for fi8 := 0; fi8 < 12; fi8++ {
            if false {
                v9 := GenS1{F0: 479}
                v10 := GenS1{F0: 921}
                _ = v10
                _ = v9
            } else {
                var v11 bool = ("foo bar" == ("" + "a"))
                _ = v11
            }
            v2 = (v2 & v2)
            _ = fi8
        }
        _ = v7
    } else {
        switch "hello" {
        case "hello":
            v2 -= v2
        case "\n\t":
            v12 := GenS1{F0: 616}
            _ = v12
        case "":
            var v13 float64 = 1.5
            _ = v13
        }
        switch ("hello" + ("hello" + "hello")) {
        case "":
            var v14 uint32 = (v2 & v2)
            _ = v14
        default:
            for fi15 := 0; fi15 < 12; fi15++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v16 uint64 = 42
                if (v16 >= (^(v16))) {
                    var v17 int32 = -2147483648
                    _ = v17
                } else {
                    v18 := GenS2{F0: 18446744073709551615}
                    v3 = (^(v3))
                    var v19 uint64 = (v16 % (((^(v16))) | 1))
                    _ = v18
                    _ = v19
                }
                _ = fi15
                _ = v16
            }
        }
    }
    v20 := GenS3{F0: 550, F1: v3}
    var v21 float32 = 1.5
    if false {
        v2 += 682
        if (((0 <= 234) && false) || false) {
            var v22 uint = 255
            var v23 int = 1
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v22
            _ = v23
        }
        fmt.Println(v21)
    } else {
        v3 = v3
    }
    v2 = v2
    v24 := GenS2{F0: 42}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v21 -= (v21 - (1e10 - v21))
    fmt.Println(v2, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v3
    _ = v5
    _ = v6
}
