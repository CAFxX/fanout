package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint }
type GenS2 struct { F0 uint64; F1 uint32; F2 int16 }
type GenS3 struct { F0 uint64; F1 uint16; F2 int32 }
func genFn0(p0 uint, p1 int, p2 uint16) uint32 {
    if (false && true) {
        var v6 float64 = 1.5
        var v7 int8 = -1
        p2 += p2
        _ = v6
        _ = v7
    }
    if (false && true) {
        v8 := GenS1{F0: 1, F1: 1, F2: (p0 | p0)}
        var v9 uint32 = 2147483647
        _ = v8
        _ = v9
    }
    var v10 uint16 = 255
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 2
}

func main() {
    var v1 uint32 = 364
    var v2 int8 = -128
    var v3 int = -2147483648
    var v4 uint8 = 42
    var v5 int16 = -32768
    var v11 uint = 2
    if true {
        for fi12 := 0; fi12 < 3; fi12++ {
            if false {
                v13 := []int{fi12, 127, v3, v3}
                fmt.Println(v4, fi12)
                v11 += v11
                _ = v13
            } else {
                if true {
                    for fi14 := 0; fi14 < 8; fi14++ {
                        if true {
                            switch (fi12 ^ (-(fi12))) {
                            case 138:
                                var v15 uint16 = 22
                                _ = v15
                            }
                            switch "foo bar" {
                            case "a":
                                var v16 uint8 = 242
                                _ = v16
                            case "hello":
                                fmt.Println(v1)
                            default:
                                v17 := []int{int(v11), (^(fi14))}
                                _ = v17
                            }
                        } else {
                            var v18 string = "\n\t"
                            if (3.14159 <= 3.14159) {
                                v19 := []int{127, fi12, 32767}
                                var v20 int = fi14
                                _ = v19
                                _ = v20
                            } else {
                                if (false && (v4 < v4)) {
                                    fmt.Println(v2)
                                    v21 := GenS2{F0: 918, F1: 1, F2: v5}
                                    _ = v21
                                } else {
                                    v18 = "hello"
                                }
                            }
                            v22 := []int{}
                            _ = v18
                            _ = v22
                        }
                        v4 = 69
                        _ = fi14
                    }
                    if true {
                        var v23 int = (v3 % ((fi12) | 1))
                        var v24 int8 = (v2 ^ v2)
                        v23 -= v23
                        _ = v23
                        _ = v24
                    } else {
                        v25 := GenS3{F0: 42, F1: 65535, F2: 988}
                        var v26 uint32 = v1
                        _ = v25
                        _ = v26
                    }
                }
            }
            _ = fi12
        }
    }
    if (-408 >= 1) {
        var v27 float32 = 0.0
        v5 -= 1000
        for fi28 := 0; fi28 < 16; fi28++ {
            v29 := []int{29, v3}
            var v30 uint8 = (v4 << ((v11) & 63))
            _ = fi28
            _ = v29
            _ = v30
        }
        _ = v27
    } else {
        v31 := GenS3{F0: 4294967295, F1: 32767, F2: int32(v4)}
        if (!(-1 > -205)) {
            v32 := GenS1{F0: v4, F1: 609, F2: 255}
            v2 -= (v2 >> (((v11 << ((4294967295) & 63))) & 63))
            _ = v32
        }
        _ = v31
    }
    var v33 int = 112
    var v34 float64 = 0.0
    var v35 float32 = (0.0 / 1.0)
    switch ("hello" + "hello") {
    case "":
        fmt.Println(v3, v3)
    default:
        var v36 uint8 = v4
        _ = v36
    }
    v37 := GenS2{F0: 4294967295, F1: (v1 | v1), F2: v5}
    var v38 string = "foo bar"
    var v39 float32 = v35
    fmt.Println(v38)
    v39 += (v35 * (v35 + ((float32(((0.0 - (float64(v35) * (v34 / v34))) + v34)) / v35) * v35)))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
