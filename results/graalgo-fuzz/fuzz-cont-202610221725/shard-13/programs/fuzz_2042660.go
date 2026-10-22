package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0() uint {
    var v4 int32 = 42
    v5 := []int{-1}
    fmt.Println(v4)
    var v6 int64 = -799
    v5 = append(v5, int(v6))
    _ = v4
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 int16, p1 int8, p2 int8) uint {
    if false {
        var v7 float32 = ((float32((((1e10 * 3.14159) - 0.1) - 3.14159)) * 1e10) / 1.0)
        _ = v7
    }
    v8 := []int{0, 0}
    v9 := GenS1{F0: 834}
    var v10 uint8 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return 827
}

func main() {
    var v1 int32 = 79
    v2 := []string{"\n\t", "hello", ("\n\t" + "\n\t")}
    var v3 uint16 = 255
    if (false || (3.14159 < (1.5 + (((3.14159 + 1e10) - (float64(1.5) + ((3.14159 - (float64((1.5 / 1.5)) + 0.0)) + (1.5 + (1.5 + (0.0 - 1e10)))))) / (((((float64(1.5) / float64(3.14159)) + 1.5) * 1e10) * float64((0.0 + 1e10))) + ((-2.25 + -2.25) * -2.25)))))) {
        if ((!false) && ((-2.25 - (float64(1.5) * 0.1)) > (float64(-2.25) + ((3.14159 / -2.25) / 1.0)))) {
            v3 = (v3 >> ((0) & 63))
            fmt.Println(v1, v1)
            for fi11 := 0; fi11 < 6; fi11++ {
                for fi12 := 0; fi12 < 18; fi12++ {
                    v3 += (v3 / (((v3 | (v3 / ((983) | 1)))) | 1))
                    v3 = 42
                    _ = fi12
                }
                _ = fi11
            }
        }
    } else {
        if len(v2) > 0 {
            v2[0] = ""
        }
    }
    switch -731 {
    case 4294967295:
        var v13 uint32 = 4294967295
        _ = v13
    }
    for fi14 := 0; fi14 < 17; fi14++ {
        _ = len(v2)
        _ = fi14
    }
    if false {
        fmt.Println(v3, v1)
    }
    var v15 bool = false
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    var v16 int = -1
    var v17 bool = v15
    for fi18 := 0; fi18 < 10; fi18++ {
        v19 := GenS1{F0: 4294967295}
        fmt.Println(v1, v16)
        v17 = true
        _ = fi18
        _ = v19
    }
    switch ("a" + "foo bar") {
    case "foo bar":
        v16 = -350
    default:
        switch 0 {
        case -603:
            for fi20 := 0; fi20 < 8; fi20++ {
                var v21 int = (v16 << ((255) & 63))
                switch ("hello" + "a") {
                case "foo bar":
                    fmt.Println(v21)
                default:
                    fmt.Println(fi20, v3)
                }
                fmt.Println(v16, v1, v3)
                _ = fi20
                _ = v21
            }
        case -330:
            for fi22 := 0; fi22 < 12; fi22++ {
                for fi23 := 0; fi23 < 7; fi23++ {
                    fmt.Println(v17, fi23)
                    v17 = (fi23 <= fi23)
                    _ = fi23
                }
                _ = fi22
            }
        default:
            v1 = (v1 | (v1 >> ((0) & 63)))
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
