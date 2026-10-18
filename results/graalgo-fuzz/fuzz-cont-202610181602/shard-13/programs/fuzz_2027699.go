package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 int; F1 uint64; F2 int8 }
type GenS3 struct { F0 bool; F1 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int8; F1 uint8; F2 bool }

func main() {
    v1 := make(map[int]string)
    var v2 int16 = 42
    var v3 float32 = -2.25
    var v4 uint16 = 729
    var v5 float32 = v3
    v6 := make(map[string]int)
    v7 := GenS1{F0: 1, F1: 1}
    switch "\n\t" {
    case "foo bar":
        v4 += 1
    }
    v3 = -2.25
    var v8 uint32 = 4294967295
    v9 := []int{1, -1, 4294967295}
    v8 -= 2
    v10 := GenS1{F0: 18446744073709551615, F1: -2147483648}
    switch ("hello" + ("foo bar" + "a")) {
    case "hello":
        switch "\n\t" {
        case "foo bar":
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
        }
    case "foo bar":
        switch 174 {
        case -607:
            v11 := GenS2{F0: 32768, F1: 148, F2: -63}
            _ = v11
        case 2147483647:
            if true {
                for fi12 := 0; fi12 < 10; fi12++ {
                    if true {
                        v8 = (v8 >> ((255) & 63))
                        if true {
                            for fi13 := 0; fi13 < 6; fi13++ {
                                fmt.Println(v8, v5, v4)
                                if true {
                                    if (0 != 9223372036854775807) {
                                        var v14 uint8 = 67
                                        v14 = v14
                                        v5 = v3
                                        _ = v14
                                    }
                                    v3 -= 1.5
                                }
                                var v15 float64 = float64(((v3 - 1e10) / v3))
                                _ = fi13
                                _ = v15
                            }
                            delete(v1, -901)
                            v2 -= (v2 & v2)
                        } else {
                            v16 := GenS3{F0: true, F1: (v4 >> ((42) & 63))}
                            fmt.Println(v2, v5)
                            _ = v16
                        }
                        if (789 > 2147483647) {
                            if false {
                                var v17 uint8 = 42
                                var v18 float64 = 1.5
                                if len(v9) > 0 {
                                    v9[0] = (fi12 & (fi12 - fi12))
                                }
                                _ = v17
                                _ = v18
                            }
                            var v19 int = fi12
                            _ = v19
                        }
                    }
                    if ((!true) || true) {
                        fmt.Println(v4, v3)
                        v20 := GenS4{F0: (((("hello" + (("foo bar" + "a") + "")) + ("a" + "foo bar")) + (("foo bar" + "") + "foo bar")) + "hello")}
                        _ = v20
                    }
                    var v21 uint8 = 1
                    _ = fi12
                    _ = v21
                }
                if (42 != -1) {
                    v22 := GenS5{F0: 100, F1: 255, F2: (-217 <= -717)}
                    v4 -= (v4 >> ((1) & 63))
                    var v23 int = 32767
                    _ = v22
                    _ = v23
                } else {
                    switch "foo bar" {
                    case "foo bar":
                        switch "\n\t" {
                        case "a":
                            if (4294967295 >= 18446744073709551615) {
                                switch "" {
                                case "foo bar":
                                    v4 = 218
                                case "a":
                                    fmt.Println(v5, v5)
                                }
                                v1[94] = ("\n\t" + "a")
                            } else {
                                _ = len(v9)
                                var v24 int64 = 0
                                for fi25 := 0; fi25 < 9; fi25++ {
                                    v3 -= v5
                                    _ = fi25
                                }
                                _ = v24
                            }
                        default:
                            fmt.Println(v5, v2, v3)
                        }
                    }
                }
            }
        default:
            var v26 uint64 = 42
            _ = v26
        }
    case "a":
        var v27 float64 = ((0.1 / -2.25) * 3.14159)
        _ = v27
    default:
        v9 = append(v9, -9223372036854775808)
    }
    v8 = v8
    if false {
        delete(v6, "a")
    } else {
        v4 = v4
        v28 := GenS1{F0: 42, F1: -840}
        _ = v28
    }
    v5 += 1e10
    fmt.Println(v8, v5)
    var v29 float32 = 0.0
    v2 += v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
