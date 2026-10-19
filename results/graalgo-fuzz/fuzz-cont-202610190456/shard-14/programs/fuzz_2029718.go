package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 string) int8 {
    if false {
        v7 := []int{4294967295}
        var v8 float64 = (float64(3.14159) + (-2.25 + -2.25))
        _ = v7
        _ = v8
    }
    var v9 float32 = 0.1
    p0 = (("\n\t" + "a") + ("foo bar" + p0))
    v9 = float32(1e10)
    _ = p0
    _ = v9
    return 116
}
func genFn1(p0 int64) int16 {
    v10 := GenS1{F0: 2}
    var v11 float32 = 1e10
    v12 := GenS1{F0: 42}
    if false {
        var v13 string = ("\n\t" + (("hello" + "a") + "foo bar"))
        v14 := GenS1{F0: 1}
        _ = v13
        _ = v14
    }
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    return 325
}

func main() {
    v1 := make(map[int]int)
    var v2 float64 = (1.5 / (0.1 + ((3.14159 / 1.5) * 0.1)))
    v3 := make(map[int]string)
    var v4 int16 = 24
    var v5 bool = false
    var v6 float32 = 0.0
    switch "foo bar" {
    case "hello":
        for fi15 := 0; fi15 < 7; fi15++ {
            v1[fi15] = (fi15 - (-(fi15)))
            _ = fi15
        }
    }
    v2 += (v2 / v2)
    fmt.Println(v5, v6)
    if v5 {
        var v16 uint64 = 0
        if v5 {
            v17 := GenS1{F0: 65535}
            var v18 string = "foo bar"
            var v19 int16 = (v4 << ((42) & 63))
            _ = v17
            _ = v18
            _ = v19
        } else {
            fmt.Println(v16, v6)
            var v20 int16 = (v4 << ((967) & 63))
            _ = v20
        }
        for fi21 := 0; fi21 < 13; fi21++ {
            for fi22 := 0; fi22 < 9; fi22++ {
                v23 := []int{-2147483648, -41}
                v5 = v5
                v24 := GenS1{F0: 255}
                _ = fi22
                _ = v23
                _ = v24
            }
            v25 := GenS2{F0: 42}
            _ = fi21
            _ = v25
        }
        _ = v16
    }
    var v26 int32 = -1
    v2 = float64((-2.25 - v6))
    v26 -= -2147483648
    switch 127 {
    case 42:
        switch -95 {
        case -1:
            v4 = 1000
        default:
            var v27 int32 = (-(v26))
            _ = v27
        }
    case 39:
        v28 := []int{-1, 42, -9223372036854775808, -453}
        _ = v28
    case 4294967295:
        v3[-9223372036854775808] = ("foo bar" + "hello")
    }
    for fi29 := 0; fi29 < 19; fi29++ {
        for fi30 := 0; fi30 < 5; fi30++ {
            v26 -= 42
            _ = fi30
        }
        v26 -= int32(fi29)
        _ = fi29
    }
    fmt.Println(v6, v2)
    var v31 float64 = ((float64(-2.25) / 0.1) + -2.25)
    for fi32 := 0; fi32 < 7; fi32++ {
        var v33 int = fi32
        v34 := []int{-2147483648, (^(v33)), 386, fi32}
        fmt.Println(v33, v2, v33)
        _ = fi32
        _ = v33
        _ = v34
    }
    var v35 int8 = -1
    var v36 bool = ((!(0 < 18446744073709551615)) || (!(!(65535 != 2))))
    v31 += v2
    var v37 bool = (-594 <= 581)
    var v38 bool = v37
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v31
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
