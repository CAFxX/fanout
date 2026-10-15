package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint16; F1 int8; F2 int }
type GenS6 struct { F0 int16; F1 uint8 }
func genFn0(p0 int16, p1 int64) uint16 {
    v7 := []int{-9223372036854775808, 1, 4294967295, 257}
    switch "hello" {
    case "\n\t":
        p1 = p1
    case "a":
        v8 := GenS1{F0: 18446744073709551615}
        _ = v8
    case "hello":
        p1 = (p1 - -13)
    }
    _ = p0
    _ = p1
    _ = v7
    return 2
}
func genFn1(p0 int16, p1 uint16) int64 {
    v9 := GenS1{F0: 2}
    v10 := GenS1{F0: 18446744073709551615}
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 9223372036854775807
}

func main() {
    var v1 int16 = 32767
    v2 := make(map[int]string)
    var v3 uint8 = 2
    v4 := make(map[uint32]int)
    var v5 int8 = -1
    var v6 int16 = v1
    var v11 uint64 = 1
    fmt.Println(v5, v5, v3)
    switch -1 {
    case 4294967295:
        if (!false) {
            var v12 int8 = v5
            var v13 int32 = -1
            _ = v12
            _ = v13
        } else {
            v14 := GenS1{F0: v11}
            _ = v14
        }
    default:
        v1 -= (v1 + (-(v1)))
    }
    switch "a" {
    case "a":
        var v15 int16 = -129
        _ = v15
    }
    switch ("\n\t" + "\n\t") {
    case "\n\t":
        v5 = 42
    }
    var v16 int16 = 42
    if (v6 <= 0) {
        var v17 uint = 1
        v1 = v1
        _ = v17
    }
    v11 -= v11
    switch "a" {
    case "foo bar":
        v16 -= v16
    case "":
        v18 := []int{9223372036854775807, -61, 2147483647, 32768}
        _ = v18
    case "\n\t":
        var v19 int = 307
        _ = v19
    }
    v3 = 104
    v20 := GenS2{F0: 397, F1: (v6 % ((v6) | 1))}
    v5 -= 42
    for fi21 := 0; fi21 < 19; fi21++ {
        v22 := GenS3{F0: genFn1((v16 ^ v16), 1)}
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 14; fi23++ {
        if false {
            for fi24 := 0; fi24 < 10; fi24++ {
                v5 -= (v5 / ((int8(v3)) | 1))
                switch "hello" {
                case "a":
                    if false {
                        v11 = v11
                    }
                case "foo bar":
                    if (18446744073709551615 >= 4294967295) {
                        _ = len(v2)
                    }
                default:
                    v11 -= (v11 / ((0) | 1))
                }
                var v25 int64 = -988
                if true {
                    v1 += v16
                    v4[35] = (^(fi24))
                } else {
                    var v26 uint = 728
                    var v27 bool = false
                    if v27 {
                        var v28 float64 = float64(((1e10 - ((float32(float64(0.0)) * 1.5) + float32(1e10))) / float32((0.1 / (float64((-2.25 - 0.0)) / 3.14159)))))
                        v29 := GenS1{F0: v11}
                        _ = v28
                        _ = v29
                    } else {
                        v1 = 869
                        v3 += uint8(v11)
                        if (1000 >= (v6 >> (((v26 ^ v26)) & 63))) {
                            v27 = v27
                        }
                    }
                    _ = v26
                    _ = v27
                }
                v25 = v25
                _ = fi24
                _ = v25
            }
        } else {
            for fi30 := 0; fi30 < 13; fi30++ {
                if (-1 < 0) {
                    v31 := GenS3{F0: -995}
                    delete(v2, -9223372036854775808)
                    _ = v31
                } else {
                    v32 := v4[512]
                    switch (v32 + 42) {
                    case 16:
                        var v33 int64 = genFn1(-32768, 397)
                        _ = v33
                    case -1:
                        if true {
                            v34 := GenS3{F0: int64(fi30)}
                            var v35 int64 = 42
                            v36 := v2[fi30]
                            switch v36 {
                            case "foo bar":
                                fmt.Println(v5, v16, v35)
                            case "a":
                                v37 := GenS2{F0: uint32(v3), F1: (v6 >> ((1) & 63))}
                                _ = v37
                            }
                            if (true && (false && true)) {
                                if (v36 > v36) {
                                    var v38 float32 = 0.1
                                    fmt.Println(v3, v6, v38)
                                    _ = v38
                                } else {
                                    var v39 float64 = 3.14159
                                    v39 += v39
                                    v40 := GenS4{F0: 42}
                                    _ = v39
                                    _ = v40
                                }
                                var v41 uint8 = (v3 + v3)
                                v42 := GenS3{F0: 89}
                                if (2147483648 >= 42) {
                                    var v43 float64 = (3.14159 - (-2.25 / 1.0))
                                    v44 := []int{v32, (fi30 / ((v32) | 1))}
                                    _ = v43
                                    _ = v44
                                } else {
                                    v45 := GenS3{F0: v35}
                                    v46 := []int{(fi30 / ((42) | 1)), fi30, fi23, (fi30 | fi23)}
                                    _ = len(v4)
                                    _ = v45
                                    _ = v46
                                }
                                _ = v41
                                _ = v42
                            } else {
                                var v47 uint16 = 42
                                _ = v47
                            }
                            _ = v34
                            _ = v35
                            _ = v36
                        }
                    }
                    var v48 uint = 995
                    fmt.Println(v16, v32)
                    v6 = v6
                    delete(v2, fi30)
                    if (2147483647 <= -588) {
                        v1 = 1000
                        if (499 >= 255) {
                            v49 := GenS2{F0: 498, F1: v16}
                            _ = v49
                        }
                        v5 = 127
                    }
                    var v50 bool = (false && ((true || (genFn0(v1, -1) == 32767)) && (int32(v48) >= -931)))
                    _ = v32
                    _ = v48
                    _ = v50
                }
                fmt.Println(v11, fi30)
                _ = fi30
            }
        }
        for fi51 := 0; fi51 < 2; fi51++ {
            v52 := GenS5{F0: 830, F1: v5, F2: 458}
            v53 := GenS6{F0: v1, F1: v3}
            _ = fi51
            _ = v52
            _ = v53
        }
        if (2147483647 > 734) {
            v54 := GenS2{F0: 148, F1: (v1 - 42)}
            var v55 float64 = (1.5 - -2.25)
            _ = v54
            _ = v55
        } else {
            v6 = (v6 >> ((1) & 63))
        }
        _ = fi23
    }
    var v56 bool = (!true)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v6
}
