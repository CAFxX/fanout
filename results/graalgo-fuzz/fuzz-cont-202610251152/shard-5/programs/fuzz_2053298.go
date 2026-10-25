package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int8; F1 uint64; F2 uint32 }
func genFn0(p0 uint32, p1 uint64, p2 uint8) string {
    v7 := []int{-744, 48, 726, -469}
    if ("hello" > "\n\t") {
        for fi8 := 0; fi8 < 6; fi8++ {
            for fi9 := 0; fi9 < 12; fi9++ {
                fmt.Println(p0, fi9)
                if ((1e10 - ((1e10 + 0.1) - (0.1 * (3.14159 / 3.14159)))) <= (float64(0.0) / 1e10)) {
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    switch (fi9 % (((fi9 << ((0) & 63))) | 1)) {
                    case 831:
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                    }
                    fmt.Println(fi9)
                } else {
                    _ = len(v7)
                    p1 += 4294967295
                }
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                var v10 int64 = 884
                _ = fi9
                _ = v10
            }
            var v11 bool = true
            _ = fi8
            _ = v11
        }
    }
    var v12 uint = 42
    if true {
        v13 := GenS1{F0: 399}
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v7
    return "\n\t"
}
func genFn1(p0 int32) string {
    p0 -= p0
    if false {
        v14 := []int{}
        var v15 uint64 = 115
        _ = v14
        _ = v15
    }
    _ = p0
    return "a"
}

func main() {
    var v1 int64 = -1
    v2 := []uint32{42, 4294967295, 263, 2147483647}
    v3 := make(map[int]string)
    v4 := make(map[int]uint32)
    v5 := make(map[string]string)
    v6 := make(map[string]uint32)
    v1 += 1
    var v16 int64 = 9223372036854775807
    v17 := GenS1{F0: 1}
    if (!false) {
        v1 = (^(v1))
        if (false && true) {
            var v18 uint8 = 114
            for fi19 := 0; fi19 < 13; fi19++ {
                var v20 uint64 = 18446744073709551615
                if ((v20 & 2) <= 2) {
                    var v21 int64 = 9223372036854775807
                    v2 = append(v2, 42)
                    v22 := GenS1{F0: 763}
                    _ = v21
                    _ = v22
                }
                if true {
                    if true {
                        if false {
                            v23 := GenS1{F0: -388}
                            fmt.Println(v1)
                            _ = v23
                        } else {
                            var v24 float64 = ((3.14159 / float64(3.14159)) / (float64(((-2.25 - 1.5) + (((1e10 / ((1.5 - (-2.25 - 1.5)) / 0.1)) / (-2.25 + float32(1.5))) / ((-2.25 + -2.25) + float32(0.1))))) - 0.1))
                            _ = v24
                        }
                        var v25 uint64 = v20
                        _ = v25
                    }
                    v16 = v1
                    fmt.Println(fi19, fi19)
                }
                _ = fi19
                _ = v20
            }
            if ((0 > 1) || true) {
                if true {
                    var v26 uint8 = (v18 | v18)
                    if (32768 == -588) {
                        v18 = v18
                    }
                    _ = v26
                }
                for fi27 := 0; fi27 < 18; fi27++ {
                    for fi28 := 0; fi28 < 6; fi28++ {
                        fmt.Println(fi27, v16, v1)
                        _ = len(v2)
                        if false {
                            _ = len(v2)
                            var v29 bool = false
                            v16 = v16
                            _ = v29
                        } else {
                            var v30 string = ("a" + "")
                            var v31 uint16 = 2
                            _ = v30
                            _ = v31
                        }
                        _ = fi28
                    }
                    v32 := GenS1{F0: -359}
                    if false {
                        if len(v2) > 0 {
                            v2[0] = 42
                        }
                        v33 := GenS1{F0: 42}
                        _ = v33
                    } else {
                        v18 = v18
                    }
                    _ = fi27
                    _ = v32
                }
            }
            _ = v18
        } else {
            v1 = v16
        }
    }
    v34 := GenS1{F0: -1}
    v16 = v1
    var v35 float32 = (float32(3.14159) + (3.14159 - 3.14159))
    v35 += v35
    v16 += -1
    if (!(127 >= 24)) {
        for fi36 := 0; fi36 < 8; fi36++ {
            v37 := GenS2{F0: 0}
            _ = fi36
            _ = v37
        }
    } else {
        if (431 != 2147483647) {
            v16 += (v16 ^ v1)
            v38 := GenS3{F0: 27}
            var v39 bool = (!false)
            _ = v38
            _ = v39
        }
        v40 := GenS3{F0: 42}
        var v41 float64 = -2.25
        _ = v40
        _ = v41
    }
    v42 := GenS4{F0: -79, F1: 42, F2: 0}
    v1 = v16
    for fi43 := 0; fi43 < 17; fi43++ {
        v44 := GenS4{F0: 1, F1: 545, F2: 42}
        fmt.Println(v16)
        var v45 float32 = (3.14159 / v35)
        _ = fi43
        _ = v44
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
