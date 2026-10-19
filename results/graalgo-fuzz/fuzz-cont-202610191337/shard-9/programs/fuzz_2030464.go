package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint32; F1 uint32; F2 int8 }
type GenS6 struct { F0 int16; F1 int8 }
type GenS7 struct { F0 int; F1 uint32; F2 int8 }
type GenS8 struct { F0 uint16; F1 string; F2 string }
type GenS9 struct { F0 uint; F1 uint8 }
type GenS10 struct { F0 int8; F1 uint8 }
func genFn0(p0 int64, p1 int32, p2 uint) uint8 {
    v7 := GenS1{F0: -2147483648, F1: p2}
    p0 = (-(p0))
    if ((p2 % (((^(p2))) | 1)) <= (p2 | p2)) {
        for fi8 := 0; fi8 < 20; fi8++ {
            p2 = (^(p2))
            for fi9 := 0; fi9 < 15; fi9++ {
                v10 := GenS2{F0: 32768, F1: 42}
                _ = fi9
                _ = v10
            }
            _ = fi8
        }
        for fi11 := 0; fi11 < 4; fi11++ {
            p2 -= (p2 << (((p2 / (((p2 | p2)) | 1))) & 63))
            switch "" {
            case "":
                switch "\n\t" {
                case "":
                    var v12 uint8 = 128
                    _ = v12
                case "a":
                    var v13 float64 = 0.1
                    _ = v13
                case "hello":
                    if (false && (!false)) {
                        switch 32767 {
                        case 42:
                            fmt.Println(p2, fi11, fi11)
                        case 659:
                            v14 := []int{fi11, fi11, (fi11 - fi11), fi11}
                            _ = v14
                        }
                    }
                }
            }
            var v15 float64 = ((-2.25 + -2.25) + 0.1)
            _ = fi11
            _ = v15
        }
    } else {
        switch "foo bar" {
        case "\n\t":
            for fi16 := 0; fi16 < 19; fi16++ {
                p1 -= (p1 ^ int32(fi16))
                _ = fi16
            }
        case "":
            var v17 int8 = 42
            _ = v17
        }
        v18 := []int{208, 127, 128, 32768}
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 237
}
func genFn1(p0 string, p1 bool) int16 {
    p0 = p0
    v19 := GenS1{F0: 42, F1: 4294967295}
    switch p0 {
    case "\n\t":
        p1 = false
    default:
        v20 := []int{188, -9223372036854775808, 9223372036854775807, 1}
        _ = v20
    }
    p0 = p0
    var v21 uint8 = 2
    _ = p0
    _ = p1
    _ = v19
    _ = v21
    return -788
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint16 = 580
    v3 := make(map[int]int)
    var v4 float32 = 3.14159
    var v5 uint = 4294967295
    v6 := make(map[int]uint32)
    var v22 int64 = -1
    v23 := GenS3{F0: v2, F1: 485}
    for fi24 := 0; fi24 < 12; fi24++ {
        v25 := []int{1, (fi24 >> ((v5) & 63)), (fi24 + (-(fi24))), fi24}
        fmt.Println(v4)
        switch ("hello" + "foo bar") {
        case "a":
            v25 = append(v25, fi24)
        case "foo bar":
            v26 := v6[-2147483648]
            _ = v26
        case "":
            v22 -= (v22 + 367)
        }
        _ = fi24
        _ = v25
    }
    v27 := GenS4{F0: true}
    if (-1 == 1) {
        switch "" {
        case "\n\t":
            var v28 float64 = (1.5 * (0.0 - float64(v4)))
            _ = v28
        case "foo bar":
            var v29 uint = (^(v5))
            _ = v29
        case "hello":
            v2 += 2
        }
        for fi30 := 0; fi30 < 18; fi30++ {
            v5 = 976
            fmt.Println(v4, v2, v22)
            v5 = (v5 >> ((v5) & 63))
            _ = fi30
        }
    }
    v31 := GenS5{F0: 2, F1: 42, F2: -1}
    switch 32768 {
    case 42:
        switch ("hello" + ("hello" + "\n\t")) {
        case "\n\t":
            _ = len(v6)
        }
    case -2147483648:
        v32 := GenS6{F0: 1, F1: 0}
        _ = v32
    default:
        if false {
            var v33 uint8 = 30
            _ = v33
        } else {
            v2 = v2
            for fi34 := 0; fi34 < 14; fi34++ {
                v35 := v3[(^(fi34))]
                v2 = (v2 << (((v5 & (v5 - (v5 << ((65535) & 63))))) & 63))
                _ = fi34
                _ = v35
            }
            if false {
                v4 = 3.14159
                v36 := []int{9223372036854775807, 17}
                v37 := GenS7{F0: 127, F1: 2147483648, F2: 100}
                _ = v36
                _ = v37
            }
        }
    }
    v38 := GenS5{F0: 463, F1: 0, F2: 0}
    for fi39 := 0; fi39 < 8; fi39++ {
        if ((-10 != 283) || ((fi39 + fi39) > fi39)) {
            v5 += (v5 / ((310) | 1))
        } else {
            _ = len(v1)
            var v40 int8 = 100
            fmt.Println(v40)
            _ = v40
        }
        _ = fi39
    }
    v41 := []int{int(v2)}
    v4 -= v4
    var v42 float32 = (v4 / (1e10 * (0.0 / (v4 + v4))))
    v43 := GenS8{F0: (v2 << ((v5) & 63)), F1: "a", F2: "foo bar"}
    for fi44 := 0; fi44 < 18; fi44++ {
        v45 := GenS3{F0: (v2 - v2), F1: 42}
        var v46 int = 127
        for fi47 := 0; fi47 < 5; fi47++ {
            v1[593] = 42
            v48 := GenS9{F0: 255, F1: genFn0(0, 0, (v5 / ((v5) | 1)))}
            var v49 int64 = v22
            _ = fi47
            _ = v48
            _ = v49
        }
        _ = fi44
        _ = v45
        _ = v46
    }
    fmt.Println(v5, v4)
    for fi50 := 0; fi50 < 19; fi50++ {
        v42 -= ((0.1 * 0.1) + v4)
        if len(v41) >= 2 {
            v41 = v41[1:len(v41)]
        }
        var v51 bool = true
        _ = fi50
        _ = v51
    }
    for fi52 := 0; fi52 < 1; fi52++ {
        v53 := GenS10{F0: -128, F1: genFn0(v22, -723, (v5 << (((v5 | 42)) & 63)))}
        switch -1 {
        case -969:
            var v54 uint32 = 2147483647
            _ = v54
        case -116:
            fmt.Println(v22, v5)
        case 128:
            v55 := GenS6{F0: 32767, F1: -128}
            _ = v55
        default:
            switch "a" {
            case "foo bar":
                if true {
                    var v56 int = fi52
                    v57 := GenS6{F0: 511, F1: 100}
                    _ = v56
                    _ = v57
                }
            }
        }
        _ = len(v41)
        _ = fi52
        _ = v53
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v27
    _ = v3
    _ = v31
    _ = v38
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
