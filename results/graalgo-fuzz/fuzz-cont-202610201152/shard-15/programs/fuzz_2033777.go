package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 uint16; F2 uint }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint64; F1 int64 }
type GenS6 struct { F0 uint; F1 int8; F2 uint64 }
type GenS7 struct { F0 int32 }
func genFn0(p0 uint8) uint16 {
    if (true || ("" != "")) {
        for fi6 := 0; fi6 < 17; fi6++ {
            fmt.Println(p0)
            for fi7 := 0; fi7 < 16; fi7++ {
                var v8 int64 = 0
                if false {
                    v9 := GenS1{F0: 1}
                    v10 := []int{fi7}
                    var v11 int32 = 2147483647
                    _ = v10
                    _ = v11
                    _ = v9
                } else {
                    v12 := []int{(fi7 + 0)}
                    p0 = 2
                    v13 := GenS1{F0: 768}
                    fmt.Println(v8, v8, fi7)
                    var v14 string = ("" + "foo bar")
                    _ = v12
                    _ = v13
                    _ = v14
                }
                var v15 float32 = (((3.14159 * 0.0) - 1.5) + 0.1)
                _ = fi7
                _ = v15
                _ = v8
            }
            var v16 uint16 = 32767
            _ = fi6
            _ = v16
        }
        p0 -= p0
    } else {
        fmt.Println(p0)
        var v17 uint = 4294967295
        switch 408 {
        case 2147483647:
            p0 = 0
        case 128:
            for fi18 := 0; fi18 < 7; fi18++ {
                switch ("hello" + "a") {
                case "foo bar":
                    p0 = uint8(v17)
                default:
                    if true {
                        var v19 float32 = ((1e10 - 1e10) - 0.1)
                        if (9223372036854775807 < 42) {
                            v20 := GenS2{F0: v17, F1: 0, F2: (v17 - 42)}
                            v21 := GenS3{F0: "\n\t", F1: true}
                            for fi22 := 0; fi22 < 16; fi22++ {
                                var v23 bool = (-1 == -1)
                                _ = fi22
                                _ = v23
                            }
                            _ = v20
                            _ = v21
                        }
                        _ = v19
                    } else {
                        p0 = p0
                        v24 := GenS4{F0: "foo bar"}
                        _ = v24
                    }
                }
                var v25 int32 = 0
                switch 127 {
                case 128:
                    p0 += 2
                case 32768:
                    fmt.Println(fi18, fi18)
                }
                _ = fi18
                _ = v25
            }
        case 32767:
            p0 = p0
        }
        _ = v17
    }
    v26 := GenS5{F0: 18446744073709551615, F1: 9223372036854775807}
    v27 := GenS6{F0: 255, F1: 100, F2: 419}
    _ = p0
    _ = v26
    _ = v27
    return 1
}
func genFn1(p0 uint, p1 int32, p2 uint8) int {
    if ((1 > 65535) && true) {
        var v28 string = "\n\t"
        _ = v28
    } else {
        if true {
            v29 := []int{-736, -9223372036854775808, 783}
            _ = len(v29)
            _ = v29
        } else {
            v30 := []int{1, -9223372036854775808, 257, 447}
            fmt.Println(p1)
            var v31 int32 = 1
            _ = v30
            _ = v31
        }
        if true {
            fmt.Println(p0)
            v32 := []int{}
            fmt.Println(p2)
            _ = v32
        } else {
            switch "\n\t" {
            case "hello":
                p2 -= 1
            case "a":
                var v33 bool = (20 == uint64(p0))
                _ = v33
            case "foo bar":
                p0 -= p0
            }
            fmt.Println(p1)
            v34 := GenS1{F0: 0}
            _ = v34
        }
    }
    v35 := []int{}
    v36 := GenS7{F0: (-(p1))}
    _ = p0
    _ = p1
    _ = p2
    _ = v35
    _ = v36
    return 1
}

func main() {
    var v1 int16 = -32768
    v2 := make(map[int]bool)
    var v3 bool = true
    var v4 int8 = -60
    v5 := make(map[uint32]uint32)
    _ = len(v5)
    var v37 int = 128
    for fi38 := 0; fi38 < 10; fi38++ {
        v1 = (v1 >> ((2) & 63))
        _ = fi38
    }
    var v39 uint = 2
    var v40 bool = v3
    v41 := []int{(v37 ^ genFn1(489, 1, 0)), int(v1), v37}
    switch "hello" {
    case "a":
        v39 = 2
    case "hello":
        fmt.Println(v37)
    }
    fmt.Println(v4, v4)
    var v42 float64 = float64(1.5)
    fmt.Println(v3)
    var v43 int16 = 131
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
}
