package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 uint8; F1 int64; F2 int }
type GenS4 struct { F0 uint64; F1 int }

func main() {
    var v1 float64 = -2.25
    v2 := []uint32{2147483648, 42}
    var v3 uint64 = 42
    v1 = v1
    var v4 uint16 = 147
    if len(v2) > 0 {
        v2[0] = 0
    }
    v2 = append(v2, 2147483647)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    _ = len(v2)
    switch 332 {
    case 42:
        var v5 int8 = -128
        _ = v5
    case 32768:
        if false {
            if true {
                v1 = 1e10
                for fi6 := 0; fi6 < 7; fi6++ {
                    var v7 bool = false
                    var v8 int16 = 204
                    v8 -= v8
                    _ = fi6
                    _ = v7
                    _ = v8
                }
                v2 = append(v2, 4294967295)
            }
        }
    default:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    if (403 < -1) {
        var v9 uint64 = v3
        v10 := GenS1{F0: 42, F1: true}
        v3 -= (v3 ^ (v9 << ((18446744073709551615) & 63)))
        _ = v10
        _ = v9
    } else {
        if (111 > 72) {
            var v11 int16 = 723
            switch "hello" {
            case "foo bar":
                v12 := GenS1{F0: -28, F1: (2 == v3)}
                _ = v12
            case "a":
                fmt.Println(v11, v1)
            default:
                if false {
                    for fi13 := 0; fi13 < 14; fi13++ {
                        if true {
                            v14 := GenS2{F0: "foo bar", F1: v11}
                            switch "\n\t" {
                            case "":
                                v15 := GenS1{F0: 42, F1: (!false)}
                                _ = v15
                            case "hello":
                                fmt.Println(v1, fi13)
                            }
                            var v16 int16 = v11
                            _ = v14
                            _ = v16
                        } else {
                            v3 = v3
                            v17 := GenS3{F0: 1, F1: -52, F2: fi13}
                            var v18 int = (fi13 / ((fi13) | 1))
                            _ = v17
                            _ = v18
                        }
                        _ = fi13
                    }
                }
            }
            if len(v2) > 0 {
                v2[0] = 65535
            }
            _ = v11
        } else {
            v3 = 4294967295
            fmt.Println(v4)
            for fi19 := 0; fi19 < 13; fi19++ {
                v3 -= v3
                v20 := GenS4{F0: v3, F1: fi19}
                _ = fi19
                _ = v20
            }
        }
        _ = len(v2)
        if (float32((v1 / v1)) >= 1e10) {
            var v21 int8 = 127
            _ = v21
        }
    }
    fmt.Println(v4)
    var v22 uint16 = (v4 << ((42) & 63))
    v23 := GenS2{F0: ("hello" + ("hello" + "\n\t")), F1: 949}
    for fi24 := 0; fi24 < 19; fi24++ {
        if true {
            var v25 int16 = 32767
            switch (fi24 - (fi24 * fi24)) {
            case -9223372036854775808:
                if false {
                    var v26 int = fi24
                    v27 := GenS3{F0: 255, F1: -1, F2: (^(v26))}
                    var v28 float32 = 1e10
                    _ = v26
                    _ = v27
                    _ = v28
                }
            case 127:
                v25 = v25
            case 32768:
                v29 := GenS2{F0: "", F1: (v25 << ((4294967295) & 63))}
                _ = v29
            }
            var v30 float64 = (v1 * v1)
            _ = v25
            _ = v30
        } else {
            for fi31 := 0; fi31 < 15; fi31++ {
                var v32 uint8 = 0
                var v33 bool = ((v32 % (((v32 - (v32 * 192))) | 1)) != (^(v32)))
                if (int16(v3) != 1) {
                    fmt.Println(fi31, v4, fi24)
                    var v34 int = fi24
                    _ = v34
                }
                _ = fi31
                _ = v32
                _ = v33
            }
        }
        v35 := GenS4{F0: (v3 << ((42) & 63)), F1: 32768}
        var v36 int = fi24
        _ = fi24
        _ = v35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
