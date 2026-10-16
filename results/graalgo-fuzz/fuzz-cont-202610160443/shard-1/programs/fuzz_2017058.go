package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 string; F1 int64; F2 uint32 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 int8, p1 int) bool {
    p0 -= p0
    fmt.Println(p1)
    v6 := []int{(p1 / ((p1) | 1)), p1, (p1 << ((42) & 63)), -2147483648}
    var v7 int8 = (p0 >> ((255) & 63))
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return ((v7 - p0) == int8(p1))
}
func genFn1() int64 {
    var v8 bool = true
    fmt.Println(v8)
    v8 = true
    var v9 uint64 = 18446744073709551615
    _ = v8
    _ = v9
    return -1
}

func main() {
    var v1 int = 127
    var v2 string = "a"
    v3 := make(map[int]string)
    v4 := make(map[uint32]int)
    var v5 float32 = 0.0
    for fi10 := 0; fi10 < 17; fi10++ {
        v11 := []int{886, (fi10 ^ v1)}
        v2 = v2
        v12 := GenS1{F0: 42, F1: int8(fi10)}
        _ = fi10
        _ = v11
        _ = v12
    }
    if true {
        if true {
            v1 += (v1 / ((0) | 1))
            var v13 float32 = ((1.5 * v5) - (v5 - ((((v5 + v5) / (v5 + 3.14159)) * (v5 / v5)) / 1e10)))
            v14 := []int{(v1 - (v1 * v1))}
            _ = v13
            _ = v14
        }
    } else {
        v15 := []int{v1, v1}
        v16 := GenS2{F0: v2, F1: 0, F2: 2147483647}
        var v17 uint = 65535
        _ = v15
        _ = v16
        _ = v17
    }
    v18 := GenS3{F0: 100}
    v19 := GenS3{F0: 4294967295}
    v2 = "foo bar"
    v20 := []int{v1}
    switch v1 {
    case -1:
        if len(v20) > 0 {
            v20[0] = v1
        }
    case 9223372036854775807:
        v21 := v3[9223372036854775807]
        _ = v21
    case 32768:
        switch "\n\t" {
        case "hello":
            v22 := GenS2{F0: v2, F1: 42, F2: 4294967295}
            _ = v22
        default:
            v1 -= v1
        }
    default:
        var v23 float64 = (3.14159 + -2.25)
        _ = v23
    }
    if true {
        if true {
            fmt.Println(v1, v2, v1)
            v20 = append(v20, (v1 / (((v1 - v1)) | 1)))
            var v24 float64 = 0.1
            _ = v24
        }
        for fi25 := 0; fi25 < 14; fi25++ {
            v26 := GenS2{F0: (v2 + ""), F1: -9223372036854775808, F2: 2}
            var v27 int8 = -128
            switch v2 {
            case "":
                if true {
                    switch (v2 + v2) {
                    case "\n\t":
                        for fi28 := 0; fi28 < 20; fi28++ {
                            v2 = v2
                            v29 := GenS2{F0: "", F1: 9223372036854775807, F2: 65535}
                            v30 := GenS2{F0: ("hello" + "hello"), F1: -348, F2: 0}
                            _ = fi28
                            _ = v29
                            _ = v30
                        }
                    case "":
                        v31 := GenS1{F0: 255, F1: (v27 * (v27 % (((v27 | (v27 | v27))) | 1)))}
                        _ = v31
                    default:
                        for fi32 := 0; fi32 < 13; fi32++ {
                            if true {
                                var v33 int16 = 42
                                var v34 int16 = (v33 | (v33 >> ((1) & 63)))
                                _ = v33
                                _ = v34
                            }
                            fmt.Println(v5)
                            _ = fi32
                        }
                    }
                    var v35 int64 = int64(v1)
                    _ = v35
                } else {
                    if false {
                        var v36 float64 = (((1e10 - (((1e10 / (0.0 * (float64((v5 / (v5 * (-2.25 / (((0.1 - 1.5) + (0.0 / (v5 / 0.1))) * v5))))) - (0.0 * 3.14159)))) * 1.5) * (3.14159 * (0.0 * 1.5)))) / 1e10) / float64(v5))
                        if len(v20) > 0 {
                            v20[0] = fi25
                        }
                        _ = v36
                    }
                    var v37 int = (-(v1))
                    _ = v37
                }
            default:
                v1 = (v1 / (((v1 & int(v27))) | 1))
            }
            _ = fi25
            _ = v26
            _ = v27
        }
        v38 := v3[110]
        _ = v38
    } else {
        var v39 float64 = -2.25
        _ = v39
    }
    if (42 < 14) {
        if ((("a" == v2) && false) || (839 != 1)) {
            _ = len(v3)
            if (!false) {
                v40 := GenS2{F0: "\n\t", F1: 394, F2: 56}
                v5 = v5
                _ = v40
            } else {
                fmt.Println(v1, v5, v2)
            }
            for fi41 := 0; fi41 < 6; fi41++ {
                var v42 int32 = 42
                if (11 > 42) {
                    v3[v1] = (v2 + (v2 + v2))
                    if len(v20) >= 2 {
                        v20 = v20[1:len(v20)]
                    }
                    var v43 int16 = -32768
                    _ = v43
                } else {
                    if (127 != 21) {
                        v44 := GenS1{F0: 127, F1: 0}
                        v2 = ("foo bar" + "\n\t")
                        _ = v44
                    } else {
                        for fi45 := 0; fi45 < 5; fi45++ {
                            var v46 float64 = ((((-2.25 + (1.5 - (0.0 - -2.25))) + 1e10) - 3.14159) - float64(1e10))
                            var v47 int16 = 1000
                            fmt.Println(v5, v47, fi45)
                            _ = fi45
                            _ = v46
                            _ = v47
                        }
                    }
                    v48 := GenS2{F0: "hello", F1: -1, F2: 4294967295}
                    _ = v48
                }
                _ = fi41
                _ = v42
            }
        }
        var v49 uint16 = 922
        _ = v49
    } else {
        _ = len(v20)
        if false {
            v5 += ((1.5 / 1.0) * 0.1)
        }
        var v50 float32 = (3.14159 - ((1e10 * 1e10) * v5))
        _ = v50
    }
    if false {
        v5 += -2.25
        var v51 uint32 = 115
        _ = v51
    }
    var v52 int = 4294967295
    var v53 int16 = -1
    if true {
        v52 = v1
    } else {
        if len(v20) > 0 {
            v20[0] = (v1 | (v1 | v1))
        }
        for fi54 := 0; fi54 < 6; fi54++ {
            var v55 bool = true
            _ = fi54
            _ = v55
        }
    }
    var v56 bool = (true && (144 < 1))
    var v57 uint32 = 97
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v56
    _ = v57
}
