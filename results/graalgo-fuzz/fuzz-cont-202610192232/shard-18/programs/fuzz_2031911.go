package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint }
func genFn0(p0 uint32, p1 uint16, p2 int64) uint32 {
    v6 := []int{42, 42, -2147483648}
    switch 950 {
    case -2147483648:
        for fi7 := 0; fi7 < 14; fi7++ {
            v8 := GenS1{F0: 1, F1: 1}
            var v9 uint64 = 648
            p1 = p1
            _ = fi7
            _ = v8
            _ = v9
        }
    case 789:
        for fi10 := 0; fi10 < 19; fi10++ {
            var v11 uint32 = 0
            v12 := GenS2{F0: p2, F1: 127}
            _ = fi10
            _ = v11
            _ = v12
        }
    case 42:
        for fi13 := 0; fi13 < 3; fi13++ {
            for fi14 := 0; fi14 < 11; fi14++ {
                switch -859 {
                case 128:
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                }
                _ = fi14
            }
            _ = fi13
        }
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if len(v6) > 0 {
        v6[0] = -9223372036854775808
    }
    var v15 int16 = 1000
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v6
    return 388
}
func genFn1() int32 {
    v16 := GenS2{F0: 9223372036854775807, F1: 1}
    for fi17 := 0; fi17 < 12; fi17++ {
        var v18 string = ""
        _ = fi17
        _ = v18
    }
    _ = v16
    return -2147483648
}

func main() {
    var v1 uint16 = 0
    var v2 uint = 18446744073709551615
    var v3 int32 = 0
    var v4 int16 = 63
    v5 := make(map[string]string)
    if false {
        v4 -= (v4 ^ (v4 ^ 0))
        var v19 int = -505
        delete(v5, "\n\t")
        _ = v19
    }
    if true {
        v4 = (v4 ^ 286)
    } else {
        for fi20 := 0; fi20 < 11; fi20++ {
            v21 := []int{fi20, (fi20 % ((fi20) | 1)), (fi20 >> (((v2 % ((v2) | 1))) & 63)), (^(fi20))}
            _ = fi20
            _ = v21
        }
    }
    _ = len(v5)
    switch ("hello" + "hello") {
    case "hello":
        if true {
            for fi22 := 0; fi22 < 19; fi22++ {
                for fi23 := 0; fi23 < 2; fi23++ {
                    v4 = v4
                    v2 = v2
                    v3 = v3
                    v5["a"] = "a"
                    v1 = 307
                    _ = fi23
                }
                _ = fi22
            }
            if true {
                v3 -= -2147483648
            }
        }
    case "\n\t":
        var v24 int8 = -102
        _ = v24
    case "foo bar":
        switch 32768 {
        case 32768:
            fmt.Println(v4, v4)
        case 383:
            if (2 > 128) {
                v25 := v5["hello"]
                var v26 uint32 = 1
                _ = v25
                _ = v26
            }
        }
    default:
        v27 := GenS3{F0: -128, F1: v4}
        _ = v27
    }
    fmt.Println(v1)
    for fi28 := 0; fi28 < 17; fi28++ {
        var v29 bool = (true && false)
        v30 := []int{fi28, -625, fi28}
        for fi31 := 0; fi31 < 5; fi31++ {
            delete(v5, "foo bar")
            var v32 uint8 = 2
            if true {
                v4 += -32768
                var v33 uint = v2
                v34 := GenS4{F0: (v2 & v33)}
                _ = v33
                _ = v34
            }
            fmt.Println(fi28)
            v4 -= -32768
            fmt.Println(v4, v3)
            var v35 int64 = -9223372036854775808
            _ = fi31
            _ = v32
            _ = v35
        }
        _ = fi28
        _ = v29
        _ = v30
    }
    fmt.Println(v2, v2, v4)
    v36 := GenS2{F0: 42, F1: 0}
    if false {
        for fi37 := 0; fi37 < 18; fi37++ {
            for fi38 := 0; fi38 < 6; fi38++ {
                switch "hello" {
                case "foo bar":
                    fmt.Println(fi38)
                case "hello":
                    v39 := GenS5{F0: 18446744073709551615}
                    _ = v39
                }
                switch ("" + "a") {
                case "a":
                    v5[("foo bar" + ("a" + "a"))] = ("hello" + ("foo bar" + "hello"))
                    var v40 string = ("foo bar" + ("a" + "hello"))
                    switch 9223372036854775807 {
                    case -9223372036854775808:
                        v1 = v1
                    }
                    fmt.Println(v4, v40, v40)
                    _ = v40
                default:
                    v41 := []int{-279, (fi38 / (((fi38 / ((fi37) | 1))) | 1))}
                    _ = v41
                }
                var v42 float64 = float64((3.14159 * (0.1 / 1.0)))
                v43 := GenS1{F0: v3, F1: v4}
                if (v42 < v42) {
                    var v44 int8 = 1
                    fmt.Println(v1)
                    _ = v44
                }
                v45 := GenS2{F0: 1, F1: 100}
                _ = fi38
                _ = v42
                _ = v43
                _ = v45
            }
            _ = fi37
        }
    }
    v46 := []int{865, 127, -662, 127}
    var v47 int8 = 42
    _ = len(v46)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v46
    _ = v47
    _ = v5
}
