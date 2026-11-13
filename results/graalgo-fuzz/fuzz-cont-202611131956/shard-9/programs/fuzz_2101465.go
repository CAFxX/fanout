package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8; F1 int }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int; F1 uint; F2 int16 }
type GenS6 struct { F0 bool; F1 uint8 }
type GenS7 struct { F0 uint32; F1 uint32; F2 string }
func genFn0(p0 int, p1 int) int32 {
    v6 := []int{p1, p0}
    v7 := GenS1{F0: 0, F1: 42}
    switch "\n\t" {
    case "foo bar":
        p0 += (p1 >> ((255) & 63))
    default:
        fmt.Println(p0)
    }
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 uint8 = 128
        switch p1 {
        case -9223372036854775808:
            v10 := GenS2{F0: 162}
            _ = v10
        case -480:
            var v11 int32 = 2147483647
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 94
}

func main() {
    var v1 float32 = -2.25
    v2 := []uint32{2147483647, 42}
    v3 := make(map[string]string)
    var v4 uint16 = 571
    v5 := make(map[uint32]uint32)
    v12 := GenS3{F0: 1, F1: 0}
    v13 := GenS1{F0: uint32(v4), F1: 1000}
    v14 := GenS2{F0: 42}
    var v15 int8 = -1
    var v16 int8 = v15
    var v17 bool = (0.0 >= 0.1)
    var v18 float64 = ((3.14159 / 1.0) - 0.1)
    var v19 uint8 = 0
    if v17 {
        for fi20 := 0; fi20 < 11; fi20++ {
            v21 := GenS4{F0: v19}
            if v17 {
                v19 += (v19 << ((18446744073709551615) & 63))
                v3["a"] = (("a" + "hello") + "")
                _ = len(v3)
            }
            _ = fi20
            _ = v21
        }
    } else {
        if v17 {
            v22 := GenS5{F0: 32768, F1: 2, F2: 1}
            switch 2147483647 {
            case -1:
                if false {
                    for fi23 := 0; fi23 < 20; fi23++ {
                        for fi24 := 0; fi24 < 13; fi24++ {
                            switch "\n\t" {
                            case "a":
                                v25 := v3[(((("hello" + ("a" + "")) + "hello") + "\n\t") + ("hello" + "a"))]
                                _ = v25
                            }
                            v18 -= (((float64((v1 - -2.25)) + 3.14159) * (1e10 / -2.25)) + 1e10)
                            _ = fi24
                        }
                        _ = fi23
                    }
                }
            case 4294967295:
                var v26 int8 = (v16 + (v16 * v16))
                _ = v26
            default:
                v15 -= (v16 | (v15 + 0))
            }
            _ = v22
        } else {
            for fi27 := 0; fi27 < 1; fi27++ {
                switch fi27 {
                case 32768:
                    if v17 {
                        for fi28 := 0; fi28 < 10; fi28++ {
                            v18 -= 1e10
                            switch "" {
                            case "hello":
                                switch "a" {
                                case "foo bar":
                                    v19 += v19
                                }
                            case "a":
                                _ = len(v2)
                                fmt.Println(v15, v19)
                                v19 = 127
                            }
                            var v29 int16 = 1
                            v29 += -446
                            _ = fi28
                            _ = v29
                        }
                        var v30 float64 = v18
                        if (!false) {
                            var v31 string = ("foo bar" + ("\n\t" + "\n\t"))
                            v30 += 1.5
                            _ = v31
                        } else {
                            switch "foo bar" {
                            case "a":
                                v18 = (1e10 - 3.14159)
                            case "foo bar":
                                _ = len(v2)
                            }
                            v4 = 65535
                        }
                        _ = v30
                    }
                case 323:
                    for fi32 := 0; fi32 < 8; fi32++ {
                        if (1e10 < v1) {
                            v2 = append(v2, 2)
                            switch "\n\t" {
                            case "":
                                if true {
                                    fmt.Println(fi27)
                                    if (42 >= v16) {
                                        v33 := GenS5{F0: fi27, F1: 1, F2: 42}
                                        _ = v33
                                    }
                                } else {
                                    if (v17 || v17) {
                                        if (v17 && (!((!false) && (v17 && (255 >= 94))))) {
                                            v15 += 42
                                            if (4294967295 > 385) {
                                                var v34 int32 = 1
                                                v19 += v19
                                                v1 = (v1 - (1.5 / (-2.25 - float32((v18 / -2.25)))))
                                                _ = v34
                                            } else {
                                                if len(v2) > 0 {
                                                    v2[0] = 0
                                                }
                                            }
                                            if (!(65535 == 255)) {
                                                fmt.Println(fi32, v4)
                                                if ("a" < "\n\t") {
                                                    v35 := GenS5{F0: fi32, F1: 969, F2: 0}
                                                    v36 := GenS6{F0: v17, F1: (^(v19))}
                                                    v2 = append(v2, 405)
                                                    _ = v35
                                                    _ = v36
                                                }
                                                v18 -= (v18 + (float64(v1) * 3.14159))
                                                var v37 uint64 = 1
                                                _ = v37
                                            }
                                            if (1e10 > float32((v18 / (v18 / (1e10 - v18))))) {
                                                fmt.Println(v17)
                                            }
                                            v16 = v16
                                        }
                                        var v38 uint = 18446744073709551615
                                        _ = v38
                                    } else {
                                        v19 = v19
                                        fmt.Println(v17, v15, v1)
                                    }
                                }
                            }
                        }
                        var v39 string = "hello"
                        fmt.Println(v18, v16, fi27)
                        _ = fi32
                        _ = v39
                    }
                case 812:
                    if v17 {
                        _ = len(v2)
                    }
                }
                _ = fi27
            }
            if false {
                var v40 uint16 = (v4 | (v4 >> ((272) & 63)))
                _ = v40
            }
            switch ("hello" + ("foo bar" + "a")) {
            case "":
                v41 := GenS6{F0: (true && (32767 == -51)), F1: v19}
                _ = v41
            }
        }
        v42 := GenS7{F0: 0, F1: 0, F2: "\n\t"}
        _ = v42
    }
    var v43 int32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
}
