package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64 }
func genFn0(p0 uint32, p1 bool, p2 int32) uint16 {
    var v5 bool = p1
    for fi6 := 0; fi6 < 4; fi6++ {
        if false {
            for fi7 := 0; fi7 < 17; fi7++ {
                v8 := []int{(fi6 / (((fi7 ^ fi7)) | 1)), fi6, (fi7 * (fi7 / (((fi6 * 128)) | 1)))}
                _ = fi7
                _ = v8
            }
            if (42 < 42) {
                var v9 bool = v5
                v5 = (-2147483648 >= (p2 % ((p2) | 1)))
                if v5 {
                    v10 := []int{32768}
                    for fi11 := 0; fi11 < 2; fi11++ {
                        v10 = append(v10, fi6)
                        _ = fi11
                    }
                    p2 = p2
                    _ = v10
                } else {
                    v12 := GenS1{F0: -69}
                    var v13 int8 = -107
                    v14 := GenS2{F0: 399}
                    _ = v12
                    _ = v13
                    _ = v14
                }
                _ = v9
            } else {
                var v15 int16 = 1000
                _ = v15
            }
        } else {
            fmt.Println(p2, p0)
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 255
}

func main() {
    var v1 uint16 = 255
    var v2 int32 = 139
    var v3 int16 = 1000
    var v4 int32 = (v2 & v2)
    if (87 > 0) {
        v16 := GenS1{F0: -1}
        for fi17 := 0; fi17 < 8; fi17++ {
            var v18 int = 451
            v19 := GenS1{F0: 49}
            _ = fi17
            _ = v18
            _ = v19
        }
        for fi20 := 0; fi20 < 2; fi20++ {
            v4 = (v4 - v4)
            _ = fi20
        }
        _ = v16
    }
    v21 := GenS2{F0: -9223372036854775808}
    fmt.Println(v4, v4)
    v22 := GenS2{F0: -1}
    v2 -= v2
    v1 += genFn0(705, (false && ((v3 >> ((360) & 63)) < -1)), (v2 + v4))
    v23 := []int{32767, 127, 448}
    fmt.Println(v2, v4, v2)
    var v24 int32 = v4
    switch "hello" {
    case "a":
        if false {
            v1 = (v1 & (^(v1)))
            var v25 uint = 18446744073709551615
            _ = v25
        } else {
            _ = len(v23)
            v1 -= 0
        }
    case "":
        fmt.Println(v3, v24)
    }
    for fi26 := 0; fi26 < 20; fi26++ {
        v27 := GenS2{F0: -1}
        fmt.Println(v1, fi26, fi26)
        _ = fi26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
