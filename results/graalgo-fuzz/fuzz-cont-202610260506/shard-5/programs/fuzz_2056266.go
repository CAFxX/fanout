package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint64 }
type GenS2 struct { F0 int64; F1 uint64; F2 uint16 }
func genFn0(p0 int16, p1 uint32) int8 {
    fmt.Println(p1)
    var v6 int16 = (p0 % ((32767) | 1))
    _ = p0
    _ = p1
    _ = v6
    return 100
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]bool)
    v3 := make(map[string]int)
    v4 := make(map[uint32]int)
    var v5 float32 = 1e10
    var v7 string = ""
    var v8 uint16 = 65535
    for fi9 := 0; fi9 < 10; fi9++ {
        v10 := GenS1{F0: 4294967295, F1: uint64(v8), F2: 4294967295}
        if (645 <= -32768) {
            var v11 uint64 = 0
            fmt.Println(v8)
            _ = v11
        }
        for fi12 := 0; fi12 < 11; fi12++ {
            _ = len(v1)
            var v13 int16 = 32767
            _ = fi12
            _ = v13
        }
        _ = fi9
        _ = v10
    }
    v8 -= 32767
    if true {
        switch v7 {
        case "hello":
            v14 := GenS1{F0: 4294967295, F1: 2, F2: 1}
            _ = v14
        case "a":
            v8 = (^(v8))
        case "\n\t":
            if (623 <= 42) {
                fmt.Println(v8, v5, v7)
            }
        default:
            v8 -= v8
        }
        if (-1 == 42) {
            for fi15 := 0; fi15 < 6; fi15++ {
                fmt.Println(v8, v5, v8)
                v4[833] = 0
                _ = fi15
            }
            v8 = v8
        }
        var v16 uint32 = 42
        _ = v16
    } else {
        switch 32768 {
        case 32768:
            var v17 uint32 = 600
            _ = v17
        case 2147483647:
            var v18 uint16 = 997
            _ = v18
        case -979:
            v7 = ((((v7 + v7) + "foo bar") + v7) + ("foo bar" + (v7 + "\n\t")))
        }
    }
    v19 := []int{}
    v7 = v7
    switch -9223372036854775808 {
    case 496:
        v20 := GenS2{F0: 1, F1: 0, F2: v8}
        _ = v20
    case -9223372036854775808:
        for fi21 := 0; fi21 < 12; fi21++ {
            switch ((v7 + v7) + "") {
            case "\n\t":
                _ = len(v4)
            case "hello":
                var v22 int8 = 1
                _ = v22
            case "a":
                var v23 string = "hello"
                _ = v23
            default:
                var v24 int64 = -956
                _ = v24
            }
            fmt.Println(v5)
            _ = fi21
        }
    case 4294967295:
        v7 = v7
    default:
        delete(v1, 65535)
    }
    fmt.Println(v5, v7, v5)
    if (!false) {
        v25 := GenS1{F0: 2, F1: 314, F2: 2}
        switch 42 {
        case -9223372036854775808:
            var v26 int16 = 598
            _ = v26
        case 821:
            for fi27 := 0; fi27 < 13; fi27++ {
                var v28 uint64 = 2
                _ = fi27
                _ = v28
            }
        }
        for fi29 := 0; fi29 < 10; fi29++ {
            if len(v19) > 0 {
                v19[0] = fi29
            }
            switch v7 {
            case "foo bar":
                var v30 float64 = 0.0
                _ = v30
            case "hello":
                v7 = v7
            case "\n\t":
                if false {
                    var v31 float64 = 0.1
                    _ = v31
                }
            }
            _ = fi29
        }
        _ = v25
    } else {
        for fi32 := 0; fi32 < 20; fi32++ {
            for fi33 := 0; fi33 < 2; fi33++ {
                var v34 bool = (!false)
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
                var v35 uint = 4294967295
                _ = fi33
                _ = v34
                _ = v35
            }
            _ = fi32
        }
        var v36 uint64 = 18446744073709551615
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
