package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int; F1 int }
func genFn0(p0 bool, p1 string) int16 {
    var v5 int64 = 42
    var v6 int = -446
    p1 = "\n\t"
    p1 = ""
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return -839
}
func genFn1(p0 uint32, p1 bool) int8 {
    v7 := GenS1{F0: 4294967295, F1: -46}
    for fi8 := 0; fi8 < 15; fi8++ {
        p0 += p0
        v9 := []int{fi8, -1, 42, (fi8 * -1)}
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: 647, F1: 32767}
    for fi11 := 0; fi11 < 12; fi11++ {
        var v12 int = (fi11 << ((65535) & 63))
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    return -118
}

func main() {
    v1 := []string{"foo bar", "foo bar"}
    v2 := make(map[int]bool)
    v3 := make(map[string]bool)
    var v4 uint16 = 42
    v4 += v4
    if false {
        v4 = v4
    }
    var v13 bool = (genFn0((!(true || true)), (("" + (("hello" + "hello") + ("\n\t" + ((("" + ("foo bar" + "")) + "hello") + "a")))) + ("a" + (("hello" + ("hello" + "hello")) + "\n\t")))) > -623)
    var v14 float64 = -2.25
    if true {
        if v13 {
            if (genFn0(false, ("foo bar" + "")) <= 0) {
                if (v13 || true) {
                    var v15 uint = 697
                    _ = v15
                } else {
                    v14 -= (float64(-2.25) / (1.5 - (-2.25 / v14)))
                    var v16 int64 = 42
                    _ = v16
                }
                if len(v1) > 0 {
                    v1[0] = "a"
                }
                v4 = (v4 + 42)
            }
        } else {
            fmt.Println(v13, v13)
        }
    }
    if (3.14159 > -2.25) {
        v13 = v13
        var v17 uint8 = 255
        _ = v17
    }
    var v18 int16 = genFn0(("hello" >= "a"), "\n\t")
    switch "hello" {
    case "a":
        var v19 int = -1
        _ = v19
    case "hello":
        var v20 string = "foo bar"
        _ = v20
    default:
        var v21 int32 = -40
        _ = v21
    }
    if (439 != -597) {
        v14 = (float64(0.1) * (((v14 + (v14 / ((v14 / float64(0.1)) - -2.25))) * float64(float32(v14))) / ((v14 - 0.1) / 1.0)))
        if (!v13) {
            v4 = v4
            fmt.Println(v4)
        }
        v22 := GenS2{F0: "foo bar"}
        _ = v22
    } else {
        for fi23 := 0; fi23 < 1; fi23++ {
            v13 = true
            v14 = (v14 - -2.25)
            _ = fi23
        }
    }
    v24 := GenS2{F0: "\n\t"}
    v25 := GenS3{F0: 4294967295, F1: 0}
    v1 = append(v1, "hello")
    fmt.Println(v13, v18, v4)
    for fi26 := 0; fi26 < 12; fi26++ {
        if v13 {
            _ = len(v1)
            v13 = (v13 || v13)
            if false {
                if v13 {
                    if (false || v13) {
                        v18 = 32767
                    }
                    if len(v1) > 0 {
                        v1[0] = (("hello" + "") + ("hello" + ("a" + "")))
                    }
                }
                fmt.Println(v4)
            } else {
                for fi27 := 0; fi27 < 19; fi27++ {
                    v3["foo bar"] = (255 <= 246)
                    _ = fi27
                }
                switch "foo bar" {
                case "hello":
                    v28 := GenS3{F0: fi26, F1: 128}
                    _ = v28
                case "a":
                    for fi29 := 0; fi29 < 19; fi29++ {
                        fmt.Println(v4, fi26, v4)
                        _ = fi29
                    }
                case "foo bar":
                    v30 := GenS2{F0: "\n\t"}
                    _ = v30
                }
            }
        } else {
            v18 = v18
            v31 := GenS1{F0: 484, F1: v18}
            _ = v31
        }
        var v32 uint8 = 2
        var v33 uint64 = 42
        _ = fi26
        _ = v32
        _ = v33
    }
    var v34 int = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v34
    _ = v4
}
