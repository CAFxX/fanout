package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint32; F1 int8; F2 int64 }
type GenS3 struct { F0 int16; F1 int16; F2 int16 }
func genFn0() int {
    var v7 uint64 = 18446744073709551615
    var v8 int = 32768
    v7 -= 185
    if false {
        if (253 != 14) {
            var v9 int8 = 1
            _ = v9
        }
    }
    _ = v7
    _ = v8
    return (v8 * 2147483647)
}
func genFn1(p0 int64, p1 uint64) uint8 {
    p1 -= 2
    switch (((("hello" + "a") + "") + "\n\t") + "") {
    case "hello":
        fmt.Println(p0)
    case "\n\t":
        v10 := []int{}
        _ = v10
    case "a":
        for fi11 := 0; fi11 < 9; fi11++ {
            var v12 int32 = 0
            var v13 string = "hello"
            v14 := GenS1{F0: 255, F1: -9223372036854775808}
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
    default:
        for fi15 := 0; fi15 < 8; fi15++ {
            p0 = (p0 >> ((925) & 63))
            p1 = p1
            var v16 uint16 = 255
            _ = fi15
            _ = v16
        }
    }
    v17 := GenS2{F0: 134, F1: -77, F2: (p0 + 42)}
    v18 := []int{}
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    return 74
}

func main() {
    var v1 int64 = 1
    var v2 int8 = 1
    var v3 float32 = 1.5
    var v4 int8 = 127
    v5 := make(map[uint32]string)
    v6 := make(map[int]bool)
    v2 = 1
    for fi19 := 0; fi19 < 2; fi19++ {
        v20 := GenS3{F0: 1, F1: 1, F2: 32767}
        for fi21 := 0; fi21 < 20; fi21++ {
            v1 = v1
            _ = fi21
        }
        var v22 uint = 74
        _ = fi19
        _ = v20
        _ = v22
    }
    var v23 uint = 737
    v23 -= v23
    _ = len(v6)
    switch "" {
    case "":
        v6[32767] = false
    case "foo bar":
        v1 = (v1 ^ (v1 | (v1 / (((v1 | v1)) | 1))))
    default:
        var v24 uint32 = 227
        _ = v24
    }
    v25 := v6[-474]
    v26 := GenS3{F0: -544, F1: 32767, F2: 571}
    fmt.Println(v23, v25)
    for fi27 := 0; fi27 < 5; fi27++ {
        _ = len(v5)
        var v28 uint = v23
        v28 += (^(v23))
        _ = fi27
        _ = v28
    }
    v29 := GenS2{F0: 4294967295, F1: v4, F2: 317}
    v30 := v6[genFn0()]
    fmt.Println(v2, v4, v3)
    for fi31 := 0; fi31 < 9; fi31++ {
        var v32 uint = (v23 + v23)
        for fi33 := 0; fi33 < 17; fi33++ {
            v30 = v25
            _ = fi33
        }
        var v34 uint64 = 4294967295
        _ = fi31
        _ = v32
        _ = v34
    }
    var v35 uint64 = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
