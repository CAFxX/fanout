package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint16 }
type GenS3 struct { F0 int32; F1 uint8; F2 int8 }
type GenS4 struct { F0 uint64; F1 uint64; F2 int8 }
type GenS5 struct { F0 uint32; F1 uint32; F2 string }
type GenS6 struct { F0 bool; F1 int }

func main() {
    var v1 int16 = 42
    var v2 float64 = (((0.0 + (1e10 / float64(1.5))) + 0.0) - float64(((float32(0.1) + (1.5 + 1.5)) - 0.0)))
    var v3 uint64 = 4294967295
    var v4 uint64 = v3
    v5 := []string{(("\n\t" + "a") + "\n\t"), ""}
    var v6 int32 = -1
    var v7 float64 = 1e10
    var v8 int32 = v6
    var v9 string = (("" + "foo bar") + "foo bar")
    switch ("\n\t" + (("foo bar" + (("a" + v9) + (v9 + ("\n\t" + ("a" + ("" + (((v9 + v9) + "\n\t") + "\n\t"))))))) + "a")) {
    case "hello":
        if len(v5) > 0 {
            v5[0] = (v9 + v9)
        }
    case "foo bar":
        v10 := GenS1{F0: 18}
        _ = v10
    }
    fmt.Println(v2, v1, v9)
    if (713 == v8) {
        var v11 uint64 = 18446744073709551615
        _ = v11
    }
    fmt.Println(v3, v4, v7)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    switch 2147483647 {
    case -2147483648:
        for fi12 := 0; fi12 < 19; fi12++ {
            if false {
                v4 = v4
                if true {
                    v2 = v7
                    v8 += (v6 ^ v6)
                    switch (v9 + v9) {
                    case "a":
                        v8 -= (v8 % (((^(v8))) | 1))
                    case "foo bar":
                        var v13 uint8 = 42
                        _ = v13
                    case "hello":
                        var v14 int = (fi12 & (fi12 << ((1) & 63)))
                        _ = v14
                    }
                }
            } else {
                for fi15 := 0; fi15 < 17; fi15++ {
                    if (v8 <= int32(v1)) {
                        v3 = 42
                    }
                    if ((false && true) && true) {
                        v7 = 1e10
                        _ = len(v5)
                        if (42 < 4294967295) {
                            var v16 int64 = 0
                            var v17 float32 = 3.14159
                            _ = v16
                            _ = v17
                        } else {
                            if len(v5) > 0 {
                                v5[0] = (v9 + "\n\t")
                            }
                            v8 = 2147483647
                        }
                        fmt.Println(v3, v7)
                        v4 += v3
                    } else {
                        var v18 int16 = (v1 - (v1 & 42))
                        _ = v18
                    }
                    _ = fi15
                }
                if false {
                    var v19 bool = false
                    v20 := GenS2{F0: uint32(v4), F1: 793, F2: 868}
                    _ = v19
                    _ = v20
                } else {
                    fmt.Println(v8, fi12)
                    v5 = append(v5, "foo bar")
                    v8 = v6
                }
            }
            for fi21 := 0; fi21 < 15; fi21++ {
                v6 = v6
                switch (v9 + v9) {
                case "hello":
                    if false {
                        var v22 float64 = (v2 / v2)
                        _ = v22
                    }
                case "a":
                    v23 := GenS3{F0: v8, F1: 1, F2: 70}
                    var v24 uint8 = 128
                    v2 = 0.0
                    _ = v23
                    _ = v24
                }
                _ = fi21
            }
            var v25 uint64 = v4
            _ = fi12
            _ = v25
        }
    case 127:
        fmt.Println(v4, v3)
    default:
        switch 233 {
        case 4294967295:
            v26 := GenS3{F0: v8, F1: 1, F2: 100}
            _ = v26
        case -488:
            for fi27 := 0; fi27 < 16; fi27++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                for fi28 := 0; fi28 < 9; fi28++ {
                    v5 = append(v5, v9)
                    fmt.Println(fi27, v3)
                    v29 := GenS1{F0: (v3 | v4)}
                    _ = fi28
                    _ = v29
                }
                if true {
                    v8 = (^(v8))
                }
                _ = fi27
            }
        case -1:
            v8 = v8
        default:
            v30 := GenS4{F0: v3, F1: (v4 - v3), F2: 1}
            _ = v30
        }
    }
    for fi31 := 0; fi31 < 5; fi31++ {
        v2 += 0.1
        var v32 string = v9
        for fi33 := 0; fi33 < 12; fi33++ {
            switch fi31 {
            case 290:
                v6 -= v6
            case 9223372036854775807:
                fmt.Println(v2, v7)
            default:
                v5 = append(v5, v9)
                var v34 bool = (v6 == 426)
                _ = v34
            }
            v1 -= v1
            _ = fi33
        }
        _ = fi31
        _ = v32
    }
    _ = len(v5)
    v9 = v9
    fmt.Println(v8)
    fmt.Println(v1, v6, v8)
    var v35 uint = 0
    v35 = v35
    if (255 <= 213) {
        if true {
            for fi36 := 0; fi36 < 10; fi36++ {
                v5 = append(v5, (v9 + ""))
                for fi37 := 0; fi37 < 13; fi37++ {
                    v38 := GenS5{F0: 0, F1: uint32(v3), F2: ("hello" + "hello")}
                    _ = fi37
                    _ = v38
                }
                if (!(-2.25 < v7)) {
                    v35 = 42
                    v39 := GenS6{F0: false, F1: fi36}
                    var v40 uint = (v35 % ((v35) | 1))
                    _ = v39
                    _ = v40
                } else {
                    var v41 float32 = (1.5 / (1.5 - 1e10))
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = v41
                }
                _ = fi36
            }
        }
        var v42 uint8 = 0
        v4 += v3
        _ = v42
    }
    v4 = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
