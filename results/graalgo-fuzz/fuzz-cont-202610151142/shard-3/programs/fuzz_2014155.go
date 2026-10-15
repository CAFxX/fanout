package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint8 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 bool }
func genFn0() uint16 {
    switch 2147483647 {
    case -297:
        var v4 int16 = 42
        _ = v4
    case 520:
        v5 := []int{-9223372036854775808, -37, 0, -736}
        _ = v5
    case 32767:
        v6 := []int{32768}
        _ = v6
    default:
        v7 := GenS1{F0: true, F1: ("hello" + ("\n\t" + (("foo bar" + "hello") + ("foo bar" + ((("a" + "foo bar") + ((("a" + "foo bar") + "a") + "foo bar")) + (((("\n\t" + "hello") + "a") + "hello") + "\n\t")))))), F2: 250}
        _ = v7
    }
    v8 := []int{-733, 171, 42, 0}
    for fi9 := 0; fi9 < 3; fi9++ {
        if (!false) {
            if len(v8) > 0 {
                v8[0] = fi9
            }
        } else {
            var v10 int64 = 1
            var v11 int32 = 42
            if true {
                if (2 != 42) {
                    if true {
                        var v12 uint32 = 2147483648
                        _ = len(v8)
                        v10 = (v10 % ((v10) | 1))
                        _ = v12
                    } else {
                        var v13 int32 = (v11 * -1)
                        if len(v8) > 0 {
                            v8[0] = fi9
                        }
                        _ = v13
                    }
                } else {
                    var v14 uint16 = 42
                    v11 -= 42
                    v15 := GenS2{F0: "a", F1: (32767 < 0)}
                    _ = v14
                    _ = v15
                }
            } else {
                var v16 uint8 = 2
                _ = v16
            }
            _ = v10
            _ = v11
        }
        var v17 bool = (42 != 2147483647)
        _ = fi9
        _ = v17
    }
    var v18 uint32 = 42
    _ = v18
    _ = v8
    return 32767
}
func genFn1() uint {
    v19 := []int{}
    switch "hello" {
    case "\n\t":
        var v20 float64 = 1.5
        _ = v20
    case "hello":
        var v21 uint32 = 4294967295
        _ = v21
    case "foo bar":
        for fi22 := 0; fi22 < 8; fi22++ {
            var v23 uint32 = 0
            _ = fi22
            _ = v23
        }
    default:
        var v24 int = 128
        _ = v24
    }
    var v25 string = "foo bar"
    v26 := GenS2{F0: "", F1: true}
    v25 = "hello"
    _ = v19
    _ = v25
    _ = v26
    return 338
}

func main() {
    var v1 float32 = 0.1
    var v2 int8 = -128
    var v3 uint8 = 220
    for fi27 := 0; fi27 < 19; fi27++ {
        if true {
            v1 = float32(float64(1e10))
            switch "a" {
            case "a":
                v3 += 87
            case "":
                v3 = v3
            case "\n\t":
                v28 := GenS3{F0: 4294967295, F1: -1}
                _ = v28
            }
        }
        _ = fi27
    }
    var v29 uint16 = 0
    v29 += (v29 >> ((2) & 63))
    if false {
        v1 = 1e10
        v29 = (v29 - 0)
    }
    v30 := GenS4{F0: (!((215 != 9223372036854775807) && false))}
    var v31 bool = false
    fmt.Println(v29, v31, v1)
    for fi32 := 0; fi32 < 19; fi32++ {
        v33 := []int{fi32}
        _ = fi32
        _ = v33
    }
    for fi34 := 0; fi34 < 17; fi34++ {
        var v35 int16 = -1
        _ = fi34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
