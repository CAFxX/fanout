package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32; F1 int16; F2 int }
func genFn0() int64 {
    if false {
        var v6 int8 = 100
        switch 323 {
        case 32767:
            switch 2147483647 {
            case 32767:
                for fi7 := 0; fi7 < 11; fi7++ {
                    for fi8 := 0; fi8 < 13; fi8++ {
                        v9 := GenS1{F0: 2, F1: 18446744073709551615}
                        fmt.Println(v6)
                        _ = fi8
                        _ = v9
                    }
                    _ = fi7
                }
            default:
                fmt.Println(v6)
            }
        }
        v10 := GenS1{F0: 0, F1: 0}
        _ = v10
        _ = v6
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        fmt.Println(fi11)
        var v12 int64 = 1
        var v13 uint16 = 269
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 int = 293
    v14 += v14
    _ = v14
    return -853
}
func genFn1() int8 {
    for fi15 := 0; fi15 < 3; fi15++ {
        for fi16 := 0; fi16 < 7; fi16++ {
            v17 := []int{}
            if (!true) {
                var v18 string = "foo bar"
                if len(v17) > 0 {
                    v17[0] = (fi15 << ((255) & 63))
                }
                _ = v18
            } else {
                if len(v17) > 0 {
                    v17[0] = (fi16 + fi15)
                }
            }
            v19 := GenS2{F0: -66}
            _ = fi16
            _ = v17
            _ = v19
        }
        _ = fi15
    }
    for fi20 := 0; fi20 < 16; fi20++ {
        for fi21 := 0; fi21 < 18; fi21++ {
            var v22 uint16 = 0
            v22 = v22
            _ = fi21
            _ = v22
        }
        _ = fi20
    }
    var v23 uint8 = 1
    _ = v23
    return -106
}

func main() {
    var v1 int32 = -542
    var v2 uint = 42
    var v3 int = 521
    var v4 uint8 = 128
    v5 := make(map[string]bool)
    v24 := v5["a"]
    v1 = 42
    switch "" {
    case "a":
        if (genFn0() < -996) {
            v4 = (v4 ^ v4)
            if v24 {
                v25 := []int{(v3 | (^(v3))), int(v2), 237, v3}
                var v26 float64 = 1e10
                if false {
                    _ = len(v25)
                    v3 += v3
                    v27 := GenS3{F0: 429, F1: 694, F2: (v3 / ((v3) | 1))}
                    _ = v27
                }
                _ = v25
                _ = v26
            }
        }
    case "hello":
        for fi28 := 0; fi28 < 9; fi28++ {
            var v29 uint = 4294967295
            _ = fi28
            _ = v29
        }
    case "":
        switch v3 {
        case 0:
            v3 -= (v3 - (v3 % (((v3 - (v3 | (^(v3))))) | 1)))
        }
    }
    v30 := []int{(v3 - v3), v3, v3, (v3 << ((v2) & 63))}
    for fi31 := 0; fi31 < 6; fi31++ {
        v32 := GenS3{F0: 0, F1: -1, F2: -9223372036854775808}
        var v33 int16 = 1000
        fmt.Println(v4)
        _ = fi31
        _ = v32
        _ = v33
    }
    _ = len(v5)
    v34 := GenS2{F0: 0}
    v35 := GenS2{F0: -128}
    for fi36 := 0; fi36 < 16; fi36++ {
        switch v3 {
        case -55:
            fmt.Println(v24, fi36)
        case -759:
            switch (fi36 << ((v2) & 63)) {
            case 165:
                var v37 string = "\n\t"
                _ = v37
            case 0:
                var v38 float32 = 0.1
                _ = v38
            case -1:
                v3 -= (v3 & fi36)
            }
        default:
            v4 = (^(v4))
        }
        var v39 int32 = (v1 & (v1 / (((v1 ^ v1)) | 1)))
        if (!((v4 - (v4 * (v4 >> ((uint(v4)) & 63)))) > v4)) {
            if (!(!((1 >= v3) && (7 <= 1)))) {
                v40 := v5[""]
                if v40 {
                    for fi41 := 0; fi41 < 18; fi41++ {
                        v42 := v5[""]
                        var v43 uint32 = 1
                        _ = fi41
                        _ = v42
                        _ = v43
                    }
                    delete(v5, "hello")
                } else {
                    var v44 int = (fi36 - v3)
                    _ = v44
                }
                for fi45 := 0; fi45 < 1; fi45++ {
                    v40 = false
                    _ = fi45
                }
                _ = v40
            }
        } else {
            if ("foo bar" < ("" + "")) {
                v46 := GenS2{F0: -128}
                v39 += (v39 | v39)
                _ = v46
            }
            v24 = (!(0.0 != 3.14159))
        }
        _ = fi36
        _ = v39
    }
    v1 -= int32(v2)
    fmt.Println(v24, v4, v1)
    fmt.Println(v24, v2, v2)
    var v47 int = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v4
    _ = v47
    _ = v5
}
