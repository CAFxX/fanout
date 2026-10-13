package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 uint64; F1 string; F2 string }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 int16; F1 uint16 }
type GenS5 struct { F0 int8 }

func main() {
    var v1 float64 = (0.0 / (1e10 - -2.25))
    var v2 uint32 = 573
    v3 := []string{"foo bar", ("a" + "hello"), "hello"}
    var v4 int8 = 1
    var v5 float32 = 0.0
    var v6 float32 = (v5 + ((v5 / 1.0) + v5))
    v7 := GenS1{F0: 1, F1: 2147483647}
    v8 := GenS2{F0: 2, F1: ("a" + ("\n\t" + "hello")), F2: ""}
    var v9 uint8 = 215
    if (true && false) {
        for fi10 := 0; fi10 < 13; fi10++ {
            v9 -= (v9 - (^(v9)))
            _ = fi10
        }
    } else {
        v9 += (v9 ^ (v9 % ((236) | 1)))
        v2 += (^(v2))
    }
    if false {
        switch "hello" {
        case "foo bar":
            fmt.Println(v9, v9, v4)
        default:
            if len(v3) > 0 {
                v3[0] = ("\n\t" + ("\n\t" + "foo bar"))
            }
        }
        v4 = v4
    } else {
        v11 := GenS3{F0: "a", F1: 218}
        v12 := GenS1{F0: 0, F1: 2147483647}
        _ = v11
        _ = v12
    }
    v13 := GenS1{F0: 0, F1: -1}
    v3 = append(v3, ("a" + "foo bar"))
    var v14 int = 128
    if true {
        v15 := GenS4{F0: -32768, F1: 2}
        for fi16 := 0; fi16 < 1; fi16++ {
            v14 += (v14 - v14)
            _ = fi16
        }
        v9 += (v9 << ((18446744073709551615) & 63))
        _ = v15
    }
    var v17 uint64 = 296
    var v18 int = (v14 & (v14 >> ((1) & 63)))
    fmt.Println(v18, v6, v1)
    var v19 uint64 = (v17 - uint64(v9))
    v17 = (^(v19))
    switch "a" {
    case "hello":
        if ((!true) || (!((v6 * (1.5 + (v5 - ((v5 + v6) + ((v6 / (v5 - -2.25)) / v5))))) < v5))) {
            v17 = v17
            v17 = 2
        } else {
            v20 := GenS5{F0: 0}
            _ = v20
        }
    case "\n\t":
        var v21 int8 = (v4 >> ((42) & 63))
        _ = v21
    case "a":
        fmt.Println(v1, v19)
    default:
        switch "a" {
        case "foo bar":
            switch "hello" {
            case "a":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case "hello":
                fmt.Println(v6, v17, v5)
            }
        }
    }
    switch "a" {
    case "foo bar":
        if ("hello" > "\n\t") {
            var v22 uint = 104
            if (true || (int32(v2) < -2147483648)) {
                v1 -= v1
                switch (v14 << (((v22 - (^(v22)))) & 63)) {
                case 127:
                    var v23 int16 = -989
                    _ = v23
                case 2147483647:
                    for fi24 := 0; fi24 < 18; fi24++ {
                        var v25 uint = (v22 >> ((720) & 63))
                        fmt.Println(fi24, v9)
                        _ = fi24
                        _ = v25
                    }
                default:
                    fmt.Println(v22, v14)
                }
                switch "a" {
                case "\n\t":
                    if (!false) {
                        var v26 uint8 = (v9 - v9)
                        v3 = append(v3, ((("\n\t" + "foo bar") + "hello") + "hello"))
                        _ = v26
                    }
                case "":
                    v14 = (v14 - v14)
                case "foo bar":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
            }
            _ = v22
        } else {
            v27 := GenS5{F0: (v4 + (v4 % (((v4 - v4)) | 1)))}
            _ = v27
        }
    case "hello":
        if true {
            fmt.Println(v6)
        }
    case "a":
        fmt.Println(v9)
    default:
        for fi28 := 0; fi28 < 5; fi28++ {
            _ = len(v3)
            v6 = v6
            _ = fi28
        }
    }
    for fi29 := 0; fi29 < 2; fi29++ {
        v30 := GenS1{F0: (v17 & 670), F1: -2147483648}
        v18 = 127
        var v31 float64 = (v1 - v1)
        _ = fi29
        _ = v30
        _ = v31
    }
    v4 = (v4 % ((100) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
