package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 int64; F1 uint8; F2 uint8 }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 int8; F1 int32; F2 uint64 }

func main() {
    var v1 uint8 = 42
    var v2 float32 = 1e10
    v3 := []int{127}
    var v4 int = 128
    var v5 int = (v4 >> ((255) & 63))
    v6 := make(map[string]int)
    for fi7 := 0; fi7 < 18; fi7++ {
        fmt.Println(fi7)
        for fi8 := 0; fi8 < 17; fi8++ {
            var v9 int32 = 729
            if false {
                v10 := GenS1{F0: v1, F1: 793}
                if true {
                    var v11 bool = (1e10 > 3.14159)
                    v12 := GenS2{F0: 42, F1: 0, F2: v1}
                    _ = v11
                    _ = v12
                } else {
                    v13 := GenS1{F0: (^(v1)), F1: 432}
                    var v14 uint32 = 42
                    _ = v13
                    _ = v14
                }
                var v15 uint16 = 2
                if false {
                    switch fi8 {
                    case 9223372036854775807:
                        v16 := GenS2{F0: 42, F1: v1, F2: v1}
                        v1 = (v1 / (((v1 - (v1 - 0))) | 1))
                        _ = v16
                    case -2147483648:
                        var v17 bool = (!(false && true))
                        _ = v17
                    case -302:
                        v18 := v6[("hello" + "a")]
                        _ = v18
                    default:
                        v4 -= 0
                    }
                    fmt.Println(fi8, fi7)
                    switch "hello" {
                    case "":
                        if ((v4 << ((918) & 63)) == fi7) {
                            var v19 float64 = 1e10
                            var v20 uint = 255
                            v6[("" + ("\n\t" + ""))] = int(v1)
                            _ = v19
                            _ = v20
                        } else {
                            v21 := GenS1{F0: 1, F1: 1}
                            _ = v21
                        }
                    case "\n\t":
                        v6["\n\t"] = v4
                    case "foo bar":
                        var v22 uint16 = v15
                        _ = v22
                    }
                } else {
                    v5 -= (fi7 >> ((0) & 63))
                    v3 = append(v3, (fi7 + (v5 ^ (v5 / ((4294967295) | 1)))))
                }
                v23 := GenS1{F0: v1, F1: 70}
                var v24 uint64 = 260
                var v25 int16 = -1
                _ = v10
                _ = v15
                _ = v23
                _ = v24
                _ = v25
            } else {
                if false {
                    v4 += 42
                }
                var v26 float32 = v2
                _ = v26
            }
            switch (v4 / ((fi8) | 1)) {
            case 4294967295:
                if (-1 <= 9223372036854775807) {
                    if (2147483648 == 69) {
                        var v27 uint32 = 0
                        fmt.Println(fi7, fi7, v5)
                        _ = v27
                    }
                } else {
                    var v28 int8 = 1
                    _ = v28
                }
            case 32768:
                v5 -= (fi7 >> ((0) & 63))
            default:
                v2 += 0.1
            }
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v3 = append(v3, v5)
    if true {
        v4 = v5
        v29 := GenS3{F0: v5, F1: (v1 - 2)}
        _ = v29
    }
    if true {
        if ((963 <= 42) || false) {
            _ = len(v6)
        } else {
            var v30 int8 = 26
            _ = v30
        }
    }
    switch (("a" + "hello") + ("" + "\n\t")) {
    case "a":
        v31 := GenS3{F0: (v4 & v5), F1: v1}
        _ = v31
    case "foo bar":
        v1 = 0
    case "\n\t":
        if len(v3) > 0 {
            v3[0] = 32768
        }
    }
    v1 -= (v1 / (((v1 - (^(v1)))) | 1))
    switch ("" + (("hello" + ("\n\t" + ("" + ("foo bar" + (("foo bar" + "a") + (("\n\t" + ("a" + "a")) + "")))))) + "")) {
    case "hello":
        switch v4 {
        case 760:
            var v32 uint64 = 770
            _ = v32
        case -987:
            var v33 bool = (1 > 4294967295)
            _ = v33
        case 4294967295:
            v34 := GenS4{F0: 100, F1: 1, F2: 4294967295}
            _ = v34
        default:
            switch (v4 - v4) {
            case 717:
                if (42 != 32767) {
                    if (false || (false || false)) {
                        _ = len(v6)
                    }
                } else {
                    v4 = 127
                    if true {
                        fmt.Println(v5, v2, v4)
                    }
                    var v35 int32 = 42
                    _ = v35
                }
            case 32767:
                fmt.Println(v1)
            default:
                var v36 string = ("\n\t" + "foo bar")
                _ = v36
            }
        }
    case "":
        v4 -= -392
    case "a":
        v2 -= 3.14159
    default:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v37 := GenS3{F0: v5, F1: 167}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
