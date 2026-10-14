package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 int8; F1 bool; F2 uint32 }
func genFn0(p0 int8, p1 int64, p2 bool) int {
    for fi4 := 0; fi4 < 3; fi4++ {
        for fi5 := 0; fi5 < 12; fi5++ {
            var v6 uint = 0
            if p2 {
                var v7 uint32 = 2
                v8 := []int{fi4}
                _ = v7
                _ = v8
            }
            if false {
                if (p2 && (false || false)) {
                    var v9 uint = (v6 * v6)
                    var v10 float64 = (0.0 + 1e10)
                    if (0 != 32767) {
                        v11 := []int{9223372036854775807, fi5, (fi4 ^ fi4)}
                        _ = v11
                    } else {
                        if true {
                            v10 += ((-2.25 - (v10 / float64((float32(v10) / ((1e10 * (1e10 - -2.25)) * -2.25))))) * 3.14159)
                            v12 := []int{fi4, 127, (-(fi5))}
                            v13 := GenS1{F0: 32767, F1: 0, F2: -2147483648}
                            _ = v12
                            _ = v13
                        }
                        v10 = 0.0
                    }
                    _ = v10
                    _ = v9
                }
            } else {
                v14 := GenS1{F0: 32767, F1: -1, F2: (fi5 & (-(fi5)))}
                if (p2 || p2) {
                    p1 -= -230
                    var v15 bool = false
                    v16 := GenS1{F0: 32767, F1: 1000, F2: (fi5 % (((fi5 - 127)) | 1))}
                    _ = v15
                    _ = v16
                }
                _ = v14
            }
            if (42 == 476) {
                p0 -= (p0 / ((p0) | 1))
                p1 = -9223372036854775808
                switch ("" + "foo bar") {
                case "hello":
                    v17 := GenS2{F0: v6}
                    _ = v17
                default:
                    v18 := []int{fi5, (fi5 >> ((v6) & 63))}
                    _ = v18
                }
            } else {
                p0 += (p0 * -128)
            }
            p2 = p2
            _ = fi5
            _ = v6
        }
        _ = fi4
    }
    var v19 int8 = (p0 % ((p0) | 1))
    p2 = (!(1e10 <= float64((0.0 / 0.1))))
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return int(p0)
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]string)
    var v3 uint8 = 128
    _ = len(v2)
    v3 += (v3 >> ((825) & 63))
    v1[1] = 545
    var v20 int16 = 1
    v2[""] = ("" + "a")
    var v21 float64 = 0.1
    v3 = v3
    var v22 int32 = 717
    for fi23 := 0; fi23 < 6; fi23++ {
        var v24 uint32 = 214
        for fi25 := 0; fi25 < 16; fi25++ {
            fmt.Println(v3, v20, fi23)
            _ = fi25
        }
        _ = fi23
        _ = v24
    }
    switch "a" {
    case "":
        switch ("a" + "foo bar") {
        case "hello":
            var v26 int8 = -31
            _ = v26
        default:
            v3 += 1
        }
    default:
        if true {
            v21 -= v21
            fmt.Println(v20, v3, v22)
        }
    }
    fmt.Println(v21)
    if (127 != 41) {
        fmt.Println(v22, v22, v20)
        if true {
            v22 = (v22 ^ 0)
            v27 := []int{}
            v21 -= 0.0
            _ = v27
        } else {
            v28 := GenS3{F0: 118, F1: 160}
            var v29 uint = 0
            for fi30 := 0; fi30 < 3; fi30++ {
                v21 = v21
                _ = fi30
            }
            _ = v28
            _ = v29
        }
        for fi31 := 0; fi31 < 5; fi31++ {
            v20 += v20
            _ = fi31
        }
    } else {
        v32 := GenS4{F0: 127, F1: true, F2: 952}
        v3 = v3
        _ = v32
    }
    v33 := []int{128, genFn0(-1, -161, (!true)), -484, 0}
    if true {
        if true {
            v21 -= v21
        }
    } else {
        var v34 int = 0
        fmt.Println(v34, v21)
        _ = v34
    }
    var v35 int16 = v20
    delete(v2, "foo bar")
    for fi36 := 0; fi36 < 1; fi36++ {
        v37 := GenS3{F0: 0, F1: v35}
        v22 = v22
        _ = fi36
        _ = v37
    }
    v38 := GenS4{F0: 42, F1: true, F2: 4294967295}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v33
    _ = v35
    _ = v38
}
