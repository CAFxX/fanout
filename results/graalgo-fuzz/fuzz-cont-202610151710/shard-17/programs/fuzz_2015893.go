package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 int; F2 bool }
type GenS3 struct { F0 int64; F1 string }
func genFn0(p0 int16, p1 uint) int64 {
    p1 = 49
    for fi4 := 0; fi4 < 9; fi4++ {
        for fi5 := 0; fi5 < 20; fi5++ {
            v6 := []int{}
            v7 := GenS1{F0: 255}
            _ = fi5
            _ = v6
            _ = v7
        }
        _ = fi4
    }
    var v8 int32 = -268
    _ = p0
    _ = p1
    _ = v8
    return 1
}
func genFn1(p0 uint8, p1 uint64) string {
    var v9 int64 = -1
    var v10 int16 = int16(p0)
    v11 := GenS1{F0: 0}
    p1 = 18446744073709551615
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return (((("a" + "\n\t") + "\n\t") + ("" + "\n\t")) + ("foo bar" + "hello"))
}

func main() {
    var v1 uint16 = 42
    var v2 uint = uint(v1)
    var v3 uint16 = v1
    v1 = (^(v3))
    var v12 float64 = 1e10
    v1 = 65535
    var v13 int16 = -582
    v14 := GenS2{F0: v1, F1: 2147483647, F2: true}
    for fi15 := 0; fi15 < 3; fi15++ {
        for fi16 := 0; fi16 < 12; fi16++ {
            v17 := []int{}
            _ = fi16
            _ = v17
        }
        for fi18 := 0; fi18 < 20; fi18++ {
            v19 := []int{(fi18 / (((^(fi15))) | 1)), fi15, (fi15 / ((fi18) | 1)), fi15}
            var v20 uint = (v2 & 65535)
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = fi15
    }
    fmt.Println(v12)
    if false {
        var v21 int = -9223372036854775808
        var v22 uint32 = 0
        v23 := GenS3{F0: 31, F1: "a"}
        _ = v21
        _ = v22
        _ = v23
    } else {
        fmt.Println(v12, v12)
        v24 := []int{9223372036854775807, 127, 0, -2147483648}
        _ = v24
    }
    fmt.Println(v2)
    var v25 int8 = 42
    if (1 <= 479) {
        var v26 int16 = (v13 * 442)
        v13 += (v13 % (((v26 * (v26 << (((v2 * 808)) & 63)))) | 1))
        for fi27 := 0; fi27 < 4; fi27++ {
            fmt.Println(v12, v1, v12)
            for fi28 := 0; fi28 < 8; fi28++ {
                var v29 uint16 = v3
                v1 = v29
                var v30 uint32 = 65535
                _ = fi28
                _ = v29
                _ = v30
            }
            _ = fi27
        }
        _ = v26
    } else {
        var v31 float32 = 0.0
        for fi32 := 0; fi32 < 18; fi32++ {
            var v33 float64 = v12
            _ = fi32
            _ = v33
        }
        _ = v31
    }
    for fi34 := 0; fi34 < 4; fi34++ {
        v35 := []int{fi34}
        v13 = (v13 % ((v13) | 1))
        _ = fi34
        _ = v35
    }
    fmt.Println(v13, v13)
    var v36 string = ("hello" + ("" + ""))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v25
    _ = v3
    _ = v36
}
