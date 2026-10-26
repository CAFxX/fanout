package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 int16 }
func genFn0() uint64 {
    if (112 > 42) {
        var v5 uint = 348
        fmt.Println(v5)
        _ = v5
    } else {
        v6 := GenS1{F0: 864, F1: 912}
        v7 := []int{0, 4294967295}
        var v8 int16 = -1
        _ = v6
        _ = v7
        _ = v8
    }
    if ((!true) || false) {
        if true {
            for fi9 := 0; fi9 < 3; fi9++ {
                v10 := []int{-950, (fi9 << ((0) & 63)), (-(fi9)), 128}
                for fi11 := 0; fi11 < 17; fi11++ {
                    var v12 uint16 = 255
                    var v13 uint16 = v12
                    v14 := GenS1{F0: 1000, F1: 18446744073709551615}
                    v13 -= 767
                    if (fi9 == (fi11 + fi11)) {
                        v13 -= 1
                        v15 := GenS1{F0: 0, F1: uint64(v13)}
                        v12 -= 449
                        _ = v15
                    }
                    var v16 float64 = (1e10 - 1.5)
                    var v17 int = 1
                    _ = fi11
                    _ = v12
                    _ = v13
                    _ = v14
                    _ = v16
                    _ = v17
                }
                _ = fi9
                _ = v10
            }
            var v18 uint8 = 128
            _ = v18
        } else {
            if false {
                if (true && true) {
                    for fi19 := 0; fi19 < 6; fi19++ {
                        fmt.Println(fi19)
                        if true {
                            var v20 uint64 = 38
                            _ = v20
                        }
                        v21 := []int{}
                        _ = fi19
                        _ = v21
                    }
                }
                for fi22 := 0; fi22 < 16; fi22++ {
                    var v23 float32 = float32((0.0 * (1e10 * (1.5 * (1.5 / float64(((((1e10 * ((3.14159 + (-2.25 - (((3.14159 / 1.5) * 3.14159) + 3.14159))) + float32((-2.25 / (0.0 - float64(3.14159)))))) / float32(1e10)) + 1e10) + -2.25)))))))
                    if true {
                        var v24 int = (fi22 << ((255) & 63))
                        v23 = float32(((0.1 + (-2.25 / 1.5)) - -2.25))
                        v25 := GenS1{F0: 42, F1: 2}
                        _ = v24
                        _ = v25
                    }
                    for fi26 := 0; fi26 < 13; fi26++ {
                        var v27 int = (fi22 ^ fi22)
                        var v28 string = "foo bar"
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                    _ = fi22
                    _ = v23
                }
                var v29 float64 = (0.1 - 1.5)
                _ = v29
            }
            if true {
                var v30 int16 = 42
                _ = v30
            }
        }
    }
    var v31 int8 = 100
    var v32 int32 = 435
    v33 := GenS2{F0: 871}
    _ = v31
    _ = v32
    _ = v33
    return 2
}
func genFn1(p0 int8, p1 uint64) uint {
    v34 := GenS1{F0: 1, F1: 0}
    fmt.Println(p0)
    var v35 int8 = p0
    _ = p0
    _ = p1
    _ = v34
    _ = v35
    return 42
}

func main() {
    var v1 int8 = 0
    v2 := []uint32{42, 2, 65535, 2147483648}
    var v3 float64 = 0.1
    v4 := []int{}
    var v36 int8 = v1
    var v37 bool = false
    if (!v37) {
        v38 := GenS3{F0: (v1 * (v1 >> ((65535) & 63))), F1: -32768}
        _ = v38
    } else {
        v3 = 0.0
    }
    fmt.Println(v37)
    var v39 uint = 162
    v1 = (v36 - v36)
    v3 -= (v3 * (0.0 - v3))
    for fi40 := 0; fi40 < 20; fi40++ {
        var v41 uint = (v39 & genFn1((v36 | v1), 4294967295))
        var v42 int = fi40
        _ = fi40
        _ = v41
        _ = v42
    }
    v1 = 127
    for fi43 := 0; fi43 < 4; fi43++ {
        v4 = append(v4, (^(fi43)))
        _ = fi43
    }
    var v44 int16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v39
    _ = v4
    _ = v44
}
