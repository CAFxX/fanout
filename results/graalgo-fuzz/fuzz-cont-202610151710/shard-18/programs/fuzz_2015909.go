package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 int8 }

func main() {
    var v1 uint32 = 2147483648
    var v2 int8 = 30
    var v3 int16 = -1
    v4 := make(map[int]string)
    v5 := []int{32767, 4294967295}
    var v6 bool = (false && false)
    var v7 uint16 = 498
    v1 += v1
    if v6 {
        if true {
            if (!v6) {
                if len(v5) > 0 {
                    v5[0] = 1
                }
            }
        }
    }
    v6 = true
    if (0 >= -1) {
        var v8 uint = 65535
        if false {
            v9 := GenS1{F0: (((!v6) && (v7 <= (^(v7)))) || ((^(v8)) == (v8 ^ (v8 % ((v8) | 1)))))}
            _ = v9
        }
        fmt.Println(v6)
        _ = v8
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        v2 = (v2 >> ((65535) & 63))
        v2 = v2
        switch 2147483647 {
        case -9223372036854775808:
            switch "foo bar" {
            case "foo bar":
                if (!v6) {
                    v11 := v4[4294967295]
                    _ = v11
                } else {
                    v6 = (18 > 42)
                    v12 := GenS1{F0: v6}
                    if v6 {
                        v13 := GenS1{F0: ((!v6) && true)}
                        _ = len(v4)
                        _ = v13
                    }
                    _ = v12
                }
            case "hello":
                fmt.Println(v3, v7)
            case "\n\t":
                switch ("hello" + "\n\t") {
                case "":
                    if len(v5) > 0 {
                        v5[0] = 0
                    }
                case "a":
                    fmt.Println(fi10, v6)
                case "hello":
                    if (fi10 <= (-(fi10))) {
                        fmt.Println(v2, fi10)
                    } else {
                        v1 -= (v1 + v1)
                        if v6 {
                            v3 = v3
                            for fi14 := 0; fi14 < 15; fi14++ {
                                if (fi10 == fi14) {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    _ = len(v4)
                                }
                                v15 := GenS2{F0: false}
                                switch (fi14 * 4294967295) {
                                case 9223372036854775807:
                                    v5 = append(v5, (fi14 - 42))
                                case 128:
                                    fmt.Println(v2)
                                case 1:
                                    if (v6 && v6) {
                                        _ = len(v5)
                                        if (0 != -776) {
                                            var v16 int16 = (v3 * (v3 & v3))
                                            _ = v16
                                        }
                                    } else {
                                        v6 = v6
                                    }
                                default:
                                    fmt.Println(fi14, v6)
                                }
                                _ = fi14
                                _ = v15
                            }
                        } else {
                            if true {
                                var v17 uint64 = 1
                                _ = len(v5)
                                var v18 float32 = 3.14159
                                _ = v17
                                _ = v18
                            }
                            v19 := GenS2{F0: (42 >= 635)}
                            _ = v19
                        }
                    }
                default:
                    fmt.Println(fi10, v6, v7)
                }
            }
        case 127:
            v7 -= v7
        case 42:
            v4[fi10] = "\n\t"
        default:
            v3 += v3
        }
        _ = fi10
    }
    var v20 uint = 1
    v21 := GenS3{F0: v20, F1: 42}
    if (int32(v20) != 1) {
        _ = len(v5)
        v20 = (v20 | 1)
    } else {
        v22 := GenS2{F0: ((((v20 - (v20 << (((v20 - v20)) & 63))) > (v20 % ((42) | 1))) || (((-2147483648 != -2147483648) || true) && v6)) && (!v6))}
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
