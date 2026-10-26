package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16 }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int8 = 42
    v3 := make(map[uint32]uint32)
    var v4 uint64 = 0
    var v5 int = 127
    v1 = (v1 & v1)
    var v6 float64 = 0.1
    switch "a" {
    case "hello":
        v7 := GenS1{F0: v5, F1: 766}
        _ = v7
    case "foo bar":
        v8 := []int{326}
        _ = v8
    default:
        var v9 float32 = 0.0
        _ = v9
    }
    var v10 int = int(v4)
    for fi11 := 0; fi11 < 9; fi11++ {
        var v12 uint64 = 0
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 11; fi13++ {
        v14 := GenS1{F0: -2147483648, F1: int16(v1)}
        switch 0 {
        case -1:
            v15 := []int{int(v1), 127, fi13, 869}
            _ = v15
        case 1:
            if ((v6 < float64((0.0 + ((float32((v6 / float64((1.5 / 1.5)))) / 1.0) / 0.1)))) || false) {
                v16 := []int{42, (v5 + fi13)}
                v17 := GenS2{F0: 0}
                _ = v16
                _ = v17
            }
        case 9223372036854775807:
            v18 := GenS3{F0: 681}
            _ = v18
        }
        _ = fi13
        _ = v14
    }
    switch "\n\t" {
    case "":
        var v19 string = ("foo bar" + ("hello" + "foo bar"))
        _ = v19
    }
    v20 := []int{}
    for fi21 := 0; fi21 < 15; fi21++ {
        v22 := GenS2{F0: 463}
        v23 := v3[565]
        _ = fi21
        _ = v22
        _ = v23
    }
    fmt.Println(v1, v10)
    v4 += (v4 * (^(v4)))
    var v24 int8 = (v2 & (v2 * (v2 - v2)))
    switch ("a" + "") {
    case "\n\t":
        switch "\n\t" {
        case "\n\t":
            v24 += v2
        default:
            for fi25 := 0; fi25 < 20; fi25++ {
                v26 := v3[656]
                if (0 > 1) {
                    var v27 int8 = v2
                    _ = v27
                } else {
                    delete(v3, v26)
                    var v28 uint32 = (v26 % ((49) | 1))
                    _ = v28
                }
                for fi29 := 0; fi29 < 9; fi29++ {
                    switch 32767 {
                    case -316:
                        if false {
                            fmt.Println(v4)
                            v24 += v24
                            if ("\n\t" <= (("foo bar" + "") + ("hello" + "hello"))) {
                                v6 = 0.0
                            } else {
                                v26 = (v26 & (^(v26)))
                            }
                        }
                        switch ("foo bar" + "") {
                        case "a":
                            v2 -= v24
                        case "":
                            fmt.Println(v24, v5)
                        case "hello":
                            fmt.Println(fi25)
                        default:
                            fmt.Println(v2, v6)
                        }
                        v10 = fi29
                    default:
                        if true {
                            if true {
                                var v30 int32 = -1
                                _ = len(v20)
                                _ = v30
                            }
                            v6 = v6
                        } else {
                            v31 := GenS3{F0: 0}
                            v24 = (v24 >> ((1) & 63))
                            _ = v31
                        }
                    }
                    _ = fi29
                }
                _ = fi25
                _ = v26
            }
        }
    case "foo bar":
        v32 := GenS2{F0: 0}
        _ = v32
    }
    v20 = append(v20, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
