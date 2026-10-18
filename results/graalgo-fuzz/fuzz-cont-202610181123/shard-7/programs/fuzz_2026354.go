package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 int32) int32 {
    if ((((1e10 * 1e10) / 1.0) + (1e10 / (1e10 - float64(0.1)))) < 0.0) {
        var v7 int32 = p0
        var v8 uint = 4294967295
        for fi9 := 0; fi9 < 4; fi9++ {
            p0 = p0
            v7 += (p0 + v7)
            _ = fi9
        }
        _ = v7
        _ = v8
    }
    v10 := []int{}
    _ = p0
    _ = v10
    return (p0 >> ((42) & 63))
}

func main() {
    v1 := []uint32{}
    var v2 uint8 = 255
    v3 := []int{-9223372036854775808, 2147483647, 32768, 2147483647}
    var v4 uint16 = 32767
    v5 := []int{127, -303, 128}
    var v6 uint = 447
    v6 += 998
    v11 := GenS1{F0: v6}
    fmt.Println(v6)
    for fi12 := 0; fi12 < 16; fi12++ {
        if (0 < 0) {
            var v13 uint8 = (v2 / ((205) | 1))
            var v14 bool = false
            _ = v13
            _ = v14
        } else {
            fmt.Println(v2, v6, v4)
            v1 = append(v1, 2147483648)
            switch fi12 {
            case 4294967295:
                var v15 int16 = 1
                _ = v15
            case -483:
                if (!(!false)) {
                    v6 = 592
                } else {
                    _ = len(v1)
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            default:
                for fi16 := 0; fi16 < 18; fi16++ {
                    var v17 bool = ((!false) && true)
                    var v18 bool = v17
                    if (v18 || true) {
                        v6 = (v6 * (v6 * 1))
                    } else {
                        v18 = ((fi12 & 1) == fi16)
                        v19 := GenS1{F0: v6}
                        var v20 string = "\n\t"
                        _ = v19
                        _ = v20
                    }
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            }
        }
        v6 = v6
        v6 = 0
        _ = fi12
    }
    if false {
        fmt.Println(v6, v4)
        var v21 uint32 = 42
        v22 := GenS1{F0: v6}
        _ = v21
        _ = v22
    }
    v23 := GenS1{F0: uint(v2)}
    v24 := GenS1{F0: (^(v6))}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if len(v5) > 0 {
        v5[0] = -2147483648
    }
    switch ("" + ("\n\t" + "foo bar")) {
    case "hello":
        fmt.Println(v4, v4)
    case "":
        v2 += (v2 / (((v2 >> (((v6 & v6)) & 63))) | 1))
    case "foo bar":
        switch "a" {
        case "a":
            var v25 int64 = 0
            _ = v25
        default:
            var v26 int32 = -1
            _ = v26
        }
    }
    switch ("a" + ("hello" + ("a" + ("foo bar" + "foo bar")))) {
    case "":
        fmt.Println(v6, v2, v4)
    case "\n\t":
        switch ("" + "foo bar") {
        case "a":
            switch ("hello" + (("hello" + "a") + "foo bar")) {
            case "hello":
                v6 = v6
            case "a":
                var v27 uint64 = 0
                _ = v27
            }
        default:
            fmt.Println(v4, v6)
        }
    case "foo bar":
        fmt.Println(v6)
    default:
        fmt.Println(v4, v6, v2)
    }
    fmt.Println(v6, v6)
    v6 = (v6 - 368)
    fmt.Println(v6, v2)
    var v28 int = 32767
    for fi29 := 0; fi29 < 18; fi29++ {
        v30 := GenS1{F0: (v6 | (v6 >> ((v6) & 63)))}
        v28 = fi29
        _ = fi29
        _ = v30
    }
    switch ("foo bar" + "a") {
    case "foo bar":
        fmt.Println(v2)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v23
    _ = v24
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
