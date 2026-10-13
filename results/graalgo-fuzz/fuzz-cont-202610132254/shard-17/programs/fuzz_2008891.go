package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 uint32; F1 int16 }
func genFn0(p0 int64) uint {
    p0 = -1
    p0 += (-(p0))
    v4 := []int{-324, -126, 242}
    _ = p0
    _ = v4
    return 0
}
func genFn1(p0 bool) uint32 {
    v5 := GenS1{F0: 1, F1: 0}
    switch "foo bar" {
    case "foo bar":
        p0 = ((0 < 59) && (0 != 42))
    default:
        fmt.Println(p0)
    }
    if p0 {
        v6 := GenS2{F0: 2147483647, F1: 1}
        v7 := []int{-9223372036854775808, 42, -1, 987}
        _ = v6
        _ = v7
    } else {
        v8 := []int{-551, 472}
        fmt.Println(p0)
        _ = v8
    }
    if (42 <= 65535) {
        var v9 int16 = 1
        _ = v9
    }
    var v10 int64 = 0
    _ = p0
    _ = v10
    _ = v5
    return 698
}

func main() {
    var v1 uint32 = 42
    var v2 uint64 = 925
    v3 := make(map[uint32]string)
    v1 -= 4294967295
    fmt.Println(v2)
    var v11 int8 = 100
    var v12 uint = genFn0(1)
    if false {
        switch 735 {
        case -9223372036854775808:
            v1 += 697
        }
    } else {
        v13 := []int{1, 38, 267}
        fmt.Println(v2)
        _ = v13
    }
    var v14 int64 = -1
    v15 := []int{-1, -2147483648, 9223372036854775807, 32767}
    switch "hello" {
    case "hello":
        v3[4294967295] = "a"
    case "foo bar":
        for fi16 := 0; fi16 < 16; fi16++ {
            v17 := GenS2{F0: v1, F1: 873}
            v1 = genFn1(false)
            if true {
                if (!(42 > 754)) {
                    fmt.Println(v14, v2)
                    fmt.Println(v1, v12, v1)
                    if (fi16 <= fi16) {
                        v14 -= v14
                        v1 = v1
                    }
                }
                for fi18 := 0; fi18 < 9; fi18++ {
                    if (false && (2 == 0)) {
                        _ = len(v3)
                    }
                    switch fi16 {
                    case -2147483648:
                        switch ("foo bar" + "a") {
                        case "a":
                            switch (fi18 / ((int(v2)) | 1)) {
                            case -9223372036854775808:
                                v14 = int64(v11)
                            case -864:
                                delete(v3, v1)
                            }
                        }
                    case -74:
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                    case 346:
                        var v19 string = (("a" + "\n\t") + ("foo bar" + "a"))
                        _ = v19
                    }
                    fmt.Println(fi16, v12, v12)
                    _ = fi18
                }
            } else {
                v14 = v14
                v20 := GenS2{F0: v1, F1: -378}
                for fi21 := 0; fi21 < 10; fi21++ {
                    var v22 string = ("hello" + "hello")
                    fmt.Println(v12, fi21, v22)
                    v23 := GenS2{F0: v1, F1: 42}
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                _ = v20
            }
            _ = fi16
            _ = v17
        }
    }
    v1 -= (v1 ^ (v1 - (^(v1))))
    var v24 uint16 = 591
    delete(v3, v1)
    v12 = v12
    v12 = (v12 % ((v12) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v3
}
