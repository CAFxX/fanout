package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint8; F1 uint64 }
type GenS7 struct { F0 int32; F1 int8 }
type GenS8 struct { F0 string; F1 int8; F2 uint8 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 uint; F1 int16; F2 int8 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]uint32)
    var v3 uint = 2
    if true {
        _ = len(v2)
        v4 := GenS1{F0: 2}
        v5 := GenS1{F0: 209}
        _ = v4
        _ = v5
    } else {
        for fi6 := 0; fi6 < 12; fi6++ {
            for fi7 := 0; fi7 < 4; fi7++ {
                if (true && true) {
                    v8 := GenS2{F0: ("" + ("" + (("" + ("a" + (("\n\t" + ("hello" + "\n\t")) + ("hello" + "a")))) + ("hello" + "a"))))}
                    switch fi7 {
                    case -458:
                        var v9 uint8 = 42
                        _ = v9
                    case 127:
                        v10 := GenS2{F0: "hello"}
                        _ = v10
                    case -450:
                        switch "\n\t" {
                        case "hello":
                            if true {
                                v3 += (v3 >> ((670) & 63))
                                v11 := []int{}
                                var v12 uint64 = 1
                                _ = v11
                                _ = v12
                            }
                        case "":
                            v13 := []int{}
                            _ = v13
                        default:
                            var v14 int64 = 612
                            _ = v14
                        }
                    }
                    var v15 int8 = -1
                    _ = v15
                    _ = v8
                } else {
                    if true {
                        v16 := []int{(fi7 >> ((4294967295) & 63))}
                        _ = len(v2)
                        v1[2] = true
                        _ = v16
                    }
                    var v17 uint8 = 0
                    v3 -= (v3 / ((v3) | 1))
                    _ = v17
                }
                if (-40 <= 100) {
                    var v18 int16 = 42
                    _ = v18
                }
                v19 := GenS2{F0: "\n\t"}
                var v20 int = 32767
                if (!true) {
                    switch "\n\t" {
                    case "":
                        var v21 int16 = -1
                        _ = v21
                    case "\n\t":
                        switch "hello" {
                        case "\n\t":
                            v20 = fi6
                            v22 := GenS3{F0: 1}
                            if (1 != 603) {
                                v23 := GenS4{F0: 0}
                                _ = v23
                            } else {
                                var v24 uint64 = 2
                                _ = v24
                            }
                            switch "foo bar" {
                            case "":
                                v25 := GenS4{F0: 18446744073709551615}
                                _ = v25
                            case "a":
                                if false {
                                    if false {
                                        var v26 int32 = 393
                                        var v27 float32 = ((1e10 / float32(0.1)) - 3.14159)
                                        v28 := GenS5{F0: -9223372036854775808}
                                        _ = v26
                                        _ = v27
                                        _ = v28
                                    } else {
                                        switch ("\n\t" + (("" + "foo bar") + (("hello" + "foo bar") + ""))) {
                                        case "hello":
                                            var v29 float32 = -2.25
                                            _ = v29
                                        case "\n\t":
                                            var v30 bool = (1 <= 782)
                                            var v31 uint16 = 742
                                            _ = v30
                                            _ = v31
                                        case "a":
                                            var v32 uint8 = 2
                                            _ = v32
                                        default:
                                            v3 -= 183
                                        }
                                        v20 -= fi6
                                        var v33 int64 = 43
                                        _ = v33
                                    }
                                    v34 := []int{}
                                    _ = v34
                                }
                            case "\n\t":
                                fmt.Println(v20)
                                v35 := []int{-1, fi7, fi6, fi7}
                                _ = v35
                            }
                            _ = v22
                        case "foo bar":
                            fmt.Println(fi6)
                        }
                    default:
                        if (2 < 0) {
                            v3 = (^(v3))
                            var v36 uint = v3
                            switch "" {
                            case "hello":
                                _ = len(v2)
                            default:
                                fmt.Println(v20, v36, v20)
                            }
                            _ = v36
                        } else {
                            v37 := []int{}
                            v38 := v1[454]
                            _ = v37
                            _ = v38
                        }
                    }
                    v20 += 32767
                    if (("a" <= "a") || false) {
                        var v39 bool = false
                        v40 := []int{(v20 ^ fi6)}
                        _ = v39
                        _ = v40
                    } else {
                        var v41 uint64 = 4294967295
                        v20 = (^(fi6))
                        v42 := GenS4{F0: v41}
                        _ = v41
                        _ = v42
                    }
                }
                var v43 string = ("" + "hello")
                v44 := GenS6{F0: 191, F1: 42}
                var v45 float64 = ((-2.25 * float64(0.0)) + -2.25)
                _ = fi7
                _ = v19
                _ = v20
                _ = v43
                _ = v44
                _ = v45
            }
            switch -912 {
            case 9223372036854775807:
                var v46 uint32 = 811
                _ = v46
            default:
                v3 = 18446744073709551615
            }
            delete(v2, (-(fi6)))
            _ = fi6
        }
    }
    if ((0.0 - float64(-2.25)) >= -2.25) {
        v47 := GenS7{F0: 42, F1: -31}
        v3 = v3
        _ = v47
    }
    switch 0 {
    case 4294967295:
        var v48 int16 = 794
        _ = v48
    default:
        if (47 != 32) {
            v3 += v3
        } else {
            for fi49 := 0; fi49 < 9; fi49++ {
                v3 += (v3 * v3)
                _ = fi49
            }
            switch "" {
            case "":
                v3 = (v3 / (((v3 + (v3 >> ((695) & 63)))) | 1))
            case "foo bar":
                v50 := GenS8{F0: ("hello" + ""), F1: 1, F2: 127}
                _ = v50
            }
        }
    }
    fmt.Println(v3)
    for fi51 := 0; fi51 < 8; fi51++ {
        v52 := GenS9{F0: 107}
        _ = fi51
        _ = v52
    }
    v3 += (v3 + v3)
    var v53 bool = ("hello" > "")
    v3 -= (v3 % (((^(v3))) | 1))
    var v54 uint64 = 322
    v55 := GenS10{F0: v3, F1: 42, F2: 127}
    var v56 int8 = -1
    var v57 int64 = 1
    var v58 bool = (0.0 != -2.25)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v53
    _ = v54
    _ = v55
    _ = v56
    _ = v57
    _ = v58
}
