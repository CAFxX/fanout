package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int32; F1 uint64 }
func genFn0(p0 string, p1 int, p2 int8) uint8 {
    switch -123 {
    case 42:
        p2 -= p2
    }
    var v4 int64 = 0
    v5 := []int{128, -242, p1}
    _ = len(v5)
    v6 := GenS1{F0: v4, F1: 2}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 10
}
func genFn1(p0 uint8) int8 {
    var v7 string = (("a" + ("foo bar" + ("\n\t" + "\n\t"))) + "foo bar")
    switch 32768 {
    case 127:
        var v8 uint16 = 48
        _ = v8
    case 9223372036854775807:
        switch ("a" + v7) {
        case "":
            if false {
                if true {
                    v9 := []int{1, 722, 127, 0}
                    _ = v9
                }
                for fi10 := 0; fi10 < 9; fi10++ {
                    var v11 int64 = 0
                    _ = fi10
                    _ = v11
                }
                var v12 uint64 = uint64(p0)
                _ = v12
            } else {
                v13 := GenS2{F0: (p0 / ((p0) | 1))}
                _ = v13
            }
        }
    default:
        for fi14 := 0; fi14 < 3; fi14++ {
            v15 := []int{}
            _ = fi14
            _ = v15
        }
    }
    var v16 int8 = 42
    v17 := GenS2{F0: (p0 - 255)}
    _ = p0
    _ = v16
    _ = v17
    _ = v7
    return (^(v16))
}

func main() {
    var v1 int64 = -565
    var v2 int64 = (^(v1))
    v3 := make(map[int]string)
    v18 := []int{-2147483648}
    var v19 float32 = float32((1.5 / 1e10))
    v3[9223372036854775807] = "\n\t"
    var v20 int16 = -1
    var v21 int8 = 100
    if (((-819 == 2147483647) || true) || true) {
        for fi22 := 0; fi22 < 16; fi22++ {
            var v23 uint32 = 0
            _ = fi22
            _ = v23
        }
    } else {
        for fi24 := 0; fi24 < 9; fi24++ {
            v19 -= 0.1
            v20 = v20
            _ = fi24
        }
    }
    var v25 uint = 594
    v20 = v20
    var v26 uint64 = 0
    v1 -= (^(v1))
    var v27 int32 = 779
    if true {
        v28 := GenS2{F0: 47}
        _ = v28
    }
    if (v2 <= v1) {
        v21 += 7
    }
    if true {
        var v29 int16 = 1
        fmt.Println(v27, v20)
        _ = v29
    }
    var v30 uint8 = 127
    for fi31 := 0; fi31 < 8; fi31++ {
        fmt.Println(v26, v21, v2)
        switch 533 {
        case 127:
            var v32 int16 = (v20 % ((v20) | 1))
            _ = v32
        case 2147483647:
            for fi33 := 0; fi33 < 4; fi33++ {
                v18 = append(v18, fi33)
                _ = fi33
            }
        case 1:
            switch fi31 {
            case 0:
                if (0.1 != 1e10) {
                    v34 := GenS3{F0: 1000}
                    v35 := GenS4{F0: genFn1((v30 | genFn0("a", fi31, (v21 + 83))))}
                    _ = v34
                    _ = v35
                }
            case 1:
                v36 := GenS5{F0: (v27 | v27), F1: 216}
                _ = v36
            case 128:
                v26 = v26
            }
        default:
            if len(v18) > 0 {
                v18[0] = 300
            }
        }
        _ = fi31
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v30
}
