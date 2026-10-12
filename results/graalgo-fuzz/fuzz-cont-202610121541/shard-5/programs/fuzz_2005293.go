package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 uint64 }
type GenS2 struct { F0 int8; F1 int32; F2 int16 }
type GenS3 struct { F0 string; F1 int16; F2 int64 }
type GenS4 struct { F0 bool }
func genFn0(p0 uint64) bool {
    v5 := []int{725}
    if ((((0.1 + float32(1e10)) * (-2.25 + (-2.25 / (-2.25 + 0.1)))) != 1.5) || false) {
        var v6 uint16 = 32767
        for fi7 := 0; fi7 < 1; fi7++ {
            v6 = 255
            switch 32768 {
            case -336:
                var v8 uint64 = (p0 / ((p0) | 1))
                _ = v8
            }
            _ = fi7
        }
        _ = v6
    }
    var v9 uint = 65535
    for fi10 := 0; fi10 < 3; fi10++ {
        if (42 <= -128) {
            v11 := GenS1{F0: 2147483647, F1: 1, F2: p0}
            for fi12 := 0; fi12 < 16; fi12++ {
                v9 += 42
                fmt.Println(v9)
                _ = fi12
            }
            v5 = append(v5, fi10)
            _ = v11
        } else {
            for fi13 := 0; fi13 < 17; fi13++ {
                v14 := GenS2{F0: 0, F1: 2147483647, F2: 32767}
                var v15 string = ("foo bar" + "a")
                v15 = (((((("a" + (((v15 + (v15 + ("" + ""))) + (((v15 + v15) + ("hello" + (v15 + (v15 + v15)))) + (((v15 + v15) + v15) + v15))) + v15)) + (((("a" + "") + v15) + v15) + (v15 + v15))) + "hello") + "") + v15) + ("a" + ""))
                _ = fi13
                _ = v14
                _ = v15
            }
        }
        if (1 > 1000) {
            v16 := GenS2{F0: int8(fi10), F1: -515, F2: -1}
            _ = v16
        } else {
            var v17 uint8 = uint8(p0)
            _ = v17
        }
        fmt.Println(v9)
        _ = fi10
    }
    p0 = p0
    _ = p0
    _ = v5
    _ = v9
    return true
}

func main() {
    var v1 uint8 = 101
    var v2 int = 790
    var v3 int16 = 324
    var v4 int = v2
    for fi18 := 0; fi18 < 17; fi18++ {
        var v19 float32 = 3.14159
        v20 := []int{}
        var v21 int8 = 94
        _ = fi18
        _ = v19
        _ = v20
        _ = v21
    }
    v22 := GenS1{F0: -2147483648, F1: -165, F2: 42}
    var v23 uint16 = 502
    v1 += (v1 | (v1 << ((429) & 63)))
    v24 := GenS3{F0: ("hello" + ("hello" + ("a" + "hello"))), F1: 1000, F2: 0}
    fmt.Println(v2, v23)
    v25 := GenS4{F0: ((-(v4)) < 127)}
    var v26 string = "\n\t"
    var v27 int = (v2 << ((285) & 63))
    if true {
        switch ((((v26 + v26) + v26) + v26) + v26) {
        case "\n\t":
            v28 := GenS4{F0: (401 < -1)}
            _ = v28
        case "":
            var v29 uint = 0
            _ = v29
        case "hello":
            if true {
                v30 := []int{(v27 - (v4 + (v2 * (v4 & -2147483648)))), v2, (-(v4)), (v4 * (v27 << ((86) & 63)))}
                _ = v30
            }
        default:
            var v31 uint16 = (v23 ^ v23)
            _ = v31
        }
        v32 := []int{}
        _ = v32
    } else {
        v2 -= v4
        var v33 string = ("a" + "a")
        var v34 uint16 = v23
        _ = v33
        _ = v34
    }
    var v35 uint = 18446744073709551615
    switch (v2 / (((v4 * 32768)) | 1)) {
    case -589:
        v36 := []int{(^(v4)), (v2 << ((uint(v1)) & 63)), (v2 * v4)}
        _ = v36
    case -2147483648:
        v37 := GenS2{F0: 127, F1: 0, F2: (-(v3))}
        _ = v37
    }
    fmt.Println(v2, v2, v26)
    for fi38 := 0; fi38 < 9; fi38++ {
        fmt.Println(v35)
        _ = fi38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v35
    _ = v4
}
