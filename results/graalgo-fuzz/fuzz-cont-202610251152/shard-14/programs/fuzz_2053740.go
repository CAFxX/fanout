package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 int32 = 350
    var v2 uint16 = 2
    var v3 int16 = -32768
    var v4 bool = (true || (false || true))
    v5 := GenS1{F0: 516, F1: v3}
    var v6 int = 128
    var v7 uint64 = 42
    switch "" {
    case "\n\t":
        v8 := GenS1{F0: (^(v6)), F1: v3}
        _ = v8
    case "a":
        if (((v7 >> ((59) & 63)) <= 2) && (v7 >= v7)) {
            for fi9 := 0; fi9 < 20; fi9++ {
                if (0 <= 42) {
                    if (128 == (fi9 % (((v6 / ((fi9) | 1))) | 1))) {
                        if (0 < v6) {
                            v10 := GenS1{F0: 32767, F1: 253}
                            v1 = 893
                            _ = v10
                        } else {
                            for fi11 := 0; fi11 < 1; fi11++ {
                                var v12 uint8 = 80
                                v13 := []int{9223372036854775807, -1}
                                v6 = (^(v6))
                                _ = fi11
                                _ = v12
                                _ = v13
                            }
                        }
                        fmt.Println(v1, v3)
                        if false {
                            for fi14 := 0; fi14 < 1; fi14++ {
                                switch v6 {
                                case 128:
                                    var v15 float64 = (3.14159 - 0.0)
                                    _ = v15
                                case -576:
                                    var v16 uint16 = v2
                                    _ = v16
                                case -1:
                                    fmt.Println(v3, fi14)
                                default:
                                    var v17 uint64 = (v7 >> ((4294967295) & 63))
                                    _ = v17
                                }
                                v6 = v6
                                _ = fi14
                            }
                        } else {
                            v3 = v3
                            var v18 int64 = 9223372036854775807
                            _ = v18
                        }
                    } else {
                        v19 := []int{(fi9 >> ((255) & 63)), v6, v6}
                        _ = v19
                    }
                } else {
                    for fi20 := 0; fi20 < 14; fi20++ {
                        v21 := []int{}
                        v3 = (-(v3))
                        _ = fi20
                        _ = v21
                    }
                }
                var v22 uint32 = 65535
                switch "hello" {
                case "hello":
                    fmt.Println(v7)
                case "":
                    if v4 {
                        v23 := []int{v6}
                        var v24 float64 = (-2.25 - -2.25)
                        _ = v23
                        _ = v24
                    } else {
                        v25 := []int{int(v7), (fi9 & (fi9 % (((fi9 + int(v22))) | 1))), (v6 + (^(v6))), (v6 & (v6 / ((v6) | 1)))}
                        var v26 uint = 0
                        _ = v25
                        _ = v26
                    }
                }
                _ = fi9
                _ = v22
            }
            switch ("\n\t" + ("foo bar" + "hello")) {
            case "\n\t":
                var v27 uint = 65535
                _ = v27
            }
            v28 := []int{v6}
            _ = v28
        } else {
            v3 = (v3 / (((v3 ^ 32767)) | 1))
            v4 = (291 != 1)
            var v29 uint8 = 0
            _ = v29
        }
    case "foo bar":
        fmt.Println(v7, v6, v2)
    default:
        switch v6 {
        case 669:
            var v30 int32 = int32(v2)
            _ = v30
        case 127:
            switch v6 {
            case 2147483647:
                switch v6 {
                case 128:
                    var v31 uint16 = 825
                    _ = v31
                }
            case -1:
                v4 = true
            case 409:
                if ((!("\n\t" == ("a" + "a"))) && v4) {
                    v4 = false
                } else {
                    v4 = false
                    var v32 uint32 = 138
                    for fi33 := 0; fi33 < 5; fi33++ {
                        for fi34 := 0; fi34 < 18; fi34++ {
                            var v35 uint32 = 2147483648
                            v36 := GenS1{F0: fi34, F1: v3}
                            v4 = ("foo bar" == ("" + "a"))
                            _ = fi34
                            _ = v35
                            _ = v36
                        }
                        var v37 uint8 = 127
                        _ = fi33
                        _ = v37
                    }
                    _ = v32
                }
            }
        default:
            for fi38 := 0; fi38 < 12; fi38++ {
                if (!v4) {
                    fmt.Println(v7, v7)
                    var v39 bool = (9223372036854775807 >= -9223372036854775808)
                    _ = v39
                }
                v40 := GenS2{F0: 42}
                fmt.Println(v6)
                _ = fi38
                _ = v40
            }
        }
    }
    var v41 float64 = float64(float32(1e10))
    var v42 int64 = 765
    if v4 {
        for fi43 := 0; fi43 < 11; fi43++ {
            var v44 int8 = 0
            _ = fi43
            _ = v44
        }
        v45 := []int{(^(v6)), (v6 >> ((777) & 63)), v6, v6}
        v45 = append(v45, (v6 & v6))
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
    _ = v7
}
