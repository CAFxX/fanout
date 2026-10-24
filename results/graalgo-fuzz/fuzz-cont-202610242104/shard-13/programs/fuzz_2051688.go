package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint8 }
func genFn0(p0 string, p1 int32, p2 int) int {
    if true {
        var v5 float64 = ((((1e10 / 3.14159) / 1.5) * 1e10) * (-2.25 + (3.14159 - (0.1 / 1e10))))
        v5 -= (0.0 + (v5 * (float64(0.0) - v5)))
        _ = v5
    }
    var v6 uint64 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 127
}
func genFn1(p0 bool, p1 string, p2 uint8) int8 {
    switch (("\n\t" + p1) + (p1 + "hello")) {
    case "a":
        p1 = "a"
    }
    p0 = p0
    v7 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -128
}

func main() {
    var v1 float64 = 1e10
    var v2 float64 = v1
    var v3 int64 = 1
    var v4 int8 = -128
    switch 127 {
    case 905:
        switch "hello" {
        case "hello":
            var v8 float32 = 0.0
            _ = v8
        default:
            v9 := GenS1{F0: ("" + "foo bar")}
            _ = v9
        }
    case 32767:
        if false {
            v2 = v2
            v10 := GenS1{F0: ("a" + "\n\t")}
            _ = v10
        } else {
            if (38 == 65535) {
                v1 += ((v2 * 1.5) * (v2 * (3.14159 * (v1 + (float64(1e10) + (((v1 + ((v1 * (-2.25 / v1)) - ((v2 * -2.25) - 1e10))) * v2) - (v2 / v1)))))))
                v4 = -12
                v1 = 1e10
            }
            v11 := GenS2{F0: 959, F1: -832}
            v12 := GenS2{F0: 23, F1: 1}
            _ = v11
            _ = v12
        }
    case 128:
        switch "hello" {
        case "":
            for fi13 := 0; fi13 < 4; fi13++ {
                for fi14 := 0; fi14 < 4; fi14++ {
                    var v15 uint8 = 2
                    if false {
                        v16 := GenS1{F0: "foo bar"}
                        if ((v3 + v3) != v3) {
                            v17 := []int{fi14, -2147483648, (fi13 + -483)}
                            var v18 float32 = 3.14159
                            _ = v17
                            _ = v18
                        } else {
                            var v19 string = "hello"
                            v3 += v3
                            _ = v19
                        }
                        if (true && true) {
                            fmt.Println(v15)
                        }
                        _ = v16
                    } else {
                        var v20 float64 = ((v1 + 1e10) - (0.1 * 3.14159))
                        var v21 float32 = float32(v1)
                        if true {
                            if true {
                                fmt.Println(v1, v2, v21)
                            } else {
                                var v22 uint8 = (^(v15))
                                _ = v22
                            }
                        } else {
                            var v23 uint64 = 816
                            _ = v23
                        }
                        v3 += v3
                        var v24 uint64 = 1
                        v21 = (3.14159 / (-2.25 / (float32(1e10) * v21)))
                        var v25 int = genFn0(("hello" + "foo bar"), -2147483648, (fi14 + 32767))
                        _ = v20
                        _ = v21
                        _ = v24
                        _ = v25
                    }
                    _ = fi14
                    _ = v15
                }
                v26 := GenS3{F0: 4294967295}
                _ = fi13
                _ = v26
            }
        case "foo bar":
            v27 := GenS1{F0: (((("foo bar" + "\n\t") + "a") + ("foo bar" + (((("foo bar" + "hello") + "a") + "foo bar") + (("a" + "") + "hello")))) + "foo bar")}
            _ = v27
        case "hello":
            var v28 uint32 = 4294967295
            _ = v28
        }
    default:
        var v29 string = ("a" + "")
        _ = v29
    }
    fmt.Println(v1, v4)
    switch "a" {
    case "a":
        for fi30 := 0; fi30 < 4; fi30++ {
            var v31 string = "foo bar"
            _ = fi30
            _ = v31
        }
    case "foo bar":
        if false {
            if ((-384 >= (v3 + 42)) || (!false)) {
                var v32 uint = 0
                _ = v32
            }
            var v33 uint16 = 32767
            _ = v33
        }
    }
    fmt.Println(v3)
    if true {
        var v34 uint64 = 1
        var v35 uint16 = 32767
        _ = v34
        _ = v35
    } else {
        switch "\n\t" {
        case "\n\t":
            v36 := []int{}
            _ = v36
        case "foo bar":
            v3 += 9223372036854775807
        default:
            v37 := []int{690, -1, 9223372036854775807, 1}
            _ = v37
        }
        var v38 string = ("hello" + "hello")
        var v39 int8 = (v4 / ((genFn1(false, "", 127)) | 1))
        _ = v38
        _ = v39
    }
    v1 -= ((0.1 * v1) / float64(1.5))
    v40 := GenS4{F0: 255}
    v41 := []int{}
    var v42 string = ("\n\t" + "hello")
    v4 -= (v4 * v4)
    fmt.Println(v3, v4, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
}
