package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 uint64 }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 bool; F1 int8 }
type GenS5 struct { F0 uint16; F1 int16 }
func genFn0() uint {
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 float64 = (1.5 - -2.25)
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 654, F1: "hello", F2: 4294967295}
    var v10 uint64 = 4294967295
    if false {
        v10 = 2
    }
    v11 := []int{2147483647, 9223372036854775807, 562, -9223372036854775808}
    _ = v10
    _ = v11
    _ = v9
    return 0
}
func genFn1(p0 uint64) int {
    p0 += (p0 << ((255) & 63))
    var v12 int32 = 42
    var v13 bool = true
    _ = p0
    _ = v12
    _ = v13
    return 869
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float64 = 0.0
    var v3 int64 = 783
    var v4 int16 = -342
    v5 := make(map[int]string)
    v6 := []int{-837, -1, 2147483647}
    v14 := GenS1{F0: 42, F1: "\n\t", F2: 1}
    v2 = 0.0
    var v15 uint32 = 42
    v3 += v3
    if (485 == 805) {
        delete(v5, 2147483647)
        var v16 int8 = -1
        _ = v16
    } else {
        fmt.Println(v3)
        for fi17 := 0; fi17 < 2; fi17++ {
            switch fi17 {
            case 671:
                if true {
                    if (27 <= 2) {
                        v18 := GenS2{F0: 18446744073709551615, F1: (fi17 / ((fi17) | 1))}
                        _ = v18
                    } else {
                        _ = len(v6)
                    }
                    var v19 float64 = 1.5
                    _ = v19
                }
            case 71:
                v4 -= -98
            case -307:
                v4 = -1
            default:
                if false {
                    v20 := GenS3{F0: true, F1: 65535}
                    _ = v20
                } else {
                    var v21 uint8 = 255
                    for fi22 := 0; fi22 < 5; fi22++ {
                        v23 := GenS4{F0: (uint(v15) > 18446744073709551615), F1: 0}
                        _ = fi22
                        _ = v23
                    }
                    switch fi17 {
                    case 32768:
                        var v24 uint64 = 2
                        _ = v24
                    }
                    _ = v21
                }
            }
            v1[42] = v15
            _ = fi17
        }
        if false {
            if (42 <= (v15 - (v15 ^ v15))) {
                var v25 uint8 = 2
                _ = v25
            }
            v3 += 1
            if false {
                v4 = v4
                for fi26 := 0; fi26 < 5; fi26++ {
                    if len(v6) > 0 {
                        v6[0] = -1
                    }
                    _ = fi26
                }
            }
        } else {
            fmt.Println(v4)
            _ = len(v1)
            var v27 uint16 = 1
            _ = v27
        }
    }
    var v28 int = 9223372036854775807
    var v29 uint64 = 542
    for fi30 := 0; fi30 < 2; fi30++ {
        v29 += v29
        if true {
            var v31 int64 = v3
            _ = len(v1)
            fmt.Println(v3)
            _ = v31
        } else {
            v28 = 42
        }
        _ = fi30
    }
    var v32 bool = true
    var v33 string = "hello"
    for fi34 := 0; fi34 < 10; fi34++ {
        var v35 uint8 = 227
        _ = len(v5)
        for fi36 := 0; fi36 < 2; fi36++ {
            fmt.Println(fi34)
            v6 = append(v6, genFn1(0))
            v4 = v4
            if len(v6) > 0 {
                v6[0] = (fi36 + 2147483647)
            }
            _ = fi36
        }
        _ = fi34
        _ = v35
    }
    if (int(v4) != genFn1(v29)) {
        var v37 float64 = v2
        delete(v5, v28)
        switch -1 {
        case 842:
            v5[198] = ("a" + "foo bar")
        default:
            for fi38 := 0; fi38 < 13; fi38++ {
                var v39 bool = false
                _ = fi38
                _ = v39
            }
        }
        _ = v37
    } else {
        switch ("" + ("\n\t" + "a")) {
        case "hello":
            fmt.Println(v28, v3, v29)
        }
    }
    if false {
        var v40 int8 = -1
        v3 -= v3
        if true {
            var v41 uint32 = v15
            v42 := GenS5{F0: 788, F1: v4}
            v15 -= (v15 & uint32(v29))
            _ = v41
            _ = v42
        } else {
            v3 = (v3 - v3)
            var v43 uint16 = 937
            v28 = genFn1((v29 % ((v29) | 1)))
            _ = v43
        }
        _ = v40
    }
    v44 := GenS1{F0: (^(v29)), F1: (("a" + (v33 + ("a" + "foo bar"))) + (((v33 + (("foo bar" + (v33 + "\n\t")) + v33)) + "") + (v33 + "\n\t"))), F2: v29}
    fmt.Println(v32)
    switch (v28 + v28) {
    case 847:
        var v45 float64 = -2.25
        _ = v45
    }
    for fi46 := 0; fi46 < 16; fi46++ {
        var v47 int64 = (v3 + v3)
        fmt.Println(v32, v33, v3)
        _ = fi46
        _ = v47
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v44
    _ = v5
    _ = v6
}
