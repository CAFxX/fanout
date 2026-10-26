package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 uint; F1 int64 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{9223372036854775807, 1, 32767, 2147483647}
    v3 := []uint32{42}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v4 int16 = 227
    v5 := GenS1{F0: "hello", F1: 42}
    var v6 float64 = (-2.25 + 1.5)
    var v7 string = "\n\t"
    delete(v1, 2147483648)
    if false {
        for fi8 := 0; fi8 < 3; fi8++ {
            fmt.Println(v7, v4, fi8)
            if true {
                if false {
                    if true {
                        v4 -= v4
                        var v9 uint32 = 4294967295
                        if false {
                            if false {
                                for fi10 := 0; fi10 < 11; fi10++ {
                                    v11 := GenS1{F0: v7, F1: 150}
                                    _ = fi10
                                    _ = v11
                                }
                            } else {
                                v6 -= ((-2.25 * 0.1) * v6)
                                if true {
                                    v6 += -2.25
                                    v12 := GenS1{F0: (((v7 + "a") + (((((("\n\t" + "foo bar") + "") + v7) + (("foo bar" + v7) + ("\n\t" + "a"))) + "") + v7)) + ("a" + ("hello" + v7))), F1: 388}
                                    _ = v12
                                }
                            }
                        } else {
                            fmt.Println(v6, fi8)
                        }
                        _ = v9
                    }
                } else {
                    delete(v1, 2147483648)
                    _ = len(v2)
                    var v13 int8 = 42
                    _ = v13
                }
                fmt.Println(v4)
            }
            _ = fi8
        }
    } else {
        if true {
            for fi14 := 0; fi14 < 17; fi14++ {
                var v15 float32 = 0.1
                var v16 float64 = (v6 + (1.5 - (3.14159 * float64(((v15 + float32(v6)) + float32((float64((1.5 + v15)) * 3.14159)))))))
                _ = fi14
                _ = v15
                _ = v16
            }
        }
    }
    switch 930 {
    case -756:
        for fi17 := 0; fi17 < 10; fi17++ {
            v6 -= v6
            var v18 uint32 = 4294967295
            _ = fi17
            _ = v18
        }
    case 127:
        switch 9223372036854775807 {
        case 128:
            var v19 float32 = float32(v6)
            _ = v19
        default:
            switch -397 {
            case -125:
                if false {
                    v6 += float64(1.5)
                } else {
                    v20 := GenS2{F0: 2, F1: 1}
                    var v21 uint64 = 18446744073709551615
                    for fi22 := 0; fi22 < 14; fi22++ {
                        v4 = v4
                        var v23 int32 = -2147483648
                        for fi24 := 0; fi24 < 4; fi24++ {
                            v21 += v21
                            if (false && (fi24 <= fi22)) {
                                v4 = 32767
                            } else {
                                fmt.Println(v23, v7, v7)
                                fmt.Println(v6)
                            }
                            v2 = append(v2, fi22)
                            _ = fi24
                        }
                        _ = fi22
                        _ = v23
                    }
                    _ = v20
                    _ = v21
                }
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
