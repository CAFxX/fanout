package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 string }
type GenS2 struct { F0 uint8 }
func genFn0() int8 {
    var v4 string = ""
    if false {
        var v5 int32 = -1
        var v6 int = -905
        for fi7 := 0; fi7 < 2; fi7++ {
            v8 := []int{(fi7 % ((fi7) | 1)), 102}
            var v9 int = fi7
            for fi10 := 0; fi10 < 4; fi10++ {
                var v11 int = 683
                _ = fi10
                _ = v11
            }
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = v5
        _ = v6
    } else {
        for fi12 := 0; fi12 < 13; fi12++ {
            v13 := []int{}
            for fi14 := 0; fi14 < 18; fi14++ {
                var v15 float64 = ((float64(-2.25) / (3.14159 * (((0.0 / 1.0) + (0.0 - 0.1)) / -2.25))) - (-2.25 / (0.1 + (((3.14159 + 0.1) * -2.25) - 0.1))))
                v4 = "foo bar"
                v16 := GenS1{F0: fi14, F1: 42, F2: ((v4 + (v4 + (("" + (v4 + v4)) + ((v4 + v4) + v4)))) + ("" + v4))}
                _ = fi14
                _ = v15
                _ = v16
            }
            _ = fi12
            _ = v13
        }
    }
    _ = v4
    return 42
}

func main() {
    v1 := []int{460, -599}
    v2 := make(map[int]uint32)
    v3 := make(map[string]int)
    var v17 uint16 = 32767
    var v18 int64 = -825
    v18 = (v18 | v18)
    if ((100 != -1) || (true || (!true))) {
        v19 := GenS1{F0: -8, F1: genFn0(), F2: ""}
        _ = v19
    }
    switch "foo bar" {
    case "hello":
        switch "\n\t" {
        case "foo bar":
            _ = len(v2)
        case "\n\t":
            var v20 uint16 = v17
            _ = v20
        case "hello":
            if true {
                v21 := GenS2{F0: 0}
                _ = v21
            }
        }
    }
    switch "\n\t" {
    case "foo bar":
        for fi22 := 0; fi22 < 1; fi22++ {
            v17 = (v17 * 32767)
            v18 -= v18
            _ = fi22
        }
    case "":
        switch ("a" + "foo bar") {
        case "hello":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case "a":
            for fi23 := 0; fi23 < 18; fi23++ {
                v2[32768] = 2
                v24 := GenS2{F0: 85}
                _ = fi23
                _ = v24
            }
        default:
            var v25 int16 = 42
            _ = v25
        }
    }
    if (1 <= 0) {
        fmt.Println(v17, v18)
        switch "foo bar" {
        case "hello":
            _ = len(v1)
        }
        v18 -= (v18 % ((v18) | 1))
    } else {
        for fi26 := 0; fi26 < 11; fi26++ {
            switch (fi26 - 2147483647) {
            case 1:
                if (false && true) {
                    v18 += v18
                } else {
                    v3["a"] = fi26
                    v17 = 42
                    v27 := GenS2{F0: 2}
                    _ = v27
                }
            default:
                _ = len(v1)
            }
            v28 := GenS1{F0: (fi26 - fi26), F1: -83, F2: ("" + "hello")}
            _ = fi26
            _ = v28
        }
        v18 += (^(v18))
    }
    fmt.Println(v17)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
