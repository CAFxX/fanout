package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 int16; F1 int8; F2 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool }
func genFn0() string {
    switch "a" {
    case "":
        var v6 uint32 = 2147483647
        _ = v6
    default:
        var v7 uint64 = 0
        _ = v7
    }
    v8 := GenS1{F0: -353, F1: 2}
    v9 := GenS1{F0: 986, F1: 0}
    var v10 float64 = 3.14159
    if (((0.1 * v10) <= v10) || (((255 < 0) || false) && (250 >= 58))) {
        if (58 == 127) {
            if ("hello" < ("\n\t" + "")) {
                v10 = v10
                switch "foo bar" {
                case "hello":
                    v10 += (((v10 * 0.1) - float64(((0.1 * -2.25) / 0.1))) / v10)
                case "":
                    switch -972 {
                    case 0:
                        var v11 float32 = 0.0
                        _ = v11
                    case -1:
                        var v12 int32 = 601
                        _ = v12
                    case 127:
                        var v13 int32 = -1
                        _ = v13
                    default:
                        v14 := []int{1, 4294967295, 862}
                        _ = v14
                    }
                default:
                    var v15 uint16 = 42
                    _ = v15
                }
            } else {
                if (!(!true)) {
                    v16 := GenS1{F0: -1, F1: 32767}
                    _ = v16
                }
                v10 -= (v10 * 0.1)
                v10 += (-2.25 - (v10 - v10))
            }
        } else {
            v10 += -2.25
            var v17 int8 = -4
            _ = v17
        }
        v18 := []int{0, 128, 242, 9223372036854775807}
        _ = len(v18)
        _ = v18
    }
    _ = v10
    _ = v8
    _ = v9
    return "hello"
}
func genFn1(p0 string) uint {
    v19 := []int{66}
    p0 = ""
    for fi20 := 0; fi20 < 10; fi20++ {
        if false {
            if (!(!(171 <= 32767))) {
                for fi21 := 0; fi21 < 2; fi21++ {
                    p0 = p0
                    _ = fi21
                }
                p0 = ("" + p0)
            }
        }
        _ = fi20
    }
    _ = p0
    _ = v19
    return 896
}

func main() {
    var v1 uint16 = 2
    var v2 uint8 = 1
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]bool)
    var v5 int8 = 100
    switch 1 {
    case -651:
        if false {
            var v22 uint = 2
            _ = v22
        } else {
            var v23 int64 = 1
            var v24 int8 = 42
            var v25 int64 = (v23 ^ v23)
            _ = v23
            _ = v24
            _ = v25
        }
    case 32768:
        var v26 int = 2147483647
        _ = v26
    }
    var v27 int64 = 0
    fmt.Println(v5)
    for fi28 := 0; fi28 < 7; fi28++ {
        v27 = (-(v27))
        _ = fi28
    }
    var v29 float64 = (float64(1e10) - (-2.25 * (-2.25 - 1.5)))
    v30 := v4[199]
    var v31 int64 = (v27 / ((v27) | 1))
    for fi32 := 0; fi32 < 11; fi32++ {
        v33 := []int{(fi32 / ((32768) | 1)), 811}
        if false {
            fmt.Println(v29, fi32, v5)
            var v34 int = -9223372036854775808
            if (2147483648 >= 732) {
                v35 := GenS2{F0: -1, F1: (v5 / ((v5) | 1)), F2: "a"}
                _ = v35
            }
            _ = v34
        }
        v36 := GenS3{F0: 910}
        _ = fi32
        _ = v33
        _ = v36
    }
    for fi37 := 0; fi37 < 17; fi37++ {
        v1 += v1
        _ = fi37
    }
    v2 -= 1
    v31 = v27
    v31 = v31
    v38 := GenS4{F0: (v30 || (("a" + "a") < ""))}
    v39 := GenS4{F0: (-2147483648 == 0)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
