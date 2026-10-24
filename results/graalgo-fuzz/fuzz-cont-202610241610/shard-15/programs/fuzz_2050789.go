package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 uint64; F1 uint64 }
func genFn0() bool {
    var v5 int8 = 1
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := []int{fi6, (fi6 / ((42) | 1)), (fi6 * fi6)}
        _ = fi6
        _ = v7
    }
    _ = v5
    return false
}

func main() {
    var v1 string = "\n\t"
    var v2 int = 42
    v3 := make(map[int]uint32)
    var v4 float32 = (0.0 - 3.14159)
    var v8 float32 = v4
    v1 = ((v1 + v1) + v1)
    for fi9 := 0; fi9 < 19; fi9++ {
        switch (v1 + (("" + v1) + "foo bar")) {
        case "hello":
            switch 643 {
            case -2147483648:
                v10 := GenS1{F0: 1, F1: v2}
                _ = v10
            }
        case "":
            v11 := GenS1{F0: 908, F1: 42}
            _ = v11
        default:
            v4 = ((0.1 - (v4 + 1.5)) * ((((v4 * v8) - v8) + 1e10) / ((v8 * (v8 / v8)) / v4)))
        }
        var v12 bool = true
        v4 = (v4 + v4)
        _ = fi9
        _ = v12
    }
    var v13 float64 = 1e10
    if false {
        v14 := GenS2{F0: 1}
        v15 := []int{(v2 << ((349) & 63)), -2147483648}
        _ = v14
        _ = v15
    }
    v8 += v4
    for fi16 := 0; fi16 < 8; fi16++ {
        if (false && ((false || true) || (3.14159 < (float32((v13 + (0.0 + v13))) - (3.14159 / 1.0))))) {
            if false {
                for fi17 := 0; fi17 < 8; fi17++ {
                    switch (((v1 + v1) + (v1 + v1)) + "foo bar") {
                    case "a":
                        var v18 int = (^(fi16))
                        _ = v18
                    default:
                        if (v2 <= -1) {
                            var v19 int64 = 9223372036854775807
                            var v20 uint8 = 42
                            v20 = (v20 % ((127) | 1))
                            _ = v19
                            _ = v20
                        }
                    }
                    v2 = -866
                    var v21 uint8 = 0
                    _ = fi17
                    _ = v21
                }
            }
            if (!true) {
                if true {
                    var v22 float32 = (v4 - 1.5)
                    _ = v22
                } else {
                    var v23 int32 = 460
                    _ = v23
                }
                for fi24 := 0; fi24 < 20; fi24++ {
                    fmt.Println(v8, fi24)
                    _ = fi24
                }
                v25 := v3[0]
                _ = v25
            } else {
                fmt.Println(fi16, fi16)
                for fi26 := 0; fi26 < 19; fi26++ {
                    var v27 int8 = 0
                    _ = fi26
                    _ = v27
                }
            }
        } else {
            if (502 > 735) {
                v28 := []int{330, (v2 | (fi16 | (fi16 * v2))), (-(v2)), (fi16 | 127)}
                v1 = v1
                var v29 float64 = (v13 - 1.5)
                _ = v28
                _ = v29
            }
            v2 += 127
        }
        _ = fi16
    }
    var v30 int8 = 23
    for fi31 := 0; fi31 < 15; fi31++ {
        var v32 uint64 = 2
        fmt.Println(v4, v2)
        _ = fi31
        _ = v32
    }
    var v33 int64 = -982
    var v34 uint = 42
    if (-496 != int32(v33)) {
        v30 = 1
    }
    v3[v2] = 33
    v35 := []int{718, v2}
    fmt.Println(v8)
    for fi36 := 0; fi36 < 16; fi36++ {
        if len(v35) > 0 {
            v35[0] = 42
        }
        if (!(v33 < 1)) {
            if (!true) {
                v1 = v1
                for fi37 := 0; fi37 < 20; fi37++ {
                    delete(v3, (fi37 / (((fi37 | (fi37 & 2147483647))) | 1)))
                    _ = fi37
                }
                v4 = 1.5
            }
            if false {
                v38 := GenS3{F0: 32767, F1: (((v33 % (((v33 + v33)) | 1)) < 42) || (true || false))}
                _ = v38
            } else {
                var v39 int16 = -991
                fmt.Println(fi36, v4)
                _ = v39
            }
            if (!((float64((0.1 / (v8 / v4))) / (v13 * v13)) > v13)) {
                switch fi36 {
                case -1:
                    if len(v35) >= 2 {
                        v35 = v35[1:len(v35)]
                    }
                case 42:
                    if true {
                        if len(v35) > 0 {
                            v35[0] = (fi36 * fi36)
                        }
                        v40 := GenS3{F0: -32768, F1: (v30 == 40)}
                        for fi41 := 0; fi41 < 16; fi41++ {
                            v8 = float32(v13)
                            if (1 < 0) {
                                switch v2 {
                                case -1:
                                    v42 := GenS4{F0: 4294967295, F1: 2}
                                    _ = v42
                                case 9223372036854775807:
                                    var v43 uint = v34
                                    _ = v43
                                case 32767:
                                    var v44 int = (fi41 * fi36)
                                    _ = v44
                                }
                                v8 += 1e10
                            }
                            fmt.Println(v34)
                            var v45 int8 = v30
                            var v46 uint32 = 674
                            _ = fi41
                            _ = v45
                            _ = v46
                        }
                        _ = v40
                    }
                case 127:
                    if len(v35) >= 2 {
                        v35 = v35[1:len(v35)]
                    }
                default:
                    _ = len(v35)
                }
                var v47 uint32 = 65535
                _ = v47
            }
        }
        v48 := v3[(-(v2))]
        _ = fi36
        _ = v48
    }
    v1 = (v1 + ("foo bar" + v1))
    var v49 uint8 = 237
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v49
    _ = v8
}
