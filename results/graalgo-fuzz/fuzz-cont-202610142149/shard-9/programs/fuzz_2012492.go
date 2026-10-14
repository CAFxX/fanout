package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 int32; F1 int16 }
type GenS5 struct { F0 int32; F1 uint16 }
func genFn0(p0 uint64, p1 int8, p2 uint) uint64 {
    for fi6 := 0; fi6 < 19; fi6++ {
        for fi7 := 0; fi7 < 6; fi7++ {
            var v8 uint32 = 165
            v9 := []int{int(p0), (^(fi6))}
            v8 = (v8 % (((^(v8))) | 1))
            _ = fi7
            _ = v8
            _ = v9
        }
        v10 := GenS1{F0: 2147483648, F1: uint8(p0)}
        var v11 float64 = 0.0
        _ = fi6
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 17; fi12++ {
        p2 = (p2 & 977)
        _ = fi12
    }
    switch 319 {
    case 127:
        if (1.5 >= 3.14159) {
            for fi13 := 0; fi13 < 3; fi13++ {
                p0 -= uint64(p2)
                v14 := []int{(fi13 >> ((p2) & 63)), (fi13 >> ((p2) & 63)), (fi13 * int(p2))}
                _ = fi13
                _ = v14
            }
            var v15 uint64 = (^(p0))
            _ = v15
        } else {
            if (true || false) {
                v16 := GenS1{F0: 42, F1: 127}
                var v17 string = ((("\n\t" + "hello") + ("hello" + "a")) + ("\n\t" + "\n\t"))
                _ = v16
                _ = v17
            } else {
                v18 := GenS2{F0: 957, F1: 2147483647}
                v19 := []int{-966, -373}
                if len(v19) > 0 {
                    v19[0] = 32767
                }
                _ = v18
                _ = v19
            }
            switch int(p2) {
            case 32768:
                if false {
                    v20 := []int{}
                    p1 = (p1 | p1)
                    _ = v20
                }
            case -267:
                p1 = p1
            default:
                var v21 uint32 = 16
                _ = v21
            }
            if (-265 != -1) {
                var v22 bool = (p1 > (p1 + p1))
                v23 := []int{}
                _ = v22
                _ = v23
            }
        }
    case 32767:
        var v24 uint = 890
        _ = v24
    case 42:
        v25 := []int{2147483647}
        _ = v25
    }
    var v26 uint16 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    return (p0 % (((p0 | (p0 % ((42) | 1)))) | 1))
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]int)
    v3 := []int{-122}
    var v4 uint16 = 32767
    var v5 uint64 = 0
    v27 := GenS3{F0: "", F1: 12}
    _ = len(v1)
    var v28 uint32 = 252
    for fi29 := 0; fi29 < 7; fi29++ {
        var v30 int16 = 0
        for fi31 := 0; fi31 < 7; fi31++ {
            switch ("\n\t" + "a") {
            case "foo bar":
                fmt.Println(v30, v4)
            case "hello":
                v4 = 32767
            default:
                var v32 int16 = v30
                _ = v32
            }
            _ = fi31
        }
        for fi33 := 0; fi33 < 9; fi33++ {
            var v34 uint8 = 226
            v35 := GenS4{F0: 42, F1: 1}
            v36 := GenS1{F0: 2147483647, F1: (v34 % ((v34) | 1))}
            var v37 float32 = 1e10
            _ = fi33
            _ = v34
            _ = v35
            _ = v36
            _ = v37
        }
        _ = fi29
        _ = v30
    }
    v5 = (v5 - 2)
    v38 := GenS3{F0: "hello", F1: 255}
    fmt.Println(v5, v5, v28)
    var v39 int16 = -32768
    v40 := GenS2{F0: (v39 * v39), F1: 0}
    v4 -= (v4 - v4)
    switch "\n\t" {
    case "hello":
        fmt.Println(v5)
    default:
        fmt.Println(v4)
    }
    var v41 float64 = 0.1
    v39 = -774
    switch 42 {
    case -9223372036854775808:
        var v42 string = "foo bar"
        _ = v42
    }
    var v43 uint64 = v5
    v44 := GenS5{F0: -1, F1: 255}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v43
    _ = v44
    _ = v5
}
