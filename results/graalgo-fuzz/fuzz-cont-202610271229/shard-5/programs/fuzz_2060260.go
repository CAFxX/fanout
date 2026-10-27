package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 bool; F1 uint32; F2 int }

func main() {
    var v1 bool = false
    var v2 uint16 = 42
    var v3 float64 = (0.0 + -2.25)
    if true {
        v4 := []int{4294967295, 2147483647, 548}
        _ = v4
    } else {
        v1 = (827 <= 18446744073709551615)
        if (!(67 >= 693)) {
            v5 := []int{4294967295, -1, int(v2)}
            _ = v5
        }
    }
    if v1 {
        for fi6 := 0; fi6 < 9; fi6++ {
            var v7 uint8 = 42
            _ = fi6
            _ = v7
        }
        var v8 int = -241
        var v9 uint16 = v2
        _ = v8
        _ = v9
    }
    fmt.Println(v1)
    var v10 uint8 = 1
    fmt.Println(v1, v1, v2)
    fmt.Println(v2, v1)
    for fi11 := 0; fi11 < 17; fi11++ {
        switch "hello" {
        case "a":
            switch "a" {
            case "hello":
                v3 -= v3
            default:
                v3 = (v3 * v3)
            }
        case "hello":
            for fi12 := 0; fi12 < 11; fi12++ {
                var v13 uint64 = 1
                switch -831 {
                case 9223372036854775807:
                    v2 += (v2 + 255)
                    v14 := []int{}
                    _ = v14
                case 32767:
                    v2 -= (^(v2))
                case 42:
                    v3 = float64(1.5)
                    var v15 string = (("a" + "\n\t") + "\n\t")
                    fmt.Println(v13, v15)
                    _ = v15
                }
                var v16 float64 = (v3 / -2.25)
                if (v1 && false) {
                    var v17 uint64 = 727
                    v16 -= (v3 + 0.1)
                    if (0 >= 42) {
                        fmt.Println(v1)
                        var v18 int = fi12
                        _ = v18
                    }
                    _ = v17
                } else {
                    v19 := []int{32767, fi12, fi12, fi12}
                    _ = v19
                }
                _ = fi12
                _ = v13
                _ = v16
            }
        case "":
            v20 := []int{fi11}
            _ = v20
        }
        for fi21 := 0; fi21 < 18; fi21++ {
            v1 = ((^(v2)) != v2)
            _ = fi21
        }
        _ = fi11
    }
    v22 := []int{}
    if (!v1) {
        v23 := GenS1{F0: 2147483647, F1: 189}
        if len(v22) > 0 {
            v22[0] = 42
        }
        switch "foo bar" {
        case "\n\t":
            var v24 string = "hello"
            _ = v24
        case "foo bar":
            fmt.Println(v2)
        default:
            v2 -= (^(v2))
        }
        _ = v23
    }
    for fi25 := 0; fi25 < 7; fi25++ {
        v2 = v2
        v1 = v1
        _ = fi25
    }
    switch 42 {
    case -1:
        fmt.Println(v2)
    case 4294967295:
        switch ("" + ("hello" + "foo bar")) {
        case "foo bar":
            if ((v3 - v3) <= 0.1) {
                for fi26 := 0; fi26 < 13; fi26++ {
                    v2 += (^(v2))
                    for fi27 := 0; fi27 < 7; fi27++ {
                        var v28 uint8 = v10
                        _ = fi27
                        _ = v28
                    }
                    _ = fi26
                }
                var v29 uint32 = 0
                _ = v29
            }
        case "hello":
            if false {
                var v30 int32 = -37
                var v31 float32 = 1.5
                _ = v30
                _ = v31
            }
        case "":
            switch 127 {
            case -9223372036854775808:
                v32 := GenS1{F0: -244, F1: 0}
                _ = v32
            default:
                var v33 string = "\n\t"
                _ = v33
            }
        default:
            v10 = uint8(v2)
        }
    case 776:
        v22 = append(v22, 32767)
    default:
        if len(v22) >= 2 {
            v22 = v22[1:len(v22)]
        }
    }
    if (9223372036854775807 >= 32767) {
        v22 = append(v22, -362)
    } else {
        var v34 int32 = -467
        var v35 float32 = 0.0
        for fi36 := 0; fi36 < 13; fi36++ {
            for fi37 := 0; fi37 < 5; fi37++ {
                _ = len(v22)
                _ = fi37
            }
            if (-828 > 1) {
                v10 = (v10 + (^(v10)))
                switch (("" + (("" + "a") + "")) + "foo bar") {
                case "foo bar":
                    switch (-(fi36)) {
                    case 9223372036854775807:
                        var v38 uint32 = 1
                        _ = v38
                    case 1:
                        v39 := GenS2{F0: v1, F1: 0, F2: (fi36 / ((194) | 1))}
                        _ = v39
                    default:
                        if len(v22) > 0 {
                            v22[0] = (fi36 + 827)
                        }
                    }
                case "hello":
                    if len(v22) > 0 {
                        v22[0] = (fi36 ^ -1)
                    }
                case "\n\t":
                    fmt.Println(v2)
                default:
                    v35 += 0.0
                }
                var v40 uint16 = v2
                _ = v40
            } else {
                v1 = (v1 && (18446744073709551615 == 1))
            }
            _ = fi36
        }
        _ = v34
        _ = v35
    }
    if len(v22) >= 2 {
        v22 = v22[1:len(v22)]
    }
    var v41 int16 = -1
    v10 -= v10
    v10 -= v10
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v3
    _ = v41
}
