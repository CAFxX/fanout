package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 int8 }
type GenS2 struct { F0 int; F1 uint8 }
func genFn0(p0 int16, p1 string, p2 uint8) uint32 {
    if true {
        p1 = p1
        if true {
            v6 := GenS1{F0: -752, F1: (^(p2)), F2: -128}
            for fi7 := 0; fi7 < 18; fi7++ {
                if true {
                    if true {
                        p1 = (p1 + p1)
                    }
                    p0 = 42
                }
                _ = fi7
            }
            _ = v6
        } else {
            v8 := []int{}
            _ = v8
        }
        p0 -= p0
    } else {
        switch 0 {
        case 0:
            if true {
                var v9 float32 = (3.14159 + 0.0)
                _ = v9
            }
        }
        v10 := []int{42, 779, 938}
        _ = v10
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        var v12 float64 = float64(0.0)
        if (!(!(0 == -1))) {
            v13 := GenS1{F0: 1, F1: (p2 & (p2 | 0)), F2: -5}
            _ = v13
        } else {
            var v14 int = fi11
            var v15 uint8 = (p2 - p2)
            _ = v14
            _ = v15
        }
        _ = fi11
        _ = v12
    }
    fmt.Println(p2, p1, p2)
    _ = p0
    _ = p1
    _ = p2
    return 2
}

func main() {
    var v1 uint = 1
    v2 := []int{-607}
    var v3 int = 32768
    var v4 int32 = 791
    var v5 float64 = 1.5
    if ((0.0 * (0.1 * -2.25)) <= 1e10) {
        v4 = v4
        if true {
            v5 -= v5
            if (v3 > 964) {
                var v16 int = 32767
                if ((((float64((-2.25 / float32(v5))) / v5) > 3.14159) && ((("hello" >= (("" + "") + ("" + "foo bar"))) || (v4 != v4)) && (4294967295 >= 18446744073709551615))) || ((!true) && (!(v5 >= ((float64(0.1) - v5) + 0.0))))) {
                    if len(v2) > 0 {
                        v2[0] = (v3 << ((v1) & 63))
                    }
                    v2 = append(v2, -9223372036854775808)
                    var v17 string = ("" + "\n\t")
                    _ = v17
                } else {
                    var v18 float64 = (((1.5 - v5) / 0.1) * (v5 * 1e10))
                    v4 = (v4 - (v4 ^ v4))
                    switch (v16 & 128) {
                    case 127:
                        fmt.Println(v4, v18, v4)
                    case -774:
                        var v19 uint16 = 65535
                        _ = v19
                    case 0:
                        v3 += (v16 % (((v3 << ((v1) & 63))) | 1))
                    default:
                        var v20 uint16 = 2
                        _ = v20
                    }
                    _ = v18
                }
                if false {
                    for fi21 := 0; fi21 < 11; fi21++ {
                        var v22 uint32 = 4294967295
                        v22 += genFn0(32767, "", 2)
                        for fi23 := 0; fi23 < 19; fi23++ {
                            var v24 string = "hello"
                            if len(v2) > 0 {
                                v2[0] = (^(fi23))
                            }
                            var v25 int16 = 32767
                            _ = fi23
                            _ = v24
                            _ = v25
                        }
                        _ = fi21
                        _ = v22
                    }
                } else {
                    switch ("a" + ("hello" + (("a" + "hello") + (("foo bar" + "a") + ("" + "hello"))))) {
                    case "a":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case "hello":
                        if true {
                            var v26 uint = v1
                            v16 = (v16 / ((v16) | 1))
                            fmt.Println(v26)
                            _ = v26
                        }
                    default:
                        v27 := GenS1{F0: -9223372036854775808, F1: 1, F2: -58}
                        _ = v27
                    }
                    v3 += (v16 / (((v16 / (((v3 ^ 2147483647)) | 1))) | 1))
                }
                _ = v16
            } else {
                for fi28 := 0; fi28 < 19; fi28++ {
                    v4 = -2147483648
                    v29 := GenS2{F0: (-(v3)), F1: 25}
                    _ = fi28
                    _ = v29
                }
                v1 += v1
                var v30 uint = v1
                _ = v30
            }
        } else {
            v4 -= v4
        }
        fmt.Println(v3)
    }
    var v31 float64 = float64((1e10 * 3.14159))
    v31 = 3.14159
    var v32 int = (v3 / ((576) | 1))
    var v33 int8 = 100
    v4 = 0
    var v34 int8 = v33
    var v35 float32 = 1.5
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v5, v31)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
