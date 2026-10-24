package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 uint16; F2 int32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 bool; F1 uint16 }
type GenS7 struct { F0 bool; F1 uint32; F2 int64 }
func genFn0(p0 int64) uint {
    var v6 float32 = -2.25
    for fi7 := 0; fi7 < 3; fi7++ {
        switch fi7 {
        case -99:
            var v8 int = fi7
            _ = v8
        case 42:
            p0 -= 1
        default:
            v6 = -2.25
        }
        _ = fi7
    }
    switch 127 {
    case 127:
        var v9 string = ("a" + "a")
        _ = v9
    case -502:
        if ((32768 >= 456) || (243 <= 42)) {
            var v10 uint32 = 2147483647
            switch "a" {
            case "hello":
                v11 := []int{-822, 9223372036854775807, -9223372036854775808}
                _ = v11
            case "":
                if false {
                    v6 += (-2.25 * -2.25)
                }
            default:
                for fi12 := 0; fi12 < 7; fi12++ {
                    if ((-290 > -258) || false) {
                        if (!false) {
                            for fi13 := 0; fi13 < 5; fi13++ {
                                v14 := GenS1{F0: -52, F1: "foo bar"}
                                var v15 uint64 = 1
                                v16 := GenS1{F0: 1, F1: "hello"}
                                _ = fi13
                                _ = v14
                                _ = v15
                                _ = v16
                            }
                            var v17 int32 = -1
                            _ = v17
                        } else {
                            switch (fi12 & (fi12 | fi12)) {
                            case -1:
                                v10 = 0
                            }
                        }
                        v18 := []int{-821, fi12}
                        var v19 int64 = (p0 >> ((42) & 63))
                        _ = v18
                        _ = v19
                    } else {
                        v20 := GenS2{F0: ("a" + "hello")}
                        var v21 uint16 = 1
                        _ = v20
                        _ = v21
                    }
                    fmt.Println(v6)
                    _ = fi12
                }
            }
            _ = v10
        } else {
            if (v6 <= ((((1e10 / v6) / float32(((0.0 / ((float64(((v6 + (float32(3.14159) + (v6 + 1e10))) / 1e10)) - 1e10) + (((0.1 * (0.1 * -2.25)) - 0.1) + float64((v6 / v6))))) - -2.25))) / v6) * v6)) {
                fmt.Println(p0)
                for fi22 := 0; fi22 < 7; fi22++ {
                    if false {
                        var v23 int = fi22
                        if true {
                            if true {
                                if false {
                                    p0 -= p0
                                    v24 := []int{fi22, fi22, (^(fi22))}
                                    _ = v24
                                }
                                fmt.Println(p0, fi22, fi22)
                            }
                        } else {
                            var v25 int64 = (-(p0))
                            fmt.Println(p0)
                            v26 := []int{v23, 313}
                            _ = v25
                            _ = v26
                        }
                        var v27 float64 = 0.1
                        _ = v23
                        _ = v27
                    }
                    p0 = (p0 >> ((4294967295) & 63))
                    if (524 > 42) {
                        for fi28 := 0; fi28 < 12; fi28++ {
                            fmt.Println(v6, v6)
                            v29 := []int{(fi22 / ((fi22) | 1)), (^(fi28)), fi28}
                            _ = fi28
                            _ = v29
                        }
                        v30 := GenS3{F0: 322, F1: 313, F2: 2147483647}
                        v31 := GenS4{F0: p0}
                        _ = v30
                        _ = v31
                    } else {
                        var v32 float32 = (float32(((((0.0 - 0.0) + (0.0 + float64(((v6 / ((float32(0.0) + v6) + ((float32(((0.0 + (1.5 - 0.0)) / (-2.25 * 1.5))) + float32(0.1)) + float32(0.1)))) - 0.1)))) + -2.25) * (float64(v6) / (float64(v6) + -2.25)))) * (v6 / v6))
                        if true {
                            fmt.Println(p0, v32)
                        } else {
                            v33 := []int{}
                            switch fi22 {
                            case 920:
                                v32 = (0.0 - float32(((3.14159 + 0.0) - 1.5)))
                            }
                            _ = v33
                        }
                        v34 := []int{(fi22 - (fi22 % (((fi22 - fi22)) | 1))), -9223372036854775808, fi22, 127}
                        _ = v32
                        _ = v34
                    }
                    _ = fi22
                }
            }
            var v35 uint = 402
            _ = v35
        }
    case 0:
        var v36 int64 = p0
        _ = v36
    default:
        switch 127 {
        case -2147483648:
            if (18446744073709551615 >= 1) {
                v37 := GenS4{F0: (p0 ^ (p0 - p0))}
                _ = v37
            }
        case 1:
            var v38 uint32 = 1
            _ = v38
        default:
            switch "" {
            case "a":
                var v39 bool = true
                _ = v39
            default:
                v6 -= v6
            }
        }
    }
    v6 = float32((3.14159 + -2.25))
    _ = p0
    _ = v6
    return 647
}
func genFn1(p0 uint64) uint32 {
    v40 := []int{32767, 837, -412, 42}
    switch ("foo bar" + "\n\t") {
    case "":
        v41 := GenS5{F0: 1}
        _ = v41
    default:
        var v42 int32 = -23
        _ = v42
    }
    p0 += 4294967295
    _ = p0
    _ = v40
    return 593
}

func main() {
    v1 := []int{}
    v2 := []int{4294967295, -53, 128, 127}
    var v3 float32 = 1e10
    v4 := []int{1, 633, -39}
    v5 := make(map[uint32]int)
    if (138 >= 255) {
        _ = len(v1)
        v4 = append(v4, -974)
    }
    switch 2147483647 {
    case 32767:
        fmt.Println(v3)
    case 127:
        v3 = -2.25
    case 128:
        var v43 string = "\n\t"
        _ = v43
    default:
        var v44 int = 32768
        _ = v44
    }
    switch ("\n\t" + ("\n\t" + ("foo bar" + ("foo bar" + "")))) {
    case "a":
        if false {
            v45 := GenS2{F0: "a"}
            v3 = float32(float64(v3))
            _ = v45
        } else {
            switch "hello" {
            case "hello":
                v3 -= ((((v3 / v3) + (1e10 / (v3 - v3))) / ((v3 + v3) + (v3 + 1.5))) * (v3 - v3))
            case "a":
                v46 := GenS2{F0: (((("foo bar" + "hello") + "foo bar") + ("" + "a")) + "\n\t")}
                _ = v46
            case "foo bar":
                var v47 int32 = 102
                _ = v47
            }
            fmt.Println(v3)
        }
    }
    var v48 int32 = -1
    v3 -= 0.0
    v49 := GenS6{F0: true, F1: 269}
    switch 42 {
    case 32767:
        if false {
            v48 = (v48 >> ((18446744073709551615) & 63))
            _ = len(v5)
            _ = len(v1)
        } else {
            switch "a" {
            case "hello":
                var v50 uint64 = 509
                _ = v50
            case "\n\t":
                v3 -= ((float32((float64(((1e10 + 0.1) + 3.14159)) - 1e10)) / v3) + (1.5 / v3))
            default:
                delete(v5, 42)
            }
        }
    default:
        _ = len(v5)
    }
    v51 := GenS7{F0: ((0.0 != (1e10 * 0.0)) && (0 > 1)), F1: 229, F2: 9223372036854775807}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v51
}
