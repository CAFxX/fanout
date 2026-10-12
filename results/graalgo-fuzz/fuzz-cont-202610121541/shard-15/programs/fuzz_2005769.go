package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64 }

func main() {
    v1 := []uint32{653, 1}
    var v2 uint32 = 459
    v3 := []int{-9223372036854775808, -908}
    if (!false) {
        var v4 string = "foo bar"
        _ = v4
    }
    v5 := GenS1{F0: ""}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v6 uint32 = v2
    var v7 int64 = -543
    if (810 != int32(v7)) {
        switch 937 {
        case -582:
            switch "hello" {
            case "a":
                v6 = 4294967295
            }
        case 2147483647:
            for fi8 := 0; fi8 < 2; fi8++ {
                if false {
                    if true {
                        v7 = int64(fi8)
                    } else {
                        var v9 bool = (fi8 == (fi8 % (((-(fi8))) | 1)))
                        _ = v9
                    }
                    fmt.Println(fi8)
                    for fi10 := 0; fi10 < 9; fi10++ {
                        _ = len(v1)
                        if true {
                            var v11 uint64 = 18446744073709551615
                            if true {
                                fmt.Println(v2)
                                v2 -= v2
                            } else {
                                var v12 int8 = -1
                                fmt.Println(v2, fi10)
                                _ = v12
                            }
                            switch "hello" {
                            case "hello":
                                var v13 int32 = 42
                                _ = v13
                            default:
                                fmt.Println(fi10, fi8)
                            }
                            fmt.Println(fi10, fi10)
                            v11 = v11
                            _ = v11
                        }
                        var v14 float64 = 1.5
                        _ = fi10
                        _ = v14
                    }
                } else {
                    for fi15 := 0; fi15 < 18; fi15++ {
                        if len(v1) > 0 {
                            v1[0] = v6
                        }
                        var v16 float64 = float64(0.0)
                        _ = fi15
                        _ = v16
                    }
                    fmt.Println(v2, v6, v2)
                    switch (^(fi8)) {
                    case 975:
                        v17 := GenS2{F0: 42}
                        _ = v17
                    case -405:
                        for fi18 := 0; fi18 < 19; fi18++ {
                            var v19 int8 = 0
                            _ = fi18
                            _ = v19
                        }
                    default:
                        v7 -= 42
                    }
                }
                v2 += 0
                _ = fi8
            }
        default:
            v7 = 480
        }
        if len(v3) > 0 {
            v3[0] = 661
        }
        v3 = append(v3, 32768)
    }
    var v20 int = 128
    var v21 string = ("foo bar" + "hello")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v5
    _ = v6
    _ = v7
}
