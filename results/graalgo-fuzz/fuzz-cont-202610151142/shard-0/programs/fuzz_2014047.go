package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 int; F1 uint64 }
func genFn0(p0 int64, p1 int) uint16 {
    var v7 string = ((("" + "a") + ("" + "a")) + "\n\t")
    var v8 int64 = (p0 % ((p0) | 1))
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 65535
}
func genFn1(p0 string) bool {
    fmt.Println(p0)
    switch (p0 + p0) {
    case "foo bar":
        p0 = p0
    case "\n\t":
        fmt.Println(p0)
    case "hello":
        v9 := []int{-1, 42, 9223372036854775807}
        _ = v9
    default:
        fmt.Println(p0)
    }
    v10 := []int{}
    if true {
        if false {
            if ((-2.25 < 0.1) && true) {
                fmt.Println(p0)
            } else {
                if ((2 > 137) && true) {
                    for fi11 := 0; fi11 < 7; fi11++ {
                        for fi12 := 0; fi12 < 10; fi12++ {
                            if false {
                                var v13 string = ((p0 + "\n\t") + p0)
                                p0 = v13
                                var v14 int32 = 42
                                _ = v13
                                _ = v14
                            }
                            p0 = p0
                            var v15 uint64 = 1
                            _ = fi12
                            _ = v15
                        }
                        _ = len(v10)
                        _ = fi11
                    }
                } else {
                    var v16 uint = 255
                    _ = v16
                }
            }
            v10 = append(v10, 127)
            if (false || (!true)) {
                p0 = ""
            }
        }
    } else {
        var v17 float64 = 0.0
        _ = v17
    }
    _ = p0
    _ = v10
    return true
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 uint = 255
    v3 := make(map[uint32]string)
    v4 := make(map[int]string)
    var v5 float32 = (-2.25 + (0.1 / 1e10))
    var v6 int32 = -2147483648
    if true {
        if true {
            var v18 bool = (!true)
            v3[1] = "a"
            var v19 string = ("a" + "")
            _ = v18
            _ = v19
        }
        v6 += v6
        var v20 string = "\n\t"
        _ = v20
    } else {
        v21 := GenS1{F0: genFn0((-(v1)), -9223372036854775808), F1: 9223372036854775807, F2: 42}
        v6 -= v6
        _ = v21
    }
    for fi22 := 0; fi22 < 9; fi22++ {
        v23 := []int{}
        _ = fi22
        _ = v23
    }
    v1 += v1
    v24 := GenS2{F0: 4294967295, F1: 127}
    var v25 int16 = -32768
    if true {
        v26 := []int{-9223372036854775808, 2147483647, 2147483647}
        switch 32768 {
        case -2147483648:
            var v27 uint64 = 2
            _ = v27
        case 32768:
            var v28 int = 207
            _ = v28
        case 39:
            var v29 int16 = v25
            _ = v29
        default:
            var v30 string = ("" + "foo bar")
            _ = v30
        }
        switch "foo bar" {
        case "hello":
            switch 2147483647 {
            case 42:
                if true {
                    v2 = v2
                    if len(v26) > 0 {
                        v26[0] = 32768
                    }
                }
            default:
                for fi31 := 0; fi31 < 18; fi31++ {
                    for fi32 := 0; fi32 < 12; fi32++ {
                        var v33 float32 = (v5 + 3.14159)
                        _ = fi32
                        _ = v33
                    }
                    var v34 int8 = 1
                    _ = len(v26)
                    _ = fi31
                    _ = v34
                }
            }
        case "foo bar":
            v2 = (v2 >> (((v2 & v2)) & 63))
        case "\n\t":
            var v35 uint8 = 128
            _ = v35
        }
        _ = v26
    } else {
        if true {
            for fi36 := 0; fi36 < 3; fi36++ {
                for fi37 := 0; fi37 < 5; fi37++ {
                    v2 -= (v2 % ((v2) | 1))
                    v38 := []int{int(v1), fi36, 1}
                    _ = fi37
                    _ = v38
                }
                _ = fi36
            }
            if true {
                fmt.Println(v25, v25, v1)
                v39 := GenS3{F0: -9223372036854775808, F1: 354}
                _ = v39
            }
            for fi40 := 0; fi40 < 5; fi40++ {
                fmt.Println(fi40)
                var v41 float64 = (0.0 - (1e10 - 0.1))
                _ = fi40
                _ = v41
            }
        }
        v42 := GenS1{F0: 449, F1: 0, F2: 226}
        _ = v42
    }
    v25 -= v25
    v2 = (v2 % ((v2) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
