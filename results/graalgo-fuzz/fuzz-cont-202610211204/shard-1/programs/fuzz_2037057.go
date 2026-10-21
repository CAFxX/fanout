package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 uint16; F2 uint8 }
func genFn0(p0 int16, p1 int64) uint64 {
    p1 -= (p1 - (p1 & 1))
    fmt.Println(p1, p0)
    for fi5 := 0; fi5 < 12; fi5++ {
        v6 := GenS1{F0: 207}
        switch fi5 {
        case -775:
            p0 -= 1000
        case -180:
            var v7 uint64 = 42
            _ = v7
        case 127:
            p1 = -1
        }
        _ = fi5
        _ = v6
    }
    var v8 int8 = 100
    _ = p0
    _ = p1
    _ = v8
    return 760
}

func main() {
    var v1 string = ("" + "")
    var v2 string = "hello"
    var v3 uint8 = 255
    v4 := make(map[int]bool)
    if (752 > int16(v3)) {
        for fi9 := 0; fi9 < 17; fi9++ {
            var v10 float32 = 1e10
            v3 = 1
            _ = fi9
            _ = v10
        }
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        for fi12 := 0; fi12 < 19; fi12++ {
            v13 := []int{(fi11 & (-(fi11))), fi12}
            v14 := GenS1{F0: 285}
            _ = fi12
            _ = v13
            _ = v14
        }
        for fi15 := 0; fi15 < 4; fi15++ {
            v2 = ""
            v1 = "a"
            v2 = ("" + v2)
            _ = fi15
        }
        v16 := GenS1{F0: 898}
        _ = fi11
        _ = v16
    }
    v17 := GenS2{F0: 1000, F1: 0, F2: 2}
    switch v1 {
    case "\n\t":
        v18 := v4[-396]
        _ = v18
    case "":
        if (!(false && (!true))) {
            v3 += v3
        }
    case "a":
        fmt.Println(v1, v3, v3)
    }
    v19 := GenS1{F0: 198}
    v1 = v2
    fmt.Println(v1, v3)
    switch 1 {
    case -243:
        v20 := []int{32767, 128}
        _ = v20
    default:
        var v21 int = -52
        _ = v21
    }
    var v22 float32 = 0.1
    var v23 string = v1
    v3 = v3
    var v24 int = 42
    for fi25 := 0; fi25 < 2; fi25++ {
        v26 := GenS1{F0: 2}
        for fi27 := 0; fi27 < 9; fi27++ {
            v22 = (v22 - 3.14159)
            var v28 uint16 = 2
            v23 = v2
            var v29 bool = (65535 != uint16(v3))
            var v30 int32 = 42
            var v31 uint16 = 1
            _ = fi27
            _ = v28
            _ = v29
            _ = v30
            _ = v31
        }
        delete(v4, v24)
        _ = fi25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
