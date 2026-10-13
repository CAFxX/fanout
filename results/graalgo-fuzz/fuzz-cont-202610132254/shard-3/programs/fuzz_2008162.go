package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 bool }
type GenS2 struct { F0 uint; F1 uint32; F2 uint32 }
type GenS3 struct { F0 uint64; F1 int; F2 uint64 }
type GenS4 struct { F0 int8; F1 int }
func genFn0() string {
    if false {
        v7 := []int{794, 127, -437, 1}
        var v8 int32 = -845
        switch ("\n\t" + "\n\t") {
        case "foo bar":
            for fi9 := 0; fi9 < 15; fi9++ {
                for fi10 := 0; fi10 < 17; fi10++ {
                    if (503 >= 4294967295) {
                        if len(v7) > 0 {
                            v7[0] = fi9
                        }
                        var v11 string = (((("\n\t" + "hello") + ("\n\t" + "\n\t")) + "a") + "")
                        if (-128 < int8(fi9)) {
                            v11 = (v11 + v11)
                        } else {
                            switch fi9 {
                            case 695:
                                fmt.Println(fi10)
                            case -245:
                                v12 := GenS1{F0: 154, F1: -94, F2: (161 < 0)}
                                _ = v12
                            case 127:
                                var v13 float64 = (1.5 + (float64((0.0 + (3.14159 - (((1e10 / 0.1) / 1.5) - 1.5)))) * (1e10 * (float64(3.14159) * 1e10))))
                                _ = v13
                            default:
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                            }
                            var v14 int = (-(fi10))
                            v15 := GenS1{F0: 152, F1: 731, F2: (-2.25 == (1e10 * 3.14159))}
                            _ = v14
                            _ = v15
                        }
                        v16 := GenS1{F0: 4294967295, F1: -1, F2: (!(-1 != 0))}
                        v17 := GenS1{F0: 383, F1: 59, F2: true}
                        if (!true) {
                            v8 = 1
                        } else {
                            if true {
                                if true {
                                    var v18 int = fi10
                                    v7 = append(v7, 32768)
                                    v7 = append(v7, fi9)
                                    if len(v7) > 0 {
                                        v7[0] = -264
                                    }
                                    _ = v18
                                }
                                switch v11 {
                                case "a":
                                    var v19 bool = false
                                    v7 = append(v7, (fi10 * fi10))
                                    v8 -= v8
                                    _ = v19
                                default:
                                    v8 += (^(v8))
                                }
                                if false {
                                    if false {
                                        var v20 uint16 = 65535
                                        fmt.Println(v8, v11, fi9)
                                        _ = v20
                                    }
                                    v11 = v11
                                    v7 = append(v7, (fi10 << ((177) & 63)))
                                }
                            }
                        }
                        _ = v11
                        _ = v16
                        _ = v17
                    }
                    if true {
                        var v21 bool = true
                        var v22 int16 = 0
                        fmt.Println(fi10)
                        _ = v21
                        _ = v22
                    } else {
                        var v23 int32 = v8
                        v23 = -2147483648
                        var v24 int32 = (^(v8))
                        _ = v23
                        _ = v24
                    }
                    _ = fi10
                }
                _ = fi9
            }
        }
        _ = v7
        _ = v8
    }
    v25 := GenS1{F0: 2, F1: -246, F2: (2147483648 >= 889)}
    v26 := GenS2{F0: 877, F1: 2147483648, F2: 703}
    if ((-2.25 - 0.0) >= (((1.5 / -2.25) + 3.14159) - 3.14159)) {
        var v27 int8 = 79
        v28 := GenS3{F0: 793, F1: -485, F2: 0}
        var v29 string = "a"
        _ = v27
        _ = v28
        _ = v29
    }
    _ = v25
    _ = v26
    return ""
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint32 = 65535
    var v3 uint64 = 2
    var v4 uint64 = 317
    var v5 int8 = int8(v3)
    v6 := []int{-9223372036854775808}
    v4 += 4294967295
    var v30 float32 = -2.25
    v31 := GenS4{F0: (v5 & (v5 % ((v5) | 1))), F1: 42}
    var v32 uint32 = (v2 * (v2 - 149))
    v6 = append(v6, -2147483648)
    for fi33 := 0; fi33 < 8; fi33++ {
        for fi34 := 0; fi34 < 7; fi34++ {
            v3 -= v3
            _ = fi34
        }
        _ = len(v1)
        _ = fi33
    }
    fmt.Println(v30, v30, v2)
    if (false && (127 <= 0)) {
        v6 = append(v6, 623)
    }
    switch "hello" {
    case "":
        if false {
            v35 := v1[v2]
            _ = v35
        }
    case "\n\t":
        v32 = v32
    default:
        if ((v4 - (^(v3))) == v4) {
            var v36 float32 = (float32(3.14159) * 1.5)
            var v37 uint16 = 371
            if (0.0 > ((float64(v30) - -2.25) + ((0.1 * (1.5 + 0.1)) / (1e10 - ((-2.25 / 1.0) * 3.14159))))) {
                v6 = append(v6, 1)
            }
            _ = v36
            _ = v37
        } else {
            v32 += v2
        }
    }
    fmt.Println(v30, v2, v2)
    v1[v2] = "a"
    v38 := v1[(v2 / (((v2 % ((2147483648) | 1))) | 1))]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
