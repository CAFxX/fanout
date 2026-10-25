package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 string; F2 uint8 }
type GenS3 struct { F0 int32; F1 int32 }
func genFn0() uint32 {
    var v7 uint16 = 255
    if false {
        v8 := GenS1{F0: 880}
        fmt.Println(v7)
        v9 := GenS1{F0: 3}
        _ = v8
        _ = v9
    } else {
        v7 = v7
        var v10 uint32 = 203
        if true {
            fmt.Println(v10)
            v11 := GenS1{F0: (v7 >> ((42) & 63))}
            _ = v11
        } else {
            fmt.Println(v7, v7)
            for fi12 := 0; fi12 < 14; fi12++ {
                var v13 string = ("\n\t" + "")
                _ = fi12
                _ = v13
            }
        }
        _ = v10
    }
    v7 = (v7 / (((v7 & (^(v7)))) | 1))
    var v14 int64 = 175
    _ = v14
    _ = v7
    return 302
}
func genFn1() uint {
    if (330 != 1000) {
        if true {
            switch "\n\t" {
            case "foo bar":
                v15 := GenS1{F0: 32767}
                _ = v15
            case "a":
                var v16 bool = false
                _ = v16
            case "hello":
                if ("hello" < "hello") {
                    for fi17 := 0; fi17 < 1; fi17++ {
                        v18 := GenS1{F0: 318}
                        _ = fi17
                        _ = v18
                    }
                    var v19 bool = (1 != 2)
                    v20 := GenS1{F0: 2}
                    _ = v19
                    _ = v20
                }
            }
            if false {
                v21 := []int{4294967295, 128, 1}
                var v22 string = ("foo bar" + (("hello" + ("foo bar" + "hello")) + ("" + ("hello" + ""))))
                v23 := GenS1{F0: 255}
                _ = v21
                _ = v22
                _ = v23
            } else {
                var v24 string = (("hello" + "foo bar") + "hello")
                for fi25 := 0; fi25 < 16; fi25++ {
                    for fi26 := 0; fi26 < 1; fi26++ {
                        if true {
                            var v27 int64 = 252
                            _ = v27
                        }
                        _ = fi26
                    }
                    if (42 == -82) {
                        var v28 int64 = -9223372036854775808
                        if (588 >= 574) {
                            v28 = (v28 ^ (v28 - v28))
                            fmt.Println(v28, fi25)
                            if (1e10 < 0.0) {
                                v28 = 0
                                var v29 int64 = (-(v28))
                                v30 := GenS2{F0: 12, F1: v24, F2: 1}
                                _ = v29
                                _ = v30
                            } else {
                                var v31 float64 = float64((1.5 * 1e10))
                                v32 := []int{(fi25 - 128), fi25}
                                v33 := GenS2{F0: 127, F1: v24, F2: 4}
                                _ = v31
                                _ = v32
                                _ = v33
                            }
                        } else {
                            v28 = (v28 * (v28 / (((-(v28))) | 1)))
                        }
                        v34 := GenS2{F0: 116, F1: v24, F2: 2}
                        _ = v28
                        _ = v34
                    }
                    v24 = v24
                    _ = fi25
                }
                fmt.Println(v24)
                _ = v24
            }
        }
    }
    var v35 uint = 18446744073709551615
    var v36 int = -978
    var v37 uint = v35
    var v38 uint16 = 98
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    return v37
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 2
    var v3 int64 = 42
    var v4 float32 = float32((3.14159 / 1e10))
    var v5 float64 = (1.5 / 3.14159)
    var v6 uint64 = 1
    var v39 int8 = 127
    if (1 != 1) {
        v6 = v6
    }
    v6 = v6
    v40 := GenS3{F0: 0, F1: -174}
    var v41 float32 = v4
    var v42 int = 188
    v6 += (v6 * (^(v6)))
    var v43 float64 = (((0.1 * (v5 - v5)) + (((v5 * 3.14159) + v5) * v5)) / (((((((v5 / v5) / float64(v41)) * (v5 * 1e10)) - v5) + 3.14159) + (0.0 + v5)) + (((0.0 + v5) + ((float64(v4) * (v5 * 3.14159)) / v5)) + (v5 + 0.0))))
    v4 = v4
    v44 := []int{v42, v42}
    for fi45 := 0; fi45 < 4; fi45++ {
        var v46 uint8 = 2
        _ = fi45
        _ = v46
    }
    fmt.Println(v5, v5, v42)
    var v47 bool = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v47
    _ = v5
    _ = v6
}
