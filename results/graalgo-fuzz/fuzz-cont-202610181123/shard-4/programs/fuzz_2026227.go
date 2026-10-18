package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 uint32 }
func genFn0(p0 string) uint {
    p0 = ("" + ("foo bar" + ""))
    fmt.Println(p0)
    _ = p0
    return 345
}
func genFn1() string {
    for fi4 := 0; fi4 < 12; fi4++ {
        v5 := GenS1{F0: 0, F1: "foo bar"}
        _ = fi4
        _ = v5
    }
    v6 := []int{-653, 9223372036854775807, 32767}
    var v7 uint64 = 42
    _ = v6
    _ = v7
    return "\n\t"
}

func main() {
    v1 := []int{}
    var v2 bool = (((false || false) && (0.0 == 1e10)) && true)
    var v3 uint32 = 476
    v8 := GenS1{F0: -9223372036854775808, F1: "foo bar"}
    v3 -= v3
    v3 += 4294967295
    if (42 > -862) {
        v2 = (true && false)
        v3 += (v3 - v3)
        if (-671 != 306) {
            var v9 int64 = 9223372036854775807
            for fi10 := 0; fi10 < 13; fi10++ {
                v9 -= v9
                _ = fi10
            }
            v11 := GenS1{F0: v9, F1: "foo bar"}
            _ = v11
            _ = v9
        } else {
            v2 = (v2 && (!(-354 != 1)))
            var v12 bool = (0 == -1)
            var v13 uint32 = (^(v3))
            _ = v12
            _ = v13
        }
    } else {
        v3 = (v3 - (v3 * 2147483647))
        v14 := GenS2{F0: 2147483648}
        _ = v14
    }
    for fi15 := 0; fi15 < 16; fi15++ {
        if len(v1) > 0 {
            v1[0] = fi15
        }
        _ = fi15
    }
    var v16 int = 127
    for fi17 := 0; fi17 < 4; fi17++ {
        for fi18 := 0; fi18 < 11; fi18++ {
            var v19 uint8 = 42
            var v20 bool = (2 > v19)
            switch 2147483647 {
            case 0:
                v20 = (v19 <= v19)
            }
            _ = fi18
            _ = v19
            _ = v20
        }
        v1 = append(v1, (v16 * v16))
        for fi21 := 0; fi21 < 6; fi21++ {
            var v22 string = "foo bar"
            var v23 int8 = 1
            v22 = v22
            v23 += 42
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = fi17
    }
    v3 -= (^(v3))
    var v24 uint32 = (v3 & 65535)
    switch "\n\t" {
    case "":
        switch ("" + "hello") {
        case "a":
            fmt.Println(v24, v3, v16)
        }
    case "hello":
        if (1 <= 183) {
            var v25 int32 = 121
            _ = v25
        }
    default:
        if false {
            var v26 uint8 = 211
            v1 = append(v1, v16)
            _ = v26
        }
    }
    v2 = (!true)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v2
    _ = v24
    _ = v3
    _ = v8
}
