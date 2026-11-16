package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 uint32; F1 int32; F2 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8; F1 int32 }
func genFn0(p0 int16) uint32 {
    p0 = p0
    v6 := []int{0}
    _ = p0
    _ = v6
    return 0
}

func main() {
    var v1 float32 = float32(0.0)
    var v2 int64 = -1
    var v3 uint32 = 4294967295
    var v4 uint8 = 255
    var v5 int16 = -625
    if (100 != -128) {
        if true {
            v5 -= v5
        } else {
            if false {
                if ((283 < 1) || false) {
                    fmt.Println(v3, v1, v3)
                } else {
                    v7 := GenS1{F0: false, F1: genFn0(v5)}
                    v8 := GenS2{F0: v3, F1: 1, F2: v2}
                    fmt.Println(v4)
                    _ = v7
                    _ = v8
                }
            } else {
                v9 := []int{32767, 426}
                _ = v9
            }
        }
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        v11 := GenS3{F0: v5}
        for fi12 := 0; fi12 < 11; fi12++ {
            v13 := GenS3{F0: v5}
            if (0 >= 99) {
                var v14 float32 = 3.14159
                v15 := GenS3{F0: (v5 << ((255) & 63))}
                _ = v14
                _ = v15
            }
            if (true || (true || (true || false))) {
                switch fi10 {
                case 1:
                    switch (("foo bar" + ("\n\t" + "a")) + "foo bar") {
                    case "\n\t":
                        var v16 int = fi12
                        _ = v16
                    default:
                        v17 := GenS3{F0: int16(v2)}
                        _ = v17
                    }
                case 4294967295:
                    var v18 bool = (42 == 127)
                    _ = v18
                case 32768:
                    if false {
                        var v19 uint = 65535
                        switch fi12 {
                        case 2147483647:
                            v19 -= v19
                        case 4294967295:
                            if true {
                                v20 := []int{127, (fi10 | int(v2))}
                                _ = v20
                            }
                        default:
                            if (!true) {
                                v19 -= (v19 / (((v19 << (((^(v19))) & 63))) | 1))
                                v1 = 1.5
                                var v21 uint = (v19 - v19)
                                _ = v21
                            } else {
                                v22 := GenS4{F0: 0, F1: -1}
                                v19 += (^(v19))
                                v2 -= -1
                                _ = v22
                            }
                        }
                        _ = v19
                    } else {
                        if (("hello" == "foo bar") && ((32767 <= 65535) || (uint(v3) <= 4294967295))) {
                            v23 := GenS2{F0: v3, F1: 1, F2: (v2 % ((v2) | 1))}
                            _ = v23
                        } else {
                            v3 = genFn0(v5)
                            switch 32767 {
                            case -948:
                                fmt.Println(fi12)
                            case 1:
                                v3 -= v3
                            case 127:
                                var v24 int = (fi10 << ((335) & 63))
                                _ = v24
                            default:
                                v25 := []int{fi10}
                                _ = v25
                            }
                        }
                    }
                }
                v26 := GenS1{F0: (v4 <= v4), F1: genFn0(32767)}
                _ = v26
            } else {
                var v27 float64 = 3.14159
                v3 -= v3
                switch "\n\t" {
                case "hello":
                    var v28 uint64 = 497
                    _ = v28
                }
                _ = v27
            }
            _ = fi12
            _ = v13
        }
        v29 := []int{(^(fi10)), fi10, (fi10 / ((2147483647) | 1)), fi10}
        _ = fi10
        _ = v11
        _ = v29
    }
    for fi30 := 0; fi30 < 1; fi30++ {
        v31 := []int{-874, fi30, (^(fi30)), (fi30 >> ((4294967295) & 63))}
        _ = fi30
        _ = v31
    }
    switch 0 {
    case 32767:
        if true {
            fmt.Println(v2, v1)
            var v32 string = ""
            _ = v32
        }
    }
    v2 += v2
    v33 := []int{32768, 955, 127, 482}
    v2 -= v2
    switch ("" + "\n\t") {
    case "":
        var v34 int64 = v2
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
