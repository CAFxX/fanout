package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 int16; F2 uint }
type GenS3 struct { F0 int }
func genFn0() uint8 {
    var v6 float32 = 0.0
    fmt.Println(v6)
    switch 2147483647 {
    case -9223372036854775808:
        if (0.0 > v6) {
            v6 = v6
            v6 = 1e10
        }
    case -2147483648:
        if (false || true) {
            var v7 int16 = 0
            v8 := []int{1, 0}
            _ = v7
            _ = v8
        } else {
            fmt.Println(v6)
            switch "" {
            case "":
                if false {
                    var v9 bool = false
                    v10 := []int{-2147483648}
                    if (127 != -2147483648) {
                        v6 += (float32(1e10) * float32(0.1))
                    } else {
                        if v9 {
                            v9 = (!false)
                        } else {
                            if (2147483647 < 0) {
                                var v11 uint32 = 44
                                var v12 int64 = 526
                                _ = v11
                                _ = v12
                            } else {
                                var v13 uint16 = 0
                                var v14 string = ((("" + (("a" + ("hello" + "foo bar")) + "")) + (("foo bar" + ("hello" + ("a" + ""))) + ("a" + ("" + "hello")))) + ("foo bar" + ("a" + (("\n\t" + ("foo bar" + "hello")) + "hello"))))
                                v9 = (-1 >= 0)
                                _ = v13
                                _ = v14
                            }
                            if (-1 == -849) {
                                v9 = ((!v9) && true)
                                if v9 {
                                    if ((v9 && false) || (!v9)) {
                                        fmt.Println(v9, v9)
                                    } else {
                                        v15 := GenS1{F0: 42}
                                        v16 := GenS1{F0: 1}
                                        _ = v15
                                        _ = v16
                                    }
                                    v9 = v9
                                    if false {
                                        var v17 int64 = 0
                                        _ = v17
                                    } else {
                                        v18 := GenS2{F0: -336, F1: 1000, F2: 42}
                                        if len(v10) > 0 {
                                            v10[0] = 127
                                        }
                                        switch -9223372036854775808 {
                                        case -2147483648:
                                            v19 := GenS3{F0: 9223372036854775807}
                                            _ = v19
                                        case -454:
                                            v9 = false
                                        default:
                                            for fi20 := 0; fi20 < 20; fi20++ {
                                                v6 += v6
                                                fmt.Println(v6)
                                                _ = fi20
                                            }
                                        }
                                        _ = v18
                                    }
                                }
                                var v21 uint32 = 0
                                _ = v21
                            }
                        }
                    }
                    _ = v10
                    _ = v9
                }
            case "a":
                v22 := []int{}
                _ = v22
            default:
                v23 := []int{450, -8, 128, 1}
                _ = v23
            }
        }
    default:
        var v24 float32 = 0.0
        _ = v24
    }
    var v25 bool = ((1 <= 767) || (-2147483648 != 128))
    _ = v25
    _ = v6
    return 255
}
func genFn1() int64 {
    v26 := GenS3{F0: 4294967295}
    var v27 int32 = 0
    _ = v26
    _ = v27
    return -805
}

func main() {
    var v1 string = "a"
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]int)
    v4 := make(map[int]int)
    var v5 int8 = -1
    var v28 uint16 = 32767
    v29 := GenS3{F0: 32767}
    fmt.Println(v5, v5)
    v28 = v28
    switch (("foo bar" + v1) + (v1 + ("a" + ("" + ((v1 + ((v1 + ("a" + v1)) + (v1 + v1))) + v1))))) {
    case "":
        v28 += 0
    case "foo bar":
        v30 := []int{}
        _ = v30
    default:
        for fi31 := 0; fi31 < 4; fi31++ {
            fmt.Println(fi31, fi31, v1)
            if true {
                var v32 string = ((v1 + "") + v1)
                v33 := GenS1{F0: 2147483647}
                v5 = v5
                _ = v32
                _ = v33
            } else {
                v4[(fi31 >> ((0) & 63))] = fi31
            }
            _ = fi31
        }
    }
    v5 = (v5 >> ((18446744073709551615) & 63))
    var v34 uint64 = 347
    v1 = "foo bar"
    var v35 uint16 = (^(v28))
    for fi36 := 0; fi36 < 10; fi36++ {
        var v37 uint16 = v35
        _ = fi36
        _ = v37
    }
    var v38 float64 = 3.14159
    v34 -= 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v34
    _ = v35
    _ = v38
    _ = v4
    _ = v5
}
