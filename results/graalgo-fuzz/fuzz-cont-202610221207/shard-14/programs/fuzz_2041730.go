package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 uint; F2 int16 }
func genFn0(p0 uint64) int {
    if (21 != 0) {
        v6 := []int{}
        _ = v6
    }
    v7 := GenS1{F0: 555}
    v8 := []int{}
    _ = p0
    _ = v7
    _ = v8
    return 32767
}
func genFn1() int8 {
    switch genFn0(42) {
    case -9223372036854775808:
        if false {
            v9 := []int{}
            v10 := GenS2{F0: 1}
            _ = v10
            _ = v9
        }
    case 9223372036854775807:
        var v11 uint = 18446744073709551615
        _ = v11
    }
    switch genFn0(539) {
    case 128:
        for fi12 := 0; fi12 < 7; fi12++ {
            var v13 uint = 702
            v13 += (v13 >> ((v13) & 63))
            var v14 int32 = -92
            _ = fi12
            _ = v13
            _ = v14
        }
    default:
        var v15 uint32 = 0
        _ = v15
    }
    var v16 bool = true
    v16 = false
    v17 := []int{4294967295, genFn0(1)}
    _ = v16
    _ = v17
    return 0
}

func main() {
    var v1 int64 = 0
    var v2 float32 = 0.1
    var v3 uint64 = 18446744073709551615
    v4 := make(map[string]string)
    v5 := []uint32{65535, 65535}
    _ = len(v5)
    var v18 int16 = -1
    for fi19 := 0; fi19 < 12; fi19++ {
        fmt.Println(v1)
        switch -688 {
        case 1:
            if (3.14159 <= -2.25) {
                for fi20 := 0; fi20 < 18; fi20++ {
                    if (int32(v3) <= -2147483648) {
                        fmt.Println(v1, v1, fi19)
                        var v21 uint = 305
                        delete(v4, "")
                        _ = v21
                    }
                    if false {
                        if (!(803 < 65535)) {
                            fmt.Println(v3, v3)
                            var v22 uint32 = 2147483647
                            _ = v22
                        } else {
                            var v23 uint8 = 133
                            if (!((((0.1 + (1e10 * (0.1 / -2.25))) * 1.5) - -2.25) >= -2.25)) {
                                fmt.Println(v3, v2, fi19)
                                if true {
                                    v23 = 180
                                } else {
                                    var v24 bool = true
                                    if len(v5) > 0 {
                                        v5[0] = 1
                                    }
                                    v2 = ((3.14159 * ((1.5 / v2) - v2)) * ((float32(3.14159) * (v2 / 1e10)) * ((((v2 + ((v2 + v2) - (float32(1e10) * -2.25))) * v2) + v2) / float32(3.14159))))
                                    if v24 {
                                        v25 := GenS3{F0: 1, F1: 1, F2: v18}
                                        _ = v25
                                    } else {
                                        var v26 uint32 = 679
                                        _ = v26
                                    }
                                    _ = v24
                                }
                                delete(v4, ("" + ""))
                            } else {
                                var v27 string = ("hello" + "")
                                if true {
                                    delete(v4, "hello")
                                }
                                if true {
                                    v18 -= v18
                                }
                                _ = v27
                            }
                            _ = v23
                        }
                        v3 = v3
                    }
                    _ = fi20
                }
                v3 = (v3 / ((v3) | 1))
                var v28 int64 = v1
                _ = v28
            } else {
                v5 = append(v5, 1)
            }
        case 695:
            var v29 uint32 = 2147483648
            _ = v29
        case -39:
            v2 = v2
        }
        var v30 uint8 = 2
        _ = fi19
        _ = v30
    }
    if (-2147483648 >= genFn0(722)) {
        for fi31 := 0; fi31 < 14; fi31++ {
            if (((!true) && false) || false) {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            } else {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            _ = fi31
        }
        for fi32 := 0; fi32 < 13; fi32++ {
            v2 += float32(-2.25)
            _ = fi32
        }
        if (false && true) {
            for fi33 := 0; fi33 < 10; fi33++ {
                fmt.Println(v18, v1, v1)
                _ = fi33
            }
            for fi34 := 0; fi34 < 8; fi34++ {
                if false {
                    var v35 int = 4294967295
                    _ = v35
                }
                _ = fi34
            }
            for fi36 := 0; fi36 < 15; fi36++ {
                var v37 int = (fi36 ^ fi36)
                v38 := GenS2{F0: v3}
                _ = fi36
                _ = v37
                _ = v38
            }
        }
    } else {
        for fi39 := 0; fi39 < 4; fi39++ {
            v1 = (v1 + -518)
            v2 = v2
            v2 = (v2 / ((v2 + 1.5) * float32(1.5)))
            _ = fi39
        }
        _ = len(v5)
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    var v40 float32 = (((-2.25 - v2) * v2) + (-2.25 * (v2 / (((v2 - float32((0.1 + 3.14159))) * (v2 * float32(-2.25))) / -2.25))))
    for fi41 := 0; fi41 < 13; fi41++ {
        v5 = append(v5, 312)
        fmt.Println(v18, v3)
        var v42 string = "\n\t"
        _ = fi41
        _ = v42
    }
    _ = len(v4)
    var v43 int16 = v18
    v44 := GenS2{F0: (v3 + v3)}
    v45 := GenS3{F0: 1, F1: 918, F2: (^(v18))}
    v43 = v18
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v43
    _ = v44
    _ = v45
    _ = v5
}
