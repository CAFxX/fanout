package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 int8; F1 int; F2 string }
type GenS5 struct { F0 bool; F1 int32 }
type GenS6 struct { F0 uint8; F1 string; F2 uint32 }
func genFn0() uint64 {
    var v5 uint = 42
    var v6 int = 9223372036854775807
    if (false || (3.14159 <= 1.5)) {
        var v7 int64 = 1
        for fi8 := 0; fi8 < 17; fi8++ {
            v9 := GenS1{F0: 0}
            _ = fi8
            _ = v9
        }
        _ = v7
    } else {
        switch (^(v6)) {
        case 9223372036854775807:
            v5 -= (^(v5))
        default:
            if ((v6 % (((-(v6))) | 1)) < (v6 >> (((v5 / (((v5 | (v5 / ((v5) | 1)))) | 1))) & 63))) {
                fmt.Println(v5)
                var v10 uint32 = 943
                v11 := GenS2{F0: "hello"}
                _ = v10
                _ = v11
            } else {
                v12 := GenS3{F0: 1, F1: v5}
                var v13 int16 = -412
                _ = v12
                _ = v13
            }
        }
    }
    _ = v5
    _ = v6
    return 4294967295
}

func main() {
    v1 := []int{}
    var v2 string = ""
    v3 := make(map[string]bool)
    var v4 uint64 = 42
    var v14 int = 4294967295
    v4 = (^(v4))
    v2 = ""
    var v15 uint = 42
    for fi16 := 0; fi16 < 1; fi16++ {
        v17 := GenS4{F0: -1, F1: 32767, F2: v2}
        var v18 int16 = 32767
        var v19 uint16 = 2
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    if (true && (42 >= 55)) {
        v20 := GenS2{F0: ((v2 + "") + "\n\t")}
        if (2147483647 != 4294967295) {
            switch ("\n\t" + (v2 + v2)) {
            case "\n\t":
                for fi21 := 0; fi21 < 4; fi21++ {
                    for fi22 := 0; fi22 < 11; fi22++ {
                        if false {
                            var v23 string = v2
                            _ = v23
                        }
                        _ = fi22
                    }
                    if true {
                        _ = len(v3)
                        var v24 bool = false
                        switch 0 {
                        case 0:
                            var v25 int32 = -1
                            _ = v25
                        case -752:
                            v26 := GenS5{F0: ((("" + ((v2 + ("hello" + v2)) + v2)) + (("hello" + "\n\t") + "a")) != "a"), F1: 2147483647}
                            _ = v26
                        default:
                            if (v4 > 18446744073709551615) {
                                switch (fi21 * (^(v14))) {
                                case -9223372036854775808:
                                    var v27 uint8 = 128
                                    _ = v27
                                }
                                v2 = "a"
                            } else {
                                v28 := GenS3{F0: -128, F1: uint(v4)}
                                var v29 uint = v15
                                _ = v28
                                _ = v29
                            }
                        }
                        _ = v24
                    } else {
                        fmt.Println(v15, fi21)
                        v4 = 4294967295
                    }
                    v15 = v15
                    _ = fi21
                }
            case "hello":
                v15 += (^(v15))
            }
        } else {
            for fi30 := 0; fi30 < 11; fi30++ {
                v31 := GenS4{F0: -94, F1: 9223372036854775807, F2: (v2 + (v2 + (v2 + "hello")))}
                v15 += (v15 - v15)
                switch v2 {
                case "\n\t":
                    v32 := GenS6{F0: 127, F1: "\n\t", F2: 65535}
                    _ = v32
                }
                _ = fi30
                _ = v31
            }
            var v33 int8 = -103
            _ = v33
        }
        _ = v20
    }
    var v34 int8 = 127
    fmt.Println(v34, v2)
    for fi35 := 0; fi35 < 2; fi35++ {
        for fi36 := 0; fi36 < 20; fi36++ {
            v4 = v4
            v15 -= v15
            v37 := GenS4{F0: v34, F1: (fi36 * fi35), F2: v2}
            _ = fi36
            _ = v37
        }
        _ = len(v1)
        _ = fi35
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v34
    _ = v4
}
