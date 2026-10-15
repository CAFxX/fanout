package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 int32 }

func main() {
    var v1 int32 = 599
    var v2 int = -966
    v3 := make(map[uint32]uint32)
    fmt.Println(v1, v1)
    var v4 uint = 65535
    fmt.Println(v4, v1, v2)
    for fi5 := 0; fi5 < 5; fi5++ {
        for fi6 := 0; fi6 < 15; fi6++ {
            v7 := []int{v2, v2, (v2 & (v2 / ((-1) | 1)))}
            if (v4 <= v4) {
                switch "hello" {
                case "\n\t":
                    switch ("foo bar" + "") {
                    case "hello":
                        v2 = fi6
                    case "\n\t":
                        if (73 < 255) {
                            v8 := GenS1{F0: "", F1: 2, F2: (v1 - 1)}
                            _ = v8
                        } else {
                            if (2147483647 <= 0) {
                                v4 = 4294967295
                                if (32767 < -1) {
                                    if (-1 >= v1) {
                                        if len(v7) >= 2 {
                                            v7 = v7[1:len(v7)]
                                        }
                                    } else {
                                        fmt.Println(v2, v1, v2)
                                        if true {
                                            var v9 float64 = 1e10
                                            var v10 string = ("foo bar" + "foo bar")
                                            var v11 uint64 = 42
                                            _ = v10
                                            _ = v11
                                            _ = v9
                                        } else {
                                            fmt.Println(fi6, v2, fi6)
                                            var v12 uint64 = 414
                                            var v13 float64 = (float64(float32(float64(-2.25))) / 3.14159)
                                            _ = v12
                                            _ = v13
                                        }
                                    }
                                }
                            }
                        }
                    case "a":
                        v14 := v3[2147483647]
                        _ = v14
                    default:
                        v15 := GenS1{F0: "", F1: 42, F2: 2147483647}
                        _ = v15
                    }
                default:
                    var v16 int32 = (v1 << ((v4) & 63))
                    _ = v16
                }
            }
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    var v17 int16 = 1
    var v18 int32 = int32(v2)
    var v19 int32 = (v18 >> ((4294967295) & 63))
    var v20 uint16 = 946
    fmt.Println(v1, v17)
    v21 := []int{}
    delete(v3, 231)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
