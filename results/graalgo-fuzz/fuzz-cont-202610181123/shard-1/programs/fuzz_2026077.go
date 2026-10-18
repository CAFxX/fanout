package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64; F1 uint }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int8; F1 uint16 }
type GenS8 struct { F0 int16; F1 int }
func genFn0(p0 string) int32 {
    v4 := []int{32768, 9223372036854775807, 750}
    fmt.Println(p0)
    for fi5 := 0; fi5 < 13; fi5++ {
        v6 := GenS1{F0: 2147483647}
        v7 := GenS1{F0: 645}
        _ = fi5
        _ = v6
        _ = v7
    }
    if (42 == 0) {
        v8 := GenS1{F0: 2}
        _ = v8
    }
    _ = p0
    _ = v4
    return 2147483647
}
func genFn1() int {
    var v9 int8 = 127
    switch "" {
    case "hello":
        fmt.Println(v9)
    case "\n\t":
        v10 := []int{1}
        _ = v10
    }
    for fi11 := 0; fi11 < 17; fi11++ {
        v9 = 1
        _ = fi11
    }
    _ = v9
    return 312
}

func main() {
    var v1 int8 = 127
    v2 := []uint32{}
    v3 := make(map[int]uint32)
    v12 := GenS1{F0: 595}
    v13 := GenS1{F0: 2147483648}
    var v14 string = "foo bar"
    v15 := GenS2{F0: -1}
    v16 := GenS3{F0: 0}
    var v17 float32 = (0.1 / (1e10 * (0.0 + 3.14159)))
    v18 := GenS3{F0: 743}
    v17 += v17
    fmt.Println(v1)
    v2 = append(v2, 1)
    for fi19 := 0; fi19 < 18; fi19++ {
        for fi20 := 0; fi20 < 16; fi20++ {
            v21 := GenS4{F0: 0, F1: 65535}
            var v22 uint8 = 17
            _ = len(v2)
            if true {
                var v23 bool = false
                _ = v23
            } else {
                if (243 >= -1) {
                    v24 := v3[genFn1()]
                    if (!((v1 + v1) < (v1 / ((v1) | 1)))) {
                        switch fi20 {
                        case 427:
                            var v25 uint64 = 7
                            v26 := GenS5{F0: 160}
                            _ = v25
                            _ = v26
                        case -304:
                            if len(v2) > 0 {
                                v2[0] = 357
                            }
                        default:
                            v14 = "a"
                            var v27 string = v14
                            v24 = v24
                            _ = v27
                        }
                        v28 := GenS5{F0: 1}
                        _ = v28
                    } else {
                        switch 0 {
                        case -2147483648:
                            if true {
                                var v29 int8 = (-(v1))
                                _ = v29
                            }
                        case 1:
                            var v30 int64 = -416
                            _ = v30
                        case 251:
                            switch v14 {
                            case "hello":
                                var v31 int64 = -9223372036854775808
                                _ = v31
                            case "a":
                                var v32 int64 = -83
                                _ = v32
                            default:
                                fmt.Println(fi20, v22)
                            }
                        }
                        if false {
                            _ = len(v2)
                            v1 = 109
                        }
                    }
                    _ = v24
                } else {
                    fmt.Println(fi20, v17)
                    var v33 int64 = -9223372036854775808
                    _ = v33
                }
                switch (((v14 + v14) + "") + (v14 + "")) {
                case "":
                    v1 += 42
                case "hello":
                    if true {
                        var v34 uint = 868
                        var v35 int = fi19
                        v36 := GenS6{F0: "foo bar"}
                        _ = v34
                        _ = v35
                        _ = v36
                    } else {
                        v37 := GenS6{F0: v14}
                        v17 = v17
                        v38 := GenS7{F0: v1, F1: 0}
                        _ = v37
                        _ = v38
                    }
                case "a":
                    fmt.Println(v22)
                }
                if true {
                    v39 := GenS8{F0: 520, F1: genFn1()}
                    _ = v39
                }
            }
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = fi19
    }
    if len(v2) > 0 {
        v2[0] = 842
    }
    fmt.Println(v17, v17)
    switch (v14 + (v14 + ((v14 + "foo bar") + v14))) {
    case "foo bar":
        v17 = (v17 + float32((3.14159 + (0.1 / 1.0))))
    case "hello":
        var v40 uint64 = 1
        _ = v40
    case "a":
        var v41 int32 = genFn0(v14)
        _ = v41
    }
    var v42 uint16 = 15
    fmt.Println(v1)
    for fi43 := 0; fi43 < 15; fi43++ {
        delete(v3, 128)
        v42 = (v42 & v42)
        v44 := v3[127]
        _ = fi43
        _ = v44
    }
    v42 += (v42 | v42)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v42:", v42)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v42
}
