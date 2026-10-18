package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 bool; F2 int32 }
type GenS3 struct { F0 int8 }
func genFn0(p0 uint32, p1 uint16) uint8 {
    var v7 int16 = 0
    switch 42 {
    case 32767:
        for fi8 := 0; fi8 < 15; fi8++ {
            v7 -= v7
            _ = fi8
        }
    }
    _ = p0
    _ = p1
    _ = v7
    return 255
}

func main() {
    var v1 string = ("" + ("foo bar" + "hello"))
    v2 := []string{v1, ((v1 + (v1 + "foo bar")) + v1)}
    v3 := []uint32{2147483648, 42, 42, 0}
    v4 := make(map[int]int)
    v5 := make(map[string]int)
    var v6 bool = (false && ((-2.25 / -2.25) >= 1.5))
    switch "" {
    case "foo bar":
        v6 = (0.1 <= 1e10)
    case "hello":
        fmt.Println(v1, v6)
    case "":
        var v9 uint64 = 1
        _ = v9
    }
    fmt.Println(v1, v6)
    switch v1 {
    case "":
        fmt.Println(v6)
    }
    if (true && v6) {
        var v10 float32 = (1e10 / 3.14159)
        _ = v10
    }
    v1 = (("\n\t" + ((("hello" + "hello") + "foo bar") + "")) + "")
    v11 := GenS1{F0: -2147483648}
    v12 := GenS2{F0: 128, F1: (100 == -1), F2: 171}
    for fi13 := 0; fi13 < 19; fi13++ {
        switch "\n\t" {
        case "\n\t":
            fmt.Println(v6, fi13)
        case "":
            v1 = v1
        case "foo bar":
            if false {
                if false {
                    var v14 string = "foo bar"
                    for fi15 := 0; fi15 < 16; fi15++ {
                        v6 = false
                        v1 = (((v1 + "") + "hello") + (v14 + v1))
                        v6 = v6
                        _ = fi15
                    }
                    var v16 int16 = 947
                    _ = v14
                    _ = v16
                } else {
                    for fi17 := 0; fi17 < 10; fi17++ {
                        var v18 bool = v6
                        _ = len(v2)
                        if v6 {
                            switch fi13 {
                            case 1:
                                v19 := GenS2{F0: 2, F1: v18, F2: 648}
                                _ = v19
                            case 84:
                                v20 := GenS2{F0: 0, F1: v6, F2: -926}
                                _ = v20
                            }
                            var v21 bool = true
                            _ = v21
                        } else {
                            v6 = (-1 <= 0)
                            v1 = v1
                            v1 = v1
                            v1 = ("\n\t" + ("\n\t" + v1))
                            v1 = ""
                        }
                        var v22 uint8 = 128
                        var v23 int = fi13
                        v22 -= (v22 ^ (v22 & genFn0(2, 207)))
                        fmt.Println(fi13, v6)
                        if v18 {
                            v24 := GenS3{F0: -1}
                            _ = v24
                        } else {
                            var v25 int16 = -135
                            var v26 uint = uint(v22)
                            _ = v25
                            _ = v26
                        }
                        if (4294967295 <= 906) {
                            v22 -= (^(v22))
                            var v27 int8 = -72
                            v18 = ((0 <= 621) || (-9223372036854775808 != 0))
                            v5["a"] = fi13
                            _ = v27
                        }
                        _ = fi17
                        _ = v18
                        _ = v22
                        _ = v23
                    }
                    if (-2.25 == -2.25) {
                        v6 = false
                        for fi28 := 0; fi28 < 3; fi28++ {
                            v29 := GenS2{F0: 2, F1: false, F2: -1}
                            switch v1 {
                            case "a":
                                v6 = false
                            case "hello":
                                v30 := GenS3{F0: 127}
                                _ = v30
                            case "":
                                v31 := GenS3{F0: -25}
                                _ = v31
                            }
                            var v32 int16 = -1
                            _ = fi28
                            _ = v29
                            _ = v32
                        }
                        var v33 uint64 = 2
                        _ = v33
                    } else {
                        if (!false) {
                            var v34 string = (((v1 + v1) + "a") + v1)
                            _ = v34
                        }
                        var v35 int16 = 1
                        _ = v35
                    }
                }
                v1 = (v1 + ((v1 + ((v1 + ((((v1 + v1) + "\n\t") + "foo bar") + v1)) + "\n\t")) + ("a" + "foo bar")))
            }
        }
        for fi36 := 0; fi36 < 11; fi36++ {
            fmt.Println(fi36, fi36)
            var v37 uint8 = 128
            _ = fi36
            _ = v37
        }
        v38 := GenS1{F0: (fi13 << ((18446744073709551615) & 63))}
        _ = fi13
        _ = v38
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
