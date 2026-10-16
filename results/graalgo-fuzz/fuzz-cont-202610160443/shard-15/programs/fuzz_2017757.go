package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]int)
    v3 := make(map[int]string)
    var v4 string = "a"
    for fi5 := 0; fi5 < 8; fi5++ {
        fmt.Println(fi5, v4)
        _ = fi5
    }
    for fi6 := 0; fi6 < 5; fi6++ {
        v7 := GenS1{F0: false}
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: true}
    v9 := GenS1{F0: false}
    if true {
        switch (v4 + v4) {
        case "\n\t":
            v4 = v4
        case "":
            v4 = "\n\t"
        case "a":
            v10 := GenS1{F0: true}
            _ = v10
        default:
            var v11 uint16 = 42
            _ = v11
        }
    } else {
        v12 := []int{412, 728, 32768, 566}
        v4 = (v4 + "")
        _ = v12
    }
    fmt.Println(v4)
    _ = len(v3)
    for fi13 := 0; fi13 < 17; fi13++ {
        fmt.Println(fi13, v4)
        if (-1 <= -9223372036854775808) {
            v4 = v4
            if (1 != 27) {
                var v14 uint64 = 326
                v4 = "a"
                _ = v14
            }
        } else {
            if (false && ((false || (32767 != 653)) || (!false))) {
                var v15 int16 = 1
                v16 := []int{(fi13 >> ((65535) & 63)), fi13, fi13, 128}
                _ = v15
                _ = v16
            } else {
                v17 := []int{-2147483648, fi13, (fi13 >> ((774) & 63))}
                v4 = v4
                _ = v17
            }
            v4 = "a"
            var v18 bool = ((!false) && false)
            _ = v18
        }
        _ = fi13
    }
    v4 = v4
    var v19 int16 = 1
    v19 -= (v19 | (v19 | v19))
    switch 32767 {
    case 656:
        v4 = v4
    case 32767:
        fmt.Println(v19)
    default:
        if true {
            var v20 uint64 = 383
            for fi21 := 0; fi21 < 9; fi21++ {
                if (9223372036854775807 > 1) {
                    v22 := []int{fi21, (^(fi21)), (fi21 << ((0) & 63)), (fi21 / ((fi21) | 1))}
                    v4 = ((v4 + v4) + v4)
                    _ = v22
                } else {
                    v19 = v19
                }
                if true {
                    if false {
                        v23 := GenS1{F0: true}
                        v24 := []int{128, fi21, (fi21 * 128)}
                        fmt.Println(v19, v4)
                        _ = v23
                        _ = v24
                    }
                } else {
                    fmt.Println(fi21)
                    var v25 int32 = -370
                    _ = v25
                }
                _ = fi21
            }
            v4 = v4
            _ = v20
        } else {
            var v26 int32 = 0
            if (false && true) {
                var v27 uint64 = 221
                fmt.Println(v19)
                fmt.Println(v27, v27, v4)
                _ = v27
            } else {
                v28 := GenS1{F0: false}
                _ = v28
            }
            _ = v26
        }
    }
    v4 = v4
    fmt.Println(v19)
    v29 := GenS1{F0: true}
    var v30 uint16 = 139
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v8
    _ = v9
}
