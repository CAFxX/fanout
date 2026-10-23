package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 int16; F2 uint64 }
type GenS3 struct { F0 uint; F1 uint8; F2 uint32 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 780
    v3 := make(map[uint32]int)
    var v4 int8 = 42
    var v5 bool = true
    var v6 uint16 = 1
    v7 := GenS1{F0: 1}
    switch -192 {
    case 473:
        v5 = (42 < 18446744073709551615)
    case 32767:
        fmt.Println(v4, v2, v5)
    case -675:
        v8 := []int{148, 42, 127, -2147483648}
        _ = v8
    }
    switch "\n\t" {
    case "foo bar":
        var v9 int32 = -1
        _ = v9
    case "a":
        v10 := GenS2{F0: -274, F1: -1, F2: 18446744073709551615}
        _ = v10
    }
    var v11 string = "a"
    v12 := []int{128}
    if false {
        v3[(^(v2))] = 32768
        var v13 uint32 = (v2 + 0)
        v4 += v4
        _ = v13
    } else {
        var v14 int = -166
        for fi15 := 0; fi15 < 11; fi15++ {
            var v16 int32 = -1
            for fi17 := 0; fi17 < 5; fi17++ {
                v18 := GenS2{F0: v16, F1: 42, F2: 0}
                switch fi17 {
                case 44:
                    switch ("\n\t" + "") {
                    case "\n\t":
                        v19 := GenS1{F0: 2}
                        _ = v19
                    default:
                        var v20 int = -376
                        _ = v20
                    }
                case 2147483647:
                    delete(v1, v2)
                }
                var v21 int8 = (^(v4))
                _ = fi17
                _ = v18
                _ = v21
            }
            for fi22 := 0; fi22 < 9; fi22++ {
                var v23 int32 = v16
                fmt.Println(v2, v4)
                v5 = v5
                _ = fi22
                _ = v23
            }
            _ = fi15
            _ = v16
        }
        _ = v14
    }
    v12 = append(v12, -25)
    var v24 float32 = -2.25
    if ((v24 + float32(-2.25)) <= v24) {
        var v25 int16 = -32768
        v6 = (^(v6))
        _ = v25
    } else {
        if (1 > uint(v6)) {
            var v26 int32 = -2147483648
            v27 := v1[v2]
            v28 := GenS3{F0: 4294967295, F1: 2, F2: v2}
            _ = v26
            _ = v27
            _ = v28
        } else {
            for fi29 := 0; fi29 < 15; fi29++ {
                switch (-(fi29)) {
                case -1:
                    v30 := GenS1{F0: 136}
                    _ = v30
                }
                _ = fi29
            }
        }
    }
    v4 -= v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
