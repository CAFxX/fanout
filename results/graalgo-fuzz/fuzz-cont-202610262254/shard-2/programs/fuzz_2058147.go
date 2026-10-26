package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 uint16; F2 uint }
type GenS4 struct { F0 string; F1 bool }
type GenS5 struct { F0 bool; F1 int; F2 bool }
type GenS6 struct { F0 int32 }
func genFn0(p0 uint16, p1 uint16) int {
    for fi7 := 0; fi7 < 6; fi7++ {
        p1 = (p1 ^ (p1 | (p1 ^ p0)))
        var v8 float64 = 1e10
        v9 := GenS1{F0: 4294967295, F1: 184}
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 string = ""
    v11 := GenS1{F0: 7, F1: uint8(p1)}
    v12 := []int{826, 2147483647}
    if (false || (-2.25 == -2.25)) {
        if false {
            _ = len(v12)
        }
        p1 -= (p0 >> ((0) & 63))
        var v13 float32 = ((0.1 + 1.5) / 1.0)
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    return 32767
}

func main() {
    v1 := make(map[int]int)
    var v2 uint32 = 2147483647
    v3 := make(map[string]string)
    var v4 float64 = 0.1
    var v5 float32 = 1e10
    v6 := []uint32{(^(v2)), v2, v2}
    var v14 uint16 = 32767
    for fi15 := 0; fi15 < 1; fi15++ {
        fmt.Println(v4, v14)
        _ = fi15
    }
    v4 -= v4
    var v16 int8 = 102
    if (false && (0 != 580)) {
        fmt.Println(v5)
    }
    _ = len(v1)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if (556 >= 65535) {
        for fi17 := 0; fi17 < 10; fi17++ {
            v4 = ((-2.25 * v4) * (v4 - v4))
            var v18 bool = (!false)
            switch -9223372036854775808 {
            case -9223372036854775808:
                fmt.Println(fi17)
            case 127:
                fmt.Println(fi17)
            case 128:
                v19 := GenS1{F0: (v2 >> ((0) & 63)), F1: 127}
                _ = v19
            default:
                var v20 string = ("hello" + "")
                _ = v20
            }
            _ = fi17
            _ = v18
        }
    } else {
        v21 := v3["a"]
        _ = len(v1)
        if ((("hello" + (v21 + ((v21 + v21) + (v21 + "hello")))) + v21) != ("" + "hello")) {
            switch (v21 + (("\n\t" + (((v21 + (v21 + ("" + (v21 + "foo bar")))) + v21) + ("a" + (v21 + "hello")))) + ((((((v21 + "foo bar") + v21) + "hello") + (("" + v21) + (v21 + (v21 + v21)))) + v21) + "a"))) {
            case "foo bar":
                v22 := GenS2{F0: (^(v2))}
                _ = v22
            case "a":
                switch (v21 + v21) {
                case "hello":
                    v21 = ""
                default:
                    for fi23 := 0; fi23 < 4; fi23++ {
                        var v24 float32 = (v5 - (float32(0.1) + v5))
                        v25 := GenS1{F0: (v2 | (v2 ^ v2)), F1: 0}
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                }
            }
            v5 = float32((v4 - 0.1))
            var v26 int8 = -116
            _ = v26
        }
        _ = v21
    }
    v1[982] = genFn0((v14 * 65535), v14)
    v27 := GenS2{F0: 945}
    for fi28 := 0; fi28 < 1; fi28++ {
        fmt.Println(v4, v4)
        for fi29 := 0; fi29 < 3; fi29++ {
            v30 := GenS3{F0: (fi28 - fi29), F1: (v14 + (v14 | 42)), F2: 0}
            v31 := GenS4{F0: (("foo bar" + ("hello" + ((("hello" + ((("a" + "") + ("\n\t" + (("" + "hello") + "a"))) + "")) + "\n\t") + ("" + "foo bar")))) + "hello"), F1: (255 < uint(v2))}
            switch 9223372036854775807 {
            case 171:
                var v32 uint64 = 769
                _ = v32
            case 1:
                var v33 float64 = v4
                _ = v33
            case 378:
                if true {
                    switch fi28 {
                    case -304:
                        if len(v6) > 0 {
                            v6[0] = (v2 ^ v2)
                        }
                    }
                } else {
                    fmt.Println(fi28, fi29)
                }
            }
            _ = fi29
            _ = v30
            _ = v31
        }
        v34 := GenS5{F0: (float32(v4) > (float32(v4) * 3.14159)), F1: (^(fi28)), F2: ((0 >= -235) || (!false))}
        _ = fi28
        _ = v34
    }
    if false {
        for fi35 := 0; fi35 < 1; fi35++ {
            fmt.Println(v4)
            _ = fi35
        }
    } else {
        _ = len(v1)
        var v36 float32 = (v5 - v5)
        delete(v3, "foo bar")
        _ = v36
    }
    if true {
        var v37 int8 = -83
        fmt.Println(v14, v2, v4)
        if false {
            v16 = 1
            var v38 string = ("hello" + "hello")
            if true {
                fmt.Println(v5)
                if false {
                    var v39 int8 = v16
                    fmt.Println(v16)
                    fmt.Println(v38, v14, v38)
                    _ = v39
                }
                v5 += (v5 + v5)
            }
            _ = v38
        } else {
            v40 := GenS6{F0: 0}
            _ = v40
        }
        _ = v37
    } else {
        fmt.Println(v2, v14, v2)
        var v41 string = "hello"
        _ = v41
    }
    var v42 float32 = ((v5 / ((v5 / (float32(((v4 / -2.25) - (v4 / float64(v5)))) - ((0.0 - ((0.0 + v5) + v5)) + (1.5 - v5)))) - (1.5 - (float32((((v4 + (0.1 / v4)) / v4) + (v4 * (v4 + (v4 * (v4 * float64(v5))))))) * ((((float32(((1.5 / v4) / v4)) * (v5 / v5)) + v5) * float32((3.14159 / 1e10))) + (v5 * -2.25)))))) + v5)
    var v43 bool = (("\n\t" + ("" + "foo bar")) < ("foo bar" + ""))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v16
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
