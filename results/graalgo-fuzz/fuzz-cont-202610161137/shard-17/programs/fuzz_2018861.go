package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 int; F1 bool }
func genFn0(p0 uint64) string {
    p0 = p0
    v7 := []int{1, 32768, 9223372036854775807}
    _ = p0
    _ = v7
    return ("a" + "a")
}
func genFn1(p0 int64) int {
    v8 := GenS1{F0: 0}
    switch (("hello" + (("a" + "foo bar") + ("hello" + ("foo bar" + ("foo bar" + ""))))) + "hello") {
    case "foo bar":
        switch 32768 {
        case 812:
            v9 := []int{2147483647, 32767}
            _ = v9
        case 1:
            v10 := []int{128, 1}
            _ = v10
        default:
            p0 = -1
        }
    case "a":
        var v11 bool = (!(127 > 2147483647))
        _ = v11
    case "hello":
        for fi12 := 0; fi12 < 8; fi12++ {
            fmt.Println(p0)
            _ = fi12
        }
    }
    _ = p0
    _ = v8
    return 213
}

func main() {
    var v1 float32 = (1.5 + 0.1)
    v2 := make(map[string]int)
    var v3 uint64 = 175
    var v4 uint8 = 69
    var v5 int32 = 42
    v6 := make(map[uint32]bool)
    v13 := v6[0]
    switch 127 {
    case -413:
        v14 := GenS1{F0: 32767}
        _ = v14
    case 899:
        var v15 string = ("a" + ("\n\t" + ("\n\t" + "\n\t")))
        _ = v15
    }
    fmt.Println(v3, v5)
    if v13 {
        v4 = (v4 - (v4 ^ (v4 / (((v4 >> ((737) & 63))) | 1))))
    }
    if (v4 >= v4) {
        v16 := []int{}
        _ = v16
    }
    v17 := GenS2{F0: v4, F1: 255}
    var v18 int = -326
    v18 = 9223372036854775807
    if v13 {
        switch v18 {
        case 1:
            if v13 {
                fmt.Println(v18, v18)
                if (("\n\t" + "foo bar") > "\n\t") {
                    _ = len(v6)
                }
                v19 := []int{v18, genFn1(1), -343}
                _ = v19
            }
        }
        for fi20 := 0; fi20 < 17; fi20++ {
            v21 := GenS2{F0: 136, F1: 822}
            _ = fi20
            _ = v21
        }
    }
    fmt.Println(v3, v4, v18)
    for fi22 := 0; fi22 < 6; fi22++ {
        switch "\n\t" {
        case "hello":
            v5 -= (v5 + 2147483647)
        case "\n\t":
            var v23 int8 = -128
            _ = v23
        case "foo bar":
            v2["foo bar"] = (-(fi22))
        default:
            fmt.Println(v4)
        }
        var v24 int = v18
        _ = fi22
        _ = v24
    }
    for fi25 := 0; fi25 < 17; fi25++ {
        if (true || (-17 > 42)) {
            var v26 int16 = -1
            _ = v26
        } else {
            v27 := []int{-859, (fi25 | (v18 >> ((0) & 63))), (fi25 + 127), (fi25 + int(v3))}
            if (!v13) {
                for fi28 := 0; fi28 < 20; fi28++ {
                    v1 -= (v1 + ((((float32((0.1 + 0.1)) * float32(0.1)) / v1) + v1) + v1))
                    var v29 uint = 231
                    var v30 uint64 = (v3 | (v3 / ((18446744073709551615) | 1)))
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                for fi31 := 0; fi31 < 7; fi31++ {
                    if ("\n\t" <= "hello") {
                        var v32 int64 = -1
                        v3 = (v3 + v3)
                        _ = v32
                    }
                    switch "\n\t" {
                    case "foo bar":
                        v27 = append(v27, (fi31 << ((4294967295) & 63)))
                    case "a":
                        v4 = (^(v4))
                    }
                    fmt.Println(fi25, v4, v5)
                    _ = fi31
                }
                fmt.Println(v5, v18, v3)
            }
            _ = len(v27)
            _ = v27
        }
        for fi33 := 0; fi33 < 20; fi33++ {
            v4 -= 0
            var v34 int32 = (v5 % ((v5) | 1))
            if (!true) {
                if v13 {
                    v3 = v3
                }
                var v35 uint64 = (v3 >> ((18446744073709551615) & 63))
                var v36 uint8 = v4
                _ = v35
                _ = v36
            }
            _ = fi33
            _ = v34
        }
        _ = fi25
    }
    if false {
        fmt.Println(v13, v4)
        switch ("hello" + "foo bar") {
        case "":
            v37 := GenS3{F0: (v5 + (v5 | (v5 | 2147483647))), F1: 321}
            _ = v37
        case "hello":
            var v38 uint32 = 4294967295
            _ = v38
        case "a":
            var v39 float64 = 3.14159
            _ = v39
        }
    }
    for fi40 := 0; fi40 < 13; fi40++ {
        for fi41 := 0; fi41 < 18; fi41++ {
            var v42 float64 = (0.0 + 1e10)
            _ = fi41
            _ = v42
        }
        _ = fi40
    }
    v43 := GenS4{F0: 1, F1: v13}
    var v44 uint8 = 233
    var v45 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
