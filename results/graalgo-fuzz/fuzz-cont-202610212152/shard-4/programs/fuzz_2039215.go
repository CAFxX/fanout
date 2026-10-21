package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }

func main() {
    var v1 float64 = 0.1
    var v2 int16 = 91
    v3 := make(map[int]uint32)
    v4 := []int{0, -9223372036854775808, 127}
    for fi5 := 0; fi5 < 18; fi5++ {
        v6 := GenS1{F0: -108, F1: "\n\t"}
        v2 = -1
        v2 = v2
        _ = fi5
        _ = v6
    }
    fmt.Println(v1, v1)
    var v7 uint32 = 2
    if false {
        v7 += (v7 + (v7 % ((v7) | 1)))
    } else {
        switch -104 {
        case -2147483648:
            var v8 int = 32768
            _ = v8
        case 42:
            for fi9 := 0; fi9 < 17; fi9++ {
                v2 -= 32767
                fmt.Println(v2, v2, fi9)
                var v10 float64 = (v1 - float64((1e10 / -2.25)))
                _ = fi9
                _ = v10
            }
        default:
            if true {
                v7 = 7
                for fi11 := 0; fi11 < 7; fi11++ {
                    v2 -= (-(v2))
                    _ = fi11
                }
            }
        }
    }
    var v12 uint = 736
    var v13 uint16 = uint16(v12)
    fmt.Println(v12, v7, v2)
    switch ("a" + "hello") {
    case "foo bar":
        for fi14 := 0; fi14 < 8; fi14++ {
            fmt.Println(v1, v2, v13)
            _ = fi14
        }
    }
    if true {
        v15 := GenS1{F0: -57, F1: "foo bar"}
        var v16 int8 = 84
        _ = v15
        _ = v16
    } else {
        var v17 uint16 = v13
        _ = v17
    }
    if true {
        for fi18 := 0; fi18 < 3; fi18++ {
            v1 += v1
            var v19 string = ("foo bar" + "a")
            _ = fi18
            _ = v19
        }
    } else {
        var v20 uint8 = 2
        _ = v20
    }
    v13 -= v13
    for fi21 := 0; fi21 < 17; fi21++ {
        var v22 bool = (("foo bar" <= "a") || (false && (!true)))
        if v22 {
            v7 = v7
            if (0 != -909) {
                for fi23 := 0; fi23 < 4; fi23++ {
                    v7 -= 1
                    _ = fi23
                }
                v12 = uint(v13)
            }
        }
        v7 += (v7 / (((v7 ^ (v7 - v7))) | 1))
        _ = fi21
        _ = v22
    }
    var v24 string = "hello"
    fmt.Println(v7, v24)
    var v25 float64 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v7
}
