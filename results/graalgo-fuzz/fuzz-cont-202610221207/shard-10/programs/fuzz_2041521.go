package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 string; F1 uint16; F2 int8 }

func main() {
    v1 := make(map[int]uint32)
    v2 := []int{4294967295, 112, 32768, 128}
    v3 := []int{-2147483648, -594, -647}
    var v4 int64 = 1
    switch 4294967295 {
    case 9223372036854775807:
        switch 610 {
        case 1:
            fmt.Println(v4)
        }
    case 4294967295:
        fmt.Println(v4)
    case 72:
        var v5 uint32 = 42
        _ = v5
    }
    v4 = (-(v4))
    if false {
        for fi6 := 0; fi6 < 18; fi6++ {
            v7 := GenS1{F0: 42}
            v4 = v4
            if false {
                fmt.Println(v4)
                if false {
                    switch fi6 {
                    case 2147483647:
                        if false {
                            for fi8 := 0; fi8 < 16; fi8++ {
                                if false {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    v4 -= v4
                                    var v9 int64 = v4
                                    switch "foo bar" {
                                    case "\n\t":
                                        if false {
                                            fmt.Println(v9)
                                        }
                                    }
                                    v4 += v4
                                    _ = v9
                                }
                                _ = fi8
                            }
                            var v10 uint16 = 255
                            v11 := GenS2{F0: fi6, F1: 0}
                            _ = v10
                            _ = v11
                        }
                    case 42:
                        v12 := GenS1{F0: 531}
                        _ = v12
                    case 601:
                        v4 += v4
                    }
                    for fi13 := 0; fi13 < 7; fi13++ {
                        var v14 uint32 = 735
                        _ = fi13
                        _ = v14
                    }
                    if (2 == 1) {
                        for fi15 := 0; fi15 < 17; fi15++ {
                            var v16 int = fi6
                            _ = fi15
                            _ = v16
                        }
                    }
                }
                v4 -= v4
            } else {
                v4 += (v4 / ((v4) | 1))
                for fi17 := 0; fi17 < 10; fi17++ {
                    if true {
                        fmt.Println(v4)
                        var v18 int32 = -1
                        _ = v18
                    }
                    v4 = v4
                    fmt.Println(fi17, fi17)
                    _ = fi17
                }
                var v19 int32 = 2147483647
                _ = v19
            }
            _ = fi6
            _ = v7
        }
    }
    var v20 uint16 = 32767
    if true {
        v21 := GenS1{F0: 109}
        v2 = append(v2, 1)
        _ = v21
    }
    v20 = v20
    v20 = v20
    v22 := GenS1{F0: 255}
    if false {
        fmt.Println(v20, v20)
        fmt.Println(v4, v4)
        if (!(339 > 255)) {
            if (!(0 != 11)) {
                var v23 uint = 18446744073709551615
                var v24 float64 = ((3.14159 * (((1.5 * 1.5) - 3.14159) - 1e10)) - 3.14159)
                v25 := GenS1{F0: (v23 ^ 225)}
                _ = v23
                _ = v24
                _ = v25
            }
        } else {
            if (!false) {
                switch ("foo bar" + "foo bar") {
                case "\n\t":
                    v20 = 338
                }
                fmt.Println(v20, v20)
                v20 = 1
            } else {
                v20 += (v20 % ((579) | 1))
            }
            if false {
                var v26 int64 = v4
                var v27 uint16 = (v20 >> ((826) & 63))
                _ = v26
                _ = v27
            }
            var v28 int32 = -1
            _ = v28
        }
    }
    var v29 bool = false
    var v30 uint32 = 2
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch 898 {
    case 1:
        for fi31 := 0; fi31 < 8; fi31++ {
            _ = len(v1)
            if false {
                var v32 uint8 = 42
                _ = v32
            } else {
                if v29 {
                    v20 = (v20 | 32767)
                    v33 := GenS3{F0: "hello", F1: v20, F2: 58}
                    var v34 uint8 = 42
                    _ = v33
                    _ = v34
                } else {
                    switch ("hello" + "") {
                    case "a":
                        fmt.Println(v20, v29)
                    default:
                        v35 := GenS1{F0: 815}
                        _ = v35
                    }
                }
            }
            _ = fi31
        }
    default:
        fmt.Println(v30, v30)
    }
    for fi36 := 0; fi36 < 17; fi36++ {
        v20 = 42
        var v37 string = ("a" + "\n\t")
        _ = fi36
        _ = v37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
