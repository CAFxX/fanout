package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint32; F1 int16; F2 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int32; F1 uint16; F2 uint }
func genFn0(p0 uint32, p1 int16, p2 string) int8 {
    if false {
        p2 = ""
    } else {
        var v6 int16 = p1
        _ = v6
    }
    p1 += (p1 >> ((65535) & 63))
    _ = p0
    _ = p1
    _ = p2
    return -1
}
func genFn1(p0 uint16, p1 uint64) uint {
    p0 = (^(p0))
    var v7 bool = (false && (p0 != (p0 >> ((784) & 63))))
    _ = p0
    _ = p1
    _ = v7
    return 18446744073709551615
}

func main() {
    v1 := []string{"a", ("a" + "foo bar"), "a"}
    var v2 int16 = -32768
    var v3 uint64 = 531
    var v4 uint32 = 802
    v5 := make(map[int]uint32)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (true || ((uint(v4) >= 259) && (-826 > 1))) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        var v8 bool = false
        var v9 bool = (float32(1e10) <= 3.14159)
        _ = v8
        _ = v9
    }
    v3 += (^(v3))
    switch -291 {
    case -677:
        for fi10 := 0; fi10 < 8; fi10++ {
            switch 127 {
            case -2147483648:
                switch "\n\t" {
                case "\n\t":
                    if true {
                        var v11 uint32 = (v4 & v4)
                        for fi12 := 0; fi12 < 1; fi12++ {
                            v13 := v5[fi12]
                            v14 := GenS1{F0: v13, F1: v2, F2: 255}
                            v4 += 42
                            v15 := GenS1{F0: (v4 ^ v11), F1: v2, F2: 128}
                            _ = fi12
                            _ = v13
                            _ = v14
                            _ = v15
                        }
                        if ((fi10 << ((118) & 63)) > (fi10 * (fi10 - (fi10 - (fi10 >> ((18446744073709551615) & 63)))))) {
                            if false {
                                var v16 uint8 = 42
                                v17 := GenS2{F0: (v11 << ((389) & 63)), F1: 32767, F2: 463}
                                switch 127 {
                                case -9223372036854775808:
                                    fmt.Println(v2, v3)
                                }
                                _ = v16
                                _ = v17
                            }
                        } else {
                            for fi18 := 0; fi18 < 6; fi18++ {
                                fmt.Println(v2, v4)
                                v19 := v5[(fi18 << ((154) & 63))]
                                var v20 float32 = float32(0.1)
                                _ = fi18
                                _ = v19
                                _ = v20
                            }
                            if (((0.0 + 0.1) == (1.5 * 0.0)) || true) {
                                if len(v1) > 0 {
                                    v1[0] = ""
                                }
                                v5[(-(fi10))] = (v4 | v11)
                            } else {
                                for fi21 := 0; fi21 < 19; fi21++ {
                                    v11 += 2147483647
                                    var v22 float64 = ((-2.25 - ((float64(((0.1 - 1.5) / (float32((0.0 - float64(float32((3.14159 - float64(1.5)))))) / 3.14159))) + 1e10) / (1e10 / -2.25))) - 3.14159)
                                    var v23 float32 = (1e10 / 3.14159)
                                    v24 := GenS2{F0: 477, F1: v2, F2: 36}
                                    var v25 uint32 = 0
                                    _ = fi21
                                    _ = v22
                                    _ = v23
                                    _ = v24
                                    _ = v25
                                }
                                switch -1 {
                                case 128:
                                    var v26 float64 = float64((1e10 / -2.25))
                                    _ = v26
                                }
                                var v27 bool = (true || (-2.25 > 1.5))
                                _ = v27
                            }
                        }
                        _ = v11
                    } else {
                        v28 := GenS2{F0: 1, F1: 42, F2: 65535}
                        _ = v28
                    }
                case "foo bar":
                    var v29 uint16 = 42
                    _ = v29
                }
            default:
                fmt.Println(v4, v4)
            }
            for fi30 := 0; fi30 < 9; fi30++ {
                fmt.Println(v3, fi10, fi30)
                var v31 string = "hello"
                var v32 int = fi10
                _ = fi30
                _ = v31
                _ = v32
            }
            _ = fi10
        }
    }
    v33 := GenS3{F0: -9223372036854775808}
    var v34 uint32 = v4
    v34 = (v34 / (((v34 / (((v34 - (v34 / ((350) | 1)))) | 1))) | 1))
    if (genFn1(0, v3) > 408) {
        v35 := GenS2{F0: v34, F1: (^(v2)), F2: 255}
        _ = v35
    }
    v36 := GenS4{F0: 528}
    switch -342 {
    case 42:
        v4 = v4
    }
    for fi37 := 0; fi37 < 13; fi37++ {
        v3 = v3
        _ = fi37
    }
    v38 := GenS5{F0: -247, F1: 65535, F2: 4294967295}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v36
    _ = v38
    _ = v4
    _ = v5
}
