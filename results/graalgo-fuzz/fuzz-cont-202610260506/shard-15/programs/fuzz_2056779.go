package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 bool; F1 int16; F2 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int }
func genFn0() uint32 {
    var v4 int8 = 100
    for fi5 := 0; fi5 < 11; fi5++ {
        v4 = -128
        _ = fi5
    }
    for fi6 := 0; fi6 < 20; fi6++ {
        v4 += (v4 & (v4 & v4))
        switch ("hello" + ("" + "\n\t")) {
        case "":
            var v7 uint16 = 2
            _ = v7
        case "a":
            v8 := GenS1{F0: 18446744073709551615, F1: (^(v4))}
            _ = v8
        }
        _ = fi6
    }
    v4 += v4
    fmt.Println(v4)
    _ = v4
    return 1
}
func genFn1() int64 {
    var v9 uint8 = 2
    v10 := []int{}
    var v11 uint64 = 2
    v9 = 255
    v11 += (^(v11))
    _ = v10
    _ = v11
    _ = v9
    return 643
}

func main() {
    v1 := []string{("a" + ("hello" + "")), "a", "foo bar", ("foo bar" + "")}
    var v2 uint8 = 1
    var v3 float32 = -2.25
    if len(v1) > 0 {
        v1[0] = (("foo bar" + (("" + ("a" + (("\n\t" + ("\n\t" + "foo bar")) + (("" + ("a" + (("foo bar" + "hello") + ((("\n\t" + "foo bar") + "hello") + "hello")))) + ((("" + "\n\t") + ("a" + "hello")) + "hello"))))) + (("" + ((("\n\t" + "\n\t") + ("\n\t" + ("\n\t" + ""))) + ("a" + "\n\t"))) + "hello"))) + ("hello" + "foo bar"))
    }
    switch "a" {
    case "\n\t":
        v2 = v2
    case "hello":
        if false {
            v1 = append(v1, "\n\t")
            v2 += (v2 | v2)
        } else {
            if (v2 < v2) {
                v2 = 42
            }
        }
    case "foo bar":
        v12 := GenS2{F0: (18446744073709551615 >= 42), F1: 32767, F2: -15}
        _ = v12
    }
    var v13 int8 = -128
    var v14 int8 = 108
    var v15 int64 = 403
    v13 = v14
    var v16 float64 = (((float64(((3.14159 + v3) - v3)) / 1.5) / (0.0 - 1e10)) * 1e10)
    var v17 string = ""
    v3 -= (v3 / (float32(0.0) + v3))
    v15 -= (v15 << ((0) & 63))
    _ = len(v1)
    for fi18 := 0; fi18 < 8; fi18++ {
        if false {
            if true {
                v1 = append(v1, ((((("\n\t" + (v17 + v17)) + "\n\t") + "") + v17) + v17))
            } else {
                v19 := GenS3{F0: v13}
                var v20 int = fi18
                _ = v19
                _ = v20
            }
            fmt.Println(v15)
        } else {
            var v21 int16 = 1
            _ = v21
        }
        _ = fi18
    }
    if false {
        v1 = append(v1, (((v17 + ("foo bar" + "a")) + v17) + ((v17 + (((v17 + "\n\t") + ("" + ("foo bar" + v17))) + "foo bar")) + ("" + ("hello" + "hello")))))
        v2 = (v2 ^ (v2 << ((18446744073709551615) & 63)))
    } else {
        v3 += (v3 * v3)
    }
    v22 := GenS4{F0: (4294967295 == 664)}
    v23 := GenS5{F0: -1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v3
}
