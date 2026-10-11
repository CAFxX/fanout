package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64; F1 uint32 }
type GenS4 struct { F0 int64 }
func genFn0(p0 bool, p1 int16) uint8 {
    p0 = ((("a" <= "a") && (185 < 128)) && p0)
    switch "foo bar" {
    case "a":
        switch 464 {
        case -873:
            v7 := GenS1{F0: ("\n\t" + (("a" + ("hello" + "")) + ("\n\t" + ("foo bar" + "\n\t"))))}
            _ = v7
        default:
            var v8 int64 = 42
            _ = v8
        }
    case "\n\t":
        p0 = (0.0 >= ((1e10 + 1e10) - 0.0))
    }
    v9 := []int{}
    _ = p0
    _ = p1
    _ = v9
    return 42
}

func main() {
    v1 := []uint32{0, 42, 2147483647}
    var v2 int64 = -1
    v3 := make(map[string]bool)
    var v4 float32 = (0.0 / 1.0)
    var v5 uint64 = 436
    var v6 int16 = 32767
    var v10 uint = 1
    for fi11 := 0; fi11 < 14; fi11++ {
        var v12 int = fi11
        _ = fi11
        _ = v12
    }
    var v13 string = ("hello" + (("foo bar" + "hello") + "a"))
    var v14 float64 = 1e10
    v15 := GenS1{F0: (v13 + ((("" + "foo bar") + "hello") + "foo bar"))}
    if false {
        for fi16 := 0; fi16 < 16; fi16++ {
            v17 := GenS2{F0: fi16}
            v10 += v10
            for fi18 := 0; fi18 < 9; fi18++ {
                var v19 float64 = v14
                if (v4 != float32(v19)) {
                    v1 = append(v1, 0)
                    fmt.Println(fi16, v14, v14)
                }
                _ = fi18
                _ = v19
            }
            _ = fi16
            _ = v17
        }
        v20 := GenS2{F0: 0}
        if false {
            var v21 float64 = ((1e10 * (v14 - float64(v4))) / v14)
            for fi22 := 0; fi22 < 7; fi22++ {
                v23 := GenS3{F0: (v2 * 567), F1: 271}
                v1 = append(v1, 2147483648)
                v3[("foo bar" + v13)] = (false || false)
                _ = fi22
                _ = v23
            }
            if false {
                v5 += v5
                v13 = ("a" + v13)
            } else {
                if (!false) {
                    var v24 bool = (!(95 < -91))
                    fmt.Println(v2, v14)
                    var v25 string = v13
                    _ = v24
                    _ = v25
                } else {
                    v4 -= (v4 * 1e10)
                }
                v1 = append(v1, 503)
                v13 = "foo bar"
            }
            _ = v21
        }
        _ = v20
    }
    switch 9223372036854775807 {
    case 373:
        fmt.Println(v4)
    case 0:
        fmt.Println(v5, v5, v10)
    default:
        switch int(v6) {
        case 2147483647:
            v26 := GenS2{F0: -180}
            _ = v26
        case 0:
            var v27 int8 = -1
            _ = v27
        default:
            var v28 int32 = -680
            _ = v28
        }
    }
    v29 := GenS4{F0: (-(v2))}
    fmt.Println(v5, v13)
    fmt.Println(v14, v13, v14)
    var v30 uint16 = 32767
    v3[v13] = false
    for fi31 := 0; fi31 < 17; fi31++ {
        v30 -= (v30 + 870)
        _ = fi31
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
