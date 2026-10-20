package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 int32 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{}
    v3 := make(map[int]int)
    var v4 uint16 = 65535
    v4 = v4
    v3[-378] = 127
    var v5 uint16 = v4
    v5 = 655
    var v6 float32 = 3.14159
    fmt.Println(v5)
    v5 = 42
    fmt.Println(v5, v5, v6)
    if false {
        v7 := GenS1{F0: 364, F1: 740, F2: 42}
        _ = v7
    }
    v4 += v5
    v6 = v6
    v4 += 42
    var v8 uint16 = v4
    var v9 float32 = (((3.14159 - ((v6 + -2.25) * (-2.25 * float32(((3.14159 * 0.1) / 0.1))))) / float32(((3.14159 + (-2.25 - (0.0 + 1e10))) * 3.14159))) + 0.0)
    switch "\n\t" {
    case "hello":
        fmt.Println(v9, v8)
    }
    if true {
        v5 = (v4 * (^(v4)))
    } else {
        var v10 uint64 = 0
        v4 = v5
        switch 42 {
        case -9223372036854775808:
            switch ("foo bar" + ("" + "hello")) {
            case "foo bar":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case "hello":
                v10 = 416
            }
        default:
            for fi11 := 0; fi11 < 3; fi11++ {
                for fi12 := 0; fi12 < 13; fi12++ {
                    if false {
                        v2 = append(v2, "")
                    } else {
                        var v13 int8 = int8(v8)
                        if len(v2) > 0 {
                            v2[0] = ((("foo bar" + ("\n\t" + "foo bar")) + (("a" + ("\n\t" + "foo bar")) + "hello")) + "foo bar")
                        }
                        _ = v13
                    }
                    var v14 bool = (-128 > 127)
                    _ = fi12
                    _ = v14
                }
                switch "hello" {
                case "\n\t":
                    switch "foo bar" {
                    case "\n\t":
                        for fi15 := 0; fi15 < 3; fi15++ {
                            switch "a" {
                            case "a":
                                var v16 uint64 = 2
                                _ = v16
                            case "\n\t":
                                fmt.Println(v8, fi15, fi11)
                            default:
                                fmt.Println(v10)
                            }
                            v6 -= ((v6 * ((v9 + v9) + v9)) / 1.0)
                            _ = len(v1)
                            var v17 uint8 = 142
                            _ = fi15
                            _ = v17
                        }
                    default:
                        var v18 int8 = 0
                        _ = v18
                    }
                case "a":
                    for fi19 := 0; fi19 < 9; fi19++ {
                        v3[fi19] = (fi19 / (((fi19 / (((fi11 & int(v10))) | 1))) | 1))
                        var v20 uint16 = (v4 ^ v4)
                        _ = fi19
                        _ = v20
                    }
                case "hello":
                    fmt.Println(v10, fi11, v9)
                }
                v5 -= (^(v8))
                _ = fi11
            }
        }
        _ = v10
    }
    v8 = (v4 / ((v4) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
