package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 uint; F1 int16; F2 bool }
type GenS3 struct { F0 int64; F1 int32 }
type GenS4 struct { F0 int; F1 uint8 }
func genFn0(p0 int, p1 uint64, p2 int8) int32 {
    if (!false) {
        if ("" >= "\n\t") {
            for fi5 := 0; fi5 < 4; fi5++ {
                switch p0 {
                case -1:
                    if false {
                        var v6 uint8 = 42
                        for fi7 := 0; fi7 < 20; fi7++ {
                            p1 -= p1
                            v8 := []int{(fi7 / ((fi5) | 1)), (p0 | fi7), fi5}
                            v6 += 134
                            _ = fi7
                            _ = v8
                        }
                        _ = v6
                    }
                case 1:
                    p2 -= (p2 ^ p2)
                }
                if (true || true) {
                    p1 = (p1 % ((42) | 1))
                }
                v9 := GenS1{F0: ("" + ("foo bar" + "")), F1: 42}
                _ = fi5
                _ = v9
            }
        }
        if ((1.5 + float64(float32((1e10 + -2.25)))) == 0.0) {
            v10 := GenS2{F0: 18446744073709551615, F1: 1000, F2: true}
            if true {
                fmt.Println(p0, p1, p1)
                var v11 string = ("\n\t" + "foo bar")
                var v12 int16 = -32768
                _ = v11
                _ = v12
            } else {
                for fi13 := 0; fi13 < 19; fi13++ {
                    v14 := GenS1{F0: ("a" + ("" + "")), F1: 1}
                    fmt.Println(p2, p0, fi13)
                    switch (fi13 | (p0 * p0)) {
                    case 32768:
                        if false {
                            p2 -= (p2 * (p2 / ((p2) | 1)))
                            var v15 bool = (65535 > 4294967295)
                            p2 -= p2
                            _ = v15
                        } else {
                            for fi16 := 0; fi16 < 4; fi16++ {
                                v17 := []int{0, fi13, -9223372036854775808, (fi16 | fi16)}
                                v18 := GenS2{F0: 693, F1: -32768, F2: (-1 > -1)}
                                var v19 bool = (!(1 != -1))
                                _ = fi16
                                _ = v17
                                _ = v18
                                _ = v19
                            }
                            p1 = p1
                        }
                    }
                    _ = fi13
                    _ = v14
                }
                v20 := GenS2{F0: 898, F1: 1, F2: true}
                p0 += -1
                _ = v20
            }
            _ = v10
        } else {
            v21 := GenS1{F0: ("\n\t" + ("" + (("foo bar" + "a") + ("hello" + "\n\t")))), F1: 2}
            p1 -= p1
            _ = v21
        }
        v22 := []int{p0, 9223372036854775807, p0, p0}
        _ = v22
    }
    var v23 int8 = (p2 % ((127) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    return 1
}

func main() {
    var v1 int64 = -1
    v2 := make(map[string]bool)
    v3 := []int{42, -1, -293, 9223372036854775807}
    var v4 uint16 = 42
    v24 := GenS2{F0: 42, F1: -32768, F2: true}
    v4 = (^(v4))
    var v25 uint64 = 1
    v4 = (^(v4))
    _ = len(v2)
    if (true || true) {
        v25 = v25
        fmt.Println(v4)
    } else {
        v26 := GenS3{F0: v1, F1: 1}
        _ = v26
    }
    if (false || ((!(!((true && false) || false))) && false)) {
        v27 := GenS4{F0: -1, F1: uint8(v4)}
        var v28 int16 = 1000
        v29 := GenS1{F0: "", F1: 1}
        _ = v27
        _ = v28
        _ = v29
    } else {
        switch "hello" {
        case "a":
            v2[""] = ("a" <= "a")
        }
        fmt.Println(v25, v1)
        var v30 uint = uint(v4)
        _ = v30
    }
    switch "foo bar" {
    case "foo bar":
        fmt.Println(v1, v25, v25)
    case "hello":
        var v31 string = ("hello" + "")
        _ = v31
    default:
        v32 := GenS3{F0: (v1 + v1), F1: 1}
        _ = v32
    }
    v25 = v25
    v25 = (v25 / ((v25) | 1))
    var v33 int32 = 454
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v33
    _ = v4
}
