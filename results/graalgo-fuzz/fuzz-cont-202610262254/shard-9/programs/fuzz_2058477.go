package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 bool }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 int64; F1 uint16; F2 int8 }
func genFn0(p0 string, p1 uint8) uint16 {
    v6 := GenS1{F0: 24, F1: -412, F2: (!true)}
    v7 := []int{42, 9223372036854775807, 128}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 2
}

func main() {
    var v1 uint16 = 0
    var v2 int16 = -1
    v3 := make(map[string]bool)
    v4 := make(map[int]bool)
    v5 := make(map[uint32]uint32)
    if (128 < 128) {
        for fi8 := 0; fi8 < 8; fi8++ {
            v3["a"] = false
            var v9 bool = false
            _ = fi8
            _ = v9
        }
        var v10 int64 = 1
        _ = v10
    }
    var v11 float32 = (float32((((-2.25 + -2.25) - -2.25) - (0.0 * (-2.25 - (0.1 - -2.25))))) / (-2.25 * ((((3.14159 + 1e10) / 0.1) + 0.0) * 1.5)))
    v12 := GenS1{F0: 42, F1: 2147483647, F2: false}
    for fi13 := 0; fi13 < 4; fi13++ {
        fmt.Println(v11, v1)
        if (!true) {
            for fi14 := 0; fi14 < 18; fi14++ {
                if (!true) {
                    switch ("" + ("a" + "")) {
                    case "a":
                        fmt.Println(v1, v11)
                    default:
                        var v15 int16 = (v2 >> ((4294967295) & 63))
                        _ = v15
                    }
                }
                v2 -= (v2 / ((-32768) | 1))
                _ = fi14
            }
            var v16 float64 = 1e10
            v2 = (v2 >> ((1) & 63))
            _ = v16
        }
        v17 := GenS1{F0: 255, F1: -87, F2: true}
        _ = fi13
        _ = v17
    }
    var v18 uint8 = 128
    var v19 uint32 = 4294967295
    switch 32768 {
    case -391:
        for fi20 := 0; fi20 < 9; fi20++ {
            _ = len(v4)
            var v21 int = (fi20 / ((32768) | 1))
            for fi22 := 0; fi22 < 6; fi22++ {
                if true {
                    v19 -= v19
                    v18 = (v18 << ((4294967295) & 63))
                    v23 := GenS2{F0: (fi22 << ((520) & 63)), F1: -2147483648}
                    _ = v23
                } else {
                    switch ("\n\t" + ("" + ((("foo bar" + "\n\t") + "\n\t") + "hello"))) {
                    case "a":
                        switch ("a" + "") {
                        case "a":
                            switch "hello" {
                            case "":
                                var v24 int = (^(v21))
                                _ = v24
                            case "hello":
                                if true {
                                    var v25 bool = (((3.14159 * 0.1) / float64(1.5)) != -2.25)
                                    _ = v25
                                }
                            default:
                                if (-9223372036854775808 == 42) {
                                    v11 = 1e10
                                }
                            }
                        case "foo bar":
                            v26 := []int{2147483647, v21, v21}
                            _ = v26
                        case "":
                            v19 -= v19
                        default:
                            v1 -= v1
                        }
                    case "foo bar":
                        var v27 float32 = (float32(3.14159) * v11)
                        _ = v27
                    default:
                        v28 := []int{(^(fi20)), v21}
                        _ = v28
                    }
                    switch (^(fi20)) {
                    case 127:
                        v29 := []int{(fi22 * (fi20 / ((fi20) | 1)))}
                        _ = v29
                    case 1:
                        v30 := GenS2{F0: fi20, F1: (v21 >> ((42) & 63))}
                        _ = v30
                    case 42:
                        v11 += (0.0 * v11)
                    default:
                        v1 = 42
                    }
                    v31 := GenS3{F0: 42, F1: (v1 << ((0) & 63)), F2: 0}
                    _ = v31
                }
                _ = fi22
            }
            _ = fi20
            _ = v21
        }
    case -2147483648:
        v32 := []int{893}
        _ = v32
    case 1:
        v11 = v11
    default:
        for fi33 := 0; fi33 < 9; fi33++ {
            v19 = v19
            _ = fi33
        }
    }
    var v34 int = 42
    v35 := GenS3{F0: 445, F1: 32767, F2: -122}
    switch "" {
    case "foo bar":
        v18 = (v18 & v18)
    case "\n\t":
        var v36 uint = 18446744073709551615
        _ = v36
    case "":
        fmt.Println(v11, v2, v19)
    default:
        v34 -= v34
    }
    v37 := []int{v34, (^(v34)), v34}
    for fi38 := 0; fi38 < 12; fi38++ {
        var v39 uint64 = 4294967295
        _ = fi38
        _ = v39
    }
    for fi40 := 0; fi40 < 3; fi40++ {
        var v41 uint32 = (v19 + v19)
        v42 := v4[(^(fi40))]
        _ = fi40
        _ = v41
        _ = v42
    }
    v43 := GenS1{F0: v18, F1: 1, F2: (false || (127 > 21))}
    if len(v37) >= 2 {
        v37 = v37[1:len(v37)]
    }
    for fi44 := 0; fi44 < 19; fi44++ {
        var v45 int64 = -1
        _ = fi44
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v37
    _ = v4
    _ = v43
    _ = v5
}
