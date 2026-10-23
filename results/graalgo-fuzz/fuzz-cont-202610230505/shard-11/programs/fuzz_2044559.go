package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 int16; F1 string; F2 int32 }
func genFn0(p0 uint8, p1 bool) int {
    if false {
        if false {
            fmt.Println(p0)
            if ((0 != 578) && false) {
                for fi4 := 0; fi4 < 2; fi4++ {
                    var v5 int = (fi4 ^ (fi4 >> ((255) & 63)))
                    p0 = 42
                    _ = fi4
                    _ = v5
                }
                fmt.Println(p1)
                if (100 < -94) {
                    p1 = p1
                    switch 32768 {
                    case 788:
                        v6 := GenS1{F0: -864, F1: p1}
                        _ = v6
                    case 32767:
                        for fi7 := 0; fi7 < 4; fi7++ {
                            for fi8 := 0; fi8 < 10; fi8++ {
                                var v9 int = fi7
                                _ = fi8
                                _ = v9
                            }
                            v10 := GenS1{F0: 521, F1: (1 > -117)}
                            _ = fi7
                            _ = v10
                        }
                    }
                    var v11 uint64 = 18446744073709551615
                    _ = v11
                } else {
                    p0 -= 104
                    p1 = p1
                    var v12 uint32 = 2
                    _ = v12
                }
            }
        }
    } else {
        var v13 int16 = -290
        if p1 {
            if p1 {
                fmt.Println(p1, p0, p1)
            } else {
                if false {
                    v14 := []int{-796}
                    if len(v14) >= 2 {
                        v14 = v14[1:len(v14)]
                    }
                    _ = v14
                }
                var v15 uint64 = 1
                _ = v15
            }
            var v16 uint = 0
            _ = v16
        } else {
            p0 = (^(p0))
        }
        _ = v13
    }
    fmt.Println(p1, p0)
    switch "\n\t" {
    case "foo bar":
        p0 = p0
    case "hello":
        for fi17 := 0; fi17 < 10; fi17++ {
            var v18 bool = (!false)
            v18 = (0 > 2147483647)
            _ = fi17
            _ = v18
        }
    default:
        if (65535 <= 11) {
            fmt.Println(p1, p0)
            v19 := []int{-360, 42}
            v19 = append(v19, 127)
            _ = v19
        }
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 int8 = -128
    var v2 int64 = 9223372036854775807
    var v3 uint64 = 18446744073709551615
    var v20 int16 = -830
    v21 := GenS1{F0: 641, F1: true}
    fmt.Println(v2)
    switch ("hello" + "\n\t") {
    case "hello":
        for fi22 := 0; fi22 < 2; fi22++ {
            if true {
                switch "hello" {
                case "":
                    var v23 int8 = (v1 >> ((42) & 63))
                    _ = v23
                case "a":
                    fmt.Println(v1, fi22, v3)
                }
            }
            _ = fi22
        }
    case "a":
        v2 -= v2
    default:
        fmt.Println(v3, v3, v3)
    }
    var v24 int = -931
    var v25 uint64 = v3
    if (v2 < v2) {
        fmt.Println(v3)
        v1 = v1
    }
    var v26 bool = false
    v24 -= (v24 % ((v24) | 1))
    v3 += 4294967295
    switch v24 {
    case 245:
        var v27 string = "a"
        _ = v27
    case 422:
        var v28 uint = 18446744073709551615
        _ = v28
    }
    v29 := []int{(v24 ^ (v24 ^ v24)), (v24 * v24), -899, (-(v24))}
    v30 := GenS1{F0: genFn0(128, (v2 < v2)), F1: true}
    if len(v29) > 0 {
        v29[0] = 32767
    }
    fmt.Println(v1)
    if (v26 || (!(true && false))) {
        fmt.Println(v2, v2)
        v31 := GenS1{F0: (v24 - v24), F1: (v26 && (-527 != -2147483648))}
        _ = v31
    }
    _ = len(v29)
    v32 := GenS2{F0: v20, F1: "hello", F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v32
}
