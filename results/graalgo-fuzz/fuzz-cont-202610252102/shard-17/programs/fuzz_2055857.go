package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string; F1 uint64; F2 uint64 }
func genFn0() uint32 {
    var v5 string = ""
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 int = fi6
        _ = fi6
        _ = v7
    }
    _ = v5
    return 65535
}

func main() {
    var v1 uint32 = 2
    var v2 bool = true
    var v3 uint16 = 32767
    var v4 int32 = 42
    for fi8 := 0; fi8 < 18; fi8++ {
        switch ("\n\t" + "") {
        case "a":
            fmt.Println(v4, v4)
        default:
            v9 := GenS1{F0: fi8, F1: 2}
            _ = v9
        }
        var v10 int32 = -1
        _ = fi8
        _ = v10
    }
    switch -2147483648 {
    case 42:
        if (874 > int16(v3)) {
            for fi11 := 0; fi11 < 12; fi11++ {
                for fi12 := 0; fi12 < 16; fi12++ {
                    v4 = v4
                    switch "a" {
                    case "hello":
                        v13 := []int{}
                        _ = v13
                    }
                    v2 = (!v2)
                    _ = fi12
                }
                v4 += v4
                if v2 {
                    if (v2 && ((v3 + v3) < (v3 & v3))) {
                        var v14 string = "\n\t"
                        _ = v14
                    }
                    v15 := GenS1{F0: fi11, F1: v3}
                    _ = v15
                }
                _ = fi11
            }
            v16 := []int{}
            _ = v16
        }
    }
    fmt.Println(v2, v2)
    var v17 float32 = 0.1
    v17 += 0.0
    if (9223372036854775807 != 605) {
        v18 := []int{-461, 652}
        for fi19 := 0; fi19 < 20; fi19++ {
            switch "hello" {
            case "a":
                fmt.Println(v1)
            default:
                if v2 {
                    for fi20 := 0; fi20 < 15; fi20++ {
                        if ((v3 / ((v3) | 1)) < v3) {
                            if ("foo bar" == "a") {
                                v18 = append(v18, 128)
                            } else {
                                if (!(280 == 1)) {
                                    var v21 float32 = ((1e10 * v17) * float32(1e10))
                                    if len(v18) > 0 {
                                        v18[0] = (fi20 + (fi20 - fi19))
                                    }
                                    _ = v21
                                }
                            }
                        } else {
                            if len(v18) > 0 {
                                v18[0] = (fi19 / (((fi19 & (fi19 | fi20))) | 1))
                            }
                        }
                        _ = fi20
                    }
                    _ = len(v18)
                    fmt.Println(v17, v2, fi19)
                }
            }
            var v22 uint32 = genFn0()
            fmt.Println(v4)
            _ = fi19
            _ = v22
        }
        _ = v18
    } else {
        v23 := GenS2{F0: 1}
        var v24 int8 = -123
        if false {
            for fi25 := 0; fi25 < 3; fi25++ {
                var v26 uint64 = 18446744073709551615
                _ = fi25
                _ = v26
            }
            var v27 string = "\n\t"
            var v28 float32 = (v17 / v17)
            _ = v27
            _ = v28
        } else {
            v29 := []int{9223372036854775807, 4294967295, -9223372036854775808, 51}
            v30 := GenS3{F0: "foo bar", F1: 42, F2: 42}
            _ = v29
            _ = v30
        }
        _ = v23
        _ = v24
    }
    if v2 {
        switch "\n\t" {
        case "hello":
            var v31 uint32 = v1
            _ = v31
        case "":
            fmt.Println(v2, v4)
        default:
            for fi32 := 0; fi32 < 11; fi32++ {
                v33 := GenS2{F0: 2}
                for fi34 := 0; fi34 < 19; fi34++ {
                    v35 := []int{fi34, (fi32 * fi34), fi34}
                    var v36 int32 = v4
                    v37 := GenS1{F0: fi32, F1: v3}
                    _ = fi34
                    _ = v35
                    _ = v36
                    _ = v37
                }
                _ = fi32
                _ = v33
            }
        }
        v17 += (((v17 - v17) / v17) / ((0.0 - 3.14159) - 1.5))
        fmt.Println(v3)
    }
    var v38 uint64 = 42
    if ((v1 & (v1 << ((42) & 63))) >= v1) {
        var v39 int32 = (v4 | v4)
        v40 := []int{-785}
        _ = v39
        _ = v40
    } else {
        var v41 float64 = 0.0
        v17 = float32((0.0 / v41))
        _ = v41
    }
    var v42 uint8 = 0
    v43 := []int{127}
    _ = len(v43)
    for fi44 := 0; fi44 < 10; fi44++ {
        var v45 uint8 = 127
        if ((v1 & v1) > (v1 * genFn0())) {
            v43 = append(v43, -16)
            fmt.Println(v4, v2, v1)
        } else {
            v38 += 1
            var v46 uint16 = v3
            _ = v46
        }
        _ = fi44
        _ = v45
    }
    v42 = 126
    if (((("hello" + (("foo bar" + "") + "a")) + ("a" + "\n\t")) + "foo bar") != "hello") {
        for fi47 := 0; fi47 < 16; fi47++ {
            var v48 uint16 = (v3 - 1)
            v38 -= 4294967295
            _ = fi47
            _ = v48
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v43:", len(v43), v43)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v42
    _ = v43
}
