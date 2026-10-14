package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 uint32 }
type GenS2 struct { F0 uint32; F1 bool }
func genFn0() uint {
    var v6 float32 = float32(float64((-2.25 / 1e10)))
    v7 := []int{113}
    v6 = (v6 + (-2.25 / float32(0.1)))
    var v8 int = 128
    _ = len(v7)
    _ = v6
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 uint64 = 226
    var v2 bool = ((116 >= 0) && (true && true))
    v3 := make(map[string]bool)
    var v4 int8 = -128
    v5 := make(map[uint32]uint32)
    delete(v5, 2147483648)
    var v9 int = -472
    v10 := []int{}
    if ((((float32(((1.5 * float64(((1e10 / -2.25) + 0.1))) + (1e10 / 1.0))) + (((3.14159 - float32(1e10)) + 0.1) + -2.25)) + (1e10 - ((1e10 - float32(1.5)) + (-2.25 * (-2.25 + (0.0 * 1.5)))))) > 1.5) && true) {
        if (-985 == 1) {
            var v11 int32 = -1
            for fi12 := 0; fi12 < 18; fi12++ {
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                var v13 float64 = ((float64(((0.0 + (((float32(float64(float32(1.5))) + 1e10) - 1e10) + ((1.5 + 3.14159) - float32(1e10)))) * (1e10 - ((3.14159 * 0.0) + (3.14159 / 1e10))))) + 1.5) + float64(((0.0 * (((3.14159 * float32(((3.14159 - 3.14159) - 1.5))) * 0.1) / 1.0)) * float32((float64(3.14159) + 1.5)))))
                _ = fi12
                _ = v13
            }
            var v14 int = int(v4)
            _ = v11
            _ = v14
        } else {
            v9 = (v9 % (((v9 - (v9 + (v9 | (-(v9)))))) | 1))
            fmt.Println(v2, v4)
            v15 := v3["hello"]
            _ = v15
        }
        v16 := GenS1{F0: 2147483647, F1: 120, F2: 0}
        _ = v16
    } else {
        if v2 {
            v3["\n\t"] = ((!(42 == 0)) || (false || ((!(0 <= 0)) || (!((168 > v1) && v2)))))
        } else {
            switch "" {
            case "foo bar":
                v17 := GenS1{F0: 65535, F1: -2147483648, F2: 980}
                _ = v17
            case "hello":
                for fi18 := 0; fi18 < 15; fi18++ {
                    for fi19 := 0; fi19 < 14; fi19++ {
                        v4 -= 42
                        _ = fi19
                    }
                    var v20 bool = v2
                    v1 -= v1
                    _ = fi18
                    _ = v20
                }
            default:
                switch 32767 {
                case 4294967295:
                    v21 := GenS1{F0: 2, F1: -2147483648, F2: 2147483648}
                    _ = v21
                case -198:
                    v9 += v9
                }
            }
        }
    }
    fmt.Println(v9)
    delete(v5, 42)
    var v22 uint = 18446744073709551615
    if (v2 || (false || v2)) {
        if (1 > 9223372036854775807) {
            fmt.Println(v4, v2)
            var v23 bool = ((0.0 / -2.25) == 1.5)
            delete(v5, 65535)
            _ = v23
        } else {
            v24 := GenS2{F0: 65535, F1: v2}
            _ = v24
        }
    } else {
        var v25 string = "a"
        _ = v25
    }
    if ("" != "\n\t") {
        var v26 float64 = (((-2.25 * 3.14159) * 0.1) - (float64(0.1) / 1e10))
        var v27 int8 = v4
        _ = v26
        _ = v27
    }
    for fi28 := 0; fi28 < 18; fi28++ {
        v2 = (v1 <= 4294967295)
        _ = len(v5)
        _ = fi28
    }
    delete(v3, "a")
    var v29 uint64 = (v1 | (v1 ^ v1))
    var v30 uint8 = 2
    var v31 int16 = -132
    v2 = (!((int64(v30) != 0) && v2))
    for fi32 := 0; fi32 < 2; fi32++ {
        for fi33 := 0; fi33 < 3; fi33++ {
            if len(v10) > 0 {
                v10[0] = 9223372036854775807
            }
            v22 += (v22 + 2)
            var v34 int = v9
            v9 = -9223372036854775808
            v35 := GenS2{F0: 2, F1: (1 > 9223372036854775807)}
            _ = fi33
            _ = v34
            _ = v35
        }
        var v36 string = (("foo bar" + "") + "a")
        var v37 int = 2147483647
        _ = fi32
        _ = v36
        _ = v37
    }
    delete(v5, 2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v9
}
