package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 string }
type GenS2 struct { F0 int; F1 uint; F2 uint8 }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8; F1 uint64 }
func genFn0(p0 uint, p1 int16, p2 uint) int64 {
    if false {
        var v5 int32 = -1
        p1 -= (p1 % ((1000) | 1))
        _ = v5
    } else {
        fmt.Println(p2, p2)
        v6 := GenS1{F0: p0, F1: p1, F2: ("a" + "a")}
        _ = v6
    }
    switch "foo bar" {
    case "hello":
        p2 += (p0 / ((p2) | 1))
    }
    v7 := []int{128, -557}
    var v8 int8 = 1
    v9 := GenS2{F0: 42, F1: (p0 >> ((881) & 63)), F2: 255}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return -703
}
func genFn1(p0 int16, p1 int64, p2 uint64) int32 {
    var v10 int16 = int16(p1)
    if ((794 > 1) || ("hello" > ("" + "\n\t"))) {
        for fi11 := 0; fi11 < 8; fi11++ {
            for fi12 := 0; fi12 < 3; fi12++ {
                fmt.Println(fi12, fi12, fi12)
                var v13 bool = true
                if ((p1 > (p1 / ((genFn0(65535, -578, 0)) | 1))) || ((487 >= 0) || v13)) {
                    v14 := GenS3{F0: 42, F1: p2}
                    if (1 <= 32767) {
                        switch ("foo bar" + "a") {
                        case "":
                            if (((1e10 / 1.5) + (-2.25 - 1e10)) != 0.1) {
                                if ("\n\t" >= ("\n\t" + "\n\t")) {
                                    var v15 bool = (true || false)
                                    v16 := []int{fi11, 569, fi12}
                                    _ = v15
                                    _ = v16
                                }
                                p1 -= p1
                            }
                        default:
                            v17 := GenS1{F0: 1, F1: v10, F2: "foo bar"}
                            _ = v17
                        }
                        if (-659 > -1) {
                            v10 += v10
                        } else {
                            v18 := GenS2{F0: fi12, F1: 4294967295, F2: 0}
                            v13 = ((253 >= 42) && (-2.25 >= 1.5))
                            p1 -= genFn0(0, (p0 / ((p0) | 1)), 42)
                            _ = v18
                        }
                        var v19 int8 = -128
                        _ = v19
                    }
                    _ = v14
                } else {
                    v20 := []int{fi12, (fi11 ^ fi12), fi11}
                    _ = v20
                }
                _ = fi12
                _ = v13
            }
            _ = fi11
        }
        v21 := GenS2{F0: -2147483648, F1: 1, F2: 125}
        _ = v21
    } else {
        p0 -= 1
    }
    for fi22 := 0; fi22 < 13; fi22++ {
        fmt.Println(v10, p1)
        _ = fi22
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return -1
}

func main() {
    v1 := []string{"a", ((("\n\t" + "a") + ("\n\t" + ("a" + "foo bar"))) + "foo bar"), "\n\t"}
    var v2 int = -441
    v3 := make(map[int]uint32)
    v4 := make(map[string]int)
    fmt.Println(v2)
    v2 += (v2 ^ (v2 / ((v2) | 1)))
    v23 := GenS2{F0: 128, F1: 4294967295, F2: 42}
    if true {
        var v24 uint32 = 724
        var v25 uint64 = 4294967295
        if (110 <= 612) {
            v2 += (v2 * (v2 >> ((18446744073709551615) & 63)))
        } else {
            _ = len(v1)
        }
        _ = v24
        _ = v25
    }
    if (true || false) {
        v26 := GenS4{F0: 1}
        _ = v26
    } else {
        v2 = v2
        v27 := GenS5{F0: 128, F1: 4294967295}
        for fi28 := 0; fi28 < 6; fi28++ {
            v2 = (fi28 ^ fi28)
            _ = fi28
        }
        _ = v27
    }
    var v29 bool = false
    var v30 float32 = (3.14159 - 0.1)
    var v31 int = (v2 + v2)
    v32 := GenS5{F0: 127, F1: 181}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
