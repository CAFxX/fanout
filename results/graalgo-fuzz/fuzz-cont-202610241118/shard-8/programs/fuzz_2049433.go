package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 int32; F1 uint32 }

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{}
    var v3 uint32 = 0
    v4 := make(map[int]bool)
    var v5 int16 = 0
    v3 += 4294967295
    switch 758 {
    case -235:
        v6 := GenS1{F0: 242, F1: -2147483648}
        _ = v6
    case 127:
        var v7 uint64 = 1
        _ = v7
    case 32767:
        var v8 int64 = -143
        _ = v8
    default:
        delete(v4, -660)
    }
    if (1 == 2) {
        v9 := v1["\n\t"]
        for fi10 := 0; fi10 < 1; fi10++ {
            if (65535 > 0) {
                if ("foo bar" > ("" + (((("a" + (("\n\t" + "foo bar") + (("\n\t" + "") + "foo bar"))) + ("" + "a")) + "\n\t") + "foo bar"))) {
                    v11 := GenS2{F0: 39, F1: false}
                    _ = v11
                }
                switch "a" {
                case "foo bar":
                    fmt.Println(v9)
                case "hello":
                    switch (("a" + "foo bar") + "") {
                    case "foo bar":
                        for fi12 := 0; fi12 < 4; fi12++ {
                            var v13 string = "hello"
                            _ = fi12
                            _ = v13
                        }
                    default:
                        if true {
                            var v14 int32 = 0
                            v9 = (^(v9))
                            _ = v14
                        } else {
                            for fi15 := 0; fi15 < 12; fi15++ {
                                v5 += v5
                                _ = fi15
                            }
                        }
                    }
                default:
                    switch fi10 {
                    case 2147483647:
                        fmt.Println(v5)
                    }
                }
            }
            v16 := GenS2{F0: 581, F1: (!((v5 >> ((18446744073709551615) & 63)) <= -1))}
            _ = fi10
            _ = v16
        }
        delete(v1, "")
        _ = v9
    }
    switch 2147483647 {
    case 1:
        var v17 uint16 = 187
        _ = v17
    }
    fmt.Println(v5, v5)
    v18 := GenS3{F0: 1, F1: (v3 + 42)}
    v3 = (v3 << ((18446744073709551615) & 63))
    var v19 uint64 = 222
    v3 += v3
    for fi20 := 0; fi20 < 12; fi20++ {
        var v21 bool = true
        v3 -= 0
        if false {
            var v22 float32 = (-2.25 / -2.25)
            var v23 int16 = (v5 % ((v5) | 1))
            _ = v22
            _ = v23
        }
        _ = fi20
        _ = v21
    }
    fmt.Println(v19)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
