package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool; F1 int32; F2 int8 }

func main() {
    var v1 float32 = ((1e10 - 1e10) - 0.0)
    var v2 string = "\n\t"
    var v3 uint64 = 2
    var v4 int32 = -1
    var v5 int32 = v4
    v6 := []int{-622, int(v4), -9223372036854775808, -9223372036854775808}
    var v7 bool = false
    if len(v6) > 0 {
        v6[0] = 128
    }
    v1 = v1
    fmt.Println(v5, v5)
    v8 := GenS1{F0: -554, F1: 2147483647, F2: v4}
    if (0 == -9223372036854775808) {
        if v7 {
            var v9 uint = 0
            if v7 {
                fmt.Println(v3, v4, v1)
                _ = len(v6)
                var v10 uint32 = 1
                _ = v10
            } else {
                var v11 uint8 = 115
                _ = v11
            }
            var v12 string = v2
            _ = v12
            _ = v9
        }
        v7 = false
    } else {
        for fi13 := 0; fi13 < 8; fi13++ {
            for fi14 := 0; fi14 < 19; fi14++ {
                var v15 uint16 = 42
                v3 = (v3 >> ((65535) & 63))
                _ = fi14
                _ = v15
            }
            switch "hello" {
            case "\n\t":
                switch "" {
                case "":
                    switch fi13 {
                    case 450:
                        if v7 {
                            var v16 uint8 = 2
                            _ = v16
                        } else {
                            v17 := GenS1{F0: (fi13 + fi13), F1: fi13, F2: 42}
                            _ = v17
                        }
                    case 32768:
                        switch (("a" + v2) + "a") {
                        case "hello":
                            var v18 int16 = 42
                            _ = v18
                        case "":
                            _ = len(v6)
                        default:
                            v19 := GenS1{F0: fi13, F1: fi13, F2: -1}
                            _ = v19
                        }
                    default:
                        _ = len(v6)
                    }
                case "\n\t":
                    for fi20 := 0; fi20 < 12; fi20++ {
                        var v21 uint64 = (v3 & 0)
                        v2 = v2
                        _ = fi20
                        _ = v21
                    }
                case "a":
                    if v7 {
                        if true {
                            for fi22 := 0; fi22 < 5; fi22++ {
                                v23 := GenS2{F0: 2147483647}
                                _ = fi22
                                _ = v23
                            }
                            fmt.Println(v4, v7)
                        } else {
                            switch (fi13 % ((-442) | 1)) {
                            case -1:
                                var v24 uint64 = (v3 / ((v3) | 1))
                                _ = v24
                            case -271:
                                if v7 {
                                    fmt.Println(fi13, v2, v7)
                                    for fi25 := 0; fi25 < 14; fi25++ {
                                        v26 := GenS2{F0: 4294967295}
                                        var v27 uint = 255
                                        var v28 int = (fi25 >> (((v27 >> ((v27) & 63))) & 63))
                                        _ = fi25
                                        _ = v26
                                        _ = v27
                                        _ = v28
                                    }
                                    var v29 int8 = -128
                                    _ = v29
                                } else {
                                    for fi30 := 0; fi30 < 6; fi30++ {
                                        v5 = (v5 >> ((254) & 63))
                                        v31 := GenS2{F0: 65535}
                                        _ = fi30
                                        _ = v31
                                    }
                                }
                            }
                            if false {
                                var v32 float32 = ((((((((float32((float64((v1 + float32((1.5 * 0.0)))) / (-2.25 + (((0.1 * 0.1) - 1.5) - 0.1)))) / float32((-2.25 * float64(v1)))) + (v1 / (((3.14159 * (1.5 - (1.5 / 1.0))) + v1) / 1.0))) - (v1 - v1)) * v1) * v1) / 0.1) / (float32((1.5 * ((-2.25 / 1.5) / 1e10))) - float32((0.0 + -2.25)))) + v1)
                                _ = v32
                            } else {
                                fmt.Println(v3, v3, fi13)
                            }
                            var v33 uint8 = 66
                            _ = v33
                        }
                        v34 := GenS3{F0: -86}
                        v3 = v3
                        _ = v34
                    } else {
                        _ = len(v6)
                    }
                }
            default:
                if v7 {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    for fi35 := 0; fi35 < 5; fi35++ {
                        if true {
                            var v36 uint32 = 1
                            var v37 uint16 = 993
                            switch "foo bar" {
                            case "a":
                                var v38 uint = 18446744073709551615
                                if (v7 && false) {
                                    fmt.Println(fi35)
                                }
                                _ = v38
                            }
                            _ = v36
                            _ = v37
                        } else {
                            var v39 uint64 = v3
                            _ = v39
                        }
                        v40 := GenS2{F0: 2}
                        _ = fi35
                        _ = v40
                    }
                    switch (fi13 | fi13) {
                    case -58:
                        var v41 int8 = 42
                        _ = v41
                    }
                } else {
                    _ = len(v6)
                }
            }
            v42 := GenS2{F0: 2}
            _ = fi13
            _ = v42
        }
        v2 = ""
    }
    v4 = v4
    v43 := GenS4{F0: v7, F1: v5, F2: -98}
    _ = len(v6)
    if v7 {
        var v44 int = -1
        _ = v44
    }
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
