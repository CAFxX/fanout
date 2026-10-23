package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 uint32 }
type GenS2 struct { F0 int; F1 bool }
func genFn0(p0 string) int {
    fmt.Println(p0)
    var v4 float64 = 0.1
    v5 := []int{462, 2147483647, 32768}
    v4 = 3.14159
    fmt.Println(p0, v4)
    _ = p0
    _ = v4
    _ = v5
    return -9223372036854775808
}
func genFn1() string {
    var v6 uint32 = 65535
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := []int{-2147483648, genFn0(("a" + "foo bar"))}
        for fi9 := 0; fi9 < 15; fi9++ {
            v6 = (v6 >> ((255) & 63))
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    if (4294967295 == 4294967295) {
        v10 := []int{-435}
        _ = v10
    } else {
        v11 := GenS1{F0: 255, F1: 741, F2: (v6 / (((v6 | (v6 ^ v6))) | 1))}
        if (283 > -401) {
            v12 := []int{-1, -2147483648, -9223372036854775808, 42}
            _ = v12
        } else {
            for fi13 := 0; fi13 < 20; fi13++ {
                if (v6 < (v6 | v6)) {
                    if ((-32768 != 0) && true) {
                        if false {
                            for fi14 := 0; fi14 < 1; fi14++ {
                                fmt.Println(v6, fi14)
                                var v15 float32 = (0.0 + 1e10)
                                var v16 uint32 = (v6 | v6)
                                fmt.Println(fi13, fi13, v6)
                                v6 -= v6
                                _ = fi14
                                _ = v15
                                _ = v16
                            }
                            var v17 float32 = 0.0
                            var v18 uint64 = 18446744073709551615
                            _ = v17
                            _ = v18
                        }
                        v6 -= (v6 | v6)
                        if true {
                            v19 := []int{(fi13 * 1), fi13, (fi13 + fi13), 1}
                            _ = v19
                        } else {
                            v6 -= 432
                        }
                    } else {
                        switch (-(fi13)) {
                        case 4294967295:
                            var v20 uint32 = 65535
                            _ = v20
                        case 32767:
                            v6 = v6
                        case -252:
                            var v21 int64 = 9223372036854775807
                            _ = v21
                        default:
                            var v22 string = "\n\t"
                            _ = v22
                        }
                        fmt.Println(v6)
                    }
                    v6 = (v6 - 578)
                    v6 = (^(v6))
                } else {
                    v6 = (v6 / (((v6 | v6)) | 1))
                }
                var v23 uint8 = 127
                v24 := []int{fi13}
                _ = fi13
                _ = v23
                _ = v24
            }
            fmt.Println(v6)
        }
        _ = v11
    }
    _ = v6
    return ""
}

func main() {
    v1 := []string{"foo bar"}
    var v2 int = 2147483647
    var v3 uint16 = 1
    var v25 int32 = 42
    if false {
        v25 -= 1
        var v26 int32 = (v25 + (v25 - (v25 + v25)))
        v3 -= v3
        _ = v26
    } else {
        v3 = v3
        v25 -= (^(v25))
        v1 = append(v1, ("" + "hello"))
    }
    switch int(v25) {
    case 127:
        var v27 float32 = float32(0.1)
        _ = v27
    default:
        for fi28 := 0; fi28 < 7; fi28++ {
            var v29 int64 = 787
            _ = fi28
            _ = v29
        }
    }
    switch -1 {
    case 127:
        v3 = 939
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v30 := GenS1{F0: 83, F1: 42, F2: 4294967295}
    for fi31 := 0; fi31 < 17; fi31++ {
        var v32 int16 = int16(v3)
        _ = fi31
        _ = v32
    }
    var v33 uint32 = 2
    var v34 uint8 = uint8(v3)
    v33 = v33
    v33 -= 42
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    if true {
        v35 := GenS1{F0: (v34 >> ((65535) & 63)), F1: 615, F2: (v33 >> ((uint(v34)) & 63))}
        if (("\n\t" + "") < (("a" + ("" + "a")) + "hello")) {
            fmt.Println(v34)
            v34 = 129
            if false {
                for fi36 := 0; fi36 < 18; fi36++ {
                    var v37 string = "a"
                    v2 -= (v2 >> ((2) & 63))
                    _ = fi36
                    _ = v37
                }
                fmt.Println(v25, v33)
            } else {
                v34 -= v34
                var v38 uint8 = v34
                if false {
                    for fi39 := 0; fi39 < 8; fi39++ {
                        if len(v1) > 0 {
                            v1[0] = ("foo bar" + "a")
                        }
                        v40 := GenS2{F0: fi39, F1: false}
                        _ = len(v1)
                        _ = fi39
                        _ = v40
                    }
                    v34 -= (v38 - 180)
                }
                _ = v38
            }
        }
        _ = v35
    } else {
        v25 -= 104
    }
    if false {
        fmt.Println(v33)
        fmt.Println(v25, v33)
    } else {
        if false {
            v25 = v25
        }
    }
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v30
    _ = v33
    _ = v34
}
