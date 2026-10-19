package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8; F1 uint; F2 uint32 }
type GenS4 struct { F0 int; F1 uint32; F2 uint8 }
func genFn0(p0 uint, p1 uint, p2 int16) bool {
    var v4 uint16 = 84
    switch 0 {
    case -353:
        var v5 float32 = 1e10
        _ = v5
    case 32767:
        var v6 uint16 = (v4 * (^(v4)))
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return false
}

func main() {
    var v1 int8 = 0
    v2 := make(map[uint32]int)
    var v3 uint64 = 2
    var v7 int = 2147483647
    var v8 string = (((("foo bar" + "a") + ("a" + "foo bar")) + ((("a" + "foo bar") + ("" + ("foo bar" + "\n\t"))) + "a")) + ((("" + ("hello" + "hello")) + "a") + "foo bar"))
    v8 = v8
    if false {
        var v9 uint16 = 661
        _ = v9
    } else {
        fmt.Println(v8)
        var v10 string = v8
        _ = v10
    }
    v11 := []int{(v7 << ((2) & 63))}
    v12 := GenS1{F0: true, F1: -9223372036854775808}
    v2[42] = (v7 / (((v7 / ((127) | 1))) | 1))
    var v13 int32 = -2147483648
    v8 = v8
    for fi14 := 0; fi14 < 15; fi14++ {
        v15 := GenS1{F0: false, F1: -968}
        _ = fi14
        _ = v15
    }
    if ((65535 < 2147483647) || (!(!true))) {
        v3 += (v3 & v3)
        v16 := GenS1{F0: (-1 >= 32767), F1: 9223372036854775807}
        _ = v16
    } else {
        v11 = append(v11, 32768)
        v1 += v1
    }
    if ((0 > 240) || (!false)) {
        var v17 float32 = -2.25
        var v18 float32 = v17
        _ = v17
        _ = v18
    } else {
        fmt.Println(v3, v7, v13)
        var v19 string = v8
        v20 := v2[65535]
        _ = v19
        _ = v20
    }
    if true {
        for fi21 := 0; fi21 < 12; fi21++ {
            if ((!((((v8 + v8) + "hello") + (((("hello" + "hello") + "a") + v8) + v8)) > v8)) || (!(float64(((3.14159 - 0.0) - (0.1 - (0.1 + 0.0)))) != (1.5 * ((float64(float32(0.0)) + (float64(-2.25) * ((1e10 - 1e10) * (3.14159 * 1.5)))) - ((0.1 - -2.25) + (0.0 - (0.1 * 0.1)))))))) {
                for fi22 := 0; fi22 < 15; fi22++ {
                    fmt.Println(v1, v1)
                    if false {
                        fmt.Println(v3)
                    } else {
                        v13 -= v13
                        if false {
                            if true {
                                fmt.Println(fi22, v3)
                                var v23 float32 = -2.25
                                _ = v23
                            } else {
                                v24 := GenS2{F0: 65535}
                                _ = v24
                            }
                            if len(v11) >= 2 {
                                v11 = v11[1:len(v11)]
                            }
                        }
                        var v25 int = fi22
                        _ = v25
                    }
                    v13 -= v13
                    _ = fi22
                }
                fmt.Println(fi21, v7, v13)
                var v26 int16 = 32767
                _ = v26
            }
            _ = fi21
        }
        v8 = "\n\t"
    } else {
        v7 += (v7 & int(v3))
    }
    fmt.Println(v3, v7, v8)
    v1 = 127
    for fi27 := 0; fi27 < 11; fi27++ {
        var v28 uint = 65535
        _ = fi27
        _ = v28
    }
    for fi29 := 0; fi29 < 9; fi29++ {
        var v30 string = v8
        if true {
            v31 := GenS1{F0: (false || false), F1: int64(v13)}
            _ = v31
        } else {
            if ((("\n\t" + "foo bar") <= v30) && (false || true)) {
                var v32 float32 = 3.14159
                if false {
                    v33 := GenS3{F0: (v1 ^ 42), F1: 18446744073709551615, F2: 885}
                    v1 = (v1 % (((v1 & v1)) | 1))
                    switch "a" {
                    case "":
                        if true {
                            v34 := GenS2{F0: 4294967295}
                            v1 -= v1
                            fmt.Println(v3, v7, v8)
                            _ = v34
                        }
                    default:
                        v13 = 1
                    }
                    _ = v33
                } else {
                    fmt.Println(v3, v3)
                    var v35 uint32 = 1
                    for fi36 := 0; fi36 < 5; fi36++ {
                        if (-445 != 517) {
                            v37 := GenS4{F0: v7, F1: (v35 + v35), F2: 224}
                            _ = v37
                        }
                        _ = fi36
                    }
                    _ = v35
                }
                var v38 uint8 = 128
                _ = v32
                _ = v38
            } else {
                if true {
                    switch v30 {
                    case "a":
                        fmt.Println(v1)
                    }
                }
                fmt.Println(v30, v7)
                if ((("foo bar" == "foo bar") && true) && false) {
                    if ((true || true) && (true || (((v1 << ((255) & 63)) != v1) && (-1 != -1)))) {
                        for fi39 := 0; fi39 < 17; fi39++ {
                            var v40 uint32 = 2147483648
                            _ = fi39
                            _ = v40
                        }
                    }
                } else {
                    fmt.Println(v7, fi29, v1)
                    v3 = (v3 % (((v3 * 1)) | 1))
                }
            }
        }
        _ = fi29
        _ = v30
    }
    switch v7 {
    case 128:
        v2[997] = v7
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
