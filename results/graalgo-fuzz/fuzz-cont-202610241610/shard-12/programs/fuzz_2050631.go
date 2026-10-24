package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int }
type GenS2 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 bool = (1 <= 2147483647)
    var v3 uint16 = 255
    var v4 string = ("foo bar" + "hello")
    fmt.Println(v3, v2)
    for fi5 := 0; fi5 < 1; fi5++ {
        v3 -= (v3 * 187)
        _ = fi5
    }
    v6 := GenS1{F0: 560, F1: (v3 ^ 1), F2: 0}
    v7 := []int{9223372036854775807, -9223372036854775808, -186}
    var v8 uint16 = (^(v3))
    if v2 {
        var v9 int = 4294967295
        for fi10 := 0; fi10 < 18; fi10++ {
            if v2 {
                v4 = "a"
            } else {
                var v11 string = v4
                _ = v11
            }
            fmt.Println(v8, v9, v8)
            if ((92 <= 113) || v2) {
                v1[4294967295] = 57
            }
            _ = fi10
        }
        _ = v9
    } else {
        switch v4 {
        case "\n\t":
            if (1e10 != 1e10) {
                v12 := GenS1{F0: -1, F1: 2, F2: 9223372036854775807}
                if (!v2) {
                    switch 2147483647 {
                    case -649:
                        v4 = ((v4 + v4) + (v4 + "hello"))
                    case 1:
                        var v13 uint = 0
                        _ = v13
                    case 32767:
                        v14 := GenS1{F0: 32768, F1: v3, F2: 76}
                        _ = v14
                    default:
                        var v15 uint32 = 0
                        _ = v15
                    }
                } else {
                    v1[2] = 42
                    v8 -= (v3 * v8)
                    switch -662 {
                    case -444:
                        if (((^(v8)) > (v3 * (v8 % ((1) | 1)))) || (-1 < 0)) {
                            var v16 bool = (v2 || (true && (1 >= -100)))
                            v7 = append(v7, -746)
                            _ = v16
                        }
                    case -9223372036854775808:
                        v7 = append(v7, -129)
                    case 128:
                        if v2 {
                            v17 := GenS1{F0: 2147483647, F1: v8, F2: -363}
                            v8 = (v3 + v8)
                            v18 := GenS1{F0: 43, F1: (v8 << ((0) & 63)), F2: -2147483648}
                            _ = v17
                            _ = v18
                        } else {
                            for fi19 := 0; fi19 < 3; fi19++ {
                                var v20 int8 = 1
                                _ = fi19
                                _ = v20
                            }
                            var v21 int32 = 273
                            for fi22 := 0; fi22 < 12; fi22++ {
                                var v23 bool = (32767 == -339)
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                                var v24 int8 = -1
                                _ = fi22
                                _ = v23
                                _ = v24
                            }
                            _ = v21
                        }
                    }
                }
                _ = v12
            }
        case "hello":
            if false {
                switch 2147483647 {
                case 127:
                    v7 = append(v7, 4294967295)
                }
                v3 = 987
                _ = len(v1)
            } else {
                delete(v1, 65535)
            }
        case "a":
            v4 = "a"
        }
        var v25 int = 0
        _ = v25
    }
    if ("hello" < (v4 + (v4 + ((("foo bar" + ("hello" + "")) + "foo bar") + ("hello" + "hello"))))) {
        for fi26 := 0; fi26 < 17; fi26++ {
            for fi27 := 0; fi27 < 6; fi27++ {
                _ = len(v7)
                _ = fi27
            }
            _ = fi26
        }
    }
    var v28 string = "foo bar"
    var v29 uint8 = 32
    fmt.Println(v4)
    if (!true) {
        v30 := GenS2{F0: 211}
        var v31 int32 = -595
        fmt.Println(v3)
        _ = v30
        _ = v31
    } else {
        var v32 int64 = 42
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
