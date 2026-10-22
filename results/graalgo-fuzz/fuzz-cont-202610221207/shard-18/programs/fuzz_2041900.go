package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 int; F1 int16; F2 int }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int16; F1 bool }
func genFn0() bool {
    v7 := []int{132}
    var v8 int16 = -32768
    v7 = append(v7, 9223372036854775807)
    if (v8 < 1000) {
        var v9 uint64 = 1
        _ = v9
    }
    _ = v7
    _ = v8
    return false
}
func genFn1(p0 bool, p1 int8, p2 uint) uint64 {
    p0 = p0
    if p0 {
        v10 := GenS1{F0: 2}
        var v11 int32 = 0
        _ = v10
        _ = v11
    }
    v12 := []int{224, -550}
    if false {
        v13 := GenS2{F0: 255, F1: (p2 - p2)}
        for fi14 := 0; fi14 < 14; fi14++ {
            _ = len(v12)
            var v15 int16 = 867
            _ = fi14
            _ = v15
        }
        _ = v13
    } else {
        var v16 string = ""
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 18446744073709551615
}

func main() {
    v1 := []uint32{4294967295}
    v2 := make(map[int]bool)
    var v3 uint8 = 99
    var v4 float32 = -2.25
    v5 := make(map[uint32]string)
    v6 := []string{("foo bar" + "\n\t"), ("hello" + ("a" + "\n\t")), (("foo bar" + ("\n\t" + "")) + "\n\t"), (("foo bar" + "hello") + "")}
    v3 = v3
    var v17 float64 = 1e10
    for fi18 := 0; fi18 < 2; fi18++ {
        v19 := GenS1{F0: 18446744073709551615}
        var v20 uint8 = v3
        _ = fi18
        _ = v19
        _ = v20
    }
    v21 := GenS2{F0: v3, F1: 0}
    fmt.Println(v3, v3, v3)
    if (!true) {
        var v22 float64 = (-2.25 * 3.14159)
        v23 := GenS2{F0: 1, F1: 292}
        if (666 >= 2) {
            v22 = v22
            switch 32768 {
            case -510:
                fmt.Println(v17, v3)
            case 1:
                v24 := GenS2{F0: v3, F1: 2}
                _ = v24
            default:
                if ((v3 << ((0) & 63)) != v3) {
                    var v25 int16 = -400
                    v4 -= float32(float64(v4))
                    _ = v25
                }
            }
            v17 = (v22 / 3.14159)
        }
        _ = v22
        _ = v23
    } else {
        v26 := GenS3{F0: 127, F1: -1, F2: 388}
        _ = len(v5)
        v4 -= (((float32((v17 / v17)) - (v4 + -2.25)) / v4) - 1e10)
        _ = v26
    }
    v4 = v4
    v27 := GenS4{F0: 2}
    v3 = (v3 ^ (v3 << ((42) & 63)))
    var v28 bool = (!false)
    var v29 string = (("foo bar" + ("" + "a")) + "")
    if v28 {
        var v30 string = v29
        v31 := v2[-2147483648]
        for fi32 := 0; fi32 < 15; fi32++ {
            if false {
                var v33 int64 = -790
                _ = v33
            }
            switch fi32 {
            case 295:
                if false {
                    _ = len(v6)
                } else {
                    v34 := GenS5{F0: -1, F1: (v30 >= (v30 + "\n\t"))}
                    for fi35 := 0; fi35 < 13; fi35++ {
                        v36 := GenS3{F0: (-(fi32)), F1: -58, F2: int(v3)}
                        var v37 float32 = v4
                        var v38 float32 = (float32((((float64((v4 - -2.25)) + (3.14159 * float64(v4))) / (-2.25 / float64((v37 - (v4 / v4))))) - v17)) + v37)
                        _ = fi35
                        _ = v36
                        _ = v37
                        _ = v38
                    }
                    var v39 uint64 = genFn1(true, 19, 596)
                    _ = v34
                    _ = v39
                }
            case 42:
                v17 = -2.25
            case -920:
                v40 := GenS4{F0: 32767}
                _ = v40
            }
            _ = fi32
        }
        _ = v30
        _ = v31
    }
    _ = len(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
