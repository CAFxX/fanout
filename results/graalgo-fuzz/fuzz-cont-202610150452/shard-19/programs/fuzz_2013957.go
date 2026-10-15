package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int; F1 uint8; F2 uint16 }
type GenS4 struct { F0 int64; F1 bool; F2 uint32 }
func genFn0(p0 int64, p1 int8) int16 {
    for fi7 := 0; fi7 < 17; fi7++ {
        if (0 <= 1000) {
            switch "foo bar" {
            case "a":
                switch (fi7 << ((65535) & 63)) {
                case 236:
                    var v8 bool = (-128 > (p1 - p1))
                    _ = v8
                case 13:
                    v9 := GenS1{F0: 2, F1: 42, F2: 1}
                    _ = v9
                }
            case "hello":
                v10 := GenS2{F0: 719}
                _ = v10
            default:
                p1 = p1
            }
        } else {
            v11 := []int{fi7}
            _ = v11
        }
        v12 := GenS2{F0: 2}
        _ = fi7
        _ = v12
    }
    v13 := []int{2147483647, -209}
    switch "a" {
    case "a":
        var v14 uint64 = 4294967295
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v13
    return -394
}
func genFn1(p0 uint64, p1 uint8, p2 uint64) string {
    if true {
        v15 := []int{2147483647, -9223372036854775808, -9223372036854775808}
        _ = v15
    }
    fmt.Println(p2)
    if (!true) {
        switch ("a" + "foo bar") {
        case "\n\t":
            v16 := GenS1{F0: 2, F1: 1000, F2: 151}
            _ = v16
        case "foo bar":
            fmt.Println(p0, p2)
        case "hello":
            v17 := []int{-1, -9223372036854775808}
            _ = v17
        }
        fmt.Println(p2, p1, p0)
        var v18 int = 779
        _ = v18
    } else {
        var v19 int = -83
        if (!((!(false && (!false))) && false)) {
            v20 := []int{(v19 * (v19 - 9223372036854775807)), 0}
            for fi21 := 0; fi21 < 1; fi21++ {
                p2 = p2
                v19 = 0
                fmt.Println(v19, p2, fi21)
                _ = fi21
            }
            _ = v20
        }
        _ = v19
    }
    _ = p0
    _ = p1
    _ = p2
    return "a"
}

func main() {
    var v1 int16 = 1000
    v2 := make(map[uint32]int)
    var v3 int = 1
    var v4 int64 = 0
    var v5 uint8 = 127
    v6 := make(map[string]string)
    switch (-(v3)) {
    case 418:
        v22 := GenS1{F0: 908, F1: 757, F2: 0}
        _ = v22
    case 2147483647:
        delete(v2, uint32(v5))
    }
    for fi23 := 0; fi23 < 10; fi23++ {
        v5 -= (^(v5))
        if false {
            if true {
                var v24 float64 = ((0.1 - 1e10) / 1.0)
                v25 := v2[745]
                if true {
                    for fi26 := 0; fi26 < 4; fi26++ {
                        var v27 int8 = 42
                        var v28 uint8 = v5
                        var v29 string = ("foo bar" + "foo bar")
                        delete(v2, 42)
                        _ = fi26
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                    v24 += v24
                } else {
                    var v30 uint8 = (v5 + v5)
                    _ = v30
                }
                _ = v24
                _ = v25
            } else {
                for fi31 := 0; fi31 < 18; fi31++ {
                    v32 := GenS3{F0: 128, F1: v5, F2: 458}
                    v33 := []int{32767, (v3 | (fi23 & -512)), v3}
                    _ = fi31
                    _ = v32
                    _ = v33
                }
                if ((-2.25 * (1.5 + 0.0)) < -2.25) {
                    var v34 float64 = ((1.5 * 0.1) + (1.5 * 3.14159))
                    _ = v34
                }
                for fi35 := 0; fi35 < 16; fi35++ {
                    v36 := GenS4{F0: v4, F1: (4294967295 >= 4294967295), F2: 656}
                    _ = fi35
                    _ = v36
                }
            }
            var v37 float32 = 0.1
            _ = v37
        }
        _ = fi23
    }
    fmt.Println(v5, v4, v3)
    v3 = (^(v3))
    var v38 uint32 = 753
    var v39 uint32 = (^(v38))
    fmt.Println(v5, v5, v3)
    var v40 int64 = v4
    for fi41 := 0; fi41 < 16; fi41++ {
        if false {
            for fi42 := 0; fi42 < 14; fi42++ {
                fmt.Println(v1, v1, v39)
                _ = fi42
            }
            v39 = (v38 * v38)
        }
        if false {
            v39 = v39
            var v43 string = ""
            _ = v43
        } else {
            v1 += v1
            v3 = fi41
            for fi44 := 0; fi44 < 18; fi44++ {
                v40 = 42
                _ = fi44
            }
        }
        var v45 uint8 = v5
        _ = fi41
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
