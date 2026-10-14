package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string }
func genFn0(p0 uint64) int64 {
    for fi6 := 0; fi6 < 15; fi6++ {
        p0 = 4294967295
        v7 := []int{32768}
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 379}
    for fi9 := 0; fi9 < 2; fi9++ {
        v10 := []int{fi9, (fi9 << ((189) & 63)), fi9}
        _ = fi9
        _ = v10
    }
    if false {
        p0 = (p0 % (((p0 / ((p0) | 1))) | 1))
        if (true || true) {
            fmt.Println(p0)
            if (!(p0 < (p0 << ((4294967295) & 63)))) {
                fmt.Println(p0)
                for fi11 := 0; fi11 < 9; fi11++ {
                    if true {
                        v12 := []int{}
                        p0 = (p0 % (((^(p0))) | 1))
                        var v13 int8 = -128
                        _ = v12
                        _ = v13
                    } else {
                        for fi14 := 0; fi14 < 3; fi14++ {
                            var v15 int64 = 42
                            var v16 int16 = 42
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                    }
                    _ = fi11
                }
            }
        }
    }
    var v17 int64 = -1
    _ = p0
    _ = v17
    _ = v8
    return v17
}
func genFn1(p0 int, p1 uint32) uint {
    var v18 int16 = 268
    var v19 bool = false
    v20 := []int{32767, (^(p0)), p0}
    if (1e10 < 3.14159) {
        switch p0 {
        case 42:
            var v21 uint = 0
            _ = v21
        }
    }
    _ = p0
    _ = p1
    _ = v18
    _ = v19
    _ = v20
    return 2
}

func main() {
    var v1 string = ""
    v2 := make(map[int]bool)
    var v3 int32 = 2147483647
    var v4 int8 = -1
    v5 := make(map[uint32]bool)
    fmt.Println(v1, v3)
    for fi22 := 0; fi22 < 20; fi22++ {
        var v23 int8 = (v4 % (((v4 & v4)) | 1))
        fmt.Println(v1, v23, v4)
        v24 := []int{(fi22 % ((-9223372036854775808) | 1))}
        _ = fi22
        _ = v23
        _ = v24
    }
    switch 9223372036854775807 {
    case 2147483647:
        if (0.1 != ((((3.14159 * 0.0) / (((1e10 / -2.25) - -2.25) + 3.14159)) - float64(1.5)) - (1e10 + -2.25))) {
            _ = len(v2)
        }
    case 4294967295:
        var v25 uint64 = 4294967295
        _ = v25
    case 1:
        if true {
            for fi26 := 0; fi26 < 9; fi26++ {
                var v27 bool = (!true)
                v3 = v3
                _ = fi26
                _ = v27
            }
            var v28 string = (v1 + v1)
            _ = v28
        }
    default:
        delete(v5, 2147483648)
    }
    if (1 == 530) {
        if (true || true) {
            v29 := GenS1{F0: v3}
            v4 = (v4 | v4)
            _ = v29
        }
    } else {
        fmt.Println(v1)
        switch "foo bar" {
        case "hello":
            v4 += (v4 + (-(v4)))
        case "foo bar":
            v30 := GenS1{F0: 0}
            _ = v30
        default:
            v31 := GenS2{F0: ("\n\t" + v1)}
            _ = v31
        }
        var v32 int32 = v3
        _ = v32
    }
    switch 280 {
    case 32767:
        var v33 float64 = (-2.25 - 3.14159)
        _ = v33
    case 4294967295:
        var v34 int = 2147483647
        _ = v34
    default:
        for fi35 := 0; fi35 < 11; fi35++ {
            fmt.Println(fi35)
            _ = fi35
        }
    }
    v36 := GenS1{F0: v3}
    v37 := []int{0, -230, -1}
    v38 := GenS1{F0: -296}
    if ((v4 | v4) < v4) {
        v4 += 127
    } else {
        v37 = append(v37, -9223372036854775808)
        var v39 int = 646
        v39 = (v39 >> ((18446744073709551615) & 63))
        _ = v39
    }
    v4 -= (-(v4))
    _ = len(v37)
    var v40 uint8 = 128
    delete(v2, -1)
    _ = len(v37)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v40
    _ = v5
}
