package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 int8 }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint; F1 uint8; F2 int64 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 string }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]bool)
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]bool)
    var v5 uint8 = 127
    delete(v4, 0)
    var v6 int64 = 0
    v7 := GenS1{F0: ((true && false) || (0 <= -741)), F1: 2147483647, F2: 127}
    v8 := []int{32767, 598, 2147483647}
    v9 := v3[1]
    for fi10 := 0; fi10 < 13; fi10++ {
        if (true && (670 > 18446744073709551615)) {
            for fi11 := 0; fi11 < 15; fi11++ {
                _ = len(v8)
                switch fi11 {
                case -2147483648:
                    v12 := GenS2{F0: 42, F1: 42}
                    _ = v12
                case 9223372036854775807:
                    _ = len(v8)
                    _ = len(v8)
                case 4294967295:
                    if true {
                        v13 := GenS3{F0: 32767, F1: 2147483648}
                        _ = v13
                    }
                default:
                    v6 = (v6 ^ 42)
                }
                if true {
                    v14 := GenS1{F0: ((0 != 1) || (!(-1 <= fi11))), F1: -5, F2: 0}
                    _ = v14
                } else {
                    if (1 >= 0) {
                        var v15 int = -1
                        v6 -= 543
                        var v16 bool = false
                        _ = v15
                        _ = v16
                    } else {
                        if (32767 > 255) {
                            if false {
                                if (32767 <= 2) {
                                    fmt.Println(fi11)
                                    v17 := GenS3{F0: 255, F1: (^(v9))}
                                    _ = v17
                                }
                                fmt.Println(v5, v9, v9)
                                if false {
                                    var v18 int32 = 1
                                    v18 = v18
                                    switch 0 {
                                    case 0:
                                        v19 := GenS3{F0: 2, F1: v9}
                                        _ = v19
                                    case -2147483648:
                                        v20 := GenS2{F0: 1, F1: 18446744073709551615}
                                        _ = v20
                                    case 210:
                                        v6 = int64(v9)
                                    }
                                    _ = v18
                                }
                            } else {
                                if (1 != 127) {
                                    if len(v8) >= 2 {
                                        v8 = v8[1:len(v8)]
                                    }
                                    var v21 uint16 = 622
                                    _ = v21
                                } else {
                                    switch "foo bar" {
                                    case "a":
                                        v22 := GenS2{F0: -1, F1: 42}
                                        _ = v22
                                    case "hello":
                                        v2["hello"] = (!false)
                                    case "foo bar":
                                        v5 = v5
                                    }
                                }
                                if false {
                                    v23 := GenS2{F0: 1, F1: 1}
                                    delete(v1, "foo bar")
                                    var v24 int = (fi11 | (fi11 ^ fi11))
                                    _ = v23
                                    _ = v24
                                }
                                var v25 uint8 = 127
                                _ = v25
                            }
                            v5 -= v5
                        } else {
                            if (-2.25 == ((float32(1.5) + ((float32(-2.25) + float32(((1.5 / (0.1 + (-2.25 / -2.25))) + 1.5))) / 0.1)) / 0.1)) {
                                var v26 uint16 = 32767
                                _ = v26
                            } else {
                                if false {
                                    var v27 uint32 = (v9 % ((v9) | 1))
                                    if (-29 == -128) {
                                        v27 -= (v27 + v27)
                                        v27 = v9
                                        if true {
                                            _ = len(v8)
                                        } else {
                                            _ = len(v8)
                                            if len(v8) > 0 {
                                                v8[0] = fi10
                                            }
                                            var v28 uint = 574
                                            var v29 float64 = (-2.25 * -2.25)
                                            _ = v28
                                            _ = v29
                                        }
                                    } else {
                                        var v30 uint32 = v27
                                        v31 := GenS4{F0: 1}
                                        fmt.Println(fi10)
                                        var v32 uint8 = v5
                                        var v33 uint8 = v32
                                        _ = v30
                                        _ = v31
                                        _ = v32
                                        _ = v33
                                    }
                                    _ = v27
                                }
                            }
                            fmt.Println(fi10, v5, fi10)
                            v9 = v9
                        }
                    }
                    var v34 int64 = v6
                    var v35 string = "a"
                    _ = v34
                    _ = v35
                }
                _ = fi11
            }
        } else {
            v36 := GenS5{F0: 255, F1: 168, F2: v6}
            v37 := GenS6{F0: (("" + ((("foo bar" + "") + "hello") + "")) + ("foo bar" + ("hello" + "")))}
            _ = v36
            _ = v37
        }
        _ = fi10
    }
    for fi38 := 0; fi38 < 16; fi38++ {
        switch (fi38 - (fi38 + (-(fi38)))) {
        case -9223372036854775808:
            var v39 int = 9223372036854775807
            _ = v39
        default:
            v9 += v9
        }
        v8 = append(v8, fi38)
        for fi40 := 0; fi40 < 18; fi40++ {
            if (((false && true) || false) && ((!((-(v6)) != 362)) || false)) {
                switch "foo bar" {
                case "\n\t":
                    if (!true) {
                        var v41 int8 = -44
                        fmt.Println(fi38, v41, v6)
                        _ = v41
                    } else {
                        v5 = v5
                        v42 := GenS7{F0: "hello"}
                        _ = v42
                    }
                default:
                    v5 = (v5 << ((2) & 63))
                }
                v9 = 947
            }
            v9 = (v9 & v9)
            _ = fi40
        }
        _ = fi38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
