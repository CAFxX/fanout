package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int8 }
func genFn0(p0 uint16, p1 int32) string {
    switch 4294967295 {
    case 1:
        switch "\n\t" {
        case "":
            var v5 int16 = -590
            _ = v5
        case "hello":
            v6 := []int{-9223372036854775808, 1, 459, 2147483647}
            _ = v6
        }
    case 128:
        v7 := []int{-424}
        _ = v7
    default:
        fmt.Println(p0, p0)
    }
    fmt.Println(p0)
    p0 = 255
    _ = p0
    _ = p1
    return ("hello" + "a")
}

func main() {
    v1 := make(map[int]string)
    v2 := []string{"foo bar"}
    v3 := []uint32{2147483648, 2147483647}
    v4 := make(map[uint32]bool)
    var v8 int64 = -1
    v8 -= (v8 * v8)
    if true {
        fmt.Println(v8)
        var v9 int64 = (v8 % (((v8 | -251)) | 1))
        v8 = v8
        _ = v9
    }
    fmt.Println(v8)
    for fi10 := 0; fi10 < 3; fi10++ {
        if (329 == 42) {
            v11 := v4[2]
            _ = v11
        } else {
            for fi12 := 0; fi12 < 2; fi12++ {
                switch "a" {
                case "foo bar":
                    fmt.Println(v8)
                case "hello":
                    var v13 bool = true
                    _ = v13
                }
                v8 -= (v8 ^ -9223372036854775808)
                _ = fi12
            }
            if false {
                v14 := GenS1{F0: 127, F1: false}
                _ = v14
            }
        }
        fmt.Println(fi10)
        v8 -= 1
        _ = fi10
    }
    switch 4294967295 {
    case -195:
        v15 := GenS2{F0: true}
        _ = v15
    }
    v16 := GenS2{F0: false}
    if (0 < 4294967295) {
        v8 += (v8 % ((v8) | 1))
        v8 = (v8 | v8)
    }
    var v17 int8 = 1
    v18 := v4[1]
    for fi19 := 0; fi19 < 19; fi19++ {
        fmt.Println(fi19)
        _ = fi19
    }
    v4[340] = v18
    v8 = v8
    var v20 float32 = 3.14159
    _ = len(v4)
    switch "foo bar" {
    case "":
        if v18 {
            if (v17 == (v17 & v17)) {
                v17 = (v17 << ((42) & 63))
                switch -346 {
                case 3:
                    v21 := GenS3{F0: 0, F1: 42}
                    _ = v21
                case 4294967295:
                    var v22 int64 = (v8 << ((4294967295) & 63))
                    _ = v22
                default:
                    fmt.Println(v8)
                }
                _ = len(v2)
            } else {
                v23 := GenS4{F0: -2147483648}
                var v24 uint = 806
                for fi25 := 0; fi25 < 11; fi25++ {
                    v8 -= (v8 ^ -1)
                    switch ("foo bar" + "") {
                    case "hello":
                        for fi26 := 0; fi26 < 16; fi26++ {
                            if v18 {
                                _ = len(v2)
                                v27 := GenS3{F0: 2147483648, F1: 830}
                                if len(v2) > 0 {
                                    v2[0] = "a"
                                }
                                _ = v27
                            } else {
                                var v28 string = ("\n\t" + "\n\t")
                                _ = v28
                            }
                            v29 := GenS2{F0: (true || true)}
                            _ = fi26
                            _ = v29
                        }
                    default:
                        var v30 uint32 = 187
                        _ = v30
                    }
                    _ = fi25
                }
                _ = v23
                _ = v24
            }
            v31 := GenS5{F0: ((v20 * 1.5) > (0.0 + ((v20 + v20) - v20)))}
            _ = v31
        } else {
            for fi32 := 0; fi32 < 16; fi32++ {
                delete(v1, 1)
                v33 := GenS6{F0: v17}
                _ = fi32
                _ = v33
            }
            var v34 uint = 2
            var v35 int8 = (-(v17))
            _ = v34
            _ = v35
        }
    case "a":
        _ = len(v1)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v8
}
