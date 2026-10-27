package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{}
    v3 := make(map[uint32]string)
    v4 := make(map[string]uint32)
    var v5 uint64 = 2
    v6 := []int{}
    var v7 string = ""
    v5 += v5
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 bool = ((v7 + v7) <= (("\n\t" + "") + ("foo bar" + (v7 + (("\n\t" + ("hello" + v7)) + (v7 + v7))))))
        _ = fi8
        _ = v9
    }
    if (v5 < (v5 % ((324) | 1))) {
        var v10 uint32 = 1
        _ = v10
    } else {
        for fi11 := 0; fi11 < 20; fi11++ {
            switch v7 {
            case "a":
                v7 = (v7 + "\n\t")
            default:
                if (!(255 >= 932)) {
                    for fi12 := 0; fi12 < 12; fi12++ {
                        var v13 uint8 = 209
                        delete(v1, 0)
                        _ = fi12
                        _ = v13
                    }
                    v7 = v7
                } else {
                    v14 := GenS1{F0: 1, F1: 1}
                    for fi15 := 0; fi15 < 7; fi15++ {
                        var v16 int = (fi11 * fi15)
                        switch (-(fi11)) {
                        case 127:
                            if ((-1 >= 1) || false) {
                                switch (fi15 ^ (fi15 >> ((40) & 63))) {
                                case 42:
                                    _ = len(v2)
                                }
                                switch (("" + v7) + v7) {
                                case "foo bar":
                                    v16 += v16
                                case "hello":
                                    var v17 uint64 = (v5 & v5)
                                    _ = v17
                                case "":
                                    if len(v2) > 0 {
                                        v2[0] = ""
                                    }
                                default:
                                    var v18 int8 = 127
                                    _ = v18
                                }
                            }
                        case 42:
                            v19 := GenS1{F0: -128, F1: 42}
                            _ = v19
                        }
                        fmt.Println(v5)
                        _ = fi15
                        _ = v16
                    }
                    switch (("foo bar" + v7) + v7) {
                    case "hello":
                        v5 += (v5 >> ((197) & 63))
                    case "foo bar":
                        var v20 uint16 = 845
                        _ = v20
                    default:
                        v5 = v5
                    }
                    _ = v14
                }
            }
            fmt.Println(fi11, fi11)
            v7 = ("\n\t" + ("foo bar" + (v7 + ("a" + ("a" + "hello")))))
            _ = fi11
        }
        v7 = "foo bar"
    }
    var v21 float32 = 0.1
    v21 = (1e10 - v21)
    if false {
        v3[2] = "hello"
    }
    var v22 string = "a"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
