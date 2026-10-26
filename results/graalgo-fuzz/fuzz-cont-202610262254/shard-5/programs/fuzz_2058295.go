package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint; F1 uint32; F2 uint }
type GenS5 struct { F0 int32; F1 int; F2 int32 }
func genFn0(p0 string) int16 {
    var v6 int = -970
    v6 = 127
    if (!(42 < 65535)) {
        p0 = p0
        v7 := []int{680}
        v6 = 506
        _ = v7
    } else {
        fmt.Println(v6, v6)
    }
    var v8 string = (((p0 + "foo bar") + (p0 + (p0 + ((p0 + ("foo bar" + "\n\t")) + (((p0 + p0) + (p0 + p0)) + "a"))))) + p0)
    switch v8 {
    case "\n\t":
        v9 := GenS1{F0: -1, F1: -32768}
        _ = v9
    default:
        for fi10 := 0; fi10 < 20; fi10++ {
            for fi11 := 0; fi11 < 1; fi11++ {
                switch "\n\t" {
                case "\n\t":
                    fmt.Println(v6, v8, fi10)
                    p0 = "foo bar"
                    fmt.Println(p0)
                case "foo bar":
                    v12 := GenS1{F0: -32768, F1: 1}
                    var v13 float64 = (-2.25 * 1.5)
                    v14 := GenS2{F0: 135, F1: -451}
                    _ = v12
                    _ = v13
                    _ = v14
                case "hello":
                    fmt.Println(v8, v8, v8)
                }
                v15 := GenS3{F0: (p0 + p0)}
                _ = fi11
                _ = v15
            }
            v16 := GenS3{F0: v8}
            _ = fi10
            _ = v16
        }
    }
    _ = p0
    _ = v6
    _ = v8
    return -32768
}
func genFn1(p0 bool, p1 uint8, p2 uint8) int8 {
    var v17 int64 = 769
    switch ("\n\t" + "foo bar") {
    case "":
        v18 := GenS4{F0: 902, F1: 42, F2: 2}
        _ = v18
    case "a":
        var v19 bool = (76 < 42)
        _ = v19
    case "foo bar":
        fmt.Println(p0, p1)
    }
    v17 = (v17 + (-(v17)))
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return 42
}

func main() {
    var v1 int64 = -554
    v2 := make(map[uint32]string)
    var v3 int16 = 1000
    v4 := []int{435, -2147483648, 127}
    v5 := make(map[int]bool)
    _ = len(v2)
    v3 -= v3
    fmt.Println(v3)
    v1 -= (v1 & v1)
    _ = len(v2)
    v20 := GenS1{F0: v3, F1: genFn0("\n\t")}
    _ = len(v4)
    switch -1 {
    case -2147483648:
        var v21 uint64 = 1
        _ = v21
    case 128:
        v1 = v1
    case -630:
        var v22 float64 = 0.0
        _ = v22
    }
    for fi23 := 0; fi23 < 15; fi23++ {
        var v24 uint = 18446744073709551615
        _ = fi23
        _ = v24
    }
    var v25 uint32 = 42
    if true {
        for fi26 := 0; fi26 < 11; fi26++ {
            for fi27 := 0; fi27 < 17; fi27++ {
                var v28 int = fi26
                v3 += (-(v3))
                _ = fi27
                _ = v28
            }
            v29 := GenS1{F0: 42, F1: v3}
            _ = fi26
            _ = v29
        }
    }
    switch ("a" + "a") {
    case "hello":
        for fi30 := 0; fi30 < 13; fi30++ {
            var v31 int32 = 2147483647
            var v32 uint32 = v25
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi30
            _ = v31
            _ = v32
        }
    }
    if (-2.25 == (float64(1.5) / 1.0)) {
        v33 := GenS5{F0: -80, F1: -116, F2: 2147483647}
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
