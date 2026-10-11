package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 uint8 }
type GenS2 struct { F0 uint64; F1 int32; F2 uint16 }
type GenS3 struct { F0 uint16; F1 string; F2 string }
type GenS4 struct { F0 int; F1 int64; F2 int32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint32; F1 int; F2 uint }
func genFn0(p0 uint16, p1 uint8) int16 {
    for fi7 := 0; fi7 < 5; fi7++ {
        p0 = (p0 & (p0 + p0))
        var v8 uint32 = 65535
        switch "foo bar" {
        case "\n\t":
            p0 = (p0 ^ p0)
        case "foo bar":
            switch "\n\t" {
            case "\n\t":
                if ((((1 <= 1) || false) && false) || false) {
                    p1 += (p1 >> ((1) & 63))
                    var v9 int16 = -94
                    _ = v9
                } else {
                    fmt.Println(v8)
                    fmt.Println(v8, v8, p0)
                }
            }
        case "hello":
            for fi10 := 0; fi10 < 3; fi10++ {
                var v11 int8 = -1
                v12 := []int{(fi10 >> ((2) & 63)), (fi7 * fi7), (fi7 ^ fi7)}
                var v13 int16 = 32767
                _ = fi10
                _ = v11
                _ = v12
                _ = v13
            }
        default:
            v8 += (v8 % ((v8) | 1))
        }
        _ = fi7
        _ = v8
    }
    var v14 uint32 = 2147483647
    p1 = p1
    _ = p0
    _ = p1
    _ = v14
    return 1000
}

func main() {
    var v1 string = "\n\t"
    var v2 uint16 = 42
    v3 := make(map[string]bool)
    var v4 uint8 = 31
    v5 := make(map[int]uint32)
    var v6 bool = true
    fmt.Println(v1)
    var v15 uint8 = 127
    v4 += v15
    for fi16 := 0; fi16 < 10; fi16++ {
        v17 := GenS1{F0: (!false), F1: 42, F2: 1}
        _ = fi16
        _ = v17
    }
    var v18 uint = uint(v2)
    for fi19 := 0; fi19 < 5; fi19++ {
        var v20 bool = ((0.0 == 0.0) && false)
        v2 = (v2 - v2)
        _ = fi19
        _ = v20
    }
    v21 := GenS1{F0: (v2 != 1), F1: -1, F2: v15}
    v5[9223372036854775807] = 1
    var v22 float64 = float64(3.14159)
    fmt.Println(v1)
    if (1e10 == (1e10 + float32((3.14159 / v22)))) {
        v6 = v6
        for fi23 := 0; fi23 < 6; fi23++ {
            v15 += 10
            v24 := GenS1{F0: v6, F1: 772, F2: v15}
            var v25 int8 = 0
            _ = fi23
            _ = v24
            _ = v25
        }
        for fi26 := 0; fi26 < 12; fi26++ {
            if (false && true) {
                v27 := GenS2{F0: 18446744073709551615, F1: -2147483648, F2: v2}
                switch (v1 + ((v1 + "a") + v1)) {
                case "a":
                    for fi28 := 0; fi28 < 6; fi28++ {
                        v18 = v18
                        _ = fi28
                    }
                }
                switch v1 {
                case "\n\t":
                    var v29 int64 = 42
                    _ = v29
                }
                _ = v27
            }
            for fi30 := 0; fi30 < 16; fi30++ {
                var v31 uint = v18
                _ = fi30
                _ = v31
            }
            _ = fi26
        }
    } else {
        var v32 uint = (v18 & 65535)
        v6 = (-106 != 81)
        switch -1 {
        case 207:
            if v6 {
                if v6 {
                    v33 := GenS3{F0: v2, F1: v1, F2: (v1 + "")}
                    for fi34 := 0; fi34 < 6; fi34++ {
                        switch "" {
                        case "a":
                            v35 := GenS1{F0: (27 < v4), F1: 2147483647, F2: (v15 - 0)}
                            _ = v35
                        case "foo bar":
                            v18 -= uint(v2)
                        case "\n\t":
                            v36 := v5[(fi34 & (fi34 - (fi34 ^ (fi34 | 128))))]
                            _ = v36
                        }
                        switch "foo bar" {
                        case "a":
                            v22 = float64(0.1)
                        }
                        v18 -= 18446744073709551615
                        _ = fi34
                    }
                    _ = v33
                }
                fmt.Println(v18, v1, v32)
                switch "foo bar" {
                case "a":
                    var v37 bool = (float64(1e10) != float64((((-2.25 - 3.14159) * 0.1) / 3.14159)))
                    _ = v37
                case "foo bar":
                    _ = len(v3)
                }
            } else {
                var v38 int32 = -2147483648
                _ = v38
            }
        case -2147483648:
            if (v6 || (1.5 == 3.14159)) {
                v39 := GenS4{F0: 223, F1: 42, F2: 488}
                v40 := []int{}
                switch -973 {
                case -309:
                    var v41 int8 = 0
                    _ = v41
                }
                _ = v39
                _ = v40
            }
        case -9223372036854775808:
            fmt.Println(v32)
        default:
            var v42 bool = true
            _ = v42
        }
        _ = v32
    }
    v43 := GenS5{F0: 1}
    v44 := GenS5{F0: 4294967295}
    var v45 uint = (^(v18))
    var v46 int8 = 100
    switch (v1 + "hello") {
    case "hello":
        v47 := GenS6{F0: 2, F1: 42, F2: 955}
        _ = v47
    case "\n\t":
        var v48 float64 = float64((1e10 / (1e10 * 0.1)))
        _ = v48
    default:
        var v49 float32 = ((3.14159 - 1.5) + 1.5)
        _ = v49
    }
    if v6 {
        var v50 int16 = genFn0((v2 - v2), v4)
        _ = v50
    } else {
        var v51 int8 = -60
        v18 = 255
        _ = v51
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
