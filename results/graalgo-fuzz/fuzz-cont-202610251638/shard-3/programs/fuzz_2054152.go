package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint }
func genFn0() int64 {
    var v7 bool = (((-2147483648 > -1) && true) && (("hello" + "a") <= "\n\t"))
    var v8 uint = 0
    v9 := GenS1{F0: 0, F1: false, F2: 1}
    _ = v7
    _ = v8
    _ = v9
    return 9223372036854775807
}

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[int]uint32)
    v3 := make(map[int]bool)
    var v4 uint32 = 575
    v5 := []int{32768}
    var v6 string = "\n\t"
    v10 := v3[-9223372036854775808]
    for fi11 := 0; fi11 < 10; fi11++ {
        switch "a" {
        case "foo bar":
            if len(v5) > 0 {
                v5[0] = -600
            }
        }
        _ = fi11
    }
    switch 9223372036854775807 {
    case 32768:
        fmt.Println(v6, v4, v1)
    case 1:
        if v10 {
            v10 = v10
        } else {
            var v12 int = 1
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = v12
        }
    }
    if v10 {
        for fi13 := 0; fi13 < 7; fi13++ {
            v14 := GenS1{F0: 0, F1: v10, F2: 105}
            _ = fi13
            _ = v14
        }
        v4 += v4
    }
    var v15 bool = v10
    switch 1 {
    case 42:
        var v16 uint64 = 1
        _ = v16
    case 1:
        if (0 > 846) {
            if v10 {
                _ = len(v2)
            }
        }
    case -9223372036854775808:
        fmt.Println(v6, v10)
    }
    var v17 int8 = 42
    var v18 uint32 = v4
    switch 852 {
    case 0:
        switch -2147483648 {
        case 456:
            v15 = (v15 && true)
        }
    default:
        if v10 {
            fmt.Println(v4, v10, v10)
            fmt.Println(v4)
            fmt.Println(v4)
        }
    }
    var v19 uint64 = 18446744073709551615
    v19 += (v19 >> ((42) & 63))
    switch 32768 {
    case 4294967295:
        var v20 int = -2147483648
        _ = v20
    case -1:
        if ((!((uint(v19) > 71) && true)) && false) {
            fmt.Println(v18, v19, v10)
            if v15 {
                for fi21 := 0; fi21 < 13; fi21++ {
                    for fi22 := 0; fi22 < 13; fi22++ {
                        fmt.Println(fi22)
                        _ = fi22
                    }
                    v6 = "foo bar"
                    v18 = v18
                    _ = fi21
                }
            }
        } else {
            v6 = v6
            fmt.Println(v10, v4, v6)
            for fi23 := 0; fi23 < 18; fi23++ {
                v24 := GenS2{F0: (v6 + v6)}
                _ = fi23
                _ = v24
            }
        }
    case 1:
        if false {
            var v25 float64 = v1
            _ = v25
        }
    }
    if false {
        v26 := GenS3{F0: 903}
        for fi27 := 0; fi27 < 17; fi27++ {
            fmt.Println(v15, fi27)
            var v28 int32 = -292
            for fi29 := 0; fi29 < 14; fi29++ {
                delete(v3, (fi27 & fi29))
                _ = fi29
            }
            _ = fi27
            _ = v28
        }
        _ = v26
    }
    var v30 float32 = float32(1.5)
    var v31 uint16 = 32767
    var v32 uint8 = 1
    _ = len(v3)
    v2[0] = (v18 % ((v4) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
