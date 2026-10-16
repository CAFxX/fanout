package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 uint32 }
type GenS2 struct { F0 int32 }
func genFn0(p0 int32, p1 uint64) int16 {
    for fi7 := 0; fi7 < 16; fi7++ {
        v8 := GenS1{F0: (p1 >> ((42) & 63)), F1: (^(p1)), F2: 2147483647}
        _ = fi7
        _ = v8
    }
    p0 -= 1
    var v9 int8 = -1
    for fi10 := 0; fi10 < 16; fi10++ {
        v11 := GenS1{F0: p1, F1: p1, F2: 67}
        var v12 float32 = -2.25
        var v13 int8 = (v9 * (v9 - -128))
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v9
    return -245
}
func genFn1() int32 {
    switch "hello" {
    case "":
        var v14 int8 = 42
        _ = v14
    default:
        v15 := []int{128}
        _ = v15
    }
    v16 := []int{-694, 4294967295}
    if (32767 <= 400) {
        switch -2147483648 {
        case 2147483647:
            if len(v16) > 0 {
                v16[0] = -1
            }
        default:
            if (!(42 < 96)) {
                var v17 uint16 = 72
                v16 = append(v16, 32768)
                for fi18 := 0; fi18 < 10; fi18++ {
                    for fi19 := 0; fi19 < 2; fi19++ {
                        if false {
                            if len(v16) >= 2 {
                                v16 = v16[1:len(v16)]
                            }
                            v17 += (v17 << ((42) & 63))
                        }
                        fmt.Println(v17, fi19)
                        var v20 int64 = -1
                        _ = fi19
                        _ = v20
                    }
                    var v21 float64 = 1.5
                    for fi22 := 0; fi22 < 14; fi22++ {
                        v21 = (float64(1.5) - ((((v21 - (1.5 * (float64(0.1) * 1e10))) - (3.14159 / ((1.5 * 1e10) * ((float64(1e10) / (v21 / (-2.25 + v21))) / float64(3.14159))))) + (float64(0.1) * 0.1)) - 0.0))
                        _ = fi22
                    }
                    _ = fi18
                    _ = v21
                }
                _ = v17
            }
        }
        var v23 int16 = genFn0(-666, 82)
        _ = v23
    } else {
        var v24 float64 = 0.1
        v24 = float64(1e10)
        _ = v24
    }
    _ = v16
    return -1
}

func main() {
    v1 := []string{(("hello" + "") + (("foo bar" + "foo bar") + "")), (("" + (("" + "a") + "hello")) + "hello"), ((("\n\t" + "hello") + ("a" + "foo bar")) + ("a" + "a"))}
    var v2 uint64 = 0
    var v3 uint16 = 2
    var v4 int8 = 42
    var v5 float32 = 3.14159
    var v6 int32 = -2147483648
    v25 := GenS2{F0: v6}
    if false {
        if (!((1.5 - v5) < v5)) {
            v5 += v5
            for fi26 := 0; fi26 < 4; fi26++ {
                var v27 uint8 = 255
                for fi28 := 0; fi28 < 8; fi28++ {
                    v4 -= v4
                    _ = fi28
                }
                _ = fi26
                _ = v27
            }
            v3 += 2
        }
    } else {
        v4 -= v4
        var v29 uint64 = (v2 & v2)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v29
    }
    for fi30 := 0; fi30 < 1; fi30++ {
        v1 = append(v1, "foo bar")
        _ = fi30
    }
    v5 -= ((v5 * float32((float64(v5) * 0.0))) / v5)
    fmt.Println(v4, v2, v2)
    var v31 uint32 = 2
    v3 = 975
    var v32 uint32 = 2147483648
    for fi33 := 0; fi33 < 2; fi33++ {
        var v34 uint32 = v31
        _ = fi33
        _ = v34
    }
    v35 := GenS2{F0: genFn1()}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v31
    _ = v32
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
