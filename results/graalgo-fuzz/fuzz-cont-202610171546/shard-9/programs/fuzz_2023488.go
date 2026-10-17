package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool; F1 int64; F2 uint32 }
func genFn0(p0 int16, p1 uint64) string {
    switch -198 {
    case 42:
        v4 := []int{}
        _ = v4
    case 502:
        if false {
            var v5 int = 581
            p0 -= p0
            _ = v5
        }
    case 4294967295:
        p0 = 0
    }
    p0 -= 0
    if (1 <= 905) {
        fmt.Println(p0, p0)
        v6 := GenS1{F0: 65535, F1: 1}
        switch "" {
        case "hello":
            p0 = (p0 * 0)
        case "\n\t":
            p1 = (^(p1))
        default:
            p0 = p0
        }
        _ = v6
    } else {
        p1 = (p1 ^ (p1 >> ((673) & 63)))
        var v7 uint16 = 42
        fmt.Println(p1, p0, p1)
        _ = v7
    }
    _ = p0
    _ = p1
    return ((("a" + "a") + "hello") + (("a" + "\n\t") + ("foo bar" + "")))
}

func main() {
    var v1 string = "hello"
    var v2 int = 32767
    var v3 uint32 = 194
    var v8 float64 = -2.25
    v9 := []int{(v2 >> ((744) & 63)), v2}
    v8 = 0.1
    v3 = 4294967295
    for fi10 := 0; fi10 < 17; fi10++ {
        switch v2 {
        case 127:
            if (!true) {
                if true {
                    v11 := GenS2{F0: -1}
                    fmt.Println(fi10)
                    v1 = (v1 + (("hello" + "hello") + "foo bar"))
                    _ = v11
                } else {
                    for fi12 := 0; fi12 < 16; fi12++ {
                        fmt.Println(v1, v3, v8)
                        var v13 int32 = int32(fi12)
                        fmt.Println(v1, fi12)
                        _ = fi12
                        _ = v13
                    }
                    for fi14 := 0; fi14 < 10; fi14++ {
                        var v15 float32 = 0.0
                        _ = fi14
                        _ = v15
                    }
                }
                if (v1 >= v1) {
                    v16 := GenS2{F0: 1}
                    v3 += v3
                    _ = v16
                }
                var v17 float32 = ((1e10 / 0.1) + 0.1)
                _ = v17
            } else {
                v3 += v3
                for fi18 := 0; fi18 < 11; fi18++ {
                    if false {
                        if true {
                            var v19 uint8 = 218
                            if true {
                                var v20 uint32 = v3
                                var v21 float32 = float32(1e10)
                                v22 := GenS2{F0: 0}
                                _ = v20
                                _ = v21
                                _ = v22
                            } else {
                                fmt.Println(v1)
                                var v23 uint = 42
                                v2 += v2
                                _ = v23
                            }
                            v9 = append(v9, (fi18 + 32767))
                            if false {
                                v19 += v19
                                switch fi18 {
                                case -2147483648:
                                    var v24 int32 = 0
                                    _ = v24
                                default:
                                    v19 += v19
                                }
                                v25 := GenS1{F0: 30, F1: -2147483648}
                                _ = v25
                            }
                            _ = v19
                        } else {
                            v26 := GenS1{F0: v3, F1: 2147483647}
                            _ = v26
                        }
                    } else {
                        v27 := GenS2{F0: -1}
                        _ = v27
                    }
                    var v28 int32 = 1
                    var v29 uint16 = 32767
                    v2 -= 127
                    v9 = append(v9, (v2 - (fi18 & (v2 << ((18446744073709551615) & 63)))))
                    _ = fi18
                    _ = v28
                    _ = v29
                }
            }
        case -9223372036854775808:
            var v30 uint32 = (^(v3))
            _ = v30
        default:
            for fi31 := 0; fi31 < 5; fi31++ {
                switch (fi10 - (^(fi31))) {
                case 2147483647:
                    var v32 bool = (!false)
                    _ = v32
                case -9223372036854775808:
                    switch v1 {
                    case "":
                        v33 := GenS3{F0: 932}
                        _ = v33
                    case "\n\t":
                        v34 := GenS4{F0: true, F1: -839, F2: 4294967295}
                        _ = v34
                    }
                }
                var v35 int64 = 687
                _ = fi31
                _ = v35
            }
        }
        if len(v9) > 0 {
            v9[0] = fi10
        }
        _ = fi10
    }
    fmt.Println(v8, v3)
    for fi36 := 0; fi36 < 7; fi36++ {
        v2 -= (^(fi36))
        var v37 bool = false
        v38 := GenS1{F0: (v3 & 2147483647), F1: -1}
        _ = fi36
        _ = v37
        _ = v38
    }
    fmt.Println(v2, v3, v2)
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
