package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
func genFn0(p0 uint8) uint32 {
    v5 := []int{9223372036854775807}
    var v6 int64 = -9223372036854775808
    fmt.Println(p0)
    _ = p0
    _ = v5
    _ = v6
    return 2147483648
}

func main() {
    var v1 uint16 = 2
    v2 := make(map[string]int)
    v3 := make(map[uint32]string)
    v4 := make(map[int]string)
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 int16 = 0
        _ = fi7
        _ = v8
    }
    var v9 uint32 = 1
    var v10 float32 = -2.25
    v10 -= v10
    var v11 int = 42
    v12 := GenS1{F0: -1, F1: v1}
    switch v11 {
    case 241:
        v1 += v1
    case 32768:
        for fi13 := 0; fi13 < 9; fi13++ {
            _ = len(v2)
            if false {
                var v14 uint = 255
                for fi15 := 0; fi15 < 11; fi15++ {
                    var v16 float64 = -2.25
                    if false {
                        if (-9223372036854775808 < 9223372036854775807) {
                            switch (("\n\t" + "\n\t") + "foo bar") {
                            case "hello":
                                v2[("hello" + "foo bar")] = fi15
                            case "":
                                var v17 uint8 = 128
                                _ = v17
                            default:
                                fmt.Println(v16, fi13, v16)
                            }
                            switch "hello" {
                            case "hello":
                                v14 = 0
                            default:
                                switch "a" {
                                case "":
                                    var v18 uint16 = v1
                                    _ = v18
                                }
                            }
                        } else {
                            switch "foo bar" {
                            case "\n\t":
                                fmt.Println(fi15, v11)
                            }
                            v1 = 42
                        }
                    }
                    _ = fi15
                    _ = v16
                }
                _ = v14
            } else {
                v19 := GenS1{F0: -948, F1: 927}
                _ = v19
            }
            _ = fi13
        }
    }
    v10 += v10
    for fi20 := 0; fi20 < 5; fi20++ {
        v9 -= v9
        _ = fi20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
