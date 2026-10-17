package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 uint; F2 uint }
type GenS4 struct { F0 uint }
func genFn0(p0 int64, p1 string, p2 int32) int32 {
    if true {
        if (!(!false)) {
            v6 := []int{4294967295, 1}
            _ = v6
        }
    }
    if true {
        var v7 int16 = 1000
        p2 -= p2
        if (222 != 1) {
            switch -9223372036854775808 {
            case -462:
                if ((3.14159 / float32(3.14159)) < 0.0) {
                    if true {
                        if true {
                            p2 -= (p2 / ((p2) | 1))
                            v8 := GenS1{F0: -2147483648}
                            switch 721 {
                            case 0:
                                if true {
                                    var v9 int64 = p0
                                    switch 9223372036854775807 {
                                    case 42:
                                        var v10 int32 = p2
                                        _ = v10
                                    default:
                                        v11 := []int{32768, -556, 128}
                                        _ = v11
                                    }
                                    if (!(677 < 65535)) {
                                        for fi12 := 0; fi12 < 7; fi12++ {
                                            v13 := GenS2{F0: "\n\t"}
                                            fmt.Println(v9)
                                            var v14 int16 = 42
                                            _ = fi12
                                            _ = v13
                                            _ = v14
                                        }
                                        v7 = v7
                                    }
                                    _ = v9
                                } else {
                                    fmt.Println(p0, p1)
                                    var v15 uint = 224
                                    var v16 uint8 = 77
                                    _ = v15
                                    _ = v16
                                }
                            }
                            _ = v8
                        } else {
                            fmt.Println(v7, v7)
                        }
                        switch p1 {
                        case "a":
                            v17 := []int{-637, int(v7)}
                            _ = v17
                        case "foo bar":
                            var v18 uint8 = 1
                            _ = v18
                        case "hello":
                            var v19 uint = 484
                            _ = v19
                        default:
                            var v20 int16 = -830
                            _ = v20
                        }
                    }
                    var v21 int = 23
                    v22 := GenS2{F0: ""}
                    _ = v21
                    _ = v22
                }
            }
            var v23 uint32 = 1
            var v24 int8 = 0
            _ = v23
            _ = v24
        } else {
            var v25 float32 = ((0.0 / 1.0) * (float32(((-2.25 - 0.0) / 1.0)) * (1.5 / 0.1)))
            _ = v25
        }
        _ = v7
    } else {
        switch p1 {
        case "":
            var v26 float32 = 1e10
            _ = v26
        case "a":
            v27 := GenS3{F0: 32767, F1: 905, F2: 4294967295}
            _ = v27
        default:
            for fi28 := 0; fi28 < 20; fi28++ {
                for fi29 := 0; fi29 < 10; fi29++ {
                    var v30 int32 = int32(p0)
                    p1 = "foo bar"
                    switch fi28 {
                    case -2147483648:
                        v31 := []int{(-(fi29)), fi29, (fi29 | -906)}
                        _ = v31
                    case 0:
                        var v32 string = p1
                        _ = v32
                    }
                    _ = fi29
                    _ = v30
                }
                _ = fi28
            }
        }
    }
    p1 = ((p1 + p1) + (("\n\t" + p1) + (p1 + "\n\t")))
    var v33 float64 = 1e10
    v33 = 0.1
    _ = p0
    _ = p1
    _ = p2
    _ = v33
    return p2
}
func genFn1() int8 {
    var v34 uint64 = 383
    v35 := []int{1, 4294967295, -2147483648, 127}
    _ = v34
    _ = v35
    return -121
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[int]uint32)
    var v3 uint8 = 127
    var v4 uint16 = 65535
    var v5 int32 = 1
    v36 := GenS2{F0: "\n\t"}
    switch 42 {
    case 1:
        if true {
            var v37 int32 = (v5 ^ 597)
            if true {
                v1 += 1e10
                v4 = 255
            }
            _ = v37
        }
    case -648:
        fmt.Println(v5, v5)
    case -956:
        for fi38 := 0; fi38 < 3; fi38++ {
            v3 += (v3 >> ((65535) & 63))
            _ = fi38
        }
    default:
        v3 = 0
    }
    fmt.Println(v1)
    v39 := GenS2{F0: "a"}
    v40 := GenS2{F0: "a"}
    v41 := GenS4{F0: 1}
    v42 := []int{}
    v43 := GenS3{F0: -678, F1: 18446744073709551615, F2: 4294967295}
    v44 := v2[32767]
    v3 += (v3 + v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v5
}
