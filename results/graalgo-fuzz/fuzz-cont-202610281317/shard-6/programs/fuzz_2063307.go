package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 uint64; F2 bool }
type GenS3 struct { F0 string; F1 uint32; F2 uint8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int16; F1 uint16; F2 uint64 }
type GenS6 struct { F0 string; F1 int8; F2 uint16 }
type GenS7 struct { F0 string; F1 int32 }
func genFn0() bool {
    var v4 float32 = (3.14159 / 0.1)
    if (233 >= 687) {
        for fi5 := 0; fi5 < 18; fi5++ {
            var v6 string = ""
            _ = fi5
            _ = v6
        }
    } else {
        fmt.Println(v4)
        for fi7 := 0; fi7 < 20; fi7++ {
            for fi8 := 0; fi8 < 20; fi8++ {
                v9 := []int{}
                _ = fi8
                _ = v9
            }
            if false {
                var v10 int8 = 100
                for fi11 := 0; fi11 < 10; fi11++ {
                    if true {
                        var v12 bool = (245 > 0)
                        _ = v12
                    }
                    v4 += (float32(0.0) + float32((0.0 * (1e10 - (3.14159 * 1e10)))))
                    _ = fi11
                }
                var v13 uint64 = 2
                _ = v10
                _ = v13
            }
            v14 := GenS1{F0: fi7}
            _ = fi7
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        v4 -= v4
        switch "" {
        case "hello":
            for fi16 := 0; fi16 < 10; fi16++ {
                v4 += (((0.1 * 0.1) / v4) / 1.0)
                _ = fi16
            }
        case "":
            var v17 float32 = float32(float64((float32(float64(((0.0 / v4) - v4))) * v4)))
            _ = v17
        }
        _ = fi15
    }
    v4 += (float32(3.14159) * (v4 + 1e10))
    _ = v4
    return false
}
func genFn1(p0 int8, p1 string) bool {
    v18 := GenS2{F0: 1, F1: 4294967295, F2: (false || true)}
    var v19 int8 = p0
    v20 := GenS3{F0: (p1 + (p1 + p1)), F1: 95, F2: 211}
    for fi21 := 0; fi21 < 8; fi21++ {
        v22 := GenS4{F0: fi21}
        for fi23 := 0; fi23 < 9; fi23++ {
            v24 := []int{fi23, (fi23 << ((328) & 63)), (^(fi21)), (fi21 + fi21)}
            if ((((!false) && (20 >= 0)) && true) && false) {
                var v25 int = fi21
                fmt.Println(p1, fi21, v25)
                v19 -= (v19 + p0)
                if false {
                    var v26 int32 = 42
                    _ = v26
                }
                var v27 int = (-(fi21))
                v28 := GenS5{F0: 20, F1: 826, F2: 1}
                _ = v25
                _ = v27
                _ = v28
            }
            if (!(3.14159 == (0.0 / (((1e10 + 0.0) / 0.1) - -2.25)))) {
                var v29 uint = 788
                v30 := GenS3{F0: p1, F1: 1, F2: 255}
                if false {
                    v24 = append(v24, (fi23 | (fi23 - 4294967295)))
                    if false {
                        var v31 string = p1
                        switch p1 {
                        case "hello":
                            v19 = p0
                        case "a":
                            v29 = (^(v29))
                        }
                        var v32 int32 = 0
                        _ = v31
                        _ = v32
                    }
                } else {
                    _ = len(v24)
                }
                _ = v29
                _ = v30
            }
            _ = fi23
            _ = v24
        }
        _ = fi21
        _ = v22
    }
    _ = p0
    _ = p1
    _ = v18
    _ = v19
    _ = v20
    return (-839 <= -32768)
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{-1, 2147483647, 515}
    var v3 bool = (0.1 != ((((3.14159 * float64(0.1)) + (float64(0.0) - float64((0.0 + 3.14159)))) - 0.1) + 1.5))
    var v33 uint16 = 32767
    fmt.Println(v3, v33)
    if (255 == 42) {
        if len(v2) > 0 {
            v2[0] = -2147483648
        }
    } else {
        fmt.Println(v33)
        v3 = true
    }
    var v34 float64 = 0.1
    if v3 {
        fmt.Println(v3, v33)
    } else {
        v34 = float64((1.5 * (0.0 / float32(v34))))
        for fi35 := 0; fi35 < 15; fi35++ {
            var v36 int8 = 42
            if (v33 < (^(v33))) {
                var v37 uint16 = v33
                _ = v37
            }
            _ = fi35
            _ = v36
        }
        if true {
            v34 = float64((1e10 - (1e10 + (-2.25 + (((3.14159 * float32((1.5 + ((v34 + (v34 - 0.1)) - ((v34 - v34) * v34))))) * -2.25) / 1e10)))))
            v38 := GenS6{F0: ("hello" + ("a" + ("\n\t" + ("a" + (("a" + "hello") + "a"))))), F1: 1, F2: (v33 & (v33 / ((v33) | 1)))}
            _ = v38
        } else {
            var v39 bool = false
            _ = v39
        }
    }
    if (!((3.14159 * 0.1) > 3.14159)) {
        v40 := GenS4{F0: 9223372036854775807}
        _ = v40
    }
    v41 := GenS7{F0: ("" + "a"), F1: 0}
    v3 = (4294967295 == 2147483648)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v41
}
