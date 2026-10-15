package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 int16; F1 string }
func genFn0() string {
    if true {
        v6 := GenS1{F0: -32768, F1: -1}
        var v7 int8 = -64
        if true {
            switch "a" {
            case "a":
                if (!(true || false)) {
                    v8 := []int{-1}
                    var v9 uint = 65535
                    _ = v8
                    _ = v9
                }
            case "":
                if true {
                    if ((((-2.25 / (3.14159 * 1e10)) == -2.25) || (false || (1 >= 2))) || true) {
                        for fi10 := 0; fi10 < 4; fi10++ {
                            fmt.Println(fi10, v7)
                            v7 -= v7
                            _ = fi10
                        }
                    } else {
                        if true {
                            fmt.Println(v7)
                        } else {
                            for fi11 := 0; fi11 < 11; fi11++ {
                                if (!(-1 != 999)) {
                                    for fi12 := 0; fi12 < 4; fi12++ {
                                        var v13 int32 = -1
                                        _ = fi12
                                        _ = v13
                                    }
                                    if true {
                                        for fi14 := 0; fi14 < 19; fi14++ {
                                            v7 = (v7 << ((18446744073709551615) & 63))
                                            fmt.Println(fi14)
                                            _ = fi14
                                        }
                                    } else {
                                        switch "a" {
                                        case "hello":
                                            v7 += v7
                                        case "a":
                                            v15 := []int{}
                                            _ = v15
                                        case "foo bar":
                                            v7 = (v7 & (v7 << ((0) & 63)))
                                        }
                                    }
                                }
                                _ = fi11
                            }
                        }
                    }
                    for fi16 := 0; fi16 < 4; fi16++ {
                        v7 += -128
                        _ = fi16
                    }
                    if (-475 >= 44) {
                        v7 = v7
                    } else {
                        if ((true && (347 != 4294967295)) || (v7 <= (v7 ^ v7))) {
                            v7 = v7
                            v17 := GenS2{F0: (v7 / ((v7) | 1)), F1: 65535}
                            v18 := []int{}
                            _ = v17
                            _ = v18
                        }
                        fmt.Println(v7)
                    }
                }
            default:
                v7 = (v7 / ((v7) | 1))
            }
            v19 := []int{16}
            _ = v19
        } else {
            v7 -= (v7 / ((0) | 1))
            for fi20 := 0; fi20 < 7; fi20++ {
                v7 += -1
                _ = fi20
            }
            v21 := GenS1{F0: 218, F1: 33}
            _ = v21
        }
        _ = v6
        _ = v7
    }
    var v22 int32 = 1
    if false {
        for fi23 := 0; fi23 < 2; fi23++ {
            fmt.Println(fi23, fi23)
            v22 -= (^(v22))
            _ = fi23
        }
        for fi24 := 0; fi24 < 7; fi24++ {
            var v25 uint16 = 320
            _ = fi24
            _ = v25
        }
        var v26 string = "hello"
        _ = v26
    } else {
        fmt.Println(v22)
        v22 -= v22
    }
    fmt.Println(v22)
    fmt.Println(v22)
    _ = v22
    return "hello"
}

func main() {
    v1 := make(map[int]int)
    var v2 float32 = 1e10
    var v3 int64 = -1
    v4 := make(map[uint32]string)
    v5 := make(map[uint32]bool)
    var v27 uint16 = 602
    v28 := GenS2{F0: 127, F1: 2147483647}
    fmt.Println(v2)
    delete(v1, 1)
    switch 2147483647 {
    case 32767:
        v3 = int64(v27)
    }
    v3 = -1
    var v29 int32 = 42
    v3 += 42
    var v30 int = 2147483647
    v31 := GenS3{F0: -1, F1: ("a" + ("foo bar" + ("\n\t" + "a")))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
