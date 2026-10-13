package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 int = 127
    var v2 int8 = 56
    v3 := make(map[uint32]int)
    var v4 float64 = -2.25
    var v5 int32 = -2147483648
    var v6 float32 = ((((0.1 + -2.25) * ((((-2.25 + float32(v4)) / float32(v4)) / 1.0) * 3.14159)) + (float32(((-2.25 / 0.1) + float64((0.0 - 1e10)))) / 1.5)) - 1e10)
    var v7 int8 = (v2 + v2)
    var v8 int16 = 32767
    v9 := GenS1{F0: (v5 * (v5 & 5))}
    fmt.Println(v5)
    v1 = -2147483648
    v10 := []int{v1, (v1 * 2147483647), v1, v1}
    if true {
        if (true && true) {
            if true {
                fmt.Println(v2, v5)
                v4 += v4
                fmt.Println(v6)
            } else {
                var v11 uint8 = 127
                if (((v6 * -2.25) * v6) != (v6 + v6)) {
                    v3[42] = v1
                    var v12 int64 = -24
                    _ = v12
                }
                v1 = v1
                _ = v11
            }
            var v13 string = "a"
            _ = v13
        } else {
            fmt.Println(v4)
            v14 := GenS1{F0: (v5 * (-(v5)))}
            switch ("\n\t" + "a") {
            case "a":
                _ = len(v10)
            default:
                for fi15 := 0; fi15 < 17; fi15++ {
                    for fi16 := 0; fi16 < 6; fi16++ {
                        v10 = append(v10, (fi16 + 0))
                        if ((v2 % (((v2 & (v7 % ((v7) | 1)))) | 1)) == -1) {
                            if true {
                                fmt.Println(fi16)
                                switch (v1 & v1) {
                                case 2147483647:
                                    fmt.Println(fi16, fi16, fi16)
                                case 9223372036854775807:
                                    fmt.Println(v7, v5)
                                }
                                var v17 int64 = 272
                                _ = v17
                            }
                            v3[65535] = (fi16 << ((65535) & 63))
                            fmt.Println(v2)
                        }
                        v4 = (((((float64(0.0) * v4) / (0.1 - (v4 - (1e10 + 3.14159)))) + (-2.25 - (0.1 + 1e10))) - (v4 - 1.5)) + v4)
                        _ = fi16
                    }
                    fmt.Println(v8, v2)
                    fmt.Println(v1)
                    _ = fi15
                }
            }
            _ = v14
        }
        if len(v10) > 0 {
            v10[0] = (v1 / (((v1 * (v1 / (((v1 >> ((309) & 63))) | 1)))) | 1))
        }
    } else {
        for fi18 := 0; fi18 < 15; fi18++ {
            if (0.1 > 0.0) {
                var v19 int = (^(fi18))
                _ = v19
            }
            _ = fi18
        }
    }
    v5 -= 2147483647
    v8 += v8
    var v20 uint8 = 42
    var v21 int = 4294967295
    var v22 string = "\n\t"
    v4 = (3.14159 * (v4 + v4))
    var v23 float64 = v4
    v2 += (v2 ^ v7)
    _ = len(v10)
    var v24 bool = (-1 > -379)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
