package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64; F1 int; F2 uint8 }
type GenS4 struct { F0 int32; F1 int32; F2 string }
func genFn0() string {
    var v5 float32 = (1.5 + (((3.14159 / -2.25) * -2.25) - 1e10))
    switch "\n\t" {
    case "a":
        v5 = v5
    }
    v6 := []int{-1, 1, 42, -9223372036854775808}
    _ = v5
    _ = v6
    return "hello"
}

func main() {
    v1 := make(map[int]uint32)
    var v2 string = "hello"
    v3 := []string{v2, v2, ("foo bar" + ""), v2}
    v4 := []string{v2, "foo bar", ((v2 + v2) + "a")}
    var v7 int16 = -494
    var v8 int8 = 42
    if false {
        var v9 float32 = 1e10
        v10 := GenS1{F0: int32(v7)}
        _ = v10
        _ = v9
    } else {
        switch -937 {
        case 376:
            v7 += 121
        }
        if (42 != 1) {
            var v11 int = -643
            _ = v11
        }
        v7 -= -32768
    }
    if (255 > 128) {
        if false {
            v2 = ("foo bar" + "\n\t")
        } else {
            var v12 int16 = -619
            fmt.Println(v2, v7)
            var v13 int16 = -454
            _ = v12
            _ = v13
        }
    } else {
        var v14 uint8 = 128
        _ = v14
    }
    v15 := GenS2{F0: v2}
    for fi16 := 0; fi16 < 14; fi16++ {
        v7 = v7
        v17 := v1[(fi16 * (fi16 << ((303) & 63)))]
        _ = fi16
        _ = v17
    }
    if (true && false) {
        v2 = v2
        var v18 uint64 = 247
        for fi19 := 0; fi19 < 7; fi19++ {
            delete(v1, -216)
            var v20 int8 = v8
            _ = fi19
            _ = v20
        }
        _ = v18
    }
    var v21 int64 = -1
    if (2 != 4294967295) {
        if (65535 < 65535) {
            v22 := GenS2{F0: v2}
            var v23 uint8 = 2
            fmt.Println(v7)
            _ = v22
            _ = v23
        }
        v4 = append(v4, (("foo bar" + ((((v2 + v2) + v2) + v2) + v2)) + (v2 + "hello")))
    }
    var v24 float64 = ((float64(3.14159) + float64(1.5)) + float64((((0.0 * 1e10) / 3.14159) * 0.0)))
    v25 := GenS3{F0: (v21 & v21), F1: 2147483647, F2: 0}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v21 = v21
    v21 = 9223372036854775807
    v26 := GenS1{F0: 0}
    var v27 float64 = v24
    v28 := GenS4{F0: 2147483647, F1: 2147483647, F2: ((v2 + v2) + (v2 + (v2 + v2)))}
    v21 += (v21 % (((v21 / ((42) | 1))) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
