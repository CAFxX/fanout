package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 uint64 }
func genFn0() int32 {
    v5 := []int{}
    switch 1 {
    case 32768:
        v6 := GenS1{F0: 2, F1: 1}
        _ = v6
    case -2147483648:
        var v7 string = "a"
        _ = v7
    case 9223372036854775807:
        var v8 int16 = 32767
        _ = v8
    }
    var v9 int8 = 100
    for fi10 := 0; fi10 < 14; fi10++ {
        if len(v5) > 0 {
            v5[0] = -874
        }
        var v11 uint16 = 255
        _ = fi10
        _ = v11
    }
    switch -1 {
    case 285:
        var v12 int32 = 603
        _ = v12
    case 537:
        v9 += (v9 - (v9 + (v9 % (((v9 << ((65535) & 63))) | 1))))
    case 128:
        fmt.Println(v9)
    }
    _ = v5
    _ = v9
    return 2147483647
}
func genFn1(p0 uint64, p1 uint16, p2 uint32) uint32 {
    v13 := []int{}
    p1 = 2
    fmt.Println(p2, p2)
    if false {
        v14 := GenS1{F0: 1, F1: 128}
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return p2
}

func main() {
    v1 := make(map[string]bool)
    var v2 string = "hello"
    v3 := make(map[string]string)
    var v4 float32 = (1.5 + float32(1e10))
    for fi15 := 0; fi15 < 6; fi15++ {
        v4 = v4
        _ = fi15
    }
    v2 = v2
    v2 = "\n\t"
    if true {
        fmt.Println(v2)
        switch ("a" + v2) {
        case "foo bar":
            v16 := GenS2{F0: false}
            _ = v16
        case "\n\t":
            var v17 uint16 = 410
            _ = v17
        case "hello":
            var v18 int64 = -1
            _ = v18
        }
        switch 2147483647 {
        case 145:
            var v19 bool = (!true)
            _ = v19
        case 9223372036854775807:
            var v20 uint32 = 2
            _ = v20
        case 127:
            for fi21 := 0; fi21 < 9; fi21++ {
                var v22 uint64 = 18446744073709551615
                _ = fi21
                _ = v22
            }
        default:
            for fi23 := 0; fi23 < 16; fi23++ {
                _ = len(v3)
                var v24 bool = true
                _ = fi23
                _ = v24
            }
        }
    }
    var v25 uint16 = 32767
    for fi26 := 0; fi26 < 3; fi26++ {
        v27 := GenS3{F0: 242, F1: 48}
        _ = fi26
        _ = v27
    }
    var v28 int32 = 112
    if false {
        if false {
            for fi29 := 0; fi29 < 6; fi29++ {
                var v30 uint16 = (^(v25))
                var v31 uint64 = 4294967295
                _ = fi29
                _ = v30
                _ = v31
            }
        }
        v32 := GenS4{F0: 1}
        v2 = (v2 + (v2 + (v2 + (v2 + ("hello" + v2)))))
        _ = v32
    }
    switch 0 {
    case 32768:
        v2 = v2
    case -254:
        v33 := []int{910, 42, -9223372036854775808, 998}
        _ = v33
    case 9223372036854775807:
        v28 += 969
    default:
        v34 := GenS2{F0: false}
        _ = v34
    }
    var v35 int8 = 100
    v36 := []int{658, 32767, 4294967295, 9223372036854775807}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v28
    _ = v3
    _ = v35
    _ = v36
    _ = v4
}
