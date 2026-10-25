package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 uint32 }
func genFn0(p0 string, p1 int8) uint32 {
    fmt.Println(p0, p0)
    p0 = p0
    var v5 int64 = 726
    var v6 bool = false
    v7 := GenS1{F0: -1}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 2147483647
}

func main() {
    var v1 int = -923
    v2 := make(map[uint32]uint32)
    var v3 float32 = 1.5
    v4 := make(map[uint32]bool)
    var v8 int32 = -2147483648
    v9 := GenS1{F0: 569}
    var v10 uint16 = 2
    var v11 uint32 = 42
    v10 -= (v10 | (^(v10)))
    if (false || true) {
        switch ("" + "") {
        case "\n\t":
            for fi12 := 0; fi12 < 20; fi12++ {
                v4[genFn0("", 100)] = ((1 != int16(v11)) || false)
                var v13 string = "hello"
                v14 := GenS1{F0: (fi12 * (fi12 << ((4294967295) & 63)))}
                _ = fi12
                _ = v13
                _ = v14
            }
        default:
            _ = len(v4)
        }
        v15 := GenS2{F0: 2, F1: genFn0(("a" + ("a" + "hello")), 127)}
        _ = v15
    }
    fmt.Println(v8, v11, v11)
    v1 = (-(v1))
    if false {
        var v16 int32 = v8
        _ = v16
    } else {
        if ((v1 - 4294967295) < 1) {
            v1 -= (^(v1))
            if ((((1e10 / -2.25) / (0.1 * 0.1)) + 0.1) > 1e10) {
                if (v1 == v1) {
                    var v17 int64 = 797
                    _ = v17
                } else {
                    var v18 int32 = v8
                    switch (v1 + v1) {
                    case 0:
                        fmt.Println(v1)
                    case 32767:
                        if true {
                            var v19 float32 = (((v3 * (v3 * v3)) * ((((v3 * ((((((v3 - v3) - -2.25) - (-2.25 + v3)) - (float32(1.5) * v3)) / v3) / (0.0 + 1e10))) - (v3 / -2.25)) + v3) / (float32(1.5) - v3))) * v3)
                            fmt.Println(v19, v11)
                            var v20 float64 = ((3.14159 - 0.0) * (0.0 + 0.1))
                            _ = v19
                            _ = v20
                        } else {
                            var v21 float64 = 0.1
                            _ = v21
                        }
                    case -1:
                        v1 = (v1 / ((int(v8)) | 1))
                    default:
                        delete(v4, v11)
                    }
                    _ = v18
                }
            } else {
                var v22 string = ("a" + "hello")
                _ = v22
            }
            var v23 float32 = (v3 - float32(0.1))
            _ = v23
        } else {
            if true {
                for fi24 := 0; fi24 < 3; fi24++ {
                    v3 += (-2.25 / (v3 / ((v3 * (0.1 + (1.5 / -2.25))) / (v3 / 1.0))))
                    _ = fi24
                }
            }
            switch "a" {
            case "\n\t":
                fmt.Println(v8, v10, v8)
            case "foo bar":
                v25 := GenS1{F0: (v1 & v1)}
                _ = v25
            case "":
                fmt.Println(v3, v8, v1)
            default:
                v26 := []int{(v1 % (((v1 / (((v1 << ((65535) & 63))) | 1))) | 1))}
                _ = v26
            }
        }
        if true {
            for fi27 := 0; fi27 < 3; fi27++ {
                var v28 string = ""
                var v29 uint64 = 4294967295
                if false {
                    v11 = (v11 / ((835) | 1))
                    var v30 uint32 = uint32(v10)
                    _ = v30
                } else {
                    fmt.Println(v1, v3, v8)
                }
                _ = fi27
                _ = v28
                _ = v29
            }
        } else {
            fmt.Println(v3)
            v11 = (^(v11))
        }
    }
    switch v1 {
    case 4294967295:
        v8 += -1
    case -1:
        if false {
            v31 := GenS2{F0: 42, F1: 0}
            _ = v31
        }
    default:
        if true {
            if (128 < 1) {
                v3 += v3
                v32 := []int{v1, v1, v1, (v1 & (v1 << ((42) & 63)))}
                _ = v32
            }
            v2[v11] = (v11 & v11)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
