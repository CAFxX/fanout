package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16; F1 uint64; F2 uint }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 bool; F1 int8; F2 int32 }
type GenS7 struct { F0 bool }

func main() {
    var v1 uint32 = 259
    var v2 int64 = 42
    v3 := make(map[int]uint32)
    var v4 int16 = -203
    v3[4294967295] = (v1 * v1)
    for fi5 := 0; fi5 < 5; fi5++ {
        v2 = v2
        for fi6 := 0; fi6 < 17; fi6++ {
            var v7 float32 = (-2.25 + (0.0 + (3.14159 - 0.1)))
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    v2 += v2
    v1 -= v1
    var v8 int16 = (v4 + (v4 ^ (v4 | v4)))
    v2 = v2
    switch (("a" + "") + (("\n\t" + ("a" + ("" + "hello"))) + "foo bar")) {
    case "a":
        delete(v3, 69)
    case "hello":
        for fi9 := 0; fi9 < 12; fi9++ {
            if false {
                fmt.Println(v2, fi9, fi9)
                fmt.Println(v2)
            } else {
                if (("hello" > "") && true) {
                    switch fi9 {
                    case 32767:
                        var v10 bool = ((v2 != v2) || (211 == 127))
                        _ = v10
                    default:
                        var v11 float32 = 3.14159
                        _ = v11
                    }
                }
            }
            _ = fi9
        }
    }
    if true {
        var v12 int8 = 124
        if true {
            v2 = (v2 & v2)
        }
        _ = v12
    }
    if false {
        v13 := GenS1{F0: 1}
        var v14 int32 = int32(v1)
        var v15 float64 = 1.5
        _ = v13
        _ = v14
        _ = v15
    }
    var v16 uint = 838
    v17 := []int{32768}
    if false {
        if (((v2 >> (((^(v16))) & 63)) >= (v2 ^ v2)) && false) {
            if false {
                delete(v3, -415)
            } else {
                var v18 float32 = float32(1.5)
                v19 := GenS2{F0: (^(v1))}
                fmt.Println(v2, v2, v8)
                _ = v18
                _ = v19
            }
        } else {
            var v20 float32 = float32(1.5)
            var v21 float32 = v20
            _ = v20
            _ = v21
        }
        var v22 int32 = 0
        _ = v22
    }
    v23 := GenS2{F0: 65535}
    var v24 string = (("foo bar" + "foo bar") + "hello")
    if (!(true || ((v2 + (v2 - v2)) != (^(v2))))) {
        for fi25 := 0; fi25 < 12; fi25++ {
            var v26 int8 = int8(v2)
            switch (fi25 / (((^(fi25))) | 1)) {
            case -9223372036854775808:
                v8 = v4
            }
            _ = len(v3)
            _ = fi25
            _ = v26
        }
        for fi27 := 0; fi27 < 14; fi27++ {
            switch fi27 {
            case 128:
                var v28 int = fi27
                _ = v28
            case 32767:
                if ((false && (!false)) && (!(73 > 77))) {
                    v8 = v4
                    switch fi27 {
                    case -86:
                        v29 := GenS3{F0: 65535, F1: uint64(v1), F2: (v16 - (^(v16)))}
                        _ = v29
                    }
                    v30 := GenS4{F0: uint(v1)}
                    _ = v30
                } else {
                    var v31 bool = false
                    switch fi27 {
                    case 32768:
                        var v32 uint = 0
                        _ = v32
                    case 128:
                        v33 := GenS5{F0: v16}
                        _ = v33
                    case -887:
                        v34 := GenS6{F0: true, F1: -37, F2: 1}
                        _ = v34
                    default:
                        v35 := GenS7{F0: (("hello" + (v24 + "hello")) >= v24)}
                        _ = v35
                    }
                    v1 -= (v1 / (((^(v1))) | 1))
                    _ = v31
                }
            }
            _ = fi27
        }
        delete(v3, 2147483647)
    } else {
        v4 = (^(v8))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v8
}
