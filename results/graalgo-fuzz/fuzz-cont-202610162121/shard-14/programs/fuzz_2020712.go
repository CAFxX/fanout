package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 uint16 }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 uint64; F1 uint64 }
type GenS5 struct { F0 uint; F1 int32; F2 int64 }
type GenS6 struct { F0 uint16; F1 uint64 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int16; F1 uint8; F2 int }
type GenS10 struct { F0 int64; F1 bool; F2 string }

func main() {
    var v1 int8 = 0
    v2 := make(map[uint32]int)
    var v3 int8 = (v1 / ((v1) | 1))
    var v4 int64 = 0
    v5 := []string{"hello", "\n\t", ""}
    var v6 int = -609
    var v7 uint16 = 65535
    var v8 int8 = -74
    switch (^(v6)) {
    case 32767:
        fmt.Println(v7)
    case 9223372036854775807:
        v7 = (^(v7))
    case -9223372036854775808:
        v1 = (v8 * v1)
    }
    var v9 uint = 65535
    v10 := v2[925]
    switch "" {
    case "hello":
        v6 = v6
    case "\n\t":
        if (float32(1e10) == ((1e10 - (0.1 + 0.1)) * (1e10 - (1.5 * (3.14159 + 0.1))))) {
            v10 = v10
            for fi11 := 0; fi11 < 18; fi11++ {
                v10 += (v10 ^ v10)
                v3 = v1
                if (float64((1.5 + 1e10)) == 0.1) {
                    for fi12 := 0; fi12 < 5; fi12++ {
                        fmt.Println(fi11)
                        var v13 float64 = 1.5
                        var v14 uint32 = 2147483647
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                }
                _ = fi11
            }
            v4 += v4
        } else {
            switch "a" {
            case "foo bar":
                if false {
                    v15 := GenS1{F0: (v10 - int(v9)), F1: "a", F2: 42}
                    v5 = append(v5, "\n\t")
                    _ = v15
                } else {
                    if (42 >= 2147483648) {
                        var v16 int64 = v4
                        v17 := GenS2{F0: (v8 % ((1) | 1)), F1: int16(v6)}
                        _ = v16
                        _ = v17
                    }
                    var v18 string = ("" + "\n\t")
                    delete(v2, 2)
                    _ = v18
                }
            case "hello":
                v19 := GenS3{F0: "foo bar", F1: 128}
                _ = v19
            default:
                if true {
                    v5 = append(v5, "foo bar")
                    v3 += (v8 << ((4294967295) & 63))
                } else {
                    for fi20 := 0; fi20 < 5; fi20++ {
                        var v21 uint16 = v7
                        if (true || (v3 != (v3 * v3))) {
                            v2[2147483647] = (fi20 / ((int(v7)) | 1))
                            v8 += -128
                            v22 := GenS4{F0: 490, F1: 18446744073709551615}
                            _ = v22
                        }
                        _ = fi20
                        _ = v21
                    }
                    v2[2147483648] = (v6 - v6)
                    if true {
                        if false {
                            v23 := GenS5{F0: (v9 & v9), F1: 1, F2: (v4 >> (((v9 / ((v9) | 1))) & 63))}
                            _ = v23
                        } else {
                            v5 = append(v5, "\n\t")
                        }
                    } else {
                        if (3.14159 == ((-2.25 * 0.1) * 0.1)) {
                            v24 := GenS6{F0: v7, F1: 1}
                            v25 := GenS2{F0: v3, F1: 253}
                            _ = v24
                            _ = v25
                        }
                        if (0.0 < 0.0) {
                            v3 += (v3 / (((-(v1))) | 1))
                            v26 := GenS3{F0: "a", F1: 0}
                            _ = v26
                        }
                    }
                }
            }
            var v27 uint64 = 4294967295
            _ = v27
        }
    default:
        if false {
            v6 += (^(v10))
            v9 += 315
        }
    }
    var v28 string = "a"
    v29 := GenS3{F0: ("foo bar" + ("hello" + "\n\t")), F1: 5}
    v28 = "foo bar"
    v7 = 2
    var v30 float64 = (3.14159 - 0.1)
    fmt.Println(v30, v9, v1)
    v10 += v10
    for fi31 := 0; fi31 < 5; fi31++ {
        v8 -= 0
        for fi32 := 0; fi32 < 15; fi32++ {
            v33 := GenS7{F0: (^(v4))}
            _ = fi32
            _ = v33
        }
        if ((v28 + v28) == v28) {
            var v34 int16 = 32767
            var v35 uint = v9
            var v36 int = 32768
            _ = v34
            _ = v35
            _ = v36
        }
        _ = fi31
    }
    for fi37 := 0; fi37 < 8; fi37++ {
        if true {
            var v38 int32 = -2147483648
            _ = v38
        } else {
            var v39 uint = 255
            var v40 uint32 = 2147483647
            switch v28 {
            case "":
                v41 := GenS8{F0: v39}
                _ = v41
            }
            _ = v39
            _ = v40
        }
        switch "a" {
        case "":
            v8 += (v3 ^ 71)
        default:
            switch v10 {
            case 127:
                if false {
                    v28 = (v28 + (v28 + (v28 + v28)))
                    v42 := GenS9{F0: 43, F1: 42, F2: (-(v10))}
                    if false {
                        switch 32767 {
                        case 32767:
                            var v43 float64 = float64((float32((v30 * v30)) * ((3.14159 - (1e10 + ((1e10 - 1e10) + 0.1))) - 0.0)))
                            _ = v43
                        case 231:
                            v3 += v3
                        default:
                            var v44 int64 = v4
                            _ = v44
                        }
                        var v45 uint = v9
                        var v46 int16 = 1
                        _ = v45
                        _ = v46
                    }
                    _ = v42
                } else {
                    if false {
                        v47 := GenS10{F0: v4, F1: true, F2: "hello"}
                        _ = v47
                    }
                }
            case -1:
                for fi48 := 0; fi48 < 12; fi48++ {
                    var v49 uint8 = 127
                    _ = fi48
                    _ = v49
                }
            default:
                v50 := GenS9{F0: -117, F1: 2, F2: fi37}
                _ = v50
            }
        }
        _ = fi37
    }
    v28 = ""
    switch v28 {
    case "a":
        var v51 int16 = 0
        _ = v51
    default:
        fmt.Println(v7, v10)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
