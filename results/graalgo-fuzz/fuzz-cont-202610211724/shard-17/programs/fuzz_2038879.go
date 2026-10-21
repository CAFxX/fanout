package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 bool; F1 string; F2 int16 }
type GenS6 struct { F0 int8; F1 uint32; F2 uint }

func main() {
    var v1 float32 = 1.5
    var v2 int64 = 42
    var v3 int64 = 0
    var v4 int8 = 100
    v5 := make(map[uint32]string)
    v6 := GenS1{F0: 0, F1: 701, F2: 4294967295}
    v7 := v5[1]
    delete(v5, 2)
    var v8 int16 = 42
    for fi9 := 0; fi9 < 16; fi9++ {
        if false {
            fmt.Println(v4)
            if true {
                for fi10 := 0; fi10 < 9; fi10++ {
                    fmt.Println(v2, v1)
                    var v11 uint8 = 147
                    _ = fi10
                    _ = v11
                }
                v5[1] = ""
            } else {
                var v12 int32 = -1
                _ = v12
            }
            for fi13 := 0; fi13 < 11; fi13++ {
                if true {
                    v8 = v8
                    v7 = v7
                } else {
                    if false {
                        v3 = v3
                        v7 = v7
                        v14 := GenS1{F0: 2147483648, F1: 0, F2: 860}
                        _ = v14
                    }
                    var v15 float64 = 0.1
                    switch fi9 {
                    case -9223372036854775808:
                        v16 := GenS2{F0: 42}
                        _ = v16
                    case -693:
                        if false {
                            var v17 uint8 = 1
                            _ = v17
                        } else {
                            v18 := []int{}
                            v15 -= (v15 * ((v15 / -2.25) - ((1.5 / v15) - v15)))
                            fmt.Println(v15, v4, fi9)
                            _ = v18
                        }
                    default:
                        switch v7 {
                        case "foo bar":
                            if (((((4294967295 < 1) || false) || false) || false) || true) {
                                var v19 uint16 = 0
                                v20 := GenS2{F0: 273}
                                if true {
                                    if false {
                                        var v21 float64 = v15
                                        switch v7 {
                                        case "\n\t":
                                            var v22 uint8 = 42
                                            _ = v22
                                        default:
                                            var v23 float32 = v1
                                            _ = v23
                                        }
                                        var v24 int64 = 598
                                        _ = v21
                                        _ = v24
                                    }
                                    v25 := GenS3{F0: (fi9 >> ((18446744073709551615) & 63))}
                                    _ = v25
                                } else {
                                    var v26 string = (("a" + (v7 + ((((v7 + v7) + v7) + (("hello" + v7) + (v7 + v7))) + v7))) + "")
                                    _ = v26
                                }
                                _ = v19
                                _ = v20
                            }
                        case "\n\t":
                            v4 -= 100
                        case "hello":
                            switch 695 {
                            case 480:
                                fmt.Println(fi13, v4)
                            case -160:
                                var v27 string = v7
                                _ = v27
                            default:
                                v2 += v3
                            }
                        }
                    }
                    _ = v15
                }
                v8 = v8
                switch (v7 + v7) {
                case "hello":
                    switch ("" + "") {
                    case "\n\t":
                        if false {
                            v28 := v5[42]
                            _ = v28
                        }
                    case "":
                        if true {
                            v29 := GenS4{F0: 0}
                            if (!(!true)) {
                                v30 := GenS4{F0: 1}
                                _ = v30
                            } else {
                                fmt.Println(fi13, v4, v2)
                            }
                            v31 := GenS3{F0: (fi13 ^ fi13)}
                            _ = v29
                            _ = v31
                        }
                    case "hello":
                        var v32 int = (-(fi9))
                        _ = v32
                    }
                case "foo bar":
                    switch v7 {
                    case "foo bar":
                        v2 -= (-(v2))
                    case "hello":
                        var v33 int32 = -1
                        _ = v33
                    default:
                        if false {
                            var v34 bool = (!true)
                            _ = v34
                        }
                        v35 := []int{(fi9 / (((fi9 % ((fi9) | 1))) | 1)), (fi9 / ((fi13) | 1))}
                        v36 := GenS5{F0: true, F1: "a", F2: 32767}
                        _ = v35
                        _ = v36
                    }
                    v3 = v2
                case "a":
                    v7 = "foo bar"
                }
                fmt.Println(v4, v3, v1)
                v3 += 638
                v37 := []int{(fi13 / ((fi13) | 1)), 1, int(v4)}
                v7 = (v7 + v7)
                _ = fi13
                _ = v37
            }
        } else {
            fmt.Println(v1, v8)
        }
        _ = fi9
    }
    v38 := []int{-9223372036854775808, -454, 0, 128}
    v3 = v2
    v39 := GenS5{F0: (true && true), F1: "", F2: 0}
    v40 := GenS6{F0: (v4 << ((490) & 63)), F1: 2147483648, F2: 18446744073709551615}
    v41 := GenS5{F0: (v8 == 1000), F1: v7, F2: v8}
    v42 := v5[42]
    v43 := GenS6{F0: v4, F1: 0, F2: 255}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
