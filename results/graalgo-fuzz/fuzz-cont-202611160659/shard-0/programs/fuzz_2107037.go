package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64 }
func genFn0() int32 {
    if (-1 > 825) {
        var v6 int32 = 0
        var v7 uint = 4294967295
        v8 := []int{-2147483648}
        _ = v6
        _ = v7
        _ = v8
    } else {
        var v9 uint16 = 848
        _ = v9
    }
    switch -94 {
    case 42:
        v10 := GenS1{F0: 0, F1: "foo bar"}
        _ = v10
    case -354:
        var v11 uint = 42
        _ = v11
    case -260:
        for fi12 := 0; fi12 < 12; fi12++ {
            var v13 uint8 = 91
            if (1 != 1) {
                for fi14 := 0; fi14 < 12; fi14++ {
                    switch "\n\t" {
                    case "\n\t":
                        fmt.Println(fi12, fi12)
                    }
                    var v15 int16 = 0
                    var v16 float32 = ((((3.14159 + 3.14159) - float32(-2.25)) + 1e10) - (1e10 / (0.1 / (((float32(0.0) / -2.25) + 3.14159) - 1.5))))
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            }
            _ = fi12
            _ = v13
        }
    }
    v17 := []int{1, -9223372036854775808, -1, 32767}
    if len(v17) > 0 {
        v17[0] = 32768
    }
    var v18 uint8 = 127
    _ = v17
    _ = v18
    return 2147483647
}

func main() {
    var v1 string = "hello"
    var v2 uint8 = 255
    var v3 string = v1
    var v4 uint8 = 2
    v5 := make(map[string]uint32)
    v19 := GenS1{F0: 4294967295, F1: (v1 + v3)}
    v20 := []int{32768, 32768}
    var v21 uint64 = 42
    v22 := GenS1{F0: -1, F1: ("foo bar" + (v1 + (v1 + v1)))}
    fmt.Println(v3, v21, v21)
    v4 -= 2
    for fi23 := 0; fi23 < 12; fi23++ {
        for fi24 := 0; fi24 < 10; fi24++ {
            v25 := GenS1{F0: 32768, F1: "\n\t"}
            switch v3 {
            case "\n\t":
                v2 = (v4 >> ((811) & 63))
            case "foo bar":
                fmt.Println(v1, v21)
            default:
                if (v21 < (^(v21))) {
                    if false {
                        delete(v5, ("a" + (v1 + ((v3 + v1) + v1))))
                        var v26 uint8 = (^(v4))
                        v27 := GenS1{F0: (fi24 % ((0) | 1)), F1: v1}
                        v1 = v1
                        if false {
                            switch "" {
                            case "":
                                switch int(v26) {
                                case -39:
                                    switch v1 {
                                    case "a":
                                        var v28 int = fi23
                                        _ = v28
                                    }
                                case -489:
                                    var v29 int8 = 100
                                    _ = v29
                                default:
                                    var v30 int = (fi23 + (fi23 - fi23))
                                    if (true || (!(9223372036854775807 == -1))) {
                                        v1 = v1
                                        var v31 uint64 = v21
                                        delete(v5, v3)
                                        _ = v31
                                    }
                                    fmt.Println(fi23)
                                    _ = v30
                                }
                            case "foo bar":
                                if false {
                                    v32 := GenS1{F0: fi23, F1: v1}
                                    switch "" {
                                    case "\n\t":
                                        v33 := GenS2{F0: 192}
                                        _ = v33
                                    default:
                                        v26 -= (v26 * (v4 & 2))
                                        delete(v5, "")
                                        var v34 float64 = ((0.1 + (((0.1 * (1.5 - 1.5)) * 1e10) - 1.5)) + ((3.14159 + 1e10) - 0.0))
                                        _ = v34
                                    }
                                    _ = v32
                                } else {
                                    switch 0 {
                                    case -9:
                                        var v35 int = fi23
                                        _ = v35
                                    case 32767:
                                        v3 = v1
                                    case 127:
                                        var v36 uint64 = (v21 << ((4294967295) & 63))
                                        _ = v36
                                    default:
                                        var v37 int16 = 1000
                                        _ = v37
                                    }
                                    v38 := GenS3{F0: 9223372036854775807}
                                    _ = v38
                                }
                                v1 = "foo bar"
                                var v39 uint16 = 42
                                _ = v39
                            }
                        }
                        _ = v26
                        _ = v27
                    } else {
                        fmt.Println(v3, v1, v3)
                    }
                    v4 = 42
                }
                v5[v1] = 0
                v3 = ((v1 + ((v3 + (v3 + "a")) + v1)) + v3)
            }
            _ = fi24
            _ = v25
        }
        v40 := GenS2{F0: 1}
        fmt.Println(v4, v1, v4)
        _ = fi23
        _ = v40
    }
    v41 := v5[v1]
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v41
    _ = v5
}
