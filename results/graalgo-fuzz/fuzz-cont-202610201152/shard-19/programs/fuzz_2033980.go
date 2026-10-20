package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 int8; F1 int16 }
func genFn0(p0 uint64) string {
    var v5 int64 = 1
    for fi6 := 0; fi6 < 17; fi6++ {
        v7 := GenS1{F0: 173}
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = v5
    return ("" + "a")
}
func genFn1(p0 int64, p1 uint32, p2 uint32) string {
    p2 = 2147483647
    var v8 float32 = 3.14159
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return "\n\t"
}

func main() {
    var v1 int64 = 981
    var v2 string = ("a" + ("foo bar" + "hello"))
    v3 := make(map[int]bool)
    var v4 int16 = 42
    for fi9 := 0; fi9 < 6; fi9++ {
        fmt.Println(v4, v2)
        _ = fi9
    }
    v3[-9223372036854775808] = true
    if (!((v4 != v4) && (false || false))) {
        if true {
            var v10 int16 = (v4 + (-(v4)))
            _ = v10
        }
        for fi11 := 0; fi11 < 11; fi11++ {
            var v12 int16 = v4
            fmt.Println(v1, v12, v1)
            v13 := []int{}
            _ = fi11
            _ = v12
            _ = v13
        }
    } else {
        for fi14 := 0; fi14 < 1; fi14++ {
            fmt.Println(v2, v4, fi14)
            var v15 uint32 = 2
            var v16 int32 = -2147483648
            _ = fi14
            _ = v15
            _ = v16
        }
        v1 -= (v1 * v1)
    }
    v2 = v2
    var v17 float64 = 0.0
    v17 = v17
    v18 := GenS2{F0: (-(v1))}
    if (!((0 > -1) || false)) {
        v1 = v1
    }
    if true {
        for fi19 := 0; fi19 < 2; fi19++ {
            v2 = v2
            var v20 int16 = v4
            _ = fi19
            _ = v20
        }
        var v21 string = "a"
        if ((v2 > "") || true) {
            fmt.Println(v1, v4)
            var v22 int16 = v4
            v3[-989] = false
            _ = v22
        }
        _ = v21
    } else {
        if true {
            v23 := GenS2{F0: v1}
            v24 := GenS1{F0: v4}
            _ = v23
            _ = v24
        }
        for fi25 := 0; fi25 < 18; fi25++ {
            fmt.Println(v17)
            if true {
                for fi26 := 0; fi26 < 15; fi26++ {
                    fmt.Println(v4)
                    var v27 uint32 = 702
                    _ = fi26
                    _ = v27
                }
                fmt.Println(fi25)
                var v28 int16 = (v4 >> ((397) & 63))
                _ = v28
            } else {
                v2 = v2
            }
            _ = fi25
        }
    }
    fmt.Println(v1, v2)
    fmt.Println(v2, v4)
    if false {
        if (false || ((v4 & v4) == 1000)) {
            var v29 int8 = 100
            var v30 int64 = (v1 - v1)
            v1 = v1
            _ = v29
            _ = v30
        }
        switch "" {
        case "\n\t":
            fmt.Println(v4)
        default:
            v31 := GenS3{F0: 2, F1: 2147483648}
            _ = v31
        }
    }
    if true {
        v1 = v1
        v32 := GenS2{F0: (-(v1))}
        for fi33 := 0; fi33 < 6; fi33++ {
            v3[2147483647] = (!(!true))
            switch 0 {
            case -1:
                v34 := GenS3{F0: 309, F1: 0}
                _ = v34
            case 1:
                var v35 uint32 = 2147483647
                _ = v35
            default:
                for fi36 := 0; fi36 < 6; fi36++ {
                    var v37 int16 = v4
                    _ = fi36
                    _ = v37
                }
            }
            _ = fi33
        }
        _ = v32
    }
    v38 := GenS4{F0: -1, F1: (v4 * v4)}
    var v39 int64 = (v1 % ((758) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
}
