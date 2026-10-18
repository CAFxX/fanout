package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int16 }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 bool; F1 uint64 }
func genFn0(p0 int8, p1 string) int32 {
    if (!false) {
        v6 := []int{127}
        var v7 float64 = 0.0
        p1 = ""
        _ = v6
        _ = v7
    } else {
        var v8 uint16 = 2
        var v9 int64 = 42
        _ = v8
        _ = v9
    }
    switch "\n\t" {
    case "":
        p0 += p0
    case "foo bar":
        switch p1 {
        case "":
            v10 := []int{287, -2147483648}
            _ = v10
        }
    case "\n\t":
        var v11 uint = 255
        _ = v11
    }
    for fi12 := 0; fi12 < 17; fi12++ {
        switch (p1 + "\n\t") {
        case "foo bar":
            v13 := GenS1{F0: p0, F1: (p1 + (("\n\t" + p1) + "foo bar")), F2: 0}
            _ = v13
        case "\n\t":
            p0 += -1
        case "":
            var v14 float32 = 1.5
            _ = v14
        default:
            switch (p1 + (((p1 + ("foo bar" + (p1 + "a"))) + p1) + ("\n\t" + (p1 + ((p1 + "a") + (p1 + p1)))))) {
            case "foo bar":
                var v15 uint = 35
                _ = v15
            case "a":
                v16 := GenS2{F0: fi12, F1: 2147483647}
                _ = v16
            case "hello":
                p0 = p0
            }
        }
        _ = fi12
    }
    fmt.Println(p0)
    p1 = "\n\t"
    _ = p0
    _ = p1
    return 0
}
func genFn1(p0 uint, p1 uint8) int64 {
    v17 := []int{-2147483648, 9223372036854775807, 4294967295}
    var v18 int8 = -1
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    return -1
}

func main() {
    v1 := []int{42, 42, -1, -640}
    v2 := []int{-2147483648, -2147483648, 32767}
    var v3 int8 = -1
    var v4 float64 = 0.0
    v5 := make(map[string]uint32)
    v3 = (-(v3))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (v4 == v4) {
        var v19 int16 = -32768
        v3 = v3
        _ = v19
    }
    v3 = -128
    v4 -= 3.14159
    if true {
        var v20 bool = ((v3 >= v3) && (1 < 2147483647))
        v3 += (^(v3))
        var v21 uint8 = 2
        _ = v20
        _ = v21
    }
    if (("foo bar" + "foo bar") >= "foo bar") {
        var v22 uint = 382
        v3 -= -121
        _ = v22
    } else {
        _ = len(v1)
    }
    var v23 uint64 = 86
    if (true || false) {
        v4 -= 0.1
        var v24 string = ("a" + ("" + "hello"))
        v25 := GenS3{F0: (!((-32768 > 1) || ((!((((v4 * v4) - 0.0) > (v4 + (v4 + 1.5))) || ((v24 != v24) && (!false)))) && (genFn0(int8(v23), "foo bar") != 1)))), F1: (v23 & v23)}
        _ = v24
        _ = v25
    } else {
        v3 += v3
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v3 += v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
