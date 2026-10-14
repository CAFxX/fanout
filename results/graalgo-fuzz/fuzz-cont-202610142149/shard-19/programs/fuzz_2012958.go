package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 int64 }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 int64; F1 bool; F2 bool }
func genFn0(p0 bool, p1 uint32) int {
    for fi6 := 0; fi6 < 18; fi6++ {
        p0 = (!(!p0))
        v7 := []int{-9223372036854775808}
        _ = fi6
        _ = v7
    }
    switch 9223372036854775807 {
    case 32768:
        var v8 int = int(p1)
        _ = v8
    case 9223372036854775807:
        var v9 bool = false
        _ = v9
    case -989:
        switch 42 {
        case 4294967295:
            var v10 int = -534
            _ = v10
        }
    default:
        p1 = (p1 + (p1 % ((35) | 1)))
    }
    p0 = p0
    _ = p0
    _ = p1
    return 727
}
func genFn1(p0 int16, p1 bool, p2 uint8) uint64 {
    switch "\n\t" {
    case "hello":
        v11 := []int{-2147483648, 32767, -468, -2147483648}
        _ = v11
    case "\n\t":
        p0 += p0
    case "a":
        if true {
            var v12 int16 = p0
            _ = v12
        } else {
            var v13 int = 9223372036854775807
            fmt.Println(p1)
            _ = v13
        }
    }
    v14 := []int{9223372036854775807, genFn0(false, uint32(p2)), 4294967295, 128}
    var v15 int8 = 127
    for fi16 := 0; fi16 < 15; fi16++ {
        if (42 != 762) {
            var v17 uint16 = 0
            var v18 uint64 = 741
            _ = v17
            _ = v18
        }
        _ = fi16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    return 2
}

func main() {
    v1 := []int{-86}
    v2 := []int{76, 32767, -9223372036854775808, -1}
    v3 := make(map[string]uint32)
    v4 := make(map[string]int)
    v5 := []string{"\n\t", (("\n\t" + (("hello" + "hello") + "\n\t")) + ("\n\t" + "foo bar")), "\n\t"}
    v19 := GenS1{F0: 267, F1: 4294967295, F2: 573}
    var v20 bool = (true && true)
    var v21 uint64 = 271
    var v22 int64 = 679
    v22 = v22
    v23 := GenS1{F0: 165, F1: (v21 | v21), F2: v22}
    fmt.Println(v22)
    var v24 int = 32767
    v25 := GenS1{F0: 0, F1: v21, F2: v22}
    v5 = append(v5, "foo bar")
    if false {
        fmt.Println(v21)
    } else {
        switch "" {
        case "a":
            v21 = v21
        case "":
            fmt.Println(v22, v24, v20)
        }
        v24 += v24
        if false {
            v26 := GenS2{F0: ("a" + ""), F1: (v21 * 42)}
            _ = v26
        } else {
            v5 = append(v5, "")
            v2 = append(v2, (-(v24)))
            if false {
                switch ("hello" + "\n\t") {
                case "\n\t":
                    if v20 {
                        v27 := GenS2{F0: ("a" + ("hello" + "")), F1: genFn1(-257, (!(-1 > 1)), 42)}
                        _ = v27
                    } else {
                        for fi28 := 0; fi28 < 6; fi28++ {
                            var v29 int16 = -72
                            v22 += v22
                            var v30 uint8 = 127
                            _ = fi28
                            _ = v29
                            _ = v30
                        }
                    }
                case "hello":
                    var v31 int32 = 42
                    _ = v31
                case "a":
                    v22 = (v22 << ((65535) & 63))
                default:
                    v24 -= 9223372036854775807
                }
            } else {
                if (v22 >= (-(v22))) {
                    var v32 float64 = 0.1
                    _ = v32
                } else {
                    if false {
                        v33 := GenS2{F0: "hello", F1: v21}
                        v21 = v21
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = v33
                    } else {
                        var v34 uint8 = 0
                        if (!(0 == -1)) {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            var v35 float32 = ((0.0 + -2.25) * ((3.14159 - ((1.5 * 0.0) / (0.1 + 3.14159))) / float32((0.1 + ((0.1 - -2.25) + (0.0 + (0.1 + ((-2.25 - 3.14159) + ((3.14159 + 0.1) / 0.1)))))))))
                            _ = v35
                        }
                        _ = v34
                    }
                }
                v36 := GenS3{F0: 2, F1: 65535}
                v37 := GenS4{F0: (v22 + 9223372036854775807), F1: v20, F2: (!(32767 < 65535))}
                _ = v36
                _ = v37
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
