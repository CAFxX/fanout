package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int64 }
func genFn0(p0 int) int16 {
    v6 := []int{p0, (p0 | 9223372036854775807), 4294967295}
    var v7 float32 = 0.0
    var v8 int8 = 0
    fmt.Println(v8, v7)
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return -32768
}
func genFn1(p0 uint32, p1 uint16) int8 {
    p0 = uint32(p1)
    v9 := []int{-1}
    var v10 int = int(p0)
    if true {
        if (255 <= 136) {
            v11 := GenS1{F0: 285, F1: 1}
            for fi12 := 0; fi12 < 8; fi12++ {
                if true {
                    fmt.Println(fi12, fi12)
                    p0 += (^(p0))
                } else {
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    p0 = (p0 / (((p0 >> ((42) & 63))) | 1))
                    v13 := GenS1{F0: 345, F1: (p1 & (p1 ^ p1))}
                    _ = v13
                }
                _ = fi12
            }
            _ = v11
        }
        fmt.Println(v10, p0)
        var v14 int32 = 1
        _ = v14
    } else {
        for fi15 := 0; fi15 < 5; fi15++ {
            v16 := GenS1{F0: 0, F1: 756}
            _ = fi15
            _ = v16
        }
        v17 := GenS2{F0: (v10 >> ((4294967295) & 63))}
        _ = v17
    }
    var v18 float64 = -2.25
    _ = p0
    _ = p1
    _ = v10
    _ = v18
    _ = v9
    return 127
}

func main() {
    v1 := make(map[int]int)
    v2 := []string{}
    var v3 int8 = 0
    v4 := make(map[int]string)
    v5 := make(map[string]int)
    fmt.Println(v3)
    var v19 int32 = 0
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi20 := 0; fi20 < 16; fi20++ {
        switch "foo bar" {
        case "hello":
            if true {
                var v21 uint8 = 128
                if (!(!(true && (!true)))) {
                    var v22 bool = false
                    v21 -= (v21 >> ((0) & 63))
                    switch (fi20 + (fi20 + fi20)) {
                    case -766:
                        fmt.Println(fi20, v22)
                    case 0:
                        switch "a" {
                        case "a":
                            v19 += v19
                        }
                    }
                    _ = v22
                } else {
                    v19 -= -1
                    fmt.Println(v3, v3)
                    var v23 uint32 = 1
                    _ = v23
                }
                _ = v21
            }
        case "":
            for fi24 := 0; fi24 < 8; fi24++ {
                fmt.Println(v3)
                v19 -= (v19 + 2147483647)
                v19 -= -586
                _ = fi24
            }
        case "a":
            var v25 uint64 = 4294967295
            _ = v25
        }
        for fi26 := 0; fi26 < 11; fi26++ {
            v27 := GenS3{F0: 52}
            var v28 uint32 = 1
            fmt.Println(fi26, fi26)
            fmt.Println(fi26, v19)
            v29 := GenS1{F0: 635, F1: 0}
            _ = fi26
            _ = v27
            _ = v28
            _ = v29
        }
        _ = fi20
    }
    if false {
        if (!false) {
            if len(v2) > 0 {
                v2[0] = ""
            }
            v30 := GenS4{F0: v19}
            v31 := GenS4{F0: v19}
            _ = v30
            _ = v31
        }
        if len(v2) > 0 {
            v2[0] = "hello"
        }
        v19 -= v19
    }
    fmt.Println(v19, v3)
    var v32 float64 = 0.0
    v33 := GenS5{F0: -1}
    fmt.Println(v3, v3, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
