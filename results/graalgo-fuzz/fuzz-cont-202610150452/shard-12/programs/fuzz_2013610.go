package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 string; F2 int32 }
type GenS4 struct { F0 int64; F1 uint32; F2 uint64 }

func main() {
    var v1 int32 = 2147483647
    v2 := []string{"a", (("foo bar" + "a") + ("\n\t" + ("a" + "foo bar"))), "a", ((((("\n\t" + ((("\n\t" + "a") + "\n\t") + "\n\t")) + "\n\t") + "hello") + ("a" + ("\n\t" + ("a" + "\n\t")))) + "\n\t")}
    v3 := []uint32{4294967295, 65535, 65535}
    var v4 uint = 1
    var v5 float64 = 1e10
    if (1e10 == 1e10) {
        if (!false) {
            fmt.Println(v4)
            var v6 float64 = v5
            _ = v6
        } else {
            for fi7 := 0; fi7 < 12; fi7++ {
                v8 := GenS1{F0: 0, F1: 681, F2: 73}
                _ = fi7
                _ = v8
            }
            if false {
                v9 := GenS2{F0: "\n\t"}
                v10 := GenS3{F0: 65535, F1: ("a" + ""), F2: (v1 * v1)}
                _ = len(v2)
                _ = v10
                _ = v9
            } else {
                if true {
                    v1 += (v1 / ((v1) | 1))
                }
                fmt.Println(v1)
                v1 += (v1 - v1)
            }
        }
    }
    if false {
        switch 1 {
        case -2147483648:
            if (0 >= 302) {
                var v11 int8 = -128
                var v12 bool = false
                var v13 bool = v12
                _ = v11
                _ = v12
                _ = v13
            } else {
                v1 = (v1 + -702)
                if (4294967295 < 2147483648) {
                    fmt.Println(v4, v4, v1)
                } else {
                    var v14 uint8 = 2
                    v1 = (v1 ^ (v1 / ((v1) | 1)))
                    _ = v14
                }
            }
        }
        v15 := GenS2{F0: ("foo bar" + "foo bar")}
        _ = v15
    }
    v3 = append(v3, 290)
    for fi16 := 0; fi16 < 9; fi16++ {
        var v17 bool = true
        v17 = v17
        if v17 {
            var v18 int64 = 42
            _ = v18
        }
        _ = fi16
        _ = v17
    }
    v19 := GenS2{F0: (("a" + ((("a" + ("" + ("foo bar" + (((("foo bar" + "a") + (("a" + ("a" + "\n\t")) + ("\n\t" + (("a" + ("hello" + ("hello" + "hello"))) + ("" + ("\n\t" + ((("hello" + ("foo bar" + "")) + "") + "hello"))))))) + (("" + "hello") + "\n\t")) + "hello")))) + ("a" + ("foo bar" + "foo bar"))) + "foo bar")) + ("hello" + "a"))}
    for fi20 := 0; fi20 < 12; fi20++ {
        v21 := GenS4{F0: -1, F1: 1, F2: 0}
        for fi22 := 0; fi22 < 6; fi22++ {
            v5 += v5
            _ = fi22
        }
        if (974 != 141) {
            if len(v2) > 0 {
                v2[0] = "hello"
            }
            v2 = append(v2, "\n\t")
        } else {
            v23 := GenS1{F0: 2, F1: 0, F2: 32767}
            v5 -= ((1e10 - v5) - 3.14159)
            fmt.Println(fi20, fi20)
            _ = v23
        }
        _ = fi20
        _ = v21
    }
    var v24 uint8 = 0
    var v25 int16 = 46
    v1 = (v1 % ((v1) | 1))
    v26 := GenS1{F0: (^(v4)), F1: 4294967295, F2: 2}
    v4 += (v4 / (((^(v4))) | 1))
    switch ("" + "foo bar") {
    case "a":
        v27 := GenS1{F0: 4294967295, F1: 970, F2: 2}
        _ = v27
    case "hello":
        if len(v3) > 0 {
            v3[0] = 2147483648
        }
    case "":
        fmt.Println(v24, v5, v25)
    default:
        for fi28 := 0; fi28 < 11; fi28++ {
            var v29 int64 = 0
            for fi30 := 0; fi30 < 5; fi30++ {
                v24 = v24
                _ = fi30
            }
            var v31 uint = v4
            _ = fi28
            _ = v29
            _ = v31
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
