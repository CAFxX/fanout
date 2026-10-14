package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32; F1 int; F2 int }
type GenS5 struct { F0 int; F1 int64; F2 int8 }
type GenS6 struct { F0 bool; F1 uint8 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint8; F1 int32; F2 int8 }
func genFn0(p0 string, p1 uint32, p2 uint64) int64 {
    fmt.Println(p2)
    if true {
        var v7 uint32 = p1
        _ = v7
    } else {
        p2 += (p2 | 18446744073709551615)
        v8 := []int{9223372036854775807, -9223372036854775808, 99}
        fmt.Println(p1, p1, p2)
        _ = v8
    }
    v9 := []int{487, 9223372036854775807}
    v10 := GenS1{F0: 2147483647, F1: ("hello" + (("foo bar" + p0) + p0))}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return -9223372036854775808
}

func main() {
    var v1 uint8 = 29
    var v2 string = "foo bar"
    var v3 uint16 = 185
    v4 := make(map[string]bool)
    v5 := make(map[uint32]int)
    v6 := make(map[uint32]string)
    if false {
        v2 = ""
        v1 = v1
    }
    var v11 int = -27
    v6[2147483647] = "a"
    if (((!true) && false) && true) {
        if true {
            if true {
                v1 += (v1 ^ uint8(v3))
            } else {
                for fi12 := 0; fi12 < 2; fi12++ {
                    for fi13 := 0; fi13 < 10; fi13++ {
                        var v14 int = 0
                        switch (-(fi13)) {
                        case 42:
                            v14 += v11
                        }
                        _ = fi13
                        _ = v14
                    }
                    _ = fi12
                }
                fmt.Println(v11, v11, v2)
            }
        } else {
            switch v11 {
            case 9223372036854775807:
                v15 := []int{v11, 127}
                _ = v15
            case 0:
                for fi16 := 0; fi16 < 14; fi16++ {
                    if (!((true && false) && true)) {
                        for fi17 := 0; fi17 < 10; fi17++ {
                            v18 := GenS2{F0: true}
                            switch 128 {
                            case 1:
                                v19 := GenS2{F0: (236 >= 851)}
                                _ = v19
                            case 2147483647:
                                v3 = v3
                            }
                            var v20 int32 = 815
                            _ = fi17
                            _ = v18
                            _ = v20
                        }
                        v2 = v2
                        var v21 uint8 = v1
                        _ = v21
                    } else {
                        if false {
                            v22 := []int{v11}
                            _ = v22
                        }
                        fmt.Println(v3)
                    }
                    v23 := GenS3{F0: "hello"}
                    _ = fi16
                    _ = v23
                }
            case 504:
                var v24 float32 = (0.1 * -2.25)
                _ = v24
            default:
                v11 = v11
            }
        }
        v25 := v5[765]
        switch v25 {
        case 249:
            v25 -= (v11 << ((1) & 63))
        case 32767:
            if (genFn0((((v2 + v2) + v2) + ("foo bar" + "hello")), 2, 312) == 495) {
                var v26 int16 = -137
                v27 := GenS4{F0: 0, F1: v25, F2: v11}
                v26 += 332
                _ = v26
                _ = v27
            } else {
                v28 := GenS5{F0: (v25 % (((v11 ^ v25)) | 1)), F1: 0, F2: 57}
                if ((-(v11)) != v11) {
                    var v29 uint = 42
                    v30 := GenS6{F0: true, F1: v1}
                    if false {
                        v31 := GenS7{F0: uint64(v29)}
                        v32 := []int{}
                        _ = v31
                        _ = v32
                    } else {
                        var v33 int = (^(v11))
                        for fi34 := 0; fi34 < 8; fi34++ {
                            var v35 int32 = 0
                            if false {
                                var v36 int = (v11 * v11)
                                _ = v36
                            } else {
                                v37 := []int{}
                                v38 := GenS2{F0: (-2.25 > 1.5)}
                                _ = v37
                                _ = v38
                            }
                            _ = fi34
                            _ = v35
                        }
                        if (0 != -9223372036854775808) {
                            switch v2 {
                            case "foo bar":
                                delete(v5, 0)
                            case "\n\t":
                                v39 := GenS1{F0: -2147483648, F1: v2}
                                _ = v39
                            default:
                                switch v2 {
                                case "hello":
                                    v40 := GenS8{F0: (v1 >> ((v29) & 63)), F1: -1, F2: -1}
                                    _ = v40
                                case "":
                                    var v41 bool = false
                                    _ = v41
                                case "\n\t":
                                    if false {
                                        v42 := v4[((v2 + (v2 + v2)) + "\n\t")]
                                        _ = v42
                                    } else {
                                        if ((v1 + (v1 ^ v1)) > (v1 >> ((v29) & 63))) {
                                            v43 := []int{}
                                            _ = v43
                                        }
                                        var v44 uint = 1
                                        _ = v44
                                    }
                                default:
                                    fmt.Println(v2, v3)
                                }
                            }
                            fmt.Println(v11, v25)
                        }
                        _ = v33
                    }
                    _ = v29
                    _ = v30
                }
                _ = v28
            }
        }
        _ = v25
    }
    var v45 uint8 = v1
    var v46 int = v11
    var v47 uint = 0
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
