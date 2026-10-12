package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int64; F1 uint; F2 uint16 }
type GenS6 struct { F0 string; F1 bool; F2 uint64 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint64 = 18446744073709551615
    v3 := make(map[uint32]uint32)
    var v4 uint64 = v2
    v5 := []string{"", "hello", "\n\t", "hello"}
    if ((v2 != 2) && false) {
        v6 := GenS1{F0: 18446744073709551615, F1: "\n\t"}
        _ = v6
    }
    if false {
        var v7 uint = 405
        v4 = v4
        _ = v7
    } else {
        switch 32767 {
        case 1:
            for fi8 := 0; fi8 < 4; fi8++ {
                fmt.Println(fi8)
                _ = fi8
            }
        case -9223372036854775808:
            v9 := GenS1{F0: 0, F1: ("" + "hello")}
            _ = v9
        case 42:
            fmt.Println(v2, v4)
        default:
            var v10 int32 = -202
            _ = v10
        }
        var v11 int64 = 0
        _ = v11
    }
    for fi12 := 0; fi12 < 16; fi12++ {
        if true {
            v13 := GenS2{F0: 1}
            _ = v13
        } else {
            fmt.Println(v2, fi12, fi12)
        }
        for fi14 := 0; fi14 < 20; fi14++ {
            delete(v3, 2147483647)
            switch (fi12 / (((-(fi12))) | 1)) {
            case 683:
                if true {
                    v2 = v2
                    v15 := GenS1{F0: 658, F1: ""}
                    if true {
                        if (!false) {
                            if false {
                                var v16 int32 = 42
                                _ = v16
                            }
                        } else {
                            fmt.Println(fi14)
                        }
                        v2 += 711
                    } else {
                        delete(v1, 2147483647)
                        fmt.Println(v2, v4)
                    }
                    v17 := GenS3{F0: 488}
                    v2 = (^(v4))
                    fmt.Println(v2, v4, fi14)
                    _ = v15
                    _ = v17
                }
                fmt.Println(v4, v2)
            default:
                v4 = 564
            }
            _ = fi14
        }
        _ = fi12
    }
    if false {
        for fi18 := 0; fi18 < 10; fi18++ {
            fmt.Println(v2, v2, v2)
            for fi19 := 0; fi19 < 15; fi19++ {
                if false {
                    v20 := GenS4{F0: (fi19 >> ((18446744073709551615) & 63))}
                    var v21 uint32 = 861
                    v1[v21] = v21
                    _ = v20
                    _ = v21
                } else {
                    var v22 float32 = 0.0
                    if (!true) {
                        v22 += ((v22 + (v22 * v22)) + v22)
                        v2 = v2
                    } else {
                        v23 := GenS1{F0: 0, F1: "\n\t"}
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        v24 := GenS2{F0: 0}
                        var v25 uint = 2
                        var v26 float32 = (v22 / v22)
                        _ = v23
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                    fmt.Println(v2, fi18, fi18)
                    _ = v22
                }
                switch fi19 {
                case 767:
                    v2 = 2
                case 1:
                    v27 := GenS5{F0: -1, F1: 1, F2: 0}
                    _ = v27
                default:
                    v1[2] = 1
                }
                var v28 bool = true
                _ = fi19
                _ = v28
            }
            _ = fi18
        }
        v4 += v4
    } else {
        v4 -= v2
        var v29 bool = ((1000 != 0) && false)
        _ = v29
    }
    var v30 uint16 = 32767
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v31 int8 = 0
    switch "hello" {
    case "foo bar":
        var v32 uint = 18446744073709551615
        _ = v32
    case "a":
        var v33 float32 = float32(float64((1e10 + (((3.14159 * (-2.25 + 0.0)) + 0.0) / 1.0))))
        _ = v33
    }
    v30 = (v30 - (v30 & 705))
    v31 = (v31 << ((834) & 63))
    _ = len(v5)
    v30 = (^(v30))
    if false {
        v2 = 4294967295
    }
    v34 := GenS6{F0: "hello", F1: (-1 >= 267), F2: 255}
    if len(v5) > 0 {
        v5[0] = "hello"
    }
    v30 = v30
    var v35 uint = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
