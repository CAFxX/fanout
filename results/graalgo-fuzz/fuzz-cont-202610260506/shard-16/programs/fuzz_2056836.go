package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64; F1 uint; F2 int32 }
func genFn0(p0 int8, p1 int32) int8 {
    fmt.Println(p0, p0)
    v5 := []int{}
    _ = p0
    _ = p1
    _ = v5
    return (p0 >> ((65535) & 63))
}
func genFn1() int {
    v6 := GenS1{F0: (("a" + ("" + "a")) == "foo bar")}
    var v7 int8 = 42
    switch -121 {
    case 9223372036854775807:
        v8 := []int{-713, 4294967295, 32768, 9223372036854775807}
        _ = v8
    case 2147483647:
        for fi9 := 0; fi9 < 20; fi9++ {
            v10 := []int{-400, (fi9 >> ((255) & 63)), (fi9 & fi9)}
            var v11 int = fi9
            if (0.1 <= 1e10) {
                _ = len(v10)
            } else {
                if (v7 >= (v7 >> ((255) & 63))) {
                    for fi12 := 0; fi12 < 7; fi12++ {
                        if true {
                            v11 -= fi9
                        } else {
                            if (-289 < int64(v7)) {
                                v11 = (v11 - 42)
                            } else {
                                v11 = -760
                            }
                            var v13 int32 = 0
                            if false {
                                if (v13 == 0) {
                                    v14 := GenS2{F0: 65535}
                                    fmt.Println(v11)
                                    var v15 uint = 589
                                    var v16 string = "hello"
                                    _ = v14
                                    _ = v15
                                    _ = v16
                                }
                                v13 += (-(v13))
                                var v17 float32 = 0.1
                                _ = v17
                            }
                            _ = v13
                        }
                        _ = fi12
                    }
                } else {
                    v18 := GenS2{F0: 677}
                    for fi19 := 0; fi19 < 2; fi19++ {
                        fmt.Println(v11, fi19, fi19)
                        fmt.Println(v11, fi9)
                        v20 := GenS3{F0: -9223372036854775808, F1: 535, F2: -1}
                        _ = fi19
                        _ = v20
                    }
                    if false {
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                        for fi21 := 0; fi21 < 11; fi21++ {
                            v22 := GenS2{F0: 65535}
                            var v23 string = "hello"
                            if len(v10) > 0 {
                                v10[0] = (fi9 & -19)
                            }
                            _ = fi21
                            _ = v22
                            _ = v23
                        }
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                    } else {
                        v10 = append(v10, (v11 % (((v11 - (v11 + 42))) | 1)))
                        var v24 string = ("\n\t" + "hello")
                        _ = v24
                    }
                    _ = v18
                }
                var v25 uint32 = 889
                _ = v25
            }
            _ = fi9
            _ = v10
            _ = v11
        }
    case 1:
        var v26 uint = 65535
        _ = v26
    default:
        var v27 uint16 = 312
        _ = v27
    }
    for fi28 := 0; fi28 < 17; fi28++ {
        var v29 uint16 = 32767
        var v30 int8 = (-(v7))
        _ = fi28
        _ = v29
        _ = v30
    }
    v7 = v7
    _ = v6
    _ = v7
    return int(v7)
}

func main() {
    var v1 uint8 = 1
    var v2 uint = 1
    v3 := []uint32{1}
    v4 := make(map[string]bool)
    if false {
        fmt.Println(v1)
    }
    v1 = uint8(v2)
    v1 = v1
    var v31 float32 = (-2.25 / (((0.1 + (1e10 / 3.14159)) + 1.5) - float32(3.14159)))
    v31 = v31
    var v32 uint64 = 502
    switch "foo bar" {
    case "foo bar":
        var v33 float32 = (1e10 - float32(1e10))
        _ = v33
    case "hello":
        var v34 int = 9223372036854775807
        _ = v34
    default:
        var v35 int64 = -336
        _ = v35
    }
    switch 9223372036854775807 {
    case -545:
        var v36 uint8 = (v1 >> (((v2 ^ (v2 ^ (v2 * (v2 ^ (v2 / (((v2 << ((v2) & 63))) | 1))))))) & 63))
        _ = v36
    case 0:
        for fi37 := 0; fi37 < 9; fi37++ {
            if (v1 < 2) {
                var v38 int64 = 42
                _ = v38
            }
            var v39 int16 = int16(v1)
            _ = fi37
            _ = v39
        }
    case 9223372036854775807:
        v32 += v32
    default:
        v32 = (v32 ^ v32)
    }
    v4[""] = (!false)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
}
