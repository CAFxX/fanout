package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool; F1 int32; F2 int64 }
func genFn0(p0 uint8, p1 int32, p2 int8) uint {
    p1 = p1
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint16 = 32767
    v3 := make(map[int]string)
    v4 := make(map[uint32]uint32)
    switch -354 {
    case 32768:
        for fi5 := 0; fi5 < 17; fi5++ {
            v2 = (v2 - v2)
            var v6 uint32 = 4294967295
            v7 := []int{}
            _ = fi5
            _ = v6
            _ = v7
        }
    case 1:
        for fi8 := 0; fi8 < 6; fi8++ {
            var v9 int = fi8
            _ = fi8
            _ = v9
        }
    case -1:
        if ((("hello" + "hello") != ("hello" + ("a" + ""))) || false) {
            var v10 float32 = 0.1
            if (!(32767 >= 32767)) {
                v11 := GenS1{F0: 100, F1: 2}
                if true {
                    v12 := GenS1{F0: 1, F1: 32767}
                    switch "" {
                    case "hello":
                        v13 := GenS1{F0: 42, F1: (v2 | v2)}
                        _ = v13
                    case "":
                        var v14 uint8 = 255
                        _ = v14
                    default:
                        v15 := []int{1, 4294967295}
                        _ = v15
                    }
                    v10 -= ((v10 / (v10 / float32((((((1.5 - 1.5) - 1.5) - 3.14159) - 1e10) - ((3.14159 - 0.1) - -2.25))))) * v10)
                    _ = v12
                } else {
                    if true {
                        v2 = 65535
                        for fi16 := 0; fi16 < 12; fi16++ {
                            for fi17 := 0; fi17 < 13; fi17++ {
                                switch "hello" {
                                case "\n\t":
                                    var v18 uint32 = 1
                                    _ = v18
                                }
                                _ = fi17
                            }
                            _ = fi16
                        }
                    }
                }
                _ = v11
            } else {
                var v19 float32 = ((v10 * -2.25) / (v10 * (v10 / ((1e10 + (3.14159 + (v10 - (float32(0.0) + v10)))) - 1e10))))
                var v20 string = ("\n\t" + "hello")
                _ = v19
                _ = v20
            }
            _ = v10
        }
    }
    v2 = (v2 % ((v2) | 1))
    var v21 uint = 186
    var v22 uint32 = 0
    v23 := GenS2{F0: 127, F1: 505}
    v24 := []int{-9223372036854775808, 127, 1}
    var v25 int16 = -1
    v26 := GenS2{F0: 32767, F1: 135}
    if len(v24) >= 2 {
        v24 = v24[1:len(v24)]
    }
    fmt.Println(v21, v22)
    var v27 bool = (427 <= (v2 - 32767))
    if (643 == 1) {
        var v28 int64 = 9223372036854775807
        v29 := GenS3{F0: v27}
        _ = v28
        _ = v29
    } else {
        var v30 uint = (v21 ^ 1)
        v31 := GenS1{F0: 100, F1: 526}
        _ = v30
        _ = v31
    }
    v32 := GenS4{F0: (v27 && (uint64(v21) <= 2)), F1: -1, F2: 9223372036854775807}
    v22 += v22
    v27 = (128 != 2)
    var v33 uint8 = 128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}
