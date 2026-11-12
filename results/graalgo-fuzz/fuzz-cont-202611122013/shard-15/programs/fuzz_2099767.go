package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 uint8 }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int16; F1 int8 }
func genFn0() uint64 {
    var v6 float64 = (((-2.25 / 1e10) - 1e10) / 1.5)
    v7 := []int{128, 0, 9223372036854775807}
    v6 = v6
    var v8 int8 = 127
    if false {
        if len(v7) > 0 {
            v7[0] = 32768
        }
        v8 = v8
        if len(v7) > 0 {
            v7[0] = 128
        }
    } else {
        v6 -= (v6 / float64(1.5))
        for fi9 := 0; fi9 < 20; fi9++ {
            v6 -= v6
            fmt.Println(v6, fi9, v6)
            var v10 int16 = -32768
            _ = fi9
            _ = v10
        }
        v11 := GenS1{F0: "", F1: 9223372036854775807, F2: 102}
        _ = v11
    }
    _ = v6
    _ = v7
    _ = v8
    return 826
}

func main() {
    v1 := make(map[int]int)
    var v2 string = "hello"
    var v3 uint = 2
    var v4 float64 = 0.1
    var v5 string = v2
    for fi12 := 0; fi12 < 14; fi12++ {
        v13 := []int{fi12, (fi12 << ((18446744073709551615) & 63)), 824}
        for fi14 := 0; fi14 < 6; fi14++ {
            v2 = v2
            v15 := GenS2{F0: 0, F1: 18446744073709551615}
            fmt.Println(v3, fi12, v3)
            var v16 float32 = float32(-2.25)
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi12
        _ = v13
    }
    v17 := GenS3{F0: 32767}
    var v18 uint16 = 255
    v19 := []int{4294967295, -9223372036854775808}
    v2 = v2
    if (32767 < 770) {
        v20 := GenS4{F0: -544}
        _ = v20
    }
    v18 = (v18 << (((v3 % ((1) | 1))) & 63))
    var v21 int64 = 42
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    v22 := v1[4294967295]
    fmt.Println(v2, v3, v21)
    for fi23 := 0; fi23 < 18; fi23++ {
        var v24 uint64 = 446
        var v25 int8 = 0
        _ = fi23
        _ = v24
        _ = v25
    }
    switch "foo bar" {
    case "a":
        if true {
            v26 := GenS5{F0: 1000, F1: 127}
            _ = v26
        } else {
            switch v22 {
            case 9223372036854775807:
                if true {
                    for fi27 := 0; fi27 < 1; fi27++ {
                        for fi28 := 0; fi28 < 6; fi28++ {
                            var v29 int32 = -1
                            _ = fi28
                            _ = v29
                        }
                        var v30 uint64 = 0
                        _ = fi27
                        _ = v30
                    }
                    v2 = v2
                }
            case 1:
                v3 += v3
            }
            if true {
                var v31 uint32 = 796
                fmt.Println(v22, v5)
                _ = v31
            }
        }
    }
    var v32 float32 = (1e10 + 0.1)
    for fi33 := 0; fi33 < 10; fi33++ {
        var v34 int = (v22 << (((v3 + v3)) & 63))
        _ = fi33
        _ = v34
    }
    v1[(v22 - v22)] = (v22 | 4294967295)
    var v35 uint16 = uint16(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v32
    _ = v35
    _ = v4
    _ = v5
}
