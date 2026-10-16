package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int }
type GenS2 struct { F0 bool; F1 int64; F2 uint32 }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint16, p1 int32, p2 string) int64 {
    v5 := GenS1{F0: (p0 >> ((956) & 63)), F1: 255, F2: -54}
    var v6 int32 = (p1 << ((42) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return -857
}
func genFn1(p0 uint16, p1 int, p2 uint16) int16 {
    v7 := []int{(p1 % (((p1 + p1)) | 1)), p1}
    switch 0 {
    case 0:
        var v8 int32 = int32(p1)
        _ = v8
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        fmt.Println(fi9, p0, fi9)
        for fi10 := 0; fi10 < 15; fi10++ {
            v11 := GenS2{F0: (4294967295 < 4294967295), F1: 791, F2: 4294967295}
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    for fi12 := 0; fi12 < 17; fi12++ {
        p2 = (p0 - p0)
        switch "foo bar" {
        case "a":
            fmt.Println(p2, fi12, p1)
        case "hello":
            if (1 >= (-(p1))) {
                p0 = (p0 * p0)
            } else {
                var v13 uint8 = 127
                var v14 float32 = (3.14159 / -2.25)
                _ = v13
                _ = v14
            }
        case "\n\t":
            var v15 int = fi12
            _ = v15
        }
        switch "foo bar" {
        case "a":
            v16 := GenS2{F0: (42 >= p2), F1: 42, F2: 794}
            _ = v16
        case "foo bar":
            p1 -= (fi12 * (^(p1)))
        default:
            var v17 float32 = (3.14159 / 1.5)
            _ = v17
        }
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 0
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint16 = 65535
    var v3 uint8 = 2
    var v4 bool = false
    v18 := GenS1{F0: v2, F1: 0, F2: 42}
    var v19 uint32 = 0
    for fi20 := 0; fi20 < 5; fi20++ {
        if v4 {
            v4 = (-1 <= -32768)
        } else {
            switch ((("" + "foo bar") + ("" + "foo bar")) + ("\n\t" + ("\n\t" + "\n\t"))) {
            case "foo bar":
                if ((v2 / ((v2) | 1)) <= v2) {
                    var v21 float64 = (0.1 + 1e10)
                    _ = v21
                }
            }
            for fi22 := 0; fi22 < 2; fi22++ {
                v3 -= (v3 & v3)
                _ = fi22
            }
            fmt.Println(fi20, v4)
        }
        _ = fi20
    }
    delete(v1, -2147483648)
    v23 := []int{-850, -2147483648}
    for fi24 := 0; fi24 < 14; fi24++ {
        var v25 string = ("hello" + "a")
        for fi26 := 0; fi26 < 20; fi26++ {
            v27 := GenS3{F0: 0}
            _ = fi26
            _ = v27
        }
        if len(v23) >= 2 {
            v23 = v23[1:len(v23)]
        }
        _ = fi24
        _ = v25
    }
    var v28 int = 4294967295
    if true {
        v28 -= v28
        v19 -= (v19 * v19)
    } else {
        v29 := GenS1{F0: (v2 % (((v2 ^ (v2 % ((v2) | 1)))) | 1)), F1: (^(v3)), F2: (v28 << ((1) & 63))}
        _ = v29
    }
    if (genFn1((v2 % (((v2 - v2)) | 1)), 937, v2) >= 42) {
        v3 = 128
        var v30 uint64 = 276
        _ = v30
    }
    delete(v1, v28)
    if v4 {
        v31 := GenS3{F0: -2147483648}
        fmt.Println(v2, v4, v19)
        v32 := GenS3{F0: -138}
        _ = v31
        _ = v32
    } else {
        var v33 int16 = 1000
        var v34 int = (v28 ^ 127)
        _ = v33
        _ = v34
    }
    for fi35 := 0; fi35 < 18; fi35++ {
        v3 += v3
        _ = fi35
    }
    v28 += v28
    fmt.Println(v19, v3)
    var v36 int16 = genFn1(32767, 42, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v28
    _ = v3
    _ = v36
    _ = v4
}
