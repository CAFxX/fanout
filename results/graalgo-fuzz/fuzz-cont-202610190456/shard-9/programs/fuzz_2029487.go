package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 uint64; F1 bool }
func genFn0() uint16 {
    var v5 string = "a"
    v5 = "foo bar"
    var v6 int = 358
    v5 = v5
    _ = v5
    _ = v6
    return 71
}
func genFn1() bool {
    v7 := []int{}
    if len(v7) > 0 {
        v7[0] = -587
    }
    var v8 bool = false
    _ = v7
    _ = v8
    return false
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]int)
    var v3 uint = 255
    var v4 uint8 = 42
    switch "foo bar" {
    case "a":
        var v9 uint = uint(v4)
        _ = v9
    case "hello":
        var v10 float64 = -2.25
        _ = v10
    case "":
        switch -709 {
        case 32767:
            v4 -= v4
        }
    default:
        if false {
            v3 = (v3 % ((0) | 1))
        }
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        v12 := []int{0, (-(fi11)), fi11}
        var v13 string = ((("foo bar" + "hello") + "a") + "")
        _ = fi11
        _ = v12
        _ = v13
    }
    if (-534 >= -2147483648) {
        var v14 int8 = 44
        for fi15 := 0; fi15 < 3; fi15++ {
            fmt.Println(v3, v14, v4)
            switch (("a" + "\n\t") + "") {
            case "a":
                v14 -= (v14 | -1)
            case "":
                v3 = 815
            }
            _ = fi15
        }
        v16 := GenS1{F0: v3}
        _ = v14
        _ = v16
    }
    switch 32767 {
    case -285:
        v3 = (v3 << ((255) & 63))
    case 4294967295:
        if (((4294967295 != 18446744073709551615) && true) || true) {
            if false {
                v17 := []int{}
                if (!(-2.25 <= 1e10)) {
                    var v18 float32 = 0.1
                    _ = v18
                } else {
                    fmt.Println(v4, v4)
                }
                var v19 uint64 = 720
                _ = v17
                _ = v19
            }
        }
    default:
        switch ("foo bar" + "foo bar") {
        case "hello":
            switch "foo bar" {
            case "a":
                var v20 string = "a"
                _ = v20
            case "hello":
                v4 += 127
            case "\n\t":
                for fi21 := 0; fi21 < 17; fi21++ {
                    for fi22 := 0; fi22 < 13; fi22++ {
                        v4 -= v4
                        _ = fi22
                    }
                    var v23 uint32 = 584
                    _ = fi21
                    _ = v23
                }
            default:
                var v24 uint8 = (v4 - 0)
                _ = v24
            }
        case "":
            v25 := GenS2{F0: 1, F1: ""}
            _ = v25
        }
    }
    v26 := GenS1{F0: (v3 | (^(v3)))}
    v2[("a" + (("\n\t" + "\n\t") + "a"))] = 0
    var v27 int8 = 77
    if true {
        v28 := GenS3{F0: 743, F1: true}
        _ = v28
    } else {
        v27 += v27
    }
    var v29 int8 = v27
    fmt.Println(v27, v29)
    v2[("" + "\n\t")] = 18
    v30 := []int{-1}
    v29 += v29
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
