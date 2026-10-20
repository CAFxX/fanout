package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 int8; F1 string }
func genFn0(p0 uint16, p1 uint) int {
    p0 = p0
    for fi5 := 0; fi5 < 7; fi5++ {
        for fi6 := 0; fi6 < 12; fi6++ {
            fmt.Println(p1)
            switch int(p1) {
            case 965:
                v7 := GenS1{F0: "a", F1: false, F2: (p1 ^ p1)}
                _ = v7
            default:
                switch (-(fi6)) {
                case 32768:
                    v8 := []int{}
                    _ = v8
                case 42:
                    var v9 int = 4294967295
                    _ = v9
                case 4294967295:
                    fmt.Println(p0, p1, p0)
                }
            }
            _ = fi6
        }
        _ = fi5
    }
    _ = p0
    _ = p1
    return -9223372036854775808
}
func genFn1(p0 uint16, p1 uint, p2 string) int16 {
    for fi10 := 0; fi10 < 19; fi10++ {
        v11 := GenS1{F0: ("" + ("hello" + p2)), F1: true, F2: p1}
        var v12 int64 = 42
        v13 := []int{(fi10 >> (((^(p1))) & 63)), (fi10 & (fi10 * 32767)), -2147483648}
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    v14 := []int{4294967295, -2147483648, 590, 0}
    var v15 uint = (p1 & p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    return 42
}

func main() {
    v1 := make(map[string]bool)
    v2 := []string{"", "hello", "hello"}
    var v3 uint = 826
    var v4 int = 223
    if true {
        for fi16 := 0; fi16 < 16; fi16++ {
            if len(v2) > 0 {
                v2[0] = ("\n\t" + "")
            }
            _ = fi16
        }
    }
    var v17 int64 = 0
    var v18 int64 = 0
    v17 += 0
    switch -659 {
    case 2147483647:
        var v19 uint8 = 127
        _ = v19
    case 127:
        v17 += v18
    case 0:
        if (0.0 >= ((0.1 + -2.25) + 0.1)) {
            for fi20 := 0; fi20 < 7; fi20++ {
                v21 := GenS1{F0: ("a" + ""), F1: false, F2: v3}
                if len(v2) > 0 {
                    v2[0] = "hello"
                }
                if (true || true) {
                    for fi22 := 0; fi22 < 10; fi22++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v4 -= (v4 >> (((v3 - (^(v3)))) & 63))
                        v1[""] = false
                        _ = fi22
                    }
                    v17 += v18
                    v1["\n\t"] = (!true)
                } else {
                    if (false && true) {
                        v23 := GenS1{F0: ("\n\t" + "foo bar"), F1: true, F2: (v3 / ((65535) | 1))}
                        _ = v23
                    } else {
                        v3 = (v3 * (^(v3)))
                        fmt.Println(v17)
                        v17 += v17
                    }
                    v17 = -540
                }
                _ = fi20
                _ = v21
            }
            fmt.Println(v3, v18)
        } else {
            v24 := GenS2{F0: 2, F1: (-(v18))}
            _ = v24
        }
    }
    for fi25 := 0; fi25 < 6; fi25++ {
        var v26 string = ("foo bar" + ("hello" + "hello"))
        v27 := GenS3{F0: 100, F1: ("a" + v26)}
        fmt.Println(v26, fi25)
        _ = fi25
        _ = v26
        _ = v27
    }
    fmt.Println(v17, v17, v18)
    v3 += 0
    v3 += v3
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    if (true || true) {
        switch (-(v4)) {
        case -747:
            if (float32((1e10 - (((((1.5 + (float64(0.0) / 1.0)) * 0.0) - 0.0) / -2.25) - float64((0.1 * float32((-2.25 + (float64(-2.25) * ((float64(0.1) - 0.1) / -2.25))))))))) == (3.14159 * -2.25)) {
                _ = len(v2)
                var v28 float64 = (3.14159 - (1.5 - 0.1))
                var v29 float32 = (3.14159 - 3.14159)
                _ = v28
                _ = v29
            } else {
                var v30 int8 = -128
                _ = v30
            }
        case 4294967295:
            if true {
                var v31 int8 = -98
                _ = v31
            }
        }
        _ = len(v2)
    } else {
        var v32 int32 = 2147483647
        v1["hello"] = false
        _ = v32
    }
    var v33 int = 1
    if (!(v18 != 0)) {
        var v34 uint32 = 0
        v35 := GenS2{F0: 0, F1: (v17 & (^(v17)))}
        var v36 int = (v4 << ((v3) & 63))
        _ = v34
        _ = v35
        _ = v36
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v33
    _ = v4
}
