package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 uint }
func genFn0(p0 int8) uint8 {
    if (1 == 855) {
        p0 -= p0
        v7 := GenS1{F0: 1}
        _ = v7
    } else {
        p0 = (p0 * (p0 % ((p0) | 1)))
        var v8 uint8 = 117
        _ = v8
    }
    fmt.Println(p0)
    var v9 float32 = ((1e10 + -2.25) + ((0.1 + ((float32(float64(0.0)) - (1e10 * ((0.1 / -2.25) * (0.1 * float32(-2.25))))) / (1.5 - (0.0 / 1.5)))) + 1e10))
    var v10 uint8 = 135
    _ = p0
    _ = v10
    _ = v9
    return (v10 % ((v10) | 1))
}

func main() {
    var v1 int64 = -977
    var v2 float64 = (1.5 - float64(1e10))
    v3 := []int{32767, 1}
    var v4 float64 = v2
    v5 := make(map[int]string)
    v6 := make(map[int]int)
    for fi11 := 0; fi11 < 15; fi11++ {
        if false {
            if (0.0 == 1.5) {
                v12 := GenS2{F0: (v1 / ((-9223372036854775808) | 1)), F1: (fi11 >> ((42) & 63))}
                _ = len(v6)
                _ = v12
            } else {
                var v13 string = "a"
                _ = v13
            }
            var v14 uint16 = 2
            if (false && (false && false)) {
                for fi15 := 0; fi15 < 6; fi15++ {
                    var v16 float32 = 3.14159
                    _ = len(v6)
                    if ("\n\t" < "") {
                        fmt.Println(v14, v1, v4)
                    }
                    _ = fi15
                    _ = v16
                }
                fmt.Println(v1)
            }
            _ = v14
        } else {
            fmt.Println(fi11, v4, v4)
        }
        _ = len(v6)
        _ = fi11
    }
    v17 := GenS3{F0: 255}
    if false {
        var v18 int8 = 120
        for fi19 := 0; fi19 < 4; fi19++ {
            var v20 uint = 18446744073709551615
            switch "" {
            case "\n\t":
                v2 = (((((v2 / v2) / v4) - 0.0) - 0.0) / (float64(((((1e10 / 0.1) / 0.1) * 1.5) * ((0.1 - ((((((0.1 + (float32((v2 - v4)) * 1e10)) * -2.25) - 1.5) * (1e10 - 3.14159)) - float32((((v2 * ((float64((0.0 * 1e10)) * (v2 + v2)) * 3.14159)) / v2) * v4))) * 0.1)) / 1e10))) - (1e10 + v4)))
            case "hello":
                v20 = v20
            }
            _ = fi19
            _ = v20
        }
        for fi21 := 0; fi21 < 10; fi21++ {
            if false {
                v22 := GenS3{F0: 65535}
                _ = v22
            } else {
                v23 := GenS1{F0: 1}
                var v24 int16 = 42
                _ = v23
                _ = v24
            }
            fmt.Println(v1, v4)
            for fi25 := 0; fi25 < 3; fi25++ {
                var v26 float64 = v4
                _ = fi25
                _ = v26
            }
            _ = fi21
        }
        _ = v18
    } else {
        v4 = v4
        if (42 > 1) {
            fmt.Println(v2)
            var v27 uint8 = 42
            v1 += (v1 * -1)
            _ = v27
        }
        for fi28 := 0; fi28 < 7; fi28++ {
            if false {
                var v29 int16 = 1
                v1 = (v1 << ((641) & 63))
                _ = v29
            }
            var v30 int = 4294967295
            if false {
                v2 -= float64(((((3.14159 * 1e10) * (((0.1 - 1.5) * 1e10) / 1.0)) / 1.0) + (1e10 * (-2.25 * float32(float64(((1.5 + 0.1) + float32(1e10))))))))
            }
            _ = fi28
            _ = v30
        }
    }
    if ((0.0 > (0.1 - -2.25)) || true) {
        fmt.Println(v1)
        v2 += v2
    }
    if (1e10 < (3.14159 - ((float32(-2.25) + 0.1) + 0.1))) {
        switch 785 {
        case 116:
            fmt.Println(v4, v2, v1)
        default:
            var v31 int64 = 0
            _ = v31
        }
    }
    var v32 float64 = float64(1.5)
    v2 -= v32
    v1 -= 0
    v1 += (v1 << ((42) & 63))
    v5[568] = ("\n\t" + "foo bar")
    v33 := GenS1{F0: 42}
    if false {
        if false {
            var v34 uint8 = 127
            _ = v34
        }
        var v35 uint16 = 65535
        v2 = float64(0.1)
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
