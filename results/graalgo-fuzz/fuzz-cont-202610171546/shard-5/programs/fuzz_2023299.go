package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
func genFn0() int64 {
    var v6 float32 = 3.14159
    v7 := []int{2147483647, -2147483648, 2147483647}
    _ = v6
    _ = v7
    return 124
}

func main() {
    var v1 float64 = 0.1
    v2 := make(map[int]string)
    v3 := make(map[string]uint32)
    v4 := make(map[string]string)
    var v5 uint = 430
    if (1 > -32768) {
        v1 = (((v1 / 3.14159) + (v1 / float64(((3.14159 - -2.25) + 1e10)))) / (v1 - 3.14159))
    }
    v8 := []int{}
    _ = len(v8)
    delete(v3, "\n\t")
    for fi9 := 0; fi9 < 20; fi9++ {
        v5 -= v5
        v10 := GenS1{F0: fi9, F1: 1}
        var v11 float32 = -2.25
        _ = fi9
        _ = v10
        _ = v11
    }
    v1 -= float64(0.1)
    fmt.Println(v1)
    if true {
        if ((v5 >> (((v5 << ((v5) & 63))) & 63)) <= (v5 << ((305) & 63))) {
            v1 = (1.5 + v1)
        } else {
            for fi12 := 0; fi12 < 3; fi12++ {
                for fi13 := 0; fi13 < 5; fi13++ {
                    _ = len(v8)
                    if (41 != 9223372036854775807) {
                        if (!(!(((32767 < 65535) && (((0 <= 2) || false) || (255 == 65535))) || false))) {
                            switch (fi13 ^ fi13) {
                            case 2147483647:
                                fmt.Println(fi12)
                            case 32768:
                                v1 = 1e10
                            default:
                                var v14 int = 2147483647
                                _ = v14
                            }
                            fmt.Println(fi13)
                        }
                        if ((2147483648 != 585) && true) {
                            var v15 uint8 = 1
                            var v16 int32 = 42
                            _ = v15
                            _ = v16
                        }
                        v17 := GenS1{F0: (fi12 + (-(fi13))), F1: 42}
                        var v18 int16 = -1
                        _ = v17
                        _ = v18
                    } else {
                        v5 -= (v5 * v5)
                        var v19 bool = (true || (false && (263 < 4294967295)))
                        _ = v19
                    }
                    _ = len(v8)
                    var v20 uint = v5
                    _ = fi13
                    _ = v20
                }
                v8 = append(v8, (fi12 - (fi12 * 128)))
                _ = len(v8)
                _ = fi12
            }
            if len(v8) > 0 {
                v8[0] = 127
            }
        }
        switch (("hello" + "hello") + "") {
        case "foo bar":
            var v21 int32 = -919
            _ = v21
        case "a":
            v4[((("hello" + "a") + "") + "\n\t")] = ("hello" + ("hello" + "hello"))
        default:
            v5 = v5
        }
        var v22 string = "foo bar"
        _ = v22
    } else {
        var v23 int = 1
        if ((v23 + v23) >= v23) {
            var v24 string = "\n\t"
            _ = v24
        }
        _ = v23
    }
    if true {
        v5 -= (v5 % ((882) | 1))
        v5 += 80
    }
    if len(v8) > 0 {
        v8[0] = 1
    }
    if true {
        for fi25 := 0; fi25 < 14; fi25++ {
            fmt.Println(v5, v1)
            _ = fi25
        }
    }
    for fi26 := 0; fi26 < 11; fi26++ {
        var v27 int8 = 1
        _ = fi26
        _ = v27
    }
    switch -101 {
    case -9223372036854775808:
        for fi28 := 0; fi28 < 19; fi28++ {
            var v29 bool = false
            switch "" {
            case "a":
                var v30 uint8 = 42
                _ = v30
            }
            _ = fi28
            _ = v29
        }
    default:
        switch 4294967295 {
        case -654:
            switch 0 {
            case 42:
                v8 = append(v8, 296)
            case 32768:
                v1 = 1e10
            case -965:
                v5 -= 894
            }
        default:
            v1 -= (0.0 + (v1 - v1))
        }
    }
    var v31 uint32 = 2147483648
    _ = len(v8)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v8
}
