package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 bool }
type GenS2 struct { F0 int; F1 uint16; F2 uint16 }
type GenS3 struct { F0 uint64; F1 bool; F2 uint }
type GenS4 struct { F0 int8; F1 uint }
type GenS5 struct { F0 string; F1 int8; F2 uint8 }

func main() {
    v1 := []uint32{322, 277}
    var v2 uint = 397
    var v3 uint64 = 42
    var v4 int = 4294967295
    v4 += 9223372036854775807
    var v5 int16 = -235
    v4 = v4
    v3 -= (v3 * 1)
    for fi6 := 0; fi6 < 14; fi6++ {
        var v7 string = "foo bar"
        _ = fi6
        _ = v7
    }
    if len(v1) > 0 {
        v1[0] = 981
    }
    v8 := GenS1{F0: (v2 >> (((v2 << (((v2 - 0)) & 63))) & 63)), F1: (v2 >> ((227) & 63)), F2: false}
    v3 -= v3
    var v9 uint8 = 127
    if (!false) {
        for fi10 := 0; fi10 < 20; fi10++ {
            if (2147483647 <= fi10) {
                var v11 int16 = int16(v4)
                _ = v11
            } else {
                if (!(false || true)) {
                    var v12 float64 = 3.14159
                    _ = v12
                }
                fmt.Println(v9)
                switch v4 {
                case -288:
                    fmt.Println(v2, v4)
                case 9223372036854775807:
                    v9 = v9
                case 441:
                    v13 := GenS1{F0: (v2 >> ((1) & 63)), F1: (v2 * 399), F2: ((0.1 < -2.25) || (false || false))}
                    _ = v13
                }
            }
            if (((128 < (v4 + (fi10 + 4294967295))) || ((0 >= (v3 | v3)) && true)) && (!(!((1 > 2) && (fi10 > (v4 + int(v2))))))) {
                v1 = append(v1, 0)
            }
            switch v4 {
            case 32767:
                var v14 uint8 = v9
                _ = v14
            case 32768:
                if (!(!true)) {
                    _ = len(v1)
                } else {
                    v15 := GenS2{F0: v4, F1: uint16(v2), F2: 40}
                    if len(v1) > 0 {
                        v1[0] = 42
                    }
                    for fi16 := 0; fi16 < 8; fi16++ {
                        v17 := GenS2{F0: (v4 / (((fi10 + 1)) | 1)), F1: 42, F2: 32767}
                        var v18 float64 = (float64(-2.25) / 1.5)
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                    _ = v15
                }
            case 110:
                for fi19 := 0; fi19 < 20; fi19++ {
                    fmt.Println(v5, fi19)
                    _ = fi19
                }
            }
            _ = fi10
        }
    }
    if true {
        var v20 float32 = -2.25
        if (1e10 <= v20) {
            fmt.Println(v5, v5, v4)
            var v21 uint8 = v9
            _ = v21
        }
        _ = v20
    }
    if ((86 < 42) && true) {
        _ = len(v1)
        if false {
            v1 = append(v1, 718)
            v9 += v9
            for fi22 := 0; fi22 < 18; fi22++ {
                v23 := GenS1{F0: (v2 - 0), F1: v2, F2: false}
                v1 = append(v1, 2147483647)
                _ = fi22
                _ = v23
            }
        } else {
            if (145 >= -214) {
                for fi24 := 0; fi24 < 9; fi24++ {
                    var v25 int16 = (v5 ^ v5)
                    _ = fi24
                    _ = v25
                }
                v9 -= v9
            } else {
                for fi26 := 0; fi26 < 15; fi26++ {
                    var v27 int32 = 407
                    switch "\n\t" {
                    case "foo bar":
                        v2 += v2
                    case "hello":
                        for fi28 := 0; fi28 < 7; fi28++ {
                            v5 += (v5 >> ((v2) & 63))
                            _ = fi28
                        }
                    case "":
                        var v29 uint = (^(v2))
                        _ = v29
                    }
                    switch "a" {
                    case "\n\t":
                        v3 = 0
                    default:
                        v9 += (^(v9))
                    }
                    _ = fi26
                    _ = v27
                }
                if (847 <= -894) {
                    _ = len(v1)
                } else {
                    v30 := GenS1{F0: (v2 / ((742) | 1)), F1: (v2 % (((v2 - (^(v2)))) | 1)), F2: ((v2 / (((v2 << ((v2) & 63))) | 1)) > (v2 % ((uint(v3)) | 1)))}
                    switch "a" {
                    case "a":
                        v2 += v2
                    case "foo bar":
                        for fi31 := 0; fi31 < 9; fi31++ {
                            for fi32 := 0; fi32 < 18; fi32++ {
                                v33 := GenS3{F0: (^(v3)), F1: true, F2: 18446744073709551615}
                                var v34 int = 1
                                var v35 int16 = int16(v2)
                                var v36 bool = ("hello" > "a")
                                _ = fi32
                                _ = v33
                                _ = v34
                                _ = v35
                                _ = v36
                            }
                            _ = fi31
                        }
                    case "":
                        var v37 string = ((("hello" + "a") + "hello") + "\n\t")
                        _ = v37
                    }
                    _ = v30
                }
            }
        }
    }
    var v38 uint8 = 202
    fmt.Println(v3)
    v39 := GenS4{F0: -57, F1: (v2 >> ((4294967295) & 63))}
    for fi40 := 0; fi40 < 8; fi40++ {
        if len(v1) > 0 {
            v1[0] = 42
        }
        _ = fi40
    }
    if (("\n\t" + "") > "foo bar") {
        var v41 bool = false
        switch (v4 >> (((v2 - (v2 & v2))) & 63)) {
        case -2147483648:
            fmt.Println(v4)
        }
        if len(v1) > 0 {
            v1[0] = 65535
        }
        _ = v41
    } else {
        v42 := GenS5{F0: (("" + "") + "\n\t"), F1: -1, F2: v9}
        _ = v42
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
