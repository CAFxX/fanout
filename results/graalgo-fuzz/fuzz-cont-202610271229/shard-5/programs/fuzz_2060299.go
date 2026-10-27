package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 int32 }
func genFn0(p0 int64, p1 bool) bool {
    var v5 uint64 = 311
    if true {
        if false {
            v6 := GenS1{F0: 255, F1: 0}
            for fi7 := 0; fi7 < 18; fi7++ {
                if ((1e10 * float64((1.5 + (0.0 * (0.0 * (0.1 - 0.0)))))) < 3.14159) {
                    p0 -= -9223372036854775808
                    switch ("foo bar" + "hello") {
                    case "a":
                        v8 := []int{fi7, fi7}
                        _ = v8
                    case "\n\t":
                        for fi9 := 0; fi9 < 16; fi9++ {
                            var v10 uint32 = 2
                            fmt.Println(p0, p0)
                            _ = fi9
                            _ = v10
                        }
                    }
                    var v11 int8 = -128
                    _ = v11
                } else {
                    v12 := GenS1{F0: 0, F1: 42}
                    for fi13 := 0; fi13 < 4; fi13++ {
                        var v14 int = (fi7 + 0)
                        _ = fi13
                        _ = v14
                    }
                    if false {
                        if false {
                            v15 := []int{}
                            _ = v15
                        } else {
                            fmt.Println(fi7, p1, fi7)
                            v16 := GenS1{F0: 2, F1: 96}
                            _ = v16
                        }
                    } else {
                        v5 += 18446744073709551615
                        v17 := []int{}
                        _ = v17
                    }
                    _ = v12
                }
                _ = fi7
            }
            _ = v6
        } else {
            p0 = (-(p0))
            var v18 float32 = -2.25
            _ = v18
        }
        if true {
            if (((("a" + "a") + "hello") <= ("hello" + "\n\t")) || (0 < 4294967295)) {
                v5 = v5
                if ((128 >= 1) || p1) {
                    var v19 uint64 = v5
                    _ = v19
                }
            } else {
                if (!(!p1)) {
                    v20 := GenS2{F0: 0}
                    v5 -= (v5 ^ (^(v5)))
                    _ = v20
                }
                for fi21 := 0; fi21 < 19; fi21++ {
                    var v22 float64 = 0.1
                    var v23 int8 = 0
                    _ = fi21
                    _ = v22
                    _ = v23
                }
            }
        } else {
            if (-1 <= 0) {
                var v24 uint = 18446744073709551615
                v24 = (v24 >> ((v24) & 63))
                _ = v24
            } else {
                v25 := []int{-2147483648, 968, -456}
                v5 -= 0
                _ = v25
            }
            var v26 int8 = 100
            fmt.Println(p1, v5)
            _ = v26
        }
    } else {
        p1 = (!p1)
        var v27 int8 = -128
        var v28 int = 42
        _ = v27
        _ = v28
    }
    var v29 int8 = 127
    fmt.Println(p1, p0, p1)
    var v30 int64 = p0
    _ = p0
    _ = p1
    _ = v29
    _ = v30
    _ = v5
    return p1
}

func main() {
    var v1 int = -1
    var v2 uint8 = 121
    v3 := make(map[string]string)
    var v4 string = "\n\t"
    _ = len(v3)
    v3["foo bar"] = (v4 + "\n\t")
    var v31 uint16 = 2
    if (false && ((v31 % ((v31) | 1)) != v31)) {
        var v32 uint64 = 575
        fmt.Println(v4, v2, v1)
        _ = v32
    }
    if (!(2147483647 > 2147483647)) {
        fmt.Println(v4, v1, v31)
    }
    var v33 int = v1
    v34 := []int{32768, (-(v33))}
    var v35 int64 = 501
    var v36 int = (v1 | (v33 % (((v33 << ((0) & 63))) | 1)))
    if false {
        var v37 uint16 = v31
        if (((true && ((1000 >= -1) || ((654 < -2147483648) || (!(!(!true)))))) && (((((((3.14159 * 0.0) + (1.5 - float32(1.5))) / 0.1) + (1.5 * (1e10 - 1.5))) > 0.1) || true) || (42 == 2))) || (false || true)) {
            var v38 int = v36
            if true {
                v39 := GenS3{F0: 0, F1: -2147483648}
                _ = v39
            }
            _ = v38
        } else {
            switch 42 {
            case -9223372036854775808:
                v33 = 1
            case 2147483647:
                var v40 uint8 = (^(v2))
                _ = v40
            }
            v36 = (v36 % ((-9223372036854775808) | 1))
        }
        var v41 uint32 = uint32(v2)
        _ = v37
        _ = v41
    }
    v2 = (v2 ^ (v2 & (v2 >> ((255) & 63))))
    fmt.Println(v1, v2)
    if len(v34) > 0 {
        v34[0] = 9223372036854775807
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
