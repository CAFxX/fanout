package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 bool; F1 int64; F2 uint16 }
type GenS3 struct { F0 uint8; F1 uint16 }
func genFn0(p0 string, p1 uint16, p2 uint8) uint64 {
    switch p0 {
    case "\n\t":
        var v5 int64 = 9223372036854775807
        _ = v5
    case "a":
        v6 := []int{0, -2147483648}
        _ = v6
    case "foo bar":
        fmt.Println(p2)
    default:
        if false {
            if (!false) {
                if false {
                    if false {
                        v7 := []int{42, -56}
                        fmt.Println(p0, p1, p2)
                        _ = v7
                    }
                } else {
                    v8 := []int{-9223372036854775808}
                    _ = v8
                }
                for fi9 := 0; fi9 < 12; fi9++ {
                    p1 = (^(p1))
                    v10 := GenS1{F0: "a", F1: 0}
                    var v11 uint8 = 205
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            } else {
                var v12 uint64 = 18446744073709551615
                v13 := GenS2{F0: false, F1: 9223372036854775807, F2: p1}
                _ = v12
                _ = v13
            }
            v14 := []int{-640, 4294967295}
            for fi15 := 0; fi15 < 2; fi15++ {
                p0 = ("hello" + (p0 + ""))
                p1 = p1
                var v16 int = fi15
                _ = fi15
                _ = v16
            }
            _ = v14
        }
    }
    p1 = 0
    _ = p0
    _ = p1
    _ = p2
    return 2
}
func genFn1() bool {
    var v17 float64 = 0.1
    fmt.Println(v17)
    _ = v17
    return false
}

func main() {
    var v1 float32 = 1.5
    var v2 uint32 = 931
    var v3 uint64 = 4294967295
    var v4 float32 = ((float32(1.5) * v1) * (-2.25 / 1.0))
    fmt.Println(v4, v4)
    switch 897 {
    case -9223372036854775808:
        var v18 uint = 621
        _ = v18
    case -250:
        if true {
            v2 = (v2 + (^(v2)))
            fmt.Println(v1)
        } else {
            var v19 int64 = -568
            for fi20 := 0; fi20 < 14; fi20++ {
                for fi21 := 0; fi21 < 5; fi21++ {
                    v19 = v19
                    v22 := []int{(^(fi20)), fi20, 128}
                    var v23 string = "foo bar"
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                switch "\n\t" {
                case "":
                    var v24 int = (fi20 / ((fi20) | 1))
                    _ = v24
                case "\n\t":
                    if (!(!(true && false))) {
                        v25 := []int{fi20}
                        var v26 int32 = 2147483647
                        var v27 uint64 = 18446744073709551615
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                case "a":
                    v2 += (v2 / ((0) | 1))
                default:
                    fmt.Println(v19)
                }
                _ = fi20
            }
            _ = v19
        }
    default:
        v2 += (v2 * (v2 + v2))
    }
    var v28 uint32 = v2
    if true {
        v3 = (^(v3))
        var v29 int32 = -1
        _ = v29
    }
    v30 := []int{0, 4294967295}
    v28 -= v28
    if false {
        fmt.Println(v28, v4, v1)
        v2 += (v2 + 632)
        var v31 float32 = v4
        _ = v31
    }
    fmt.Println(v2, v3, v4)
    var v32 uint32 = (v2 ^ (v2 * uint32(v3)))
    switch ("" + "") {
    case "a":
        var v33 uint = 732
        _ = v33
    case "\n\t":
        var v34 int8 = -128
        _ = v34
    default:
        v4 = (v4 * float32((0.1 * (float64((((v1 * (((v1 + v4) * (float32(0.1) + ((v1 * -2.25) * v1))) / 0.1)) - v4) * v1)) + 3.14159))))
    }
    v1 = v4
    v35 := GenS2{F0: false, F1: 42, F2: 191}
    for fi36 := 0; fi36 < 8; fi36++ {
        fmt.Println(fi36, v32)
        if len(v30) > 0 {
            v30[0] = (^(fi36))
        }
        _ = fi36
    }
    if true {
        if ("" == ("\n\t" + "hello")) {
            for fi37 := 0; fi37 < 19; fi37++ {
                for fi38 := 0; fi38 < 16; fi38++ {
                    fmt.Println(v3, v3, fi38)
                    _ = fi38
                }
                if (!true) {
                    fmt.Println(v2, v3, v3)
                }
                _ = len(v30)
                _ = fi37
            }
            var v39 uint32 = 2
            _ = v39
        } else {
            v30 = append(v30, 249)
            if (int8(v28) == 80) {
                var v40 float64 = float64(-2.25)
                _ = v40
            }
            var v41 int64 = 1
            _ = v41
        }
        if (-168 < -390) {
            v28 = v28
        } else {
            var v42 bool = false
            _ = v42
        }
        v43 := GenS3{F0: 29, F1: 752}
        _ = v43
    } else {
        for fi44 := 0; fi44 < 3; fi44++ {
            var v45 string = "a"
            var v46 string = v45
            _ = fi44
            _ = v45
            _ = v46
        }
        for fi47 := 0; fi47 < 16; fi47++ {
            v30 = append(v30, fi47)
            _ = fi47
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v32
    _ = v35
    _ = v4
}
