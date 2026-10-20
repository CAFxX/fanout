package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int8 }
func genFn0(p0 int16, p1 int16) int16 {
    if (!(p1 != p1)) {
        var v6 float64 = (1e10 / 1.0)
        _ = v6
    }
    v7 := []int{}
    v7 = append(v7, -341)
    var v8 uint = 2
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return p1
}
func genFn1(p0 uint32, p1 uint16, p2 uint16) bool {
    for fi9 := 0; fi9 < 5; fi9++ {
        if true {
            for fi10 := 0; fi10 < 3; fi10++ {
                p1 = 42
                _ = fi10
            }
            var v11 uint = 18446744073709551615
            _ = v11
        } else {
            var v12 int32 = -508
            _ = v12
        }
        _ = fi9
    }
    fmt.Println(p2, p0, p0)
    v13 := GenS1{F0: p1, F1: -73}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return true
}

func main() {
    var v1 uint8 = 255
    v2 := make(map[uint32]bool)
    var v3 uint32 = 779
    v4 := make(map[int]int)
    var v5 uint16 = 65535
    var v14 uint8 = v1
    v3 = (v3 - (v3 + 0))
    var v15 int16 = genFn0(-1, 0)
    var v16 uint32 = (v3 * v3)
    var v17 uint16 = (v5 + v5)
    switch 32768 {
    case 42:
        v14 -= (v1 + 255)
    case 906:
        if true {
            var v18 int8 = 40
            v18 -= v18
            v19 := GenS2{F0: (v3 * uint32(v1)), F1: (v14 / ((v14) | 1))}
            _ = v18
            _ = v19
        } else {
            switch "\n\t" {
            case "\n\t":
                var v20 string = "\n\t"
                _ = v20
            default:
                var v21 int8 = -55
                _ = v21
            }
        }
    default:
        switch 0 {
        case 303:
            var v22 bool = false
            _ = v22
        }
    }
    var v23 int32 = 42
    var v24 int8 = 1
    var v25 bool = (!true)
    for fi26 := 0; fi26 < 20; fi26++ {
        switch ("a" + "foo bar") {
        case "foo bar":
            for fi27 := 0; fi27 < 4; fi27++ {
                v23 += 419
                var v28 uint16 = (v5 << ((2) & 63))
                _ = fi27
                _ = v28
            }
        case "":
            if (221 > 42) {
                v29 := GenS3{F0: v3}
                _ = v29
            }
        }
        v17 = v17
        if (v25 && v25) {
            for fi30 := 0; fi30 < 14; fi30++ {
                var v31 float64 = 3.14159
                _ = fi30
                _ = v31
            }
            var v32 uint = 93
            _ = v32
        } else {
            var v33 string = "\n\t"
            _ = len(v4)
            fmt.Println(v17, v14)
            _ = v33
        }
        _ = fi26
    }
    v1 -= v14
    v34 := GenS4{F0: v24}
    var v35 int16 = v15
    v36 := GenS4{F0: v24}
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v14, v3)
    case "foo bar":
        var v37 uint16 = 65535
        _ = v37
    default:
        if false {
            if (v23 < v23) {
                var v38 int = 961
                _ = v38
            }
        } else {
            var v39 int8 = v24
            _ = v39
        }
    }
    switch "" {
    case "hello":
        fmt.Println(v1, v5)
    case "a":
        v40 := GenS1{F0: 0, F1: 0}
        _ = v40
    case "\n\t":
        fmt.Println(v1)
    }
    var v41 string = ("hello" + "a")
    var v42 uint32 = (v3 - (v16 >> ((743) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v41
    _ = v42
    _ = v5
}
