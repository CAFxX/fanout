package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint64 }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16; F1 int16; F2 uint16 }
type GenS5 struct { F0 string }
func genFn0(p0 uint16, p1 int16, p2 uint32) int16 {
    v5 := []int{4294967295, 32768, 9223372036854775807}
    var v6 uint8 = 67
    var v7 uint32 = (p2 % (((p2 % (((p2 % ((2147483648) | 1))) | 1))) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 918
}
func genFn1(p0 uint, p1 int) int32 {
    switch "\n\t" {
    case "":
        var v8 bool = ((65535 == 42) && true)
        _ = v8
    case "hello":
        var v9 bool = false
        _ = v9
    }
    v10 := GenS1{F0: -1, F1: (p1 * (p1 + p1)), F2: 18446744073709551615}
    if true {
        switch ("\n\t" + ("" + "a")) {
        case "\n\t":
            fmt.Println(p0, p1)
        case "":
            for fi11 := 0; fi11 < 17; fi11++ {
                p1 = p1
                v12 := []int{(fi11 / ((p1) | 1)), (fi11 / ((-9223372036854775808) | 1))}
                var v13 uint32 = 2147483647
                _ = fi11
                _ = v12
                _ = v13
            }
        case "a":
            v14 := []int{p1, (^(p1)), (p1 & 4294967295), (p1 ^ (p1 - -658))}
            _ = v14
        default:
            v15 := GenS1{F0: 0, F1: p1, F2: 493}
            _ = v15
        }
    }
    fmt.Println(p1, p1)
    switch "hello" {
    case "hello":
        var v16 int16 = 42
        _ = v16
    case "a":
        v17 := GenS2{F0: p1, F1: (p0 - (p0 * p0))}
        _ = v17
    }
    _ = p0
    _ = p1
    _ = v10
    return 0
}

func main() {
    var v1 float32 = -2.25
    var v2 int16 = 1000
    var v3 float64 = 1.5
    var v4 uint32 = 1
    if false {
        if false {
            if true {
                switch 128 {
                case 127:
                    v18 := []int{619}
                    _ = v18
                }
                switch "" {
                case "foo bar":
                    v3 += v3
                case "hello":
                    v19 := GenS1{F0: 1, F1: 127, F2: 207}
                    _ = v19
                case "\n\t":
                    var v20 uint8 = 1
                    _ = v20
                default:
                    var v21 int64 = -9223372036854775808
                    _ = v21
                }
                v2 = v2
            }
            var v22 uint16 = 255
            if false {
                var v23 uint16 = 2
                v22 = (v22 & v23)
                if (42 == 0) {
                    v1 = (((1e10 * 1.5) / 1e10) - float32(v3))
                }
                _ = v23
            } else {
                for fi24 := 0; fi24 < 9; fi24++ {
                    v25 := GenS2{F0: -1, F1: 65535}
                    _ = fi24
                    _ = v25
                }
            }
            _ = v22
        } else {
            v26 := []int{-190, -7, 32768, -121}
            if (180 > 323) {
                var v27 float64 = (v3 - ((((1.5 / (1.5 - (3.14159 - -2.25))) / v3) * v3) + float64(1e10)))
                for fi28 := 0; fi28 < 17; fi28++ {
                    v29 := GenS3{F0: (v4 | v4)}
                    _ = fi28
                    _ = v29
                }
                _ = v27
            }
            _ = v26
        }
        var v30 bool = (127 > -2147483648)
        _ = v30
    }
    if ((!false) || (!(v3 >= float64(((v1 + (float32(v3) / v1)) - (v1 * v1)))))) {
        for fi31 := 0; fi31 < 20; fi31++ {
            var v32 int = 32767
            v33 := GenS2{F0: v32, F1: 18446744073709551615}
            v1 = (1e10 - float32(float64(v1)))
            _ = fi31
            _ = v32
            _ = v33
        }
    } else {
        v4 -= v4
    }
    for fi34 := 0; fi34 < 1; fi34++ {
        v35 := []int{}
        v1 -= (float32(v3) + (v1 / -2.25))
        if true {
            for fi36 := 0; fi36 < 3; fi36++ {
                var v37 uint = 2
                _ = fi36
                _ = v37
            }
            var v38 float64 = float64(1.5)
            _ = v38
        }
        _ = fi34
        _ = v35
    }
    v39 := []int{598, -1}
    v2 = v2
    var v40 float32 = 3.14159
    for fi41 := 0; fi41 < 18; fi41++ {
        v39 = append(v39, (fi41 << ((18446744073709551615) & 63)))
        _ = fi41
    }
    if true {
        if (false && (!false)) {
            if (87 >= 255) {
                var v42 uint16 = 911
                fmt.Println(v1)
                var v43 int64 = 571
                _ = v42
                _ = v43
            }
            v40 -= 3.14159
        }
    }
    fmt.Println(v1, v2)
    if (false && (genFn1(0, -668) <= -329)) {
        v40 = float32(v3)
    }
    v44 := GenS4{F0: v2, F1: genFn0(255, (v2 & (v2 & -1)), 1), F2: 2}
    v45 := GenS4{F0: (v2 / ((v2) | 1)), F1: (v2 & (-(v2))), F2: 2}
    v4 = (^(v4))
    for fi46 := 0; fi46 < 10; fi46++ {
        v47 := GenS5{F0: ("" + "hello")}
        _ = fi46
        _ = v47
    }
    var v48 uint16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v44
    _ = v45
    _ = v48
}
