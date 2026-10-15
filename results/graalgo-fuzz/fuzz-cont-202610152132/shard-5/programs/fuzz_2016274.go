package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint32 }
func genFn0(p0 int, p1 uint8) uint8 {
    var v6 uint64 = 4294967295
    p0 += int(p1)
    _ = p0
    _ = p1
    _ = v6
    return 127
}

func main() {
    v1 := []uint32{2, 2147483647}
    var v2 string = "hello"
    v3 := []int{-2147483648, 32767, -21}
    var v4 float64 = 1e10
    var v5 uint = 0
    switch ((((v2 + v2) + "") + "") + v2) {
    case "":
        v7 := GenS1{F0: 984}
        _ = v7
    case "hello":
        for fi8 := 0; fi8 < 14; fi8++ {
            fmt.Println(v5)
            var v9 bool = false
            _ = fi8
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        for fi11 := 0; fi11 < 5; fi11++ {
            switch v2 {
            case "a":
                if (!true) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                } else {
                    if true {
                        fmt.Println(v5)
                        fmt.Println(v5, fi11)
                    } else {
                        switch ((("a" + ("hello" + v2)) + "hello") + "\n\t") {
                        case "hello":
                            v12 := GenS2{F0: 261, F1: 0}
                            _ = v12
                        case "foo bar":
                            if (false || (true && (!true))) {
                                v2 = "a"
                                var v13 float64 = 0.1
                                _ = v13
                            } else {
                                if ((1.5 == float32(((v4 - v4) / v4))) || false) {
                                    v5 += v5
                                }
                                v14 := GenS2{F0: (fi10 % (((fi10 << ((18446744073709551615) & 63))) | 1)), F1: 127}
                                _ = v14
                            }
                        default:
                            v15 := GenS3{F0: (v2 + "")}
                            _ = v15
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                }
            case "foo bar":
                _ = len(v1)
            }
            var v16 uint64 = 2
            if (false && true) {
                var v17 bool = ((false && true) && ((1 < 65535) && false))
                v5 -= v5
                _ = v17
            }
            _ = fi11
            _ = v16
        }
        v5 -= (v5 - 18446744073709551615)
        _ = fi10
    }
    v4 = v4
    v18 := GenS3{F0: v2}
    var v19 int8 = -1
    v4 -= (0.1 * v4)
    if true {
        var v20 int64 = 1
        _ = v20
    } else {
        for fi21 := 0; fi21 < 10; fi21++ {
            var v22 bool = true
            var v23 bool = false
            var v24 uint64 = 570
            _ = fi21
            _ = v22
            _ = v23
            _ = v24
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v19)
    v25 := GenS3{F0: ("\n\t" + ((("foo bar" + "a") + v2) + v2))}
    if len(v1) > 0 {
        v1[0] = 2
    }
    v26 := GenS4{F0: 928}
    fmt.Println(v4, v5, v4)
    fmt.Println(v2, v5)
    if true {
        if len(v1) > 0 {
            v1[0] = 2
        }
        var v27 int = 42
        _ = v27
    } else {
        if false {
            var v28 uint32 = 42
            switch (v2 + v2) {
            case "":
                for fi29 := 0; fi29 < 12; fi29++ {
                    var v30 uint32 = (v28 % ((2147483648) | 1))
                    _ = fi29
                    _ = v30
                }
            case "\n\t":
                if false {
                    fmt.Println(v5, v2)
                    for fi31 := 0; fi31 < 17; fi31++ {
                        for fi32 := 0; fi32 < 8; fi32++ {
                            fmt.Println(fi32, v5)
                            v3 = append(v3, fi31)
                            _ = fi32
                        }
                        var v33 float32 = float32(0.1)
                        _ = fi31
                        _ = v33
                    }
                }
            case "foo bar":
                v28 = 65535
            }
            switch "hello" {
            case "":
                var v34 uint16 = 255
                _ = v34
            case "\n\t":
                _ = len(v3)
            case "foo bar":
                fmt.Println(v28)
            default:
                for fi35 := 0; fi35 < 14; fi35++ {
                    var v36 int8 = (^(v19))
                    if len(v1) > 0 {
                        v1[0] = v28
                    }
                    v3 = append(v3, (fi35 & fi35))
                    _ = fi35
                    _ = v36
                }
            }
            _ = v28
        }
        var v37 float32 = float32(3.14159)
        fmt.Println(v19)
        _ = v37
    }
    v19 += (v19 & int8(v5))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
