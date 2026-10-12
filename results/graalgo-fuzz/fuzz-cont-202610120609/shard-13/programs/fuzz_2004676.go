package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint8, p1 uint) int16 {
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := GenS1{F0: 0, F1: (fi7 * (fi7 & 871))}
        _ = fi7
        _ = v8
    }
    fmt.Println(p0, p1)
    fmt.Println(p0)
    for fi9 := 0; fi9 < 2; fi9++ {
        v10 := []int{(fi9 ^ 42), (fi9 * 157)}
        var v11 int16 = 32767
        p1 = 0
        _ = fi9
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 2; fi12++ {
        for fi13 := 0; fi13 < 12; fi13++ {
            v14 := GenS2{F0: 9223372036854775807}
            switch "a" {
            case "\n\t":
                var v15 float64 = 0.0
                _ = v15
            case "hello":
                var v16 int = (fi13 + fi12)
                fmt.Println(fi12, fi13)
                _ = v16
            case "":
                v17 := GenS1{F0: 255, F1: fi13}
                _ = v17
            default:
                fmt.Println(p1, p0, p1)
            }
            var v18 string = ("a" + ("a" + "\n\t"))
            _ = fi13
            _ = v14
            _ = v18
        }
        v19 := []int{(fi12 * 127)}
        _ = fi12
        _ = v19
    }
    _ = p0
    _ = p1
    return 892
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{2, 0}
    var v3 int16 = -1
    v4 := []string{"a", "foo bar", ""}
    v5 := make(map[int]string)
    v6 := make(map[uint32]int)
    if len(v2) > 0 {
        v2[0] = 2
    }
    if (42 == 0) {
        v20 := GenS1{F0: 2, F1: 32768}
        _ = v20
    } else {
        v4 = append(v4, "foo bar")
    }
    for fi21 := 0; fi21 < 6; fi21++ {
        for fi22 := 0; fi22 < 7; fi22++ {
            v3 = -1
            fmt.Println(fi21)
            v3 += (v3 & v3)
            _ = fi22
        }
        v3 = 0
        _ = fi21
    }
    v3 += (v3 - v3)
    var v23 uint = 255
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch "foo bar" {
    case "a":
        v6[532] = -87
    case "hello":
        v23 += 65535
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch 266 {
    case 0:
        if (true && true) {
            fmt.Println(v23)
            var v24 int = -308
            _ = v24
        } else {
            var v25 bool = (0 == 255)
            var v26 bool = v25
            _ = v25
            _ = v26
        }
    case 401:
        v23 -= 181
    }
    fmt.Println(v23, v23)
    switch -2147483648 {
    case 640:
        v23 -= (v23 * (^(v23)))
    case 127:
        v27 := GenS3{F0: (v3 * v3)}
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
