package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 bool; F2 uint8 }
type GenS4 struct { F0 int64 }
func genFn0() uint16 {
    var v7 uint64 = 1
    var v8 uint64 = v7
    _ = v7
    _ = v8
    return 32767
}

func main() {
    var v1 int16 = -422
    var v2 int = 128
    v3 := make(map[uint32]int)
    v4 := make(map[string]int)
    v5 := make(map[string]string)
    var v6 int64 = -9223372036854775808
    v2 += 32768
    _ = len(v3)
    v2 -= v2
    if (((1e10 * 0.0) + float32((0.1 + -2.25))) >= 0.1) {
        if (true || (0 > 454)) {
            for fi9 := 0; fi9 < 18; fi9++ {
                for fi10 := 0; fi10 < 18; fi10++ {
                    if false {
                        v11 := GenS1{F0: 255}
                        var v12 float32 = 0.0
                        v13 := []int{(^(fi9)), 737}
                        _ = v11
                        _ = v12
                        _ = v13
                    } else {
                        v6 -= (v6 / ((v6) | 1))
                    }
                    if true {
                        if (!(!false)) {
                            v14 := GenS2{F0: 1}
                            if (9223372036854775807 >= fi10) {
                                v2 = (fi9 >> ((255) & 63))
                            } else {
                                v15 := GenS3{F0: v6, F1: (!(2147483648 < 2)), F2: 255}
                                _ = v15
                            }
                            _ = v14
                        }
                        var v16 uint32 = 65535
                        v17 := GenS2{F0: 0}
                        _ = v16
                        _ = v17
                    } else {
                        if (false || true) {
                            v6 += v6
                        } else {
                            v2 = fi10
                            var v18 float32 = 0.1
                            v6 += v6
                            v19 := GenS4{F0: v6}
                            _ = v18
                            _ = v19
                        }
                        v20 := GenS4{F0: 1}
                        _ = v20
                    }
                    if (20 <= -3) {
                        var v21 int16 = 0
                        v2 -= 4294967295
                        v21 += (v1 + -32768)
                        _ = v21
                    } else {
                        v22 := GenS3{F0: -16, F1: (true && (0 != 65535)), F2: 42}
                        v1 = -32768
                        _ = v22
                    }
                    _ = fi10
                }
                fmt.Println(v2, v2, v1)
                var v23 int8 = -82
                _ = fi9
                _ = v23
            }
        } else {
            v2 = (-(v2))
            v1 = (v1 << ((4294967295) & 63))
            switch (v2 + v2) {
            case 42:
                delete(v3, 1)
            case 128:
                if (!false) {
                    for fi24 := 0; fi24 < 8; fi24++ {
                        v25 := []int{127, (fi24 | (fi24 | (fi24 ^ v2))), fi24, 42}
                        _ = len(v4)
                        var v26 uint16 = 1
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                    fmt.Println(v6, v1, v2)
                }
            default:
                var v27 int8 = 127
                _ = v27
            }
        }
        if false {
            var v28 int64 = (v6 + (v6 >> ((4294967295) & 63)))
            _ = v28
        }
        v1 = (^(v1))
    } else {
        var v29 uint64 = 18446744073709551615
        var v30 int32 = 1
        fmt.Println(v6, v30)
        _ = v29
        _ = v30
    }
    v6 -= (v6 / ((-976) | 1))
    v2 = 0
    v31 := GenS3{F0: v6, F1: false, F2: 2}
    var v32 uint32 = 804
    v33 := GenS3{F0: int64(v32), F1: (!false), F2: 128}
    fmt.Println(v1, v6, v6)
    v32 = v32
    if false {
        v34 := []int{9223372036854775807, (v2 & -9223372036854775808)}
        v2 = v2
        _ = v34
    } else {
        v35 := GenS2{F0: 1}
        _ = v35
    }
    for fi36 := 0; fi36 < 2; fi36++ {
        fmt.Println(v6, fi36, v6)
        _ = fi36
    }
    var v37 uint = 1
    var v38 int32 = 2147483647
    fmt.Println(v32, v32, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
