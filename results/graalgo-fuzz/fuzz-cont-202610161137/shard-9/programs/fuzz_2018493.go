package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 string; F1 string; F2 int16 }
type GenS3 struct { F0 bool; F1 int; F2 uint32 }
func genFn0() uint8 {
    var v4 bool = (65535 <= 255)
    v5 := []int{32768, 32767, -2147483648, 2147483647}
    v6 := GenS1{F0: -536, F1: 1}
    _ = v4
    _ = v5
    _ = v6
    return 42
}

func main() {
    v1 := []int{42, 127, 2147483647, 128}
    var v2 uint32 = 4294967295
    v3 := make(map[string]uint32)
    v7 := GenS2{F0: ("" + ""), F1: ("foo bar" + "\n\t"), F2: -1}
    v8 := GenS2{F0: (("" + "\n\t") + "foo bar"), F1: "\n\t", F2: -466}
    fmt.Println(v2)
    if false {
        v3["a"] = 2147483647
        fmt.Println(v2)
        v2 = 2147483648
    } else {
        fmt.Println(v2)
        v9 := GenS1{F0: 9223372036854775807, F1: 4294967295}
        if (true && (((2147483647 != 42) || (42 <= 42)) && true)) {
            if (!(1000 == 32767)) {
                v2 = 2
            } else {
                var v10 uint = 255
                _ = v10
            }
        } else {
            fmt.Println(v2)
        }
        _ = v9
    }
    v11 := GenS2{F0: "hello", F1: "a", F2: 0}
    v2 += v2
    v2 = (v2 * v2)
    v2 = 197
    var v12 float64 = 3.14159
    v13 := GenS1{F0: 590, F1: 969}
    var v14 float64 = 0.1
    var v15 uint64 = 42
    v12 = float64(1.5)
    switch 127 {
    case -564:
        fmt.Println(v15)
    case 9223372036854775807:
        delete(v3, ("foo bar" + ("a" + "hello")))
    case -9223372036854775808:
        var v16 int64 = -9223372036854775808
        _ = v16
    }
    if (!false) {
        fmt.Println(v15, v12, v15)
        var v17 uint16 = 2
        _ = v17
    } else {
        if true {
            for fi18 := 0; fi18 < 18; fi18++ {
                fmt.Println(v14, fi18)
                _ = fi18
            }
        } else {
            for fi19 := 0; fi19 < 4; fi19++ {
                switch "hello" {
                case "hello":
                    var v20 float32 = (0.0 * 3.14159)
                    _ = v20
                case "":
                    var v21 int16 = 1
                    _ = v21
                case "a":
                    v14 -= (3.14159 - (0.1 - v12))
                default:
                    switch fi19 {
                    case 127:
                        for fi22 := 0; fi22 < 12; fi22++ {
                            var v23 uint32 = v2
                            v12 -= v12
                            if (true || false) {
                                v24 := GenS3{F0: (42 > 42), F1: fi22, F2: v23}
                                if (!(true || false)) {
                                    v12 += 3.14159
                                }
                                _ = v24
                            }
                            _ = fi22
                            _ = v23
                        }
                    case -1:
                        v15 = (v15 >> ((4294967295) & 63))
                    }
                }
                _ = fi19
            }
        }
        v3["foo bar"] = (v2 & v2)
        var v25 int16 = -292
        _ = v25
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
