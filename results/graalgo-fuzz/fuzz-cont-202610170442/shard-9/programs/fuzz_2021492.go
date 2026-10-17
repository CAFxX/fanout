package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool; F1 uint }
type GenS4 struct { F0 int16; F1 uint16 }
type GenS5 struct { F0 uint32; F1 uint8 }
func genFn0() uint64 {
    var v5 int8 = 1
    if false {
        fmt.Println(v5)
        for fi6 := 0; fi6 < 14; fi6++ {
            if (!false) {
                for fi7 := 0; fi7 < 12; fi7++ {
                    v5 += (-(v5))
                    switch "\n\t" {
                    case "a":
                        switch 9223372036854775807 {
                        case 32768:
                            v5 -= v5
                            v5 += (v5 | (v5 - (^(v5))))
                            switch (fi6 ^ fi7) {
                            case 127:
                                v8 := []int{(fi6 / (((^(fi7))) | 1)), fi6, (fi6 * fi7)}
                                _ = v8
                            case 923:
                                fmt.Println(fi6)
                            }
                            if false {
                                switch ("" + "foo bar") {
                                case "\n\t":
                                    v5 -= v5
                                case "hello":
                                    if (18446744073709551615 == 804) {
                                        var v9 int16 = 42
                                        v9 -= v9
                                        _ = v9
                                    } else {
                                        if true {
                                            var v10 string = "hello"
                                            _ = v10
                                        } else {
                                            v5 += 127
                                        }
                                        var v11 int32 = -2147483648
                                        v12 := []int{}
                                        v5 += v5
                                        _ = v11
                                        _ = v12
                                    }
                                    v5 -= -126
                                    v5 -= v5
                                }
                            }
                            v5 += v5
                            var v13 float32 = 3.14159
                            var v14 uint8 = 239
                            _ = v13
                            _ = v14
                        case 128:
                            var v15 int8 = 87
                            _ = v15
                        case -191:
                            if (1000 == 0) {
                                v5 = (-(v5))
                                var v16 int16 = 1000
                                _ = v16
                            }
                        }
                    case "\n\t":
                        var v17 uint64 = 1
                        _ = v17
                    case "foo bar":
                        fmt.Println(v5, fi7, fi7)
                    }
                    _ = fi7
                }
            }
            v5 = v5
            _ = fi6
        }
    }
    fmt.Println(v5)
    _ = v5
    return 4294967295
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]int)
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]string)
    v18 := GenS1{F0: -12, F1: 32767, F2: 42}
    var v19 uint64 = 495
    v20 := GenS2{F0: false}
    v19 = v19
    if true {
        var v21 int = 0
        for fi22 := 0; fi22 < 16; fi22++ {
            v23 := GenS2{F0: true}
            delete(v3, 0)
            fmt.Println(v21)
            _ = fi22
            _ = v23
        }
        _ = v21
    } else {
        if (32767 == 1) {
            v24 := GenS3{F0: true, F1: 415}
            for fi25 := 0; fi25 < 6; fi25++ {
                v19 = 913
                _ = fi25
            }
            delete(v4, 65535)
            _ = v24
        }
        var v26 uint32 = 1
        v27 := []int{32767, 9223372036854775807, 42}
        _ = v26
        _ = v27
    }
    switch -9223372036854775808 {
    case 42:
        v19 -= (v19 * v19)
    }
    v28 := GenS4{F0: -1, F1: 65535}
    if false {
        if (1.5 < 3.14159) {
            fmt.Println(v19)
            var v29 float32 = (-2.25 * 0.1)
            if false {
                v30 := []int{-9223372036854775808, -952, -487}
                for fi31 := 0; fi31 < 18; fi31++ {
                    var v32 uint8 = 127
                    fmt.Println(v32)
                    var v33 uint16 = 1
                    _ = fi31
                    _ = v32
                    _ = v33
                }
                for fi34 := 0; fi34 < 8; fi34++ {
                    var v35 float64 = ((3.14159 / 1.0) / 3.14159)
                    v36 := GenS4{F0: 42, F1: 685}
                    _ = fi34
                    _ = v35
                    _ = v36
                }
                _ = v30
            }
            _ = v29
        }
        if ((42 != -1) || (0 < 488)) {
            v37 := []int{4294967295, 127}
            v38 := GenS1{F0: 1, F1: 42, F2: 381}
            var v39 uint32 = 0
            _ = v37
            _ = v38
            _ = v39
        } else {
            var v40 float64 = 1.5
            _ = v40
        }
    }
    v41 := GenS5{F0: 4294967295, F1: 249}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v4
    _ = v41
}
