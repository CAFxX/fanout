package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int16; F1 uint; F2 bool }
type GenS5 struct { F0 uint }
func genFn0(p0 uint16, p1 int, p2 uint32) uint16 {
    switch ("" + "") {
    case "a":
        v5 := GenS1{F0: p2}
        _ = v5
    default:
        p1 += (p1 / (((-(p1))) | 1))
    }
    v6 := GenS1{F0: 2}
    var v7 int8 = 0
    if (true && (531 != 42)) {
        var v8 string = ""
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (p0 + p0)
}
func genFn1(p0 int64, p1 int32) uint {
    if (true && true) {
        v9 := []int{32767, -9223372036854775808, -2147483648}
        _ = v9
    } else {
        v10 := GenS2{F0: -32768, F1: 18446744073709551615}
        var v11 uint32 = 2147483648
        _ = v10
        _ = v11
    }
    if true {
        var v12 uint16 = 0
        switch "hello" {
        case "\n\t":
            p1 = -2147483648
        case "":
            v13 := []int{-1}
            _ = v13
        case "a":
            var v14 int64 = (^(p0))
            _ = v14
        default:
            switch 32767 {
            case 619:
                v15 := GenS3{F0: -128}
                _ = v15
            case 128:
                p1 += (p1 * p1)
            }
        }
        _ = v12
    }
    _ = p0
    _ = p1
    return 65535
}

func main() {
    v1 := []string{"", "a", "hello", "hello"}
    v2 := []uint32{}
    var v3 uint16 = 1
    var v4 int16 = 0
    for fi16 := 0; fi16 < 2; fi16++ {
        fmt.Println(v4, v4, v4)
        switch "foo bar" {
        case "\n\t":
            var v17 uint8 = 0
            _ = v17
        case "a":
            if false {
                for fi18 := 0; fi18 < 1; fi18++ {
                    var v19 uint64 = 357
                    v20 := GenS2{F0: -724, F1: v19}
                    switch ("" + ("\n\t" + ("hello" + "\n\t"))) {
                    case "foo bar":
                        _ = len(v1)
                    case "a":
                        v19 += v19
                    }
                    v21 := GenS4{F0: v4, F1: 0, F2: true}
                    fmt.Println(fi18, v3)
                    var v22 int32 = 2147483647
                    fmt.Println(v4, v3, v4)
                    _ = fi18
                    _ = v19
                    _ = v20
                    _ = v21
                    _ = v22
                }
                v3 = genFn0(v3, 32767, 2147483648)
            } else {
                switch -75 {
                case 32768:
                    if ((-2.25 + 0.1) > ((((1e10 + 3.14159) - 1.5) * 3.14159) / (1e10 + -2.25))) {
                        _ = len(v1)
                        v2 = append(v2, 65535)
                        v23 := GenS1{F0: 681}
                        _ = v23
                    } else {
                        var v24 int64 = -1
                        _ = v24
                    }
                case 4294967295:
                    var v25 bool = (0 == -2147483648)
                    _ = v25
                case 1:
                    fmt.Println(v4, fi16, v4)
                }
                if ((fi16 | fi16) < fi16) {
                    v4 = int16(v3)
                    v4 += (v4 * 32767)
                    fmt.Println(v4, fi16)
                } else {
                    switch "\n\t" {
                    case "hello":
                        v3 -= (v3 << ((2) & 63))
                    }
                    switch fi16 {
                    case 619:
                        fmt.Println(v4, fi16)
                    default:
                        v1 = append(v1, "foo bar")
                    }
                    var v26 bool = (-854 > 302)
                    _ = v26
                }
                v4 += -32768
            }
        }
        v4 -= v4
        _ = fi16
    }
    v27 := GenS4{F0: (v4 << ((255) & 63)), F1: 942, F2: (0 <= 1)}
    var v28 uint8 = 1
    if len(v2) > 0 {
        v2[0] = 65535
    }
    var v29 string = ""
    var v30 float32 = (-2.25 + 0.0)
    v31 := GenS1{F0: 570}
    var v32 string = ("foo bar" + v29)
    switch (v32 + "") {
    case "\n\t":
        switch (v32 + "") {
        case "\n\t":
            v30 = ((-2.25 * v30) + v30)
        default:
            if (42 > -2147483648) {
                v29 = v29
            } else {
                if (58 <= 4294967295) {
                    fmt.Println(v28, v30, v29)
                }
                switch ((((v32 + "foo bar") + v32) + "a") + (v32 + ((v32 + v29) + v29))) {
                case "a":
                    v4 -= 42
                }
                for fi33 := 0; fi33 < 1; fi33++ {
                    var v34 int8 = 100
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi33
                    _ = v34
                }
            }
        }
    default:
        v35 := GenS5{F0: 42}
        _ = v35
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
