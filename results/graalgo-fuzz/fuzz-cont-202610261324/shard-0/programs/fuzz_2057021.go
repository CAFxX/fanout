package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 int16 }
func genFn0(p0 uint64) uint32 {
    var v7 float64 = 0.1
    var v8 uint8 = 2
    fmt.Println(p0)
    p0 = p0
    _ = p0
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 int8, p1 bool, p2 int16) uint32 {
    if (-2.25 == (0.1 + -2.25)) {
        p1 = false
        var v9 bool = false
        _ = v9
    } else {
        v10 := GenS1{F0: 18446744073709551615}
        var v11 uint = 0
        v12 := GenS1{F0: 1}
        _ = v10
        _ = v11
        _ = v12
    }
    v13 := []int{42, 32768, -414}
    switch "" {
    case "a":
        var v14 float64 = (0.0 * 0.0)
        _ = v14
    }
    fmt.Println(p0, p0, p0)
    var v15 int = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v15
    return genFn0(0)
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = 1
    v3 := make(map[int]int)
    v4 := make(map[string]bool)
    var v5 uint16 = 0
    var v6 int8 = (v2 / (((v2 % ((v2) | 1))) | 1))
    var v16 uint = 2
    if false {
        v17 := GenS1{F0: 895}
        _ = v17
    }
    v5 -= v5
    v16 = v16
    if ((65535 < 0) && (18 <= v5)) {
        var v18 uint = 815
        var v19 float64 = (0.0 / -2.25)
        _ = v18
        _ = v19
    } else {
        v16 = (v16 & (v16 | v16))
        v16 -= v16
        v5 = (v5 ^ (v5 | uint16(v16)))
    }
    _ = len(v3)
    v16 -= (v16 % (((v16 ^ v16)) | 1))
    if true {
        if true {
            switch (("a" + "hello") + ("hello" + "hello")) {
            case "\n\t":
                switch 175 {
                case -2147483648:
                    if false {
                        v2 = 112
                    }
                case 0:
                    fmt.Println(v16)
                }
            default:
                var v20 uint16 = v5
                _ = v20
            }
            fmt.Println(v5, v6)
        }
        fmt.Println(v5, v5)
        switch 32767 {
        case 9223372036854775807:
            if ((2147483647 != 32767) && true) {
                var v21 uint16 = v5
                _ = v21
            }
        case 4294967295:
            var v22 float64 = (3.14159 * ((((1e10 + (-2.25 - float64((0.0 - 1.5)))) + (((float64((0.0 * float32(1e10))) - -2.25) * 1e10) * (-2.25 * float64(-2.25)))) / -2.25) + 0.0))
            _ = v22
        case 2147483647:
            if (!(((-(v6)) == (v6 + v6)) || (3.14159 <= 1.5))) {
                var v23 int16 = 1000
                _ = v23
            } else {
                switch "hello" {
                case "":
                    var v24 int32 = -130
                    _ = v24
                case "a":
                    var v25 int16 = -1
                    _ = v25
                default:
                    var v26 bool = true
                    _ = v26
                }
                v27 := GenS1{F0: 2}
                v6 -= (v6 | v2)
                _ = v27
            }
        default:
            fmt.Println(v5, v16, v16)
        }
    }
    v5 -= (^(v5))
    for fi28 := 0; fi28 < 17; fi28++ {
        v6 -= v2
        fmt.Println(v5, v6)
        _ = fi28
    }
    v29 := GenS2{F0: "a"}
    var v30 int64 = 9223372036854775807
    for fi31 := 0; fi31 < 12; fi31++ {
        var v32 uint16 = (v5 / (((v5 | 255)) | 1))
        fmt.Println(v16)
        _ = fi31
        _ = v32
    }
    var v33 int = -782
    v34 := GenS2{F0: "foo bar"}
    if false {
        if true {
            v35 := []int{}
            v33 = (v33 - v33)
            _ = v35
        } else {
            v5 = 610
            v36 := []int{(v33 ^ v33), (v33 ^ (v33 - (^(v33)))), (v33 >> ((v16) & 63)), (v33 >> (((v16 >> ((v16) & 63))) & 63))}
            _ = v36
        }
    }
    for fi37 := 0; fi37 < 10; fi37++ {
        v38 := GenS1{F0: 42}
        for fi39 := 0; fi39 < 10; fi39++ {
            if true {
                var v40 string = ""
                if (((v2 / ((v2) | 1)) <= -80) || true) {
                    var v41 int16 = 1000
                    v42 := GenS2{F0: v40}
                    v43 := v3[v33]
                    v44 := GenS3{F0: (v5 << ((v16) & 63)), F1: v30}
                    if true {
                        v45 := GenS4{F0: (v41 - 42)}
                        fmt.Println(v6, v43)
                        _ = v45
                    }
                    delete(v4, (("" + v40) + v40))
                    fmt.Println(v40, v30, v33)
                    var v46 int64 = (v30 + v30)
                    fmt.Println(v33, v5, fi39)
                    v47 := GenS4{F0: (v41 >> (((v16 / (((^(v16))) | 1))) & 63))}
                    _ = v41
                    _ = v42
                    _ = v43
                    _ = v44
                    _ = v46
                    _ = v47
                }
                fmt.Println(v16, v40)
                _ = v40
            }
            var v48 float32 = (3.14159 + float32(((1e10 / (((0.1 + ((1e10 + float64(-2.25)) / float64(float32(1.5)))) * (3.14159 / (-2.25 - (float64(((1e10 + 0.0) - 0.1)) / float64(0.1))))) * float64(0.1))) * 0.0)))
            _ = fi39
            _ = v48
        }
        var v49 float32 = float32(3.14159)
        _ = fi37
        _ = v38
        _ = v49
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
