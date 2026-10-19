package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 uint64; F2 uint }
func genFn0() uint16 {
    var v5 uint8 = 217
    v5 += v5
    v5 = 127
    if (184 != 1) {
        var v6 string = ""
        _ = v6
    }
    switch "a" {
    case "hello":
        for fi7 := 0; fi7 < 7; fi7++ {
            for fi8 := 0; fi8 < 5; fi8++ {
                v5 = v5
                if (true && (!(1 >= 745))) {
                    switch "\n\t" {
                    case "a":
                        v5 = 197
                    case "":
                        var v9 float64 = 1e10
                        _ = v9
                    case "foo bar":
                        v5 -= 2
                    }
                    var v10 int = fi8
                    fmt.Println(fi7)
                    fmt.Println(fi7, fi7, v5)
                    if (2 <= 42) {
                        v11 := GenS1{F0: 433, F1: fi8}
                        var v12 int = 1
                        if ((3.14159 == -2.25) && false) {
                            v10 = v10
                            var v13 int16 = -21
                            v5 += 0
                            var v14 int8 = -1
                            _ = v13
                            _ = v14
                        } else {
                            var v15 int16 = 403
                            v16 := GenS2{F0: 255}
                            _ = v15
                            _ = v16
                        }
                        _ = v11
                        _ = v12
                    }
                    _ = v10
                } else {
                    v17 := GenS3{F0: 42, F1: 0, F2: 1}
                    _ = v17
                }
                _ = fi8
            }
            for fi18 := 0; fi18 < 11; fi18++ {
                fmt.Println(v5, fi18, fi18)
                fmt.Println(v5)
                v5 = v5
                v5 -= 80
                var v19 int16 = 882
                _ = fi18
                _ = v19
            }
            if false {
                fmt.Println(v5, v5)
            } else {
                for fi20 := 0; fi20 < 9; fi20++ {
                    if false {
                        var v21 string = "a"
                        _ = v21
                    } else {
                        v5 -= v5
                        var v22 uint16 = 2
                        _ = v22
                    }
                    _ = fi20
                }
            }
            _ = fi7
        }
    default:
        for fi23 := 0; fi23 < 1; fi23++ {
            if false {
                v5 = v5
            }
            var v24 uint = 2
            _ = fi23
            _ = v24
        }
    }
    _ = v5
    return 65535
}
func genFn1() int16 {
    var v25 uint32 = 1
    v26 := []int{32767, 32768}
    if len(v26) > 0 {
        v26[0] = 32768
    }
    if (2 != 4294967295) {
        fmt.Println(v25)
        _ = len(v26)
        var v27 int64 = -1
        _ = v27
    }
    _ = v25
    _ = v26
    return -1
}

func main() {
    var v1 int16 = 42
    var v2 uint = 4294967295
    v3 := make(map[string]string)
    v4 := []string{"", "", ("\n\t" + "foo bar"), "a"}
    v1 += (^(v1))
    var v28 int64 = int64(v2)
    v28 += (v28 % ((-877) | 1))
    v2 = (v2 / ((v2) | 1))
    var v29 int = 4294967295
    fmt.Println(v28, v28, v1)
    v1 = (v1 & (v1 >> ((v2) & 63)))
    var v30 bool = ((((((0 > 18446744073709551615) && false) && false) && true) && true) || (true || true))
    v31 := GenS2{F0: v2}
    for fi32 := 0; fi32 < 13; fi32++ {
        if len(v4) > 0 {
            v4[0] = ""
        }
        var v33 uint = (v2 % ((v2) | 1))
        v30 = (32767 > 4294967295)
        _ = fi32
        _ = v33
    }
    var v34 string = "a"
    v35 := GenS1{F0: v1, F1: (v29 / ((4294967295) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v35
    _ = v4
}
