package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int8 }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 int64; F1 int32; F2 uint32 }
func genFn0() int64 {
    var v7 uint = 18446744073709551615
    v7 = (v7 % ((42) | 1))
    v7 += 42
    _ = v7
    return 613
}
func genFn1(p0 bool, p1 uint64) int8 {
    var v8 uint64 = 330
    if false {
        v8 = (^(p1))
        for fi9 := 0; fi9 < 14; fi9++ {
            for fi10 := 0; fi10 < 16; fi10++ {
                if p0 {
                    v11 := []int{}
                    p1 += v8
                    if false {
                        p1 = 0
                    } else {
                        var v12 int64 = -703
                        switch 128 {
                        case 127:
                            v13 := GenS1{F0: 2, F1: -1, F2: 100}
                            _ = v13
                        case -751:
                            fmt.Println(fi9)
                        case -2147483648:
                            var v14 uint32 = 836
                            _ = v14
                        default:
                            fmt.Println(v12, fi9)
                        }
                        _ = v12
                    }
                    _ = v11
                } else {
                    v8 = 1
                }
                fmt.Println(fi10, fi10)
                var v15 uint64 = (v8 & (v8 << ((58) & 63)))
                _ = fi10
                _ = v15
            }
            fmt.Println(p0)
            v16 := []int{fi9}
            _ = fi9
            _ = v16
        }
    } else {
        for fi17 := 0; fi17 < 6; fi17++ {
            var v18 int = (fi17 - fi17)
            var v19 uint = 1
            _ = fi17
            _ = v18
            _ = v19
        }
    }
    _ = p0
    _ = p1
    _ = v8
    return 42
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]uint32)
    var v3 uint64 = 18446744073709551615
    var v4 uint32 = 812
    v5 := []string{}
    v6 := make(map[string]int)
    var v20 uint64 = v3
    v5 = append(v5, "hello")
    if true {
        if len(v5) > 0 {
            v5[0] = "\n\t"
        }
        var v21 float32 = ((1e10 * -2.25) * 1e10)
        v22 := GenS2{F0: (v20 - (v3 >> ((42) & 63))), F1: -101}
        _ = v21
        _ = v22
    } else {
        fmt.Println(v20, v20)
        if len(v5) > 0 {
            v5[0] = ("foo bar" + "hello")
        }
        v1[127] = true
    }
    v4 += v4
    _ = len(v5)
    v5 = append(v5, "hello")
    fmt.Println(v4, v4, v3)
    v23 := GenS1{F0: 0, F1: 1, F2: 44}
    var v24 int32 = 2147483647
    if len(v5) > 0 {
        v5[0] = "hello"
    }
    switch -9223372036854775808 {
    case 541:
        for fi25 := 0; fi25 < 11; fi25++ {
            var v26 uint32 = 4294967295
            fmt.Println(v3, v20)
            _ = fi25
            _ = v26
        }
    }
    var v27 uint = 492
    if (1e10 >= -2.25) {
        v20 -= v3
        for fi28 := 0; fi28 < 4; fi28++ {
            var v29 uint16 = 355
            v30 := GenS3{F0: genFn0(), F1: v24, F2: v4}
            _ = fi28
            _ = v29
            _ = v30
        }
        if ((("hello" + "hello") + "foo bar") >= "a") {
            if (false && ((127 < -128) && (float64(((3.14159 * float32(float64(-2.25))) - -2.25)) == float64(0.1)))) {
                switch -296 {
                case -306:
                    for fi31 := 0; fi31 < 6; fi31++ {
                        if (!true) {
                            var v32 uint32 = 4294967295
                            _ = v32
                        }
                        _ = fi31
                    }
                case -9223372036854775808:
                    for fi33 := 0; fi33 < 11; fi33++ {
                        for fi34 := 0; fi34 < 18; fi34++ {
                            var v35 string = ("a" + ("foo bar" + "a"))
                            fmt.Println(v3, fi34, v27)
                            _ = fi34
                            _ = v35
                        }
                        fmt.Println(v27)
                        if len(v5) > 0 {
                            v5[0] = "hello"
                        }
                        _ = fi33
                    }
                }
                for fi36 := 0; fi36 < 14; fi36++ {
                    fmt.Println(v4, v4, v20)
                    v20 -= 4294967295
                    _ = fi36
                }
            } else {
                var v37 string = "hello"
                _ = v37
            }
        } else {
            if len(v5) > 0 {
                v5[0] = (("\n\t" + ((("\n\t" + "foo bar") + "a") + "")) + ("\n\t" + ""))
            }
            v4 += (v4 & (v4 + (v4 % ((v4) | 1))))
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
