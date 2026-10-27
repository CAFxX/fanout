package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 int; F1 uint32; F2 bool }
type GenS4 struct { F0 uint64 }
func genFn0() int {
    v6 := []int{}
    v6 = append(v6, 42)
    var v7 uint8 = 107
    var v8 int16 = -32768
    _ = v6
    _ = v7
    _ = v8
    return 32768
}
func genFn1(p0 uint32) uint8 {
    p0 = (p0 << ((42) & 63))
    p0 -= 2147483647
    for fi9 := 0; fi9 < 20; fi9++ {
        v10 := GenS1{F0: 4294967295, F1: true}
        v11 := []int{(^(fi9)), (fi9 >> ((18446744073709551615) & 63)), (^(fi9))}
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 int = 127
    _ = p0
    _ = v12
    return 128
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]bool)
    var v3 float32 = (float32(((1e10 - (3.14159 / 1.5)) * (1e10 * 1.5))) * 1.5)
    var v4 int16 = 382
    var v5 bool = false
    v5 = true
    for fi13 := 0; fi13 < 10; fi13++ {
        var v14 uint16 = 1
        v15 := GenS2{F0: fi13, F1: -128}
        switch genFn0() {
        case 204:
            var v16 float32 = float32((3.14159 / ((1e10 + 1e10) / 1e10)))
            _ = v16
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    for fi17 := 0; fi17 < 17; fi17++ {
        var v18 float64 = 1.5
        for fi19 := 0; fi19 < 11; fi19++ {
            v18 = v18
            v5 = true
            _ = len(v1)
            _ = fi19
        }
        if v5 {
            v20 := GenS3{F0: 2147483647, F1: 4294967295, F2: true}
            _ = v20
        }
        _ = fi17
        _ = v18
    }
    for fi21 := 0; fi21 < 9; fi21++ {
        v22 := GenS2{F0: fi21, F1: 13}
        v23 := GenS1{F0: 42, F1: v5}
        v24 := GenS4{F0: 0}
        _ = fi21
        _ = v22
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 6; fi25++ {
        v5 = false
        fmt.Println(v4, v5, v4)
        if false {
            var v26 int8 = -128
            switch "" {
            case "":
                v3 = v3
            case "\n\t":
                v26 -= (-(v26))
            default:
                if (-2147483648 != 42) {
                    if true {
                        var v27 string = "hello"
                        _ = v27
                    }
                }
            }
            _ = v26
        }
        _ = fi25
    }
    for fi28 := 0; fi28 < 20; fi28++ {
        var v29 float64 = (0.1 + (1e10 * 0.1))
        _ = fi28
        _ = v29
    }
    if (-9223372036854775808 >= 1) {
        v30 := GenS2{F0: -124, F1: -1}
        var v31 int8 = 42
        for fi32 := 0; fi32 < 20; fi32++ {
            fmt.Println(fi32)
            for fi33 := 0; fi33 < 13; fi33++ {
                fmt.Println(v31, v4, v3)
                _ = fi33
            }
            _ = fi32
        }
        _ = v30
        _ = v31
    }
    v34 := GenS1{F0: 2147483647, F1: v5}
    v4 = (v4 << ((2) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
