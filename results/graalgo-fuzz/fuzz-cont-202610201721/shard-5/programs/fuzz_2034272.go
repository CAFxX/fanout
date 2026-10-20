package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 int32; F2 uint16 }
func genFn0(p0 int8, p1 int64) int {
    var v6 int32 = 42
    var v7 uint64 = 1
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 32768
}
func genFn1(p0 string, p1 uint8) int16 {
    v8 := GenS1{F0: p0}
    if false {
        v9 := GenS1{F0: "a"}
        p1 = 1
        p0 = p0
        _ = v9
    }
    if false {
        v10 := []int{}
        _ = v10
    }
    v11 := GenS1{F0: "hello"}
    var v12 uint16 = 42
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v8
    return -510
}

func main() {
    var v1 bool = false
    var v2 int16 = -1
    v3 := make(map[uint32]uint32)
    var v4 int32 = -338
    v5 := make(map[string]string)
    v1 = (32768 <= genFn0(100, -9223372036854775808))
    v2 -= v2
    var v13 uint32 = 2
    v14 := GenS2{F0: (v2 & (v2 | (v2 | -1)))}
    var v15 uint64 = 623
    for fi16 := 0; fi16 < 20; fi16++ {
        v17 := GenS3{F0: 2, F1: -2147483648, F2: 32767}
        _ = fi16
        _ = v17
    }
    var v18 uint16 = uint16(v15)
    v19 := GenS3{F0: (v18 % ((589) | 1)), F1: (^(v4)), F2: (^(v18))}
    v18 = 65535
    var v20 float32 = 0.1
    if (("\n\t" + ((("foo bar" + "foo bar") + "a") + "foo bar")) <= "\n\t") {
        fmt.Println(v13, v1, v4)
        var v21 bool = false
        _ = v21
    }
    fmt.Println(v13, v18)
    if true {
        var v22 int32 = (v4 ^ -183)
        var v23 int8 = 100
        var v24 int8 = v23
        _ = v22
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 17; fi25++ {
        fmt.Println(v15, fi25, v13)
        var v26 uint64 = v15
        _ = fi25
        _ = v26
    }
    fmt.Println(v4)
    v20 += 0.1
    var v27 int = 9223372036854775807
    var v28 float64 = ((float64(v20) / 1.5) * float64(v20))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
