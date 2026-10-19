package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 int; F2 int }

func main() {
    v1 := []uint32{688, 2147483648}
    var v2 bool = true
    var v3 string = (("hello" + "a") + (("a" + ("foo bar" + "hello")) + "foo bar"))
    v4 := []int{127, 55, 278, 1}
    if v2 {
        v3 = "foo bar"
    } else {
        fmt.Println(v3, v3)
        for fi5 := 0; fi5 < 9; fi5++ {
            var v6 int32 = -714
            v7 := GenS1{F0: -286}
            _ = fi5
            _ = v6
            _ = v7
        }
    }
    var v8 uint16 = 65535
    var v9 uint16 = (^(v8))
    fmt.Println(v2)
    var v10 float64 = 1e10
    _ = len(v4)
    var v11 uint16 = 255
    if (65535 <= 887) {
        var v12 uint64 = 2
        _ = v12
    } else {
        var v13 int8 = -21
        switch 32767 {
        case 707:
            var v14 string = (v3 + (v3 + ("a" + "foo bar")))
            _ = v14
        case -9223372036854775808:
            var v15 float64 = (((v10 / v10) - v10) - 1.5)
            _ = v15
        case -1:
            _ = len(v1)
        default:
            if v2 {
                var v16 float32 = 1e10
                var v17 float32 = (0.1 + v16)
                _ = v16
                _ = v17
            }
        }
        var v18 float32 = 3.14159
        _ = v13
        _ = v18
    }
    for fi19 := 0; fi19 < 14; fi19++ {
        if (-469 > -2) {
            if v2 {
                if (v2 && v2) {
                    if (-9223372036854775808 < 0) {
                        v20 := GenS1{F0: 1}
                        _ = v20
                    }
                    switch v3 {
                    case "\n\t":
                        if (260 >= -1) {
                            _ = len(v1)
                        }
                    }
                }
                if ((74 == 42) && (fi19 != fi19)) {
                    var v21 uint16 = v11
                    _ = v21
                } else {
                    for fi22 := 0; fi22 < 10; fi22++ {
                        fmt.Println(v9, v2)
                        v4 = append(v4, (fi19 - fi22))
                        if len(v1) > 0 {
                            v1[0] = 615
                        }
                        _ = fi22
                    }
                    if (-928 == 42) {
                        v9 -= v11
                        var v23 int = (fi19 | 2147483647)
                        var v24 bool = true
                        _ = v23
                        _ = v24
                    } else {
                        if false {
                            switch (fi19 | fi19) {
                            case 0:
                                for fi25 := 0; fi25 < 20; fi25++ {
                                    v3 = v3
                                    switch 42 {
                                    case 32768:
                                        _ = len(v4)
                                    default:
                                        if ("foo bar" <= v3) {
                                            var v26 string = v3
                                            var v27 uint16 = (^(v9))
                                            v28 := GenS1{F0: 1}
                                            v2 = false
                                            _ = v26
                                            _ = v27
                                            _ = v28
                                        }
                                    }
                                    switch "a" {
                                    case "a":
                                        v29 := GenS2{F0: ("a" + ("" + v3))}
                                        if (!v2) {
                                            v11 = (^(v8))
                                            if len(v1) > 0 {
                                                v1[0] = 2147483647
                                            }
                                            var v30 float64 = 0.0
                                            _ = v30
                                        }
                                        _ = v29
                                    case "\n\t":
                                        var v31 int16 = 32767
                                        _ = v31
                                    }
                                    _ = fi25
                                }
                            }
                            var v32 uint16 = 720
                            var v33 int32 = 2147483647
                            _ = v32
                            _ = v33
                        } else {
                            v34 := GenS3{F0: 32767, F1: 114, F2: (fi19 * 548)}
                            if len(v1) > 0 {
                                v1[0] = 4294967295
                            }
                            for fi35 := 0; fi35 < 16; fi35++ {
                                if len(v4) > 0 {
                                    v4[0] = fi35
                                }
                                if len(v4) > 0 {
                                    v4[0] = (fi19 - (^(fi35)))
                                }
                                _ = fi35
                            }
                            _ = v34
                        }
                        fmt.Println(v9, v9, v2)
                        for fi36 := 0; fi36 < 13; fi36++ {
                            v3 = v3
                            var v37 uint32 = 0
                            _ = fi36
                            _ = v37
                        }
                    }
                    switch -9223372036854775808 {
                    case 652:
                        var v38 uint16 = (^(v9))
                        _ = v38
                    }
                }
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v2 = true
        }
        _ = fi19
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
