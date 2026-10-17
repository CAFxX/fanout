package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64 }

func main() {
    var v1 int8 = 100
    var v2 int = -343
    var v3 float64 = 1.5
    var v4 float64 = (3.14159 / float64(1e10))
    var v5 float64 = (v3 - v4)
    var v6 uint8 = 127
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 int16 = -1
        for fi9 := 0; fi9 < 3; fi9++ {
            v10 := []int{42}
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    if (!true) {
        v11 := []int{}
        _ = v11
    } else {
        var v12 uint = 18446744073709551615
        var v13 string = "foo bar"
        _ = v12
        _ = v13
    }
    if false {
        v2 -= v2
        var v14 uint = 135
        v14 -= v14
        _ = v14
    }
    var v15 int = (v2 - (v2 % ((v2) | 1)))
    v4 += (-2.25 * v3)
    for fi16 := 0; fi16 < 4; fi16++ {
        if (!true) {
            switch "a" {
            case "a":
                v17 := []int{v15, (v15 ^ v2)}
                _ = v17
            default:
                v18 := GenS1{F0: 127, F1: 0}
                _ = v18
            }
            for fi19 := 0; fi19 < 19; fi19++ {
                v1 = v1
                v20 := GenS2{F0: -1, F1: ("\n\t" + "")}
                if (true && (-1 > -1)) {
                    v5 -= v4
                } else {
                    v3 -= v5
                    fmt.Println(fi19)
                    v21 := GenS3{F0: v1}
                    _ = v21
                }
                _ = fi19
                _ = v20
            }
        } else {
            var v22 uint16 = 65535
            v23 := GenS4{F0: 42}
            if true {
                var v24 uint16 = v22
                _ = v24
            }
            _ = v22
            _ = v23
        }
        _ = fi16
    }
    v1 = 72
    for fi25 := 0; fi25 < 8; fi25++ {
        fmt.Println(v6, v4, v6)
        _ = fi25
    }
    var v26 uint64 = 18446744073709551615
    v27 := []int{v15}
    if (1 == 771) {
        v28 := GenS1{F0: v1, F1: 1}
        if ((!false) || false) {
            fmt.Println(v2, v5, v4)
        } else {
            if (v1 >= v1) {
                v15 -= (v15 / ((32767) | 1))
            }
        }
        _ = v28
    } else {
        fmt.Println(v2, v2, v26)
    }
    v2 = v15
    switch "\n\t" {
    case "foo bar":
        var v29 float32 = 0.1
        _ = v29
    default:
        switch v2 {
        case -2147483648:
            var v30 int16 = -32768
            _ = v30
        default:
            v6 -= 0
        }
    }
    v1 -= (v1 / ((48) | 1))
    var v31 int = 32768
    v32 := GenS1{F0: (v1 & 19), F1: 65535}
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
