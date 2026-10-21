package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint16 }
type GenS2 struct { F0 int8; F1 int32; F2 int64 }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 uint32; F1 int; F2 int8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint; F1 uint16 }
func genFn0(p0 uint, p1 uint32, p2 uint8) uint {
    for fi7 := 0; fi7 < 9; fi7++ {
        for fi8 := 0; fi8 < 1; fi8++ {
            v9 := GenS1{F0: 0, F1: -32768, F2: uint16(p0)}
            p0 -= p0
            var v10 uint = (^(p0))
            var v11 uint = (v10 / (((^(v10))) | 1))
            p0 = p0
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
        fmt.Println(fi7, p1)
        if (4294967295 < 853) {
            for fi12 := 0; fi12 < 7; fi12++ {
                p1 = 2
                var v13 string = "hello"
                var v14 float32 = 1e10
                _ = fi12
                _ = v13
                _ = v14
            }
        } else {
            if true {
                fmt.Println(p0, p0)
                p1 -= 428
                var v15 uint32 = 2147483647
                _ = v15
            }
        }
        _ = fi7
    }
    v16 := []int{-436}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return 18446744073709551615
}
func genFn1() int16 {
    var v17 uint64 = 4294967295
    v18 := []int{4294967295, 212, 819}
    if (4294967295 < v17) {
        v19 := GenS2{F0: -58, F1: 43, F2: -9223372036854775808}
        v20 := GenS1{F0: -43, F1: -1, F2: 1}
        if false {
            if false {
                v21 := GenS1{F0: 125, F1: 0, F2: 42}
                if false {
                    switch "foo bar" {
                    case "a":
                        v22 := GenS3{F0: 65535, F1: 9223372036854775807}
                        _ = v22
                    case "foo bar":
                        v17 += v17
                    }
                }
                var v23 uint16 = 32767
                _ = v21
                _ = v23
            }
            if (!false) {
                v17 += (^(v17))
                fmt.Println(v17)
                if false {
                    v24 := GenS4{F0: 4294967295, F1: 128, F2: -1}
                    _ = v24
                } else {
                    if (false && true) {
                        fmt.Println(v17)
                    } else {
                        if (-1 >= -418) {
                            v17 = (v17 ^ v17)
                            v25 := GenS1{F0: -80, F1: 1, F2: 255}
                            _ = v25
                        } else {
                            for fi26 := 0; fi26 < 17; fi26++ {
                                v17 = (v17 & v17)
                                var v27 int = (-(fi26))
                                _ = fi26
                                _ = v27
                            }
                            v28 := GenS5{F0: 2147483647}
                            v29 := GenS6{F0: 716, F1: 65535}
                            _ = v28
                            _ = v29
                        }
                    }
                }
            }
        } else {
            fmt.Println(v17)
            if false {
                for fi30 := 0; fi30 < 5; fi30++ {
                    switch "a" {
                    case "foo bar":
                        for fi31 := 0; fi31 < 20; fi31++ {
                            v32 := GenS4{F0: 1, F1: fi30, F2: -1}
                            var v33 int16 = -1
                            v17 = v17
                            _ = fi31
                            _ = v32
                            _ = v33
                        }
                    case "\n\t":
                        v18 = append(v18, 32768)
                    case "hello":
                        fmt.Println(fi30)
                    }
                    fmt.Println(fi30, v17)
                    fmt.Println(v17, fi30)
                    _ = fi30
                }
                switch "hello" {
                case "foo bar":
                    v17 = 317
                }
                v18 = append(v18, -662)
            }
        }
        _ = v19
        _ = v20
    }
    _ = v17
    _ = v18
    return 413
}

func main() {
    var v1 int = 128
    var v2 uint16 = 1
    v3 := make(map[int]bool)
    v4 := make(map[int]string)
    v5 := make(map[int]int)
    v6 := make(map[int]bool)
    v1 = (v1 / ((127) | 1))
    var v34 string = "foo bar"
    v35 := GenS2{F0: 127, F1: 42, F2: 0}
    v34 = (("" + v34) + (("a" + "") + ""))
    fmt.Println(v34)
    v34 = ""
    fmt.Println(v34, v1, v1)
    fmt.Println(v34, v1)
    var v36 int32 = 2147483647
    v37 := GenS6{F0: 2, F1: v2}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
