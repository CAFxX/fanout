package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 uint32; F2 int64 }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 int8; F1 uint8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int16 }
func genFn0() bool {
    switch 32768 {
    case -542:
        v6 := GenS1{F0: 757}
        _ = v6
    }
    if (3.14159 < (1.5 / 3.14159)) {
        if (("" + "\n\t") == "\n\t") {
            switch "" {
            case "foo bar":
                var v7 int8 = 1
                _ = v7
            case "hello":
                var v8 float32 = 3.14159
                _ = v8
            case "a":
                if true {
                    var v9 int64 = 1
                    if true {
                        v10 := GenS2{F0: 23, F1: 2147483648, F2: v9}
                        var v11 bool = (!(65535 > 2))
                        _ = v10
                        _ = v11
                    }
                    if true {
                        if false {
                            for fi12 := 0; fi12 < 13; fi12++ {
                                var v13 uint = 42
                                _ = fi12
                                _ = v13
                            }
                        }
                        fmt.Println(v9)
                        var v14 float32 = 0.1
                        _ = v14
                    }
                    _ = v9
                }
            }
            switch "a" {
            case "hello":
                switch 4294967295 {
                case 32767:
                    var v15 int8 = -56
                    _ = v15
                case -836:
                    switch -58 {
                    case -275:
                        if false {
                            switch "foo bar" {
                            case "\n\t":
                                for fi16 := 0; fi16 < 7; fi16++ {
                                    var v17 int32 = 42
                                    var v18 uint16 = 2
                                    _ = fi16
                                    _ = v17
                                    _ = v18
                                }
                            case "hello":
                                var v19 float32 = 0.1
                                _ = v19
                            default:
                                if (128 <= 2) {
                                    var v20 int = 2147483647
                                    _ = v20
                                }
                            }
                        } else {
                            var v21 bool = (((false && false) || true) && (!true))
                            v22 := GenS3{F0: ((!(29 >= 128)) && true), F1: 1}
                            _ = v21
                            _ = v22
                        }
                    case 42:
                        var v23 string = ("foo bar" + "a")
                        _ = v23
                    case 4294967295:
                        if (4294967295 == -9223372036854775808) {
                            v24 := GenS4{F0: 42, F1: 135}
                            switch (("hello" + "foo bar") + "") {
                            case "":
                                switch 32767 {
                                case 128:
                                    var v25 int = 973
                                    _ = v25
                                case 2147483647:
                                    var v26 uint64 = 4294967295
                                    _ = v26
                                case 0:
                                    var v27 bool = (true || true)
                                    _ = v27
                                }
                            default:
                                var v28 int32 = -933
                                _ = v28
                            }
                            _ = v24
                        } else {
                            var v29 string = "\n\t"
                            _ = v29
                        }
                    default:
                        v30 := []int{-9223372036854775808, 1}
                        _ = v30
                    }
                case -475:
                    var v31 string = "hello"
                    _ = v31
                default:
                    var v32 float64 = float64(0.0)
                    _ = v32
                }
            case "\n\t":
                if true {
                    switch "" {
                    case "\n\t":
                        v33 := []int{512, -905}
                        _ = v33
                    case "hello":
                        v34 := GenS5{F0: 1}
                        _ = v34
                    case "a":
                        var v35 int32 = 1
                        _ = v35
                    default:
                        var v36 uint = 464
                        _ = v36
                    }
                }
            default:
                v37 := []int{-1, 430}
                _ = v37
            }
        } else {
            if (!(2 >= 42)) {
                for fi38 := 0; fi38 < 16; fi38++ {
                    var v39 float32 = (-2.25 / ((1.5 * 1e10) / -2.25))
                    _ = fi38
                    _ = v39
                }
                var v40 float64 = (((-2.25 - -2.25) * float64(0.1)) / (0.0 - ((0.1 + (3.14159 * (-2.25 * 1e10))) / float64(float32(3.14159)))))
                v41 := []int{-1, 930}
                _ = v40
                _ = v41
            } else {
                if false {
                    var v42 int64 = -1
                    v43 := []int{2147483647}
                    _ = v42
                    _ = v43
                }
                var v44 int = 9223372036854775807
                _ = v44
            }
        }
    }
    return true
}

func main() {
    var v1 int32 = -793
    v2 := make(map[uint32]bool)
    var v3 bool = false
    var v4 int64 = -729
    var v5 int16 = 250
    var v45 float32 = float32((-2.25 + -2.25))
    for fi46 := 0; fi46 < 13; fi46++ {
        v47 := []int{}
        fmt.Println(v4, v5)
        _ = fi46
        _ = v47
    }
    fmt.Println(v4, v1)
    var v48 string = ("hello" + (("a" + "") + "foo bar"))
    if (499 == 255) {
        var v49 int32 = (v1 & (v1 >> ((18446744073709551615) & 63)))
        var v50 int16 = int16(v4)
        v5 = v50
        _ = v49
        _ = v50
    } else {
        fmt.Println(v3, v5)
        fmt.Println(v1, v3)
    }
    v5 += 652
    if v3 {
        for fi51 := 0; fi51 < 1; fi51++ {
            fmt.Println(v48, v48)
            _ = fi51
        }
        for fi52 := 0; fi52 < 1; fi52++ {
            if v3 {
                v53 := GenS6{F0: (v5 | -32768)}
                _ = v53
            } else {
                for fi54 := 0; fi54 < 13; fi54++ {
                    var v55 uint16 = 42
                    fmt.Println(fi52)
                    _ = fi54
                    _ = v55
                }
                v56 := []int{}
                _ = v56
            }
            _ = fi52
        }
    }
    v57 := GenS1{F0: 255}
    v58 := []int{9223372036854775807, 32768}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v58:", len(v58), v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v48
    _ = v5
    _ = v57
    _ = v58
}
