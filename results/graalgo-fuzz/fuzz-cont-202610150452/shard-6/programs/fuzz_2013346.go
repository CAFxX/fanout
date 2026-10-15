package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 uint; F1 int16; F2 uint }
type GenS4 struct { F0 uint64; F1 int64; F2 uint }
func genFn0(p0 uint8) int {
    v6 := GenS1{F0: 0, F1: 729, F2: 2}
    p0 = p0
    var v7 int32 = 42
    v7 -= (^(v7))
    _ = p0
    _ = v6
    _ = v7
    return -278
}
func genFn1(p0 uint16, p1 int64, p2 string) uint64 {
    switch -129 {
    case 42:
        p2 = ("a" + "")
    case 128:
        var v8 string = p2
        _ = v8
    case 581:
        p1 += (p1 * -9223372036854775808)
    }
    v9 := GenS1{F0: -1, F1: 32767, F2: 18446744073709551615}
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 4294967295
}

func main() {
    var v1 uint16 = 1
    v2 := []int{-1}
    var v3 int32 = -835
    var v4 uint16 = v1
    var v5 string = "foo bar"
    for fi10 := 0; fi10 < 20; fi10++ {
        var v11 bool = false
        if len(v2) > 0 {
            v2[0] = 937
        }
        fmt.Println(v1)
        _ = fi10
        _ = v11
    }
    switch 2147483647 {
    case -1:
        var v12 int32 = v3
        _ = v12
    }
    for fi13 := 0; fi13 < 20; fi13++ {
        v2 = append(v2, 32768)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if true {
            for fi14 := 0; fi14 < 4; fi14++ {
                if false {
                    v15 := GenS2{F0: 82, F1: -1}
                    _ = v15
                } else {
                    _ = len(v2)
                    fmt.Println(v4, v4, v3)
                    if false {
                        fmt.Println(v4)
                        v16 := GenS2{F0: 1, F1: -1}
                        if len(v2) > 0 {
                            v2[0] = fi13
                        }
                        v17 := GenS1{F0: int32(fi13), F1: 32767, F2: 2}
                        _ = v16
                        _ = v17
                    } else {
                        if true {
                            fmt.Println(v1, v4)
                        } else {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = len(v2)
                            v1 = v1
                            v18 := GenS3{F0: 2, F1: 32767, F2: 807}
                            _ = v18
                        }
                        fmt.Println(v5, fi13)
                    }
                }
                v2 = append(v2, fi13)
                _ = fi14
            }
            var v19 uint = 255
            v2 = append(v2, fi13)
            _ = v19
        }
        _ = fi13
    }
    if ((v3 / ((568) | 1)) <= 0) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    if false {
        v5 = ""
        v20 := GenS3{F0: 255, F1: 669, F2: 42}
        var v21 bool = ((0 < 2) && true)
        _ = v20
        _ = v21
    } else {
        switch v5 {
        case "":
            for fi22 := 0; fi22 < 17; fi22++ {
                if false {
                    _ = len(v2)
                    v4 -= 2
                    for fi23 := 0; fi23 < 11; fi23++ {
                        v24 := GenS1{F0: v3, F1: -536, F2: 1}
                        _ = fi23
                        _ = v24
                    }
                } else {
                    v1 = (v1 - 0)
                    var v25 int32 = 1
                    _ = v25
                }
                _ = fi22
            }
        case "foo bar":
            fmt.Println(v5, v3)
        case "hello":
            var v26 bool = (!(!((128 == 62) && ((float32(0.0) + 1.5) <= (float32((1e10 * -2.25)) / 1.0)))))
            _ = v26
        }
    }
    v27 := GenS4{F0: 1, F1: 1, F2: 435}
    if false {
        if false {
            var v28 int8 = 100
            var v29 uint32 = 2147483648
            _ = v28
            _ = v29
        } else {
            var v30 uint8 = 42
            _ = v30
        }
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        switch v5 {
        case "a":
            v3 = v3
        case "hello":
            v4 -= v4
        default:
            v1 = (v1 * (v1 ^ (v4 ^ (v1 - (v1 & (v1 ^ v4))))))
        }
    }
    v4 -= v4
    for fi31 := 0; fi31 < 16; fi31++ {
        var v32 float32 = 3.14159
        _ = fi31
        _ = v32
    }
    var v33 string = "a"
    var v34 int32 = v3
    var v35 uint16 = (v1 * (^(v1)))
    if false {
        if len(v2) > 0 {
            v2[0] = -9223372036854775808
        }
    } else {
        v36 := GenS3{F0: 255, F1: 1, F2: 18446744073709551615}
        var v37 int16 = 42
        _ = v36
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
