package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 int16 }
type GenS4 struct { F0 uint32; F1 uint32; F2 uint16 }
func genFn0(p0 uint8, p1 bool, p2 uint8) uint32 {
    if false {
        v4 := []int{-9223372036854775808, 32767, 32767}
        var v5 float32 = (((3.14159 * float32((0.1 + 0.0))) - 3.14159) * (float32(0.0) * 1e10))
        _ = v4
        _ = v5
    }
    fmt.Println(p2, p0, p2)
    var v6 int16 = 93
    v7 := GenS1{F0: 573, F1: 2147483647, F2: 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 2
}
func genFn1(p0 bool, p1 int) int {
    v8 := GenS2{F0: "a"}
    p0 = (!p0)
    var v9 bool = false
    p0 = true
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return p1
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = -2.25
    var v3 bool = false
    var v10 int = -9223372036854775808
    fmt.Println(v2, v3)
    for fi11 := 0; fi11 < 1; fi11++ {
        if (253 != 864) {
            v12 := GenS1{F0: (^(v10)), F1: 1, F2: genFn1(false, -456)}
            var v13 uint32 = 80
            var v14 int = -628
            _ = v12
            _ = v13
            _ = v14
        }
        if ((!v3) && (genFn0(127, v3, 127) == 65535)) {
            if ((-(fi11)) < genFn1((1 < -1), (v10 ^ -2147483648))) {
                v15 := []int{v10, (fi11 ^ -9223372036854775808), 828, fi11}
                if v3 {
                    if v3 {
                        var v16 uint32 = 982
                        var v17 uint16 = 255
                        _ = v16
                        _ = v17
                    } else {
                        if v3 {
                            fmt.Println(v2, v10, v3)
                            v10 = fi11
                            v3 = v3
                        } else {
                            v18 := GenS3{F0: 1, F1: -982}
                            if (false || true) {
                                switch (("\n\t" + "") + "foo bar") {
                                case "a":
                                    if v3 {
                                        if (((((("a" + "hello") > ("\n\t" + "a")) || v3) || v3) || (1 <= 42)) || false) {
                                            v19 := GenS1{F0: (-(fi11)), F1: 0, F2: 32767}
                                            switch "foo bar" {
                                            case "foo bar":
                                                v2 -= 0.0
                                            case "\n\t":
                                                v3 = false
                                            case "a":
                                                v20 := GenS2{F0: "a"}
                                                _ = v20
                                            default:
                                                for fi21 := 0; fi21 < 7; fi21++ {
                                                    var v22 bool = (381 != 2)
                                                    v1[2147483647] = v3
                                                    _ = fi21
                                                    _ = v22
                                                }
                                            }
                                            fmt.Println(v3, fi11, v10)
                                            _ = v19
                                        } else {
                                            v23 := v1[2147483648]
                                            v2 = float64(0.1)
                                            _ = v23
                                        }
                                    }
                                case "":
                                    v10 = genFn1(v3, (v10 / ((v10) | 1)))
                                default:
                                    var v24 string = "hello"
                                    _ = v24
                                }
                                v10 = 2147483647
                            } else {
                                var v25 uint32 = 1
                                for fi26 := 0; fi26 < 19; fi26++ {
                                    v10 = (v10 >> ((42) & 63))
                                    _ = fi26
                                }
                                if v3 {
                                    var v27 int32 = -1
                                    fmt.Println(fi11, v25, v3)
                                    _ = len(v15)
                                    _ = v27
                                }
                                _ = v25
                            }
                            _ = v18
                        }
                    }
                    var v28 uint32 = 0
                    _ = v28
                } else {
                    var v29 uint64 = 18446744073709551615
                    _ = v29
                }
                delete(v1, 375)
                _ = v15
            }
            if (255 >= 128) {
                v10 = v10
                var v30 uint64 = 2
                _ = v30
            } else {
                switch 2147483647 {
                case 42:
                    _ = len(v1)
                case 155:
                    v31 := GenS4{F0: 4294967295, F1: 0, F2: 42}
                    _ = v31
                }
            }
        }
        var v32 int32 = -1
        _ = fi11
        _ = v32
    }
    var v33 uint8 = 42
    switch v10 {
    case 127:
        _ = len(v1)
    default:
        v34 := []int{42, (v10 * (v10 ^ v10))}
        _ = v34
    }
    v35 := []int{genFn1((v3 && (!(290 == 1))), v10), v10}
    var v36 uint16 = 65535
    fmt.Println(v2, v33, v3)
    fmt.Println(v2, v2, v33)
    v33 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v33
    _ = v35
    _ = v36
}
