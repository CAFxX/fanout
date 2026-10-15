package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int }
func genFn0(p0 uint64, p1 uint8, p2 uint64) bool {
    v6 := []int{-787, 1, 4294967295}
    var v7 uint = 18446744073709551615
    v7 = v7
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return false
}
func genFn1(p0 uint16, p1 bool, p2 int32) bool {
    if p1 {
        fmt.Println(p2, p1)
        fmt.Println(p0, p0)
        var v8 float64 = (0.0 + (1e10 * ((3.14159 / ((0.1 * -2.25) - (1e10 + 3.14159))) / (0.0 - ((((float64(1e10) / -2.25) * ((1.5 * 1e10) / float64(0.1))) - 1e10) + 0.1)))))
        _ = v8
    } else {
        for fi9 := 0; fi9 < 2; fi9++ {
            fmt.Println(p0)
            _ = fi9
        }
        v10 := GenS1{F0: -32768, F1: 82}
        for fi11 := 0; fi11 < 5; fi11++ {
            p1 = p1
            v12 := []int{32767, fi11}
            _ = fi11
            _ = v12
        }
        _ = v10
    }
    var v13 bool = (!p1)
    v14 := GenS1{F0: 182, F1: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return (false && (1 > 18446744073709551615))
}

func main() {
    var v1 bool = (4294967295 > 65535)
    var v2 int16 = -1
    var v3 string = "\n\t"
    var v4 float32 = -2.25
    var v5 int = -354
    for fi15 := 0; fi15 < 3; fi15++ {
        for fi16 := 0; fi16 < 4; fi16++ {
            var v17 int = 808
            var v18 bool = v1
            v19 := []int{}
            var v20 int = fi16
            v20 += (fi15 | 42)
            _ = fi16
            _ = v17
            _ = v18
            _ = v19
            _ = v20
        }
        v1 = (0 != 0)
        _ = fi15
    }
    v2 = v2
    for fi21 := 0; fi21 < 1; fi21++ {
        if v1 {
            v22 := []int{2147483647}
            v1 = (255 >= 885)
            _ = v22
        } else {
            var v23 int8 = -127
            v23 = (v23 * 0)
            _ = v23
        }
        v5 += (v5 | (fi21 & (fi21 & (v5 | v5))))
        for fi24 := 0; fi24 < 18; fi24++ {
            v4 = -2.25
            v4 = 0.0
            _ = fi24
        }
        _ = fi21
    }
    var v25 float32 = v4
    if (513 != 493) {
        v26 := GenS2{F0: (v5 * v5)}
        _ = v26
    } else {
        v1 = v1
    }
    var v27 int16 = 42
    for fi28 := 0; fi28 < 20; fi28++ {
        fmt.Println(fi28, v3, v3)
        fmt.Println(v25, v5)
        v3 = "foo bar"
        _ = fi28
    }
    var v29 string = "foo bar"
    v30 := GenS1{F0: (v2 | v27), F1: 128}
    switch v5 {
    case 1:
        v31 := GenS3{F0: v5}
        _ = v31
    case -1:
        fmt.Println(v4, v3, v4)
    default:
        if (127 < 119) {
            switch 380 {
            case 611:
                v32 := []int{v5, (v5 * (^(v5))), (v5 % ((v5) | 1))}
                _ = v32
            default:
                fmt.Println(v29, v29, v3)
            }
        }
    }
    for fi33 := 0; fi33 < 4; fi33++ {
        var v34 int16 = v2
        _ = fi33
        _ = v34
    }
    fmt.Println(v25)
    v35 := []int{}
    if len(v35) > 0 {
        v35[0] = v5
    }
    if len(v35) >= 2 {
        v35 = v35[1:len(v35)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v35
    _ = v4
    _ = v5
}
