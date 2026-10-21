package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int32 }

func main() {
    var v1 uint = 502
    v2 := make(map[uint32]bool)
    var v3 uint8 = 1
    fmt.Println(v3, v1)
    if (4294967295 > 2147483647) {
        var v4 uint64 = 587
        _ = v4
    }
    for fi5 := 0; fi5 < 13; fi5++ {
        v3 = v3
        _ = fi5
    }
    if ((-298 > 216) && false) {
        v1 -= (v1 / ((v1) | 1))
        v6 := GenS1{F0: 1}
        _ = v6
    } else {
        v7 := []int{-9223372036854775808, -2147483648, 917}
        v1 = (v1 * v1)
        _ = v7
    }
    v8 := GenS2{F0: 2, F1: (("hello" + "\n\t") + "foo bar")}
    if false {
        if (((-431 < -2147483648) || true) && (!false)) {
            v9 := GenS1{F0: 42}
            var v10 int16 = 42
            v11 := GenS3{F0: -1}
            _ = v10
            _ = v11
            _ = v9
        } else {
            v1 = 1
        }
        v12 := GenS2{F0: 906, F1: ("hello" + "foo bar")}
        v13 := []int{}
        _ = v12
        _ = v13
    } else {
        v3 -= 2
        var v14 bool = false
        _ = v14
    }
    fmt.Println(v1, v1)
    var v15 uint = (v1 ^ v1)
    if true {
        v16 := GenS4{F0: 103}
        if true {
            for fi17 := 0; fi17 < 1; fi17++ {
                v18 := []int{fi17, fi17, 2147483647, fi17}
                _ = fi17
                _ = v18
            }
        } else {
            var v19 int = 128
            var v20 string = "hello"
            _ = v19
            _ = v20
        }
        v21 := GenS1{F0: 42}
        _ = v16
        _ = v21
    }
    v22 := []int{0}
    var v23 string = "\n\t"
    if false {
        if true {
            var v24 uint16 = 905
            if true {
                v24 = 17
                v23 = "a"
            } else {
                v2[4294967295] = (!false)
            }
            _ = v24
        }
        v15 -= (v1 % ((v1) | 1))
        if len(v22) > 0 {
            v22[0] = 242
        }
    }
    if false {
        var v25 int8 = 19
        _ = v25
    } else {
        v3 = (v3 % ((v3) | 1))
        v23 = "hello"
    }
    for fi26 := 0; fi26 < 2; fi26++ {
        for fi27 := 0; fi27 < 9; fi27++ {
            v15 = v1
            _ = fi27
        }
        fmt.Println(v15, v23, v1)
        v15 -= (v1 | (v1 - v1))
        _ = fi26
    }
    v2[1] = false
    v1 += 4294967295
    v23 = ("\n\t" + v23)
    var v28 int = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v28
    _ = v3
    _ = v8
}
