package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint; F1 string }

func main() {
    v1 := []string{"foo bar"}
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]bool)
    var v4 string = "hello"
    v4 = "a"
    switch 481 {
    case 321:
        for fi5 := 0; fi5 < 3; fi5++ {
            v6 := GenS1{F0: 0, F1: 1, F2: 1}
            v4 = v4
            _ = fi5
            _ = v6
        }
    default:
        v7 := GenS2{F0: -869}
        _ = v7
    }
    v1 = append(v1, (v4 + "foo bar"))
    var v8 int8 = -112
    var v9 uint32 = 42
    v9 = (v9 + 176)
    var v10 uint32 = 419
    delete(v2, v9)
    fmt.Println(v9, v10)
    fmt.Println(v10, v8)
    for fi11 := 0; fi11 < 16; fi11++ {
        var v12 uint16 = 811
        var v13 int = (fi11 * (fi11 - fi11))
        v2[(v10 & v10)] = v13
        _ = fi11
        _ = v12
        _ = v13
    }
    if true {
        v14 := GenS3{F0: 0, F1: ("a" + "\n\t")}
        switch 720 {
        case 128:
            if false {
                if (255 >= 65535) {
                    fmt.Println(v10, v8)
                    v8 = v8
                }
                switch "hello" {
                case "a":
                    fmt.Println(v8, v4, v9)
                case "hello":
                    v10 -= 2147483647
                case "":
                    for fi15 := 0; fi15 < 14; fi15++ {
                        v8 -= v8
                        for fi16 := 0; fi16 < 2; fi16++ {
                            var v17 uint16 = 757
                            var v18 uint32 = v10
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                        _ = fi15
                    }
                }
                v9 = 404
            }
        case 32768:
            var v19 int64 = 0
            _ = v19
        default:
            for fi20 := 0; fi20 < 13; fi20++ {
                var v21 float32 = -2.25
                var v22 int8 = 100
                _ = fi20
                _ = v21
                _ = v22
            }
        }
        for fi23 := 0; fi23 < 5; fi23++ {
            _ = len(v1)
            v8 = v8
            if false {
                if (!true) {
                    switch 899 {
                    case 127:
                        switch int(v10) {
                        case -9223372036854775808:
                            fmt.Println(v4, v4)
                        case 42:
                            var v24 uint = 553
                            _ = v24
                        default:
                            if true {
                                fmt.Println(v10)
                            } else {
                                for fi25 := 0; fi25 < 2; fi25++ {
                                    var v26 int = fi23
                                    fmt.Println(v10)
                                    _ = fi25
                                    _ = v26
                                }
                            }
                        }
                    case 0:
                        if (v8 <= (v8 | (v8 * 0))) {
                            var v27 uint16 = 75
                            v28 := GenS2{F0: 371}
                            _ = v27
                            _ = v28
                        } else {
                            v8 += v8
                        }
                    default:
                        v3[v9] = true
                    }
                }
                for fi29 := 0; fi29 < 7; fi29++ {
                    if true {
                        v30 := GenS3{F0: 80, F1: (("" + "a") + "")}
                        var v31 float32 = (1e10 / 0.1)
                        _ = v30
                        _ = v31
                    }
                    var v32 float32 = (0.1 / (-2.25 + (3.14159 + 1.5)))
                    _ = fi29
                    _ = v32
                }
                v4 = ((v4 + v4) + ("\n\t" + "a"))
            }
            _ = fi23
        }
        _ = v14
    } else {
        var v33 bool = true
        fmt.Println(v33, v8, v9)
        _ = v33
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
