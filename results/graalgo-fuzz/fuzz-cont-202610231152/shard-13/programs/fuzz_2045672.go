package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 bool; F2 int32 }
func genFn0() uint {
    v6 := []int{32767, 128, 128}
    var v7 float32 = (1e10 * 0.0)
    for fi8 := 0; fi8 < 18; fi8++ {
        v7 += ((v7 / (v7 - 1.5)) * v7)
        var v9 float64 = float64(0.1)
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        for fi11 := 0; fi11 < 20; fi11++ {
            v12 := GenS1{F0: 255, F1: fi11}
            _ = fi11
            _ = v12
        }
        v13 := GenS1{F0: 159, F1: fi10}
        var v14 string = ((("\n\t" + ("hello" + "hello")) + ("foo bar" + "a")) + "hello")
        _ = fi10
        _ = v13
        _ = v14
    }
    fmt.Println(v7)
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 int32, p1 string, p2 bool) int64 {
    var v15 uint8 = 248
    p1 = ((p1 + p1) + p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return -410
}

func main() {
    var v1 uint = 0
    var v2 int16 = 702
    var v3 uint64 = 246
    var v4 int8 = 100
    var v5 int8 = v4
    var v16 float32 = 3.14159
    v17 := GenS2{F0: 2147483647}
    v18 := []int{0, -1, -465}
    v4 -= (-(v5))
    if (795 >= 42) {
        for fi19 := 0; fi19 < 6; fi19++ {
            v3 += 2
            _ = fi19
        }
        v2 = int16(v3)
    }
    if (!("" <= "\n\t")) {
        var v20 bool = true
        _ = v20
    }
    v16 = v16
    v16 = v16
    var v21 uint16 = 2
    fmt.Println(v3)
    for fi22 := 0; fi22 < 6; fi22++ {
        if len(v18) > 0 {
            v18[0] = 410
        }
        if true {
            _ = len(v18)
            v23 := GenS1{F0: 184, F1: fi22}
            v18 = append(v18, fi22)
            _ = v23
        }
        _ = fi22
    }
    fmt.Println(v5, v1)
    var v24 uint16 = (v21 / (((v21 & (v21 + (v21 - 317)))) | 1))
    if ((v3 == v3) && true) {
        v1 = (v1 >> ((v1) & 63))
        switch ("" + "foo bar") {
        case "\n\t":
            v25 := GenS2{F0: -2147483648}
            _ = v25
        case "a":
            v18 = append(v18, 4294967295)
        }
        v18 = append(v18, 32768)
    }
    for fi26 := 0; fi26 < 10; fi26++ {
        var v27 int8 = v5
        v5 = v5
        fmt.Println(v5, v5, v5)
        _ = fi26
        _ = v27
    }
    var v28 uint32 = 4294967295
    switch 128 {
    case -1:
        var v29 int64 = 0
        _ = v29
    case 0:
        var v30 int8 = (v5 / ((42) | 1))
        _ = v30
    case -771:
        switch 32767 {
        case 127:
            var v31 int16 = 1000
            _ = v31
        case 42:
            if true {
                v32 := GenS3{F0: v28, F1: (true && true), F2: -1}
                _ = v32
            }
        case 0:
            v21 += 353
        default:
            if len(v18) >= 2 {
                v18 = v18[1:len(v18)]
            }
        }
    }
    var v33 uint32 = uint32(v21)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v24
    _ = v28
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
