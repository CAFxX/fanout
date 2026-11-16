package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 int8; F2 uint }
type GenS3 struct { F0 uint16; F1 uint; F2 int16 }
type GenS4 struct { F0 uint }

func main() {
    var v1 int32 = -337
    var v2 float64 = (3.14159 + (0.1 / 3.14159))
    v3 := make(map[string]bool)
    v4 := make(map[uint32]uint32)
    var v5 string = ""
    if true {
        var v6 string = v5
        _ = v6
    } else {
        v7 := []int{-1}
        var v8 int64 = 42
        var v9 uint64 = 4294967295
        _ = v7
        _ = v8
        _ = v9
    }
    switch "foo bar" {
    case "\n\t":
        v10 := v3[v5]
        _ = v10
    case "a":
        fmt.Println(v5, v1)
    default:
        switch "hello" {
        case "a":
            v11 := GenS1{F0: (v1 - v1)}
            _ = v11
        default:
            fmt.Println(v2)
        }
    }
    var v12 int64 = -797
    for fi13 := 0; fi13 < 8; fi13++ {
        for fi14 := 0; fi14 < 7; fi14++ {
            v2 = v2
            if ((v1 * v1) > 0) {
                if true {
                    v12 -= (v12 - v12)
                    delete(v3, "foo bar")
                } else {
                    var v15 int32 = (v1 + v1)
                    v16 := []int{}
                    if (fi13 > (^(fi14))) {
                        var v17 uint32 = 65535
                        fmt.Println(v1)
                        _ = v17
                    }
                    _ = v15
                    _ = v16
                }
            }
            _ = fi14
        }
        _ = fi13
    }
    var v18 int32 = (v1 % ((v1) | 1))
    v4[276] = 278
    if (-1 > 2147483647) {
        switch 32767 {
        case -1:
            for fi19 := 0; fi19 < 5; fi19++ {
                fmt.Println(fi19)
                var v20 float32 = float32(-2.25)
                var v21 int = fi19
                _ = fi19
                _ = v20
                _ = v21
            }
        case 453:
            for fi22 := 0; fi22 < 6; fi22++ {
                if false {
                    if true {
                        v18 -= v18
                        v23 := []int{-21, (fi22 ^ 421)}
                        for fi24 := 0; fi24 < 5; fi24++ {
                            v25 := GenS1{F0: v18}
                            var v26 uint16 = 575
                            v5 = v5
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                        _ = v23
                    } else {
                        switch fi22 {
                        case 0:
                            if true {
                                switch "a" {
                                case "foo bar":
                                    v12 += (-(v12))
                                case "a":
                                    v27 := []int{fi22, (fi22 ^ -336), 127}
                                    _ = v27
                                }
                                var v28 uint16 = 1
                                if (float64(1.5) <= float64((((float32(((1e10 + v2) - v2)) + float32(v2)) + 0.0) / (1.5 * -2.25)))) {
                                    v29 := []int{}
                                    fmt.Println(v28, v28, v28)
                                    fmt.Println(fi22, v2, v1)
                                    _ = v29
                                }
                                _ = v28
                            }
                        }
                        v30 := GenS2{F0: 2, F1: -45, F2: 317}
                        v31 := GenS2{F0: 188, F1: 127, F2: 42}
                        _ = v30
                        _ = v31
                    }
                    var v32 float64 = 1.5
                    v32 -= 1.5
                    _ = v32
                } else {
                    fmt.Println(fi22)
                    v2 += v2
                    fmt.Println(fi22, v5, v1)
                }
                _ = fi22
            }
        default:
            fmt.Println(v1, v2, v1)
        }
    }
    v33 := []int{127, 0, 160}
    var v34 string = (v5 + (v5 + v5))
    var v35 int = 4294967295
    var v36 int64 = (v12 * (v12 + (v12 << ((896) & 63))))
    var v37 bool = (2147483647 > 42)
    v3["hello"] = ((v36 << ((1) & 63)) == v36)
    for fi38 := 0; fi38 < 19; fi38++ {
        fmt.Println(v37, v1)
        v39 := GenS3{F0: 1, F1: 0, F2: -32768}
        var v40 bool = (0 != 40)
        _ = fi38
        _ = v39
        _ = v40
    }
    fmt.Println(v37)
    v12 = -195
    v41 := GenS4{F0: 618}
    var v42 int8 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v18
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v41
    _ = v42
    _ = v5
}
