package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32; F1 int32 }
func genFn0(p0 uint16) int16 {
    var v7 uint16 = 2
    if false {
        fmt.Println(p0, p0)
    }
    v8 := []int{2147483647}
    if false {
        switch -913 {
        case 2147483647:
            switch "hello" {
            case "\n\t":
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
            default:
                fmt.Println(p0, v7)
            }
        case 32767:
            for fi9 := 0; fi9 < 3; fi9++ {
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                _ = fi9
            }
        case 0:
            var v10 uint16 = 32767
            _ = v10
        }
    } else {
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
    }
    _ = p0
    _ = v7
    _ = v8
    return -1
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int64 = 1
    v3 := make(map[string]bool)
    var v4 uint64 = 0
    v5 := make(map[uint32]string)
    var v6 int = 979
    switch "\n\t" {
    case "foo bar":
        if (!(!true)) {
            var v11 int16 = -691
            v12 := []int{}
            switch 226 {
            case -9223372036854775808:
                var v13 int = 4294967295
                _ = v13
            case 32768:
                var v14 bool = (true || true)
                _ = v14
            }
            _ = v11
            _ = v12
        }
    }
    var v15 int32 = -2147483648
    for fi16 := 0; fi16 < 2; fi16++ {
        for fi17 := 0; fi17 < 4; fi17++ {
            v18 := GenS1{F0: 2147483647}
            v4 = 2
            _ = fi17
            _ = v18
        }
        v19 := []int{32768}
        var v20 float64 = ((1.5 - 0.1) * (1.5 - (float64(3.14159) - 3.14159)))
        _ = fi16
        _ = v19
        _ = v20
    }
    switch v6 {
    case 42:
        for fi21 := 0; fi21 < 11; fi21++ {
            if ((!true) || (true || (498 < 2147483648))) {
                switch v6 {
                case -255:
                    v22 := []int{(fi21 + fi21)}
                    _ = v22
                case 774:
                    switch "hello" {
                    case "foo bar":
                        v23 := GenS2{F0: 0, F1: v15}
                        _ = v23
                    }
                default:
                    fmt.Println(v6, v15, fi21)
                }
                v24 := []int{}
                var v25 float32 = -2.25
                _ = v24
                _ = v25
            }
            v15 += v15
            _ = fi21
        }
    }
    var v26 float32 = 0.0
    v26 -= 3.14159
    fmt.Println(v4, v15)
    var v27 float32 = float32(0.0)
    v15 -= (v15 + -145)
    v28 := []int{v6}
    if (-1 <= -32768) {
        var v29 string = "hello"
        v27 -= 1e10
        _ = v29
    } else {
        fmt.Println(v2, v4)
        v2 += v2
        var v30 int64 = 42
        _ = v30
    }
    fmt.Println(v4, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
