package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 uint64 }
func genFn0(p0 uint) int {
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 bool = false
        fmt.Println(p0)
        v8 := []int{fi6, -752, (fi6 & (fi6 ^ -9223372036854775808))}
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := []int{9223372036854775807, 0, -850, 1}
    for fi10 := 0; fi10 < 9; fi10++ {
        v11 := GenS1{F0: 65535}
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 13; fi12++ {
        if true {
            if ((!(!("hello" == ("a" + "hello")))) || ("" <= "a")) {
                if (42 != 255) {
                    if false {
                        v9 = append(v9, -1)
                    } else {
                        var v13 uint64 = 0
                        for fi14 := 0; fi14 < 9; fi14++ {
                            var v15 int64 = 1
                            v9 = append(v9, 294)
                            _ = fi14
                            _ = v15
                        }
                        v16 := GenS1{F0: 18446744073709551615}
                        _ = v13
                        _ = v16
                    }
                    if true {
                        if false {
                            if false {
                                p0 = p0
                                fmt.Println(p0, p0)
                            }
                        } else {
                            var v17 int8 = 0
                            if len(v9) >= 2 {
                                v9 = v9[1:len(v9)]
                            }
                            _ = v17
                        }
                    }
                } else {
                    _ = len(v9)
                }
                p0 -= p0
            }
        } else {
            var v18 int16 = -734
            _ = v18
        }
        var v19 int8 = 1
        _ = fi12
        _ = v19
    }
    _ = p0
    _ = v9
    return -9223372036854775808
}

func main() {
    var v1 uint8 = 128
    var v2 uint64 = 42
    v3 := []int{}
    v4 := []int{364, 562}
    var v5 int16 = 1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch -2147483648 {
    case -1:
        fmt.Println(v1, v5, v5)
    }
    var v20 uint8 = 128
    var v21 bool = true
    fmt.Println(v21, v1)
    for fi22 := 0; fi22 < 16; fi22++ {
        for fi23 := 0; fi23 < 1; fi23++ {
            v5 = 0
            v24 := GenS2{F0: v21, F1: (v2 | 0)}
            v5 = (^(v5))
            _ = fi23
            _ = v24
        }
        _ = fi22
    }
    for fi25 := 0; fi25 < 16; fi25++ {
        switch "" {
        case "hello":
            v1 += v1
        case "\n\t":
            var v26 uint32 = 2147483647
            _ = v26
        }
        _ = fi25
    }
    fmt.Println(v5, v5, v20)
    v20 = 1
    var v27 int32 = -203
    var v28 string = "hello"
    if true {
        if (true && false) {
            var v29 int8 = 0
            _ = v29
        }
    }
    var v30 float32 = 0.1
    v31 := GenS1{F0: uint(v2)}
    fmt.Println(v2)
    var v32 int = 42
    var v33 float32 = v30
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
