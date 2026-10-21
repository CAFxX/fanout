package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 uint64; F1 int32; F2 bool }
type GenS4 struct { F0 int16; F1 uint32; F2 string }
func genFn0(p0 int32) string {
    v5 := []int{10, -166, -225}
    _ = len(v5)
    _ = p0
    _ = v5
    return ""
}
func genFn1() uint16 {
    for fi6 := 0; fi6 < 7; fi6++ {
        switch "foo bar" {
        case "":
            v7 := []int{fi6, fi6, (-(fi6))}
            _ = v7
        case "hello":
            var v8 string = "\n\t"
            _ = v8
        case "\n\t":
            v9 := GenS1{F0: -9223372036854775808}
            _ = v9
        }
        fmt.Println(fi6)
        v10 := []int{}
        _ = fi6
        _ = v10
    }
    var v11 string = ("hello" + "foo bar")
    for fi12 := 0; fi12 < 14; fi12++ {
        switch v11 {
        case "hello":
            var v13 float64 = 1.5
            _ = v13
        case "":
            for fi14 := 0; fi14 < 11; fi14++ {
                var v15 uint = 2
                var v16 int64 = 42
                v11 = ((v11 + v11) + ("hello" + v11))
                _ = fi14
                _ = v15
                _ = v16
            }
        default:
            var v17 bool = (-2147483648 == 0)
            _ = v17
        }
        var v18 uint32 = 398
        var v19 float64 = 1.5
        _ = fi12
        _ = v18
        _ = v19
    }
    _ = v11
    return 414
}

func main() {
    v1 := make(map[int]int)
    var v2 int8 = 100
    v3 := []int{-524, -1}
    var v4 float32 = (float32(3.14159) * float32(1e10))
    for fi20 := 0; fi20 < 8; fi20++ {
        v2 = -1
        if false {
            v21 := GenS2{F0: "\n\t", F1: "\n\t"}
            for fi22 := 0; fi22 < 14; fi22++ {
                if false {
                    fmt.Println(fi20, v4, fi22)
                    var v23 uint8 = 128
                    if (870 != 42) {
                        fmt.Println(fi20)
                    } else {
                        fmt.Println(fi22, fi20, fi22)
                        var v24 float32 = v4
                        _ = v24
                    }
                    _ = v23
                }
                _ = fi22
            }
            if (((((3.14159 + v4) * 0.0) - 0.0) * v4) >= 0.1) {
                if (v2 == v2) {
                    fmt.Println(v2)
                    v2 += v2
                } else {
                    if false {
                        if false {
                            if (-18 < 0) {
                                fmt.Println(v2, v2)
                            } else {
                                for fi25 := 0; fi25 < 18; fi25++ {
                                    v26 := v1[(fi20 & (fi25 & fi20))]
                                    v26 -= (fi25 << ((42) & 63))
                                    _ = fi25
                                    _ = v26
                                }
                                _ = len(v1)
                            }
                            fmt.Println(v4)
                            v27 := GenS1{F0: fi20}
                            _ = v27
                        }
                    }
                }
            } else {
                v2 -= 42
            }
            _ = v21
        }
        switch ("\n\t" + "hello") {
        case "foo bar":
            fmt.Println(fi20)
        case "hello":
            v4 += v4
        }
        _ = fi20
    }
    var v28 int = -874
    v29 := GenS3{F0: 0, F1: 594, F2: false}
    switch "\n\t" {
    case "":
        v30 := v1[-984]
        _ = v30
    case "hello":
        v2 = v2
    case "foo bar":
        switch "hello" {
        case "hello":
            if (4294967295 < 2) {
                v4 += -2.25
                fmt.Println(v4, v4, v4)
                if len(v3) > 0 {
                    v3[0] = (v28 >> ((42) & 63))
                }
            }
        case "\n\t":
            v31 := GenS4{F0: -1, F1: 2, F2: ("\n\t" + "foo bar")}
            _ = v31
        case "":
            for fi32 := 0; fi32 < 13; fi32++ {
                _ = len(v3)
                _ = len(v1)
                _ = fi32
            }
        }
    }
    for fi33 := 0; fi33 < 16; fi33++ {
        for fi34 := 0; fi34 < 2; fi34++ {
            if (670 > 0) {
                fmt.Println(v4, fi33, fi33)
                _ = len(v3)
                var v35 int8 = (v2 % (((^(v2))) | 1))
                _ = v35
            } else {
                fmt.Println(v28, v2, v4)
            }
            _ = fi34
        }
        _ = fi33
    }
    fmt.Println(v4)
    v28 -= (v28 / ((v28) | 1))
    var v36 bool = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v36
    _ = v4
}
