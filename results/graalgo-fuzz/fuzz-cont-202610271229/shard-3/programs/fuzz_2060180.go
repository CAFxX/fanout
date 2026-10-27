package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
func genFn0(p0 int8) uint64 {
    fmt.Println(p0)
    fmt.Println(p0)
    for fi7 := 0; fi7 < 18; fi7++ {
        for fi8 := 0; fi8 < 12; fi8++ {
            p0 = (p0 - (p0 << ((255) & 63)))
            if (1 < 4294967295) {
                if false {
                    var v9 float64 = 1.5
                    if false {
                        var v10 uint32 = 165
                        _ = v10
                    } else {
                        v11 := []int{507, fi7}
                        _ = v11
                    }
                    _ = v9
                }
                p0 = p0
                var v12 int8 = p0
                var v13 bool = (!(!true))
                v12 -= v12
                var v14 uint16 = 42
                _ = v12
                _ = v13
                _ = v14
            }
            _ = fi8
        }
        var v15 int64 = 9223372036854775807
        var v16 int8 = (p0 * 127)
        _ = fi7
        _ = v15
        _ = v16
    }
    v17 := GenS1{F0: "\n\t", F1: 1}
    _ = p0
    _ = v17
    return 0
}

func main() {
    var v1 uint16 = 65535
    v2 := []string{"foo bar"}
    var v3 float64 = 0.0
    var v4 uint = 968
    v5 := make(map[int]string)
    var v6 float64 = v3
    v2 = append(v2, "foo bar")
    v1 -= v1
    var v18 uint8 = 128
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    switch -738 {
    case 1:
        v1 += uint16(v18)
    case 2147483647:
        for fi19 := 0; fi19 < 2; fi19++ {
            v4 -= (v4 - (v4 ^ uint(v1)))
            v1 -= v1
            var v20 int32 = -254
            _ = fi19
            _ = v20
        }
    }
    var v21 float64 = ((0.1 * (0.1 + v6)) + (v3 * float64((float32(0.0) / (float32(((float64(0.1) - v6) + v6)) * 0.1)))))
    v1 -= (v1 + 2)
    v18 += 0
    if true {
        switch "a" {
        case "hello":
            switch 127 {
            case -101:
                v22 := GenS1{F0: "a", F1: -220}
                _ = v22
            default:
                var v23 int16 = -1
                _ = v23
            }
        default:
            v2 = append(v2, ("hello" + ("a" + ("hello" + "hello"))))
        }
        v4 = v4
    }
    for fi24 := 0; fi24 < 6; fi24++ {
        for fi25 := 0; fi25 < 7; fi25++ {
            v21 = 1e10
            v21 += -2.25
            _ = fi25
        }
        for fi26 := 0; fi26 < 6; fi26++ {
            if ((139 != v18) && (false && ((23 >= 100) && true))) {
                var v27 int8 = 81
                _ = len(v2)
                _ = v27
            }
            var v28 int64 = 599
            _ = fi26
            _ = v28
        }
        _ = fi24
    }
    v29 := v5[-9223372036854775808]
    v30 := v5[-1]
    v31 := GenS1{F0: (v30 + (v29 + v30)), F1: 1000}
    v32 := GenS1{F0: "a", F1: 0}
    var v33 string = (("a" + "hello") + v29)
    if true {
        v3 = 0.1
    } else {
        _ = len(v5)
        var v34 uint64 = 42
        _ = v34
    }
    var v35 uint16 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
