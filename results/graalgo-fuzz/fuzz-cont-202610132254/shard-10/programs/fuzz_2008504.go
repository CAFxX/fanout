package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint; F1 bool }
func genFn0() int16 {
    var v5 uint16 = 65535
    var v6 float32 = 3.14159
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 bool) int32 {
    var v7 uint32 = 65535
    fmt.Println(v7, v7)
    p0 = (p0 && false)
    var v8 float64 = ((1.5 - 0.1) + 3.14159)
    _ = p0
    _ = v7
    _ = v8
    return -830
}

func main() {
    var v1 float32 = -2.25
    var v2 uint8 = 0
    var v3 float32 = v1
    v4 := make(map[int]uint32)
    var v9 uint64 = 1
    v3 = (((v1 / float32(1e10)) - float32(-2.25)) + (v1 - 0.1))
    if (0 > 65535) {
        var v10 uint8 = v2
        v11 := GenS1{F0: 234}
        _ = v10
        _ = v11
    } else {
        v9 = v9
        if (!(155 != 971)) {
            if true {
                switch ("\n\t" + ("" + "foo bar")) {
                case "":
                    v9 = 4294967295
                default:
                    v4[783] = 42
                }
            }
        }
    }
    var v12 bool = false
    if (670 != 18446744073709551615) {
        v1 = -2.25
        v3 -= v3
        if ((("\n\t" >= ("hello" + "a")) || v12) || (2147483647 <= 9223372036854775807)) {
            v3 -= 1.5
            for fi13 := 0; fi13 < 6; fi13++ {
                var v14 uint32 = 0
                _ = fi13
                _ = v14
            }
        } else {
            if true {
                switch 0 {
                case -1:
                    v2 += (v2 ^ 42)
                case 18:
                    for fi15 := 0; fi15 < 7; fi15++ {
                        v16 := GenS1{F0: 2}
                        var v17 uint32 = 2
                        if (0 > 255) {
                            v2 += 1
                        }
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                default:
                    v18 := GenS2{F0: 4294967295}
                    _ = v18
                }
            } else {
                if v12 {
                    v3 -= (-2.25 + v1)
                } else {
                    var v19 string = "a"
                    fmt.Println(v9, v1)
                    fmt.Println(v9, v2)
                    _ = v19
                }
            }
        }
    }
    var v20 int32 = 2147483647
    v21 := GenS2{F0: uint32(v2)}
    fmt.Println(v12, v2)
    for fi22 := 0; fi22 < 6; fi22++ {
        fmt.Println(v1, v3, v2)
        _ = fi22
    }
    if true {
        switch 128 {
        case 128:
            fmt.Println(v1, v12)
        }
        var v23 uint64 = (v9 * (v9 >> ((65535) & 63)))
        _ = v23
    }
    v24 := GenS1{F0: v2}
    for fi25 := 0; fi25 < 19; fi25++ {
        v26 := GenS3{F0: v12}
        _ = fi25
        _ = v26
    }
    for fi27 := 0; fi27 < 10; fi27++ {
        var v28 float64 = (0.1 * ((1.5 * 0.1) - -2.25))
        _ = fi27
        _ = v28
    }
    v12 = false
    switch "foo bar" {
    case "foo bar":
        v29 := []int{0, -540, -9223372036854775808, 276}
        _ = v29
    case "":
        v30 := GenS3{F0: (255 > 1)}
        _ = v30
    default:
        if (0 > 255) {
            if v12 {
                switch -9223372036854775808 {
                case 662:
                    v9 += v9
                default:
                    if false {
                        var v31 float32 = (0.0 - v1)
                        for fi32 := 0; fi32 < 1; fi32++ {
                            for fi33 := 0; fi33 < 20; fi33++ {
                                fmt.Println(v20)
                                _ = fi33
                            }
                            _ = fi32
                        }
                        for fi34 := 0; fi34 < 9; fi34++ {
                            var v35 int = (-(fi34))
                            _ = fi34
                            _ = v35
                        }
                        _ = v31
                    } else {
                        v2 -= v2
                    }
                }
                v36 := GenS4{F0: 255, F1: false}
                v37 := GenS2{F0: 2147483648}
                _ = v36
                _ = v37
            }
            fmt.Println(v9, v2, v1)
        }
    }
    for fi38 := 0; fi38 < 14; fi38++ {
        var v39 uint = 65535
        _ = fi38
        _ = v39
    }
    if v12 {
        v20 = (v20 >> ((0) & 63))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v9
}
