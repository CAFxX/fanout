package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 uint64; F1 int32 }
type GenS5 struct { F0 string; F1 bool }
type GenS6 struct { F0 uint32; F1 string }
func genFn0() uint16 {
    v5 := []int{628, -145, -278}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    switch "\n\t" {
    case "foo bar":
        var v6 uint = 255
        _ = v6
    }
    _ = v5
    return 1
}

func main() {
    v1 := []int{413}
    v2 := make(map[int]bool)
    var v3 int = 0
    var v4 uint16 = 32767
    v1 = append(v1, 4294967295)
    var v7 bool = false
    fmt.Println(v3, v3, v3)
    var v8 int32 = -229
    v9 := GenS1{F0: 42, F1: 2147483648}
    var v10 uint32 = 572
    v11 := GenS2{F0: 0, F1: 165}
    for fi12 := 0; fi12 < 13; fi12++ {
        var v13 float32 = 1e10
        if (1000 < 32767) {
            var v14 int16 = 42
            _ = v14
        }
        _ = fi12
        _ = v13
    }
    switch v3 {
    case 128:
        fmt.Println(v10, v4, v3)
    }
    var v15 uint = 995
    if v7 {
        v2[(v3 % ((42) | 1))] = ((v10 | 758) >= v10)
        v16 := GenS2{F0: -1, F1: 127}
        v17 := GenS1{F0: 42, F1: v10}
        _ = v16
        _ = v17
    } else {
        v2[int(v4)] = true
        var v18 uint16 = v4
        _ = v18
    }
    fmt.Println(v8, v7)
    if len(v1) > 0 {
        v1[0] = (-(v3))
    }
    var v19 float64 = (-2.25 + -2.25)
    for fi20 := 0; fi20 < 9; fi20++ {
        if false {
            switch fi20 {
            case 351:
                v21 := GenS2{F0: -772, F1: 177}
                _ = v21
            case 102:
                _ = len(v1)
            default:
                switch "a" {
                case "a":
                    v22 := GenS2{F0: 1, F1: 255}
                    _ = v22
                default:
                    for fi23 := 0; fi23 < 1; fi23++ {
                        fmt.Println(fi23, v7, v8)
                        fmt.Println(v8)
                        switch (fi20 / (((fi20 & fi20)) | 1)) {
                        case -752:
                            if v7 {
                                var v24 string = "a"
                                var v25 int = (^(fi20))
                                _ = v24
                                _ = v25
                            }
                        case 32768:
                            v26 := GenS2{F0: 1, F1: 155}
                            switch "foo bar" {
                            case "\n\t":
                                if (v7 || v7) {
                                    v7 = false
                                } else {
                                    v15 -= (v15 & (v15 >> (((v15 | v15)) & 63)))
                                    _ = len(v1)
                                }
                            }
                            switch "foo bar" {
                            case "foo bar":
                                if v7 {
                                    if (42 >= v10) {
                                        delete(v2, v3)
                                    } else {
                                        if (!v7) {
                                            var v27 uint = v15
                                            _ = v27
                                        } else {
                                            _ = len(v2)
                                        }
                                        v7 = v7
                                        fmt.Println(fi23)
                                    }
                                    delete(v2, (fi20 - 32767))
                                }
                            case "\n\t":
                                var v28 uint = v15
                                v29 := GenS3{F0: "hello", F1: 255}
                                _ = v28
                                _ = v29
                            case "hello":
                                var v30 int32 = (v8 >> ((1) & 63))
                                _ = v30
                            default:
                                var v31 float64 = (0.1 * float64(((((1.5 * 1.5) - ((0.1 * 3.14159) * 1.5)) + 1.5) + 3.14159)))
                                _ = v31
                            }
                            _ = v26
                        case -2147483648:
                            v32 := GenS4{F0: uint64(v10), F1: v8}
                            var v33 uint = v15
                            _ = v32
                            _ = v33
                        default:
                            var v34 int8 = -18
                            _ = v34
                        }
                        v8 += v8
                        var v35 bool = ((v4 % ((genFn0()) | 1)) == v4)
                        _ = fi23
                        _ = v35
                    }
                }
            }
            v10 += 4294967295
            var v36 uint8 = 178
            _ = v36
        } else {
            if (genFn0() > genFn0()) {
                v37 := GenS5{F0: ("" + "\n\t"), F1: ("hello" > ("foo bar" + "hello"))}
                if false {
                    var v38 string = ("hello" + "hello")
                    _ = v38
                }
                v15 = 0
                _ = v37
            }
            var v39 uint = (^(v15))
            v40 := GenS6{F0: 65535, F1: "a"}
            _ = v39
            _ = v40
        }
        _ = fi20
    }
    switch 2147483647 {
    case 2147483647:
        var v41 uint32 = 2147483648
        _ = v41
    }
    for fi42 := 0; fi42 < 6; fi42++ {
        var v43 uint64 = 190
        v8 -= v8
        for fi44 := 0; fi44 < 9; fi44++ {
            var v45 int32 = -166
            _ = fi44
            _ = v45
        }
        _ = fi42
        _ = v43
    }
    v46 := v2[v3]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v7
    _ = v8
    _ = v9
}
