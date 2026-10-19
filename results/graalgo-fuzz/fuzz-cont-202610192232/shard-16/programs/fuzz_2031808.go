package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int8 }
type GenS2 struct { F0 uint8; F1 uint64 }

func main() {
    v1 := []uint32{899, 42, 2, 307}
    v2 := make(map[int]uint32)
    var v3 string = "foo bar"
    var v4 bool = false
    v5 := make(map[string]int)
    var v6 bool = (65535 < 255)
    if (v3 > "foo bar") {
        for fi7 := 0; fi7 < 20; fi7++ {
            if (v6 && (966 < -1)) {
                for fi8 := 0; fi8 < 11; fi8++ {
                    if (!((true || v6) && true)) {
                        v3 = ((v3 + v3) + v3)
                        if (v3 != v3) {
                            var v9 int8 = -68
                            _ = v9
                        } else {
                            v3 = ("\n\t" + v3)
                        }
                    } else {
                        fmt.Println(v6, fi7, fi8)
                        var v10 uint16 = 1
                        _ = v10
                    }
                    _ = fi8
                }
            } else {
                v2[(fi7 | (^(fi7)))] = 1
                fmt.Println(v6)
                v3 = v3
            }
            _ = len(v1)
            var v11 string = v3
            _ = fi7
            _ = v11
        }
    } else {
        v6 = v6
    }
    v3 = ("a" + "")
    fmt.Println(v4, v4, v3)
    for fi12 := 0; fi12 < 8; fi12++ {
        v2[(fi12 >> ((42) & 63))] = 0
        if (553 > 332) {
            var v13 uint16 = 798
            var v14 int = fi12
            v13 = (v13 ^ 255)
            _ = v13
            _ = v14
        } else {
            v1 = append(v1, 2)
            if v6 {
                if true {
                    v4 = false
                    var v15 int16 = 144
                    for fi16 := 0; fi16 < 8; fi16++ {
                        var v17 uint = 42
                        var v18 bool = v4
                        switch (v3 + ("a" + (v3 + (v3 + ("foo bar" + v3))))) {
                        case "":
                            v19 := GenS1{F0: (fi16 | fi12), F1: uint16(v17), F2: -1}
                            _ = v19
                        case "a":
                            v4 = v4
                        case "foo bar":
                            if v6 {
                                switch v3 {
                                case "a":
                                    v4 = (896 != 2147483648)
                                default:
                                    v20 := GenS1{F0: 0, F1: 1, F2: -8}
                                    _ = v20
                                }
                                switch fi12 {
                                case 107:
                                    if len(v1) > 0 {
                                        v1[0] = 0
                                    }
                                case 32768:
                                    fmt.Println(v15)
                                case 4294967295:
                                    if true {
                                        v18 = true
                                    }
                                default:
                                    v21 := GenS2{F0: 2, F1: 2}
                                    _ = v21
                                }
                                fmt.Println(fi12, v18, v4)
                                fmt.Println(v3, v17, v18)
                                var v22 bool = true
                                _ = v22
                            }
                        default:
                            var v23 uint = v17
                            _ = v23
                        }
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                    _ = v15
                }
            }
            var v24 string = ""
            _ = v24
        }
        v3 = v3
        _ = fi12
    }
    _ = len(v1)
    for fi25 := 0; fi25 < 11; fi25++ {
        if v4 {
            if v6 {
                _ = len(v1)
            }
            v6 = (33 <= 22)
        }
        v6 = v6
        switch ("hello" + v3) {
        case "hello":
            fmt.Println(v3, fi25)
        case "a":
            var v26 float32 = -2.25
            _ = v26
        }
        _ = fi25
    }
    var v27 uint32 = 65535
    var v28 uint8 = 0
    var v29 string = v3
    var v30 float64 = 1.5
    var v31 uint32 = 65535
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
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
    _ = v6
}
