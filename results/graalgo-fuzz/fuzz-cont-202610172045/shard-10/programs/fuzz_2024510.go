package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 int; F2 uint8 }
func genFn0(p0 string) int {
    for fi5 := 0; fi5 < 10; fi5++ {
        p0 = p0
        _ = fi5
    }
    var v6 uint16 = 170
    p0 = p0
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 int = (fi7 ^ fi7)
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = v6
    return 2147483647
}
func genFn1(p0 bool, p1 int8) bool {
    v9 := []int{4294967295, 246, 127}
    fmt.Println(p0, p1)
    var v10 uint32 = 269
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return p0
}

func main() {
    v1 := []string{"hello", ("\n\t" + "a"), "a"}
    var v2 bool = false
    v3 := []uint32{}
    v4 := []string{"", ("hello" + ("\n\t" + "a")), ""}
    switch 32767 {
    case -679:
        var v11 bool = false
        _ = v11
    case 165:
        v12 := GenS1{F0: 65535}
        _ = v12
    case 128:
        v2 = (!v2)
    default:
        if (2147483647 <= 42) {
            v2 = (32768 <= -2147483648)
            switch 128 {
            case 2147483647:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case -457:
                for fi13 := 0; fi13 < 4; fi13++ {
                    v14 := GenS2{F0: 538, F1: 42, F2: 127}
                    _ = fi13
                    _ = v14
                }
            default:
                v2 = (v2 || (!true))
            }
        }
    }
    var v15 float64 = (0.1 * ((3.14159 * (((((1.5 - ((((1.5 + (1.5 - 0.0)) + 3.14159) - ((0.0 + 0.1) - -2.25)) + (float64(float32(-2.25)) * 0.1))) / 1.5) * 0.0) + 0.0) * (float64((float32(-2.25) + 0.1)) + 3.14159))) / (1e10 - 0.0)))
    for fi16 := 0; fi16 < 9; fi16++ {
        v17 := GenS2{F0: 26, F1: fi16, F2: 0}
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 1; fi18++ {
        v2 = (729 <= 65535)
        for fi19 := 0; fi19 < 6; fi19++ {
            var v20 uint = 84
            if v2 {
                v2 = true
                v15 = 0.1
                v21 := GenS2{F0: 4294967295, F1: fi19, F2: 144}
                _ = v21
            } else {
                if v2 {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                } else {
                    v15 += 1.5
                }
                if false {
                    v2 = true
                    v22 := GenS2{F0: 141, F1: fi19, F2: 255}
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v22
                }
                var v23 uint8 = 42
                _ = v23
            }
            _ = fi19
            _ = v20
        }
        for fi24 := 0; fi24 < 10; fi24++ {
            v2 = v2
            _ = fi24
        }
        _ = fi18
    }
    var v25 int = 2147483647
    if false {
        v25 -= v25
    }
    if false {
        var v26 int = (-(v25))
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        fmt.Println(v2, v15, v25)
        _ = v26
    }
    fmt.Println(v2, v25)
    v15 = v15
    fmt.Println(v15)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v25
    _ = v3
    _ = v4
}
