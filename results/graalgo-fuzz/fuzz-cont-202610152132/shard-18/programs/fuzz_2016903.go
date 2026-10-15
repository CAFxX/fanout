package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 string; F2 int64 }
type GenS3 struct { F0 int16; F1 bool; F2 string }
type GenS4 struct { F0 uint64; F1 uint }
func genFn0(p0 uint, p1 string) int16 {
    var v7 string = ""
    var v8 uint32 = 2147483647
    v8 = v8
    if true {
        v9 := GenS1{F0: 42}
        switch p1 {
        case "":
            v10 := GenS1{F0: 2}
            _ = v10
        default:
            for fi11 := 0; fi11 < 13; fi11++ {
                if false {
                    v12 := GenS1{F0: 157}
                    v13 := []int{fi11, 292, (fi11 << (((p0 + p0)) & 63)), fi11}
                    _ = v12
                    _ = v13
                } else {
                    if false {
                        var v14 uint = 18446744073709551615
                        _ = v14
                    }
                    p1 = p1
                }
                switch "foo bar" {
                case "hello":
                    var v15 uint16 = 42
                    _ = v15
                case "":
                    p1 = p1
                case "\n\t":
                    for fi16 := 0; fi16 < 15; fi16++ {
                        v17 := GenS1{F0: 2}
                        _ = fi16
                        _ = v17
                    }
                }
                switch "a" {
                case "foo bar":
                    for fi18 := 0; fi18 < 4; fi18++ {
                        switch (fi18 / ((2147483647) | 1)) {
                        case 2147483647:
                            if false {
                                v7 = ((("" + p1) + p1) + p1)
                            } else {
                                var v19 uint = (p0 / (((p0 - (p0 >> (((p0 * 84)) & 63)))) | 1))
                                v20 := GenS2{F0: -1, F1: (((v7 + (((p1 + ((v7 + p1) + "foo bar")) + v7) + ("\n\t" + ("hello" + v7)))) + ("" + ("a" + p1))) + (v7 + (p1 + (p1 + v7)))), F2: -495}
                                _ = v19
                                _ = v20
                            }
                        case 32767:
                            v21 := []int{fi18}
                            _ = v21
                        }
                        v22 := []int{fi11}
                        if len(v22) >= 2 {
                            v22 = v22[1:len(v22)]
                        }
                        _ = fi18
                        _ = v22
                    }
                case "a":
                    v23 := GenS3{F0: 1000, F1: true, F2: p1}
                    _ = v23
                default:
                    p0 += p0
                }
                _ = fi11
            }
        }
        _ = v9
    }
    switch (p1 + "\n\t") {
    case "":
        v8 = (^(v8))
    case "hello":
        fmt.Println(p0, v8)
    default:
        var v24 float32 = float32(-2.25)
        _ = v24
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 968
}

func main() {
    var v1 bool = false
    var v2 float64 = -2.25
    v3 := make(map[uint32]uint32)
    var v4 uint64 = 2
    v5 := make(map[uint32]bool)
    v6 := make(map[int]uint32)
    var v25 float64 = v2
    if false {
        switch -157 {
        case 669:
            var v26 float64 = 1e10
            _ = v26
        default:
            v27 := GenS1{F0: 147}
            _ = v27
        }
        for fi28 := 0; fi28 < 3; fi28++ {
            var v29 uint16 = 2
            _ = fi28
            _ = v29
        }
    } else {
        if v1 {
            if v1 {
                v2 = 0.0
            }
        } else {
            v30 := GenS2{F0: -873, F1: "a", F2: -844}
            _ = v30
        }
        v31 := v6[311]
        _ = v31
    }
    var v32 uint32 = 781
    v33 := []int{0, 4294967295, 127}
    v1 = v1
    v2 = v25
    v2 += v2
    v3[(v32 * 2)] = v32
    if false {
        for fi34 := 0; fi34 < 10; fi34++ {
            v2 = (v2 * v2)
            var v35 string = "\n\t"
            delete(v6, 4294967295)
            _ = fi34
            _ = v35
        }
        var v36 float64 = v2
        var v37 int = 742
        _ = v36
        _ = v37
    } else {
        v1 = (v1 || ((v1 || (528 > -1)) || v1))
    }
    for fi38 := 0; fi38 < 14; fi38++ {
        v32 = 42
        if len(v33) > 0 {
            v33[0] = fi38
        }
        _ = fi38
    }
    v39 := GenS4{F0: v4, F1: 1}
    var v40 float64 = (-2.25 * v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v32
    _ = v33
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
