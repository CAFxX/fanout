package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 string; F1 int; F2 int64 }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 uint64; F1 int64 }
func genFn0(p0 int8, p1 int8, p2 int) uint16 {
    p2 -= -1
    p2 = p2
    p0 = p1
    _ = p0
    _ = p1
    _ = p2
    return 32767
}
func genFn1() uint8 {
    var v5 int16 = 1000
    v6 := GenS1{F0: -481, F1: 0}
    _ = v5
    _ = v6
    return 0
}

func main() {
    var v1 int32 = 1
    v2 := make(map[string]int)
    var v3 int64 = -169
    var v4 float32 = 0.0
    for fi7 := 0; fi7 < 2; fi7++ {
        if true {
            fmt.Println(v4, fi7)
            switch fi7 {
            case 32768:
                v8 := GenS1{F0: v3, F1: (v1 + 0)}
                _ = v8
            case 0:
                v3 = v3
            }
        }
        v9 := []int{(fi7 & (-(fi7))), 220, -41, fi7}
        if true {
            var v10 int64 = -9223372036854775808
            _ = v10
        } else {
            if false {
                switch "" {
                case "a":
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                case "\n\t":
                    switch "foo bar" {
                    case "":
                        if false {
                            v4 += v4
                        } else {
                            _ = len(v2)
                        }
                    default:
                        var v11 uint8 = genFn1()
                        _ = v11
                    }
                default:
                    v4 += (float32(float64(3.14159)) + (1e10 - 0.0))
                }
                var v12 uint16 = 849
                _ = v12
            } else {
                if false {
                    var v13 int32 = (v1 | v1)
                    v1 = (-(v1))
                    v14 := GenS2{F0: "hello", F1: 42, F2: (-(v3))}
                    _ = v13
                    _ = v14
                } else {
                    _ = len(v2)
                    var v15 uint64 = 18446744073709551615
                    var v16 int64 = 1
                    _ = v15
                    _ = v16
                }
            }
        }
        _ = fi7
        _ = v9
    }
    v17 := GenS1{F0: v3, F1: v1}
    var v18 string = "a"
    v4 = (0.1 / (3.14159 / (((0.1 * (v4 * 3.14159)) / v4) * v4)))
    var v19 int32 = v1
    v20 := GenS1{F0: v3, F1: (-(v1))}
    for fi21 := 0; fi21 < 16; fi21++ {
        if false {
            v4 += 0.1
            var v22 float32 = float32(1.5)
            _ = v22
        } else {
            v4 -= (1e10 + v4)
            var v23 uint = 2
            _ = v23
        }
        v24 := GenS3{F0: 128, F1: 2}
        switch (^(fi21)) {
        case -34:
            for fi25 := 0; fi25 < 12; fi25++ {
                v4 = (0.1 + (float32(((1.5 + float64(float32(1e10))) + 0.1)) / float32(0.1)))
                var v26 uint8 = 202
                var v27 uint32 = 4294967295
                _ = fi25
                _ = v26
                _ = v27
            }
        case 1:
            var v28 float64 = 0.0
            _ = v28
        default:
            v29 := GenS4{F0: 1, F1: (v3 * v3)}
            _ = v29
        }
        _ = fi21
        _ = v24
    }
    v4 -= 1e10
    v4 = v4
    for fi30 := 0; fi30 < 2; fi30++ {
        fmt.Println(v19, v1, v3)
        _ = fi30
    }
    v1 = 1
    delete(v2, "foo bar")
    v31 := GenS3{F0: 255, F1: 18446744073709551615}
    v32 := []int{}
    var v33 int64 = (v3 / ((-1) | 1))
    switch (v18 + v18) {
    case "a":
        v19 += (v1 * v19)
    }
    switch ("hello" + ((((v18 + "") + "hello") + v18) + ((((v18 + "\n\t") + ("a" + ("\n\t" + "foo bar"))) + (v18 + v18)) + v18))) {
    case "foo bar":
        fmt.Println(v4, v18)
    case "hello":
        var v34 float32 = v4
        _ = v34
    case "a":
        v35 := GenS4{F0: 2, F1: v33}
        _ = v35
    default:
        v36 := GenS4{F0: 18446744073709551615, F1: v33}
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
