package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int; F1 int }
type GenS5 struct { F0 uint16; F1 int16 }

func main() {
    var v1 uint = 2
    var v2 bool = ((v1 >> ((0) & 63)) != v1)
    var v3 uint32 = 2147483648
    v4 := make(map[string]int)
    var v5 int16 = 42
    v4["a"] = 899
    delete(v4, "")
    if (v2 && ((!(18446744073709551615 == 864)) && false)) {
        if (9223372036854775807 == 9223372036854775807) {
            switch "a" {
            case "hello":
                if false {
                    v5 -= (v5 >> (((^(v1))) & 63))
                    if v2 {
                        _ = len(v4)
                    }
                } else {
                    v3 = (^(v3))
                    fmt.Println(v3, v3)
                }
            }
            var v6 bool = (true && (v2 && (65 >= 220)))
            var v7 float32 = (3.14159 + 0.1)
            _ = v6
            _ = v7
        } else {
            switch 128 {
            case 1:
                switch 326 {
                case 42:
                    var v8 float32 = -2.25
                    _ = v8
                case 99:
                    v9 := GenS1{F0: 1, F1: "hello", F2: 576}
                    _ = v9
                default:
                    fmt.Println(v1, v1, v1)
                }
            }
        }
        v5 = v5
    }
    fmt.Println(v5, v1)
    var v10 int8 = -107
    v11 := GenS1{F0: 128, F1: "", F2: (v1 * 42)}
    v12 := GenS2{F0: v1, F1: 1}
    var v13 uint64 = 18446744073709551615
    switch ("hello" + ("foo bar" + "a")) {
    case "foo bar":
        fmt.Println(v13)
    case "a":
        _ = len(v4)
    default:
        v1 = (v1 + (v1 * v1))
    }
    if v2 {
        v3 += uint32(v1)
    }
    var v14 uint8 = 48
    v5 += 1000
    v15 := []int{2147483647, 438}
    v16 := GenS3{F0: -495}
    v15 = append(v15, 2147483647)
    for fi17 := 0; fi17 < 10; fi17++ {
        fmt.Println(v5, v1, v2)
        _ = fi17
    }
    if ((^(v1)) != v1) {
        if v2 {
            if ((-(v10)) < (v10 % (((v10 * (v10 ^ (v10 ^ v10)))) | 1))) {
                v3 = 2147483647
                v18 := GenS1{F0: uint8(v1), F1: "", F2: 42}
                _ = v18
            }
            var v19 uint8 = v14
            var v20 uint = (^(v1))
            _ = v19
            _ = v20
        } else {
            v21 := GenS2{F0: (v1 - (^(v1))), F1: 42}
            if v2 {
                switch int(v1) {
                case 42:
                    if (42 != -452) {
                        var v22 int8 = (v10 / ((0) | 1))
                        _ = v22
                    }
                case -170:
                    fmt.Println(v1, v3)
                case 32768:
                    for fi23 := 0; fi23 < 13; fi23++ {
                        switch "" {
                        case "a":
                            _ = len(v15)
                        }
                        var v24 uint8 = 2
                        switch fi23 {
                        case 4294967295:
                            v25 := GenS3{F0: -9223372036854775808}
                            _ = v25
                        case 127:
                            var v26 uint8 = v14
                            _ = v26
                        case 32767:
                            var v27 float64 = 0.1
                            _ = v27
                        }
                        _ = fi23
                        _ = v24
                    }
                }
                v28 := GenS4{F0: -119, F1: 122}
                _ = v28
            } else {
                var v29 int64 = 9223372036854775807
                _ = v29
            }
            _ = v21
        }
        v30 := GenS4{F0: 32767, F1: 4294967295}
        fmt.Println(v10, v1)
        _ = v30
    } else {
        var v31 float64 = 3.14159
        fmt.Println(v2, v3, v2)
        if false {
            var v32 bool = ((4294967295 == 0) || (!v2))
            v2 = v32
            for fi33 := 0; fi33 < 19; fi33++ {
                var v34 uint8 = (v14 << ((v1) & 63))
                fmt.Println(v13, v31)
                _ = fi33
                _ = v34
            }
            _ = v32
        } else {
            if false {
                v35 := GenS4{F0: -666, F1: 101}
                v36 := GenS3{F0: -998}
                var v37 uint8 = (^(v14))
                _ = v35
                _ = v36
                _ = v37
            }
            var v38 int16 = 32767
            if false {
                v39 := GenS1{F0: (v14 >> (((v1 | (v1 & (v1 + (v1 * (v1 & v1)))))) & 63)), F1: "a", F2: 508}
                for fi40 := 0; fi40 < 1; fi40++ {
                    var v41 string = (((("hello" + "") + "hello") + ("hello" + "\n\t")) + ((("hello" + "a") + (("a" + "\n\t") + "\n\t")) + "foo bar"))
                    v42 := GenS5{F0: 151, F1: int16(v10)}
                    _ = fi40
                    _ = v41
                    _ = v42
                }
                _ = v39
            }
            _ = v38
        }
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
