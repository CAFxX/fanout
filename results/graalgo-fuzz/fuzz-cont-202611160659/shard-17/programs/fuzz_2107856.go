package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 bool; F1 int8; F2 uint64 }
type GenS3 struct { F0 uint; F1 bool; F2 uint8 }
type GenS4 struct { F0 int; F1 string }
type GenS5 struct { F0 bool; F1 uint8; F2 uint16 }
type GenS6 struct { F0 uint64 }

func main() {
    var v1 uint32 = 920
    var v2 string = ("" + "\n\t")
    v3 := []int{128, 32767, 0}
    _ = len(v3)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v4 := GenS1{F0: -12, F1: (v2 > v2)}
    v5 := GenS2{F0: true, F1: -128, F2: 518}
    switch -9223372036854775808 {
    case 845:
        v6 := GenS1{F0: 924, F1: false}
        _ = v6
    case 1:
        fmt.Println(v1, v1)
    }
    var v7 float64 = ((0.1 + 3.14159) * 3.14159)
    var v8 int8 = 0
    v9 := GenS3{F0: 255, F1: false, F2: 128}
    v10 := GenS1{F0: 4294967295, F1: true}
    for fi11 := 0; fi11 < 9; fi11++ {
        if false {
            v2 = ("hello" + (v2 + v2))
            v7 += 1.5
        } else {
            if false {
                v12 := GenS4{F0: 32767, F1: "\n\t"}
                _ = v12
            } else {
                fmt.Println(fi11, fi11, v2)
                for fi13 := 0; fi13 < 11; fi13++ {
                    v14 := GenS2{F0: false, F1: v8, F2: 42}
                    _ = fi13
                    _ = v14
                }
                switch "" {
                case "hello":
                    if false {
                        var v15 int64 = 719
                        var v16 int8 = (^(v8))
                        var v17 float32 = 3.14159
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                case "\n\t":
                    var v18 uint16 = 708
                    _ = v18
                case "a":
                    if (v8 <= (-(v8))) {
                        var v19 uint16 = 0
                        for fi20 := 0; fi20 < 7; fi20++ {
                            fmt.Println(fi20, fi11, v2)
                            _ = fi20
                        }
                        if (true && false) {
                            var v21 int8 = -128
                            fmt.Println(v1, v2)
                            fmt.Println(v1)
                            _ = v21
                        } else {
                            switch int(v1) {
                            case -880:
                                if (!true) {
                                    if false {
                                        fmt.Println(v19, v19)
                                        var v22 bool = false
                                        var v23 int32 = 619
                                        _ = v22
                                        _ = v23
                                    } else {
                                        v8 -= v8
                                        v3 = append(v3, fi11)
                                    }
                                    if (true || (false || true)) {
                                        _ = len(v3)
                                    }
                                } else {
                                    v24 := GenS4{F0: fi11, F1: v2}
                                    var v25 uint64 = 375
                                    for fi26 := 0; fi26 < 7; fi26++ {
                                        var v27 int = fi26
                                        switch fi26 {
                                        case 4294967295:
                                            v28 := GenS5{F0: true, F1: 1, F2: 255}
                                            _ = v28
                                        default:
                                            var v29 float32 = float32((float64(float32(v7)) + (0.1 * v7)))
                                            _ = v29
                                        }
                                        var v30 uint = 1
                                        _ = fi26
                                        _ = v27
                                        _ = v30
                                    }
                                    _ = v24
                                    _ = v25
                                }
                            default:
                                for fi31 := 0; fi31 < 7; fi31++ {
                                    var v32 uint8 = 42
                                    fmt.Println(v1, v1, v8)
                                    var v33 uint = 255
                                    _ = fi31
                                    _ = v32
                                    _ = v33
                                }
                            }
                            v1 += v1
                        }
                        _ = v19
                    }
                default:
                    v34 := GenS1{F0: fi11, F1: true}
                    _ = v34
                }
            }
            if ((v2 + "a") > v2) {
                var v35 uint32 = 2147483647
                _ = v35
            } else {
                fmt.Println(v8, fi11)
            }
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v7 -= v7
        _ = fi11
    }
    var v36 int16 = 586
    v3 = append(v3, 127)
    v2 = "a"
    fmt.Println(v2, v1)
    v37 := GenS6{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
