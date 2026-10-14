package main

import (
    "fmt"
)

func genFn0(p0 int) uint32 {
    v7 := []int{2147483647, 4294967295, 791}
    var v8 uint64 = 4294967295
    _ = p0
    _ = v7
    _ = v8
    return 974
}
func genFn1() int64 {
    var v9 int16 = 1000
    v10 := []int{0, -728}
    _ = v10
    _ = v9
    return 9223372036854775807
}

func main() {
    v1 := []int{}
    var v2 uint = 255
    var v3 int16 = 240
    var v4 int32 = -1
    v5 := make(map[int]string)
    var v6 float32 = 3.14159
    var v11 int8 = 42
    var v12 int16 = -1
    var v13 uint64 = 18446744073709551615
    v2 = v2
    var v14 bool = (0 != 255)
    switch 128 {
    case 128:
        v2 += 2
    case 127:
        v6 = 3.14159
    case 4294967295:
        v11 -= v11
    }
    if v14 {
        var v15 int16 = (v3 >> ((v2) & 63))
        switch 875 {
        case 2147483647:
            delete(v5, 224)
        case -319:
            _ = len(v5)
        default:
            if v14 {
                if (65535 == 255) {
                    var v16 bool = v14
                    switch "hello" {
                    case "foo bar":
                        for fi17 := 0; fi17 < 20; fi17++ {
                            var v18 string = "\n\t"
                            v19 := v5[(fi17 / (((-(fi17))) | 1))]
                            if (255 != 83) {
                                switch "hello" {
                                case "\n\t":
                                    var v20 float64 = 1.5
                                    _ = v20
                                case "a":
                                    v16 = true
                                }
                            }
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                    default:
                        var v21 int8 = (v11 * (v11 ^ v11))
                        _ = v21
                    }
                    if (!(327 != -199)) {
                        var v22 int16 = v12
                        _ = v22
                    }
                    _ = v16
                }
                v4 = v4
            } else {
                v2 = (v2 / ((v2) | 1))
                fmt.Println(v3, v14, v15)
            }
        }
        var v23 uint64 = (v13 << (((v2 % (((v2 | (v2 % ((2) | 1)))) | 1))) & 63))
        _ = v15
        _ = v23
    }
    for fi24 := 0; fi24 < 5; fi24++ {
        var v25 int32 = -2147483648
        var v26 uint32 = 2147483648
        _ = fi24
        _ = v25
        _ = v26
    }
    var v27 int8 = v11
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
