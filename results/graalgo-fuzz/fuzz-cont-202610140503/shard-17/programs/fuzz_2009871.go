package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 uint16; F1 uint32; F2 uint8 }

func main() {
    v1 := []string{"foo bar"}
    v2 := make(map[int]int)
    var v3 uint64 = 4294967295
    v4 := []int{42, 32767, -1}
    v5 := make(map[int]string)
    var v6 string = (("\n\t" + ("\n\t" + (("\n\t" + "foo bar") + ("hello" + "foo bar")))) + "\n\t")
    v7 := GenS1{F0: -37, F1: 255}
    var v8 int = -9223372036854775808
    for fi9 := 0; fi9 < 10; fi9++ {
        for fi10 := 0; fi10 < 3; fi10++ {
            switch v8 {
            case 9223372036854775807:
                var v11 string = v6
                v12 := GenS2{F0: 0, F1: 1}
                _ = v11
                _ = v12
            }
            fmt.Println(v3)
            v13 := GenS1{F0: 127, F1: 1}
            _ = fi10
            _ = v13
        }
        if len(v4) > 0 {
            v4[0] = int(v3)
        }
        v6 = (v6 + v6)
        _ = fi9
    }
    switch "foo bar" {
    case "":
        for fi14 := 0; fi14 < 11; fi14++ {
            var v15 uint = 1
            _ = fi14
            _ = v15
        }
    default:
        fmt.Println(v3, v6, v8)
    }
    fmt.Println(v3, v8)
    var v16 int32 = 2147483647
    v8 += (v8 - (v8 / ((-864) | 1)))
    switch "hello" {
    case "a":
        if true {
            if true {
                v17 := GenS1{F0: 0, F1: 1}
                _ = v17
            } else {
                if true {
                    switch v6 {
                    case "a":
                        for fi18 := 0; fi18 < 17; fi18++ {
                            for fi19 := 0; fi19 < 2; fi19++ {
                                if len(v4) > 0 {
                                    v4[0] = fi18
                                }
                                v20 := GenS1{F0: -128, F1: 18446744073709551615}
                                _ = fi19
                                _ = v20
                            }
                            _ = fi18
                        }
                    default:
                        for fi21 := 0; fi21 < 6; fi21++ {
                            v22 := GenS3{F0: 42, F1: 1, F2: 1}
                            fmt.Println(v3, v3)
                            v8 = v8
                            _ = fi21
                            _ = v22
                        }
                    }
                    fmt.Println(v8)
                    if (-1 <= 388) {
                        for fi23 := 0; fi23 < 17; fi23++ {
                            v24 := GenS3{F0: 347, F1: 519, F2: 0}
                            _ = fi23
                            _ = v24
                        }
                        v3 = v3
                        v6 = "foo bar"
                    }
                } else {
                    var v25 int64 = 1
                    var v26 int16 = -32768
                    _ = v25
                    _ = v26
                }
            }
        }
    case "\n\t":
        if false {
            v8 += v8
        }
    case "hello":
        if ((267 < 0) && false) {
            for fi27 := 0; fi27 < 12; fi27++ {
                switch 0 {
                case 9223372036854775807:
                    v6 = (v6 + (v6 + (("a" + ((v6 + "hello") + v6)) + (v6 + v6))))
                case 0:
                    if len(v1) > 0 {
                        v1[0] = "foo bar"
                    }
                }
                var v28 uint64 = v3
                delete(v2, fi27)
                _ = fi27
                _ = v28
            }
        }
    }
    v29 := GenS1{F0: 100, F1: 255}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
