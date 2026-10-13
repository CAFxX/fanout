package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64 }
func genFn0(p0 uint16, p1 uint8, p2 uint8) int16 {
    var v7 uint8 = p1
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1000
}
func genFn1(p0 uint, p1 uint64, p2 string) int16 {
    switch -387 {
    case 9223372036854775807:
        var v8 uint8 = 42
        _ = v8
    default:
        p1 = p1
    }
    v9 := []int{32767, -509, -1}
    for fi10 := 0; fi10 < 13; fi10++ {
        _ = len(v9)
        var v11 float32 = (0.1 / (1.5 - 0.0))
        var v12 int64 = 9223372036854775807
        _ = fi10
        _ = v11
        _ = v12
    }
    p1 = 951
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return -32768
}

func main() {
    v1 := make(map[string]string)
    var v2 float64 = 0.0
    var v3 float32 = 1e10
    v4 := make(map[string]bool)
    var v5 bool = (true || (65535 >= 424))
    var v6 uint64 = 42
    var v13 int = -9223372036854775808
    fmt.Println(v6, v3, v13)
    v1["\n\t"] = "\n\t"
    var v14 uint8 = 120
    fmt.Println(v2)
    v15 := GenS1{F0: 127}
    v13 = (v13 + v13)
    v2 += v2
    var v16 int16 = -730
    var v17 string = ("\n\t" + "hello")
    switch ("a" + v17) {
    case "hello":
        switch "\n\t" {
        case "foo bar":
            v18 := []int{}
            _ = v18
        default:
            switch v17 {
            case "a":
                v14 += (v14 % ((v14) | 1))
            default:
                v13 -= v13
            }
        }
    }
    if false {
        if ((v3 * (float32(1e10) * (v3 - (v3 / v3)))) < v3) {
            var v19 float64 = v2
            switch v13 {
            case 103:
                v20 := []int{v13, -152}
                _ = v20
            }
            var v21 uint8 = 127
            _ = v19
            _ = v21
        } else {
            var v22 uint16 = 65535
            var v23 bool = (v5 || (true && v5))
            switch v17 {
            case "hello":
                fmt.Println(v14)
            default:
                if v23 {
                    if (v3 < v3) {
                        v24 := []int{-795, (v13 | 42), -1, int(v22)}
                        _ = len(v1)
                        for fi25 := 0; fi25 < 17; fi25++ {
                            if v5 {
                                var v26 uint = 0
                                v24 = append(v24, 32767)
                                for fi27 := 0; fi27 < 4; fi27++ {
                                    fmt.Println(v6)
                                    if len(v24) > 0 {
                                        v24[0] = (-(fi27))
                                    }
                                    _ = fi27
                                }
                                _ = v26
                            }
                            fmt.Println(v22)
                            _ = fi25
                        }
                        _ = v24
                    }
                    v28 := GenS2{F0: 1}
                    _ = v28
                } else {
                    v29 := GenS1{F0: 42}
                    if true {
                        v30 := GenS1{F0: -1}
                        switch 374 {
                        case 4294967295:
                            v13 += (v13 & v13)
                        case 956:
                            _ = len(v1)
                        case -1:
                            delete(v1, (v17 + "foo bar"))
                        default:
                            v31 := GenS2{F0: -271}
                            _ = v31
                        }
                        var v32 float32 = v3
                        _ = v30
                        _ = v32
                    } else {
                        if (v2 == ((-2.25 + (v2 / 3.14159)) / v2)) {
                            fmt.Println(v5)
                            v16 = (^(v16))
                        }
                    }
                    _ = v29
                }
            }
            _ = v22
            _ = v23
        }
        fmt.Println(v2, v17)
        v16 = v16
    }
    switch v13 {
    case 32767:
        v33 := []int{-2147483648, v13, -2147483648, (v13 ^ (v13 + (^(v13))))}
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
