package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 string; F2 int64 }
type GenS3 struct { F0 int16 }
func genFn0() uint16 {
    var v7 uint32 = 1
    if true {
        var v8 int32 = -1
        if (true && (v7 <= v7)) {
            if ((-2.25 <= -2.25) && (3.14159 >= float64((3.14159 * 1.5)))) {
                v9 := []int{}
                _ = len(v9)
                _ = v9
            }
        } else {
            v10 := []int{776, 4294967295}
            _ = v10
        }
        switch "hello" {
        case "\n\t":
            var v11 float64 = 3.14159
            _ = v11
        case "":
            switch 32767 {
            case -692:
                if true {
                    if true {
                        switch 2147483647 {
                        case -23:
                            var v12 uint32 = 2147483648
                            _ = v12
                        case 32768:
                            if true {
                                if false {
                                    v13 := []int{-9223372036854775808, 1, 9223372036854775807}
                                    v8 = 517
                                    _ = v13
                                }
                                fmt.Println(v7, v8)
                                for fi14 := 0; fi14 < 12; fi14++ {
                                    v15 := GenS1{F0: "hello"}
                                    v16 := []int{fi14}
                                    v17 := GenS1{F0: "a"}
                                    _ = fi14
                                    _ = v15
                                    _ = v16
                                    _ = v17
                                }
                            }
                        case -1:
                            for fi18 := 0; fi18 < 1; fi18++ {
                                v19 := GenS2{F0: (v7 << ((42) & 63)), F1: "foo bar", F2: 42}
                                _ = fi18
                                _ = v19
                            }
                        default:
                            v20 := GenS3{F0: 0}
                            _ = v20
                        }
                        var v21 float64 = (((1.5 / 0.1) + 1.5) - 1.5)
                        _ = v21
                    }
                } else {
                    v22 := []int{2147483647, -793, -9223372036854775808}
                    if false {
                        if true {
                            v7 -= v7
                        }
                        if false {
                            switch -229 {
                            case 2147483647:
                                var v23 uint8 = 15
                                _ = v23
                            case -552:
                                v8 -= v8
                            default:
                                v8 = (v8 / (((v8 >> ((42) & 63))) | 1))
                            }
                            v8 -= -838
                        }
                    }
                    var v24 bool = ((100 < 234) || false)
                    _ = v22
                    _ = v24
                }
            default:
                v25 := GenS2{F0: (v7 << ((72) & 63)), F1: "", F2: -1}
                _ = v25
            }
        }
        _ = v8
    }
    v26 := []int{-811}
    for fi27 := 0; fi27 < 20; fi27++ {
        v7 = v7
        _ = len(v26)
        _ = fi27
    }
    v26 = append(v26, 1)
    _ = v26
    _ = v7
    return 753
}

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{1}
    v3 := make(map[string]bool)
    var v4 float64 = 1.5
    v5 := make(map[int]bool)
    var v6 int16 = -32768
    switch "foo bar" {
    case "a":
        v28 := GenS2{F0: 315, F1: "\n\t", F2: 42}
        _ = v28
    default:
        fmt.Println(v4)
    }
    v2 = append(v2, 2147483647)
    delete(v5, 9223372036854775807)
    if true {
        v6 = (^(v6))
        var v29 int16 = (v6 | v6)
        switch 307 {
        case 4294967295:
            v6 = v29
        case 128:
            v29 += v6
        }
        _ = v29
    }
    v6 += v6
    v6 += 1000
    v4 += (float64(1e10) / -2.25)
    v30 := GenS3{F0: 42}
    v4 -= 1.5
    var v31 float32 = ((-2.25 + (((float32(-2.25) * 3.14159) - 1e10) - 3.14159)) - 0.1)
    var v32 int16 = (^(v6))
    var v33 uint8 = 61
    var v34 float32 = (v31 - (float32(v4) * (v31 - 3.14159)))
    var v35 int = -30
    v31 -= 1.5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
