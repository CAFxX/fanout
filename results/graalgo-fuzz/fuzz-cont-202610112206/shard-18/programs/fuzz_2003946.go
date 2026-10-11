package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 int32 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 int32; F1 uint64 }
func genFn0(p0 bool, p1 uint32, p2 int64) string {
    v6 := []int{-9223372036854775808, int(p1)}
    p2 = (^(p2))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return ((((("foo bar" + "foo bar") + "a") + "a") + ("a" + "hello")) + "")
}

func main() {
    var v1 uint8 = 2
    var v2 uint64 = 1
    v3 := []uint32{}
    v4 := make(map[uint32]int)
    var v5 bool = (!true)
    v7 := GenS1{F0: (v2 & (v2 ^ v2)), F1: 696, F2: 42}
    _ = len(v4)
    v3 = append(v3, 1)
    var v8 uint8 = v1
    var v9 int8 = 127
    v5 = true
    switch 32768 {
    case 32768:
        for fi10 := 0; fi10 < 15; fi10++ {
            if false {
                if (((0 != -1) || (int64(fi10) >= 0)) || (128 <= v8)) {
                    var v11 int16 = 1000
                    _ = v11
                }
            } else {
                var v12 bool = true
                _ = v12
            }
            _ = fi10
        }
    default:
        v5 = (0.1 <= 1.5)
    }
    var v13 uint32 = 1
    v14 := GenS1{F0: (v2 | 1), F1: 1, F2: -600}
    if false {
        v15 := GenS2{F0: (-779 <= -285), F1: uint8(v13)}
        var v16 string = ""
        if v5 {
            for fi17 := 0; fi17 < 18; fi17++ {
                v13 = v13
                v18 := GenS2{F0: v5, F1: (v8 + v1)}
                _ = fi17
                _ = v18
            }
        }
        _ = v15
        _ = v16
    }
    if v5 {
        v8 = (v8 & 2)
        var v19 float32 = (((3.14159 + float32(0.1)) - 0.0) + 1e10)
        _ = v19
    } else {
        var v20 int32 = 42
        v21 := GenS1{F0: 18446744073709551615, F1: v20, F2: 512}
        _ = v20
        _ = v21
    }
    v22 := GenS1{F0: (^(v2)), F1: 903, F2: 0}
    v23 := GenS1{F0: v2, F1: 0, F2: -1}
    if ((false || (v5 && (-1 != -1))) && ((-410 <= 503) && (!(43 > (v8 | (v8 >> ((1) & 63))))))) {
        for fi24 := 0; fi24 < 2; fi24++ {
            v25 := GenS2{F0: true, F1: (v8 << ((886) & 63))}
            var v26 int8 = (^(v9))
            _ = fi24
            _ = v25
            _ = v26
        }
        var v27 uint = 65535
        _ = v27
    } else {
        v13 -= 0
        v28 := v4[(^(v13))]
        _ = v28
    }
    for fi29 := 0; fi29 < 9; fi29++ {
        switch (fi29 >> ((uint(v2)) & 63)) {
        case -2147483648:
            if true {
                _ = len(v3)
                fmt.Println(v1)
            }
        default:
            if ((v2 << ((18446744073709551615) & 63)) == v2) {
                if v5 {
                    var v30 uint32 = 764
                    v31 := v4[(v30 & (^(v30)))]
                    var v32 uint32 = v13
                    _ = v30
                    _ = v31
                    _ = v32
                }
                switch "" {
                case "a":
                    for fi33 := 0; fi33 < 6; fi33++ {
                        switch (("" + "foo bar") + "") {
                        case "":
                            v34 := GenS2{F0: v5, F1: (v1 | v8)}
                            _ = v34
                        case "hello":
                            delete(v4, 2)
                        case "foo bar":
                            if false {
                                var v35 uint64 = 18446744073709551615
                                v36 := v4[(v13 % (((v13 + v13)) | 1))]
                                _ = v35
                                _ = v36
                            }
                        }
                        v8 = v1
                        _ = fi33
                    }
                case "":
                    switch (^(fi29)) {
                    case 491:
                        for fi37 := 0; fi37 < 5; fi37++ {
                            v5 = ((v2 ^ (v2 << ((18446744073709551615) & 63))) != (v2 & v2))
                            v38 := GenS3{F0: -811, F1: v2}
                            _ = fi37
                            _ = v38
                        }
                    }
                }
            }
        }
        _ = fi29
    }
    if (!v5) {
        var v39 uint = 42
        v40 := GenS1{F0: (v2 ^ 905), F1: -2147483648, F2: 2147483647}
        var v41 int16 = -1
        _ = v39
        _ = v40
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
