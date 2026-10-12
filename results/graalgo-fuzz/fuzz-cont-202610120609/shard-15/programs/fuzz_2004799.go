package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 int8; F1 uint16; F2 uint32 }
type GenS4 struct { F0 int8; F1 int }
func genFn0() int32 {
    var v6 uint8 = 255
    var v7 int16 = 1
    _ = v6
    _ = v7
    return -2147483648
}

func main() {
    var v1 int8 = 1
    var v2 int = 4294967295
    var v3 bool = true
    var v4 float64 = (((3.14159 * (-2.25 - ((-2.25 - 0.1) - (float64(1e10) * 0.0)))) * (1e10 * 1.5)) / 1.0)
    v5 := make(map[uint32]int)
    v8 := GenS1{F0: v3, F1: 0}
    fmt.Println(v1, v4)
    v9 := GenS1{F0: (645 > 760), F1: int32(v1)}
    v3 = false
    var v10 string = ""
    for fi11 := 0; fi11 < 15; fi11++ {
        v3 = (false || (4294967295 <= 4294967295))
        if true {
            if (float64(3.14159) != (v4 - v4)) {
                v12 := v5[65535]
                var v13 uint16 = 65535
                _ = v12
                _ = v13
            }
        }
        switch ("\n\t" + v10) {
        case "hello":
            fmt.Println(v1, v1, v1)
        default:
            fmt.Println(v10, v4, fi11)
        }
        _ = fi11
    }
    if (0 == -1) {
        var v14 float32 = ((1e10 - 1.5) * 0.1)
        fmt.Println(v4, v10)
        v15 := []int{v2}
        _ = v14
        _ = v15
    }
    if (!(42 > -32768)) {
        v10 = v10
        v3 = (genFn0() > 42)
        v3 = v3
    }
    v3 = (v4 > (v4 * float64(0.1)))
    delete(v5, 2)
    var v16 int32 = 0
    var v17 float64 = v4
    switch 32768 {
    case 32768:
        v18 := []int{(v2 ^ v2), v2, v2}
        _ = v18
    case 9223372036854775807:
        var v19 int = v2
        _ = v19
    default:
        for fi20 := 0; fi20 < 19; fi20++ {
            var v21 int16 = 1000
            switch v10 {
            case "":
                var v22 int16 = -307
                _ = v22
            case "\n\t":
                if false {
                    v17 = v17
                    var v23 string = (v10 + (v10 + (v10 + v10)))
                    fmt.Println(v4)
                    _ = v23
                } else {
                    switch "" {
                    case "":
                        if v3 {
                            v3 = (!true)
                        } else {
                            v24 := GenS1{F0: false, F1: (v16 - (v16 << ((59) & 63)))}
                            v25 := []int{v2, (v2 / ((v2) | 1)), 117, v2}
                            if (v16 >= int32(v2)) {
                                v10 = v10
                                v26 := GenS2{F0: 1, F1: 1}
                                var v27 float64 = ((v17 - v4) * (3.14159 / v17))
                                _ = v26
                                _ = v27
                            }
                            _ = v24
                            _ = v25
                        }
                    case "foo bar":
                        v4 += v4
                    default:
                        var v28 uint = 1
                        _ = v28
                    }
                }
            default:
                v29 := GenS3{F0: (-(v1)), F1: 32767, F2: 2}
                _ = v29
            }
            v30 := GenS4{F0: (-(v1)), F1: (v2 + fi20)}
            _ = fi20
            _ = v21
            _ = v30
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
