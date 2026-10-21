package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int16, p1 string, p2 uint16) uint8 {
    if true {
        p2 = p2
        v4 := GenS1{F0: 99}
        switch "foo bar" {
        case "\n\t":
            v5 := []int{}
            _ = v5
        }
        _ = v4
    } else {
        p2 = 42
        for fi6 := 0; fi6 < 8; fi6++ {
            for fi7 := 0; fi7 < 10; fi7++ {
                var v8 int = fi7
                if false {
                    if (false && true) {
                        fmt.Println(v8, p0)
                        switch 32768 {
                        case -2147483648:
                            if false {
                                p2 = p2
                            }
                        case -1:
                            var v9 int16 = (p0 << ((311) & 63))
                            _ = v9
                        default:
                            v10 := GenS1{F0: 9223372036854775807}
                            _ = v10
                        }
                    } else {
                        if true {
                            if true {
                                v11 := GenS2{F0: 2}
                                v12 := []int{(fi6 - fi7)}
                                var v13 int8 = 1
                                _ = v11
                                _ = v12
                                _ = v13
                            } else {
                                v14 := GenS2{F0: 0}
                                if true {
                                    v8 -= v8
                                } else {
                                    var v15 int64 = -9223372036854775808
                                    _ = v15
                                }
                                p0 = -1
                                _ = v14
                            }
                        } else {
                            if (!false) {
                                v16 := []int{0, (^(fi7)), fi7, int(p2)}
                                _ = len(v16)
                                var v17 int64 = 9223372036854775807
                                _ = v16
                                _ = v17
                            }
                            var v18 int64 = 42
                            p1 = p1
                            _ = v18
                        }
                    }
                } else {
                    v19 := []int{(fi7 * fi6)}
                    _ = v19
                }
                var v20 bool = (9223372036854775807 >= -1)
                _ = fi7
                _ = v20
                _ = v8
            }
            _ = fi6
        }
        p0 += p0
    }
    if (!(0 >= 1)) {
        var v21 uint = 4294967295
        _ = v21
    }
    v22 := GenS2{F0: 861}
    var v23 bool = (42 == 518)
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    _ = v23
    return 127
}

func main() {
    var v1 float32 = 0.0
    var v2 uint = 187
    v3 := []uint32{652, 909}
    if (true || true) {
        for fi24 := 0; fi24 < 14; fi24++ {
            var v25 float32 = v1
            if len(v3) > 0 {
                v3[0] = 65535
            }
            for fi26 := 0; fi26 < 20; fi26++ {
                v2 += (v2 << ((1) & 63))
                _ = fi26
            }
            _ = fi24
            _ = v25
        }
    }
    var v27 string = ("a" + "\n\t")
    v27 = v27
    v1 += v1
    switch 127 {
    case 32767:
        var v28 int8 = -1
        _ = v28
    default:
        var v29 bool = (65535 < 170)
        _ = v29
    }
    for fi30 := 0; fi30 < 9; fi30++ {
        var v31 bool = (float64(((float32(3.14159) + (1.5 / (3.14159 - v1))) * (v1 / (((1e10 / float32((float64(v1) + 0.0))) + float32(3.14159)) * (-2.25 - ((v1 - v1) - (v1 / float32(1e10)))))))) < 3.14159)
        fmt.Println(v1, fi30)
        _ = fi30
        _ = v31
    }
    if (18446744073709551615 <= 18446744073709551615) {
        _ = len(v3)
        v3 = append(v3, 65535)
    }
    _ = len(v3)
    var v32 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v32
}
