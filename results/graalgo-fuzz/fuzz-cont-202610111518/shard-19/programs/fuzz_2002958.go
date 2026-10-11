package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 uint8; F2 uint32 }
func genFn0() int {
    var v7 float32 = ((0.1 + 1.5) + 0.0)
    switch "a" {
    case "":
        var v8 string = "a"
        _ = v8
    case "hello":
        if (42 > 2) {
            var v9 uint16 = 0
            if (false && (true || true)) {
                if (704 <= -9223372036854775808) {
                    var v10 int64 = 306
                    if (v9 > v9) {
                        v11 := GenS1{F0: 9223372036854775807}
                        if false {
                            v12 := []int{32768, -2147483648, -651}
                            v13 := GenS1{F0: (v10 >> ((2) & 63))}
                            _ = v12
                            _ = v13
                        }
                        _ = v11
                    } else {
                        v14 := []int{-125, 900, 763, -1}
                        _ = len(v14)
                        _ = v14
                    }
                    _ = v10
                }
            } else {
                var v15 string = "foo bar"
                _ = v15
            }
            v16 := []int{2147483647}
            _ = v16
            _ = v9
        }
    case "\n\t":
        var v17 string = "foo bar"
        _ = v17
    default:
        v18 := []int{-474}
        _ = v18
    }
    _ = v7
    return 960
}

func main() {
    v1 := []uint32{971, 42, 2147483648, 945}
    var v2 uint64 = 1
    var v3 uint8 = 127
    var v4 int64 = 0
    var v5 float32 = 0.0
    var v6 bool = (!false)
    v19 := GenS1{F0: (v4 | (v4 ^ 9223372036854775807))}
    fmt.Println(v5, v2)
    v2 += (v2 ^ 42)
    if (!false) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        fmt.Println(v5, v5)
        if len(v1) > 0 {
            v1[0] = 2147483648
        }
    }
    var v20 int16 = -32768
    var v21 string = ""
    v22 := GenS1{F0: v4}
    for fi23 := 0; fi23 < 1; fi23++ {
        fmt.Println(v5)
        _ = fi23
    }
    var v24 string = v21
    var v25 string = v21
    v20 -= v20
    if v6 {
        v25 = (v25 + v24)
        var v26 float32 = v5
        _ = v26
    }
    for fi27 := 0; fi27 < 14; fi27++ {
        if v6 {
            var v28 bool = (v6 && false)
            if len(v1) > 0 {
                v1[0] = 2147483647
            }
            var v29 bool = v6
            _ = v28
            _ = v29
        } else {
            var v30 int64 = -647
            fmt.Println(v20, v2)
            _ = v30
        }
        var v31 int8 = 9
        if (!(((v2 * (v2 / ((v2) | 1))) > v2) || (((v25 + ((v24 + v21) + ((v24 + v25) + ("\n\t" + (v24 + v21))))) < "") && true))) {
            var v32 int16 = v20
            _ = v32
        }
        _ = fi27
        _ = v31
    }
    switch -655 {
    case 127:
        v5 += 1e10
    case 128:
        v33 := GenS2{F0: (v2 << ((65535) & 63))}
        _ = v33
    case 2147483647:
        if len(v1) > 0 {
            v1[0] = 2147483648
        }
    default:
        if ((^(v4)) < (v4 | -661)) {
            if len(v1) > 0 {
                v1[0] = 486
            }
            switch 921 {
            case 32768:
                v6 = v6
            case 127:
                v34 := GenS3{F0: 42, F1: (v3 | (v3 | (v3 >> ((2) & 63)))), F2: 65535}
                _ = v34
            case -2147483648:
                switch "a" {
                case "\n\t":
                    var v35 string = v25
                    _ = v35
                }
            }
            fmt.Println(v2, v20, v3)
        } else {
            fmt.Println(v20)
            for fi36 := 0; fi36 < 7; fi36++ {
                v37 := GenS2{F0: 0}
                v5 += float32(1.5)
                v38 := GenS2{F0: uint64(v3)}
                _ = fi36
                _ = v37
                _ = v38
            }
        }
    }
    v24 = "hello"
    var v39 uint8 = 255
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v25
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
