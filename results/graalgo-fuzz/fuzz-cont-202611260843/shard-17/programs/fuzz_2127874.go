package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 uint32; F2 string }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 int16; F1 uint8; F2 uint64 }
type GenS5 struct { F0 int16; F1 uint8 }
type GenS6 struct { F0 int; F1 int; F2 uint8 }
type GenS7 struct { F0 int; F1 int16 }
func genFn0(p0 int32) uint {
    for fi4 := 0; fi4 < 11; fi4++ {
        v5 := []int{42, (fi4 ^ (fi4 - fi4))}
        if ((float32(1e10) - 0.1) >= 3.14159) {
            for fi6 := 0; fi6 < 9; fi6++ {
                var v7 uint16 = 32767
                v8 := GenS1{F0: 610}
                var v9 bool = true
                switch fi4 {
                case 147:
                    var v10 bool = (v9 || v9)
                    _ = v10
                case 128:
                    var v11 string = "a"
                    v7 += v7
                    if len(v5) > 0 {
                        v5[0] = fi6
                    }
                    v7 = (v7 << ((1) & 63))
                    if (-2.25 < (((0.1 / 1.0) / 1.0) / ((1.5 + -2.25) + 0.0))) {
                        if v9 {
                            fmt.Println(fi6)
                        } else {
                            p0 += p0
                        }
                        v11 = v11
                    }
                    _ = v11
                case 0:
                    v9 = v9
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi6
                _ = v7
                _ = v8
                _ = v9
            }
            p0 = p0
            for fi12 := 0; fi12 < 1; fi12++ {
                p0 += p0
                p0 += p0
                _ = fi12
            }
        }
        _ = fi4
        _ = v5
    }
    for fi13 := 0; fi13 < 5; fi13++ {
        v14 := GenS1{F0: 2}
        _ = fi13
        _ = v14
    }
    v15 := GenS2{F0: 42, F1: 2, F2: ("" + "")}
    var v16 uint16 = 0
    v16 -= 32767
    _ = p0
    _ = v15
    _ = v16
    return 18446744073709551615
}

func main() {
    var v1 uint16 = 2
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]bool)
    v1 = (v1 >> ((474) & 63))
    if (false || ((("foo bar" + "a") <= "foo bar") && (!true))) {
        v1 = 32767
        for fi17 := 0; fi17 < 8; fi17++ {
            if true {
                fmt.Println(fi17, v1)
                if (((431 == 1) && false) && (false && (1 < 0))) {
                    fmt.Println(v1, fi17)
                } else {
                    v18 := []int{127, fi17, fi17, 9223372036854775807}
                    _ = len(v3)
                    v1 = v1
                    _ = v18
                }
            } else {
                if (!false) {
                    var v19 float64 = 1.5
                    _ = v19
                } else {
                    var v20 int8 = -1
                    _ = v20
                }
            }
            switch ("foo bar" + ("\n\t" + "a")) {
            case "a":
                for fi21 := 0; fi21 < 12; fi21++ {
                    v22 := GenS3{F0: 1, F1: 42}
                    if true {
                        var v23 int64 = -9223372036854775808
                        v24 := GenS2{F0: -1, F1: 2, F2: "\n\t"}
                        v1 = 421
                        _ = v23
                        _ = v24
                    }
                    var v25 uint = 65535
                    _ = fi21
                    _ = v22
                    _ = v25
                }
            case "foo bar":
                v1 = 359
            case "hello":
                v26 := []int{32768, 9223372036854775807, (fi17 - fi17), fi17}
                _ = v26
            default:
                for fi27 := 0; fi27 < 19; fi27++ {
                    v28 := GenS2{F0: -1, F1: 117, F2: "\n\t"}
                    _ = fi27
                    _ = v28
                }
            }
            _ = fi17
        }
    } else {
        for fi29 := 0; fi29 < 2; fi29++ {
            v30 := GenS4{F0: 851, F1: 255, F2: 339}
            _ = fi29
            _ = v30
        }
        var v31 uint16 = (^(v1))
        for fi32 := 0; fi32 < 3; fi32++ {
            v31 -= (v31 | (v31 & 42))
            v33 := GenS5{F0: 32767, F1: 1}
            var v34 float64 = (float64(0.1) * 3.14159)
            _ = fi32
            _ = v33
            _ = v34
        }
        _ = v31
    }
    v35 := GenS6{F0: 0, F1: 42, F2: 128}
    v36 := GenS2{F0: 32767, F1: 88, F2: ""}
    v37 := []int{-1, 128, 9223372036854775807, 334}
    var v38 int = 0
    var v39 int8 = 100
    v1 -= 32767
    v38 = 985
    var v40 uint = 255
    switch "a" {
    case "hello":
        v39 += v39
    case "a":
        var v41 bool = (v38 >= (^(v38)))
        _ = v41
    case "foo bar":
        switch -9223372036854775808 {
        case 32768:
            fmt.Println(v1)
        case 128:
            v37 = append(v37, v38)
        case -1:
            fmt.Println(v40, v40, v39)
        default:
            v42 := GenS7{F0: 0, F1: -263}
            _ = v42
        }
    }
    v43 := GenS4{F0: -358, F1: 255, F2: 453}
    v38 += v38
    if true {
        v37 = append(v37, v38)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v40
    _ = v43
}
