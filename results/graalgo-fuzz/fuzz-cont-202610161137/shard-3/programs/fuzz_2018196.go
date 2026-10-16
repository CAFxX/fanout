package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 uint64; F1 int32 }
type GenS5 struct { F0 string; F1 int8 }
type GenS6 struct { F0 uint64; F1 string }
func genFn0(p0 uint16, p1 uint8) uint {
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 int8 = -35
        var v8 int64 = -127
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 1}
    p0 += (^(p0))
    if false {
        p0 = p0
        fmt.Println(p1, p0)
        v10 := GenS2{F0: (1 == 0)}
        _ = v10
    } else {
        if false {
            p1 -= (p1 ^ (p1 ^ 1))
        }
    }
    _ = p0
    _ = p1
    _ = v9
    return 2
}

func main() {
    var v1 int = 0
    var v2 uint32 = 1
    var v3 float64 = 3.14159
    v4 := []int{(v1 / ((-202) | 1)), v1, v1}
    var v5 int16 = 1
    for fi11 := 0; fi11 < 6; fi11++ {
        v5 += 32767
        v12 := GenS2{F0: (1 == 87)}
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi11
        _ = v12
    }
    v13 := GenS2{F0: (42 == 2147483647)}
    switch v1 {
    case 128:
        v14 := GenS3{F0: (!(12 < 135)), F1: ("\n\t" + "a")}
        _ = v14
    case -287:
        var v15 int8 = 1
        _ = v15
    case 80:
        var v16 int16 = (v5 - (v5 % ((v5) | 1)))
        _ = v16
    }
    v17 := GenS4{F0: 1, F1: -151}
    for fi18 := 0; fi18 < 11; fi18++ {
        if (true || (42 >= 42)) {
            v5 = (^(v5))
        }
        _ = fi18
    }
    v1 -= v1
    for fi19 := 0; fi19 < 5; fi19++ {
        v5 += v5
        v1 = 4294967295
        _ = fi19
    }
    fmt.Println(v2)
    v5 -= (v5 + (v5 - (v5 + v5)))
    v4 = append(v4, v1)
    v4 = append(v4, 32767)
    for fi20 := 0; fi20 < 3; fi20++ {
        v1 += fi20
        if false {
            v3 = v3
            if (-1 != (v5 * (v5 - 32767))) {
                switch fi20 {
                case -1:
                    var v21 int8 = -1
                    _ = v21
                case 32768:
                    var v22 int = v1
                    _ = v22
                case 32767:
                    var v23 int64 = 0
                    _ = v23
                default:
                    var v24 uint16 = 32767
                    _ = v24
                }
                var v25 int32 = 42
                if ((42 < -11) || true) {
                    v26 := GenS5{F0: "foo bar", F1: -71}
                    _ = v26
                }
                _ = v25
            } else {
                if len(v4) > 0 {
                    v4[0] = v1
                }
                v1 += 1
                v27 := GenS6{F0: 4294967295, F1: (("foo bar" + ("a" + "a")) + "")}
                _ = v27
            }
        } else {
            var v28 uint = genFn0(32767, 19)
            v2 -= (v2 ^ 91)
            var v29 uint32 = (v2 ^ v2)
            _ = v28
            _ = v29
        }
        _ = fi20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
