package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int64 }
type GenS2 struct { F0 int16; F1 int16; F2 int16 }
type GenS3 struct { F0 string; F1 int32; F2 int8 }
type GenS4 struct { F0 int32; F1 bool }
type GenS5 struct { F0 uint16 }
func genFn0(p0 bool, p1 bool, p2 uint16) int {
    switch 127 {
    case 953:
        p2 -= p2
    }
    var v4 int16 = 1
    switch ("" + "a") {
    case "a":
        var v5 float32 = 0.1
        _ = v5
    }
    if p0 {
        v6 := GenS1{F0: "a", F1: 0, F2: -467}
        var v7 int64 = -9223372036854775808
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 4294967295
}

func main() {
    var v1 uint16 = 255
    var v2 uint8 = 206
    var v3 uint64 = 337
    for fi8 := 0; fi8 < 3; fi8++ {
        v9 := []int{(fi8 / (((-(fi8))) | 1)), (fi8 + 4294967295), fi8, (fi8 % (((fi8 | (-(fi8)))) | 1))}
        var v10 uint32 = 2147483648
        if true {
            var v11 bool = (((fi8 / ((fi8) | 1)) == fi8) && true)
            _ = v11
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    if false {
        v2 = uint8(v3)
        v12 := []int{}
        if (!("hello" >= ("\n\t" + ""))) {
            var v13 int32 = 0
            var v14 int32 = (v13 >> ((2) & 63))
            _ = v13
            _ = v14
        } else {
            for fi15 := 0; fi15 < 20; fi15++ {
                switch "\n\t" {
                case "\n\t":
                    var v16 int16 = -485
                    _ = v16
                default:
                    v3 -= v3
                }
                for fi17 := 0; fi17 < 5; fi17++ {
                    v2 += 42
                    fmt.Println(v2)
                    _ = fi17
                }
                _ = fi15
            }
            for fi18 := 0; fi18 < 14; fi18++ {
                fmt.Println(v3, v1, v2)
                _ = fi18
            }
        }
        _ = v12
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        v2 -= (v2 * 127)
        for fi20 := 0; fi20 < 6; fi20++ {
            var v21 int = -1
            switch "a" {
            case "foo bar":
                var v22 int = (fi20 / (((v21 * 128)) | 1))
                _ = v22
            case "a":
                fmt.Println(fi20)
                switch "hello" {
                case "foo bar":
                    switch (("a" + (("a" + "") + "a")) + (("\n\t" + ("foo bar" + (("\n\t" + "") + ((("" + "\n\t") + ("foo bar" + "foo bar")) + "foo bar")))) + ("\n\t" + ""))) {
                    case "hello":
                        v23 := GenS2{F0: 1, F1: -32768, F2: -352}
                        _ = v23
                    case "foo bar":
                        fmt.Println(v1, v2)
                    }
                case "\n\t":
                    v3 -= (v3 >> ((42) & 63))
                case "":
                    v1 = (v1 & 239)
                }
                switch -9223372036854775808 {
                case -1:
                    v24 := GenS3{F0: ((("a" + "a") + "hello") + ("hello" + ("hello" + "hello"))), F1: -498, F2: 100}
                    _ = v24
                case 213:
                    var v25 string = ("" + "foo bar")
                    v26 := GenS2{F0: 1, F1: 50, F2: 1}
                    v27 := GenS4{F0: 42, F1: (!false)}
                    _ = v25
                    _ = v26
                    _ = v27
                case 1:
                    var v28 float64 = (1.5 * 0.1)
                    _ = v28
                }
            case "":
                switch fi20 {
                case -1:
                    v29 := []int{(v21 ^ (^(v21))), 4294967295, (fi20 * v21), fi19}
                    var v30 int8 = 4
                    var v31 int8 = (v30 & (v30 * 78))
                    _ = v29
                    _ = v30
                    _ = v31
                default:
                    var v32 int16 = 1
                    _ = v32
                }
            default:
                var v33 int = 42
                _ = v33
            }
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    v34 := GenS1{F0: (("\n\t" + ("foo bar" + "")) + "foo bar"), F1: -20, F2: 9223372036854775807}
    fmt.Println(v3, v2)
    v35 := []int{42}
    v1 += v1
    v36 := GenS5{F0: v1}
    v2 = (v2 - 11)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
}
