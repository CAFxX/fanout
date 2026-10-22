package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint8; F1 uint64 }
func genFn0(p0 string, p1 int64, p2 bool) int8 {
    if ((p1 <= (p1 | p1)) || p2) {
        var v5 int32 = 2147483647
        if (1 <= 42) {
            fmt.Println(p1, v5, v5)
            v6 := []int{32768}
            _ = v6
        } else {
            v7 := GenS1{F0: 0}
            v8 := GenS2{F0: 1, F1: false}
            _ = v7
            _ = v8
        }
        _ = v5
    } else {
        fmt.Println(p1, p0)
        p1 -= p1
    }
    for fi9 := 0; fi9 < 20; fi9++ {
        if p2 {
            fmt.Println(fi9, p1)
            p2 = (false && (107 == -128))
            fmt.Println(fi9)
        } else {
            v10 := []int{(fi9 | fi9), (fi9 * fi9)}
            _ = v10
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = p2
    return 127
}
func genFn1() int8 {
    v11 := []int{}
    for fi12 := 0; fi12 < 15; fi12++ {
        if ((1.5 != 0.0) && (true || (!(((false || (false && (int8(fi12) < 63))) || false) || false)))) {
            var v13 uint32 = 1
            var v14 int32 = -1
            var v15 string = ("" + "")
            _ = v13
            _ = v14
            _ = v15
        } else {
            fmt.Println(fi12)
            for fi16 := 0; fi16 < 19; fi16++ {
                var v17 uint64 = 18446744073709551615
                v18 := GenS2{F0: 4294967295, F1: false}
                _ = fi16
                _ = v17
                _ = v18
            }
            v19 := GenS1{F0: 255}
            _ = v19
        }
        var v20 uint16 = 655
        for fi21 := 0; fi21 < 10; fi21++ {
            fmt.Println(fi12, v20)
            _ = fi21
        }
        _ = fi12
        _ = v20
    }
    var v22 int8 = 5
    _ = v11
    _ = v22
    return genFn0((("a" + "hello") + "hello"), 42, (0 <= 18446744073709551615))
}

func main() {
    var v1 bool = (!true)
    v2 := make(map[int]uint32)
    var v3 uint8 = 1
    var v4 string = ("foo bar" + "")
    var v23 int = -2147483648
    fmt.Println(v1)
    for fi24 := 0; fi24 < 16; fi24++ {
        v25 := GenS3{F0: 2}
        fmt.Println(v4, v3)
        _ = fi24
        _ = v25
    }
    v1 = false
    v26 := GenS1{F0: 65535}
    var v27 int32 = 862
    if (!v1) {
        var v28 int32 = (v27 ^ 772)
        switch v23 {
        case -160:
            v29 := v2[v23]
            _ = v29
        case -331:
            v4 = (v4 + "")
        }
        _ = v28
    }
    for fi30 := 0; fi30 < 13; fi30++ {
        for fi31 := 0; fi31 < 17; fi31++ {
            _ = len(v2)
            _ = fi31
        }
        _ = fi30
    }
    v32 := GenS4{F0: (v3 ^ (v3 ^ v3)), F1: 350}
    var v33 uint32 = 82
    var v34 float64 = (1e10 / 0.1)
    switch (v4 + "foo bar") {
    case "hello":
        for fi35 := 0; fi35 < 11; fi35++ {
            v36 := GenS3{F0: 65535}
            _ = fi35
            _ = v36
        }
    case "foo bar":
        for fi37 := 0; fi37 < 16; fi37++ {
            v33 += v33
            _ = fi37
        }
    case "\n\t":
        if (!(!(false || v1))) {
            var v38 bool = (v3 >= 127)
            v39 := []int{}
            if (true && v38) {
                fmt.Println(v33)
            }
            _ = v38
            _ = v39
        } else {
            fmt.Println(v3, v27, v34)
        }
    default:
        fmt.Println(v33)
    }
    switch v4 {
    case "\n\t":
        v34 += v34
    case "hello":
        if v1 {
            v27 -= v27
        } else {
            v40 := []int{v23}
            _ = v40
        }
    case "":
        var v41 int8 = 1
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
