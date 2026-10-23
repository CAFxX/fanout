package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string; F1 string }
func genFn0(p0 uint16) int64 {
    for fi7 := 0; fi7 < 2; fi7++ {
        p0 = 65535
        _ = fi7
    }
    var v8 uint8 = 1
    fmt.Println(p0, v8)
    if (false || true) {
        v9 := GenS1{F0: (!true), F1: 95}
        _ = v9
    } else {
        fmt.Println(p0)
        var v10 uint = 18446744073709551615
        for fi11 := 0; fi11 < 1; fi11++ {
            v12 := []int{42, (fi11 ^ 773)}
            _ = fi11
            _ = v12
        }
        _ = v10
    }
    var v13 uint8 = v8
    _ = p0
    _ = v13
    _ = v8
    return -333
}

func main() {
    v1 := []string{"a"}
    var v2 int = -1
    v3 := make(map[uint32]bool)
    v4 := []uint32{2, 893}
    var v5 int8 = 69
    var v6 bool = (-191 >= 970)
    v14 := GenS1{F0: v6, F1: 409}
    for fi15 := 0; fi15 < 2; fi15++ {
        if false {
            v16 := GenS2{F0: "a"}
            v2 = (fi15 ^ fi15)
            delete(v3, 29)
            _ = v16
        } else {
            var v17 uint8 = 127
            for fi18 := 0; fi18 < 6; fi18++ {
                v19 := GenS2{F0: (((("\n\t" + "hello") + (("" + "a") + "a")) + "foo bar") + "")}
                if len(v1) > 0 {
                    v1[0] = "a"
                }
                var v20 float64 = -2.25
                v17 += (v17 % ((2) | 1))
                if v6 {
                    var v21 uint64 = 438
                    v3[2] = true
                    _ = v21
                } else {
                    v22 := GenS1{F0: true, F1: 65535}
                    v23 := GenS3{F0: "a", F1: ("a" + "foo bar")}
                    v17 -= 128
                    _ = v22
                    _ = v23
                }
                switch "" {
                case "\n\t":
                    switch "hello" {
                    case "\n\t":
                        v24 := GenS3{F0: "\n\t", F1: ("" + "\n\t")}
                        _ = v24
                    case "hello":
                        v25 := GenS1{F0: true, F1: 2}
                        _ = v25
                    default:
                        var v26 int = 0
                        _ = v26
                    }
                }
                _ = fi18
                _ = v19
                _ = v20
            }
            v27 := GenS3{F0: "\n\t", F1: ("a" + "a")}
            _ = v17
            _ = v27
        }
        var v28 string = "a"
        _ = fi15
        _ = v28
    }
    var v29 int64 = -1
    v29 -= (v29 / (((v29 / (((v29 >> ((4294967295) & 63))) | 1))) | 1))
    var v30 int8 = (-(v5))
    var v31 bool = v6
    for fi32 := 0; fi32 < 17; fi32++ {
        fmt.Println(fi32)
        _ = fi32
    }
    v29 -= -1
    var v33 int = v2
    for fi34 := 0; fi34 < 20; fi34++ {
        _ = len(v1)
        if true {
            v3[75] = false
            var v35 string = "a"
            _ = v35
        }
        _ = fi34
    }
    for fi36 := 0; fi36 < 5; fi36++ {
        v6 = (-1 != -67)
        for fi37 := 0; fi37 < 10; fi37++ {
            fmt.Println(fi36, v30, fi37)
            _ = fi37
        }
        _ = fi36
    }
    var v38 uint8 = 0
    if ((2 >= 42) && true) {
        var v39 int64 = v29
        v5 += 42
        _ = v39
    } else {
        for fi40 := 0; fi40 < 11; fi40++ {
            v41 := GenS2{F0: "foo bar"}
            if len(v4) > 0 {
                v4[0] = 4294967295
            }
            var v42 float32 = 0.1
            _ = fi40
            _ = v41
            _ = v42
        }
        var v43 uint8 = v38
        _ = v43
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v33
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
