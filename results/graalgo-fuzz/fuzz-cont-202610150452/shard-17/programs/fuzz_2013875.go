package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 string; F1 int16; F2 uint }
func genFn0() uint64 {
    if (-1 <= -1) {
        v7 := GenS1{F0: "foo bar", F1: "a"}
        _ = v7
    } else {
        var v8 int8 = -69
        _ = v8
    }
    if true {
        v9 := []int{-714, -2147483648}
        _ = v9
    }
    return 1
}

func main() {
    var v1 uint32 = 1
    var v2 uint16 = 42
    v3 := make(map[string]int)
    v4 := []int{-1, 2147483647}
    var v5 uint = 255
    var v6 string = "\n\t"
    v6 = ((v6 + "hello") + (v6 + (((v6 + v6) + v6) + "")))
    _ = len(v4)
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 int16 = 1000
        fmt.Println(v1, fi10)
        var v12 float32 = (0.1 * 0.1)
        _ = fi10
        _ = v11
        _ = v12
    }
    v6 = ("a" + (v6 + v6))
    v1 = v1
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v5 = (v5 - 231)
    v1 = (v1 ^ (^(v1)))
    if true {
        var v13 int16 = -32768
        _ = v13
    }
    v6 = v6
    if false {
        for fi14 := 0; fi14 < 6; fi14++ {
            v6 = (v6 + v6)
            switch fi14 {
            case -1:
                switch v6 {
                case "a":
                    v2 -= v2
                case "hello":
                    switch fi14 {
                    case 32768:
                        var v15 int = 127
                        _ = v15
                    case 9223372036854775807:
                        if false {
                            switch (v6 + "foo bar") {
                            case "":
                                _ = len(v4)
                            case "foo bar":
                                var v16 int16 = 665
                                _ = v16
                            case "hello":
                                v17 := GenS1{F0: "hello", F1: v6}
                                _ = v17
                            }
                            v6 = (("hello" + ("foo bar" + v6)) + ("\n\t" + (v6 + "foo bar")))
                        }
                    }
                default:
                    switch "hello" {
                    case "a":
                        v18 := GenS2{F0: 185, F1: true}
                        _ = v18
                    case "foo bar":
                        switch "hello" {
                        case "foo bar":
                            v1 += (v1 >> ((v5) & 63))
                        case "\n\t":
                            fmt.Println(fi14, v2, v5)
                        }
                    }
                }
            }
            _ = fi14
        }
    }
    var v19 int32 = 682
    delete(v3, ((v6 + v6) + (v6 + (((("foo bar" + v6) + v6) + "\n\t") + (((v6 + v6) + v6) + v6)))))
    _ = len(v3)
    v20 := v3[v6]
    var v21 int64 = -936
    v21 -= (v21 >> ((v5) & 63))
    if (2147483647 > v1) {
        v22 := GenS3{F0: "a", F1: int16(v21), F2: (v5 | (^(v5)))}
        fmt.Println(v20)
        switch (v6 + ("foo bar" + v6)) {
        case "":
            fmt.Println(v19, v19, v19)
        case "foo bar":
            v4 = append(v4, (v20 % (((v20 % ((v20) | 1))) | 1)))
        default:
            v1 = v1
        }
        _ = v22
    } else {
        fmt.Println(v2, v19)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
