package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 int; F1 bool; F2 int8 }
type GenS5 struct { F0 uint; F1 uint8 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 2
    v3 := make(map[string]uint32)
    var v4 uint8 = v2
    v5 := make(map[int]bool)
    var v6 bool = ((("\n\t" + "") + "hello") != ("" + ("" + (("\n\t" + "") + "\n\t"))))
    switch "a" {
    case "hello":
        fmt.Println(v2, v4)
    default:
        v6 = (-246 <= 718)
    }
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := GenS1{F0: 18446744073709551615, F1: 574}
        delete(v3, "a")
        _ = fi7
        _ = v8
    }
    if true {
        if (1 != 42) {
            v4 -= (^(v4))
            v9 := []int{1, -9223372036854775808}
            var v10 string = "\n\t"
            _ = v10
            _ = v9
        }
        v4 = (v4 + 1)
        if v6 {
            if (0.0 >= 0.0) {
                switch 127 {
                case 32768:
                    v11 := GenS2{F0: 127, F1: -970}
                    _ = v11
                }
            } else {
                v4 = v4
                var v12 float32 = float32(3.14159)
                _ = v12
            }
            if ((1.5 != 1.5) && (1 < 2)) {
                fmt.Println(v4, v2)
                v13 := GenS2{F0: 100, F1: 659}
                for fi14 := 0; fi14 < 4; fi14++ {
                    v2 -= 42
                    _ = fi14
                }
                _ = v13
            }
        }
    } else {
        v2 += (v4 - v2)
        fmt.Println(v6, v2)
    }
    var v15 uint32 = uint32(v2)
    for fi16 := 0; fi16 < 17; fi16++ {
        for fi17 := 0; fi17 < 4; fi17++ {
            v6 = v6
            fmt.Println(fi17, v2, v2)
            fmt.Println(fi16, v2)
            _ = fi17
        }
        fmt.Println(v15, v2)
        _ = fi16
    }
    for fi18 := 0; fi18 < 9; fi18++ {
        v2 += (^(v4))
        var v19 float64 = (1e10 * (1.5 - ((0.0 / (1e10 / (0.1 / (((0.1 + 0.0) * (3.14159 / 1.5)) / (1e10 - (3.14159 / 1e10)))))) + (0.1 * (1e10 / 1e10)))))
        switch fi18 {
        case -2147483648:
            for fi20 := 0; fi20 < 6; fi20++ {
                var v21 int8 = 0
                v19 = 0.0
                fmt.Println(v21)
                if true {
                    switch fi18 {
                    case 32768:
                        if (((v2 >> ((42) & 63)) == v2) || ("" <= ("foo bar" + (("" + "\n\t") + "\n\t")))) {
                            v22 := GenS3{F0: (v2 << ((255) & 63)), F1: (v15 - 2147483647)}
                            fmt.Println(v4, v4)
                            var v23 int64 = -9223372036854775808
                            _ = v22
                            _ = v23
                        }
                    case 0:
                        v24 := GenS4{F0: (-(fi18)), F1: ("\n\t" != "hello"), F2: (v21 >> ((1) & 63))}
                        _ = v24
                    case 197:
                        var v25 uint8 = 128
                        _ = v25
                    default:
                        v26 := GenS2{F0: v21, F1: 9223372036854775807}
                        _ = v26
                    }
                    v2 -= (v2 % (((v2 | v2)) | 1))
                    v6 = true
                }
                _ = fi20
                _ = v21
            }
        case 32767:
            v27 := []int{}
            _ = v27
        case 2147483647:
            if v6 {
                v15 = v15
                switch fi18 {
                case 9223372036854775807:
                    v28 := GenS5{F0: 2, F1: v2}
                    _ = v28
                case 32768:
                    if (v19 != v19) {
                        fmt.Println(v6)
                        v2 += (v2 ^ v4)
                    }
                }
            }
        default:
            _ = len(v1)
        }
        _ = fi18
        _ = v19
    }
    if false {
        v29 := GenS1{F0: 458, F1: -2147483648}
        _ = v29
    } else {
        v4 += v2
        v30 := GenS2{F0: int8(v15), F1: 646}
        _ = v30
    }
    if (!(!(475 == 18446744073709551615))) {
        for fi31 := 0; fi31 < 17; fi31++ {
            v32 := GenS3{F0: v4, F1: (v15 ^ (v15 * v15))}
            _ = fi31
            _ = v32
        }
        var v33 uint64 = 1
        _ = v33
    }
    var v34 float64 = 0.0
    if v6 {
        v35 := GenS1{F0: 0, F1: -2147483648}
        fmt.Println(v4, v15)
        _ = v35
    }
    for fi36 := 0; fi36 < 14; fi36++ {
        _ = len(v3)
        _ = fi36
    }
    var v37 uint = 873
    fmt.Println(v2, v6)
    for fi38 := 0; fi38 < 10; fi38++ {
        var v39 uint64 = 42
        var v40 uint32 = v15
        v4 += 127
        _ = fi38
        _ = v39
        _ = v40
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v34
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
