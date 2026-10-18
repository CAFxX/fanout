package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 uint32; F2 int8 }
type GenS4 struct { F0 uint16; F1 int }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint8; F1 int64 }
type GenS7 struct { F0 uint8; F1 int; F2 string }

func main() {
    v1 := []int{2147483647, 390, 1}
    var v2 int64 = 0
    var v3 int8 = 99
    var v4 int = 1
    switch v4 {
    case 32767:
        v4 -= -716
    }
    v5 := GenS1{F0: 128}
    v2 = (v2 * -1)
    if ((813 != 2) && (!(1 >= 69))) {
        var v6 float32 = ((1e10 + ((3.14159 + 1e10) + (0.0 + ((0.1 / -2.25) * (0.0 / ((((1e10 + (-2.25 + 1e10)) + 1.5) / (0.0 + 1.5)) + 0.1)))))) + 0.1)
        v4 -= v4
        v1 = append(v1, 9223372036854775807)
        _ = v6
    } else {
        var v7 uint64 = 1
        if ("\n\t" != "a") {
            switch v4 {
            case 9223372036854775807:
                var v8 uint64 = (v7 << ((1) & 63))
                _ = v8
            default:
                v7 = (v7 * (v7 >> ((18446744073709551615) & 63)))
            }
        }
        if true {
            switch "\n\t" {
            case "":
                for fi9 := 0; fi9 < 3; fi9++ {
                    v10 := GenS2{F0: 19}
                    _ = fi9
                    _ = v10
                }
            }
            v11 := GenS2{F0: 32767}
            fmt.Println(v7)
            _ = v11
        }
        _ = v7
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        if false {
            for fi13 := 0; fi13 < 2; fi13++ {
                fmt.Println(fi13, v3, v4)
                var v14 int64 = (v2 ^ v2)
                _ = fi13
                _ = v14
            }
            v15 := GenS3{F0: (v4 - 2147483647), F1: 319, F2: v3}
            if false {
                if false {
                    var v16 uint32 = 0
                    fmt.Println(v4)
                    for fi17 := 0; fi17 < 3; fi17++ {
                        switch (("" + "a") + "") {
                        case "foo bar":
                            var v18 int64 = int64(v16)
                            _ = v18
                        case "":
                            _ = len(v1)
                        default:
                            var v19 int64 = 0
                            _ = v19
                        }
                        fmt.Println(v16, fi17)
                        v3 -= (v3 | v3)
                        _ = fi17
                    }
                    _ = v16
                }
            }
            _ = v15
        } else {
            var v20 int = (-(v4))
            fmt.Println(v4, v2)
            v2 -= (v2 / (((v2 >> ((2) & 63))) | 1))
            _ = v20
        }
        v3 -= (v3 * v3)
        v3 = v3
        _ = fi12
    }
    var v21 bool = (true || (-1 == -32768))
    v22 := GenS4{F0: 42, F1: v4}
    v23 := GenS4{F0: 65535, F1: v4}
    v1 = append(v1, v4)
    var v24 float32 = float32(0.1)
    var v25 float64 = (1.5 - -2.25)
    var v26 float32 = v24
    if (42 == 255) {
        if false {
            for fi27 := 0; fi27 < 13; fi27++ {
                v2 -= -1
                for fi28 := 0; fi28 < 17; fi28++ {
                    var v29 uint = 59
                    _ = fi28
                    _ = v29
                }
                _ = fi27
            }
            if (v3 > (v3 % ((v3) | 1))) {
                v2 = (v2 << ((985) & 63))
                v26 -= (v26 - float32(v25))
                for fi30 := 0; fi30 < 8; fi30++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi30
                }
            } else {
                var v31 bool = v21
                _ = v31
            }
        } else {
            v1 = append(v1, v4)
        }
    }
    v21 = false
    v25 = v25
    if ((v2 % (((v2 << ((255) & 63))) | 1)) >= (v2 * (v2 - 42))) {
        if true {
            v32 := GenS1{F0: 32768}
            _ = v32
        } else {
            v33 := GenS5{F0: 2147483647}
            switch int(v3) {
            case 127:
                var v34 int64 = -1
                _ = v34
            case 42:
                v35 := GenS6{F0: 161, F1: v2}
                _ = v35
            case -239:
                var v36 uint8 = 9
                _ = v36
            }
            v4 -= (v4 >> ((971) & 63))
            _ = v33
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        fmt.Println(v24)
        var v37 float64 = (v25 - v25)
        _ = v37
    }
    v38 := GenS7{F0: 2, F1: (^(v4)), F2: "\n\t"}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v38
    _ = v4
    _ = v5
}
