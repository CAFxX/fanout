package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 bool; F2 int16 }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 int32; F1 string; F2 int }
func genFn0(p0 int8) bool {
    if true {
        var v7 float32 = (float32(1e10) * float32(1e10))
        switch "\n\t" {
        case "foo bar":
            v7 = (v7 / (((v7 + float32(1e10)) * v7) - v7))
        case "hello":
            var v8 int16 = 42
            _ = v8
        default:
            v9 := GenS1{F0: 1}
            _ = v9
        }
        _ = v7
    } else {
        fmt.Println(p0)
        var v10 int8 = 127
        _ = v10
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        v12 := GenS2{F0: "foo bar", F1: (114 == 2), F2: -1}
        var v13 uint8 = 2
        fmt.Println(fi11, v13)
        _ = fi11
        _ = v12
        _ = v13
    }
    v14 := GenS2{F0: "hello", F1: false, F2: -32768}
    p0 = (p0 >> ((1) & 63))
    _ = p0
    _ = v14
    return (42 == 564)
}

func main() {
    v1 := []uint32{2147483647}
    var v2 uint64 = 1
    var v3 int32 = int32(v2)
    v4 := make(map[uint32]int)
    v5 := []string{}
    var v6 uint64 = 4294967295
    v15 := GenS2{F0: ((("a" + "hello") + "a") + "\n\t"), F1: false, F2: -350}
    for fi16 := 0; fi16 < 18; fi16++ {
        v17 := GenS3{F0: 692, F1: 255}
        _ = fi16
        _ = v17
    }
    v6 += (v2 * (^(v2)))
    var v18 int8 = -6
    var v19 int = 0
    _ = len(v5)
    switch 650 {
    case -44:
        fmt.Println(v2, v3, v19)
    }
    if false {
        for fi20 := 0; fi20 < 7; fi20++ {
            v21 := GenS4{F0: v3, F1: "foo bar", F2: 9223372036854775807}
            v19 = fi20
            v18 -= 61
            _ = fi20
            _ = v21
        }
        v3 = v3
    } else {
        var v22 uint64 = (v2 << ((2) & 63))
        _ = v22
    }
    for fi23 := 0; fi23 < 11; fi23++ {
        v5 = append(v5, "")
        if ("hello" >= (("a" + (("\n\t" + (("\n\t" + ("" + ("\n\t" + "hello"))) + "foo bar")) + "\n\t")) + ("a" + ""))) {
            var v24 float64 = ((0.1 - 0.1) * 3.14159)
            switch (v19 ^ (fi23 | (v19 / (((-(fi23))) | 1)))) {
            case 128:
                if true {
                    v25 := GenS1{F0: 0}
                    _ = v25
                } else {
                    var v26 uint64 = (v6 << ((255) & 63))
                    if true {
                        for fi27 := 0; fi27 < 17; fi27++ {
                            fmt.Println(v26, v6, v26)
                            if true {
                                if len(v1) > 0 {
                                    v1[0] = 448
                                }
                                fmt.Println(v26, fi27, v19)
                                var v28 int32 = 0
                                _ = v28
                            }
                            var v29 uint32 = 529
                            _ = fi27
                            _ = v29
                        }
                        var v30 int64 = -1
                        _ = v30
                    } else {
                        v19 = (v19 + (v19 * v19))
                        if len(v1) > 0 {
                            v1[0] = 42
                        }
                    }
                    v24 -= 0.1
                    _ = v26
                }
            case 9223372036854775807:
                var v31 uint64 = 0
                _ = v31
            case -832:
                var v32 uint64 = 0
                _ = v32
            default:
                for fi33 := 0; fi33 < 20; fi33++ {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    fmt.Println(v6)
                    _ = fi33
                }
            }
            for fi34 := 0; fi34 < 8; fi34++ {
                var v35 string = "hello"
                var v36 uint64 = 2
                v3 += (v3 % ((v3) | 1))
                _ = fi34
                _ = v35
                _ = v36
            }
            _ = v24
        }
        _ = fi23
    }
    fmt.Println(v6, v3)
    v3 = 897
    _ = len(v4)
    switch v19 {
    case 139:
        switch v19 {
        case -106:
            v5 = append(v5, "foo bar")
        default:
            fmt.Println(v3)
        }
    case 32767:
        v37 := GenS3{F0: 4294967295, F1: 255}
        _ = v37
    case 895:
        var v38 uint = 1
        _ = v38
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
