package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 uint8; F1 string }
func genFn0() int16 {
    var v4 int16 = 0
    for fi5 := 0; fi5 < 17; fi5++ {
        fmt.Println(v4, v4)
        for fi6 := 0; fi6 < 15; fi6++ {
            v4 -= -32768
            if (1 != 156) {
                var v7 float32 = 0.1
                v7 = v7
                _ = v7
            }
            _ = fi6
        }
        switch (-(fi5)) {
        case 9223372036854775807:
            switch ("" + "foo bar") {
            case "a":
                v8 := GenS1{F0: "a"}
                _ = v8
            case "\n\t":
                fmt.Println(fi5, fi5)
            }
        case 32767:
            for fi9 := 0; fi9 < 3; fi9++ {
                v10 := GenS1{F0: "a"}
                v11 := []int{(fi9 + fi9), 9223372036854775807, -1}
                var v12 uint32 = 2147483647
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
            }
        default:
            if (!false) {
                if true {
                    v4 = (v4 % ((v4) | 1))
                }
                var v13 uint = 0
                if (!false) {
                    v13 = (v13 * 521)
                    switch 617 {
                    case -1:
                        fmt.Println(v13, v4)
                    }
                } else {
                    for fi14 := 0; fi14 < 4; fi14++ {
                        fmt.Println(v13, fi14, fi5)
                        v15 := []int{fi5, (fi14 + -1)}
                        _ = fi14
                        _ = v15
                    }
                }
                _ = v13
            }
        }
        _ = fi5
    }
    _ = v4
    return 0
}
func genFn1(p0 int, p1 int16, p2 bool) int32 {
    v16 := []int{-2147483648, 2147483647, p0, 9223372036854775807}
    switch "\n\t" {
    case "hello":
        p0 = p0
    case "":
        for fi17 := 0; fi17 < 8; fi17++ {
            fmt.Println(fi17)
            for fi18 := 0; fi18 < 12; fi18++ {
                var v19 int32 = 0
                v19 += (v19 >> ((2) & 63))
                switch "\n\t" {
                case "hello":
                    var v20 int8 = -128
                    _ = v20
                default:
                    if p2 {
                        var v21 bool = p2
                        _ = v21
                    }
                }
                _ = fi18
                _ = v19
            }
            _ = fi17
        }
    case "\n\t":
        var v22 int16 = -1
        _ = v22
    }
    var v23 int8 = 127
    p1 += genFn0()
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v23
    return -1
}

func main() {
    var v1 uint8 = 2
    v2 := make(map[int]bool)
    var v3 int = 42
    if true {
        v1 += (v1 >> ((2) & 63))
        fmt.Println(v1, v3)
    }
    v24 := GenS2{F0: -249, F1: true}
    v3 = -742
    var v25 bool = false
    switch (^(v3)) {
    case 0:
        if false {
            if v25 {
                if v25 {
                    var v26 string = "hello"
                    var v27 uint64 = 831
                    _ = v26
                    _ = v27
                } else {
                    v28 := GenS1{F0: ""}
                    _ = v28
                }
            }
            v29 := GenS3{F0: v1, F1: ("" + ("" + "foo bar"))}
            _ = v29
        }
    }
    v1 = v1
    for fi30 := 0; fi30 < 20; fi30++ {
        v1 = v1
        _ = fi30
    }
    var v31 int32 = 1
    for fi32 := 0; fi32 < 14; fi32++ {
        v1 = v1
        var v33 uint32 = 801
        var v34 uint8 = (v1 * (v1 * (v1 >> ((4294967295) & 63))))
        _ = fi32
        _ = v33
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v31
}
