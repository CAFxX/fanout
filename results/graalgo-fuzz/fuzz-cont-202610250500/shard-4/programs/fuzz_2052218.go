package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 uint64; F2 uint64 }
type GenS4 struct { F0 uint32; F1 string; F2 uint16 }
type GenS5 struct { F0 uint32; F1 int8; F2 int16 }
type GenS6 struct { F0 int16; F1 uint32; F2 int8 }
func genFn0(p0 uint32) string {
    var v7 uint8 = 42
    v8 := GenS1{F0: 18446744073709551615}
    _ = p0
    _ = v7
    _ = v8
    return "foo bar"
}

func main() {
    var v1 uint64 = 4294967295
    var v2 int16 = -734
    var v3 uint = 1
    var v4 uint64 = v1
    var v5 string = "hello"
    v6 := make(map[int]string)
    v9 := GenS2{F0: -128}
    var v10 uint = v3
    v5 = "a"
    switch 127 {
    case 128:
        for fi11 := 0; fi11 < 1; fi11++ {
            v12 := []int{}
            fmt.Println(fi11)
            if len(v12) > 0 {
                v12[0] = fi11
            }
            _ = fi11
            _ = v12
        }
    case -1:
        var v13 float64 = (3.14159 * 1.5)
        _ = v13
    }
    var v14 uint8 = 137
    v3 -= v3
    v15 := GenS2{F0: 0}
    switch v5 {
    case "a":
        switch v5 {
        case "\n\t":
            v16 := GenS3{F0: 244, F1: (v4 % ((v4) | 1)), F2: (v4 / (((^(v1))) | 1))}
            _ = v16
        }
    case "hello":
        delete(v6, 32768)
    case "":
        if true {
            switch v5 {
            case "a":
                var v17 float64 = float64(float32((0.1 - 3.14159)))
                _ = v17
            }
            switch v5 {
            case "a":
                for fi18 := 0; fi18 < 15; fi18++ {
                    v19 := GenS4{F0: 65535, F1: v5, F2: 0}
                    var v20 int8 = 1
                    if (float32(3.14159) < 1.5) {
                        switch 128 {
                        case 128:
                            v21 := GenS5{F0: 2147483648, F1: (v20 * (v20 - (-(v20)))), F2: (^(v2))}
                            _ = v21
                        case -296:
                            switch ("hello" + v5) {
                            case "":
                                v22 := []int{fi18, (^(fi18))}
                                _ = v22
                            case "a":
                                v10 = (v3 >> (((v3 + 255)) & 63))
                            default:
                                for fi23 := 0; fi23 < 4; fi23++ {
                                    var v24 uint16 = 1
                                    _ = fi23
                                    _ = v24
                                }
                            }
                        }
                    }
                    _ = fi18
                    _ = v19
                    _ = v20
                }
            case "foo bar":
                var v25 int64 = -9223372036854775808
                _ = v25
            case "\n\t":
                v26 := GenS6{F0: (^(v2)), F1: 2147483648, F2: -128}
                _ = v26
            default:
                var v27 uint64 = (v4 * v4)
                _ = v27
            }
            v28 := v6[885]
            _ = v28
        } else {
            var v29 int = 128
            for fi30 := 0; fi30 < 11; fi30++ {
                if false {
                    v2 = (v2 - v2)
                    switch v5 {
                    case "foo bar":
                        v4 = (^(v4))
                    case "a":
                        var v31 float32 = -2.25
                        _ = v31
                    case "\n\t":
                        var v32 int8 = 0
                        _ = v32
                    }
                    v1 += (^(v1))
                } else {
                    fmt.Println(v5, v2)
                    for fi33 := 0; fi33 < 17; fi33++ {
                        if (!true) {
                            switch (v29 & (v29 + (^(fi30)))) {
                            case 9223372036854775807:
                                if ((("a" + v5) + ((v5 + "foo bar") + v5)) == v5) {
                                    fmt.Println(v3, v5)
                                    if true {
                                        fmt.Println(v29, fi30, v2)
                                        v10 -= uint(v14)
                                        v29 -= (^(fi30))
                                        fmt.Println(fi33, v3, fi33)
                                    }
                                    v10 += (v10 ^ (v3 << (((v3 / (((v10 + 4294967295)) | 1))) & 63)))
                                }
                            case 341:
                                var v34 uint32 = 2
                                _ = v34
                            case 0:
                                v2 = v2
                            default:
                                switch 32768 {
                                case -1:
                                    fmt.Println(v14, v5)
                                case -2147483648:
                                    fmt.Println(v10, v2)
                                }
                            }
                            var v35 uint = v10
                            _ = v35
                        }
                        _ = fi33
                    }
                }
                var v36 string = (v5 + "\n\t")
                if (2 != 65535) {
                    var v37 string = (v36 + "\n\t")
                    _ = v37
                }
                _ = fi30
                _ = v36
            }
            if true {
                var v38 int32 = 758
                v6[-713] = v5
                if true {
                    fmt.Println(v4, v10, v29)
                    switch 128 {
                    case -1:
                        if ((v29 & (v29 << ((432) & 63))) != v29) {
                            for fi39 := 0; fi39 < 13; fi39++ {
                                v29 -= fi39
                                v40 := []int{v29, fi39, (^(v29))}
                                for fi41 := 0; fi41 < 5; fi41++ {
                                    fmt.Println(v29, v14, v10)
                                    var v42 uint16 = uint16(v3)
                                    var v43 string = v5
                                    _ = fi41
                                    _ = v42
                                    _ = v43
                                }
                                _ = fi39
                                _ = v40
                            }
                        } else {
                            v44 := GenS4{F0: 2, F1: "a", F2: 65535}
                            _ = v44
                        }
                    case 9223372036854775807:
                        v5 = v5
                    }
                }
                _ = v38
            } else {
                if (1.5 < (1e10 / ((float32(-2.25) / -2.25) + ((0.0 / 1.0) - float32(0.0))))) {
                    var v45 uint32 = 2
                    v1 = 18446744073709551615
                    v46 := []int{(v29 & int(v14)), int(v2), (v29 % ((v29) | 1)), (v29 * v29)}
                    _ = v45
                    _ = v46
                }
            }
            _ = v29
        }
    }
    if false {
        var v47 uint64 = 436
        _ = v47
    }
    v48 := GenS6{F0: (v2 & -32768), F1: 2147483647, F2: -59}
    var v49 int = 127
    fmt.Println(v10, v1, v14)
    v49 -= v49
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v6
    _ = v9
}
