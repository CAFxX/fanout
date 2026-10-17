package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 uint64; F1 int32; F2 int32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 string; F1 int8; F2 int }
func genFn0(p0 uint16, p1 string, p2 int) int {
    if (974 != 947) {
        var v6 uint8 = 2
        _ = v6
    } else {
        v7 := GenS1{F0: -1, F1: 127}
        var v8 uint16 = (p0 ^ (p0 << ((4294967295) & 63)))
        v9 := GenS1{F0: 2147483647, F1: 128}
        _ = v7
        _ = v8
        _ = v9
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    return 270
}
func genFn1(p0 int) string {
    p0 -= (p0 - (p0 - (p0 | genFn0(944, "a", p0))))
    v10 := GenS1{F0: -2147483648, F1: 255}
    _ = p0
    _ = v10
    return "a"
}

func main() {
    var v1 bool = (!(true && false))
    var v2 string = ("foo bar" + "a")
    var v3 bool = ((-1 == 9223372036854775807) && false)
    var v4 int8 = 0
    var v5 bool = true
    switch ("" + v2) {
    case "\n\t":
        var v11 uint8 = 1
        _ = v11
    case "a":
        var v12 string = v2
        _ = v12
    }
    var v13 bool = false
    v13 = false
    for fi14 := 0; fi14 < 2; fi14++ {
        v13 = false
        if ((false && v13) && v1) {
            switch ("\n\t" + v2) {
            case "":
                v15 := GenS1{F0: 315, F1: 2}
                _ = v15
            case "hello":
                var v16 int64 = 434
                _ = v16
            case "a":
                if true {
                    v3 = (!v13)
                    v2 = (("\n\t" + v2) + v2)
                    v4 = -128
                } else {
                    for fi17 := 0; fi17 < 20; fi17++ {
                        if true {
                            if (!v3) {
                                switch v2 {
                                case "foo bar":
                                    var v18 int16 = 0
                                    _ = v18
                                }
                                fmt.Println(v2)
                                v19 := []int{(^(fi17)), (fi14 & fi17), (fi17 | fi14), genFn0(0, "foo bar", (-(fi17)))}
                                _ = v19
                            } else {
                                var v20 string = v2
                                if (2 == 18) {
                                    switch 340 {
                                    case -1:
                                        var v21 int16 = -32768
                                        _ = v21
                                    case -122:
                                        switch fi14 {
                                        case 1:
                                            var v22 bool = v1
                                            _ = v22
                                        }
                                    }
                                    v20 = v20
                                    if (!(genFn0(2, v2, (fi17 - (^(fi17)))) <= 2147483647)) {
                                        if (false || (!((3.14159 / float64(-2.25)) < 0.1))) {
                                            fmt.Println(fi14)
                                            fmt.Println(v4, v3, v1)
                                        } else {
                                            v13 = (42 != 255)
                                            if ((!v1) && (!true)) {
                                                switch v20 {
                                                case "hello":
                                                    v23 := GenS2{F0: v2, F1: v13}
                                                    _ = v23
                                                }
                                            }
                                        }
                                        v24 := GenS3{F0: 2, F1: 42, F2: -485}
                                        _ = v24
                                    }
                                }
                                _ = v20
                            }
                            var v25 int64 = 831
                            _ = v25
                        } else {
                            v26 := []int{-351, fi14}
                            fmt.Println(fi17, v1, v3)
                            _ = v26
                        }
                        var v27 uint64 = 2
                        v28 := GenS4{F0: 4294967295}
                        _ = fi17
                        _ = v27
                        _ = v28
                    }
                    switch genFn0(255, v2, 4294967295) {
                    case 32768:
                        fmt.Println(v2, v3, v4)
                    }
                }
            default:
                v2 = (v2 + "foo bar")
            }
            var v29 float64 = (float64(0.1) + float64(1.5))
            v30 := []int{fi14, genFn0(255, "hello", fi14)}
            _ = v29
            _ = v30
        }
        var v31 uint8 = 1
        _ = fi14
        _ = v31
    }
    v32 := GenS2{F0: "foo bar", F1: v13}
    var v33 float64 = float64(-2.25)
    for fi34 := 0; fi34 < 12; fi34++ {
        var v35 int8 = (v4 - -1)
        v36 := GenS2{F0: v2, F1: false}
        switch fi34 {
        case 127:
            v37 := []int{(fi34 ^ fi34), fi34, genFn0(0, v2, 39), 32767}
            _ = v37
        case 9223372036854775807:
            v38 := GenS4{F0: 343}
            _ = v38
        case 2147483647:
            v39 := []int{fi34}
            _ = v39
        }
        _ = fi34
        _ = v35
        _ = v36
    }
    v40 := GenS5{F0: "hello", F1: (v4 + 0), F2: -9223372036854775808}
    var v41 string = ((v2 + "hello") + (((("hello" + "a") + v2) + v2) + v2))
    v4 = v4
    var v42 float32 = (((-2.25 + (0.0 + 1e10)) / (-2.25 / 1.5)) * (-2.25 - (float32(1.5) * 0.1)))
    v5 = (!(float64(v42) == 1e10))
    v42 = v42
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
}
