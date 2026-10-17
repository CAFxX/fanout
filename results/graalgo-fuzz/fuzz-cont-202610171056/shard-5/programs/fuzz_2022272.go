package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 uint8 }
func genFn0(p0 int64, p1 int) int16 {
    v4 := []int{(p1 + (p1 | p1)), 128}
    var v5 int8 = 100
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 0
}
func genFn1(p0 int8, p1 uint16) int32 {
    var v6 int32 = -2147483648
    p0 += 100
    _ = p0
    _ = p1
    _ = v6
    return v6
}

func main() {
    var v1 string = "\n\t"
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]int)
    v7 := GenS1{F0: 0, F1: -902, F2: 2}
    fmt.Println(v1)
    switch ("\n\t" + (v1 + (v1 + v1))) {
    case "":
        fmt.Println(v1)
    case "a":
        var v8 uint8 = 12
        _ = v8
    }
    v9 := GenS1{F0: 42, F1: 42, F2: 497}
    _ = len(v3)
    v1 = v1
    v10 := []int{42, 99}
    fmt.Println(v1)
    var v11 int8 = -128
    fmt.Println(v1, v1)
    switch -478 {
    case 1:
        v12 := GenS2{F0: false}
        _ = v12
    default:
        for fi13 := 0; fi13 < 16; fi13++ {
            v1 = v1
            v14 := GenS1{F0: 65535, F1: 764, F2: 465}
            _ = fi13
            _ = v14
        }
    }
    if (117 <= 42) {
        v15 := v3[4294967295]
        if (!((4294967295 >= 2) && true)) {
            if (65535 > 65535) {
                _ = len(v2)
            }
        }
        switch "" {
        case "":
            var v16 int = -9223372036854775808
            _ = v16
        case "hello":
            var v17 uint = 680
            _ = v17
        case "\n\t":
            var v18 float32 = ((0.1 - 1.5) - 0.1)
            _ = v18
        default:
            if true {
                fmt.Println(v1, v11)
                if false {
                    v19 := GenS2{F0: false}
                    if true {
                        v10 = append(v10, (v15 | v15))
                    } else {
                        var v20 int16 = 0
                        var v21 int32 = 42
                        fmt.Println(v11)
                        _ = v20
                        _ = v21
                    }
                    _ = v19
                } else {
                    v22 := GenS3{F0: (^(v15)), F1: 0}
                    v1 = ("" + "a")
                    _ = v22
                }
                fmt.Println(v11)
            } else {
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
            }
        }
        _ = v15
    }
    var v23 uint32 = 0
    var v24 int8 = -69
    switch (v1 + ("hello" + "foo bar")) {
    case "hello":
        fmt.Println(v24, v11)
    case "\n\t":
        for fi25 := 0; fi25 < 9; fi25++ {
            for fi26 := 0; fi26 < 15; fi26++ {
                v27 := GenS3{F0: fi25, F1: 2}
                var v28 int = (-(fi25))
                _ = fi26
                _ = v27
                _ = v28
            }
            _ = fi25
        }
    case "":
        if (!(!true)) {
            var v29 uint8 = 127
            _ = v29
        } else {
            v11 -= v24
        }
    default:
        switch 0 {
        case 544:
            v11 = (v24 + v24)
        case -407:
            fmt.Println(v1, v24, v11)
        }
    }
    v1 = (v1 + v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v7
    _ = v9
}
