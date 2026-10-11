package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 int16; F2 uint64 }
type GenS3 struct { F0 string; F1 uint16 }

func main() {
    var v1 uint16 = 616
    var v2 int8 = -108
    v3 := []uint32{uint32(v1), 2}
    for fi4 := 0; fi4 < 8; fi4++ {
        v1 = v1
        var v5 int32 = -2147483648
        switch -220 {
        case 309:
            v6 := GenS1{F0: v2}
            _ = v6
        case -9223372036854775808:
            v2 = int8(v1)
        }
        _ = fi4
        _ = v5
    }
    for fi7 := 0; fi7 < 10; fi7++ {
        v2 -= v2
        _ = fi7
    }
    for fi8 := 0; fi8 < 6; fi8++ {
        v9 := GenS1{F0: (v2 / ((42) | 1))}
        if len(v3) > 0 {
            v3[0] = 2
        }
        for fi10 := 0; fi10 < 3; fi10++ {
            fmt.Println(fi10, fi8)
            v1 -= (v1 & v1)
            fmt.Println(fi10, v2, fi10)
            _ = fi10
        }
        _ = fi8
        _ = v9
    }
    v1 += 2
    var v11 uint64 = 2
    switch -163 {
    case 0:
        v2 = v2
    case 32767:
        if (true && false) {
            for fi12 := 0; fi12 < 14; fi12++ {
                if (334 == 4294967295) {
                    fmt.Println(v2)
                }
                _ = fi12
            }
            if len(v3) > 0 {
                v3[0] = 503
            }
        }
    }
    v13 := GenS1{F0: (-(v2))}
    v11 -= (v11 % ((v11) | 1))
    var v14 uint32 = 65535
    var v15 uint32 = (v14 / (((^(v14))) | 1))
    if (42 == 0) {
        for fi16 := 0; fi16 < 3; fi16++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            for fi17 := 0; fi17 < 3; fi17++ {
                v18 := GenS1{F0: v2}
                if ((!(true && (18446744073709551615 >= 4294967295))) && (2 >= v14)) {
                    v19 := GenS1{F0: 1}
                    switch 128 {
                    case 0:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case -780:
                        if ((((true || (!true)) || true) || true) && ((-1 < 42) || (78 != 72))) {
                            var v20 uint8 = 2
                            fmt.Println(v14, fi16, fi16)
                            var v21 uint = 1
                            fmt.Println(fi16)
                            _ = v20
                            _ = v21
                        } else {
                            v22 := GenS1{F0: (v2 ^ v2)}
                            var v23 int64 = -9223372036854775808
                            fmt.Println(v2, fi17)
                            _ = v22
                            _ = v23
                        }
                    default:
                        if (2147483647 >= 0) {
                            var v24 float32 = (1.5 - float32(float64(0.0)))
                            var v25 int64 = 42
                            v26 := GenS2{F0: -35, F1: -287, F2: v11}
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                    }
                    v3 = append(v3, 42)
                    _ = v19
                }
                _ = fi17
                _ = v18
            }
            _ = fi16
        }
        var v27 int32 = -2147483648
        _ = v27
    } else {
        v11 += (^(v11))
        v3 = append(v3, v15)
        var v28 uint8 = 42
        _ = v28
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v29 := GenS3{F0: "", F1: (v1 << ((2) & 63))}
    _ = len(v3)
    var v30 uint16 = 1
    v11 += v11
    fmt.Println(v15, v11, v15)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v29
    _ = v3
    _ = v30
}
