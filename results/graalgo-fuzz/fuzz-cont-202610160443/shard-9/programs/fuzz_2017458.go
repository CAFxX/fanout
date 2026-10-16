package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 uint8; F2 uint32 }

func main() {
    var v1 uint = 897
    var v2 int = 201
    v3 := make(map[uint32]int)
    v4 := make(map[int]uint32)
    v5 := GenS1{F0: 2, F1: (((((("a" + ("\n\t" + (("" + "a") + ""))) + "foo bar") + "a") + ("foo bar" + "hello")) + (((("a" + "hello") + "") + "") + (("a" + "foo bar") + (("a" + ("\n\t" + "foo bar")) + "\n\t")))) + "a"), F2: (v2 ^ (v2 | (v2 + 763)))}
    v6 := []int{v2}
    var v7 int64 = 1
    var v8 int32 = -607
    if false {
        fmt.Println(v8, v2)
    } else {
        var v9 float32 = 0.0
        var v10 int = (v2 << (((v1 - v1)) & 63))
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 17; fi11++ {
        switch ("" + "foo bar") {
        case "":
            v7 += -9223372036854775808
        }
        var v12 int8 = 127
        _ = fi11
        _ = v12
    }
    v13 := GenS1{F0: 2147483648, F1: "\n\t", F2: 236}
    switch ("\n\t" + "a") {
    case "a":
        var v14 string = "hello"
        _ = v14
    case "foo bar":
        if true {
            v15 := GenS2{F0: -1}
            switch v2 {
            case 127:
                for fi16 := 0; fi16 < 1; fi16++ {
                    switch ((("hello" + (("" + "") + "\n\t")) + (("" + (("" + "\n\t") + "a")) + "hello")) + ("" + "hello")) {
                    case "foo bar":
                        var v17 int16 = 32767
                        _ = v17
                    case "\n\t":
                        var v18 int64 = (v7 - v7)
                        _ = v18
                    case "hello":
                        if true {
                            v2 += int(v7)
                            v19 := GenS3{F0: ((true && true) && (0 < (v1 & (v1 << ((42) & 63))))), F1: 128, F2: 515}
                            v20 := GenS2{F0: -79}
                            _ = v19
                            _ = v20
                        } else {
                            var v21 int = fi16
                            for fi22 := 0; fi22 < 18; fi22++ {
                                v23 := GenS1{F0: 1, F1: "hello", F2: -2147483648}
                                _ = fi22
                                _ = v23
                            }
                            _ = v21
                        }
                    default:
                        v7 -= (-(v7))
                    }
                    var v24 uint16 = 2
                    var v25 int64 = (v7 | v7)
                    _ = fi16
                    _ = v24
                    _ = v25
                }
            case 854:
                var v26 uint = v1
                _ = v26
            case 9223372036854775807:
                var v27 float64 = ((0.1 / (0.0 - (0.1 / ((((0.0 - 0.1) - 0.1) + 3.14159) * 3.14159)))) + float64(0.0))
                _ = v27
            }
            switch int(v8) {
            case 1:
                var v28 float32 = (-2.25 + 0.1)
                _ = v28
            }
            _ = v15
        }
    case "hello":
        _ = len(v6)
    default:
        v2 = (v2 / ((-1) | 1))
    }
    for fi29 := 0; fi29 < 9; fi29++ {
        for fi30 := 0; fi30 < 4; fi30++ {
            fmt.Println(v8, v7, v1)
            _ = fi30
        }
        v1 = (v1 + (v1 ^ (^(v1))))
        _ = fi29
    }
    v31 := GenS2{F0: -128}
    if false {
        delete(v3, 2147483648)
    }
    v2 -= v2
    v1 = v1
    v7 = v7
    switch (v2 & v2) {
    case 128:
        fmt.Println(v7, v2)
    default:
        v32 := GenS2{F0: -1}
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
