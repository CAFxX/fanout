package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8; F1 int }
type GenS5 struct { F0 int32; F1 int16 }
type GenS6 struct { F0 string; F1 int32; F2 int64 }
type GenS7 struct { F0 int64; F1 uint }
func genFn0() int32 {
    var v5 uint = 780
    for fi6 := 0; fi6 < 6; fi6++ {
        v7 := []int{}
        var v8 float32 = 0.0
        if (fi6 < fi6) {
            if ((fi6 & -9223372036854775808) <= fi6) {
                fmt.Println(v5, v5)
                for fi9 := 0; fi9 < 9; fi9++ {
                    v10 := GenS1{F0: -1, F1: 127, F2: 507}
                    v11 := GenS2{F0: 32767}
                    v12 := GenS3{F0: 255}
                    if (32767 < 1000) {
                        v7 = append(v7, fi6)
                    } else {
                        v13 := GenS4{F0: 204, F1: fi9}
                        fmt.Println(v8)
                        if true {
                            v14 := GenS5{F0: 0, F1: 42}
                            var v15 float32 = v8
                            if len(v7) > 0 {
                                v7[0] = (fi9 % (((^(fi6))) | 1))
                            }
                            _ = v14
                            _ = v15
                        }
                        _ = v13
                    }
                    var v16 uint32 = 1
                    v16 += v16
                    fmt.Println(fi6, v5, v16)
                    _ = fi9
                    _ = v10
                    _ = v11
                    _ = v12
                    _ = v16
                }
            }
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    var v17 uint64 = 42
    _ = v17
    _ = v5
    return -1
}
func genFn1(p0 int, p1 int) int16 {
    var v18 int16 = 1000
    v18 = v18
    v19 := []int{-9223372036854775808, (p0 & p1), p0}
    switch (-(p0)) {
    case -254:
        if false {
            v19 = append(v19, p1)
            var v20 bool = (421 > 0)
            _ = v20
        }
    case 0:
        v21 := GenS6{F0: "a", F1: 651, F2: 9223372036854775807}
        _ = v21
    default:
        for fi22 := 0; fi22 < 2; fi22++ {
            for fi23 := 0; fi23 < 10; fi23++ {
                if true {
                    if len(v19) > 0 {
                        v19[0] = p1
                    }
                    v24 := GenS1{F0: fi22, F1: 255, F2: 0}
                    _ = v24
                }
                p0 = -9223372036854775808
                _ = fi23
            }
            p1 = p1
            switch "hello" {
            case "":
                if len(v19) > 0 {
                    v19[0] = (-(p0))
                }
            case "hello":
                switch p0 {
                case 42:
                    var v25 float64 = (0.1 - 0.1)
                    _ = v25
                case 2147483647:
                    var v26 bool = (47 == -38)
                    _ = v26
                case -60:
                    switch 452 {
                    case 123:
                        _ = len(v19)
                    }
                }
            }
            _ = fi22
        }
    }
    var v27 int64 = 927
    _ = p0
    _ = p1
    _ = v18
    _ = v19
    _ = v27
    return (v18 - (-(v18)))
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int16 = 0
    var v3 uint16 = 2
    var v4 int8 = 127
    v28 := GenS7{F0: 0, F1: 240}
    var v29 uint32 = 2
    v4 -= v4
    fmt.Println(v3, v3, v3)
    if (59 <= 0) {
        v30 := []int{}
        for fi31 := 0; fi31 < 17; fi31++ {
            var v32 uint32 = (v29 - (v29 >> ((234) & 63)))
            _ = fi31
            _ = v32
        }
        var v33 int = 757
        _ = v30
        _ = v33
    }
    v34 := v1[(v29 - (^(v29)))]
    v29 += v34
    var v35 float64 = 0.0
    v35 = (-2.25 / 1.0)
    fmt.Println(v3, v34)
    for fi36 := 0; fi36 < 15; fi36++ {
        var v37 bool = (!false)
        _ = fi36
        _ = v37
    }
    switch "foo bar" {
    case "hello":
        v38 := GenS1{F0: -166, F1: 127, F2: 255}
        _ = v38
    case "foo bar":
        _ = len(v1)
    case "":
        var v39 int32 = -583
        _ = v39
    }
    if (19 >= 2147483647) {
        for fi40 := 0; fi40 < 17; fi40++ {
            v4 = v4
            delete(v1, (v34 - 1))
            _ = fi40
        }
    }
    switch ("a" + "foo bar") {
    case "foo bar":
        var v41 string = ""
        _ = v41
    default:
        fmt.Println(v3, v35, v2)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v34
    _ = v35
    _ = v4
}
