package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 uint }
func genFn0() uint32 {
    var v6 float32 = float32(float64(1.5))
    v7 := GenS1{F0: -1, F1: 4294967295}
    var v8 uint16 = 1
    fmt.Println(v6, v8)
    var v9 int64 = 0
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 4294967295
}

func main() {
    v1 := []int{}
    v2 := make(map[int]uint32)
    var v3 uint32 = 0
    var v4 int = 32768
    var v5 uint16 = 32767
    v10 := GenS1{F0: 1, F1: 627}
    if false {
        if false {
            if true {
                switch "a" {
                case "":
                    if len(v1) > 0 {
                        v1[0] = v4
                    }
                case "a":
                    fmt.Println(v3, v4, v4)
                }
                for fi11 := 0; fi11 < 12; fi11++ {
                    if true {
                        v5 = (v5 + (v5 - 65535))
                        v12 := v2[fi11]
                        var v13 bool = true
                        _ = v12
                        _ = v13
                    } else {
                        if (true && true) {
                            v14 := GenS1{F0: 2147483647, F1: uint64(v5)}
                            var v15 int = fi11
                            _ = v14
                            _ = v15
                        }
                        v16 := GenS2{F0: true}
                        _ = v16
                    }
                    for fi17 := 0; fi17 < 4; fi17++ {
                        switch "hello" {
                        case "hello":
                            switch "foo bar" {
                            case "foo bar":
                                switch fi11 {
                                case 963:
                                    v3 = (v3 << ((810) & 63))
                                default:
                                    v18 := GenS2{F0: true}
                                    _ = v18
                                }
                            }
                        case "\n\t":
                            v5 += (v5 & (v5 * v5))
                        }
                        v19 := GenS2{F0: (!(!true))}
                        v5 = 492
                        _ = fi17
                        _ = v19
                    }
                    v4 = v4
                    _ = fi11
                }
            } else {
                v5 -= 255
            }
            var v20 uint32 = genFn0()
            _ = v20
        } else {
            v21 := GenS2{F0: true}
            _ = v21
        }
        fmt.Println(v5, v3)
    } else {
        var v22 uint32 = genFn0()
        _ = v22
    }
    fmt.Println(v3)
    var v23 int16 = -738
    var v24 int = 4294967295
    var v25 int32 = int32(v23)
    v5 = v5
    v26 := GenS3{F0: 18446744073709551615, F1: uint(v3)}
    if false {
        var v27 float64 = 0.0
        for fi28 := 0; fi28 < 20; fi28++ {
            v29 := GenS2{F0: false}
            _ = fi28
            _ = v29
        }
        _ = v27
    } else {
        _ = len(v1)
        if (32767 == (-(v23))) {
            var v30 string = "a"
            v1 = append(v1, (v24 * (v24 / ((127) | 1))))
            _ = v30
        }
        for fi31 := 0; fi31 < 9; fi31++ {
            v25 -= v25
            _ = fi31
        }
    }
    var v32 float32 = 0.0
    v33 := GenS3{F0: 1, F1: 851}
    var v34 int16 = (v23 >> ((4294967295) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
