package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 uint8; F1 bool; F2 int32 }
type GenS3 struct { F0 int8; F1 int8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 string; F1 int16; F2 uint64 }
type GenS6 struct { F0 int8; F1 uint32; F2 uint8 }
func genFn0(p0 uint) uint8 {
    var v7 bool = (1 == -1)
    v8 := []int{9223372036854775807, 128, 2147483647, -1}
    v9 := GenS1{F0: 9223372036854775807, F1: 4294967295}
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 128
}
func genFn1(p0 uint32, p1 int32, p2 uint16) int16 {
    var v10 int32 = p1
    var v11 uint64 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 32767
}

func main() {
    var v1 int8 = 1
    v2 := make(map[string]bool)
    var v3 uint = 18446744073709551615
    var v4 uint16 = 1
    v5 := make(map[uint32]uint32)
    v6 := make(map[int]int)
    v12 := []int{-1, 32768}
    v13 := v2[("foo bar" + ("a" + ((((("a" + "foo bar") + ("hello" + "")) + "hello") + ("" + "")) + "")))]
    v13 = v13
    v3 -= 0
    v14 := GenS1{F0: 128, F1: 1}
    v15 := GenS1{F0: 127, F1: -1}
    v16 := GenS1{F0: 357, F1: 853}
    var v17 uint16 = 42
    for fi18 := 0; fi18 < 10; fi18++ {
        if len(v12) > 0 {
            v12[0] = (-(fi18))
        }
        if (!v13) {
            var v19 uint64 = 2
            _ = v19
        } else {
            if v13 {
                var v20 uint = v3
                if (("a" >= ("hello" + "a")) || (uint32(v3) < 0)) {
                    fmt.Println(v13, v3)
                    v1 = (v1 / ((int8(fi18)) | 1))
                    v3 += (v3 >> ((v3) & 63))
                }
                if v13 {
                    v21 := GenS2{F0: 1, F1: (!v13), F2: 2147483647}
                    if v13 {
                        var v22 int32 = -992
                        var v23 uint64 = uint64(v20)
                        for fi24 := 0; fi24 < 19; fi24++ {
                            v3 += 4294967295
                            if false {
                                fmt.Println(v17, v3, v23)
                            }
                            v25 := GenS3{F0: int8(v22), F1: (v1 | v1)}
                            v26 := GenS4{F0: (v22 + 0)}
                            if v13 {
                                v2["hello"] = (42 >= 918)
                                switch "" {
                                case "a":
                                    var v27 uint8 = 209
                                    _ = v27
                                case "":
                                    _ = len(v6)
                                }
                            }
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                        _ = v22
                        _ = v23
                    }
                    var v28 float64 = float64(((0.1 * float32((0.0 - 1.5))) / 1.0))
                    _ = v21
                    _ = v28
                }
                _ = v20
            }
        }
        for fi29 := 0; fi29 < 7; fi29++ {
            if (65535 > v17) {
                v17 = (v17 + v17)
                _ = len(v5)
                v30 := v6[fi29]
                _ = v30
            } else {
                if v13 {
                    var v31 string = ""
                    _ = v31
                }
                if (983 == 0) {
                    v32 := v5[uint32(v3)]
                    v33 := GenS2{F0: 255, F1: ((!((v1 >= (v1 % (((v1 << (((v3 * (^(v3)))) & 63))) | 1))) && v13)) && (!(-34 > (fi18 * (fi18 / (((fi18 % ((-9223372036854775808) | 1))) | 1)))))), F2: 1}
                    _ = v32
                    _ = v33
                }
                v34 := GenS5{F0: "foo bar", F1: 42, F2: 18446744073709551615}
                switch 32767 {
                case 0:
                    v35 := GenS6{F0: 127, F1: 0, F2: 2}
                    _ = v35
                }
                _ = v34
            }
            switch "hello" {
            case "foo bar":
                delete(v6, (^(fi29)))
            case "hello":
                var v36 uint = (^(v3))
                _ = v36
            }
            _ = len(v12)
            _ = fi29
        }
        _ = fi18
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
