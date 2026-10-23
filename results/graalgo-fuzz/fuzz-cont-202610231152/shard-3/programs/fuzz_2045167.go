package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 uint }
func genFn0(p0 uint8, p1 uint32, p2 uint) int64 {
    for fi7 := 0; fi7 < 2; fi7++ {
        p2 = (p2 << ((p2) & 63))
        switch (("foo bar" + "\n\t") + "foo bar") {
        case "":
            p2 = p2
        case "a":
            for fi8 := 0; fi8 < 18; fi8++ {
                fmt.Println(p2, p0, p1)
                v9 := []int{(fi8 - 9223372036854775807)}
                _ = fi8
                _ = v9
            }
        }
        v10 := GenS1{F0: uint(p1), F1: 2147483647, F2: p2}
        _ = fi7
        _ = v10
    }
    if (42 <= -2147483648) {
        for fi11 := 0; fi11 < 13; fi11++ {
            switch (-(fi11)) {
            case -126:
                if true {
                    var v12 int16 = 42
                    switch 1 {
                    case -1:
                        if false {
                            fmt.Println(fi11, v12)
                            v13 := []int{fi11}
                            _ = v13
                        } else {
                            fmt.Println(v12)
                            v12 += (v12 & 461)
                        }
                    default:
                        p1 -= p1
                    }
                    _ = v12
                } else {
                    for fi14 := 0; fi14 < 5; fi14++ {
                        switch (fi11 / (((fi11 + fi11)) | 1)) {
                        case 1:
                            if true {
                                p2 = 65535
                                v15 := []int{fi14}
                                _ = v15
                            }
                        case 613:
                            p0 -= 42
                        case 4294967295:
                            fmt.Println(fi14)
                        }
                        _ = fi14
                    }
                }
            case 32767:
                v16 := GenS1{F0: 1, F1: 42, F2: (^(p2))}
                _ = v16
            case 2147483647:
                var v17 uint = (p2 / ((18446744073709551615) | 1))
                _ = v17
            default:
                v18 := GenS1{F0: (p2 & (p2 | p2)), F1: -563, F2: p2}
                _ = v18
            }
            _ = fi11
        }
    }
    v19 := []int{32768, 4294967295, 9223372036854775807}
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return -370
}

func main() {
    var v1 int16 = 134
    var v2 uint16 = 42
    var v3 int64 = -1
    var v4 uint8 = 128
    var v5 uint8 = (v4 / (((v4 / ((v4) | 1))) | 1))
    var v6 float64 = float64(float32(1e10))
    if (false && true) {
        for fi20 := 0; fi20 < 3; fi20++ {
            v21 := GenS2{F0: ("" + "a")}
            var v22 int8 = -96
            _ = fi20
            _ = v21
            _ = v22
        }
    } else {
        for fi23 := 0; fi23 < 1; fi23++ {
            if (2 != 255) {
                v24 := GenS2{F0: "\n\t"}
                v3 += int64(v4)
                v25 := []int{(^(fi23)), fi23}
                _ = v24
                _ = v25
            } else {
                v26 := GenS2{F0: ("hello" + "\n\t")}
                _ = v26
            }
            var v27 uint = 42
            _ = fi23
            _ = v27
        }
    }
    v28 := GenS3{F0: -47, F1: (2147483647 <= 42)}
    if (!((-2.25 - v6) != -2.25)) {
        fmt.Println(v1)
        v29 := []int{32768, 4294967295}
        _ = v29
    }
    switch "a" {
    case "":
        var v30 uint32 = 2147483648
        _ = v30
    case "\n\t":
        v31 := GenS1{F0: 255, F1: -642, F2: 390}
        _ = v31
    case "a":
        var v32 uint64 = 4294967295
        _ = v32
    }
    if true {
        var v33 string = "hello"
        var v34 int64 = v3
        _ = v33
        _ = v34
    }
    if true {
        switch (("foo bar" + "a") + "\n\t") {
        case "":
            v35 := []int{797, 984}
            _ = v35
        case "foo bar":
            var v36 int32 = -140
            _ = v36
        case "\n\t":
            var v37 uint16 = 1
            _ = v37
        default:
            v38 := GenS1{F0: 4294967295, F1: 1, F2: 1}
            _ = v38
        }
        var v39 float32 = (1e10 * 0.1)
        _ = v39
    }
    fmt.Println(v3, v5, v3)
    switch 32767 {
    case 4294967295:
        if ((v5 >> ((25) & 63)) < v5) {
            var v40 float32 = -2.25
            if true {
                v1 = v1
                v1 = v1
                var v41 uint8 = v4
                _ = v41
            }
            if (v3 > (v3 ^ (v3 + (v3 & v3)))) {
                switch 249 {
                case 1:
                    v4 = v4
                case -2147483648:
                    var v42 float64 = 3.14159
                    _ = v42
                }
            }
            _ = v40
        }
    case -222:
        v43 := GenS4{F0: 1}
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
