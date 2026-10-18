package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
func genFn0(p0 uint8) int8 {
    p0 = (p0 ^ p0)
    v5 := []int{127}
    _ = p0
    _ = v5
    return 1
}

func main() {
    var v1 int16 = 1000
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]bool)
    var v4 float64 = 3.14159
    fmt.Println(v1, v1)
    var v6 int16 = v1
    if true {
        v1 = 1
        var v7 int16 = -504
        for fi8 := 0; fi8 < 2; fi8++ {
            v9 := GenS1{F0: 356, F1: 0}
            _ = fi8
            _ = v9
        }
        _ = v7
    }
    switch (("" + "hello") + "") {
    case "a":
        v10 := []int{-312}
        _ = v10
    case "":
        if true {
            for fi11 := 0; fi11 < 6; fi11++ {
                v1 = (^(v6))
                v1 = v1
                var v12 float64 = v4
                _ = fi11
                _ = v12
            }
            v6 -= v1
        } else {
            if false {
                v13 := GenS1{F0: 255, F1: 106}
                fmt.Println(v1)
                var v14 int = -2147483648
                _ = v13
                _ = v14
            } else {
                for fi15 := 0; fi15 < 10; fi15++ {
                    if (((v4 * ((((float64((-2.25 - float32(v4))) * v4) / float64(-2.25)) + float64(0.1)) + (v4 / (-2.25 + (v4 - (0.1 / (v4 - 0.0))))))) >= ((3.14159 - float64((-2.25 / 1.5))) * 1.5)) && false) {
                        for fi16 := 0; fi16 < 14; fi16++ {
                            var v17 uint64 = 679
                            var v18 uint8 = 128
                            if (((v4 + 1e10) - v4) == v4) {
                                v19 := []int{(fi16 % (((fi15 / ((9223372036854775807) | 1))) | 1))}
                                if true {
                                    var v20 uint8 = 2
                                    v4 = -2.25
                                    var v21 uint8 = (v18 - v20)
                                    _ = v20
                                    _ = v21
                                }
                                var v22 uint64 = v17
                                _ = v19
                                _ = v22
                            } else {
                                v23 := []int{(fi15 * fi16)}
                                var v24 bool = (false && (995 == int32(v17)))
                                _ = v23
                                _ = v24
                            }
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                    }
                    fmt.Println(v4, fi15, fi15)
                    _ = fi15
                }
                var v25 int64 = 0
                _ = v25
            }
        }
    default:
        var v26 uint32 = 392
        _ = v26
    }
    switch 2147483647 {
    case 128:
        if false {
            fmt.Println(v6, v1)
        }
    case 32768:
        v27 := GenS1{F0: 255, F1: 0}
        _ = v27
    default:
        for fi28 := 0; fi28 < 6; fi28++ {
            v29 := []int{fi28}
            _ = fi28
            _ = v29
        }
    }
    v30 := []int{}
    v2[2147483648] = ("\n\t" + "foo bar")
    fmt.Println(v4, v4, v6)
    delete(v3, 0)
    var v31 uint32 = 3
    _ = len(v30)
    fmt.Println(v1)
    var v32 uint64 = 309
    v4 -= (v4 * (1.5 * float64(3.14159)))
    for fi33 := 0; fi33 < 11; fi33++ {
        for fi34 := 0; fi34 < 20; fi34++ {
            var v35 float32 = 1.5
            _ = fi34
            _ = v35
        }
        _ = fi33
    }
    v31 = v31
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v6
}
