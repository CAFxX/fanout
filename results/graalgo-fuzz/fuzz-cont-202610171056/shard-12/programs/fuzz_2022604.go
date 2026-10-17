package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 bool; F1 uint }
func genFn0(p0 int32, p1 bool, p2 int8) int32 {
    if ((-2.25 > 1.5) || ((3.14159 != 1e10) || (p1 || (!(!false))))) {
        p1 = true
    }
    if false {
        var v6 float32 = 0.0
        _ = v6
    }
    switch -509 {
    case 32767:
        v7 := GenS1{F0: 18446744073709551615, F1: (p2 != -12), F2: 42}
        _ = v7
    default:
        v8 := []int{-991, 897, 640}
        _ = v8
    }
    if (2 > 1) {
        v9 := []int{1, -1, 1}
        switch 128 {
        case 435:
            var v10 int64 = 1
            _ = v10
        case -2147483648:
            p1 = false
        default:
            v11 := GenS1{F0: 1, F1: (false && p1), F2: 128}
            _ = v11
        }
        for fi12 := 0; fi12 < 6; fi12++ {
            v13 := GenS2{F0: 18446744073709551615}
            _ = len(v9)
            v14 := GenS1{F0: 2, F1: (4294967295 >= 694), F2: 0}
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = v9
    }
    p0 -= 869
    _ = p0
    _ = p1
    _ = p2
    return (p0 << ((4294967295) & 63))
}
func genFn1(p0 int, p1 int) uint8 {
    v15 := []int{p0, (p1 << ((2) & 63)), p0, (p1 % ((p0) | 1))}
    v16 := GenS1{F0: 1, F1: true, F2: 127}
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    return 0
}

func main() {
    v1 := make(map[int]int)
    var v2 int = 2147483647
    var v3 uint = 18446744073709551615
    v4 := []string{"a", "a", "hello"}
    var v5 float32 = ((-2.25 + 1.5) / -2.25)
    v17 := GenS3{F0: 2147483648, F1: -9223372036854775808}
    v2 = 2147483647
    for fi18 := 0; fi18 < 3; fi18++ {
        switch (^(fi18)) {
        case 1:
            v1[fi18] = 0
        case -1:
            if (!true) {
                fmt.Println(v2)
                fmt.Println(v3, v5)
            }
        case -935:
            v5 += ((v5 + -2.25) / 1.0)
        }
        v3 = 679
        v1[(fi18 * -483)] = (-(fi18))
        _ = fi18
    }
    fmt.Println(v2, v2)
    v3 -= (v3 * v3)
    switch ("" + "hello") {
    case "foo bar":
        fmt.Println(v2, v5)
    case "a":
        for fi19 := 0; fi19 < 7; fi19++ {
            v2 += (v2 & (fi19 - 32768))
            var v20 int = (fi19 ^ fi19)
            fmt.Println(v3, fi19)
            _ = fi19
            _ = v20
        }
    default:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    fmt.Println(v3)
    switch 32768 {
    case 805:
        var v21 uint = v3
        _ = v21
    case 32768:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case -414:
        _ = len(v4)
    default:
        v3 = (v3 >> ((500) & 63))
    }
    fmt.Println(v5, v3)
    v4 = append(v4, "a")
    fmt.Println(v2)
    v2 = 0
    var v22 int16 = 32767
    for fi23 := 0; fi23 < 3; fi23++ {
        v24 := GenS3{F0: 65535, F1: 0}
        for fi25 := 0; fi25 < 13; fi25++ {
            v26 := GenS2{F0: v3}
            v27 := GenS4{F0: (-941 <= -2147483648), F1: v3}
            _ = fi25
            _ = v26
            _ = v27
        }
        _ = fi23
        _ = v24
    }
    var v28 uint8 = genFn1(v2, -414)
    var v29 uint = 792
    v3 -= v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
