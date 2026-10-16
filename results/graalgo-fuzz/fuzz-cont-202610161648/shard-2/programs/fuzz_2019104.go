package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 bool }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int16; F1 uint }

func main() {
    v1 := make(map[string]bool)
    var v2 int16 = 733
    var v3 int = 128
    var v4 int = (v3 % (((v3 >> ((277) & 63))) | 1))
    var v5 uint16 = 255
    delete(v1, (("a" + "\n\t") + ("\n\t" + "a")))
    v5 -= v5
    v3 = v3
    var v6 uint8 = 42
    v7 := GenS1{F0: (v2 == (-(v2))), F1: v3, F2: false}
    switch -9223372036854775808 {
    case -635:
        v8 := []int{(v4 & v3), (v3 % (((v3 * 276)) | 1))}
        _ = v8
    default:
        var v9 int16 = -1
        _ = v9
    }
    v3 += (v3 | (v3 - v4))
    if false {
        for fi10 := 0; fi10 < 19; fi10++ {
            v3 = 127
            for fi11 := 0; fi11 < 7; fi11++ {
                switch "" {
                case "":
                    fmt.Println(v6, v4)
                case "hello":
                    v12 := []int{(fi11 * (-(fi11))), (v3 + 127), fi11}
                    var v13 uint16 = (v5 + v5)
                    var v14 uint64 = 2
                    _ = v12
                    _ = v13
                    _ = v14
                case "\n\t":
                    var v15 int64 = 9223372036854775807
                    _ = v15
                default:
                    v5 += 477
                }
                v1["foo bar"] = false
                fmt.Println(v4, v5, v5)
                _ = fi11
            }
            v16 := GenS2{F0: 0, F1: (v2 * v2)}
            _ = fi10
            _ = v16
        }
        fmt.Println(v3, v2)
    } else {
        switch "\n\t" {
        case "hello":
            switch "foo bar" {
            case "foo bar":
                fmt.Println(v3, v2)
            default:
                for fi17 := 0; fi17 < 6; fi17++ {
                    v18 := GenS2{F0: 1, F1: (v2 | (v2 / ((v2) | 1)))}
                    var v19 int = fi17
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            }
        case "a":
            v2 += v2
        }
        v20 := GenS1{F0: (!false), F1: 0, F2: false}
        _ = v20
    }
    switch ("\n\t" + ("" + ("\n\t" + (("foo bar" + "a") + ("foo bar" + (("hello" + "a") + "a")))))) {
    case "hello":
        v21 := GenS3{F0: 1, F1: 1}
        _ = v21
    }
    v22 := GenS4{F0: v3}
    if ((("a" + "hello") + "hello") <= "foo bar") {
        v23 := GenS5{F0: (v2 >> ((316) & 63)), F1: 4294967295}
        v24 := []int{v3, (v4 / (((^(v3))) | 1)), 32767}
        _ = v23
        _ = v24
    } else {
        for fi25 := 0; fi25 < 18; fi25++ {
            var v26 uint16 = (^(v5))
            switch fi25 {
            case 639:
                if (18446744073709551615 >= 161) {
                    switch (v4 | (v4 / ((fi25) | 1))) {
                    case -510:
                        v27 := GenS4{F0: v4}
                        _ = v27
                    case -1:
                        if (776 <= 2147483647) {
                            v28 := []int{}
                            _ = v28
                        } else {
                            delete(v1, "a")
                        }
                    default:
                        fmt.Println(v5, v5)
                    }
                    var v29 uint8 = (^(v6))
                    _ = v29
                } else {
                    v5 += v26
                    var v30 uint32 = 2
                    fmt.Println(v2, v26, v3)
                    _ = v30
                }
            case -232:
                var v31 uint64 = 495
                _ = v31
            }
            fmt.Println(v6)
            _ = fi25
            _ = v26
        }
        for fi32 := 0; fi32 < 5; fi32++ {
            fmt.Println(v3, v3, v6)
            v3 -= (v3 / ((-627) | 1))
            _ = fi32
        }
        v33 := GenS3{F0: 479, F1: (v5 ^ 0)}
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
