package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int; F1 uint16 }
type GenS4 struct { F0 uint16; F1 uint64; F2 uint16 }
type GenS5 struct { F0 int; F1 bool; F2 uint32 }
type GenS6 struct { F0 uint8; F1 uint64; F2 uint8 }
func genFn0(p0 uint16, p1 uint) uint32 {
    v6 := []int{-643, 32768, -50, -244}
    var v7 float64 = -2.25
    p0 += (p0 ^ p0)
    switch ((("" + ("" + "foo bar")) + (("" + ("foo bar" + "hello")) + "\n\t")) + "\n\t") {
    case "a":
        var v8 float32 = 1.5
        _ = v8
    default:
        for fi9 := 0; fi9 < 14; fi9++ {
            var v10 int8 = 0
            _ = fi9
            _ = v10
        }
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 390
}

func main() {
    var v1 uint64 = 299
    v2 := make(map[string]string)
    var v3 uint64 = v1
    v4 := []uint32{4294967295, 196}
    var v5 uint32 = 65535
    if false {
        if ((!((v3 / ((v1) | 1)) <= 490)) && true) {
            v11 := GenS1{F0: v3, F1: v5, F2: true}
            _ = v11
        } else {
            var v12 uint = 65535
            if true {
                if false {
                    v13 := v2[(("a" + "") + "a")]
                    fmt.Println(v3, v13, v3)
                    _ = len(v2)
                    _ = v13
                }
                fmt.Println(v1)
                v3 += (^(v1))
            }
            _ = v12
        }
        var v14 string = ("" + "a")
        _ = v14
    } else {
        v1 = (v1 << ((1) & 63))
        if len(v4) > 0 {
            v4[0] = genFn0(70, 42)
        }
        v3 -= v1
    }
    v15 := GenS2{F0: "hello"}
    var v16 int8 = 42
    v17 := GenS1{F0: 375, F1: (^(v5)), F2: (false || (v16 == (v16 & (v16 >> ((1) & 63)))))}
    v18 := GenS3{F0: -9223372036854775808, F1: 255}
    if (false || (1 < uint8(v5))) {
        v19 := GenS4{F0: 42, F1: 898, F2: 552}
        v16 = v16
        v20 := GenS4{F0: 1, F1: 4294967295, F2: 1}
        _ = v19
        _ = v20
    }
    v2["a"] = ((("\n\t" + "hello") + ("foo bar" + (("a" + (("hello" + "hello") + ("" + ((("foo bar" + ("hello" + ((("a" + ((("a" + "\n\t") + "hello") + ("foo bar" + "foo bar"))) + "hello") + "hello"))) + ("\n\t" + "a")) + ("" + "\n\t"))))) + "foo bar"))) + ("" + ("\n\t" + "foo bar")))
    switch 0 {
    case 128:
        v5 = (v5 % ((v5) | 1))
    case -1:
        var v21 bool = false
        _ = v21
    }
    if (v16 == v16) {
        if true {
            var v22 int64 = 545
            _ = v22
        }
        switch "\n\t" {
        case "hello":
            fmt.Println(v5, v3)
        case "foo bar":
            v23 := v2[("hello" + "\n\t")]
            _ = v23
        default:
            if true {
                switch 4294967295 {
                case -448:
                    var v24 int8 = v16
                    _ = v24
                case -366:
                    switch 32768 {
                    case 4294967295:
                        v25 := GenS5{F0: 32768, F1: false, F2: (v5 << ((18446744073709551615) & 63))}
                        _ = v25
                    }
                }
            }
        }
    } else {
        v26 := v2["a"]
        fmt.Println(v26)
        _ = v26
    }
    v5 = 1
    var v27 bool = true
    v28 := GenS6{F0: 127, F1: v3, F2: 32}
    for fi29 := 0; fi29 < 19; fi29++ {
        for fi30 := 0; fi30 < 14; fi30++ {
            if len(v4) > 0 {
                v4[0] = genFn0(1, 2)
            }
            v31 := GenS4{F0: uint16(v1), F1: (v3 / ((v1) | 1)), F2: 255}
            _ = fi30
            _ = v31
        }
        v32 := GenS3{F0: fi29, F1: 0}
        for fi33 := 0; fi33 < 16; fi33++ {
            v3 = v3
            _ = fi33
        }
        _ = fi29
        _ = v32
    }
    var v34 uint64 = v1
    v35 := GenS1{F0: 0, F1: 65535, F2: false}
    for fi36 := 0; fi36 < 1; fi36++ {
        v4 = append(v4, 1)
        delete(v2, "foo bar")
        _ = fi36
    }
    var v37 bool = v27
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v34
    _ = v35
    _ = v37
    _ = v4
    _ = v5
}
