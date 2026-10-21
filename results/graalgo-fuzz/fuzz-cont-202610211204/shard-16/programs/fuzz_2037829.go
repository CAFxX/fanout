package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 bool; F2 int32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int64; F1 string }

func main() {
    var v1 int64 = -919
    v2 := make(map[string]uint32)
    var v3 int16 = 42
    v4 := make(map[uint32]bool)
    switch "a" {
    case "\n\t":
        v3 = (v3 >> ((290) & 63))
    case "a":
        var v5 bool = true
        _ = v5
    }
    if false {
        var v6 uint16 = 1
        var v7 uint32 = 65535
        for fi8 := 0; fi8 < 19; fi8++ {
            switch 957 {
            case -740:
                switch "" {
                case "foo bar":
                    switch "hello" {
                    case "":
                        v9 := GenS1{F0: "foo bar", F1: -128}
                        _ = v9
                    case "a":
                        v7 -= (v7 << ((0) & 63))
                    case "foo bar":
                        v10 := []int{fi8, fi8, fi8}
                        _ = v10
                    }
                case "hello":
                    v6 -= 698
                case "":
                    v7 = v7
                }
            default:
                fmt.Println(v1, v7)
            }
            var v11 int16 = v3
            _ = fi8
            _ = v11
        }
        _ = v6
        _ = v7
    } else {
        fmt.Println(v3)
        v1 -= (v1 * v1)
        v3 += (v3 >> ((42) & 63))
    }
    v12 := []int{-9223372036854775808, 32767}
    var v13 int16 = (v3 % ((v3) | 1))
    var v14 uint16 = 32767
    v15 := GenS1{F0: "a", F1: 127}
    switch ("" + "a") {
    case "\n\t":
        if false {
            var v16 uint16 = (v14 % ((v14) | 1))
            _ = v16
        }
    default:
        v14 -= v14
    }
    var v17 bool = true
    switch (("\n\t" + "\n\t") + "\n\t") {
    case "\n\t":
        v2[""] = 2
    case "foo bar":
        v18 := GenS1{F0: ("a" + "hello"), F1: -1}
        _ = v18
    case "":
        if (187 != 65535) {
            if (true && false) {
                var v19 uint = 42
                if v17 {
                    if v17 {
                        if (-537 > -127) {
                            _ = len(v2)
                        }
                    }
                    v20 := GenS2{F0: 135}
                    for fi21 := 0; fi21 < 1; fi21++ {
                        v14 += uint16(v19)
                        _ = fi21
                    }
                    _ = v20
                } else {
                    if len(v12) > 0 {
                        v12[0] = 2147483647
                    }
                }
                v22 := GenS2{F0: 119}
                _ = v19
                _ = v22
            }
            v23 := GenS3{F0: -2147483648, F1: (v17 && (70 >= 183)), F2: -865}
            _ = v23
        }
    }
    var v24 uint16 = 119
    v14 = v14
    switch "" {
    case "hello":
        v1 += (v1 - v1)
    }
    v24 = 255
    v25 := GenS2{F0: 0}
    if (1 == (v1 * (v1 | (-(v1))))) {
        for fi26 := 0; fi26 < 12; fi26++ {
            var v27 int8 = 100
            _ = fi26
            _ = v27
        }
        fmt.Println(v13)
        for fi28 := 0; fi28 < 7; fi28++ {
            if (-1 < v13) {
                _ = len(v12)
            } else {
                if false {
                    v3 += (v3 - 647)
                }
                v17 = false
            }
            if v17 {
                var v29 uint32 = 1
                if true {
                    fmt.Println(fi28)
                } else {
                    v30 := GenS4{F0: 42}
                    for fi31 := 0; fi31 < 7; fi31++ {
                        if false {
                            var v32 string = ("hello" + "hello")
                            var v33 string = v32
                            v34 := GenS5{F0: (v1 % ((v1) | 1)), F1: ((v33 + v32) + v33)}
                            _ = v32
                            _ = v33
                            _ = v34
                        } else {
                            var v35 int = (-(fi28))
                            v1 = (v1 >> ((403) & 63))
                            delete(v2, "hello")
                            _ = v35
                        }
                        _ = fi31
                    }
                    if (fi28 == fi28) {
                        v29 = (v29 - (v29 % (((^(v29))) | 1)))
                    } else {
                        fmt.Println(v14, v3)
                    }
                    _ = v30
                }
                _ = v29
            }
            _ = fi28
        }
    }
    _ = len(v4)
    var v36 float32 = float32((0.1 * 1.5))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v36
    _ = v4
}
