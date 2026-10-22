package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int; F1 int64; F2 uint }

func main() {
    var v1 int = 2147483647
    v2 := []string{}
    var v3 int32 = 265
    var v4 uint8 = 133
    v5 := make(map[uint32]int)
    switch v1 {
    case 857:
        if false {
            for fi6 := 0; fi6 < 15; fi6++ {
                var v7 int8 = 72
                _ = fi6
                _ = v7
            }
        }
    case -950:
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 int16 = -1
            switch v1 {
            case -828:
                v10 := GenS1{F0: 0, F1: 515, F2: -1}
                _ = v10
            case 128:
                var v11 int64 = -569
                _ = v11
            case 32767:
                var v12 uint8 = (v4 + (v4 * (v4 << ((627) & 63))))
                _ = v12
            }
            v9 = (v9 >> ((774) & 63))
            _ = fi8
            _ = v9
        }
    case -214:
        var v13 int64 = -1
        _ = v13
    default:
        v14 := GenS1{F0: 42, F1: 979, F2: -1}
        _ = v14
    }
    v4 = (v4 * (v4 | v4))
    var v15 int64 = -1
    var v16 float32 = -2.25
    _ = len(v2)
    v17 := GenS1{F0: 416, F1: 864, F2: v15}
    fmt.Println(v15)
    var v18 int32 = (v3 * v3)
    v18 += (v3 % ((v3) | 1))
    _ = len(v2)
    if true {
        var v19 int = 32767
        _ = v19
    }
    switch "\n\t" {
    case "hello":
        _ = len(v5)
    case "a":
        switch ("a" + "hello") {
        case "\n\t":
            var v20 int32 = (v3 << ((uint(v4)) & 63))
            _ = v20
        case "":
            v21 := GenS1{F0: 2, F1: 2147483648, F2: v15}
            _ = v21
        case "foo bar":
            var v22 int64 = v15
            _ = v22
        }
    case "\n\t":
        v18 -= 1
    }
    if false {
        if (!true) {
            v23 := v5[2147483648]
            switch "\n\t" {
            case "foo bar":
                fmt.Println(v15, v1, v23)
            default:
                switch int(v4) {
                case 0:
                    var v24 uint16 = 2
                    _ = v24
                case -281:
                    switch v23 {
                    case 0:
                        fmt.Println(v3, v4)
                    case -439:
                        for fi25 := 0; fi25 < 6; fi25++ {
                            if true {
                                fmt.Println(v1)
                                v2 = append(v2, "")
                                var v26 uint64 = 18446744073709551615
                                _ = v26
                            } else {
                                if true {
                                    var v27 int16 = 0
                                    if len(v2) > 0 {
                                        v2[0] = "foo bar"
                                    }
                                    _ = v27
                                } else {
                                    fmt.Println(v3)
                                    v15 -= -489
                                }
                            }
                            v28 := v5[0]
                            _ = fi25
                            _ = v28
                        }
                    case 4294967295:
                        var v29 float32 = (float32(1e10) + float32(1.5))
                        _ = v29
                    default:
                        var v30 uint = 65535
                        _ = v30
                    }
                case 4294967295:
                    for fi31 := 0; fi31 < 15; fi31++ {
                        v32 := GenS2{F0: ((("" + "a") + (("" + ("\n\t" + "a")) + "foo bar")) + "")}
                        v33 := GenS3{F0: v1, F1: v15, F2: 255}
                        _ = fi31
                        _ = v32
                        _ = v33
                    }
                default:
                    _ = len(v2)
                }
            }
            _ = v23
        }
    } else {
        v1 += v1
    }
    v34 := GenS2{F0: ((("\n\t" + (("foo bar" + ("\n\t" + "foo bar")) + (("foo bar" + ("a" + "a")) + (("a" + "") + "a")))) + "\n\t") + "\n\t")}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
