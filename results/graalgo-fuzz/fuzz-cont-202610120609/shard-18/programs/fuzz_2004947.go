package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 string }
type GenS2 struct { F0 uint32; F1 bool }
func genFn0() uint16 {
    v6 := GenS1{F0: 18446744073709551615, F1: 32767, F2: ""}
    v7 := GenS1{F0: 42, F1: 1, F2: "a"}
    _ = v6
    _ = v7
    return 32767
}
func genFn1(p0 int32, p1 string, p2 int16) int16 {
    switch -9223372036854775808 {
    case 127:
        switch "a" {
        case "":
            v8 := []int{}
            _ = v8
        case "\n\t":
            p1 = ("" + p1)
        case "a":
            p2 += 32767
        }
    case -9223372036854775808:
        p0 += -307
    default:
        for fi9 := 0; fi9 < 15; fi9++ {
            for fi10 := 0; fi10 < 17; fi10++ {
                p2 = p2
                var v11 int8 = -118
                _ = fi10
                _ = v11
            }
            _ = fi9
        }
    }
    var v12 uint64 = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return -1
}

func main() {
    v1 := []uint32{}
    var v2 uint16 = 903
    var v3 int8 = 100
    v4 := make(map[string]int)
    v5 := []uint32{4294967295, 405}
    var v13 uint32 = 2
    var v14 bool = (v3 == v3)
    for fi15 := 0; fi15 < 13; fi15++ {
        fmt.Println(v13, v3)
        _ = fi15
    }
    var v16 uint32 = v13
    if false {
        fmt.Println(v13, v3, v14)
        for fi17 := 0; fi17 < 15; fi17++ {
            fmt.Println(v14, v14, v2)
            v14 = true
            _ = fi17
        }
    }
    v2 += genFn0()
    if (1e10 < 0.1) {
        v18 := GenS1{F0: 42, F1: (v2 * v2), F2: ("hello" + "\n\t")}
        _ = v18
    }
    var v19 string = "hello"
    if v14 {
        var v20 uint = 2
        _ = v20
    } else {
        v13 += v13
        if (v14 || v14) {
            switch v19 {
            case "foo bar":
                v2 = v2
            case "a":
                v21 := v4[v19]
                _ = v21
            case "":
                v22 := GenS1{F0: 1, F1: (v2 / ((v2) | 1)), F2: "hello"}
                _ = v22
            }
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v23 := GenS1{F0: 881, F1: genFn0(), F2: v19}
            v24 := GenS1{F0: 135, F1: 2, F2: v19}
            _ = v23
            _ = v24
        }
    }
    if len(v1) > 0 {
        v1[0] = 1
    }
    var v25 int16 = int16(v13)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi26 := 0; fi26 < 3; fi26++ {
        if v14 {
            delete(v4, "foo bar")
            for fi27 := 0; fi27 < 10; fi27++ {
                if v14 {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                } else {
                    v28 := GenS1{F0: 42, F1: v2, F2: "a"}
                    var v29 int = 764
                    var v30 int = v29
                    v31 := GenS1{F0: 0, F1: (v2 % ((42) | 1)), F2: (((v19 + (v19 + v19)) + v19) + (("hello" + v19) + v19))}
                    _ = v28
                    _ = v29
                    _ = v30
                    _ = v31
                }
                v19 = "hello"
                _ = fi27
            }
        } else {
            switch (fi26 & (fi26 >> ((0) & 63))) {
            case -707:
                if v14 {
                    v25 = v25
                }
            default:
                fmt.Println(v14, v3, v16)
            }
        }
        v32 := GenS1{F0: 642, F1: v2, F2: ""}
        var v33 int64 = 0
        _ = fi26
        _ = v32
        _ = v33
    }
    v34 := GenS2{F0: 2, F1: (!(228 > 42))}
    switch ("hello" + v19) {
    case "hello":
        var v35 uint = 253
        _ = v35
    case "\n\t":
        _ = len(v1)
    }
    for fi36 := 0; fi36 < 4; fi36++ {
        v37 := GenS1{F0: 0, F1: (v2 & genFn0()), F2: v19}
        _ = fi36
        _ = v37
    }
    if (v3 <= (v3 >> ((0) & 63))) {
        v13 += 2
        v25 -= v25
    }
    v3 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
