package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 uint16; F1 int; F2 bool }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8; F1 int }
func genFn0() uint8 {
    v4 := GenS1{F0: 4294967295, F1: 42}
    if false {
        var v5 uint64 = 2
        v6 := []int{2147483647, -194, -269, 32768}
        _ = v5
        _ = v6
    }
    v7 := []int{-248, -1}
    var v8 uint32 = 2147483647
    var v9 uint32 = v8
    _ = v4
    _ = v7
    _ = v8
    _ = v9
    return 150
}
func genFn1(p0 uint64, p1 uint64) int16 {
    for fi10 := 0; fi10 < 7; fi10++ {
        fmt.Println(fi10)
        _ = fi10
    }
    switch 211 {
    case 2147483647:
        switch 0 {
        case 431:
            for fi11 := 0; fi11 < 19; fi11++ {
                p0 += (p1 / ((150) | 1))
                switch fi11 {
                case 42:
                    fmt.Println(fi11, p0, fi11)
                }
                v12 := []int{}
                _ = fi11
                _ = v12
            }
        case -703:
            fmt.Println(p1)
        }
    }
    _ = p0
    _ = p1
    return -32768
}

func main() {
    var v1 uint = 930
    v2 := make(map[int]uint32)
    var v3 uint64 = 707
    v13 := []int{}
    var v14 string = (("\n\t" + "hello") + ("\n\t" + "a"))
    if (uint8(v3) > 255) {
        if (142 < 4294967295) {
            v14 = ("" + v14)
        }
        for fi15 := 0; fi15 < 12; fi15++ {
            var v16 uint64 = v3
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            v16 = (v16 * v3)
            _ = fi15
            _ = v16
        }
    } else {
        v1 = 0
        v17 := v2[2147483647]
        _ = v17
    }
    var v18 float64 = (float64(1e10) * 1.5)
    v19 := GenS2{F0: 0, F1: 414, F2: (!(!true))}
    v13 = append(v13, -97)
    var v20 float64 = (v18 + (0.0 - v18))
    for fi21 := 0; fi21 < 13; fi21++ {
        for fi22 := 0; fi22 < 7; fi22++ {
            v3 = (v3 << ((v1) & 63))
            v14 = ("\n\t" + "hello")
            if (v20 != -2.25) {
                v3 += 18446744073709551615
                var v23 float32 = 0.0
                switch v14 {
                case "a":
                    var v24 float64 = 0.0
                    if false {
                        if (!false) {
                            v25 := GenS2{F0: 0, F1: fi22, F2: (!(true || true))}
                            fmt.Println(v3)
                            _ = v25
                        } else {
                            _ = len(v13)
                            fmt.Println(fi21, v20)
                        }
                        var v26 int8 = 42
                        switch (fi21 | -338) {
                        case 824:
                            var v27 int32 = 2147483647
                            _ = v27
                        }
                        _ = v26
                    } else {
                        var v28 uint16 = 661
                        v24 += v20
                        _ = v28
                    }
                    v2[(fi21 * fi22)] = 309
                    var v29 uint32 = 42
                    switch "" {
                    case "a":
                        switch (fi21 + fi21) {
                        case 393:
                            v3 -= (v3 >> ((v1) & 63))
                        }
                    }
                    _ = v24
                    _ = v29
                default:
                    fmt.Println(fi22, v1)
                }
                switch "hello" {
                case "a":
                    if len(v13) > 0 {
                        v13[0] = (fi22 / ((fi22) | 1))
                    }
                case "\n\t":
                    switch v14 {
                    case "":
                        v30 := GenS2{F0: 659, F1: (fi22 / (((-(fi22))) | 1)), F2: (!true)}
                        _ = v30
                    default:
                        delete(v2, (fi22 | (fi21 % ((-9223372036854775808) | 1))))
                    }
                case "":
                    switch 128 {
                    case 27:
                        var v31 float32 = v23
                        _ = v31
                    case 32767:
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                    case -2147483648:
                        v32 := GenS2{F0: 65535, F1: -9223372036854775808, F2: true}
                        _ = v32
                    }
                }
                fmt.Println(v1, v1)
                _ = v23
            } else {
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
                v33 := GenS3{F0: 42}
                _ = v33
            }
            v20 = v20
            v20 = 3.14159
            var v34 int64 = -515
            v34 += (v34 ^ (v34 & v34))
            var v35 float32 = (float32((((v20 - ((v18 - v18) - 3.14159)) - 3.14159) / v18)) + float32(-2.25))
            if false {
                v36 := v2[1]
                var v37 int = fi22
                v14 = ""
                v36 = 1
                fmt.Println(v37)
                var v38 uint = 568
                _ = v36
                _ = v37
                _ = v38
            } else {
                v39 := GenS4{F0: 1, F1: 976}
                _ = v39
            }
            _ = fi22
            _ = v34
            _ = v35
        }
        _ = fi21
    }
    var v40 uint64 = v3
    switch (v14 + "hello") {
    case "a":
        v14 = ""
    default:
        v14 = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v40
}
