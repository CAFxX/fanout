package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int64 }
type GenS2 struct { F0 bool; F1 int32; F2 int }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 int16; F1 uint16 }
type GenS5 struct { F0 string; F1 uint16; F2 uint64 }

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[string]uint32)
    var v3 uint32 = 1
    var v4 uint32 = (v3 | (v3 % ((v3) | 1)))
    var v5 int = -965
    var v6 int32 = 42
    v3 += 988
    if false {
        var v7 uint = 0
        var v8 int = (v5 | v5)
        _ = v7
        _ = v8
    }
    var v9 float32 = 3.14159
    var v10 int16 = 0
    var v11 int16 = (v10 << ((42) & 63))
    if (127 <= -1) {
        var v12 int16 = (^(v11))
        var v13 uint32 = (v4 % ((v4) | 1))
        v14 := GenS1{F0: 0, F1: 65535, F2: 0}
        _ = v12
        _ = v13
        _ = v14
    }
    v6 = -56
    v15 := GenS1{F0: 4294967295, F1: 904, F2: (v1 & (-(v1)))}
    v10 += (v11 & 1000)
    v16 := GenS1{F0: 65535, F1: 2, F2: v1}
    if true {
        var v17 int16 = (v10 / (((-(v10))) | 1))
        delete(v2, ("hello" + "\n\t"))
        switch -1 {
        case 2147483647:
            if false {
                if false {
                    if false {
                        if (false || true) {
                            if false {
                                v18 := GenS1{F0: 18446744073709551615, F1: 18446744073709551615, F2: 1}
                                _ = v18
                            } else {
                                v17 = 1
                            }
                        }
                    }
                    v19 := []int{(v5 | -801), (v5 - (v5 / (((v5 * v5)) | 1)))}
                    switch "a" {
                    case "foo bar":
                        var v20 int8 = 125
                        _ = v20
                    default:
                        if (815 != 2) {
                            v6 -= (v6 | 2147483647)
                        }
                    }
                    _ = v19
                } else {
                    for fi21 := 0; fi21 < 6; fi21++ {
                        var v22 uint = 4294967295
                        if (!(v22 >= v22)) {
                            v23 := []int{fi21, fi21, (fi21 >> ((2) & 63)), (v5 >> (((v22 ^ v22)) & 63))}
                            v3 = v4
                            _ = v23
                        }
                        _ = fi21
                        _ = v22
                    }
                    switch 9223372036854775807 {
                    case -9223372036854775808:
                        v24 := GenS2{F0: (true && true), F1: v6, F2: (^(v5))}
                        _ = v24
                    }
                    _ = len(v2)
                }
                v25 := GenS3{F0: 1, F1: (^(v3))}
                v26 := []int{}
                _ = v25
                _ = v26
            }
        case -544:
            v17 += (v17 << ((65535) & 63))
        case 42:
            v27 := GenS1{F0: 38, F1: 4294967295, F2: (v1 | (v1 ^ 42))}
            _ = v27
        default:
            if true {
                if ((!true) || (v1 < v1)) {
                    fmt.Println(v4, v3, v1)
                }
                if (!true) {
                    v28 := GenS1{F0: 4294967295, F1: 626, F2: (v1 - v1)}
                    v1 = (v1 * (v1 >> ((644) & 63)))
                    _ = v28
                }
            }
        }
        _ = v17
    }
    v29 := GenS4{F0: (v10 | (-(v10))), F1: 65535}
    for fi30 := 0; fi30 < 11; fi30++ {
        var v31 uint64 = 346
        switch "hello" {
        case "\n\t":
            var v32 uint16 = 831
            _ = v32
        default:
            if (false || (v6 != (v6 * v6))) {
                if false {
                    fmt.Println(v9, v4)
                    v33 := GenS5{F0: ("foo bar" + "\n\t"), F1: 65535, F2: (v31 - uint64(v4))}
                    _ = v33
                }
                if (v31 > v31) {
                    var v34 float64 = 0.0
                    switch (fi30 ^ (v5 / ((v5) | 1))) {
                    case 32768:
                        for fi35 := 0; fi35 < 13; fi35++ {
                            v36 := GenS1{F0: 255, F1: 715, F2: (v1 % (((v1 >> ((42) & 63))) | 1))}
                            _ = fi35
                            _ = v36
                        }
                    case 9223372036854775807:
                        v37 := GenS5{F0: ((("\n\t" + "a") + "foo bar") + "\n\t"), F1: 1, F2: 2}
                        _ = v37
                    }
                    fmt.Println(v4, v31)
                    _ = v34
                } else {
                    v38 := []int{v5, -805, fi30}
                    _ = v38
                }
                var v39 int = fi30
                _ = v39
            }
        }
        _ = fi30
        _ = v31
    }
    for fi40 := 0; fi40 < 15; fi40++ {
        v41 := GenS5{F0: ("hello" + ("" + "")), F1: 32767, F2: 460}
        fmt.Println(v4, v5, v11)
        v42 := []int{v5}
        _ = fi40
        _ = v41
        _ = v42
    }
    v9 += ((1.5 / (float32(1.5) / v9)) * ((v9 - (v9 * (v9 * float32(0.0)))) + (float32(float64(v9)) / ((v9 * (((float32(1e10) + (v9 / -2.25)) * 0.0) / v9)) / (float32(1.5) / (float32(3.14159) * (v9 / v9)))))))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
