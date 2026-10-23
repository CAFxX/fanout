package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32; F1 uint }
func genFn0(p0 uint16, p1 uint16) uint8 {
    var v4 bool = (!(!(!true)))
    v5 := GenS1{F0: v4, F1: 320}
    v6 := GenS1{F0: v4, F1: 42}
    v4 = (!v4)
    fmt.Println(p1, p1, v4)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 53
}
func genFn1() uint16 {
    switch "hello" {
    case "foo bar":
        var v7 int8 = 0
        _ = v7
    case "a":
        if (true && true) {
            var v8 int64 = -1
            _ = v8
        } else {
            switch "\n\t" {
            case "hello":
                var v9 string = "foo bar"
                _ = v9
            case "foo bar":
                for fi10 := 0; fi10 < 13; fi10++ {
                    for fi11 := 0; fi11 < 20; fi11++ {
                        var v12 int32 = -365
                        v13 := GenS2{F0: 214}
                        v14 := GenS3{F0: 4294967295, F1: 18446744073709551615}
                        _ = fi11
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                    var v15 uint16 = 255
                    var v16 uint16 = 196
                    _ = fi10
                    _ = v15
                    _ = v16
                }
            case "a":
                var v17 string = ""
                _ = v17
            }
            if (2147483647 > 42) {
                var v18 uint16 = 255
                _ = v18
            }
        }
    case "hello":
        v19 := GenS1{F0: false, F1: 2}
        _ = v19
    default:
        v20 := []int{-2147483648, 305, 42}
        _ = v20
    }
    var v21 int16 = 778
    v21 = (v21 | (-(v21)))
    if true {
        v22 := []int{42, 901}
        var v23 int32 = 0
        _ = v22
        _ = v23
    } else {
        for fi24 := 0; fi24 < 10; fi24++ {
            v25 := GenS3{F0: 2, F1: 18446744073709551615}
            _ = fi24
            _ = v25
        }
    }
    _ = v21
    return 498
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]int)
    var v3 int32 = -2147483648
    v3 = v3
    v3 = v3
    if (true || ((0 <= -114) && true)) {
        switch "foo bar" {
        case "":
            fmt.Println(v3)
        default:
            switch ("hello" + "hello") {
            case "hello":
                v3 = (-(v3))
            }
        }
        v26 := GenS3{F0: 2, F1: 1}
        _ = v26
    } else {
        var v27 uint16 = 1
        _ = v27
    }
    v3 = (-(v3))
    switch -9223372036854775808 {
    case 9223372036854775807:
        var v28 int32 = (v3 / ((v3) | 1))
        _ = v28
    case -818:
        fmt.Println(v3)
    case 32767:
        var v29 int = 9223372036854775807
        _ = v29
    default:
        v3 = (v3 ^ (v3 + (v3 & v3)))
    }
    var v30 int32 = 42
    v2[""] = 127
    _ = len(v2)
    if false {
        v3 -= v30
        var v31 float32 = -2.25
        _ = v31
    }
    for fi32 := 0; fi32 < 16; fi32++ {
        v30 += 2147483647
        _ = fi32
    }
    var v33 float32 = 0.1
    var v34 float32 = (v33 - v33)
    v35 := GenS2{F0: genFn1()}
    v34 -= float32((1.5 - ((1.5 * (3.14159 * 1.5)) * (float64((3.14159 * v33)) + float64(v34)))))
    v36 := GenS1{F0: false, F1: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v35
    _ = v36
}
