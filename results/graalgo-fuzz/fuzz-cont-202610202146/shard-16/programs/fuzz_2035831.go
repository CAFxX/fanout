package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 uint }
func genFn0(p0 int32) int16 {
    var v6 string = "hello"
    v7 := GenS1{F0: 65535}
    fmt.Println(v6, p0)
    _ = p0
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 int) int16 {
    p0 = (p0 & -1)
    for fi8 := 0; fi8 < 14; fi8++ {
        for fi9 := 0; fi9 < 12; fi9++ {
            var v10 uint64 = 18446744073709551615
            v11 := GenS1{F0: 2}
            v10 = 222
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi8
    }
    switch "hello" {
    case "a":
        p0 = -2147483648
    case "hello":
        for fi12 := 0; fi12 < 1; fi12++ {
            var v13 float32 = (float32(0.1) / (-2.25 + 0.1))
            _ = fi12
            _ = v13
        }
    case "foo bar":
        for fi14 := 0; fi14 < 20; fi14++ {
            var v15 uint16 = 255
            _ = fi14
            _ = v15
        }
    default:
        v16 := []int{-1, (-(p0))}
        _ = v16
    }
    var v17 float64 = ((3.14159 / float64(3.14159)) * float64(float32(-2.25)))
    _ = p0
    _ = v17
    return -47
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]uint32)
    v3 := make(map[string]bool)
    var v4 uint64 = 900
    var v5 int64 = 509
    v18 := GenS1{F0: 2}
    v4 = v4
    var v19 int32 = 0
    for fi20 := 0; fi20 < 1; fi20++ {
        for fi21 := 0; fi21 < 20; fi21++ {
            var v22 float64 = (3.14159 * 1.5)
            v5 -= (v5 - (v5 & (v5 | (v5 - v5))))
            var v23 int = (-(fi20))
            switch fi20 {
            case 2147483647:
                if true {
                    v4 += v4
                } else {
                    delete(v3, (("" + "foo bar") + ""))
                    var v24 bool = true
                    _ = v24
                }
            case 127:
                v19 = v19
            default:
                fmt.Println(v23, v5, v19)
            }
            _ = fi21
            _ = v22
            _ = v23
        }
        if false {
            if false {
                v5 -= (v5 % ((-135) | 1))
                v25 := []int{}
                _ = v25
            } else {
                fmt.Println(fi20)
            }
            if (2 <= 1) {
                v5 = 1
                fmt.Println(v4, v5)
                delete(v1, 4294967295)
            }
            if ((true || (0.0 < 0.0)) || (1 > -53)) {
                fmt.Println(v5, v5, v4)
            } else {
                v4 += (^(v4))
            }
        } else {
            var v26 float32 = float32(0.0)
            for fi27 := 0; fi27 < 12; fi27++ {
                var v28 float32 = 3.14159
                _ = fi27
                _ = v28
            }
            _ = v26
        }
        _ = fi20
    }
    var v29 uint32 = 65535
    for fi30 := 0; fi30 < 13; fi30++ {
        for fi31 := 0; fi31 < 1; fi31++ {
            v32 := []int{127}
            if (((false && (true || (false && (((!true) || (1e10 <= -2.25)) || true)))) && true) && (false || (65535 >= 2))) {
                var v33 uint64 = v4
                _ = v33
            } else {
                switch "\n\t" {
                case "a":
                    v4 -= (v4 | v4)
                    v19 = v19
                case "foo bar":
                    if false {
                        var v34 float64 = 3.14159
                        _ = v34
                    } else {
                        var v35 int8 = 42
                        fmt.Println(fi31, fi31)
                        _ = v35
                    }
                default:
                    var v36 uint = 21
                    _ = v36
                }
                v37 := GenS2{F0: ("hello" + ("\n\t" + "hello")), F1: 2}
                _ = v37
            }
            v38 := v1[(v29 >> ((42) & 63))]
            _ = fi31
            _ = v32
            _ = v38
        }
        if (0 <= 127) {
            v39 := GenS1{F0: 42}
            switch "" {
            case "a":
                _ = len(v2)
            case "foo bar":
                v40 := GenS2{F0: "", F1: 4294967295}
                _ = v40
            default:
                v29 = v29
            }
            v41 := GenS1{F0: 18446744073709551615}
            _ = v39
            _ = v41
        } else {
            v4 -= (v4 + v4)
            switch (fi30 & (fi30 / (((fi30 | (fi30 << ((uint(v29)) & 63)))) | 1))) {
            case 127:
                v19 = (v19 | (v19 & -1))
            case 854:
                var v42 uint32 = 209
                _ = v42
            }
            v43 := GenS2{F0: "a", F1: 65535}
            _ = v43
        }
        switch "" {
        case "\n\t":
            var v44 bool = false
            _ = v44
        case "foo bar":
            var v45 uint32 = (v29 * v29)
            _ = v45
        case "hello":
            if false {
                var v46 int8 = -108
                _ = v46
            }
        }
        _ = fi30
    }
    for fi47 := 0; fi47 < 13; fi47++ {
        for fi48 := 0; fi48 < 5; fi48++ {
            fmt.Println(v4, v5)
            _ = fi48
        }
        _ = fi47
    }
    fmt.Println(v5, v4, v4)
    if false {
        v49 := []int{int(v29), 9223372036854775807, 1}
        _ = v49
    } else {
        var v50 uint32 = 2
        _ = v50
    }
    v5 = (v5 + v5)
    v19 += v19
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
