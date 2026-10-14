package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 int8; F2 uint16 }
func genFn0(p0 uint8) uint {
    var v7 uint64 = 1
    var v8 string = "foo bar"
    v8 = "\n\t"
    v9 := []int{32767, 186, 676, 0}
    if false {
        fmt.Println(v8, v7, v8)
        p0 += 1
        var v10 string = "\n\t"
        _ = v10
    }
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 uint16 = 65535
    v2 := []uint32{1}
    var v3 uint = 226
    var v4 float64 = 3.14159
    v5 := []int{9223372036854775807, 42, 32767}
    var v6 uint8 = 127
    v11 := GenS1{F0: (v6 ^ (v6 * v6))}
    for fi12 := 0; fi12 < 1; fi12++ {
        var v13 float32 = 3.14159
        for fi14 := 0; fi14 < 14; fi14++ {
            if false {
                fmt.Println(v4)
            } else {
                var v15 int64 = 9223372036854775807
                if false {
                    fmt.Println(v13, v13, v6)
                    if (42 >= 895) {
                        if (((((fi12 / ((fi14) | 1)) != (^(fi14))) && (v6 != v6)) || (65 < 584)) || false) {
                            if ((3.14159 - 1.5) != v13) {
                                v1 = (^(v1))
                            }
                        }
                        v16 := GenS1{F0: v6}
                        var v17 float64 = ((1.5 + float64(v13)) * -2.25)
                        v17 += float64(v13)
                        if false {
                            v18 := GenS2{F0: (v15 / (((v15 % (((v15 % (((v15 | v15)) | 1))) | 1))) | 1))}
                            _ = v18
                        }
                        _ = v16
                        _ = v17
                    }
                } else {
                    if len(v5) > 0 {
                        v5[0] = (fi14 / (((^(fi12))) | 1))
                    }
                }
                if false {
                    var v19 int32 = -2147483648
                    if ((v1 | (v1 / (((v1 ^ (^(v1)))) | 1))) >= (v1 % ((430) | 1))) {
                        v20 := GenS1{F0: v6}
                        fmt.Println(fi14, v6, v3)
                        var v21 int = (-(fi12))
                        _ = v20
                        _ = v21
                    }
                    v22 := GenS2{F0: (v15 >> ((v3) & 63))}
                    _ = v19
                    _ = v22
                } else {
                    var v23 int32 = 0
                    _ = v23
                }
                var v24 string = "hello"
                _ = v15
                _ = v24
            }
            var v25 float64 = (-2.25 - 1.5)
            v25 -= v25
            _ = fi14
            _ = v25
        }
        for fi26 := 0; fi26 < 12; fi26++ {
            v4 -= ((v4 - (v4 - 0.0)) - (v4 - ((v4 - float64(v13)) * v4)))
            fmt.Println(v13, v6, v4)
            _ = len(v5)
            _ = fi26
        }
        _ = fi12
        _ = v13
    }
    var v27 float64 = (-2.25 / float64((1e10 / 0.1)))
    var v28 bool = (-843 != 9223372036854775807)
    var v29 int8 = 0
    if (v28 && (42 < 703)) {
        v4 = v27
        v5 = append(v5, -144)
    } else {
        v30 := GenS2{F0: 756}
        fmt.Println(v28, v3)
        var v31 bool = true
        _ = v30
        _ = v31
    }
    if len(v5) > 0 {
        v5[0] = -2147483648
    }
    var v32 int32 = 2147483647
    v28 = (!v28)
    if (274 > 0) {
        fmt.Println(v6, v27, v29)
        v3 = (v3 + v3)
        v1 += (v1 / ((uint16(v3)) | 1))
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v33 := GenS3{F0: "", F1: (v29 - (v29 | v29)), F2: (^(v1))}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
