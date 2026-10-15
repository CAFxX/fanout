package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8 }
func genFn0(p0 uint32, p1 bool) uint8 {
    p1 = (!(p1 && p1))
    switch 128 {
    case 32768:
        var v7 uint32 = 1
        _ = v7
    case 1:
        var v8 uint = 255
        _ = v8
    }
    var v9 bool = (1 == -1)
    v10 := GenS1{F0: v9, F1: 1, F2: 1}
    var v11 uint = 4294967295
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return 127
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]int)
    var v3 uint64 = 4294967295
    var v4 int64 = 1
    v5 := []int{0}
    var v6 float64 = float64(0.1)
    v12 := GenS2{F0: 0}
    v6 -= v6
    v5 = append(v5, 2147483647)
    v3 = v3
    delete(v1, -964)
    switch 32768 {
    case 907:
        var v13 float32 = (float32(v6) + -2.25)
        _ = v13
    case 127:
        fmt.Println(v6)
    default:
        fmt.Println(v6)
    }
    switch "a" {
    case "a":
        for fi14 := 0; fi14 < 9; fi14++ {
            if true {
                v6 -= v6
            }
            if true {
                v15 := GenS3{F0: 57}
                v4 = v4
                _ = v15
            } else {
                switch fi14 {
                case -184:
                    v4 = (v4 * v4)
                case -2147483648:
                    switch ("hello" + "foo bar") {
                    case "foo bar":
                        fmt.Println(v6)
                    case "":
                        if true {
                            for fi16 := 0; fi16 < 4; fi16++ {
                                if false {
                                    if true {
                                        v6 = (float64((1.5 / 1.5)) + ((v6 * (v6 + v6)) * v6))
                                    }
                                }
                                _ = fi16
                            }
                            v17 := GenS3{F0: -128}
                            _ = v17
                        } else {
                            fmt.Println(v4, v4)
                            v18 := GenS1{F0: false, F1: (fi14 % ((-129) | 1)), F2: 215}
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = v18
                        }
                    default:
                        v19 := GenS2{F0: -935}
                        _ = v19
                    }
                case 32767:
                    if true {
                        var v20 bool = (false && false)
                        _ = v20
                    }
                }
            }
            fmt.Println(v4, v6, v6)
            _ = fi14
        }
    case "hello":
        var v21 uint8 = genFn0(2147483647, (386 != 2147483647))
        _ = v21
    default:
        v3 += (v3 + (v3 % (((v3 % (((v3 + 4294967295)) | 1))) | 1)))
    }
    var v22 uint16 = 65535
    v23 := GenS3{F0: -1}
    fmt.Println(v6)
    var v24 float64 = v6
    _ = len(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
