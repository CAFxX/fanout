package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 string, p1 int8, p2 int8) uint8 {
    if (!(0.0 > 1e10)) {
        v5 := GenS1{F0: 42, F1: 551}
        switch p0 {
        case "\n\t":
            for fi6 := 0; fi6 < 18; fi6++ {
                switch "a" {
                case "\n\t":
                    v7 := []int{947}
                    _ = v7
                default:
                    v8 := []int{-1, fi6}
                    _ = v8
                }
                p0 = p0
                var v9 int16 = 32767
                _ = fi6
                _ = v9
            }
        }
        for fi10 := 0; fi10 < 18; fi10++ {
            p1 += (p2 * (p2 - (p2 - p2)))
            p1 = p2
            v11 := GenS1{F0: 599, F1: 424}
            _ = fi10
            _ = v11
        }
        _ = v5
    }
    v12 := GenS1{F0: 1, F1: 42}
    fmt.Println(p1, p1)
    var v13 uint = 42
    p1 = (p2 | p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return 255
}
func genFn1(p0 int8, p1 uint8) int8 {
    if false {
        p0 = (p0 - (p0 | (p0 | (p0 - p0))))
    }
    p0 = (p0 | p0)
    p1 += 154
    var v14 uint8 = genFn0("", (p0 | p0), p0)
    _ = p0
    _ = p1
    _ = v14
    return (p0 / (((p0 - p0)) | 1))
}

func main() {
    var v1 uint8 = 152
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]string)
    var v4 uint16 = 713
    v15 := []int{238, 32768, 341, 2147483647}
    v16 := GenS2{F0: 859}
    fmt.Println(v1)
    if len(v15) > 0 {
        v15[0] = -834
    }
    switch -2147483648 {
    case 0:
        v17 := GenS2{F0: 4294967295}
        _ = v17
    case -9223372036854775808:
        switch ("" + "") {
        case "foo bar":
            if (104 <= -87) {
                fmt.Println(v4, v4)
                if (-2147483648 < -1) {
                    v1 += (v1 << ((2) & 63))
                } else {
                    v4 += (v4 - 1)
                    var v18 uint64 = 311
                    for fi19 := 0; fi19 < 13; fi19++ {
                        if (!(-1 != -1)) {
                            v1 = 2
                        } else {
                            switch "a" {
                            case "hello":
                                if len(v15) >= 2 {
                                    v15 = v15[1:len(v15)]
                                }
                            case "foo bar":
                                var v20 float32 = -2.25
                                _ = v20
                            }
                        }
                        _ = fi19
                    }
                    _ = v18
                }
            } else {
                if (-2147483648 == 1) {
                    _ = len(v3)
                } else {
                    v21 := v2[413]
                    _ = v21
                }
            }
        default:
            v1 -= v1
        }
    }
    if false {
        if (65535 != 42) {
            v4 = v4
            fmt.Println(v4, v4)
        } else {
            var v22 int = 127
            _ = v22
        }
        if (!((-661 < 42) && false)) {
            for fi23 := 0; fi23 < 12; fi23++ {
                v4 += (v4 << ((418) & 63))
                v24 := GenS1{F0: -1, F1: v4}
                _ = fi23
                _ = v24
            }
        }
    }
    var v25 uint8 = (v1 & (^(v1)))
    v25 = v1
    var v26 uint16 = v4
    v25 = (v1 % ((genFn0(("foo bar" + "a"), 100, -1)) | 1))
    v27 := GenS2{F0: 42}
    v3[9223372036854775807] = ""
    for fi28 := 0; fi28 < 8; fi28++ {
        var v29 uint32 = 65535
        var v30 int64 = 0
        v2[2147483647] = (v29 >> ((662) & 63))
        _ = fi28
        _ = v29
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
