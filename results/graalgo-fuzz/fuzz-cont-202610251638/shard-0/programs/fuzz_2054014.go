package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 int32; F1 uint16; F2 uint32 }

func main() {
    v1 := make(map[string]string)
    var v2 string = ("\n\t" + (("\n\t" + (("foo bar" + "\n\t") + "")) + "foo bar"))
    var v3 string = v2
    v4 := make(map[string]int)
    v5 := make(map[int]bool)
    if (929 == 714) {
        var v6 float32 = 0.1
        fmt.Println(v6, v2)
        _ = v6
    }
    v7 := GenS1{F0: 1, F1: 1}
    v3 = (v2 + ("a" + v3))
    var v8 uint8 = 145
    v9 := []int{32768, 2147483647, -9223372036854775808, 127}
    var v10 bool = false
    fmt.Println(v3)
    if true {
        if (42 != 42) {
            fmt.Println(v10, v8)
        }
    }
    var v11 uint16 = 255
    v10 = ((32768 <= 128) && (float32(3.14159) < 1.5))
    for fi12 := 0; fi12 < 14; fi12++ {
        v13 := GenS1{F0: -1, F1: (^(v11))}
        for fi14 := 0; fi14 < 10; fi14++ {
            v15 := v5[9223372036854775807]
            if (202 <= -588) {
                v16 := GenS1{F0: -1, F1: v11}
                v17 := GenS1{F0: 42, F1: (v11 % ((uint16(v8)) | 1))}
                _ = v16
                _ = v17
            } else {
                var v18 bool = v10
                var v19 bool = v15
                _ = v18
                _ = v19
            }
            _ = fi14
            _ = v15
        }
        var v20 int16 = -699
        _ = fi12
        _ = v13
        _ = v20
    }
    v21 := GenS2{F0: -969, F1: 863, F2: 611}
    v22 := GenS2{F0: 1, F1: (v11 & (v11 % ((uint16(v8)) | 1))), F2: 509}
    var v23 uint64 = 2
    for fi24 := 0; fi24 < 8; fi24++ {
        v25 := GenS2{F0: 2147483647, F1: (^(v11)), F2: 786}
        v8 += v8
        v4[v3] = fi24
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 15; fi26++ {
        v27 := v1[("foo bar" + "a")]
        for fi28 := 0; fi28 < 7; fi28++ {
            v10 = (int64(fi28) > -9223372036854775808)
            v1[(((v27 + v2) + ("\n\t" + v3)) + ((((v3 + ("a" + "a")) + v3) + ("a" + ((v27 + v27) + ""))) + v27))] = (v27 + (v27 + v27))
            var v29 int8 = 0
            if v10 {
                v2 = "\n\t"
                v8 += v8
            } else {
                if v10 {
                    _ = len(v9)
                } else {
                    if v10 {
                        var v30 int64 = 0
                        fmt.Println(fi26, v23)
                        _ = v30
                    } else {
                        if false {
                            fmt.Println(v27)
                        } else {
                            var v31 bool = (v8 <= v8)
                            _ = v31
                        }
                    }
                    v32 := GenS2{F0: 236, F1: v11, F2: 326}
                    _ = v32
                }
            }
            _ = fi28
            _ = v29
        }
        _ = fi26
        _ = v27
    }
    for fi33 := 0; fi33 < 4; fi33++ {
        var v34 uint8 = (v8 * v8)
        _ = fi33
        _ = v34
    }
    if true {
        v35 := GenS2{F0: 42, F1: (v11 ^ v11), F2: 1}
        delete(v4, v3)
        _ = v35
    } else {
        var v36 int16 = -1
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
