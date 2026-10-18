package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 uint16 }
type GenS2 struct { F0 int }
func genFn0(p0 bool, p1 uint16, p2 int32) int64 {
    fmt.Println(p0, p2)
    for fi6 := 0; fi6 < 8; fi6++ {
        p2 = p2
        _ = fi6
    }
    switch ("a" + ("foo bar" + "a")) {
    case "a":
        if (!p0) {
            var v7 int64 = -9223372036854775808
            p0 = (1000 == 0)
            p2 = (p2 & (p2 / (((p2 % (((p2 >> ((65535) & 63))) | 1))) | 1)))
            _ = v7
        } else {
            var v8 uint16 = (p1 ^ (p1 ^ (p1 % ((p1) | 1))))
            for fi9 := 0; fi9 < 4; fi9++ {
                var v10 int32 = int32(fi9)
                for fi11 := 0; fi11 < 17; fi11++ {
                    v12 := []int{(fi9 / ((fi11) | 1)), fi11, (fi9 + (fi9 / (((fi11 + -891)) | 1))), fi9}
                    p0 = p0
                    if true {
                        p2 = int32(fi9)
                        if len(v12) >= 2 {
                            v12 = v12[1:len(v12)]
                        }
                    }
                    fmt.Println(fi9, fi9)
                    _ = fi11
                    _ = v12
                }
                p2 += 567
                _ = fi9
                _ = v10
            }
            _ = v8
        }
    case "":
        v13 := []int{}
        _ = v13
    case "foo bar":
        for fi14 := 0; fi14 < 10; fi14++ {
            if true {
                var v15 int16 = 1
                v16 := GenS1{F0: ("foo bar" == (("hello" + "a") + "hello")), F1: 42, F2: (p1 ^ 42)}
                v17 := []int{(^(fi14)), 127, (fi14 ^ fi14)}
                _ = v15
                _ = v16
                _ = v17
            } else {
                p2 = -2147483648
            }
            _ = fi14
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return int64(p1)
}
func genFn1() bool {
    var v18 float32 = float32(0.0)
    var v19 uint64 = 18446744073709551615
    switch "hello" {
    case "hello":
        if true {
            v19 += v19
        }
    case "\n\t":
        v18 = v18
    case "a":
        v20 := []int{}
        _ = v20
    }
    _ = v18
    _ = v19
    return true
}

func main() {
    v1 := make(map[string]string)
    var v2 float64 = 0.0
    var v3 int32 = 1
    v4 := make(map[int]string)
    v5 := []uint32{}
    for fi21 := 0; fi21 < 15; fi21++ {
        switch fi21 {
        case -359:
            v3 = v3
        }
        v22 := GenS2{F0: fi21}
        _ = fi21
        _ = v22
    }
    v2 = v2
    _ = len(v1)
    v2 += float64(1e10)
    v23 := GenS2{F0: -328}
    v24 := GenS1{F0: true, F1: 42, F2: 0}
    var v25 uint8 = 135
    if ((true || (1 > 0)) || false) {
        fmt.Println(v25, v25)
        v26 := GenS2{F0: 128}
        _ = v26
    }
    fmt.Println(v3, v3, v3)
    var v27 float64 = float64(1.5)
    var v28 uint64 = 2
    var v29 int64 = -874
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
