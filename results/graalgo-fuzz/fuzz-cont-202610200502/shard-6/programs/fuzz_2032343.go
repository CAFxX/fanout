package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 uint16) uint64 {
    v4 := []int{-18, 4294967295, -1, 1}
    p0 -= p0
    _ = p0
    _ = v4
    return uint64(p0)
}
func genFn1() uint16 {
    if true {
        var v5 int8 = 110
        switch 32767 {
        case 9223372036854775807:
            v6 := GenS1{F0: 250}
            _ = v6
        case 2147483647:
            var v7 uint8 = 1
            _ = v7
        default:
            v8 := GenS1{F0: 127}
            _ = v8
        }
        var v9 int8 = (v5 & (v5 & (-(v5))))
        _ = v5
        _ = v9
    } else {
        var v10 uint16 = 65535
        var v11 string = "hello"
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        var v13 int16 = -32768
        v13 = v13
        if true {
            v14 := []int{9223372036854775807}
            _ = v14
        }
        _ = fi12
        _ = v13
    }
    switch "foo bar" {
    case "":
        var v15 int = 2147483647
        _ = v15
    default:
        var v16 uint64 = 2
        _ = v16
    }
    if ((516 == -2147483648) && (true && (("a" + "") < (("foo bar" + "a") + (("hello" + (("foo bar" + "foo bar") + "\n\t")) + "foo bar"))))) {
        if (((((float32(float64((((1e10 - 1.5) / (float32((1e10 / 1.5)) * 0.1)) + (1e10 + -2.25)))) + (0.0 / -2.25)) + (-2.25 / (0.1 + ((0.0 * ((0.1 / 1e10) / 0.1)) * 1.5)))) * 1.5) < 0.1) && false) {
            v17 := GenS1{F0: 127}
            _ = v17
        }
        var v18 string = "foo bar"
        if true {
            if true {
                v19 := []int{32767, 42}
                _ = v19
            }
        }
        _ = v18
    }
    v20 := []int{127, 42, 4294967295, 42}
    _ = v20
    return 990
}

func main() {
    var v1 uint64 = 0
    var v2 bool = (1 < 132)
    v3 := make(map[string]int)
    v2 = false
    if v2 {
        if (v2 && true) {
            v21 := []int{2147483647, -9223372036854775808, 1}
            if ((!(true && (genFn0(602) >= v1))) && v2) {
                v21 = append(v21, -915)
                fmt.Println(v2, v2)
                var v22 uint32 = 65535
                _ = v22
            }
            _ = v21
        }
        v1 = (v1 & (^(v1)))
    } else {
        var v23 bool = (0 <= -1)
        var v24 int = 1
        _ = v23
        _ = v24
    }
    v2 = v2
    switch -1 {
    case -9223372036854775808:
        var v25 int8 = 100
        _ = v25
    default:
        switch "\n\t" {
        case "":
            var v26 uint32 = 575
            _ = v26
        case "foo bar":
            switch 387 {
            case 1:
                v1 += v1
            default:
                var v27 uint = 2
                _ = v27
            }
        case "a":
            v28 := GenS1{F0: 127}
            _ = v28
        }
    }
    var v29 uint64 = (v1 & 785)
    if ("hello" == ("foo bar" + "foo bar")) {
        var v30 string = ("hello" + "a")
        _ = v30
    }
    v2 = v2
    var v31 int8 = 1
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v31
}
