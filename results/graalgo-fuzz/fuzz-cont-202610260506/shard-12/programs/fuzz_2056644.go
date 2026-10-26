package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16 }
func genFn0(p0 int32) int32 {
    v5 := []int{}
    var v6 float32 = (3.14159 / 1.0)
    if true {
        v7 := GenS1{F0: 18446744073709551615, F1: (p0 / (((-(p0))) | 1)), F2: -1}
        _ = v7
    }
    fmt.Println(p0)
    _ = p0
    _ = v5
    _ = v6
    return p0
}
func genFn1() uint16 {
    var v8 int16 = 42
    switch "foo bar" {
    case "":
        var v9 uint = 849
        _ = v9
    case "a":
        v10 := []int{0}
        _ = v10
    case "hello":
        v11 := []int{146, -762, -508}
        _ = v11
    }
    if false {
        v12 := GenS2{F0: (v8 & v8)}
        v13 := []int{}
        _ = v12
        _ = v13
    }
    fmt.Println(v8)
    _ = v8
    return 0
}

func main() {
    var v1 int = 32767
    var v2 uint8 = 240
    v3 := []uint32{}
    var v4 uint = 0
    for fi14 := 0; fi14 < 2; fi14++ {
        var v15 float64 = (float64((0.0 - -2.25)) + 0.1)
        _ = fi14
        _ = v15
    }
    var v16 uint16 = 922
    v17 := GenS2{F0: -1}
    for fi18 := 0; fi18 < 2; fi18++ {
        for fi19 := 0; fi19 < 7; fi19++ {
            var v20 uint32 = 1
            var v21 uint32 = v20
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi19
            _ = v20
            _ = v21
        }
        _ = fi18
    }
    var v22 uint16 = v16
    if len(v3) > 0 {
        v3[0] = 4294967295
    }
    switch v1 {
    case -563:
        v23 := GenS3{F0: 115}
        _ = v23
    case 9223372036854775807:
        var v24 uint64 = 740
        _ = v24
    }
    v25 := GenS1{F0: 42, F1: -1, F2: 1}
    _ = len(v3)
    var v26 uint8 = (v2 & (v2 % ((0) | 1)))
    v27 := GenS2{F0: -772}
    _ = len(v3)
    for fi28 := 0; fi28 < 13; fi28++ {
        switch (fi28 / (((-(v1))) | 1)) {
        case 128:
            if len(v3) > 0 {
                v3[0] = 694
            }
        }
        _ = fi28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
