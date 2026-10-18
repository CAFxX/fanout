package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 uint32; F2 uint32 }
func genFn0(p0 int64, p1 uint64) uint8 {
    var v6 uint8 = 1
    v7 := []int{1, -2147483648, 1}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 string, p1 string) uint64 {
    var v8 float32 = ((0.0 - (-2.25 + (((((1e10 * (3.14159 - (((((1e10 * 0.1) + 1.5) / (3.14159 / (3.14159 / float32(0.1)))) - 3.14159) / 0.1))) - 1.5) - (1.5 + float32(0.0))) / 1.5) / (1.5 - (0.1 - 1.5))))) - 0.0)
    var v9 float64 = -2.25
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 799
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float32 = 1.5
    var v3 uint32 = 0
    var v4 float32 = float32(1.5)
    var v5 uint8 = 127
    v4 = (float32((1.5 + ((0.1 * float64((v2 / v2))) + 0.1))) - ((v4 - (v2 / float32(1.5))) - v4))
    v5 += 63
    var v10 uint = 42
    switch ("\n\t" + "foo bar") {
    case "foo bar":
        for fi11 := 0; fi11 < 13; fi11++ {
            if false {
                for fi12 := 0; fi12 < 20; fi12++ {
                    v5 -= genFn0(9223372036854775807, 2)
                    if false {
                        v10 += 476
                        v13 := []int{}
                        var v14 int16 = 1
                        _ = v13
                        _ = v14
                    } else {
                        if true {
                            v3 = v3
                        } else {
                            v15 := GenS1{F0: 159}
                            v16 := v1[(("a" + ("hello" + ((("hello" + "foo bar") + ("\n\t" + "a")) + "foo bar"))) + "\n\t")]
                            fmt.Println(v3, fi11, fi12)
                            _ = v15
                            _ = v16
                        }
                        if false {
                            if true {
                                if true {
                                    v17 := GenS2{F0: (fi12 & 127)}
                                    v18 := GenS3{F0: ("foo bar" + ""), F1: v3, F2: v3}
                                    v19 := GenS3{F0: "hello", F1: uint32(v10), F2: v3}
                                    _ = v17
                                    _ = v18
                                    _ = v19
                                } else {
                                    var v20 int64 = -9223372036854775808
                                    v2 -= v2
                                    var v21 bool = ((18446744073709551615 != 461) && ((true || (false || (genFn0(1, 0) <= 0))) && (-480 >= 2147483647)))
                                    _ = v20
                                    _ = v21
                                }
                            } else {
                                v22 := []int{}
                                _ = v22
                            }
                        }
                        if false {
                            fmt.Println(v10, v10, v10)
                            if false {
                                if true {
                                    var v23 float64 = 3.14159
                                    if true {
                                        fmt.Println(v23, fi11)
                                    }
                                    _ = v23
                                } else {
                                    var v24 uint = (v10 & (v10 - (v10 >> ((uint(v5)) & 63))))
                                    _ = v24
                                }
                            } else {
                                if false {
                                    v3 += 42
                                    var v25 int16 = 1
                                    _ = v25
                                } else {
                                    var v26 int64 = -48
                                    _ = v26
                                }
                            }
                        }
                        var v27 int16 = int16(fi11)
                        _ = v27
                    }
                    _ = fi12
                }
            }
            v3 = v3
            fmt.Println(v5, v10, v4)
            _ = fi11
        }
    case "":
        switch -1 {
        case 9223372036854775807:
            v5 += (v5 | 255)
        default:
            v28 := GenS1{F0: 1}
            _ = v28
        }
    case "hello":
        v4 = v4
    }
    v29 := v1[(("\n\t" + "a") + "hello")]
    fmt.Println(v10, v3, v5)
    v4 = v4
    var v30 string = ""
    v31 := v1["hello"]
    v4 -= (float32(((1e10 + 1.5) * (0.0 * ((-2.25 + 1e10) + float64((v2 / 1.0)))))) - v2)
    v32 := GenS1{F0: 1}
    v33 := []int{2147483647, 128, 687}
    if (v30 == v30) {
        var v34 int = -2147483648
        _ = v34
    } else {
        var v35 uint32 = (v31 << ((v10) & 63))
        v31 += 116
        v36 := GenS1{F0: 2}
        _ = v35
        _ = v36
    }
    for fi37 := 0; fi37 < 5; fi37++ {
        var v38 int16 = 0
        for fi39 := 0; fi39 < 13; fi39++ {
            v4 += (0.0 + (3.14159 - v2))
            switch "a" {
            case "a":
                if (!false) {
                    var v40 bool = (true || true)
                    v5 += 255
                    _ = v40
                }
            default:
                var v41 int64 = -1
                var v42 float32 = 1e10
                var v43 int32 = 0
                _ = v41
                _ = v42
                _ = v43
            }
            v2 = v4
            _ = fi39
        }
        _ = fi37
        _ = v38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
