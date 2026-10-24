package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 int32; F2 uint }
func genFn0(p0 uint16) int16 {
    var v5 uint8 = 1
    switch ("foo bar" + ("foo bar" + ("\n\t" + "\n\t"))) {
    case "\n\t":
        v6 := []int{-9223372036854775808, 127}
        _ = v6
    case "":
        p0 = (p0 % ((42) | 1))
    case "hello":
        if true {
            switch "a" {
            case "":
                var v7 uint32 = 4294967295
                _ = v7
            case "\n\t":
                v8 := []int{-9223372036854775808, -414}
                _ = v8
            case "foo bar":
                v9 := []int{32767, 2, 42, 128}
                _ = v9
            }
            p0 -= p0
        }
    }
    var v10 uint32 = 2
    switch 127 {
    case 32768:
        if false {
            for fi11 := 0; fi11 < 18; fi11++ {
                p0 -= p0
                var v12 int16 = -1
                _ = fi11
                _ = v12
            }
            switch "\n\t" {
            case "foo bar":
                fmt.Println(v10)
            case "hello":
                var v13 int64 = -9223372036854775808
                _ = v13
            case "\n\t":
                fmt.Println(v5, v10, p0)
            }
        }
    default:
        var v14 int64 = -1
        _ = v14
    }
    var v15 uint32 = (v10 - (v10 ^ v10))
    _ = p0
    _ = v10
    _ = v15
    _ = v5
    return -1
}

func main() {
    var v1 uint8 = 2
    var v2 int32 = 2147483647
    var v3 uint8 = 227
    var v4 uint8 = (v1 % ((128) | 1))
    v16 := GenS1{F0: -9223372036854775808}
    v2 += (v2 / (((^(v2))) | 1))
    v4 = 42
    v17 := []int{0}
    v17 = append(v17, 4294967295)
    switch "foo bar" {
    case "a":
        var v18 uint8 = (v1 << ((11) & 63))
        _ = v18
    default:
        for fi19 := 0; fi19 < 8; fi19++ {
            v1 += 32
            switch fi19 {
            case 355:
                fmt.Println(v4)
            case 969:
                v20 := GenS1{F0: -564}
                _ = v20
            case 32768:
                if len(v17) >= 2 {
                    v17 = v17[1:len(v17)]
                }
            }
            switch ("\n\t" + (("a" + (("hello" + "\n\t") + "hello")) + ("\n\t" + "hello"))) {
            case "hello":
                for fi21 := 0; fi21 < 8; fi21++ {
                    var v22 string = "a"
                    switch fi21 {
                    case 4294967295:
                        v23 := GenS2{F0: 2}
                        _ = v23
                    case 23:
                        v3 = v4
                    case -860:
                        if true {
                            if (!true) {
                                var v24 int8 = 0
                                v1 -= v4
                                _ = v24
                            }
                        } else {
                            var v25 int = fi19
                            fmt.Println(v3)
                            switch (v22 + (((v22 + "\n\t") + ("hello" + v22)) + ("foo bar" + ("" + ("" + v22))))) {
                            case "a":
                                var v26 uint16 = 0
                                _ = v26
                            case "foo bar":
                                if false {
                                    switch -648 {
                                    case -9223372036854775808:
                                        var v27 uint64 = 1
                                        _ = v27
                                    case 127:
                                        _ = len(v17)
                                    }
                                    var v28 uint8 = v1
                                    _ = v28
                                }
                            case "":
                                _ = len(v17)
                            }
                            _ = v25
                        }
                    }
                    var v29 int64 = -12
                    _ = fi21
                    _ = v22
                    _ = v29
                }
            case "a":
                var v30 float32 = 3.14159
                _ = v30
            case "foo bar":
                v31 := GenS3{F0: 127, F1: 42, F2: 18446744073709551615}
                _ = v31
            }
            _ = fi19
        }
    }
    var v32 string = "\n\t"
    if len(v17) >= 2 {
        v17 = v17[1:len(v17)]
    }
    var v33 int = 0
    v33 -= 4294967295
    var v34 uint64 = 267
    fmt.Println(v34, v4, v3)
    var v35 bool = (uint8(v34) < 249)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
}
