package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 uint64; F2 int32 }
type GenS4 struct { F0 uint8; F1 uint32 }

func main() {
    var v1 uint8 = 255
    var v2 string = "\n\t"
    var v3 uint16 = 149
    v4 := make(map[string]bool)
    var v5 int = 541
    var v6 int32 = -1
    for fi7 := 0; fi7 < 12; fi7++ {
        v8 := []int{127}
        if len(v8) > 0 {
            v8[0] = v5
        }
        delete(v4, ((("foo bar" + v2) + v2) + v2))
        _ = fi7
        _ = v8
    }
    switch -9223372036854775808 {
    case 1:
        v9 := []int{}
        _ = v9
    case 0:
        v4[("foo bar" + (((("\n\t" + v2) + ((v2 + (v2 + ((("" + v2) + (v2 + ((v2 + "") + (((v2 + v2) + v2) + ("a" + v2))))) + (v2 + "hello")))) + v2)) + (v2 + "\n\t")) + v2))] = false
    case 42:
        fmt.Println(v1, v2, v1)
    default:
        fmt.Println(v2, v3, v5)
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        var v11 int = (^(fi10))
        _ = fi10
        _ = v11
    }
    var v12 int16 = 1000
    switch ("a" + "hello") {
    case "foo bar":
        if (!(!(2 > 0))) {
            var v13 float64 = (-2.25 * 0.1)
            delete(v4, "foo bar")
            v4[v2] = (true && (!true))
            _ = v13
        } else {
            for fi14 := 0; fi14 < 11; fi14++ {
                if true {
                    if true {
                        v15 := GenS1{F0: 0, F1: int(v1), F2: 0}
                        if (-2147483648 == v6) {
                            v5 -= (^(v5))
                            var v16 int32 = (-(v6))
                            v5 += 365
                            _ = v16
                        } else {
                            v17 := GenS1{F0: 2, F1: v5, F2: -1}
                            v18 := []int{fi14, v5}
                            _ = v17
                            _ = v18
                        }
                        var v19 uint = 255
                        _ = v15
                        _ = v19
                    }
                    switch fi14 {
                    case 668:
                        v20 := GenS2{F0: v2}
                        _ = v20
                    case 2147483647:
                        switch v5 {
                        case 0:
                            for fi21 := 0; fi21 < 15; fi21++ {
                                fmt.Println(v12, fi14)
                                _ = fi21
                            }
                        }
                    case -9223372036854775808:
                        switch (fi14 - (fi14 ^ fi14)) {
                        case 127:
                            fmt.Println(v2)
                        }
                    }
                }
                v22 := GenS3{F0: 255, F1: 4294967295, F2: (v6 << ((745) & 63))}
                _ = fi14
                _ = v22
            }
        }
    case "hello":
        var v23 int16 = v12
        _ = v23
    }
    var v24 int16 = (v12 | 32767)
    v25 := GenS4{F0: (v1 >> ((18446744073709551615) & 63)), F1: 42}
    v26 := []int{}
    fmt.Println(v1)
    v4[v2] = true
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
