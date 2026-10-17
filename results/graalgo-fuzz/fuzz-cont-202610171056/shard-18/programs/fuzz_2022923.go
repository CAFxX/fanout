package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 bool }
type GenS2 struct { F0 int8; F1 uint8; F2 uint8 }
type GenS3 struct { F0 int16; F1 string }
func genFn0(p0 uint32) uint32 {
    v6 := []int{}
    v7 := GenS1{F0: -649, F1: 1, F2: (!(!(0.0 == 1e10)))}
    _ = p0
    _ = v6
    _ = v7
    return 2147483647
}
func genFn1(p0 int64, p1 int8) uint32 {
    for fi8 := 0; fi8 < 12; fi8++ {
        for fi9 := 0; fi9 < 4; fi9++ {
            var v10 uint = 1
            v10 = (^(v10))
            if ((3.14159 / 1e10) < 0.0) {
                v11 := []int{(fi8 % ((42) | 1)), 0, fi8}
                p1 = p1
                _ = v11
            }
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        if true {
            p0 += (p0 / ((p0) | 1))
            if true {
                p0 -= p0
                p1 = p1
            } else {
                v13 := GenS1{F0: fi12, F1: 4294967295, F2: (2147483647 <= -1)}
                for fi14 := 0; fi14 < 10; fi14++ {
                    if (!true) {
                        v15 := GenS1{F0: (fi14 | 127), F1: 0, F2: ((^(p0)) < 42)}
                        if false {
                            switch "hello" {
                            case "":
                                p1 += 127
                            case "foo bar":
                                p0 = (^(p0))
                            }
                            var v16 int8 = (-(p1))
                            _ = v16
                        }
                        var v17 int8 = 42
                        _ = v15
                        _ = v17
                    }
                    fmt.Println(fi12, p1)
                    _ = fi14
                }
                _ = v13
            }
        } else {
            switch "" {
            case "hello":
                p1 = (^(p1))
            case "a":
                fmt.Println(fi12)
            }
            var v18 uint16 = 32767
            switch ("hello" + "a") {
            case "hello":
                v19 := []int{42, 4294967295, (-(fi12))}
                _ = v19
            case "":
                var v20 string = "a"
                _ = v20
            case "foo bar":
                p0 += p0
            default:
                if true {
                    var v21 int64 = (p0 + p0)
                    _ = v21
                } else {
                    var v22 uint16 = (^(v18))
                    fmt.Println(p0, v18)
                    if ((p1 << ((324) & 63)) <= -1) {
                        var v23 string = "hello"
                        v24 := []int{2147483647, fi12, (fi12 & fi12)}
                        _ = v23
                        _ = v24
                    }
                    _ = v22
                }
            }
            _ = v18
        }
        for fi25 := 0; fi25 < 3; fi25++ {
            var v26 int8 = p1
            var v27 bool = true
            _ = fi25
            _ = v26
            _ = v27
        }
        _ = fi12
    }
    p0 -= p0
    _ = p0
    _ = p1
    return 4294967295
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint = 848
    var v3 float64 = 0.1
    v4 := make(map[uint32]string)
    var v5 float32 = float32(v3)
    v1 = 1
    v3 = 0.0
    v3 -= v3
    v28 := GenS2{F0: 107, F1: 161, F2: 255}
    switch "hello" {
    case "":
        v1 -= (v1 >> ((v2) & 63))
    }
    if false {
        switch "hello" {
        case "":
            var v29 float64 = v3
            _ = v29
        default:
            v3 = 0.0
        }
    }
    fmt.Println(v5, v1)
    switch ("a" + "\n\t") {
    case "a":
        for fi30 := 0; fi30 < 16; fi30++ {
            fmt.Println(v2, v5)
            v5 -= v5
            var v31 string = ("\n\t" + "a")
            _ = fi30
            _ = v31
        }
    case "hello":
        v32 := GenS1{F0: 617, F1: 0, F2: true}
        _ = v32
    case "foo bar":
        var v33 int = -581
        _ = v33
    default:
        var v34 int32 = 950
        _ = v34
    }
    var v35 int8 = 42
    v36 := GenS3{F0: -32768, F1: "\n\t"}
    v1 = (v1 * (v1 >> (((^(v2))) & 63)))
    v3 = (v3 * 1.5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
