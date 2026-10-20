package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 uint8; F1 uint; F2 int32 }
func genFn0(p0 int16, p1 int8, p2 int64) int64 {
    var v4 int16 = (p0 & (p0 * (^(p0))))
    for fi5 := 0; fi5 < 6; fi5++ {
        if (!(!true)) {
            if false {
                v6 := GenS1{F0: 673}
                _ = v6
            }
        } else {
            switch "\n\t" {
            case "a":
                fmt.Println(fi5, fi5)
            case "foo bar":
                fmt.Println(v4, p1, p1)
            case "\n\t":
                var v7 uint32 = 42
                _ = v7
            default:
                v8 := []int{(fi5 / (((fi5 & fi5)) | 1)), 4294967295, fi5, (-(fi5))}
                _ = v8
            }
        }
        p1 -= p1
        var v9 uint8 = 2
        _ = fi5
        _ = v9
    }
    switch "a" {
    case "":
        var v10 uint64 = 505
        _ = v10
    case "\n\t":
        var v11 uint8 = 251
        _ = v11
    case "hello":
        fmt.Println(p0, p2)
    }
    p1 -= p1
    for fi12 := 0; fi12 < 18; fi12++ {
        p1 += (p1 >> ((749) & 63))
        var v13 int16 = (p0 / ((v4) | 1))
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return p2
}
func genFn1(p0 int) uint16 {
    for fi14 := 0; fi14 < 7; fi14++ {
        v15 := []int{-1, 545, p0}
        p0 -= p0
        _ = fi14
        _ = v15
    }
    switch ("a" + "\n\t") {
    case "a":
        p0 = 128
    case "":
        var v16 int64 = 0
        _ = v16
    default:
        v17 := GenS1{F0: 1}
        _ = v17
    }
    switch 0 {
    case 32768:
        switch "" {
        case "hello":
            var v18 uint8 = 0
            _ = v18
        default:
            v19 := GenS2{F0: "", F1: (p0 + p0)}
            _ = v19
        }
    case 9223372036854775807:
        if (127 < 42) {
            for fi20 := 0; fi20 < 15; fi20++ {
                if true {
                    var v21 int16 = 1000
                    _ = v21
                }
                if false {
                    for fi22 := 0; fi22 < 14; fi22++ {
                        v23 := GenS2{F0: "foo bar", F1: p0}
                        v24 := []int{(fi20 | (fi20 * (fi22 | 32768))), (-(p0)), fi20}
                        var v25 uint32 = 2
                        _ = fi22
                        _ = v23
                        _ = v24
                        _ = v25
                    }
                } else {
                    if false {
                        fmt.Println(p0, fi20)
                    } else {
                        var v26 int16 = -1
                        _ = v26
                    }
                    v27 := []int{(fi20 >> ((364) & 63))}
                    if true {
                        p0 += (fi20 >> ((1) & 63))
                        fmt.Println(fi20, p0)
                        var v28 float32 = -2.25
                        _ = v28
                    }
                    _ = v27
                }
                var v29 uint32 = 405
                _ = fi20
                _ = v29
            }
            v30 := GenS3{F0: 127, F1: 627, F2: -2147483648}
            var v31 int = 0
            _ = v30
            _ = v31
        } else {
            v32 := []int{0, p0, (p0 >> ((4294967295) & 63))}
            _ = v32
        }
    default:
        fmt.Println(p0)
    }
    var v33 string = "hello"
    _ = p0
    _ = v33
    return 1
}

func main() {
    v1 := []uint32{}
    var v2 bool = (0 < -2147483648)
    v3 := make(map[string]int)
    v3["hello"] = 558
    v2 = ((1 <= 4294967295) || (283 >= 807))
    v2 = (genFn1(9223372036854775807) < 65535)
    var v34 int64 = -606
    var v35 int = -509
    var v36 float32 = 3.14159
    v34 = v34
    v37 := GenS1{F0: v34}
    var v38 string = "a"
    v38 = v38
    v34 += genFn0(32767, 34, v34)
    delete(v3, v38)
    v38 = "hello"
    var v39 string = "a"
    if true {
        v2 = true
        v2 = (!true)
        v38 = v39
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
}
