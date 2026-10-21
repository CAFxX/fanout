package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 int16 }
type GenS2 struct { F0 int16; F1 uint16; F2 uint32 }
type GenS3 struct { F0 string; F1 int }
func genFn0(p0 int64) int8 {
    v4 := []int{32767, -9223372036854775808, -764}
    v5 := GenS1{F0: ("foo bar" + "foo bar"), F1: ("\n\t" + "hello"), F2: 32767}
    _ = p0
    _ = v4
    _ = v5
    return -1
}

func main() {
    var v1 bool = (2 > 42)
    var v2 int32 = -2147483648
    var v3 int8 = 100
    v6 := []int{32768, -1}
    for fi7 := 0; fi7 < 4; fi7++ {
        switch (fi7 + fi7) {
        case -2147483648:
            var v8 bool = (int8(fi7) <= (-(v3)))
            _ = v8
        case 0:
            if len(v6) > 0 {
                v6[0] = (fi7 % ((691) | 1))
            }
        case 797:
            fmt.Println(v2)
        default:
            if false {
                v3 = (v3 * v3)
            } else {
                v1 = (v1 || (v1 && (!true)))
                v1 = v1
                if (0.0 >= (-2.25 + (1.5 * -2.25))) {
                    v9 := GenS1{F0: "", F1: "", F2: -1}
                    switch (^(fi7)) {
                    case 0:
                        var v10 int32 = 2147483647
                        _ = v10
                    case -9223372036854775808:
                        for fi11 := 0; fi11 < 4; fi11++ {
                            if (((float32(-2.25) - float32((3.14159 + 1.5))) + (3.14159 / 0.1)) <= (0.1 + -2.25)) {
                                var v12 uint32 = 4294967295
                                _ = v12
                            }
                            v13 := GenS2{F0: -991, F1: 25, F2: 2147483648}
                            _ = fi11
                            _ = v13
                        }
                    case 128:
                        v6 = append(v6, fi7)
                    }
                    _ = v9
                }
            }
        }
        _ = fi7
    }
    var v14 string = ("foo bar" + ("\n\t" + "a"))
    var v15 int64 = -9223372036854775808
    v6 = append(v6, -691)
    v16 := GenS3{F0: ((v14 + v14) + v14), F1: 32767}
    var v17 int16 = -32768
    v18 := GenS1{F0: v14, F1: v14, F2: v17}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v6
}
