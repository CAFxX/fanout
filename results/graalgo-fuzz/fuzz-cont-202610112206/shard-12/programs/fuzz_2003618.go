package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 int8; F2 uint64 }
type GenS3 struct { F0 uint64; F1 uint32 }

func main() {
    var v1 uint16 = 1
    var v2 uint8 = 120
    var v3 uint8 = (v2 - (^(v2)))
    v4 := []int{1}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v5 int16 = -32768
    var v6 int8 = 42
    switch -315 {
    case -66:
        var v7 int64 = -1
        _ = v7
    case -1:
        var v8 float32 = 0.0
        _ = v8
    default:
        v6 += v6
    }
    if (!(((^(v5)) <= v5) || (42 == 816))) {
        switch int(v5) {
        case -735:
            for fi9 := 0; fi9 < 8; fi9++ {
                v2 = v3
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi9
            }
        case -138:
            fmt.Println(v2, v1, v1)
        case -945:
            v10 := GenS1{F0: 1}
            _ = v10
        }
    } else {
        for fi11 := 0; fi11 < 17; fi11++ {
            _ = len(v4)
            _ = fi11
        }
        v12 := GenS1{F0: -1}
        _ = v12
    }
    v13 := GenS2{F0: 4294967295, F1: -128, F2: 2}
    var v14 int8 = (v6 ^ v6)
    fmt.Println(v14, v3, v5)
    switch ("" + (("\n\t" + ("hello" + "")) + "")) {
    case "":
        if ((false && true) && true) {
            fmt.Println(v5, v3)
            switch "\n\t" {
            case "a":
                if (-9223372036854775808 == -9223372036854775808) {
                    var v15 int16 = (v5 << ((426) & 63))
                    switch "hello" {
                    case "\n\t":
                        if true {
                            fmt.Println(v1, v2, v2)
                            var v16 int8 = v6
                            _ = v16
                        }
                    case "foo bar":
                        for fi17 := 0; fi17 < 4; fi17++ {
                            var v18 uint16 = v1
                            var v19 int16 = 42
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                    default:
                        switch 1 {
                        case 127:
                            var v20 uint8 = v2
                            _ = v20
                        case 205:
                            v14 += v6
                        }
                    }
                    for fi21 := 0; fi21 < 11; fi21++ {
                        v14 -= v14
                        _ = fi21
                    }
                    _ = v15
                } else {
                    if len(v4) > 0 {
                        v4[0] = 127
                    }
                    v22 := GenS1{F0: 0}
                    _ = v22
                }
            case "":
                fmt.Println(v1, v5)
            }
            var v23 string = ""
            _ = v23
        } else {
            fmt.Println(v5, v2)
        }
    case "\n\t":
        if (true || true) {
            v4 = append(v4, int(v3))
            var v24 float64 = 0.1
            _ = v24
        } else {
            var v25 float64 = (float64(-2.25) + 3.14159)
            v6 += (v14 - v14)
            _ = v25
        }
    default:
        v26 := GenS3{F0: 18446744073709551615, F1: 42}
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
