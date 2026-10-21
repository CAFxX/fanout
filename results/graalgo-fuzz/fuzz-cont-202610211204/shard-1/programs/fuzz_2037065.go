package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
func genFn0() string {
    var v7 uint32 = 65535
    var v8 string = ""
    _ = v7
    _ = v8
    return v8
}
func genFn1(p0 bool, p1 int64, p2 uint64) uint16 {
    p2 = p2
    v9 := []int{32767, 847, 32767, -689}
    v10 := GenS1{F0: 32767, F1: -1}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 549
}

func main() {
    v1 := make(map[int]bool)
    var v2 int8 = -23
    var v3 string = "foo bar"
    var v4 string = v3
    var v5 uint8 = 128
    var v6 int8 = v2
    switch 127 {
    case -1:
        v3 = "hello"
    case 32767:
        var v11 int8 = (v6 - v6)
        _ = v11
    default:
        fmt.Println(v6, v5, v2)
    }
    v5 -= v5
    var v12 int = 32768
    v1[(v12 / ((v12) | 1))] = true
    var v13 float64 = (1e10 + -2.25)
    v4 = v4
    if (v5 <= v5) {
        var v14 int = 0
        _ = v14
    } else {
        var v15 string = ("" + (("a" + v4) + ""))
        for fi16 := 0; fi16 < 15; fi16++ {
            v6 = (v2 % ((v6) | 1))
            var v17 string = "foo bar"
            v18 := []int{(v12 << ((4294967295) & 63)), (v12 % ((fi16) | 1)), (fi16 + 32767), (^(fi16))}
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = v15
    }
    for fi19 := 0; fi19 < 7; fi19++ {
        if true {
            v12 = -1
            for fi20 := 0; fi20 < 11; fi20++ {
                switch fi19 {
                case 128:
                    var v21 uint16 = 42
                    _ = v21
                case 4294967295:
                    var v22 uint = 858
                    _ = v22
                default:
                    v23 := []int{fi19, fi20}
                    _ = v23
                }
                _ = fi20
            }
        } else {
            for fi24 := 0; fi24 < 14; fi24++ {
                if (v13 != (1.5 * 3.14159)) {
                    var v25 int16 = 1000
                    if (!(!("foo bar" > v3))) {
                        v4 = ((v3 + ("a" + "foo bar")) + "a")
                        v26 := v1[128]
                        _ = v26
                    } else {
                        if false {
                            delete(v1, (-(v12)))
                            v27 := []int{}
                            _ = v27
                        }
                        if ((v4 == ("foo bar" + "\n\t")) && false) {
                            fmt.Println(v13, v12, v4)
                        }
                    }
                    _ = v25
                }
                _ = fi24
            }
            fmt.Println(v5, v2, v4)
            var v28 bool = true
            _ = v28
        }
        v29 := GenS1{F0: (^(v12)), F1: 1000}
        _ = fi19
        _ = v29
    }
    v5 = (^(v5))
    switch v4 {
    case "":
        for fi30 := 0; fi30 < 16; fi30++ {
            v1[(v12 * v12)] = true
            var v31 float32 = -2.25
            v13 -= 0.0
            _ = fi30
            _ = v31
        }
    default:
        var v32 float32 = ((0.1 * 3.14159) - (-2.25 / 1e10))
        _ = v32
    }
    var v33 bool = ((1 >= 4294967295) || (((v3 + "foo bar") <= v3) || true))
    v3 = ("" + ((v4 + ("" + ("" + ("foo bar" + v3)))) + "a"))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
