package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 bool }
func genFn0() string {
    var v4 string = ((("\n\t" + "foo bar") + ("hello" + "foo bar")) + "foo bar")
    fmt.Println(v4)
    switch 128 {
    case 9223372036854775807:
        fmt.Println(v4)
    case -771:
        var v5 string = ""
        _ = v5
    default:
        if (false && (!(127 >= -1))) {
            v6 := GenS1{F0: "hello"}
            _ = v6
        } else {
            for fi7 := 0; fi7 < 9; fi7++ {
                for fi8 := 0; fi8 < 13; fi8++ {
                    var v9 uint8 = 33
                    switch v4 {
                    case "":
                        v4 = "hello"
                    case "foo bar":
                        v9 += (v9 * (v9 ^ v9))
                    default:
                        fmt.Println(v9, fi7, v9)
                    }
                    _ = fi8
                    _ = v9
                }
                for fi10 := 0; fi10 < 17; fi10++ {
                    if (!true) {
                        fmt.Println(fi7, fi7, fi7)
                        var v11 uint64 = 42
                        _ = v11
                    } else {
                        v4 = v4
                    }
                    v12 := GenS1{F0: v4}
                    v4 = "hello"
                    _ = fi10
                    _ = v12
                }
                if (!true) {
                    v4 = (v4 + v4)
                } else {
                    var v13 int8 = 127
                    v4 = (v4 + "\n\t")
                    _ = v13
                }
                _ = fi7
            }
            var v14 uint64 = 2
            _ = v14
        }
    }
    v15 := GenS2{F0: 128, F1: (v4 + v4)}
    v4 = v4
    _ = v15
    _ = v4
    return (("a" + v4) + v4)
}

func main() {
    v1 := []int{-554, -2147483648, -2147483648, 1}
    var v2 uint32 = 65535
    var v3 int32 = -2147483648
    v16 := GenS1{F0: ("" + "\n\t")}
    if (-1 > -407) {
        var v17 int = 4294967295
        if len(v1) > 0 {
            v1[0] = int(v3)
        }
        _ = v17
    }
    v18 := GenS1{F0: "\n\t"}
    if true {
        var v19 bool = false
        v20 := GenS3{F0: (!(!(32767 == 1000)))}
        if v19 {
            var v21 uint16 = 0
            var v22 int = -610
            fmt.Println(v3, v19, v22)
            _ = v21
            _ = v22
        } else {
            _ = len(v1)
        }
        _ = v19
        _ = v20
    }
    v2 -= (v2 / ((v2) | 1))
    fmt.Println(v2, v3)
    if false {
        v2 += (v2 >> ((0) & 63))
    } else {
        switch 2147483647 {
        case -300:
            if (4294967295 > 110) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v1 = append(v1, 610)
                v3 += 42
            }
        case 32767:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    }
    v23 := GenS2{F0: 255, F1: "hello"}
    _ = len(v1)
    v2 = 711
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v23
    _ = v3
}
