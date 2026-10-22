package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 int }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 string; F1 int32 }
type GenS4 struct { F0 string; F1 uint16; F2 uint8 }
type GenS5 struct { F0 uint8 }
func genFn0(p0 string) uint16 {
    var v4 uint64 = 414
    if false {
        v5 := []int{2147483647, -358}
        v6 := GenS1{F0: 424, F1: -1, F2: -102}
        if false {
            if false {
                p0 = p0
                if true {
                    for fi7 := 0; fi7 < 11; fi7++ {
                        if false {
                            v4 = (^(v4))
                            p0 = (p0 + "")
                            for fi8 := 0; fi8 < 16; fi8++ {
                                var v9 int8 = 100
                                _ = fi8
                                _ = v9
                            }
                        }
                        fmt.Println(v4)
                        _ = fi7
                    }
                    switch p0 {
                    case "hello":
                        p0 = p0
                    }
                    for fi10 := 0; fi10 < 20; fi10++ {
                        if len(v5) > 0 {
                            v5[0] = -9223372036854775808
                        }
                        v11 := GenS1{F0: 0, F1: 2147483647, F2: 2147483647}
                        _ = fi10
                        _ = v11
                    }
                } else {
                    fmt.Println(v4)
                    fmt.Println(v4, p0)
                }
            } else {
                if true {
                    _ = len(v5)
                    v12 := GenS1{F0: 908, F1: -651, F2: 127}
                    v4 += v4
                    _ = v12
                } else {
                    for fi13 := 0; fi13 < 13; fi13++ {
                        v4 = v4
                        for fi14 := 0; fi14 < 7; fi14++ {
                            v4 = v4
                            v5 = append(v5, fi14)
                            var v15 bool = (true && (false || (fi14 == 671)))
                            _ = fi14
                            _ = v15
                        }
                        var v16 bool = (int16(v4) <= 42)
                        _ = fi13
                        _ = v16
                    }
                    var v17 uint = 65535
                    _ = v17
                }
                var v18 float64 = (3.14159 - float64((1.5 / 1e10)))
                _ = v18
            }
        }
        _ = v5
        _ = v6
    }
    _ = p0
    _ = v4
    return 421
}

func main() {
    var v1 float64 = 1.5
    v2 := []int{712, -9223372036854775808, 30, -9223372036854775808}
    v3 := make(map[string]string)
    fmt.Println(v1)
    switch -1 {
    case 495:
        v1 = v1
    case 42:
        v19 := GenS2{F0: 0, F1: 32767}
        _ = v19
    }
    for fi20 := 0; fi20 < 11; fi20++ {
        switch "foo bar" {
        case "foo bar":
            v21 := GenS3{F0: (("\n\t" + ("hello" + "\n\t")) + "hello"), F1: 1}
            _ = v21
        }
        fmt.Println(fi20, fi20)
        v1 += v1
        _ = fi20
    }
    if (-32768 >= 1) {
        _ = len(v2)
        var v22 int16 = 98
        _ = v22
    }
    if false {
        if false {
            v1 -= (1e10 - v1)
            switch "" {
            case "\n\t":
                var v23 uint32 = 42
                _ = v23
            case "a":
                v2 = append(v2, 753)
            case "hello":
                for fi24 := 0; fi24 < 14; fi24++ {
                    var v25 uint64 = 0
                    v26 := GenS4{F0: "", F1: 42, F2: 2}
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            }
            switch 1 {
            case -51:
                v27 := GenS3{F0: "hello", F1: 913}
                _ = v27
            }
        }
        var v28 uint64 = 665
        v29 := GenS5{F0: 2}
        _ = v28
        _ = v29
    }
    if false {
        var v30 int32 = -588
        switch "foo bar" {
        case "":
            var v31 float64 = v1
            _ = v31
        case "\n\t":
            if true {
                for fi32 := 0; fi32 < 4; fi32++ {
                    var v33 int64 = -1
                    var v34 int16 = -328
                    for fi35 := 0; fi35 < 3; fi35++ {
                        if false {
                            v33 = (v33 % (((v33 ^ v33)) | 1))
                            _ = len(v2)
                        } else {
                            var v36 uint = 42
                            _ = v36
                        }
                        v34 -= (v34 / ((v34) | 1))
                        _ = fi35
                    }
                    _ = fi32
                    _ = v33
                    _ = v34
                }
            } else {
                fmt.Println(v1, v1)
            }
        case "hello":
            v30 = v30
        }
        fmt.Println(v30, v30)
        _ = v30
    }
    if (-2147483648 >= 2147483647) {
        v1 = ((float64(3.14159) - v1) + float64((1.5 / 3.14159)))
    }
    var v37 float32 = float32(3.14159)
    var v38 float64 = v1
    if (!false) {
        var v39 int8 = 94
        _ = v39
    } else {
        if len(v2) > 0 {
            v2[0] = 2147483647
        }
        v40 := GenS2{F0: -32768, F1: 0}
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
}
