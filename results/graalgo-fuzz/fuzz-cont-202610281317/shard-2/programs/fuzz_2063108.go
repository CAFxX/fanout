package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 int16 }
type GenS2 struct { F0 int32; F1 uint16; F2 int }
type GenS3 struct { F0 uint64; F1 bool; F2 int16 }
type GenS4 struct { F0 int; F1 int16 }
type GenS5 struct { F0 uint64; F1 int; F2 int32 }
type GenS6 struct { F0 uint32; F1 uint32 }
func genFn0() uint8 {
    v4 := []int{42, 42, -1, -2147483648}
    switch (("a" + ("" + "foo bar")) + ("hello" + "a")) {
    case "a":
        var v5 bool = true
        _ = v5
    case "\n\t":
        for fi6 := 0; fi6 < 4; fi6++ {
            var v7 int8 = -128
            if (-1 != 1) {
                v8 := GenS1{F0: 4294967295, F1: -209, F2: 1000}
                var v9 int = -750
                _ = v8
                _ = v9
            } else {
                v10 := GenS1{F0: 18446744073709551615, F1: -2147483648, F2: int16(fi6)}
                _ = len(v4)
                if (!(0 >= 0)) {
                    for fi11 := 0; fi11 < 10; fi11++ {
                        v7 = (v7 & 127)
                        var v12 int32 = 0
                        _ = fi11
                        _ = v12
                    }
                    var v13 int = fi6
                    switch (("foo bar" + "hello") + (("\n\t" + "foo bar") + "\n\t")) {
                    case "a":
                        _ = len(v4)
                    case "hello":
                        switch ("\n\t" + "") {
                        case "\n\t":
                            v13 -= (v13 + v13)
                        case "foo bar":
                            var v14 uint8 = 2
                            _ = v14
                        case "hello":
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                    }
                    _ = v13
                } else {
                    v7 = v7
                    if (("hello" > ("\n\t" + "foo bar")) || false) {
                        v15 := GenS2{F0: -214, F1: 42, F2: 9223372036854775807}
                        var v16 int32 = 2147483647
                        v16 = v16
                        _ = v15
                        _ = v16
                    } else {
                        v17 := GenS2{F0: -2147483648, F1: 255, F2: fi6}
                        var v18 uint = 42
                        _ = v17
                        _ = v18
                    }
                    _ = len(v4)
                }
                _ = v10
            }
            v7 -= -1
            _ = fi6
            _ = v7
        }
    default:
        var v19 uint8 = 42
        _ = v19
    }
    for fi20 := 0; fi20 < 6; fi20++ {
        v21 := GenS3{F0: 387, F1: false, F2: -792}
        var v22 int8 = 127
        _ = fi20
        _ = v21
        _ = v22
    }
    v4 = append(v4, -544)
    _ = v4
    return 127
}

func main() {
    var v1 int16 = -892
    var v2 int = 2147483647
    var v3 uint8 = 2
    v23 := []int{v2, (^(v2))}
    for fi24 := 0; fi24 < 1; fi24++ {
        for fi25 := 0; fi25 < 11; fi25++ {
            if true {
                var v26 uint8 = genFn0()
                v26 = (v3 % ((v26) | 1))
                _ = v26
            } else {
                var v27 uint64 = 1
                if len(v23) >= 2 {
                    v23 = v23[1:len(v23)]
                }
                _ = v27
            }
            v28 := GenS2{F0: 0, F1: 1, F2: (-(fi25))}
            v1 = v1
            _ = fi25
            _ = v28
        }
        _ = fi24
    }
    fmt.Println(v3)
    var v29 int = v2
    v30 := GenS4{F0: (v29 * 9223372036854775807), F1: 1}
    v2 -= (v29 + 2147483647)
    fmt.Println(v2, v1, v2)
    if ((v2 + v2) != (^(v2))) {
        v31 := GenS5{F0: 18446744073709551615, F1: 4294967295, F2: 1}
        _ = v31
    }
    var v32 int8 = 77
    var v33 int8 = v32
    v34 := GenS6{F0: 4294967295, F1: 651}
    var v35 int16 = v1
    if false {
        v29 -= (-(v29))
        fmt.Println(v33, v1, v29)
    }
    if true {
        if len(v23) >= 2 {
            v23 = v23[1:len(v23)]
        }
        v2 = 2147483647
        for fi36 := 0; fi36 < 4; fi36++ {
            var v37 int64 = 1
            _ = fi36
            _ = v37
        }
    } else {
        v38 := GenS4{F0: v29, F1: -32768}
        _ = v38
    }
    var v39 int64 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v39
}
