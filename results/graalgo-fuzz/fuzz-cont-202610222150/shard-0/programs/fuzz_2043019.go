package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 uint8 }
func genFn0(p0 int, p1 uint32, p2 int8) int8 {
    fmt.Println(p1, p1, p2)
    var v4 uint64 = 1
    v5 := []int{}
    var v6 uint = 18446744073709551615
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return p2
}

func main() {
    var v1 string = "a"
    var v2 float64 = 1e10
    v3 := []uint32{872, 1, 2, 4294967295}
    var v7 string = ("a" + (v1 + ""))
    v8 := GenS1{F0: 1, F1: -721, F2: 2}
    v7 = (v7 + (v1 + ("" + (v1 + (v7 + (v1 + v7))))))
    v7 = v7
    var v9 uint32 = 0
    v1 = (v1 + ((v7 + "a") + v1))
    v9 -= 2147483647
    for fi10 := 0; fi10 < 7; fi10++ {
        if false {
            v1 = ""
            switch fi10 {
            case 128:
                var v11 uint8 = 42
                _ = v11
            case -2147483648:
                v12 := GenS1{F0: 255, F1: -9223372036854775808, F2: 127}
                _ = v12
            }
            v13 := GenS1{F0: 1, F1: -1, F2: 2}
            _ = v13
        }
        _ = fi10
    }
    v3 = append(v3, 0)
    if true {
        switch (v7 + ((v7 + v1) + ("hello" + (v1 + "\n\t")))) {
        case "foo bar":
            if false {
                var v14 uint64 = 2
                switch 4294967295 {
                case -199:
                    var v15 bool = true
                    _ = v15
                default:
                    v7 = v1
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v14
            }
        case "a":
            switch 167 {
            case 32768:
                var v16 int32 = 42
                _ = v16
            case 32767:
                v3 = append(v3, (v9 ^ (^(v9))))
            default:
                v17 := GenS1{F0: 1, F1: -9223372036854775808, F2: 113}
                _ = v17
            }
        case "":
            v9 += (v9 / ((376) | 1))
        default:
            v7 = ((v1 + "hello") + v7)
        }
        if ((18446744073709551615 == 748) || (!true)) {
            fmt.Println(v2, v7)
        } else {
            switch -1 {
            case 128:
                if (399 < -1) {
                    for fi18 := 0; fi18 < 20; fi18++ {
                        _ = len(v3)
                        var v19 int32 = 2147483647
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi18
                        _ = v19
                    }
                    for fi20 := 0; fi20 < 12; fi20++ {
                        for fi21 := 0; fi21 < 15; fi21++ {
                            v7 = v1
                            v9 = 42
                            if ((^(fi21)) < fi20) {
                                v9 = v9
                            } else {
                                v1 = "a"
                                v2 = (v2 / -2.25)
                                if ((-2147483648 != -587) || true) {
                                    var v22 uint64 = 18446744073709551615
                                    fmt.Println(fi20)
                                    _ = v22
                                } else {
                                    v3 = append(v3, 2147483647)
                                    var v23 int32 = -965
                                    var v24 uint32 = (v9 - v9)
                                    var v25 int16 = 108
                                    _ = v23
                                    _ = v24
                                    _ = v25
                                }
                                fmt.Println(fi21, fi21, fi21)
                            }
                            var v26 int16 = -540
                            _ = fi21
                            _ = v26
                        }
                        v1 = ((v1 + ("foo bar" + "a")) + v7)
                        fmt.Println(v9, fi20, v2)
                        _ = fi20
                    }
                    v1 = ((("" + (("hello" + v1) + (v7 + v7))) + "") + (v1 + v1))
                } else {
                    fmt.Println(v9, v2)
                    var v27 uint32 = v9
                    v27 = v27
                    _ = v27
                }
            case 32767:
                v28 := GenS1{F0: 86, F1: int64(v9), F2: 66}
                _ = v28
            default:
                for fi29 := 0; fi29 < 18; fi29++ {
                    var v30 float32 = (((float32(v2) * (1e10 - 0.1)) * float32(v2)) * (0.0 - 0.0))
                    v7 = (v1 + "")
                    for fi31 := 0; fi31 < 6; fi31++ {
                        v9 = 2
                        v1 = "\n\t"
                        switch (-(fi31)) {
                        case 9223372036854775807:
                            fmt.Println(v7, v30)
                        case -1:
                            v32 := GenS1{F0: 42, F1: 1, F2: 2}
                            _ = v32
                        default:
                            v1 = "hello"
                        }
                        _ = fi31
                    }
                    _ = fi29
                    _ = v30
                }
            }
            fmt.Println(v7, v1, v7)
            v2 = v2
        }
        var v33 string = v7
        _ = v33
    }
    var v34 uint32 = 464
    v34 += 65535
    var v35 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v7
    _ = v8
    _ = v9
}
