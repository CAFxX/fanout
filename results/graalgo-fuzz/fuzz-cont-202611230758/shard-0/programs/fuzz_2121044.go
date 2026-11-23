package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
func genFn0() uint16 {
    v7 := []int{32768, 127, 2147483647, 9223372036854775807}
    var v8 uint8 = 128
    switch "\n\t" {
    case "":
        var v9 uint32 = 4294967295
        _ = v9
    default:
        if len(v7) > 0 {
            v7[0] = 708
        }
    }
    var v10 uint16 = 65535
    _ = v10
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 uint64 = 925
    v2 := make(map[int]bool)
    v3 := []string{"foo bar", "\n\t"}
    var v4 int16 = 1000
    var v5 float32 = -2.25
    var v6 uint64 = v1
    v11 := GenS1{F0: 922, F1: (v1 - 2)}
    var v12 int8 = -1
    v13 := GenS1{F0: 4294967295, F1: (v1 ^ v6)}
    v4 -= 1
    v5 += v5
    v14 := GenS1{F0: 1, F1: v1}
    if false {
        v12 += 121
        fmt.Println(v5, v5, v6)
        fmt.Println(v1, v5, v5)
    }
    v15 := GenS1{F0: 1, F1: v1}
    if (1 == 2147483647) {
        var v16 int8 = (v12 ^ v12)
        var v17 int8 = v12
        fmt.Println(v5, v16)
        _ = v16
        _ = v17
    } else {
        v1 = 18446744073709551615
        v2[1] = (65535 <= 291)
    }
    v5 -= 1e10
    if (2147483647 != 42) {
        v1 -= (v1 >> ((2) & 63))
        switch ("a" + ("hello" + ("a" + ("" + "a")))) {
        case "foo bar":
            switch 127 {
            case -9223372036854775808:
                if true {
                    v1 += v1
                    v4 -= v4
                    fmt.Println(v5, v4)
                } else {
                    for fi18 := 0; fi18 < 3; fi18++ {
                        for fi19 := 0; fi19 < 2; fi19++ {
                            var v20 bool = true
                            _ = fi19
                            _ = v20
                        }
                        v21 := v2[(fi18 % ((fi18) | 1))]
                        _ = fi18
                        _ = v21
                    }
                    var v22 uint64 = v1
                    v23 := GenS1{F0: 1, F1: v22}
                    _ = v22
                    _ = v23
                }
            }
        case "hello":
            v2[9223372036854775807] = true
        }
        switch 0 {
        case -177:
            var v24 string = ("" + ("foo bar" + ("hello" + "a")))
            _ = v24
        case 4294967295:
            v5 = ((v5 * (3.14159 + 1e10)) - 0.1)
        default:
            switch -985 {
            case 219:
                v6 += v1
            case 0:
                var v25 uint8 = 128
                _ = v25
            default:
                for fi26 := 0; fi26 < 1; fi26++ {
                    var v27 uint32 = 2147483647
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if (false && ((v27 ^ 818) != (v27 ^ 607))) {
                        if true {
                            v3 = append(v3, "\n\t")
                            v28 := GenS1{F0: 255, F1: (v6 * 0)}
                            v3 = append(v3, (("\n\t" + "\n\t") + "a"))
                            _ = v28
                        } else {
                            fmt.Println(v1, v1, v6)
                            for fi29 := 0; fi29 < 17; fi29++ {
                                v30 := GenS1{F0: uint(v27), F1: (v6 * (v6 >> ((862) & 63)))}
                                var v31 uint32 = 2147483648
                                _ = fi29
                                _ = v30
                                _ = v31
                            }
                            v1 = 2
                        }
                    } else {
                        var v32 int8 = v12
                        v5 = (v5 * (1e10 + ((v5 / float32(1.5)) - float32(-2.25))))
                        _ = v32
                    }
                    _ = fi26
                    _ = v27
                }
            }
        }
    } else {
        fmt.Println(v5, v6)
        _ = len(v3)
    }
    v33 := GenS1{F0: 931, F1: (v1 & 4294967295)}
    for fi34 := 0; fi34 < 13; fi34++ {
        var v35 float64 = 0.1
        var v36 int64 = 9223372036854775807
        _ = fi34
        _ = v35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
