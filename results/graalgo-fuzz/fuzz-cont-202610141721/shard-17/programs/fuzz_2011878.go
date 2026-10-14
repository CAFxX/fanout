package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 int32 }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 int64; F1 bool; F2 uint64 }
type GenS4 struct { F0 int64; F1 uint8 }
type GenS5 struct { F0 int; F1 uint16 }
func genFn0(p0 uint, p1 uint64, p2 uint) int8 {
    fmt.Println(p2)
    if true {
        var v5 uint8 = 209
        _ = v5
    } else {
        for fi6 := 0; fi6 < 18; fi6++ {
            var v7 uint16 = 42
            _ = fi6
            _ = v7
        }
        if ((^(p0)) >= (p0 % ((42) | 1))) {
            v8 := GenS1{F0: 32767, F1: false, F2: 2147483647}
            switch 95 {
            case -9223372036854775808:
                v9 := GenS2{F0: 2147483647, F1: 1}
                _ = v9
            }
            if false {
                var v10 uint8 = 93
                v11 := GenS1{F0: -9223372036854775808, F1: true, F2: 2147483647}
                for fi12 := 0; fi12 < 7; fi12++ {
                    v10 -= 127
                    var v13 float32 = 0.1
                    _ = fi12
                    _ = v13
                }
                _ = v10
                _ = v11
            } else {
                for fi14 := 0; fi14 < 10; fi14++ {
                    for fi15 := 0; fi15 < 3; fi15++ {
                        v16 := []int{(fi15 ^ (-(fi15)))}
                        if false {
                            if false {
                                var v17 uint = (p0 * 42)
                                var v18 uint32 = 2147483647
                                _ = v17
                                _ = v18
                            }
                            var v19 string = "a"
                            v20 := GenS3{F0: -9223372036854775808, F1: true, F2: p1}
                            _ = v19
                            _ = v20
                        } else {
                            fmt.Println(p0, fi15, p2)
                            v21 := GenS1{F0: -402, F1: ((!(!(2 > 65535))) || false), F2: 42}
                            switch "foo bar" {
                            case "hello":
                                v22 := GenS3{F0: 42, F1: false, F2: p1}
                                _ = v22
                            }
                            _ = v21
                        }
                        _ = fi15
                        _ = v16
                    }
                    _ = fi14
                }
            }
            _ = v8
        }
        v23 := []int{}
        _ = v23
    }
    v24 := []int{-9223372036854775808, 591, -9223372036854775808}
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    return -128
}
func genFn1() int8 {
    var v25 bool = (!(("a" > "a") || false))
    for fi26 := 0; fi26 < 4; fi26++ {
        v27 := []int{fi26}
        switch "foo bar" {
        case "foo bar":
            for fi28 := 0; fi28 < 15; fi28++ {
                _ = len(v27)
                switch "\n\t" {
                case "a":
                    v29 := GenS2{F0: fi28, F1: genFn0(4294967295, 18446744073709551615, 899)}
                    _ = v29
                case "":
                    v25 = false
                default:
                    v25 = (!("hello" == "a"))
                }
                _ = fi28
            }
        case "":
            for fi30 := 0; fi30 < 12; fi30++ {
                _ = len(v27)
                var v31 int16 = -32768
                v31 += (v31 / (((v31 & (-(v31)))) | 1))
                _ = fi30
                _ = v31
            }
        default:
            v25 = v25
        }
        v32 := GenS4{F0: -1, F1: 210}
        _ = fi26
        _ = v27
        _ = v32
    }
    v33 := GenS5{F0: 32767, F1: 0}
    for fi34 := 0; fi34 < 8; fi34++ {
        v35 := []int{-2147483648, 9223372036854775807}
        v25 = v25
        _ = fi34
        _ = v35
    }
    _ = v25
    _ = v33
    return -128
}

func main() {
    var v1 float64 = -2.25
    var v2 uint16 = 170
    v3 := make(map[int]bool)
    var v4 uint8 = 2
    for fi36 := 0; fi36 < 19; fi36++ {
        v1 -= v1
        switch fi36 {
        case 42:
            switch ("foo bar" + "hello") {
            case "foo bar":
                var v37 int32 = -154
                _ = v37
            }
        }
        _ = fi36
    }
    var v38 uint8 = 1
    fmt.Println(v4)
    v39 := v3[2147483647]
    switch -603 {
    case 742:
        var v40 uint = 255
        _ = v40
    case -368:
        var v41 float32 = -2.25
        _ = v41
    case 32767:
        var v42 uint64 = 18446744073709551615
        _ = v42
    }
    var v43 float32 = 1e10
    v38 = v38
    v2 = (^(v2))
    v39 = true
    var v44 uint32 = uint32(v4)
    v45 := GenS1{F0: 258, F1: (v39 && v39), F2: 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v43
    _ = v44
    _ = v45
}
