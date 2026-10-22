package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 uint16 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]string)
    var v3 uint = 65535
    var v4 bool = false
    for fi5 := 0; fi5 < 9; fi5++ {
        var v6 string = "\n\t"
        var v7 int32 = 0
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 bool = (203 < 2)
    v4 = v4
    var v9 int64 = 42
    v10 := []int{4294967295, 128, 462, 565}
    v11 := GenS1{F0: 42}
    v9 += v9
    switch 42 {
    case 710:
        v9 = v9
    }
    v3 += (^(v3))
    for fi12 := 0; fi12 < 3; fi12++ {
        switch (fi12 + 4294967295) {
        case 4294967295:
            v3 -= (v3 >> ((v3) & 63))
        case -973:
            if false {
                v3 += (^(v3))
            }
        case 0:
            var v13 uint = v3
            _ = v13
        }
        var v14 uint64 = uint64(v3)
        _ = fi12
        _ = v14
    }
    var v15 uint32 = 42
    if v4 {
        if v8 {
            for fi16 := 0; fi16 < 14; fi16++ {
                for fi17 := 0; fi17 < 17; fi17++ {
                    var v18 int = fi16
                    v4 = (1 == 0)
                    _ = fi17
                    _ = v18
                }
                _ = fi16
            }
            switch 1 {
            case 42:
                fmt.Println(v9, v3, v4)
            case 1:
                if true {
                    v3 = v3
                    v2[2147483648] = ("hello" + "a")
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                } else {
                    switch ("" + ("foo bar" + (((("" + ("" + "hello")) + "") + ("foo bar" + "hello")) + ""))) {
                    case "\n\t":
                        for fi19 := 0; fi19 < 14; fi19++ {
                            v4 = (((18 != -128) || v4) && true)
                            _ = fi19
                        }
                    case "hello":
                        if (!v8) {
                            v20 := GenS2{F0: 208, F1: 0}
                            if v4 {
                                v2[v15] = "\n\t"
                                fmt.Println(v8, v15, v15)
                                var v21 int64 = v9
                                _ = v21
                            }
                            var v22 uint64 = 18446744073709551615
                            _ = v20
                            _ = v22
                        } else {
                            var v23 int16 = -1
                            _ = v23
                        }
                    case "a":
                        for fi24 := 0; fi24 < 1; fi24++ {
                            var v25 uint = v3
                            _ = fi24
                            _ = v25
                        }
                    default:
                        v26 := GenS3{F0: true, F1: 100}
                        _ = v26
                    }
                    if (!v4) {
                        v27 := GenS4{F0: 1}
                        _ = v27
                    } else {
                        switch "foo bar" {
                        case "foo bar":
                            _ = len(v10)
                        case "":
                            if v4 {
                                var v28 int64 = v9
                                switch "\n\t" {
                                case "hello":
                                    if len(v10) > 0 {
                                        v10[0] = 0
                                    }
                                }
                                _ = v28
                            }
                        }
                    }
                    _ = len(v10)
                }
            case -1:
                fmt.Println(v15, v9)
            default:
                fmt.Println(v3, v15, v3)
            }
        }
    }
    for fi29 := 0; fi29 < 12; fi29++ {
        var v30 int64 = (v9 | -136)
        var v31 string = ("hello" + "")
        _ = fi29
        _ = v30
        _ = v31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
