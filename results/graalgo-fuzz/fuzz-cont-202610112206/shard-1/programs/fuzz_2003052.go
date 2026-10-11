package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 uint64; F2 int8 }
type GenS3 struct { F0 int }
func genFn0() uint64 {
    var v4 float32 = 0.1
    var v5 uint16 = 42
    var v6 int32 = -161
    fmt.Println(v4, v6, v6)
    _ = v4
    _ = v5
    _ = v6
    return 965
}
func genFn1(p0 string, p1 int64) int8 {
    fmt.Println(p0)
    if (!(!(460 != 4294967295))) {
        v7 := []int{-1, -456}
        v7 = append(v7, -750)
        _ = v7
    }
    if false {
        v8 := []int{523}
        var v9 int32 = -933
        _ = v8
        _ = v9
    } else {
        var v10 int8 = 42
        _ = v10
    }
    v11 := GenS1{F0: 0}
    _ = p0
    _ = p1
    _ = v11
    return -1
}

func main() {
    v1 := []uint32{2147483648}
    v2 := make(map[string]string)
    var v3 uint32 = 2
    v3 += v3
    for fi12 := 0; fi12 < 16; fi12++ {
        fmt.Println(fi12, fi12)
        for fi13 := 0; fi13 < 14; fi13++ {
            v3 = (^(v3))
            _ = fi13
        }
        _ = fi12
    }
    var v14 int = 828
    if (-2.25 >= 0.0) {
        var v15 float64 = 1.5
        _ = v15
    } else {
        fmt.Println(v14, v14)
        if (true || false) {
            for fi16 := 0; fi16 < 20; fi16++ {
                if true {
                    v14 -= (v14 ^ -508)
                } else {
                    switch 405 {
                    case 2147483647:
                        var v17 uint64 = genFn0()
                        _ = v17
                    case 1:
                        var v18 int = (v14 ^ fi16)
                        _ = v18
                    default:
                        var v19 string = (("\n\t" + "foo bar") + "hello")
                        _ = v19
                    }
                }
                if true {
                    for fi20 := 0; fi20 < 20; fi20++ {
                        var v21 int32 = -457
                        var v22 int16 = -423
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                    for fi23 := 0; fi23 < 10; fi23++ {
                        switch fi16 {
                        case 432:
                            switch v14 {
                            case -705:
                                var v24 int32 = 42
                                v25 := GenS2{F0: 9223372036854775807, F1: 0, F2: 42}
                                _ = v24
                                _ = v25
                            }
                        case -602:
                            v3 = 211
                        case 0:
                            v3 = v3
                        default:
                            v14 = (fi16 / ((42) | 1))
                        }
                        var v26 bool = false
                        var v27 uint64 = 31
                        _ = fi23
                        _ = v26
                        _ = v27
                    }
                } else {
                    switch 128 {
                    case 127:
                        for fi28 := 0; fi28 < 10; fi28++ {
                            v3 = (v3 / (((^(v3))) | 1))
                            _ = fi28
                        }
                    case 103:
                        delete(v2, "foo bar")
                    case 1:
                        v3 += (^(v3))
                    }
                    var v29 uint8 = 1
                    _ = v29
                }
                v30 := v2[""]
                _ = fi16
                _ = v30
            }
        } else {
            switch "a" {
            case "hello":
                var v31 uint32 = v3
                _ = v31
            case "a":
                switch "a" {
                case "foo bar":
                    switch "foo bar" {
                    case "\n\t":
                        delete(v2, "")
                    case "":
                        v32 := GenS2{F0: -204, F1: 1, F2: 42}
                        _ = v32
                    default:
                        fmt.Println(v14, v14)
                    }
                }
            case "\n\t":
                v14 += v14
            }
            v33 := GenS1{F0: 133}
            var v34 string = ""
            _ = v33
            _ = v34
        }
        var v35 int16 = 42
        _ = v35
    }
    if true {
        delete(v2, "hello")
        if true {
            var v36 uint64 = genFn0()
            v1 = append(v1, (v3 * v3))
            _ = v36
        } else {
            v3 = (v3 - v3)
        }
        var v37 int32 = 42
        _ = v37
    } else {
        fmt.Println(v14, v14)
        for fi38 := 0; fi38 < 15; fi38++ {
            v14 = (^(v14))
            _ = fi38
        }
        v39 := v2["\n\t"]
        _ = v39
    }
    fmt.Println(v3, v3)
    var v40 float64 = 3.14159
    var v41 bool = (42 <= 565)
    v42 := GenS3{F0: (v14 + v14)}
    v3 += v3
    for fi43 := 0; fi43 < 7; fi43++ {
        var v44 bool = (true || (!(!(genFn0() != 42))))
        _ = fi43
        _ = v44
    }
    var v45 float32 = 0.0
    for fi46 := 0; fi46 < 14; fi46++ {
        v14 -= (fi46 / ((fi46) | 1))
        _ = fi46
    }
    v14 -= (^(v14))
    fmt.Println(v40)
    v14 = (v14 ^ v14)
    v40 = v40
    switch ("\n\t" + "\n\t") {
    case "a":
        var v47 uint16 = 271
        _ = v47
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v40
    _ = v41
    _ = v42
    _ = v45
}
