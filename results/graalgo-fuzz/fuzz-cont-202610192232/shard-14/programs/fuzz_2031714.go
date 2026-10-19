package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 string }
type GenS2 struct { F0 uint }
func genFn0(p0 uint32, p1 uint32, p2 bool) uint16 {
    var v5 int16 = -317
    var v6 uint8 = 2
    v5 += (v5 / ((int16(v6)) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 32767
}

func main() {
    v1 := make(map[string]int)
    var v2 int = 753
    var v3 uint32 = 42
    v4 := make(map[string]uint32)
    v2 += (^(v2))
    switch "a" {
    case "\n\t":
        fmt.Println(v3, v3)
    case "hello":
        var v7 int64 = -1
        _ = v7
    case "foo bar":
        switch "" {
        case "a":
            v3 += 350
        }
    default:
        v2 = v2
    }
    fmt.Println(v3)
    for fi8 := 0; fi8 < 12; fi8++ {
        fmt.Println(fi8)
        v9 := []int{(fi8 << ((65535) & 63)), (v2 - v2)}
        _ = fi8
        _ = v9
    }
    fmt.Println(v2, v3)
    var v10 uint16 = 1
    for fi11 := 0; fi11 < 2; fi11++ {
        v12 := []int{(fi11 * 32767)}
        for fi13 := 0; fi13 < 2; fi13++ {
            switch "\n\t" {
            case "a":
                v3 = v3
            case "":
                var v14 bool = true
                _ = v14
            case "\n\t":
                var v15 uint16 = (v10 ^ 1)
                _ = v15
            default:
                switch v2 {
                case 42:
                    v16 := GenS1{F0: 100, F1: ((fi13 & -195) >= fi13), F2: "foo bar"}
                    _ = v16
                default:
                    v17 := GenS1{F0: -128, F1: (-1 <= 0), F2: ""}
                    if (uint(v10) >= 2) {
                        v18 := GenS1{F0: 100, F1: (2147483647 <= v2), F2: ""}
                        v10 = 2
                        var v19 uint8 = 128
                        _ = v18
                        _ = v19
                    }
                    v2 = (fi13 << ((255) & 63))
                    _ = v17
                }
                var v20 uint64 = 790
                _ = v20
            }
            _ = len(v4)
            _ = fi13
        }
        for fi21 := 0; fi21 < 10; fi21++ {
            var v22 int32 = 178
            fmt.Println(v22, v3, v10)
            if (fi11 > (^(v2))) {
                fmt.Println(v10, v22)
            } else {
                v3 += (v3 & (v3 & v3))
                v10 -= 74
            }
            _ = fi21
            _ = v22
        }
        _ = fi11
        _ = v12
    }
    var v23 int16 = 32767
    for fi24 := 0; fi24 < 20; fi24++ {
        if (false || (-840 != 403)) {
            var v25 string = "a"
            var v26 uint64 = 968
            v27 := GenS1{F0: 1, F1: false, F2: v25}
            _ = v25
            _ = v26
            _ = v27
        }
        _ = fi24
    }
    v3 -= v3
    switch "hello" {
    case "hello":
        switch "" {
        case "a":
            v3 = v3
        default:
            for fi28 := 0; fi28 < 17; fi28++ {
                v29 := []int{(fi28 * -803), fi28, 4294967295}
                if (false || (0 == 1)) {
                    if len(v29) >= 2 {
                        v29 = v29[1:len(v29)]
                    }
                } else {
                    switch (-(fi28)) {
                    case 127:
                        var v30 uint8 = 42
                        _ = v30
                    }
                }
                _ = fi28
                _ = v29
            }
        }
    case "a":
        v31 := GenS2{F0: 2}
        _ = v31
    case "":
        v23 = (v23 / (((v23 ^ (v23 * 0))) | 1))
    default:
        var v32 int16 = (v23 + -32768)
        _ = v32
    }
    if false {
        var v33 string = "a"
        v2 = -813
        _ = v33
    } else {
        switch v2 {
        case 40:
            v34 := GenS2{F0: 174}
            _ = v34
        case 32768:
            v3 -= 2147483648
        case 127:
            v10 = (v10 + v10)
        }
        v3 = 655
    }
    v3 -= v3
    for fi35 := 0; fi35 < 20; fi35++ {
        if false {
            fmt.Println(v23, v23)
            var v36 uint64 = 248
            _ = v36
        }
        _ = fi35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
