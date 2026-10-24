package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 string; F2 int16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint64; F1 int32; F2 uint32 }
func genFn0(p0 uint64, p1 bool, p2 int16) int32 {
    p2 -= (^(p2))
    var v4 float32 = (float32(3.14159) / (-2.25 * 3.14159))
    var v5 bool = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 945
}
func genFn1() string {
    if true {
        var v6 float32 = 1.5
        var v7 float32 = (1e10 * (v6 * v6))
        if true {
            var v8 uint = 42
            _ = v8
        } else {
            for fi9 := 0; fi9 < 3; fi9++ {
                v7 = (0.0 / (1.5 - (v7 / v6)))
                _ = fi9
            }
            v10 := []int{2147483647, -1, 32768, 1}
            _ = v10
        }
        _ = v6
        _ = v7
    }
    var v11 int8 = 1
    var v12 uint64 = 42
    v13 := GenS1{F0: false}
    _ = v11
    _ = v12
    _ = v13
    return "a"
}

func main() {
    var v1 int16 = -328
    var v2 int16 = (v1 % ((42) | 1))
    var v3 float64 = 0.1
    v14 := GenS1{F0: true}
    if ((v1 - (v2 - v1)) < v1) {
        var v15 uint = 2
        _ = v15
    } else {
        v16 := GenS2{F0: 1}
        _ = v16
    }
    switch (("\n\t" + "\n\t") + ("a" + ("foo bar" + "foo bar"))) {
    case "foo bar":
        for fi17 := 0; fi17 < 14; fi17++ {
            fmt.Println(v2, v3)
            _ = fi17
        }
    case "a":
        v18 := GenS2{F0: 43}
        _ = v18
    default:
        switch "hello" {
        case "hello":
            for fi19 := 0; fi19 < 11; fi19++ {
                var v20 int64 = -29
                _ = fi19
                _ = v20
            }
        }
    }
    var v21 uint8 = 208
    var v22 float64 = (v3 / (v3 / (v3 * 3.14159)))
    v23 := []int{-1, -871, -310}
    v24 := GenS1{F0: true}
    var v25 uint = 632
    v25 += v25
    v26 := GenS3{F0: uint32(v25), F1: "", F2: (v1 % ((v2) | 1))}
    v2 = -1
    if (!true) {
        fmt.Println(v3)
        v3 += v3
        for fi27 := 0; fi27 < 10; fi27++ {
            fmt.Println(v1, v22, fi27)
            var v28 int = fi27
            if false {
                fmt.Println(v25, v25, v22)
                for fi29 := 0; fi29 < 3; fi29++ {
                    if false {
                        var v30 float32 = (float32(v22) * 3.14159)
                        if true {
                            v31 := GenS4{F0: v21}
                            var v32 bool = true
                            var v33 float64 = (float64(0.0) + (float64((v30 * 0.0)) + v3))
                            v33 -= (float64((v30 + 3.14159)) * v22)
                            var v34 int = (fi29 | 311)
                            _ = v31
                            _ = v32
                            _ = v33
                            _ = v34
                        } else {
                            var v35 uint64 = 547
                            var v36 uint8 = v21
                            if len(v23) > 0 {
                                v23[0] = (fi29 / ((766) | 1))
                            }
                            _ = v35
                            _ = v36
                        }
                        _ = v30
                    } else {
                        v2 += (v2 | v1)
                    }
                    _ = fi29
                }
                for fi37 := 0; fi37 < 15; fi37++ {
                    v38 := GenS5{F0: 350, F1: 519, F2: 37}
                    var v39 uint8 = 255
                    if true {
                        v21 = 1
                        v3 += 0.0
                        var v40 string = ("a" + "foo bar")
                        _ = v40
                    } else {
                        var v41 uint16 = 255
                        fmt.Println(v21)
                        var v42 uint16 = v41
                        _ = v41
                        _ = v42
                    }
                    v43 := GenS1{F0: true}
                    fmt.Println(v39, v2)
                    _ = fi37
                    _ = v38
                    _ = v39
                    _ = v43
                }
            }
            _ = fi27
            _ = v28
        }
    }
    for fi44 := 0; fi44 < 12; fi44++ {
        var v45 string = ""
        _ = fi44
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}
