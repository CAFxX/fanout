package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 bool }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 int16; F1 int; F2 uint16 }
type GenS4 struct { F0 int32; F1 uint32 }

func main() {
    var v1 int32 = -1
    var v2 bool = false
    v3 := make(map[uint32]int)
    v4 := make(map[int]bool)
    v5 := []string{("foo bar" + ("foo bar" + "hello")), (("a" + "") + "a"), "a"}
    v3[4294967295] = 9223372036854775807
    delete(v4, 9223372036854775807)
    for fi6 := 0; fi6 < 5; fi6++ {
        switch (fi6 / (((fi6 ^ 2147483647)) | 1)) {
        case 0:
            if true {
                fmt.Println(v1, v2)
                switch -2147483648 {
                case 42:
                    v7 := GenS1{F0: "foo bar", F1: 1, F2: (v2 || v2)}
                    _ = v7
                }
            }
        default:
            v8 := GenS1{F0: "foo bar", F1: 639, F2: v2}
            _ = v8
        }
        if len(v5) > 0 {
            v5[0] = "\n\t"
        }
        if len(v5) > 0 {
            v5[0] = ""
        }
        _ = fi6
    }
    fmt.Println(v2)
    fmt.Println(v2)
    if true {
        v9 := GenS1{F0: ("hello" + ("\n\t" + "\n\t")), F1: 2, F2: v2}
        _ = v9
    }
    fmt.Println(v1, v1)
    var v10 uint16 = 42
    if v2 {
        var v11 uint = 65535
        _ = v11
    }
    v1 += (v1 % (((-(v1))) | 1))
    fmt.Println(v10, v10, v1)
    fmt.Println(v2)
    v12 := GenS1{F0: "\n\t", F1: 4294967295, F2: (((!v2) || ((-2.25 != 0.1) && (!v2))) || false)}
    for fi13 := 0; fi13 < 7; fi13++ {
        v14 := GenS1{F0: "hello", F1: 1, F2: false}
        for fi15 := 0; fi15 < 9; fi15++ {
            v10 += v10
            if v2 {
                var v16 int64 = -792
                v2 = ((-6 > -54) && v2)
                _ = v16
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi15
        }
        _ = fi13
        _ = v14
    }
    v1 -= 2147483647
    var v17 uint8 = 255
    if ((0.1 - 0.1) != float64(3.14159)) {
        v18 := GenS1{F0: ("a" + ("\n\t" + "\n\t")), F1: 42, F2: v2}
        if false {
            v19 := GenS2{F0: 2, F1: 0}
            var v20 int = 42
            _ = v19
            _ = v20
        }
        v3[2] = 32768
        _ = v18
    } else {
        for fi21 := 0; fi21 < 12; fi21++ {
            var v22 uint64 = 168
            if v2 {
                v23 := GenS1{F0: (("a" + "a") + "hello"), F1: 2, F2: (!(1 != (v22 | (v22 / ((1) | 1)))))}
                _ = v23
            } else {
                switch (fi21 << ((42) & 63)) {
                case -955:
                    switch ("foo bar" + (("a" + "") + "\n\t")) {
                    case "":
                        for fi24 := 0; fi24 < 3; fi24++ {
                            var v25 int16 = 0
                            switch "hello" {
                            case "a":
                                v1 -= v1
                            }
                            _ = fi24
                            _ = v25
                        }
                    }
                case 128:
                    if ((-(v1)) >= v1) {
                        v26 := GenS3{F0: int16(v22), F1: fi21, F2: v10}
                        var v27 int32 = v1
                        v5 = append(v5, ("" + ("hello" + ("foo bar" + "hello"))))
                        _ = v26
                        _ = v27
                    }
                case 1:
                    var v28 uint = 61
                    _ = v28
                }
                var v29 int = fi21
                _ = v29
            }
            switch (fi21 >> ((0) & 63)) {
            case 127:
                v2 = (!(((((("hello" + "hello") + "foo bar") + "foo bar") + ("hello" + (("foo bar" + "a") + "\n\t"))) + ("a" + "hello")) <= "hello"))
            case -635:
                if true {
                    if (!false) {
                        v4[(fi21 | 32768)] = v2
                        v2 = (-1 > 0)
                    }
                    for fi30 := 0; fi30 < 17; fi30++ {
                        v1 -= (v1 | (v1 - v1))
                        fmt.Println(v1, v22, fi30)
                        fmt.Println(fi21)
                        _ = len(v4)
                        _ = fi30
                    }
                    v31 := GenS3{F0: -1, F1: -1, F2: v10}
                    _ = v31
                } else {
                    v10 = (v10 % (((v10 % ((42) | 1))) | 1))
                    var v32 int8 = 114
                    v33 := GenS4{F0: v1, F1: 0}
                    _ = v32
                    _ = v33
                }
            case 9223372036854775807:
                var v34 uint = 4294967295
                _ = v34
            }
            _ = fi21
            _ = v22
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
