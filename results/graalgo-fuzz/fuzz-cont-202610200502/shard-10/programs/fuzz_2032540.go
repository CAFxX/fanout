package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 bool }
type GenS2 struct { F0 uint32; F1 uint }

func main() {
    var v1 uint = 892
    var v2 int32 = 715
    var v3 int8 = 1
    v4 := make(map[uint32]bool)
    var v5 uint64 = 213
    v2 = (v2 & 42)
    v6 := GenS1{F0: 42, F1: 0, F2: true}
    fmt.Println(v2)
    fmt.Println(v2)
    v2 = (v2 % (((-(v2))) | 1))
    v7 := GenS1{F0: 0, F1: v3, F2: true}
    v5 = 1
    v8 := GenS1{F0: 991, F1: v3, F2: false}
    for fi9 := 0; fi9 < 13; fi9++ {
        for fi10 := 0; fi10 < 8; fi10++ {
            v4[1] = ((true && true) || (false && (false && false)))
            _ = fi10
        }
        _ = fi9
    }
    v3 = 127
    switch -356 {
    case 42:
        var v11 int64 = 42
        _ = v11
    case 128:
        if (true && true) {
            v12 := []int{-727, 4294967295}
            var v13 uint16 = 42
            var v14 uint8 = uint8(v13)
            _ = v12
            _ = v13
            _ = v14
        }
    }
    v15 := []int{459}
    if (-2.25 <= float64(1e10)) {
        var v16 string = "\n\t"
        for fi17 := 0; fi17 < 15; fi17++ {
            fmt.Println(v5, v5, v16)
            switch (fi17 << ((2) & 63)) {
            case 32768:
                for fi18 := 0; fi18 < 11; fi18++ {
                    if (838 <= 0) {
                        var v19 uint = (v1 / ((v1) | 1))
                        if false {
                            v16 = (v16 + v16)
                            if false {
                                _ = len(v4)
                                var v20 uint = v1
                                var v21 uint32 = 42
                                _ = v20
                                _ = v21
                            }
                            fmt.Println(v19, v3, v16)
                            if false {
                                v2 -= v2
                                v5 = (v5 & (^(v5)))
                            }
                            if true {
                                fmt.Println(fi18, v5)
                            }
                        } else {
                            if false {
                                v5 += (v5 ^ v5)
                                v16 = ((v16 + ("foo bar" + v16)) + "")
                            }
                        }
                        v5 = (^(v5))
                        _ = v19
                    }
                    fmt.Println(v1)
                    _ = fi18
                }
            case 4294967295:
                v2 -= v2
            }
            _ = fi17
        }
        switch (("" + (v16 + "hello")) + (v16 + "hello")) {
        case "hello":
            var v22 int8 = v3
            _ = v22
        case "\n\t":
            if ((false || (114 > 0)) && true) {
                switch ((("" + v16) + "hello") + "") {
                case "":
                    var v23 int8 = (v3 / ((int8(v5)) | 1))
                    _ = v23
                case "\n\t":
                    v15 = append(v15, int(v3))
                case "foo bar":
                    switch int(v2) {
                    case -612:
                        var v24 int16 = -32768
                        _ = v24
                    case 2147483647:
                        var v25 uint64 = (v5 + v5)
                        _ = v25
                    default:
                        var v26 uint8 = 128
                        _ = v26
                    }
                default:
                    var v27 int32 = 42
                    _ = v27
                }
            }
        case "foo bar":
            v28 := GenS2{F0: 1, F1: v1}
            _ = v28
        default:
            var v29 float64 = float64(float32(3.14159))
            _ = v29
        }
        _ = v16
    } else {
        for fi30 := 0; fi30 < 10; fi30++ {
            fmt.Println(v1)
            _ = fi30
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
