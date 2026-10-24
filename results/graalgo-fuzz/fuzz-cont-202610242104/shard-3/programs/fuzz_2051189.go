package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16; F1 int32 }
func genFn0(p0 int16, p1 int16, p2 int) uint32 {
    var v7 uint16 = 54
    if (false || (-277 > -1)) {
        var v8 bool = ((-2.25 == (1e10 * (0.0 + float64(1.5)))) && (0 > 42))
        var v9 uint16 = (v7 * (v7 | (v7 >> ((1) & 63))))
        v10 := GenS1{F0: int32(v9), F1: v8, F2: p1}
        _ = v10
        _ = v8
        _ = v9
    }
    p1 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 2
}
func genFn1(p0 uint64, p1 string) uint16 {
    var v11 int = 1
    switch (-(v11)) {
    case 9223372036854775807:
        var v12 int = (-(v11))
        _ = v12
    case 1:
        v13 := GenS2{F0: 319}
        _ = v13
    }
    switch p1 {
    case "a":
        for fi14 := 0; fi14 < 3; fi14++ {
            switch (^(v11)) {
            case -417:
                v15 := []int{(-(v11)), (-(fi14)), (v11 * -9223372036854775808)}
                _ = v15
            case -949:
                v16 := GenS1{F0: 2147483647, F1: (true || true), F2: 32767}
                _ = v16
            case 32767:
                p1 = ("" + (p1 + ((p1 + "a") + "\n\t")))
            }
            _ = fi14
        }
    case "\n\t":
        var v17 int16 = 0
        _ = v17
    case "hello":
        var v18 int8 = -123
        _ = v18
    default:
        fmt.Println(p1)
    }
    var v19 bool = (false && ((v11 < (v11 / ((-9223372036854775808) | 1))) || false))
    v11 -= 0
    _ = p0
    _ = p1
    _ = v11
    _ = v19
    return 32767
}

func main() {
    var v1 float32 = -2.25
    var v2 uint64 = 1
    var v3 float64 = 1e10
    v4 := make(map[uint32]string)
    v5 := []string{("foo bar" + ""), "hello", (("" + "foo bar") + "foo bar"), "hello"}
    v6 := make(map[int]int)
    v3 = v3
    for fi20 := 0; fi20 < 8; fi20++ {
        var v21 string = "a"
        switch v21 {
        case "":
            v22 := GenS1{F0: 1, F1: false, F2: 32767}
            _ = v22
        case "\n\t":
            v5 = append(v5, (v21 + ((v21 + "foo bar") + "\n\t")))
        case "hello":
            if true {
                for fi23 := 0; fi23 < 8; fi23++ {
                    if (917 == 1) {
                        if (!false) {
                            if len(v5) > 0 {
                                v5[0] = v21
                            }
                            var v24 int32 = -2147483648
                            fmt.Println(v21, fi20, v1)
                            var v25 float32 = 0.1
                            var v26 int32 = v24
                            var v27 uint64 = (v2 + 613)
                            _ = v24
                            _ = v25
                            _ = v26
                            _ = v27
                        }
                    } else {
                        var v28 int = (fi20 / (((fi20 << ((1) & 63))) | 1))
                        _ = v28
                    }
                    _ = fi23
                }
                fmt.Println(v1, fi20, v3)
            } else {
                for fi29 := 0; fi29 < 10; fi29++ {
                    v1 = v1
                    _ = fi29
                }
            }
        }
        _ = fi20
        _ = v21
    }
    if true {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v1 = (0.1 / (0.0 - float32((1e10 - v3))))
    } else {
        if true {
            v2 += v2
            v1 = 1e10
        }
    }
    var v30 uint64 = v2
    if false {
        _ = len(v5)
        fmt.Println(v30, v1)
    }
    for fi31 := 0; fi31 < 11; fi31++ {
        v32 := GenS2{F0: 42}
        var v33 string = ("foo bar" + "")
        fmt.Println(v30, v3, fi31)
        _ = fi31
        _ = v32
        _ = v33
    }
    v2 += (v2 - 18446744073709551615)
    delete(v4, 65535)
    var v34 uint = 18446744073709551615
    var v35 float32 = (1.5 / 1.0)
    v6[-1] = 42
    if true {
        v2 = (v30 << (((v34 + 4294967295)) & 63))
        for fi36 := 0; fi36 < 1; fi36++ {
            var v37 int16 = -1
            var v38 int = 32767
            if true {
                fmt.Println(v3, fi36)
                v5 = append(v5, "hello")
                for fi39 := 0; fi39 < 8; fi39++ {
                    if true {
                        var v40 uint8 = 2
                        _ = v40
                    } else {
                        v41 := GenS3{F0: (v37 * (v37 | v37)), F1: 467}
                        if false {
                            if true {
                                switch 42 {
                                case -595:
                                    fmt.Println(v3)
                                default:
                                    v34 = (v34 ^ v34)
                                }
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                            if true {
                                var v42 int32 = -772
                                _ = v42
                            }
                        } else {
                            var v43 uint = (v34 / (((v34 / ((v34) | 1))) | 1))
                            var v44 uint = v43
                            _ = v43
                            _ = v44
                        }
                        v1 += float32(v3)
                        _ = v41
                    }
                    _ = fi39
                }
            } else {
                v35 -= 0.0
            }
            _ = fi36
            _ = v37
            _ = v38
        }
        v1 = (3.14159 - 1.5)
    }
    v30 = (^(v2))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
