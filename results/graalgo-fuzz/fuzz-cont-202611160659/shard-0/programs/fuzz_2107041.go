package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 bool }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 int64; F1 int8; F2 uint }
func genFn0() string {
    var v4 float32 = 0.1
    v4 += (-2.25 + ((v4 + v4) + float32(float64((((v4 + (float32(((((1e10 / (-2.25 * ((float64(v4) * 1.5) - 1.5))) / 1e10) + float64(-2.25)) + 3.14159)) - 1.5)) / (v4 - (1e10 / (-2.25 * v4)))) * 3.14159)))))
    v4 += v4
    _ = v4
    return "hello"
}
func genFn1() bool {
    for fi5 := 0; fi5 < 6; fi5++ {
        var v6 int8 = -1
        fmt.Println(fi5)
        _ = fi5
        _ = v6
    }
    if false {
        v7 := []int{128, 42, 9223372036854775807}
        if false {
            var v8 uint16 = 2
            _ = v8
        }
        var v9 uint = 18446744073709551615
        _ = v7
        _ = v9
    }
    v10 := []int{-9223372036854775808, 0, 32768, 1}
    v11 := GenS1{F0: -687, F1: "", F2: (100 > 127)}
    _ = v10
    _ = v11
    return false
}

func main() {
    var v1 int64 = -572
    v2 := make(map[int]uint32)
    v3 := []uint32{947, 256, 686}
    v3 = append(v3, 2)
    switch 97 {
    case 42:
        if len(v3) > 0 {
            v3[0] = 0
        }
    default:
        v12 := GenS1{F0: 1, F1: ("a" + "a"), F2: true}
        _ = v12
    }
    if false {
        switch "hello" {
        case "hello":
            v13 := GenS1{F0: 42, F1: ("" + "a"), F2: true}
            _ = v13
        default:
            fmt.Println(v1)
        }
        var v14 uint32 = 4294967295
        _ = v14
    }
    fmt.Println(v1)
    switch "foo bar" {
    case "":
        var v15 uint = 262
        _ = v15
    }
    for fi16 := 0; fi16 < 2; fi16++ {
        v1 += -1
        _ = fi16
    }
    var v17 int8 = 1
    for fi18 := 0; fi18 < 13; fi18++ {
        fmt.Println(v1, v1, fi18)
        if (fi18 > (fi18 << ((0) & 63))) {
            for fi19 := 0; fi19 < 1; fi19++ {
                var v20 int32 = -963
                v20 = (v20 >> ((65535) & 63))
                _ = fi19
                _ = v20
            }
            v1 -= v1
            var v21 string = (("hello" + "hello") + "foo bar")
            _ = v21
        }
        _ = fi18
    }
    v1 = v1
    for fi22 := 0; fi22 < 20; fi22++ {
        switch "foo bar" {
        case "":
            for fi23 := 0; fi23 < 7; fi23++ {
                v24 := GenS2{F0: 65535, F1: 255}
                if true {
                    if false {
                        switch fi22 {
                        case -848:
                            v25 := GenS1{F0: -967, F1: "a", F2: true}
                            _ = v25
                        case -9223372036854775808:
                            v1 = v1
                        case -716:
                            v1 = (v1 << ((507) & 63))
                        }
                    }
                    v26 := GenS2{F0: 255, F1: 4294967295}
                    v3 = append(v3, 0)
                    _ = v26
                } else {
                    v27 := GenS3{F0: v1, F1: -128, F2: 0}
                    _ = v27
                }
                v17 += (^(v17))
                v1 -= v1
                if true {
                    fmt.Println(fi22, fi22)
                }
                _ = fi23
                _ = v24
            }
        case "foo bar":
            v17 -= (v17 * (-(v17)))
        }
        _ = fi22
    }
    for fi28 := 0; fi28 < 11; fi28++ {
        for fi29 := 0; fi29 < 8; fi29++ {
            _ = len(v2)
            _ = fi29
        }
        _ = fi28
    }
    switch 666 {
    case 1:
        v30 := GenS1{F0: 42, F1: ("hello" + ("" + ("" + ("a" + "")))), F2: false}
        _ = v30
    case 475:
        fmt.Println(v1)
    default:
        v31 := v2[0]
        _ = v31
    }
    fmt.Println(v17, v17)
    v1 += (-(v1))
    v32 := GenS3{F0: (v1 * (v1 % ((v1) | 1))), F1: v17, F2: 579}
    v17 = (v17 * 127)
    switch (((("foo bar" + ("" + "")) + "hello") + "foo bar") + "\n\t") {
    case "hello":
        var v33 bool = false
        _ = v33
    case "foo bar":
        switch 1 {
        case 32768:
            v17 -= -104
        }
    default:
        v34 := GenS3{F0: v1, F1: (^(v17)), F2: 741}
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v32
}
