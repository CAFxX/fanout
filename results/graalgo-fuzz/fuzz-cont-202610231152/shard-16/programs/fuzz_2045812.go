package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 string }
func genFn0(p0 bool, p1 uint32, p2 int64) uint {
    v4 := GenS1{F0: 0}
    var v5 uint16 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return uint(v5)
}
func genFn1(p0 bool, p1 uint64) int {
    switch 4294967295 {
    case 2147483647:
        var v6 string = ((("hello" + "a") + ("foo bar" + "\n\t")) + ("foo bar" + ("a" + "foo bar")))
        _ = v6
    case 416:
        v7 := []int{684, 1, 4294967295, 467}
        _ = v7
    }
    var v8 int = -572
    v9 := []int{v8, (v8 << ((255) & 63))}
    var v10 bool = (42 < 255)
    _ = p0
    _ = p1
    _ = v10
    _ = v8
    _ = v9
    return (v8 * v8)
}

func main() {
    var v1 int64 = -1
    var v2 uint = 0
    v3 := make(map[int]bool)
    v1 = -9223372036854775808
    fmt.Println(v2, v1)
    for fi11 := 0; fi11 < 5; fi11++ {
        var v12 uint32 = 42
        for fi13 := 0; fi13 < 13; fi13++ {
            fmt.Println(v1, v2)
            v14 := GenS2{F0: v2, F1: false}
            fmt.Println(v2)
            _ = fi13
            _ = v14
        }
        _ = fi11
        _ = v12
    }
    fmt.Println(v1, v1)
    v1 = (-(v1))
    for fi15 := 0; fi15 < 5; fi15++ {
        v2 = (v2 + (v2 | genFn0(false, 2147483648, (v1 - -9223372036854775808))))
        _ = fi15
    }
    v2 += 65535
    var v16 float64 = 0.0
    v17 := GenS2{F0: 65535, F1: true}
    v18 := GenS1{F0: 1}
    if false {
        if false {
            switch ("hello" + "a") {
            case "hello":
                for fi19 := 0; fi19 < 2; fi19++ {
                    v2 = (v2 ^ genFn0(((v16 - v16) > v16), 240, v1))
                    if (!(v1 <= -188)) {
                        fmt.Println(v1)
                        if (93 < 42) {
                            v20 := []int{(fi19 >> ((65535) & 63))}
                            if false {
                                if (v2 <= 2) {
                                    v21 := GenS1{F0: uint16(v2)}
                                    if true {
                                        for fi22 := 0; fi22 < 1; fi22++ {
                                            var v23 float64 = 1e10
                                            _ = fi22
                                            _ = v23
                                        }
                                        switch "hello" {
                                        case "":
                                            v24 := GenS3{F0: 77, F1: genFn0(false, 0, v1)}
                                            _ = v24
                                        case "foo bar":
                                            var v25 float64 = (v16 / (1.5 - (0.0 / v16)))
                                            _ = v25
                                        }
                                    }
                                    v2 = (v2 / ((genFn0(false, 713, (v1 / ((-9223372036854775808) | 1)))) | 1))
                                    _ = v21
                                } else {
                                    _ = len(v20)
                                    v16 -= v16
                                    fmt.Println(fi19, v1)
                                }
                            } else {
                                for fi26 := 0; fi26 < 9; fi26++ {
                                    var v27 uint8 = 255
                                    _ = fi26
                                    _ = v27
                                }
                            }
                            fmt.Println(v2)
                            _ = v20
                        }
                    }
                    _ = fi19
                }
            }
            for fi28 := 0; fi28 < 6; fi28++ {
                v29 := GenS3{F0: 0, F1: (v2 & (v2 & v2))}
                for fi30 := 0; fi30 < 17; fi30++ {
                    switch "\n\t" {
                    case "\n\t":
                        v2 -= v2
                    case "foo bar":
                        if ((1.5 - (float32(v16) * 3.14159)) < 1.5) {
                            v31 := GenS4{F0: "foo bar"}
                            _ = v31
                        } else {
                            var v32 uint16 = 42
                            var v33 uint32 = 744
                            _ = v32
                            _ = v33
                        }
                    default:
                        switch -9223372036854775808 {
                        case 10:
                            v34 := GenS2{F0: v2, F1: true}
                            _ = v34
                        }
                    }
                    v16 += v16
                    v16 += v16
                    _ = fi30
                }
                _ = fi28
                _ = v29
            }
            var v35 uint8 = 128
            _ = v35
        } else {
            var v36 float64 = v16
            fmt.Println(v36, v16, v36)
            v2 = v2
            _ = v36
        }
        var v37 float32 = 1e10
        _ = v37
    }
    for fi38 := 0; fi38 < 17; fi38++ {
        fmt.Println(v16, v16)
        var v39 uint16 = 101
        _ = fi38
        _ = v39
    }
    fmt.Println(v16)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
