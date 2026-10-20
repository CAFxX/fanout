package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int8 }
type GenS2 struct { F0 uint64 }
func genFn0() int64 {
    if true {
        var v7 string = ""
        _ = v7
    } else {
        var v8 uint32 = 2147483647
        _ = v8
    }
    var v9 uint64 = 0
    v10 := []int{32768, 32767, 32768, 50}
    _ = v10
    _ = v9
    return 0
}
func genFn1(p0 int64, p1 uint8) bool {
    fmt.Println(p0)
    v11 := GenS1{F0: (0 < 0), F1: 4294967295, F2: -1}
    if true {
        var v12 float32 = 0.0
        _ = v12
    }
    v13 := []int{128, -1, 4294967295, 1}
    var v14 uint = 4294967295
    _ = p0
    _ = p1
    _ = v11
    _ = v13
    _ = v14
    return false
}

func main() {
    var v1 int = 4294967295
    v2 := []int{414, (v1 >> ((18446744073709551615) & 63)), (v1 >> ((4294967295) & 63)), v1}
    v3 := []string{"", "hello"}
    var v4 int16 = 789
    var v5 int16 = (v4 << ((2) & 63))
    v6 := make(map[string]bool)
    var v15 int32 = -1
    v1 = v1
    delete(v6, ("hello" + ""))
    fmt.Println(v1)
    var v16 uint8 = 0
    v15 = (-(v15))
    v1 = -733
    v5 = v5
    switch 9223372036854775807 {
    case -174:
        var v17 int = v1
        _ = v17
    case 73:
        switch v1 {
        case 2147483647:
            v16 -= v16
        case 42:
            if true {
                if ("\n\t" > "hello") {
                    var v18 float32 = -2.25
                    v19 := GenS1{F0: (0 > 42), F1: 65535, F2: 127}
                    fmt.Println(v5, v4)
                    _ = v18
                    _ = v19
                } else {
                    fmt.Println(v5, v5)
                    if (1.5 == 0.1) {
                        v1 = v1
                        for fi20 := 0; fi20 < 17; fi20++ {
                            v16 -= v16
                            _ = fi20
                        }
                        fmt.Println(v5)
                    }
                }
            } else {
                v16 -= 128
                if (2147483648 >= 1) {
                    var v21 uint64 = 4294967295
                    v21 += (v21 / ((v21) | 1))
                    v6[(("\n\t" + "") + (("\n\t" + "foo bar") + "hello"))] = (!(false || ((float64((1.5 / (3.14159 + 0.1))) * 0.1) == -2.25)))
                    _ = v21
                }
            }
        case 32767:
            if true {
                v22 := GenS1{F0: false, F1: 0, F2: 1}
                var v23 uint = 290
                v5 += (v5 / ((442) | 1))
                _ = v22
                _ = v23
            } else {
                v16 = (v16 & v16)
                if true {
                    for fi24 := 0; fi24 < 6; fi24++ {
                        fmt.Println(v4, v15, v16)
                        _ = fi24
                    }
                    for fi25 := 0; fi25 < 17; fi25++ {
                        v16 = v16
                        fmt.Println(fi25, v1, v4)
                        var v26 float64 = 1.5
                        _ = fi25
                        _ = v26
                    }
                } else {
                    v4 += v4
                    v27 := GenS1{F0: (v16 <= (v16 | 255)), F1: 255, F2: 42}
                    _ = v27
                }
            }
        default:
            v28 := GenS1{F0: (!true), F1: 743, F2: 100}
            _ = v28
        }
    default:
        var v29 uint = 334
        _ = v29
    }
    v30 := GenS2{F0: 93}
    var v31 uint8 = (v16 >> ((255) & 63))
    for fi32 := 0; fi32 < 13; fi32++ {
        var v33 int32 = v15
        _ = fi32
        _ = v33
    }
    v16 += v31
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
