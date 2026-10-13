package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 int8; F1 uint8; F2 int8 }
type GenS4 struct { F0 string; F1 int8; F2 uint }
func genFn0(p0 uint16, p1 uint64, p2 uint64) int {
    switch -2147483648 {
    case -1:
        fmt.Println(p1, p0)
    case 2147483647:
        for fi5 := 0; fi5 < 17; fi5++ {
            v6 := GenS1{F0: -2147483648, F1: 0}
            var v7 string = ("hello" + "\n\t")
            _ = fi5
            _ = v6
            _ = v7
        }
    default:
        switch 4294967295 {
        case -431:
            v8 := GenS1{F0: 42, F1: 18446744073709551615}
            _ = v8
        default:
            switch "hello" {
            case "foo bar":
                v9 := []int{4294967295, -159, 9223372036854775807, 0}
                _ = v9
            }
        }
    }
    v10 := []int{}
    fmt.Println(p2, p0)
    var v11 int16 = 571
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 128
}
func genFn1() uint16 {
    switch -1 {
    case 4294967295:
        for fi12 := 0; fi12 < 10; fi12++ {
            fmt.Println(fi12)
            _ = fi12
        }
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        switch (fi13 - fi13) {
        case 42:
            var v14 string = ("\n\t" + "")
            _ = v14
        case -1:
            var v15 int32 = -971
            _ = v15
        default:
            var v16 int32 = -2147483648
            _ = v16
        }
        v17 := GenS2{F0: 810, F1: true}
        var v18 uint = 255
        _ = fi13
        _ = v17
        _ = v18
    }
    return 1
}

func main() {
    v1 := []string{("a" + ("\n\t" + "\n\t")), ("foo bar" + "foo bar"), "foo bar", "foo bar"}
    var v2 int16 = -384
    var v3 int8 = 42
    var v4 string = "\n\t"
    var v19 string = v4
    v2 = v2
    if false {
        v20 := GenS2{F0: 1, F1: ((v19 + ((("\n\t" + "a") + (((v4 + (v19 + ((v4 + "") + ((v4 + v4) + "\n\t")))) + ("hello" + v4)) + (v19 + v4))) + (v4 + v4))) > ((v4 + (v4 + v19)) + v4))}
        _ = v20
    } else {
        for fi21 := 0; fi21 < 15; fi21++ {
            fmt.Println(v4, fi21)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v22 := GenS3{F0: 42, F1: 141, F2: (v3 | v3)}
            _ = fi21
            _ = v22
        }
        switch 2147483647 {
        case 463:
            v23 := GenS4{F0: ((("foo bar" + "") + v4) + ((v19 + ((v19 + ("\n\t" + v19)) + ((v19 + (v4 + "\n\t")) + (v4 + v4)))) + v19)), F1: (v3 + v3), F2: 906}
            _ = v23
        default:
            if (true || true) {
                if true {
                    var v24 uint16 = 255
                    _ = v24
                }
            }
        }
    }
    v1 = append(v1, v19)
    _ = len(v1)
    if len(v1) > 0 {
        v1[0] = v4
    }
    switch "\n\t" {
    case "":
        var v25 string = (((v19 + v4) + v19) + v19)
        _ = v25
    case "hello":
        var v26 float64 = -2.25
        _ = v26
    }
    v3 = 42
    v2 += -1
    var v27 int32 = -2147483648
    _ = len(v1)
    for fi28 := 0; fi28 < 15; fi28++ {
        v29 := GenS1{F0: (v27 % (((v27 << ((2) & 63))) | 1)), F1: 255}
        _ = fi28
        _ = v29
    }
    v4 = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v27
    _ = v3
    _ = v4
}
