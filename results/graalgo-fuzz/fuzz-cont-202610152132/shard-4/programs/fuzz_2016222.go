package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 int64 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 uint64; F1 int32; F2 uint16 }
type GenS5 struct { F0 uint32 }
func genFn0(p0 int8, p1 int16, p2 uint) uint8 {
    var v6 int32 = 1
    v7 := GenS1{F0: 542, F1: 699, F2: 0}
    if true {
        var v8 int32 = v6
        _ = v8
    } else {
        v6 += 336
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 150
}
func genFn1() uint16 {
    v9 := []int{}
    switch 1 {
    case 9223372036854775807:
        if len(v9) > 0 {
            v9[0] = -741
        }
    case 388:
        var v10 float32 = (-2.25 / ((0.1 - -2.25) * 0.1))
        _ = v10
    default:
        var v11 uint8 = 97
        _ = v11
    }
    _ = v9
    return 317
}

func main() {
    var v1 string = ("a" + ("foo bar" + "hello"))
    v2 := []int{4294967295, 0}
    var v3 float32 = 1.5
    v4 := make(map[uint32]uint32)
    var v5 int32 = -1
    fmt.Println(v1, v1, v3)
    v4[0] = 0
    if false {
        var v12 float32 = (v3 / (v3 / (1.5 * v3)))
        v13 := GenS2{F0: false, F1: -594}
        _ = v12
        _ = v13
    }
    v14 := GenS3{F0: 4294967295, F1: 18446744073709551615}
    switch 32768 {
    case 990:
        var v15 int32 = (v5 >> ((18446744073709551615) & 63))
        _ = v15
    default:
        v16 := GenS2{F0: false, F1: 1}
        _ = v16
    }
    v3 -= v3
    v3 -= float32(float64(float32((-2.25 / (3.14159 * (1e10 - 0.1))))))
    var v17 uint16 = 0
    fmt.Println(v3, v5, v1)
    switch -9223372036854775808 {
    case -2147483648:
        v1 = ("\n\t" + "foo bar")
    case 9223372036854775807:
        var v18 string = ((v1 + v1) + "a")
        _ = v18
    }
    v5 -= (^(v5))
    v19 := GenS4{F0: 18446744073709551615, F1: v5, F2: 255}
    var v20 int = 9223372036854775807
    var v21 uint16 = (v17 >> ((7) & 63))
    if (1e10 <= 1e10) {
        fmt.Println(v3, v21, v1)
    } else {
        switch (v20 << ((395) & 63)) {
        case 229:
            v22 := GenS4{F0: 732, F1: v5, F2: 950}
            _ = v22
        default:
            v5 -= v5
        }
    }
    _ = len(v2)
    for fi23 := 0; fi23 < 14; fi23++ {
        for fi24 := 0; fi24 < 13; fi24++ {
            var v25 bool = (0 == 2)
            fmt.Println(fi23, v3, v17)
            v25 = v25
            v26 := v4[2]
            v5 = (v5 & v5)
            delete(v4, 2147483648)
            _ = fi24
            _ = v25
            _ = v26
        }
        v27 := GenS4{F0: 1, F1: v5, F2: genFn1()}
        _ = fi23
        _ = v27
    }
    switch v20 {
    case 241:
        for fi28 := 0; fi28 < 7; fi28++ {
            var v29 int64 = 42
            switch v1 {
            case "a":
                v21 = v21
            case "\n\t":
                if ((676 > 655) || (true && (-2147483648 == fi28))) {
                    if false {
                        v30 := GenS5{F0: uint32(v21)}
                        fmt.Println(v20, v21, v20)
                        for fi31 := 0; fi31 < 15; fi31++ {
                            var v32 int32 = (v5 + (-(v5)))
                            var v33 float32 = v3
                            _ = fi31
                            _ = v32
                            _ = v33
                        }
                        _ = v30
                    } else {
                        v20 -= fi28
                        for fi34 := 0; fi34 < 1; fi34++ {
                            var v35 int8 = 0
                            v36 := GenS1{F0: 162, F1: (^(v21)), F2: v29}
                            v21 += 48
                            _ = fi34
                            _ = v35
                            _ = v36
                        }
                    }
                    if (!true) {
                        var v37 bool = (false && (v5 <= v5))
                        _ = v37
                    } else {
                        for fi38 := 0; fi38 < 10; fi38++ {
                            v5 += (v5 * v5)
                            var v39 uint32 = 2147483648
                            v3 = ((0.0 * v3) - (((0.1 - v3) / float32(float64(1e10))) - v3))
                            _ = fi38
                            _ = v39
                        }
                    }
                } else {
                    switch v1 {
                    case "foo bar":
                        v17 = (v17 * 2)
                    }
                    v1 = (v1 + (("\n\t" + v1) + "foo bar"))
                    v40 := v4[2147483647]
                    _ = v40
                }
            case "foo bar":
                v41 := GenS1{F0: 649, F1: (v21 % ((v17) | 1)), F2: v29}
                _ = v41
            }
            if len(v2) > 0 {
                v2[0] = v20
            }
            _ = fi28
            _ = v29
        }
    case 32768:
        fmt.Println(v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
