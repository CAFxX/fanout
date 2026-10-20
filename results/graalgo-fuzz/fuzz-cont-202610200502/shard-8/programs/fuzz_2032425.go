package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 int8; F2 uint64 }
type GenS4 struct { F0 int64; F1 int }
type GenS5 struct { F0 int8; F1 uint8; F2 int64 }
type GenS6 struct { F0 int16 }

func main() {
    var v1 float32 = 0.1
    v2 := make(map[string]int)
    v3 := make(map[uint32]int)
    v4 := make(map[int]uint32)
    var v5 string = "\n\t"
    v6 := []int{-794}
    var v7 string = v5
    v8 := GenS1{F0: 65535, F1: (((33 < -1) || false) || (-1 > -128))}
    switch 128 {
    case 127:
        v5 = ""
    case -9223372036854775808:
        if false {
            v7 = (((v7 + "a") + v7) + "\n\t")
            fmt.Println(v1, v7)
        }
    case 42:
        v9 := GenS1{F0: 255, F1: false}
        _ = v9
    default:
        var v10 uint8 = 128
        _ = v10
    }
    if (!true) {
        v11 := GenS2{F0: 2147483647}
        v12 := GenS1{F0: 255, F1: (535 == 801)}
        if (!false) {
            fmt.Println(v7, v5, v1)
            delete(v3, 870)
            if true {
                v6 = append(v6, -301)
                var v13 int64 = 42
                _ = v13
            } else {
                var v14 float32 = 1.5
                var v15 string = "hello"
                v5 = v7
                _ = v14
                _ = v15
            }
        } else {
            v1 = ((float32(1e10) + float32(1e10)) * v1)
            var v16 int8 = -128
            if (1 >= 1) {
                v2[(v7 + ("a" + v5))] = -351
            }
            _ = v16
        }
        _ = v11
        _ = v12
    }
    fmt.Println(v1, v5, v7)
    v5 = v7
    v17 := GenS1{F0: 65535, F1: false}
    if (!(!false)) {
        if true {
            if ((("foo bar" + v7) + (v7 + ("a" + ""))) < (((("foo bar" + "foo bar") + v7) + (v5 + v5)) + (v7 + ""))) {
                for fi18 := 0; fi18 < 11; fi18++ {
                    v5 = v5
                    _ = fi18
                }
                for fi19 := 0; fi19 < 8; fi19++ {
                    v20 := GenS2{F0: 65535}
                    v21 := GenS1{F0: 2, F1: (42 <= 2)}
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            }
        } else {
            for fi22 := 0; fi22 < 14; fi22++ {
                var v23 uint16 = 0
                _ = fi22
                _ = v23
            }
            v7 = (v7 + v7)
            var v24 string = v5
            _ = v24
        }
        v25 := GenS1{F0: 4294967295, F1: true}
        _ = v25
    }
    v6 = append(v6, 42)
    v26 := GenS3{F0: v5, F1: 100, F2: 18446744073709551615}
    v27 := GenS3{F0: v5, F1: 42, F2: 4294967295}
    v28 := GenS4{F0: -858, F1: 9223372036854775807}
    v29 := GenS5{F0: 42, F1: 127, F2: 1}
    for fi30 := 0; fi30 < 16; fi30++ {
        v31 := GenS5{F0: 127, F1: 127, F2: 9223372036854775807}
        if false {
            var v32 bool = false
            _ = v32
        } else {
            for fi33 := 0; fi33 < 13; fi33++ {
                if false {
                    if (true || (!false)) {
                        delete(v4, 128)
                        v5 = v7
                    }
                    fmt.Println(fi33, v1, fi33)
                    var v34 string = v7
                    _ = v34
                } else {
                    if (42 > 255) {
                        if false {
                            v35 := v2[((("a" + v5) + (v5 + ("foo bar" + v7))) + v7)]
                            v36 := GenS4{F0: 1, F1: fi33}
                            if (!false) {
                                fmt.Println(fi33, v1)
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                v37 := GenS1{F0: 190, F1: true}
                                _ = v37
                            }
                            _ = v35
                            _ = v36
                        }
                    }
                    v38 := GenS6{F0: 1}
                    v39 := GenS2{F0: 1}
                    _ = v38
                    _ = v39
                }
                var v40 int32 = -1
                var v41 float64 = 3.14159
                _ = fi33
                _ = v40
                _ = v41
            }
            v42 := GenS2{F0: 42}
            _ = v42
        }
        v1 = float32((1e10 + (float64((((0.0 - v1) * float32(3.14159)) * v1)) / ((0.1 - 1e10) - 1e10))))
        _ = fi30
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
