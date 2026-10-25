package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 uint16 }
func genFn0(p0 uint, p1 uint16, p2 int8) uint64 {
    var v7 float32 = 0.1
    p0 = 119
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 4294967295
}

func main() {
    var v1 int8 = 42
    var v2 string = "foo bar"
    v3 := []int{127, -240, 9223372036854775807, -2147483648}
    var v4 uint = 892
    var v5 float32 = 1.5
    var v6 uint16 = 42
    for fi8 := 0; fi8 < 1; fi8++ {
        for fi9 := 0; fi9 < 3; fi9++ {
            if len(v3) > 0 {
                v3[0] = 2147483647
            }
            v1 = v1
            v10 := GenS1{F0: v4}
            _ = fi9
            _ = v10
        }
        fmt.Println(v5, fi8)
        if ((!((v6 % (((v6 << ((v4) & 63))) | 1)) > (v6 | v6))) && (true || ((65535 <= 587) && true))) {
            _ = len(v3)
        } else {
            var v11 float32 = float32((0.1 * -2.25))
            _ = v11
        }
        _ = fi8
    }
    switch v2 {
    case "a":
        v12 := GenS1{F0: (v4 ^ (v4 / (((v4 + (v4 + 65535))) | 1)))}
        _ = v12
    }
    if false {
        v4 -= v4
        switch -595 {
        case 128:
            var v13 float64 = 0.1
            _ = v13
        default:
            if true {
                switch 1 {
                case 622:
                    v6 -= 32767
                case 42:
                    var v14 uint8 = 117
                    _ = v14
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v1 += -1
            }
        }
        v5 -= v5
    } else {
        _ = len(v3)
    }
    var v15 int16 = 1
    v3 = append(v3, int(v1))
    var v16 float32 = v5
    v3 = append(v3, 128)
    v1 = (v1 >> ((42) & 63))
    if false {
        if (true && true) {
            fmt.Println(v2, v4, v6)
            var v17 string = ""
            var v18 string = v17
            _ = v17
            _ = v18
        } else {
            if len(v3) > 0 {
                v3[0] = 42
            }
        }
    } else {
        v4 -= (v4 - v4)
    }
    var v19 string = ""
    fmt.Println(v19, v2, v15)
    if true {
        for fi20 := 0; fi20 < 14; fi20++ {
            v1 -= (-(v1))
            var v21 uint = 42
            for fi22 := 0; fi22 < 14; fi22++ {
                if false {
                    var v23 int16 = -1
                    fmt.Println(v4, v16)
                    switch 398 {
                    case -785:
                        v24 := GenS2{F0: 1, F1: 42}
                        _ = v24
                    case 4294967295:
                        var v25 int = fi22
                        _ = v25
                    case 42:
                        var v26 int16 = v15
                        _ = v26
                    }
                    _ = v23
                }
                var v27 float64 = (0.1 - 0.1)
                _ = fi22
                _ = v27
            }
            _ = fi20
            _ = v21
        }
    }
    if true {
        v16 += 0.0
        if (!((((0.0 - (1e10 - 0.1)) - 0.0) + ((-2.25 + 1e10) * (float64(float32((0.1 - 1.5))) - (((0.1 - float64(3.14159)) - 1.5) / 3.14159)))) > (((((((float64((float32((float64((v5 * 1e10)) * -2.25)) - v16)) - 0.1) / float64((float32(3.14159) * v16))) / (0.1 * 1.5)) / 1.5) - -2.25) * -2.25) - -2.25))) {
            v5 = v16
        } else {
            if len(v3) > 0 {
                v3[0] = 32768
            }
            v6 += v6
        }
    } else {
        v19 = (("" + "foo bar") + (v2 + ("\n\t" + (v19 + ("a" + ((("hello" + v19) + v2) + ((((((v19 + ("foo bar" + v2)) + "\n\t") + v19) + (((v19 + ("foo bar" + v2)) + ("hello" + v19)) + v19)) + "\n\t") + (("" + "hello") + v19))))))))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
