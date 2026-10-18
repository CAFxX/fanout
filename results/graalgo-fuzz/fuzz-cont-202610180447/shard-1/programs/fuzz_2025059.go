package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 int16; F1 bool }
type GenS5 struct { F0 int32; F1 uint64 }
func genFn0(p0 string) uint64 {
    var v6 int16 = -104
    switch 127 {
    case 32768:
        v6 = (v6 << ((1) & 63))
    default:
        fmt.Println(v6)
    }
    v7 := GenS1{F0: 673}
    v8 := []int{}
    var v9 uint = 4294967295
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return uint64(v9)
}
func genFn1(p0 int32, p1 uint64, p2 int) uint {
    if true {
        var v10 uint = 384
        _ = v10
    }
    for fi11 := 0; fi11 < 16; fi11++ {
        for fi12 := 0; fi12 < 7; fi12++ {
            fmt.Println(p2, p2)
            if (!true) {
                if (2 != 34) {
                    if (1 != 557) {
                        switch (p2 << ((0) & 63)) {
                        case 150:
                            var v13 uint8 = 255
                            _ = v13
                        case 1:
                            fmt.Println(p2, fi11)
                        case 127:
                            p2 -= (fi11 << ((42) & 63))
                        }
                        var v14 uint32 = 270
                        v14 = 335
                        _ = v14
                    } else {
                        p2 -= (p2 - (fi12 >> ((255) & 63)))
                    }
                    p2 = fi11
                }
            }
            var v15 int64 = 42
            _ = fi12
            _ = v15
        }
        _ = fi11
    }
    p1 += genFn0((("a" + "hello") + ""))
    p1 = (^(p1))
    _ = p0
    _ = p1
    _ = p2
    return 65535
}

func main() {
    var v1 uint16 = 65535
    v2 := make(map[uint32]string)
    var v3 uint = 2
    v4 := make(map[uint32]uint32)
    v5 := make(map[int]bool)
    v16 := GenS2{F0: false}
    switch "\n\t" {
    case "a":
        if true {
            v17 := GenS3{F0: -559, F1: (v1 | v1)}
            _ = v17
        }
    default:
        v3 = (v3 << ((v3) & 63))
    }
    var v18 uint32 = uint32(v3)
    v19 := []int{-825, 128, 32767}
    if true {
        for fi20 := 0; fi20 < 1; fi20++ {
            var v21 uint8 = 37
            _ = fi20
            _ = v21
        }
        v22 := GenS4{F0: 0, F1: (float64((0.0 + 0.0)) > 0.0)}
        v1 -= (v1 / ((1) | 1))
        _ = v22
    }
    fmt.Println(v3, v3, v18)
    if ((float64(float32((1e10 - (0.1 * 0.1)))) * float64(0.0)) != 1e10) {
        if true {
            if ((!(!true)) && (false || (!(-128 >= -128)))) {
                v1 = (v1 >> (((v3 - v3)) & 63))
                v1 = uint16(v3)
            } else {
                v1 = (v1 >> ((v3) & 63))
                if false {
                    fmt.Println(v3, v1, v18)
                    v1 = (v1 - (v1 << ((genFn1(-1, 4294967295, 128)) & 63)))
                }
                for fi23 := 0; fi23 < 20; fi23++ {
                    fmt.Println(v3, v1, v18)
                    if (9223372036854775807 > -844) {
                        v18 = v18
                        var v24 uint16 = uint16(v3)
                        _ = v24
                    }
                    _ = fi23
                }
            }
            v18 = v18
            v25 := GenS3{F0: 0, F1: 2}
            _ = v25
        } else {
            var v26 int64 = 904
            if (true && false) {
                if false {
                    _ = len(v19)
                    for fi27 := 0; fi27 < 10; fi27++ {
                        if true {
                            v18 = (^(v18))
                        } else {
                            v3 += v3
                        }
                        fmt.Println(v1)
                        _ = fi27
                    }
                } else {
                    var v28 uint16 = (v1 ^ (^(v1)))
                    v26 += -1
                    _ = v28
                }
            } else {
                var v29 int16 = 1000
                var v30 uint = v3
                _ = v29
                _ = v30
            }
            _ = v26
        }
        v18 += (^(v18))
        if true {
            if (float32((0.1 - (1.5 * (-2.25 - 3.14159)))) != 1.5) {
                v31 := GenS5{F0: 1, F1: 42}
                var v32 float64 = float64((0.1 + 0.0))
                _ = v31
                _ = v32
            } else {
                fmt.Println(v3, v1, v18)
                switch -9223372036854775808 {
                case -448:
                    var v33 float64 = -2.25
                    _ = v33
                default:
                    var v34 int8 = -44
                    _ = v34
                }
                delete(v2, uint32(v1))
            }
            _ = len(v5)
        } else {
            v2[(v18 | 951)] = "a"
            var v35 string = "\n\t"
            _ = v35
        }
    }
    var v36 float32 = float32(((((1.5 + (-2.25 * (1.5 * -2.25))) * 1.5) / 1e10) + 3.14159))
    for fi37 := 0; fi37 < 15; fi37++ {
        var v38 uint16 = (^(v1))
        _ = fi37
        _ = v38
    }
    var v39 int8 = int8(v1)
    v19 = append(v19, -2147483648)
    var v40 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v36
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
