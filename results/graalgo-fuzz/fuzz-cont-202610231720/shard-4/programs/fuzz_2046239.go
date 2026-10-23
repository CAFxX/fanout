package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16; F1 uint16; F2 int32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint32; F1 string; F2 uint8 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int8 }
func genFn0(p0 int16, p1 string) int64 {
    var v4 bool = (false && true)
    var v5 uint8 = 128
    v5 = (^(v5))
    v6 := []int{-803, int(p0), -2147483648, 4294967295}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 int8 = -128
    var v2 uint64 = 42
    var v3 bool = (false || (-110 > 32767))
    v1 = v1
    var v7 string = "foo bar"
    switch -9223372036854775808 {
    case -9223372036854775808:
        v2 -= 4294967295
    case -121:
        v8 := GenS1{F0: 32767, F1: "a"}
        _ = v8
    default:
        v7 = (((v7 + v7) + v7) + v7)
    }
    v2 = v2
    v2 += v2
    for fi9 := 0; fi9 < 6; fi9++ {
        fmt.Println(v1)
        _ = fi9
    }
    v1 = v1
    for fi10 := 0; fi10 < 13; fi10++ {
        var v11 int64 = 0
        for fi12 := 0; fi12 < 8; fi12++ {
            v13 := GenS1{F0: 255, F1: "foo bar"}
            if true {
                if true {
                    if v3 {
                        var v14 float32 = 0.1
                        _ = v14
                    } else {
                        if (v2 == 2) {
                            var v15 float64 = (float64((((3.14159 - 0.1) - 1.5) + 1.5)) - (0.1 - (1.5 * 0.0)))
                            var v16 int64 = (v11 & genFn0(0, ""))
                            _ = v15
                            _ = v16
                        } else {
                            switch (fi12 & -743) {
                            case 1:
                                if v3 {
                                    var v17 int = -1
                                    _ = v17
                                } else {
                                    fmt.Println(v1, fi12)
                                    var v18 int8 = (v1 % ((v1) | 1))
                                    v3 = v3
                                    _ = v18
                                }
                            case 128:
                                v19 := GenS2{F0: v1}
                                _ = v19
                            case -2147483648:
                                v20 := GenS3{F0: "a"}
                                _ = v20
                            }
                            v21 := []int{fi12, fi10, (fi10 - fi10), fi10}
                            fmt.Println(v1, fi10, fi10)
                            v11 += (v11 + (v11 ^ (v11 | v11)))
                            v22 := GenS1{F0: 1, F1: "foo bar"}
                            _ = v21
                            _ = v22
                        }
                    }
                    v23 := GenS1{F0: 527, F1: "a"}
                    v24 := GenS4{F0: 2, F1: 912, F2: 2147483647}
                    _ = v23
                    _ = v24
                }
                v25 := GenS2{F0: v1}
                _ = v25
            }
            v26 := GenS5{F0: 329}
            v27 := GenS6{F0: 848, F1: v7, F2: 1}
            _ = fi12
            _ = v13
            _ = v26
            _ = v27
        }
        for fi28 := 0; fi28 < 13; fi28++ {
            v1 -= (v1 & (^(v1)))
            _ = fi28
        }
        _ = fi10
        _ = v11
    }
    var v29 uint8 = 155
    v30 := []int{747, 128}
    switch v7 {
    case "hello":
        v1 = 127
    case "":
        if len(v30) > 0 {
            v30[0] = 4294967295
        }
    default:
        var v31 float64 = -2.25
        _ = v31
    }
    switch ((((v7 + (v7 + "a")) + v7) + ((v7 + "hello") + "hello")) + v7) {
    case "\n\t":
        if (!(float64(float32(1.5)) == 1.5)) {
            v7 = "foo bar"
        } else {
            if (!(true && false)) {
                var v32 int32 = -784
                var v33 uint32 = 0
                _ = v32
                _ = v33
            }
            if v3 {
                v29 = 2
            } else {
                if len(v30) > 0 {
                    v30[0] = -129
                }
                var v34 string = ((v7 + ("" + v7)) + "foo bar")
                _ = v34
            }
            v35 := GenS5{F0: 672}
            _ = v35
        }
    case "hello":
        v36 := GenS7{F0: v1}
        _ = v36
    case "a":
        v37 := GenS4{F0: 255, F1: 613, F2: 42}
        _ = v37
    default:
        if v3 {
            if (3.14159 != -2.25) {
                for fi38 := 0; fi38 < 11; fi38++ {
                    v39 := GenS3{F0: (v7 + (("a" + (v7 + ("" + "\n\t"))) + v7))}
                    _ = fi38
                    _ = v39
                }
                var v40 float32 = 1e10
                _ = v40
            }
        }
    }
    if (-1 >= 29) {
        for fi41 := 0; fi41 < 4; fi41++ {
            v29 = (v29 << ((uint(v2)) & 63))
            _ = fi41
        }
    }
    var v42 int32 = -2147483648
    if (!false) {
        var v43 int16 = -1
        v44 := GenS2{F0: (v1 * v1)}
        v29 -= v29
        _ = v43
        _ = v44
    } else {
        var v45 int32 = 0
        v46 := GenS2{F0: v1}
        var v47 uint = 65535
        _ = v45
        _ = v46
        _ = v47
    }
    v48 := GenS8{F0: v1}
    v29 = v29
    v2 = 883
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v42:", v42)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v42
    _ = v48
    _ = v7
}
