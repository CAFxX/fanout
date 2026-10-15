package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8; F1 uint32 }
func genFn0(p0 string, p1 uint8, p2 bool) int {
    v7 := []int{989, 42, -1}
    switch -526 {
    case 128:
        v8 := GenS1{F0: 4294967295}
        _ = v8
    case -683:
        v7 = append(v7, 42)
    case 9223372036854775807:
        p2 = (p1 > 157)
    default:
        for fi9 := 0; fi9 < 12; fi9++ {
            var v10 int = 413
            for fi11 := 0; fi11 < 18; fi11++ {
                var v12 int8 = -128
                v13 := GenS2{F0: -32768}
                var v14 int64 = -1
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
            }
            for fi15 := 0; fi15 < 16; fi15++ {
                var v16 int16 = -660
                if p2 {
                    fmt.Println(p0, p0)
                    p1 -= p1
                } else {
                    var v17 string = p0
                    var v18 bool = false
                    var v19 uint16 = 2
                    _ = v17
                    _ = v18
                    _ = v19
                }
                _ = fi15
                _ = v16
            }
            _ = fi9
            _ = v10
        }
    }
    var v20 float64 = 3.14159
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v7
    return -2147483648
}

func main() {
    var v1 int32 = -717
    var v2 float32 = ((1.5 - (-2.25 / 1.5)) / 1.5)
    var v3 int32 = (v1 << ((2) & 63))
    v4 := make(map[uint32]bool)
    var v5 string = ""
    v6 := make(map[int]uint32)
    if true {
        for fi21 := 0; fi21 < 15; fi21++ {
            fmt.Println(v2, v2, v2)
            switch v5 {
            case "foo bar":
                v22 := []int{(fi21 / ((260) | 1))}
                _ = v22
            default:
                v5 = "\n\t"
            }
            _ = fi21
        }
        var v23 uint8 = 1
        v24 := []int{}
        _ = v23
        _ = v24
    } else {
        v2 = float32((0.0 - ((3.14159 + -2.25) / 0.1)))
    }
    v3 += (v3 + (v3 | v3))
    var v25 int64 = 42
    switch (v5 + (v5 + (v5 + ((v5 + ((v5 + "\n\t") + ("" + v5))) + (("hello" + v5) + v5))))) {
    case "\n\t":
        _ = len(v4)
    case "":
        switch 734 {
        case 42:
            switch v5 {
            case "a":
                v26 := v4[42]
                _ = v26
            case "hello":
                v3 = (v1 - (v3 - v3))
            }
        case 32767:
            if (!true) {
                for fi27 := 0; fi27 < 15; fi27++ {
                    fmt.Println(v1, fi27)
                    var v28 uint = 65535
                    _ = fi27
                    _ = v28
                }
                v29 := GenS3{F0: 0, F1: 551}
                _ = v29
            }
        case 128:
            for fi30 := 0; fi30 < 14; fi30++ {
                for fi31 := 0; fi31 < 17; fi31++ {
                    var v32 int = genFn0(v5, 1, false)
                    var v33 int16 = -32768
                    _ = fi31
                    _ = v32
                    _ = v33
                }
                _ = fi30
            }
        default:
            v34 := []int{}
            _ = v34
        }
    default:
        v35 := GenS3{F0: 1, F1: 1}
        _ = v35
    }
    for fi36 := 0; fi36 < 20; fi36++ {
        fmt.Println(v25, v3, v2)
        var v37 float32 = 1e10
        _ = fi36
        _ = v37
    }
    for fi38 := 0; fi38 < 12; fi38++ {
        var v39 float32 = v2
        _ = fi38
        _ = v39
    }
    fmt.Println(v2, v1)
    v3 = (v3 ^ v3)
    var v40 int32 = v1
    v41 := GenS1{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
