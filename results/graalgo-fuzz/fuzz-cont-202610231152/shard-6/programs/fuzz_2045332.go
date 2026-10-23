package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 uint32; F1 int; F2 int32 }
type GenS4 struct { F0 uint16 }
func genFn0(p0 int, p1 uint, p2 int8) bool {
    v7 := []int{(p0 + (p0 >> ((p1) & 63))), p0, (p0 << ((255) & 63))}
    var v8 int16 = -148
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (-675 == -2147483648)
}
func genFn1(p0 uint8) bool {
    p0 = (p0 | 1)
    var v9 bool = true
    switch "" {
    case "\n\t":
        v10 := []int{-1}
        _ = v10
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        if (1 != 407) {
            p0 = 1
        } else {
            var v12 uint64 = 2
            var v13 int = (^(fi11))
            _ = v12
            _ = v13
        }
        for fi14 := 0; fi14 < 20; fi14++ {
            v15 := []int{1, -9223372036854775808, (-(fi14)), (fi11 << ((0) & 63))}
            _ = len(v15)
            if len(v15) > 0 {
                v15[0] = 899
            }
            var v16 int32 = 2147483647
            p0 = (^(p0))
            if false {
                fmt.Println(fi14)
                var v17 int = 1
                var v18 string = ""
                _ = v17
                _ = v18
            } else {
                p0 -= p0
                v19 := GenS1{F0: 350}
                _ = v19
            }
            _ = fi14
            _ = v15
            _ = v16
        }
        fmt.Println(v9)
        _ = fi11
    }
    _ = p0
    _ = v9
    return true
}

func main() {
    var v1 uint8 = 0
    v2 := []int{-1, -827}
    var v3 uint = 796
    v4 := []string{"a", "hello", "hello", "hello"}
    var v5 int64 = 914
    var v6 string = ("a" + "")
    v2 = append(v2, 812)
    for fi20 := 0; fi20 < 18; fi20++ {
        v21 := GenS1{F0: (^(v3))}
        fmt.Println(fi20)
        _ = fi20
        _ = v21
    }
    var v22 int = 121
    var v23 int64 = (v5 * v5)
    if (!(((v6 + "foo bar") <= (("a" + "hello") + (v6 + (v6 + "a")))) || false)) {
        v24 := GenS1{F0: 4294967295}
        _ = v24
    } else {
        if (!false) {
            for fi25 := 0; fi25 < 2; fi25++ {
                _ = len(v4)
                for fi26 := 0; fi26 < 7; fi26++ {
                    var v27 bool = false
                    if (v27 && v27) {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    } else {
                        if len(v2) > 0 {
                            v2[0] = fi26
                        }
                        if v27 {
                            var v28 uint32 = 101
                            _ = v28
                        }
                        v23 = v23
                    }
                    _ = fi26
                    _ = v27
                }
                if (102 <= (v1 * (v1 ^ 0))) {
                    var v29 int = fi25
                    _ = v29
                }
                _ = fi25
            }
        }
        var v30 uint64 = 42
        _ = v30
    }
    var v31 int64 = (v5 >> ((2) & 63))
    v1 = 0
    v6 = (((((v6 + v6) + v6) + v6) + (v6 + v6)) + "hello")
    v23 -= (v23 / ((v5) | 1))
    fmt.Println(v23, v23, v31)
    v32 := GenS2{F0: (-(v31)), F1: v22}
    var v33 int8 = -128
    v5 = v31
    v6 = "a"
    v34 := GenS3{F0: 1, F1: (v22 / ((v22) | 1)), F2: -1}
    v6 = v6
    switch "foo bar" {
    case "hello":
        fmt.Println(v31, v5)
    default:
        if false {
            v35 := GenS4{F0: 32767}
            var v36 float32 = (-2.25 / 1e10)
            _ = v35
            _ = v36
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
