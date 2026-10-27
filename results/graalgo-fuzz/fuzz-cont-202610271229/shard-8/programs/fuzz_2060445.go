package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int }
type GenS2 struct { F0 int16; F1 int8; F2 int }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int16 }
func genFn0(p0 uint8, p1 int32) uint32 {
    p1 += p1
    switch 205 {
    case -9223372036854775808:
        if (724 < -21) {
            v6 := []int{-9223372036854775808, 32767, 0, -608}
            _ = v6
        }
    case -1:
        var v7 bool = true
        _ = v7
    default:
        var v8 string = "hello"
        _ = v8
    }
    if true {
        v9 := GenS1{F0: 42, F1: p0, F2: -502}
        _ = v9
    } else {
        fmt.Println(p0, p1)
        v10 := []int{0}
        var v11 uint16 = 1
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    return 2
}
func genFn1(p0 int64, p1 uint8) bool {
    var v12 bool = false
    for fi13 := 0; fi13 < 4; fi13++ {
        for fi14 := 0; fi14 < 10; fi14++ {
            p1 -= (p1 + p1)
            _ = fi14
        }
        _ = fi13
    }
    v15 := []int{1, -940, 127}
    _ = p0
    _ = p1
    _ = v12
    _ = v15
    return (!v12)
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[int]uint32)
    var v3 uint32 = 83
    var v4 uint8 = 128
    var v5 int64 = 1
    v16 := GenS1{F0: v3, F1: v4, F2: 127}
    for fi17 := 0; fi17 < 13; fi17++ {
        switch fi17 {
        case -1:
            if true {
                v1 = ""
                var v18 uint32 = genFn0(249, 42)
                _ = v18
            }
        case -95:
            v5 += (-(v5))
        case 128:
            v3 = v3
        }
        v19 := GenS1{F0: 804, F1: (v4 + v4), F2: fi17}
        if false {
            fmt.Println(v5)
        } else {
            if false {
                if false {
                    fmt.Println(v3, v3, v1)
                    v5 += v5
                    v20 := GenS1{F0: v3, F1: (v4 / (((v4 * 245)) | 1)), F2: (-(fi17))}
                    _ = v20
                } else {
                    for fi21 := 0; fi21 < 12; fi21++ {
                        v22 := []int{(fi21 % ((fi17) | 1))}
                        if true {
                            v1 = v1
                            if false {
                                fmt.Println(v1, v1, fi21)
                            }
                            if false {
                                switch (v1 + v1) {
                                case "hello":
                                    v1 = (("\n\t" + v1) + ((v1 + v1) + ""))
                                case "a":
                                    var v23 int16 = 1000
                                    _ = v23
                                }
                                switch "hello" {
                                case "a":
                                    if len(v22) > 0 {
                                        v22[0] = (fi17 - (fi17 ^ (fi17 - (fi17 % ((4294967295) | 1)))))
                                    }
                                case "\n\t":
                                    switch fi17 {
                                    case 4294967295:
                                        fmt.Println(v4)
                                    }
                                default:
                                    v24 := GenS1{F0: (^(v3)), F1: (v4 | v4), F2: 32767}
                                    _ = v24
                                }
                                _ = len(v2)
                            }
                        } else {
                            v2[fi17] = genFn0(255, 0)
                            v25 := GenS2{F0: 60, F1: 0, F2: 128}
                            v1 = (((v1 + v1) + (v1 + v1)) + v1)
                            _ = v25
                        }
                        switch (fi17 >> ((0) & 63)) {
                        case 32768:
                            v1 = v1
                        }
                        v26 := GenS1{F0: 42, F1: v4, F2: (fi17 - (fi21 >> ((45) & 63)))}
                        _ = fi21
                        _ = v22
                        _ = v26
                    }
                    var v27 bool = false
                    _ = v27
                }
                v28 := GenS3{F0: -830}
                v5 = -24
                _ = v28
            }
            _ = len(v2)
            v29 := GenS1{F0: v3, F1: (v4 | (v4 & 233)), F2: 127}
            _ = v29
        }
        _ = fi17
        _ = v19
    }
    switch 2147483647 {
    case 0:
        v30 := []int{}
        _ = v30
    }
    var v31 int64 = (v5 + (v5 / ((-680) | 1)))
    fmt.Println(v31)
    var v32 float64 = 0.1
    v32 -= ((-2.25 * (1e10 - 1e10)) / 1.0)
    v33 := GenS4{F0: 32767}
    var v34 int32 = -2147483648
    fmt.Println(v32, v4)
    delete(v2, -503)
    switch "\n\t" {
    case "":
        v35 := []int{-60, 32767, 42}
        _ = v35
    default:
        v36 := []int{}
        _ = v36
    }
    var v37 int16 = 1
    if false {
        for fi38 := 0; fi38 < 18; fi38++ {
            var v39 int64 = v5
            _ = fi38
            _ = v39
        }
        v37 = (v37 & (v37 ^ (v37 >> ((255) & 63))))
    } else {
        v40 := GenS2{F0: v37, F1: 56, F2: -737}
        _ = v40
    }
    v41 := GenS3{F0: -134}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v37
    _ = v4
    _ = v41
    _ = v5
}
