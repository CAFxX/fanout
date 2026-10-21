package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint32, p1 uint16, p2 uint32) int {
    v4 := GenS1{F0: 2, F1: 42}
    p1 += (p1 >> ((255) & 63))
    fmt.Println(p1, p0, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 32768
}
func genFn1(p0 int32, p1 bool, p2 uint8) int {
    var v5 float64 = 3.14159
    var v6 float64 = v5
    if false {
        p0 = (p0 % ((p0) | 1))
        v7 := []int{genFn0(4294967295, 292, 0), int(p2), -363, genFn0(4294967295, 1, 156)}
        v8 := GenS1{F0: 18446744073709551615, F1: (p0 * (p0 >> ((222) & 63)))}
        _ = v7
        _ = v8
    }
    v5 = v6
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return -35
}

func main() {
    var v1 uint32 = 4294967295
    v2 := []uint32{v1, (v1 ^ (v1 | v1))}
    var v3 uint32 = 2
    fmt.Println(v3, v1)
    v9 := GenS2{F0: -1}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi10 := 0; fi10 < 13; fi10++ {
        v11 := GenS2{F0: -1}
        v3 -= v1
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        v2 = append(v2, (v3 - (v3 ^ (^(v1)))))
        var v13 bool = false
        _ = fi12
        _ = v13
    }
    switch 4294967295 {
    case -1:
        var v14 bool = (3.14159 == float64(0.0))
        _ = v14
    default:
        _ = len(v2)
    }
    fmt.Println(v1, v3)
    var v15 bool = false
    if (42 <= 1) {
        var v16 bool = v15
        for fi17 := 0; fi17 < 15; fi17++ {
            v1 -= v3
            v18 := GenS1{F0: 42, F1: 42}
            _ = len(v2)
            _ = fi17
            _ = v18
        }
        _ = v16
    }
    var v19 uint16 = 65535
    switch ("hello" + ("hello" + "")) {
    case "":
        fmt.Println(v3, v15)
    case "\n\t":
        if (2 > 42) {
            var v20 int = 4294967295
            for fi21 := 0; fi21 < 3; fi21++ {
                v15 = (((^(v19)) != 32767) || ((v1 != (v1 * v3)) || (!v15)))
                _ = fi21
            }
            _ = v20
        } else {
            var v22 uint16 = (v19 + 476)
            for fi23 := 0; fi23 < 20; fi23++ {
                var v24 bool = (1.5 >= (0.0 / 0.1))
                switch (^(fi23)) {
                case 899:
                    if v15 {
                        if (v15 && ((!(v24 && v24)) || v24)) {
                            v15 = false
                        } else {
                            v15 = v15
                            fmt.Println(v19, v1)
                        }
                        v22 = v22
                        if (1e10 == 3.14159) {
                            var v25 string = "foo bar"
                            _ = v25
                        }
                    }
                default:
                    v19 -= 65535
                }
                _ = fi23
                _ = v24
            }
            _ = v22
        }
    default:
        var v26 uint32 = (v3 << ((0) & 63))
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v9
}
