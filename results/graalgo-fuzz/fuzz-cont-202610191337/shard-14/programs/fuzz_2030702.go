package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint32 }
type GenS2 struct { F0 int32; F1 int64; F2 uint8 }
type GenS3 struct { F0 bool; F1 string }
func genFn0(p0 string) int16 {
    v7 := GenS1{F0: 127, F1: 556, F2: 2147483647}
    var v8 int8 = 127
    if false {
        switch 518 {
        case 1:
            if ((-317 == 1) && (!false)) {
                if (!false) {
                    var v9 int16 = -665
                    v10 := GenS1{F0: 32767, F1: 18446744073709551615, F2: 163}
                    _ = v10
                    _ = v9
                } else {
                    for fi11 := 0; fi11 < 10; fi11++ {
                        if (true || false) {
                            switch 255 {
                            case 709:
                                switch -528 {
                                case -589:
                                    var v12 float32 = (float32(((-2.25 + (1e10 + 1e10)) + -2.25)) - (float32(0.1) - (0.0 - 1.5)))
                                    _ = v12
                                case -2147483648:
                                    var v13 uint8 = 128
                                    _ = v13
                                }
                            case -2147483648:
                                for fi14 := 0; fi14 < 1; fi14++ {
                                    p0 = (p0 + "")
                                    v8 += 42
                                    v15 := []int{fi11, -528, 1}
                                    p0 = ((p0 + p0) + "")
                                    v8 = v8
                                    _ = fi14
                                    _ = v15
                                }
                            }
                            v16 := []int{}
                            var v17 int64 = 0
                            _ = v16
                            _ = v17
                        }
                        var v18 string = p0
                        v19 := []int{(fi11 >> ((4294967295) & 63))}
                        _ = fi11
                        _ = v18
                        _ = v19
                    }
                    v8 += (v8 | v8)
                    fmt.Println(p0)
                }
            }
        }
        fmt.Println(v8, p0)
    }
    _ = p0
    _ = v7
    _ = v8
    return 32767
}

func main() {
    v1 := make(map[string]bool)
    v2 := []uint32{42, 2}
    v3 := []uint32{2147483647}
    var v4 uint = 4294967295
    v5 := []string{"a", "hello", "foo bar", ("foo bar" + ("" + ("hello" + "hello")))}
    var v6 int = 970
    v20 := GenS1{F0: 9223372036854775807, F1: v4, F2: 2147483648}
    fmt.Println(v4)
    if true {
        var v21 uint16 = 8
        v22 := GenS2{F0: -1, F1: 42, F2: 73}
        var v23 uint = (^(v4))
        _ = v21
        _ = v22
        _ = v23
    }
    v24 := GenS1{F0: v6, F1: v4, F2: 65535}
    v25 := GenS1{F0: 128, F1: v4, F2: 2147483647}
    if true {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        v4 = (v4 | (v4 | (v4 + v4)))
    }
    for fi26 := 0; fi26 < 13; fi26++ {
        if (!false) {
            fmt.Println(fi26)
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v2 = append(v2, 65535)
            var v27 uint32 = 4294967295
            _ = v27
        }
        if false {
            var v28 int = (fi26 | (v6 & 0))
            _ = v28
        } else {
            fmt.Println(fi26, fi26)
            if (-1 < 0) {
                var v29 float64 = 1.5
                _ = v29
            } else {
                var v30 float32 = 0.0
                for fi31 := 0; fi31 < 7; fi31++ {
                    switch (fi26 >> (((^(v4))) & 63)) {
                    case 128:
                        fmt.Println(v4, fi26, v6)
                    default:
                        fmt.Println(v30, fi26)
                    }
                    switch -187 {
                    case 128:
                        var v32 int16 = 560
                        _ = v32
                    case -226:
                        v33 := v1["a"]
                        _ = v33
                    default:
                        v34 := v1[""]
                        _ = v34
                    }
                    var v35 bool = (109 < -128)
                    _ = fi31
                    _ = v35
                }
                var v36 int16 = genFn0("foo bar")
                _ = v30
                _ = v36
            }
        }
        _ = fi26
    }
    v37 := GenS3{F0: ((42 > 255) || true), F1: "foo bar"}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
