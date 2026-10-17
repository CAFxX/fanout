package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 uint64 }
type GenS2 struct { F0 uint8; F1 int; F2 int }

func main() {
    var v1 uint8 = 255
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]bool)
    v4 := make(map[int]bool)
    for fi5 := 0; fi5 < 4; fi5++ {
        switch (fi5 * int(v1)) {
        case -191:
            var v6 uint32 = 1
            _ = v6
        case 438:
            var v7 string = "\n\t"
            _ = v7
        case -1:
            var v8 uint16 = 42
            _ = v8
        default:
            var v9 int = fi5
            _ = v9
        }
        _ = fi5
    }
    switch "a" {
    case "hello":
        fmt.Println(v1)
    case "\n\t":
        switch -167 {
        case 1:
            switch 32768 {
            case -9223372036854775808:
                fmt.Println(v1)
            case 128:
                var v10 uint16 = 2
                _ = v10
            case 127:
                fmt.Println(v1)
            }
        default:
            var v11 float64 = -2.25
            _ = v11
        }
    }
    _ = len(v3)
    v1 = v1
    if false {
        fmt.Println(v1)
        switch "a" {
        case "\n\t":
            for fi12 := 0; fi12 < 20; fi12++ {
                v1 = (v1 & 228)
                for fi13 := 0; fi13 < 20; fi13++ {
                    v1 += (v1 + v1)
                    var v14 int8 = 20
                    var v15 uint16 = 2
                    var v16 int64 = 1
                    _ = fi13
                    _ = v14
                    _ = v15
                    _ = v16
                }
                fmt.Println(v1, v1)
                _ = fi12
            }
        }
    } else {
        if false {
            fmt.Println(v1)
            switch ((((("hello" + "\n\t") + "\n\t") + ("foo bar" + "hello")) + ("a" + "")) + "foo bar") {
            case "hello":
                switch ("foo bar" + ("hello" + ("hello" + (("hello" + ("hello" + ("foo bar" + "foo bar"))) + "foo bar")))) {
                case "foo bar":
                    v17 := []int{}
                    _ = v17
                default:
                    v18 := v2[2147483648]
                    _ = v18
                }
            case "a":
                v1 -= (v1 << ((0) & 63))
            case "foo bar":
                for fi19 := 0; fi19 < 11; fi19++ {
                    v20 := []int{-2147483648, (fi19 % (((fi19 & 9223372036854775807)) | 1)), fi19, (fi19 * (fi19 << ((42) & 63)))}
                    v1 = v1
                    if false {
                        v1 += 2
                        v21 := GenS1{F0: -32768, F1: 42, F2: 1}
                        if (100 == 0) {
                            for fi22 := 0; fi22 < 16; fi22++ {
                                var v23 uint8 = (v1 + 42)
                                var v24 float32 = ((3.14159 / 3.14159) * 0.0)
                                fmt.Println(v24, v23)
                                _ = len(v20)
                                _ = fi22
                                _ = v23
                                _ = v24
                            }
                        } else {
                            for fi25 := 0; fi25 < 2; fi25++ {
                                var v26 float32 = (1e10 / 3.14159)
                                var v27 int8 = 64
                                _ = fi25
                                _ = v26
                                _ = v27
                            }
                            fmt.Println(fi19, fi19)
                            for fi28 := 0; fi28 < 17; fi28++ {
                                var v29 string = "hello"
                                _ = fi28
                                _ = v29
                            }
                        }
                        _ = v21
                    }
                    _ = fi19
                    _ = v20
                }
            default:
                switch "hello" {
                case "":
                    switch "" {
                    case "hello":
                        if (1 >= -1) {
                            v30 := []int{9223372036854775807, 127, 0}
                            v1 -= 0
                            v3[50] = true
                            _ = v30
                        } else {
                            var v31 int64 = 179
                            _ = v31
                        }
                    default:
                        for fi32 := 0; fi32 < 18; fi32++ {
                            v33 := GenS2{F0: v1, F1: (fi32 * (fi32 * fi32)), F2: (fi32 << ((65535) & 63))}
                            fmt.Println(fi32)
                            _ = fi32
                            _ = v33
                        }
                    }
                case "hello":
                    fmt.Println(v1)
                }
            }
        } else {
            for fi34 := 0; fi34 < 4; fi34++ {
                var v35 uint64 = 0
                _ = fi34
                _ = v35
            }
        }
    }
    if false {
        var v36 string = "hello"
        v37 := v3[65535]
        _ = v36
        _ = v37
    }
    delete(v4, 9223372036854775807)
    for fi38 := 0; fi38 < 8; fi38++ {
        v39 := []int{(fi38 ^ fi38)}
        for fi40 := 0; fi40 < 4; fi40++ {
            v39 = append(v39, fi40)
            fmt.Println(v1, v1, fi38)
            _ = fi40
        }
        _ = fi38
        _ = v39
    }
    for fi41 := 0; fi41 < 8; fi41++ {
        var v42 int32 = 2147483647
        v43 := GenS1{F0: int16(v1), F1: 175, F2: 444}
        _ = fi41
        _ = v42
        _ = v43
    }
    fmt.Println(v1)
    v44 := GenS1{F0: -52, F1: (v1 ^ (^(v1))), F2: 732}
    v1 -= 127
    v1 -= 1
    v45 := GenS2{F0: (v1 << ((4294967295) & 63)), F1: 726, F2: 16}
    var v46 uint32 = 826
    v46 += v46
    v1 = v1
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
}
