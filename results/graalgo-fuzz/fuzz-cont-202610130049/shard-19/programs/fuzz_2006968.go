package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 uint64 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string; F1 bool; F2 bool }
func genFn0(p0 int8, p1 int64) uint {
    var v6 uint32 = 2147483647
    for fi7 := 0; fi7 < 16; fi7++ {
        for fi8 := 0; fi8 < 20; fi8++ {
            v9 := GenS1{F0: "hello", F1: 18446744073709551615, F2: 2}
            v10 := []int{fi7, (fi7 ^ (-(fi8))), fi7, 486}
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    v6 -= v6
    _ = p0
    _ = p1
    _ = v6
    return 325
}
func genFn1(p0 int16) int8 {
    var v11 uint32 = 746
    v12 := GenS1{F0: ("foo bar" + "\n\t"), F1: 42, F2: 2}
    v13 := []int{-308}
    if ((true || (510 > v11)) && (237 < -1)) {
        v14 := GenS2{F0: 0, F1: 2147483647}
        _ = v14
    }
    var v15 int8 = -128
    _ = p0
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    return (v15 >> ((0) & 63))
}

func main() {
    var v1 bool = false
    var v2 float64 = 1.5
    v3 := []string{"\n\t"}
    var v4 float64 = v2
    var v5 uint32 = 0
    var v16 int64 = -1
    var v17 float32 = ((3.14159 - 1.5) - -2.25)
    v18 := GenS3{F0: true}
    var v19 float64 = -2.25
    fmt.Println(v16, v1)
    v20 := GenS4{F0: "foo bar", F1: v1, F2: (!true)}
    if (-2.25 == 0.1) {
        if v1 {
            v1 = true
            var v21 uint64 = uint64(v5)
            _ = v21
        }
        for fi22 := 0; fi22 < 14; fi22++ {
            if v1 {
                v23 := GenS4{F0: ("foo bar" + "a"), F1: true, F2: (((fi22 >> ((0) & 63)) != fi22) && (!v1))}
                fmt.Println(v16)
                v24 := GenS1{F0: (("foo bar" + ((("foo bar" + "\n\t") + "") + "\n\t")) + ((("foo bar" + "foo bar") + "\n\t") + ("a" + "foo bar"))), F1: genFn0(40, -1), F2: 2}
                _ = v23
                _ = v24
            }
            _ = fi22
        }
    } else {
        for fi25 := 0; fi25 < 18; fi25++ {
            if v1 {
                var v26 uint16 = 42
                _ = v26
            } else {
                fmt.Println(v2, v1)
            }
            fmt.Println(v17, v5, v4)
            _ = fi25
        }
    }
    for fi27 := 0; fi27 < 1; fi27++ {
        fmt.Println(v17, v16)
        v1 = (695 == 1)
        _ = fi27
    }
    v2 = (((1e10 + 1e10) - 1e10) + float64((0.1 * (0.0 * (((v17 * v17) / (v17 - v17)) * v17)))))
    fmt.Println(v5, v4)
    v17 = (v17 - (v17 + ((v17 / v17) + v17)))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
