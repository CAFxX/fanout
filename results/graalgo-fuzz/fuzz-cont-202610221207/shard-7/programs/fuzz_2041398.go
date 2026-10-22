package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 uint64, p1 uint32, p2 uint64) bool {
    v7 := GenS1{F0: -2147483648}
    fmt.Println(p1, p1)
    p1 = p1
    var v8 int = 9223372036854775807
    v9 := GenS1{F0: v8}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return true
}
func genFn1() string {
    var v10 uint16 = 559
    fmt.Println(v10)
    for fi11 := 0; fi11 < 3; fi11++ {
        v12 := []int{}
        v10 += v10
        _ = fi11
        _ = v12
    }
    v13 := []int{303, -1, 127, 128}
    var v14 bool = false
    _ = v10
    _ = v13
    _ = v14
    return "hello"
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]bool)
    var v3 int16 = -885
    v4 := make(map[uint32]string)
    var v5 int64 = -859
    var v6 float32 = 1.5
    v3 = (-(v3))
    v15 := []int{32768, 4294967295, 142}
    v16 := v1["foo bar"]
    delete(v4, 83)
    var v17 bool = false
    for fi18 := 0; fi18 < 5; fi18++ {
        if ((!v17) && v17) {
            v15 = append(v15, (fi18 + fi18))
            var v19 uint64 = 42
            v17 = v17
            _ = v19
        } else {
            if v17 {
                var v20 int32 = -858
                _ = v20
            } else {
                delete(v4, 65535)
                fmt.Println(v3, fi18)
                v21 := GenS1{F0: 320}
                _ = v21
            }
            var v22 uint16 = 1
            v17 = v17
            _ = v22
        }
        _ = fi18
    }
    v6 = v6
    v5 = (v5 & (v5 ^ (v5 % (((v5 >> ((0) & 63))) | 1))))
    var v23 string = "hello"
    v3 = v3
    if (!(v17 && (!(!(-477 != 32767))))) {
        for fi24 := 0; fi24 < 16; fi24++ {
            if true {
                var v25 float32 = (1.5 * ((0.1 / 1e10) / ((1.5 + (((0.1 * 1e10) * ((0.1 - float32(0.1)) - (v6 - 3.14159))) / v6)) * 0.0)))
                _ = v25
            }
            _ = fi24
        }
        var v26 int8 = 5
        _ = v26
    } else {
        var v27 int64 = v5
        _ = v27
    }
    if false {
        v16 = ((v23 + "hello") + (v16 + (v16 + "\n\t")))
        v5 += 0
    } else {
        var v28 float64 = 0.0
        _ = v28
    }
    v6 = (0.1 * float32(3.14159))
    switch 128 {
    case -994:
        var v29 uint32 = 42
        _ = v29
    case 9223372036854775807:
        var v30 int = 42
        _ = v30
    case -2147483648:
        fmt.Println(v16, v17, v6)
    default:
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
    }
    for fi31 := 0; fi31 < 2; fi31++ {
        var v32 float64 = (1e10 + 3.14159)
        switch fi31 {
        case -2147483648:
            v16 = ""
        case 42:
            v6 -= float32((((v32 / v32) / (1.5 / float64(v6))) - v32))
        case 32768:
            for fi33 := 0; fi33 < 9; fi33++ {
                var v34 float64 = (v32 * 0.1)
                v32 = ((-2.25 * 1e10) * 1.5)
                if v17 {
                    v35 := GenS1{F0: (fi31 << ((4294967295) & 63))}
                    _ = v35
                } else {
                    if len(v15) > 0 {
                        v15[0] = (fi31 & fi31)
                    }
                }
                switch fi33 {
                case 0:
                    fmt.Println(fi33, v5, v3)
                case -9223372036854775808:
                    v5 += (^(v5))
                default:
                    v16 = v16
                }
                _ = fi33
                _ = v34
            }
        }
        fmt.Println(v3, v5, v16)
        _ = fi31
        _ = v32
    }
    v17 = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
