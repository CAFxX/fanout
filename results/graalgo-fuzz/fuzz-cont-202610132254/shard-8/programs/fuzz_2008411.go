package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 int64; F1 int32; F2 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int16; F1 uint8; F2 uint8 }
type GenS5 struct { F0 int; F1 uint32 }
func genFn0(p0 int, p1 int32, p2 uint8) int64 {
    p1 += p1
    switch (p0 + (p0 << ((944) & 63))) {
    case 557:
        v7 := []int{42, 127}
        _ = v7
    case 127:
        p2 = p2
    case -138:
        v8 := []int{p0, 32767, 4294967295, (p0 >> ((2) & 63))}
        _ = v8
    default:
        v9 := GenS1{F0: 956, F1: 100}
        _ = v9
    }
    v10 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return -608
}
func genFn1(p0 string, p1 uint8) int8 {
    p0 = ((p0 + ("hello" + "foo bar")) + "foo bar")
    if true {
        if (363 != 1) {
            for fi11 := 0; fi11 < 16; fi11++ {
                fmt.Println(fi11, p0, p0)
                _ = fi11
            }
        } else {
            p1 += 1
            fmt.Println(p0)
        }
        var v12 uint16 = 1
        _ = v12
    } else {
        v13 := GenS1{F0: 2, F1: -128}
        v14 := []int{2147483647, -514, 9223372036854775807}
        v15 := GenS2{F0: 0, F1: -620, F2: p0}
        _ = v13
        _ = v14
        _ = v15
    }
    fmt.Println(p1)
    if (-971 != 0) {
        p1 = 128
        fmt.Println(p1, p1)
    }
    _ = p0
    _ = p1
    return 127
}

func main() {
    var v1 uint64 = 42
    v2 := make(map[int]uint32)
    var v3 int = -103
    v4 := make(map[int]uint32)
    var v5 int8 = 127
    var v6 int16 = 767
    switch 127 {
    case 648:
        v16 := []int{(v3 & v3), v3}
        _ = v16
    case -1:
        switch (v3 >> ((2) & 63)) {
        case 9223372036854775807:
            v2[v3] = 676
        }
    case 0:
        var v17 uint = 2
        _ = v17
    }
    var v18 bool = ((v5 <= (v5 % ((int8(v6)) | 1))) || false)
    v19 := GenS3{F0: true}
    v5 = v5
    for fi20 := 0; fi20 < 12; fi20++ {
        v21 := GenS2{F0: 754, F1: 42, F2: "\n\t"}
        _ = fi20
        _ = v21
    }
    v22 := []int{v3, 311}
    if ((v6 << ((4294967295) & 63)) > v6) {
        switch ("" + ("hello" + "hello")) {
        case "\n\t":
            var v23 int64 = 1
            _ = v23
        case "hello":
            v24 := GenS2{F0: -407, F1: -299, F2: "hello"}
            _ = v24
        default:
            v1 += (^(v1))
        }
        fmt.Println(v5, v3, v3)
        var v25 bool = (!true)
        _ = v25
    }
    for fi26 := 0; fi26 < 13; fi26++ {
        v18 = false
        v27 := GenS2{F0: 9223372036854775807, F1: 1, F2: ((("hello" + "hello") + "foo bar") + ("foo bar" + "foo bar"))}
        var v28 int8 = (^(v5))
        _ = fi26
        _ = v27
        _ = v28
    }
    v3 = int(v5)
    v5 -= (v5 ^ genFn1("foo bar", 127))
    v6 = v6
    for fi29 := 0; fi29 < 10; fi29++ {
        v6 = int16(v5)
        v1 = v1
        v30 := GenS4{F0: int16(v1), F1: 127, F2: 128}
        _ = fi29
        _ = v30
    }
    for fi31 := 0; fi31 < 2; fi31++ {
        v32 := GenS5{F0: (-(v3)), F1: 2}
        _ = fi31
        _ = v32
    }
    v1 -= v1
    if v18 {
        for fi33 := 0; fi33 < 15; fi33++ {
            var v34 string = "\n\t"
            v5 += genFn1(v34, 255)
            if (!v18) {
                var v35 int = v3
                _ = v35
            }
            _ = fi33
            _ = v34
        }
        var v36 bool = (-1 > 0)
        var v37 uint32 = 0
        _ = v36
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
