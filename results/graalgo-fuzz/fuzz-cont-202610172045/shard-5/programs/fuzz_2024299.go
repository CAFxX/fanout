package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 int64; F2 uint }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int16 }
func genFn0(p0 string, p1 int) int64 {
    v6 := []int{(p1 - 4294967295), (p1 - (-(p1)))}
    p0 = (((p0 + "a") + "") + p0)
    var v7 uint16 = 65535
    v8 := GenS1{F0: 128, F1: p0}
    var v9 float64 = 1e10
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return -958
}

func main() {
    var v1 float32 = -2.25
    var v2 string = "a"
    var v3 int8 = 6
    var v4 uint = 63
    var v5 int32 = -898
    if (42 <= 1) {
        if true {
            fmt.Println(v4, v3, v1)
            for fi10 := 0; fi10 < 17; fi10++ {
                v11 := []int{fi10, 888, fi10, 0}
                var v12 float32 = v1
                switch (fi10 >> ((1) & 63)) {
                case 0:
                    var v13 uint = (v4 - 255)
                    _ = v13
                case 81:
                    var v14 float32 = (0.0 + 0.1)
                    _ = v14
                default:
                    if (0.0 <= (float64(v12) / (1e10 + float64((((float32(3.14159) + (float32(1.5) / v1)) / ((v1 * float32(1.5)) * ((-2.25 / v12) / (1e10 * v12)))) + v1))))) {
                        fmt.Println(v2, v4)
                    } else {
                        if true {
                            v5 -= (v5 >> (((v4 - (^(v4)))) & 63))
                            var v15 int = (^(fi10))
                            _ = v15
                        } else {
                            for fi16 := 0; fi16 < 18; fi16++ {
                                if (((0.0 / ((((3.14159 * -2.25) - 3.14159) - float64(v12)) - 3.14159)) == 1e10) && (-2.25 <= 3.14159)) {
                                    v12 -= (1e10 * (-2.25 / 3.14159))
                                    v4 += 769
                                }
                                _ = fi16
                            }
                        }
                        v11 = append(v11, 9223372036854775807)
                    }
                }
                _ = fi10
                _ = v11
                _ = v12
            }
            v5 -= v5
        } else {
            fmt.Println(v5)
            if (-32768 < -32768) {
                var v17 float32 = float32(((((3.14159 + 3.14159) + 1.5) / 1.5) * -2.25))
                if true {
                    var v18 uint8 = 255
                    _ = v18
                }
                _ = v17
            } else {
                v19 := []int{}
                v2 = ""
                _ = v19
            }
            v5 = v5
        }
        v20 := []int{0, 32768}
        _ = v20
    }
    v21 := []int{32768, -1, 0, -1}
    v4 -= v4
    if true {
        for fi22 := 0; fi22 < 4; fi22++ {
            var v23 uint8 = 128
            if len(v21) >= 2 {
                v21 = v21[1:len(v21)]
            }
            v24 := GenS2{F0: 255}
            _ = fi22
            _ = v23
            _ = v24
        }
        for fi25 := 0; fi25 < 13; fi25++ {
            v3 = 1
            _ = fi25
        }
    } else {
        if true {
            fmt.Println(v5, v3)
            v26 := GenS2{F0: 2}
            _ = v26
        }
        if (v3 >= 37) {
            if (!false) {
                switch v2 {
                case "a":
                    v2 = "\n\t"
                case "foo bar":
                    switch -1 {
                    case 486:
                        v27 := GenS3{F0: 4294967295, F1: 0, F2: (^(v4))}
                        _ = v27
                    case 843:
                        var v28 uint32 = 0
                        _ = v28
                    default:
                        var v29 int8 = (v3 - (v3 + -96))
                        _ = v29
                    }
                case "hello":
                    v4 = v4
                }
                if (1 >= 42) {
                    v30 := GenS1{F0: 2147483647, F1: v2}
                    _ = v30
                } else {
                    var v31 uint32 = 0
                    var v32 uint16 = 255
                    v33 := GenS4{F0: 128}
                    _ = v31
                    _ = v32
                    _ = v33
                }
            }
        }
        v34 := GenS2{F0: 2}
        _ = v34
    }
    var v35 int16 = 1000
    if false {
        _ = len(v21)
        switch v2 {
        case "a":
            var v36 uint64 = 18446744073709551615
            _ = v36
        default:
            v4 += 1
        }
    } else {
        v4 += v4
        var v37 uint = (v4 - v4)
        switch v2 {
        case "foo bar":
            v1 = 0.1
        }
        _ = v37
    }
    var v38 bool = true
    v39 := GenS5{F0: 1}
    fmt.Println(v38)
    v5 = (v5 >> ((196) & 63))
    fmt.Println(v2, v38)
    v40 := GenS6{F0: (v35 & -326)}
    var v41 float32 = 0.0
    fmt.Println(v41)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v35
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
