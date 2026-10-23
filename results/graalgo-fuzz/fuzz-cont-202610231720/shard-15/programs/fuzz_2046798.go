package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
func genFn0(p0 string, p1 uint16) uint32 {
    for fi4 := 0; fi4 < 7; fi4++ {
        p0 = p0
        switch (p0 + (("foo bar" + p0) + p0)) {
        case "a":
            v5 := []int{-2147483648, (fi4 ^ (fi4 >> ((18446744073709551615) & 63)))}
            _ = v5
        case "\n\t":
            p0 = p0
        case "foo bar":
            p1 = 32767
        default:
            switch (p0 + (p0 + ("" + p0))) {
            case "\n\t":
                var v6 uint = 268
                _ = v6
            default:
                if true {
                    v7 := []int{-1}
                    var v8 uint16 = (p1 - 42)
                    _ = v7
                    _ = v8
                } else {
                    v9 := GenS1{F0: 128, F1: 9223372036854775807}
                    _ = v9
                }
            }
        }
        _ = fi4
    }
    fmt.Println(p0, p0)
    fmt.Println(p1, p0)
    v10 := GenS1{F0: -2147483648, F1: -52}
    _ = p0
    _ = p1
    _ = v10
    return 42
}

func main() {
    var v1 string = "a"
    v2 := make(map[int]string)
    v3 := make(map[string]string)
    fmt.Println(v1)
    var v11 uint16 = 1
    v1 = v1
    v12 := GenS1{F0: -9223372036854775808, F1: 9223372036854775807}
    if (1000 >= 923) {
        var v13 uint32 = 540
        v14 := GenS1{F0: 128, F1: 0}
        var v15 int64 = -1
        _ = v13
        _ = v14
        _ = v15
    }
    var v16 int16 = -420
    if (114 >= 100) {
        v16 = (-(v16))
    } else {
        for fi17 := 0; fi17 < 6; fi17++ {
            if true {
                switch v1 {
                case "foo bar":
                    v1 = "foo bar"
                case "":
                    v18 := GenS1{F0: 0, F1: 57}
                    _ = v18
                }
                var v19 bool = (18446744073709551615 >= 42)
                _ = v19
            }
            delete(v3, v1)
            v20 := GenS1{F0: 127, F1: -1}
            _ = fi17
            _ = v20
        }
        var v21 string = v1
        v1 = "foo bar"
        _ = v21
    }
    if (!false) {
        switch -841 {
        case -273:
            v16 = 32767
        case 4294967295:
            v22 := []int{-2147483648, 901, -2147483648, 546}
            _ = v22
        case -2147483648:
            v16 += v16
        default:
            for fi23 := 0; fi23 < 8; fi23++ {
                v1 = ("" + v1)
                _ = fi23
            }
        }
        var v24 uint16 = (^(v11))
        _ = v24
    }
    if true {
        var v25 uint64 = 0
        v16 += v16
        v26 := GenS1{F0: 693, F1: -592}
        _ = v25
        _ = v26
    } else {
        for fi27 := 0; fi27 < 17; fi27++ {
            var v28 uint32 = 2
            fmt.Println(v1, v28, v28)
            if ((155 > 1) || ((v28 ^ 894) == (^(v28)))) {
                switch "\n\t" {
                case "a":
                    for fi29 := 0; fi29 < 20; fi29++ {
                        var v30 int8 = 0
                        v11 -= 32767
                        _ = fi29
                        _ = v30
                    }
                case "hello":
                    for fi31 := 0; fi31 < 4; fi31++ {
                        switch v1 {
                        case "foo bar":
                            v32 := GenS1{F0: fi31, F1: 935}
                            _ = v32
                        case "\n\t":
                            v28 = (v28 ^ (v28 % ((v28) | 1)))
                        }
                        var v33 uint16 = v11
                        _ = fi31
                        _ = v33
                    }
                default:
                    v28 = v28
                }
            }
            _ = fi27
            _ = v28
        }
        _ = len(v3)
        switch 12 {
        case 32768:
            for fi34 := 0; fi34 < 12; fi34++ {
                if (float64(((0.0 / 1e10) + float32(0.0))) >= 1.5) {
                    if false {
                        var v35 uint = 255
                        fmt.Println(fi34, v35)
                        _ = v35
                    }
                    fmt.Println(fi34, v1)
                }
                _ = fi34
            }
        case 45:
            fmt.Println(v16, v16)
        case -9223372036854775808:
            v16 = v16
        default:
            var v36 int = 0
            _ = v36
        }
    }
    for fi37 := 0; fi37 < 5; fi37++ {
        v16 = (v16 % ((1) | 1))
        if true {
            var v38 uint16 = (v11 + v11)
            var v39 uint = 42
            v38 += 803
            _ = v38
            _ = v39
        }
        _ = fi37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
}
