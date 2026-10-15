package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint64 }

func main() {
    var v1 int8 = -24
    v2 := make(map[uint32]string)
    var v3 string = ""
    var v4 uint8 = 0
    var v5 int16 = 0
    if false {
        switch v3 {
        case "hello":
            for fi6 := 0; fi6 < 19; fi6++ {
                v5 = (^(v5))
                delete(v2, 42)
                _ = fi6
            }
        case "\n\t":
            v1 = v1
        default:
            if true {
                v2[65535] = (v3 + v3)
                v7 := GenS1{F0: (v4 & 2)}
                var v8 uint8 = (^(v4))
                _ = v7
                _ = v8
            } else {
                v9 := []int{-926, 553, 42}
                var v10 int8 = v1
                _ = v10
                _ = v9
            }
        }
    }
    for fi11 := 0; fi11 < 18; fi11++ {
        for fi12 := 0; fi12 < 9; fi12++ {
            v5 -= -114
            _ = fi12
        }
        var v13 int32 = 42
        v14 := []int{}
        _ = fi11
        _ = v13
        _ = v14
    }
    v15 := GenS1{F0: v4}
    v16 := GenS2{F0: v1, F1: 2}
    v17 := []int{128, -768, 720, 32767}
    var v18 int16 = -588
    v19 := GenS1{F0: 128}
    v20 := v2[995]
    v4 += (v4 - v4)
    v21 := GenS3{F0: v1, F1: 255}
    v4 -= v4
    var v22 int16 = v5
    v23 := GenS2{F0: v1, F1: v4}
    if ((v1 - v1) > (v1 ^ v1)) {
        switch v3 {
        case "a":
            v5 += v18
        case "":
            v5 = (v18 * 42)
        }
        if false {
            if len(v17) > 0 {
                v17[0] = -9223372036854775808
            }
            v24 := GenS1{F0: v4}
            v2[2] = "hello"
            _ = v24
        }
        switch 625 {
        case 32767:
            var v25 int = 127
            _ = v25
        case 359:
            fmt.Println(v3, v18, v22)
        default:
            switch (v3 + v20) {
            case "foo bar":
                var v26 int8 = v1
                _ = v26
            case "":
                delete(v2, 2)
            case "hello":
                v27 := GenS4{F0: ("foo bar" >= "hello")}
                _ = v27
            }
        }
    } else {
        for fi28 := 0; fi28 < 11; fi28++ {
            var v29 uint = 18446744073709551615
            switch fi28 {
            case -513:
                fmt.Println(v3)
            case -521:
                for fi30 := 0; fi30 < 19; fi30++ {
                    var v31 int8 = (-(v1))
                    v3 = ("foo bar" + v3)
                    switch fi30 {
                    case 4294967295:
                        v32 := GenS3{F0: (v1 >> ((0) & 63)), F1: v29}
                        if len(v17) >= 2 {
                            v17 = v17[1:len(v17)]
                        }
                        _ = v32
                    case 244:
                        v20 = "foo bar"
                    case 633:
                        var v33 uint32 = 2147483647
                        _ = v33
                    default:
                        switch (^(fi30)) {
                        case -1:
                            var v34 float32 = float32(1e10)
                            _ = v34
                        default:
                            switch v3 {
                            case "\n\t":
                                if true {
                                    if len(v17) >= 2 {
                                        v17 = v17[1:len(v17)]
                                    }
                                    v35 := GenS1{F0: (^(v4))}
                                    _ = v35
                                }
                            case "":
                                v18 -= v5
                            case "foo bar":
                                if true {
                                    if true {
                                        v1 -= v31
                                        var v36 uint32 = 42
                                        _ = v36
                                    }
                                    if (false && false) {
                                        var v37 uint8 = v4
                                        fmt.Println(v29)
                                        var v38 int8 = (v1 & 8)
                                        _ = v37
                                        _ = v38
                                    } else {
                                        v29 = 2
                                        var v39 uint32 = 2147483647
                                        _ = v39
                                    }
                                    var v40 uint = (^(v29))
                                    _ = v40
                                }
                            }
                        }
                    }
                    _ = fi30
                    _ = v31
                }
            case 127:
                for fi41 := 0; fi41 < 13; fi41++ {
                    var v42 float32 = (-2.25 + -2.25)
                    _ = len(v17)
                    _ = fi41
                    _ = v42
                }
            default:
                v18 -= v5
            }
            _ = fi28
            _ = v29
        }
    }
    v43 := GenS5{F0: 4294967295}
    var v44 uint = 255
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
}
