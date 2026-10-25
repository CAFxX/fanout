package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 int16 }
type GenS2 struct { F0 int32; F1 bool; F2 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int16; F1 uint32; F2 bool }
func genFn0(p0 uint) uint {
    p0 = 18446744073709551615
    v7 := []int{42, -3, 971, 32768}
    _ = p0
    _ = v7
    return p0
}
func genFn1(p0 int64) uint32 {
    var v8 bool = false
    p0 += p0
    var v9 int16 = 241
    v8 = v8
    _ = p0
    _ = v8
    _ = v9
    return 2147483647
}

func main() {
    var v1 int64 = -438
    v2 := []string{"", (("" + (("hello" + ("hello" + "\n\t")) + ((("" + ("foo bar" + "foo bar")) + "foo bar") + ""))) + "foo bar"), "foo bar", ("a" + "")}
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]bool)
    v5 := []string{"\n\t", ("foo bar" + "foo bar"), (("a" + ("" + (("" + ("foo bar" + ("" + "a"))) + "\n\t"))) + "")}
    var v6 uint = 0
    var v10 int16 = 0
    v11 := GenS1{F0: 127, F1: "hello", F2: v10}
    v12 := GenS2{F0: 42, F1: true, F2: v6}
    v3[9223372036854775807] = 326
    if true {
        _ = len(v5)
        v10 -= 42
        var v13 int16 = v10
        _ = v13
    } else {
        if false {
            for fi14 := 0; fi14 < 3; fi14++ {
                var v15 int = 42
                for fi16 := 0; fi16 < 4; fi16++ {
                    var v17 uint16 = 618
                    _ = fi16
                    _ = v17
                }
                if false {
                    var v18 float64 = 0.0
                    var v19 float64 = v18
                    var v20 float64 = (float64((-2.25 + (0.1 + 0.0))) * float64(3.14159))
                    _ = v18
                    _ = v19
                    _ = v20
                } else {
                    if (!((1e10 * (-2.25 - (-2.25 - (0.1 + (1e10 * 1.5))))) > ((-2.25 - ((-2.25 + 0.1) + 0.1)) - float32(-2.25)))) {
                        var v21 int32 = -377
                        if true {
                            fmt.Println(fi14, v21, v1)
                        } else {
                            var v22 uint64 = 39
                            if false {
                                v10 += (v10 - v10)
                            }
                            _ = v22
                        }
                        _ = v21
                    } else {
                        for fi23 := 0; fi23 < 13; fi23++ {
                            var v24 int64 = (v1 + -1)
                            var v25 int8 = 21
                            var v26 uint = v6
                            _ = len(v2)
                            _ = fi23
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                        v15 = fi14
                        fmt.Println(fi14, fi14, fi14)
                    }
                }
                _ = fi14
                _ = v15
            }
            var v27 int8 = 127
            delete(v4, 0)
            _ = v27
        }
        v28 := GenS3{F0: v6}
        _ = v28
    }
    switch -871 {
    case 127:
        fmt.Println(v6, v1, v10)
    case 227:
        var v29 uint16 = 0
        _ = v29
    }
    v30 := GenS3{F0: (v6 / (((v6 >> ((v6) & 63))) | 1))}
    _ = len(v3)
    v31 := GenS3{F0: 459}
    var v32 uint16 = 65535
    v32 += (v32 | v32)
    switch 32768 {
    case 42:
        v1 = (-(v1))
    case 2147483647:
        v2 = append(v2, "foo bar")
    default:
        v1 = (v1 / (((v1 + v1)) | 1))
    }
    v6 = (v6 >> ((641) & 63))
    if (true || ((((("\n\t" + "\n\t") + "hello") + "foo bar") + "\n\t") > "foo bar")) {
        v33 := GenS3{F0: v6}
        _ = v33
    }
    fmt.Println(v6, v10)
    v34 := GenS4{F0: v10, F1: 4294967295, F2: true}
    var v35 uint = 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
