package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint32; F1 uint32 }
type GenS5 struct { F0 int64; F1 uint; F2 uint16 }
type GenS6 struct { F0 uint64; F1 string; F2 int }
func genFn0(p0 uint) uint8 {
    switch (("hello" + "foo bar") + "foo bar") {
    case "hello":
        p0 = (p0 & (p0 | (p0 + (p0 - 213))))
    case "a":
        v4 := GenS1{F0: -128, F1: 1, F2: (p0 / ((0) | 1))}
        _ = v4
    default:
        var v5 uint = 18446744073709551615
        _ = v5
    }
    var v6 int = 4294967295
    v7 := GenS2{F0: 922}
    switch (-(v6)) {
    case 1:
        v6 -= v6
    }
    v6 += v6
    _ = p0
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 uint32 = 173
    v2 := make(map[string]int)
    var v3 int64 = 9223372036854775807
    var v8 uint8 = 0
    for fi9 := 0; fi9 < 16; fi9++ {
        var v10 uint16 = 0
        v8 = genFn0(65535)
        fmt.Println(v10, v1)
        _ = fi9
        _ = v10
    }
    v11 := GenS3{F0: 42}
    var v12 uint8 = (^(v8))
    var v13 string = "hello"
    fmt.Println(v8, v13, v1)
    for fi14 := 0; fi14 < 15; fi14++ {
        v15 := GenS3{F0: (fi14 << ((255) & 63))}
        if (((^(v8)) != genFn0(18446744073709551615)) && true) {
            for fi16 := 0; fi16 < 5; fi16++ {
                var v17 uint8 = 127
                v18 := GenS3{F0: (fi16 >> ((1) & 63))}
                var v19 float64 = 0.0
                v20 := GenS4{F0: 275, F1: v1}
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
                _ = v20
            }
            for fi21 := 0; fi21 < 13; fi21++ {
                fmt.Println(fi14, fi21)
                v3 += (v3 - v3)
                _ = fi21
            }
        } else {
            for fi22 := 0; fi22 < 10; fi22++ {
                switch v13 {
                case "a":
                    var v23 uint32 = 434
                    _ = v23
                default:
                    _ = len(v2)
                }
                switch int(v1) {
                case 128:
                    v13 = ""
                    v3 = (v3 >> ((1) & 63))
                    switch (v13 + (v13 + v13)) {
                    case "hello":
                        if ((v1 % (((^(v1))) | 1)) < (v1 >> ((4294967295) & 63))) {
                            v3 += v3
                            if false {
                                if false {
                                    fmt.Println(fi22, v8, fi14)
                                }
                            }
                        }
                        v24 := GenS2{F0: (v1 % ((2147483647) | 1))}
                        switch ((v13 + (v13 + "foo bar")) + (v13 + v13)) {
                        case "":
                            v8 = 255
                        case "\n\t":
                            if (128 < genFn0(255)) {
                                var v25 int8 = -1
                                var v26 bool = false
                                _ = v25
                                _ = v26
                            } else {
                                if (float32(float64(1e10)) != 0.1) {
                                    if false {
                                        v3 = (^(v3))
                                    } else {
                                        var v27 uint = 1
                                        v2[((v13 + v13) + (v13 + "a"))] = (fi14 % ((fi14) | 1))
                                        v28 := []int{fi14}
                                        v12 += uint8(v27)
                                        _ = v27
                                        _ = v28
                                    }
                                } else {
                                    var v29 bool = (127 >= -76)
                                    _ = v29
                                }
                                if false {
                                    v30 := GenS5{F0: -278, F1: uint(v8), F2: uint16(v12)}
                                    var v31 uint64 = 978
                                    _ = v30
                                    _ = v31
                                } else {
                                    v2[(((v13 + (v13 + v13)) + v13) + ("hello" + ((v13 + (v13 + ((v13 + "") + (v13 + v13)))) + ((v13 + (v13 + v13)) + ((v13 + ("" + v13)) + v13)))))] = (fi14 * (-(fi22)))
                                }
                            }
                        default:
                            v8 += (v12 + (^(v8)))
                        }
                        _ = v24
                    case "\n\t":
                        switch 1 {
                        case 32768:
                            v32 := []int{32768, fi14}
                            _ = v32
                        }
                    }
                case 2147483647:
                    switch (v13 + "\n\t") {
                    case "foo bar":
                        v8 += v12
                    case "":
                        var v33 uint64 = 0
                        _ = v33
                    case "a":
                        v8 -= (v12 % ((v8) | 1))
                    default:
                        var v34 bool = true
                        _ = v34
                    }
                case 4294967295:
                    var v35 float64 = 3.14159
                    _ = v35
                default:
                    var v36 int8 = 0
                    _ = v36
                }
                _ = fi22
            }
            var v37 uint16 = 32767
            v37 = (v37 * v37)
            _ = v37
        }
        _ = fi14
        _ = v15
    }
    if true {
        if false {
            for fi38 := 0; fi38 < 10; fi38++ {
                v1 -= (v1 % (((v1 & v1)) | 1))
                for fi39 := 0; fi39 < 13; fi39++ {
                    _ = len(v2)
                    v40 := []int{(^(fi39))}
                    v3 = v3
                    _ = fi39
                    _ = v40
                }
                _ = fi38
            }
            v3 += v3
        } else {
            v41 := v2[""]
            _ = v41
        }
        v13 = v13
    }
    for fi42 := 0; fi42 < 3; fi42++ {
        v43 := GenS6{F0: 270, F1: "\n\t", F2: 0}
        v1 -= v1
        fmt.Println(v1)
        _ = fi42
        _ = v43
    }
    v3 = 0
    var v44 uint8 = (v8 % ((2) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v44:", v44)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v44
    _ = v8
}
