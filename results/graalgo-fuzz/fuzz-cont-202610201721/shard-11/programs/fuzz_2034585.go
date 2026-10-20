package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32; F1 int64 }

func main() {
    v1 := []int{32768}
    v2 := make(map[uint32]string)
    var v3 int16 = 1000
    v4 := make(map[string]string)
    v3 -= 32767
    v5 := GenS1{F0: "a"}
    if (-128 > 66) {
        fmt.Println(v3)
    } else {
        v3 -= v3
    }
    fmt.Println(v3)
    v3 -= (-(v3))
    v6 := GenS2{F0: v3}
    if false {
        switch 1 {
        case 431:
            if (!(true || false)) {
                v7 := GenS3{F0: 964, F1: 42}
                v3 += v3
                _ = v7
            } else {
                var v8 int64 = 1
                _ = v8
            }
        case 4294967295:
            if true {
                for fi9 := 0; fi9 < 20; fi9++ {
                    switch "\n\t" {
                    case "\n\t":
                        var v10 uint = 4294967295
                        _ = v10
                    }
                    fmt.Println(fi9)
                    _ = fi9
                }
            } else {
                v3 -= 1
                if ((!true) || ((128 == 32768) || (float32(-2.25) == 0.1))) {
                    switch 127 {
                    case -2147483648:
                        var v11 uint8 = 172
                        _ = v11
                    case 32767:
                        _ = len(v1)
                    case 42:
                        for fi12 := 0; fi12 < 7; fi12++ {
                            v3 -= -32768
                            if true {
                                fmt.Println(v3)
                                fmt.Println(fi12)
                                if true {
                                    for fi13 := 0; fi13 < 8; fi13++ {
                                        var v14 int64 = int64(fi12)
                                        var v15 string = "a"
                                        if false {
                                            switch -2147483648 {
                                            case 32768:
                                                v14 += 9223372036854775807
                                            default:
                                                v15 = ("" + ("\n\t" + v15))
                                            }
                                        } else {
                                            var v16 int32 = 0
                                            _ = v16
                                        }
                                        var v17 bool = true
                                        _ = fi13
                                        _ = v14
                                        _ = v15
                                        _ = v17
                                    }
                                    v3 += 138
                                }
                            }
                            _ = fi12
                        }
                    }
                    v3 = v3
                    var v18 int16 = -468
                    _ = v18
                }
            }
        case 9223372036854775807:
            var v19 uint64 = 287
            _ = v19
        }
    } else {
        fmt.Println(v3)
    }
    for fi20 := 0; fi20 < 1; fi20++ {
        v21 := GenS1{F0: ("hello" + ((("foo bar" + ((((("a" + ("hello" + "hello")) + ("" + "")) + "foo bar") + "foo bar") + "a")) + (("a" + "a") + ("\n\t" + "a"))) + "hello"))}
        var v22 uint32 = 42
        if true {
            fmt.Println(fi20, fi20)
        }
        _ = fi20
        _ = v21
        _ = v22
    }
    v3 += v3
    v3 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
