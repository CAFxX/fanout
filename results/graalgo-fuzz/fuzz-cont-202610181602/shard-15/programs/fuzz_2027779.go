package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 int; F1 int32; F2 string }
type GenS3 struct { F0 uint8; F1 int32; F2 int64 }
func genFn0(p0 bool, p1 uint32, p2 uint) uint16 {
    p1 = p1
    var v5 uint = p2
    p2 += 255
    var v6 uint32 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 65535
}

func main() {
    v1 := []uint32{1, 785, 0, 42}
    v2 := []uint32{983, 0, 65535}
    var v3 float32 = 0.0
    var v4 string = "hello"
    var v7 uint64 = 18446744073709551615
    v2 = append(v2, 65535)
    if false {
        fmt.Println(v4, v3, v7)
        if true {
            if len(v2) > 0 {
                v2[0] = 1
            }
            v3 = v3
        }
        var v8 int16 = 0
        _ = v8
    } else {
        v9 := GenS1{F0: 65535, F1: 839}
        for fi10 := 0; fi10 < 13; fi10++ {
            v3 += ((0.1 / 1.5) + (v3 - float32(1.5)))
            _ = fi10
        }
        _ = v9
    }
    v11 := GenS2{F0: -69, F1: 0, F2: v4}
    v7 = v7
    for fi12 := 0; fi12 < 18; fi12++ {
        fmt.Println(fi12, v7, fi12)
        v13 := GenS2{F0: fi12, F1: -1, F2: (v4 + "hello")}
        v4 = (v4 + "\n\t")
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 5; fi14++ {
        fmt.Println(fi14, v3, v3)
        var v15 uint16 = 255
        _ = fi14
        _ = v15
    }
    v7 = (v7 % (((v7 >> ((0) & 63))) | 1))
    v7 = (v7 & (v7 % (((v7 >> ((18446744073709551615) & 63))) | 1)))
    _ = len(v2)
    if ((((0.0 - 0.1) / 1.0) >= (1e10 * 1.5)) || ((!(((v4 + v4) < v4) && false)) && false)) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        switch (v4 + ((v4 + "\n\t") + v4)) {
        case "foo bar":
            switch 11 {
            case 127:
                var v16 int16 = 42
                _ = v16
            case 32767:
                fmt.Println(v3, v7, v4)
            case 881:
                if true {
                    if true {
                        if false {
                            v2 = append(v2, 4294967295)
                            v4 = (v4 + ("hello" + (("" + (v4 + (v4 + v4))) + v4)))
                            var v17 uint64 = (v7 | v7)
                            _ = v17
                        } else {
                            v3 = (((0.1 / ((((v3 + (0.0 * float32(0.1))) / v3) + -2.25) / v3)) * v3) / v3)
                            if (479 > 4294967295) {
                                if true {
                                    if false {
                                        if len(v1) > 0 {
                                            v1[0] = 65535
                                        }
                                        switch (v4 + v4) {
                                        case "":
                                            v18 := GenS3{F0: 2, F1: -440, F2: 1}
                                            _ = v18
                                        case "hello":
                                            if true {
                                                for fi19 := 0; fi19 < 10; fi19++ {
                                                    v4 = v4
                                                    var v20 float32 = v3
                                                    v7 -= v7
                                                    _ = fi19
                                                    _ = v20
                                                }
                                            }
                                        case "foo bar":
                                            var v21 float64 = float64(v3)
                                            _ = v21
                                        default:
                                            fmt.Println(v4)
                                        }
                                    }
                                    if true {
                                        for fi22 := 0; fi22 < 4; fi22++ {
                                            var v23 float64 = 3.14159
                                            var v24 uint16 = 732
                                            v3 -= 1.5
                                            _ = fi22
                                            _ = v23
                                            _ = v24
                                        }
                                        var v25 uint8 = 2
                                        _ = v25
                                    }
                                    v7 += v7
                                }
                                v4 = v4
                            }
                        }
                    }
                    var v26 float32 = v3
                    for fi27 := 0; fi27 < 16; fi27++ {
                        var v28 uint64 = v7
                        _ = fi27
                        _ = v28
                    }
                    _ = v26
                } else {
                    for fi29 := 0; fi29 < 12; fi29++ {
                        if len(v2) > 0 {
                            v2[0] = 2147483648
                        }
                        v3 += v3
                        fmt.Println(fi29)
                        _ = fi29
                    }
                    v4 = ("a" + "hello")
                }
            }
        case "a":
            fmt.Println(v4)
        default:
            v2 = append(v2, 2147483648)
        }
        var v30 string = v4
        _ = v30
    }
    var v31 int8 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v7
}
