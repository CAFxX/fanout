package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 int64; F1 int; F2 uint64 }
type GenS4 struct { F0 int16; F1 string; F2 int64 }
func genFn0() int16 {
    var v7 bool = true
    if (!(42 <= -58)) {
        if (!v7) {
            var v8 int16 = -1
            v7 = v7
            _ = v8
        } else {
            v7 = v7
            fmt.Println(v7)
        }
    } else {
        v9 := GenS1{F0: 1}
        v7 = (false || v7)
        _ = v9
    }
    v10 := GenS2{F0: -1, F1: -1}
    var v11 float64 = (0.0 - (1.5 * ((float64(1.5) - (0.1 * 1e10)) + 0.1)))
    _ = v10
    _ = v11
    _ = v7
    return 0
}

func main() {
    var v1 uint8 = 207
    var v2 uint32 = 130
    var v3 bool = false
    var v4 int16 = -32768
    v5 := make(map[uint32]bool)
    v6 := make(map[string]int)
    var v12 int8 = 127
    var v13 uint8 = 71
    fmt.Println(v4, v12)
    v4 = genFn0()
    if true {
        var v14 string = "\n\t"
        v15 := GenS2{F0: 2147483647, F1: -1}
        v2 -= 2147483648
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        fmt.Println(v3, v12, v2)
        switch "\n\t" {
        case "foo bar":
            fmt.Println(v3, v2)
        case "hello":
            v17 := []int{fi16, fi16, int(v2), fi16}
            _ = v17
        case "\n\t":
            var v18 bool = v3
            _ = v18
        }
        for fi19 := 0; fi19 < 6; fi19++ {
            if true {
                var v20 int = (fi16 << ((2) & 63))
                _ = v20
            } else {
                v21 := []int{fi16}
                _ = len(v21)
                if v3 {
                    v1 += 0
                } else {
                    var v22 bool = (!(!true))
                    var v23 uint = 42
                    v24 := GenS1{F0: 790}
                    v21 = append(v21, fi16)
                    v4 += v4
                    _ = v22
                    _ = v23
                    _ = v24
                }
                _ = v21
            }
            fmt.Println(v1, fi16)
            var v25 uint = 2
            _ = fi19
            _ = v25
        }
        _ = fi16
    }
    var v26 string = "hello"
    v3 = v3
    v27 := GenS3{F0: 689, F1: 0, F2: 451}
    switch v26 {
    case "hello":
        delete(v6, (v26 + (v26 + v26)))
    case "foo bar":
        v3 = (!false)
    default:
        v4 = genFn0()
    }
    fmt.Println(v13)
    v12 = 127
    v28 := GenS2{F0: -2147483648, F1: 9223372036854775807}
    switch v26 {
    case "\n\t":
        v13 = v1
    case "a":
        if (uint64(v13) < 80) {
            v29 := GenS3{F0: -1, F1: 127, F2: 29}
            _ = v29
        }
    case "":
        fmt.Println(v3, v3)
    default:
        for fi30 := 0; fi30 < 3; fi30++ {
            var v31 uint32 = (^(v2))
            v32 := GenS4{F0: (v4 << ((149) & 63)), F1: v26, F2: -654}
            _ = fi30
            _ = v31
            _ = v32
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
