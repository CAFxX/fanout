package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 uint16; F2 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 string; F1 int64 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint; F1 int8; F2 int16 }

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 int16 = 42
    v3 := []int{976, 127, 32767, 127}
    var v4 uint16 = 2
    v4 = (v4 + (v4 / (((v4 % (((v4 / ((v4) | 1))) | 1))) | 1)))
    v5 := GenS1{F0: 904}
    switch ("hello" + "") {
    case "hello":
        var v6 int16 = v2
        _ = v6
    case "\n\t":
        var v7 int = 2147483647
        _ = v7
    case "a":
        for fi8 := 0; fi8 < 14; fi8++ {
            var v9 int32 = 2147483647
            if (!((!false) && true)) {
                if false {
                    if true {
                        fmt.Println(fi8, v2)
                    } else {
                        if (!(!((0.1 - 1.5) <= 1.5))) {
                            _ = len(v3)
                            v10 := GenS2{F0: (v4 >> ((2) & 63)), F1: (v4 ^ 2), F2: 100}
                            var v11 bool = true
                            _ = v10
                            _ = v11
                        }
                    }
                    switch (fi8 + int(v1)) {
                    case 9223372036854775807:
                        var v12 uint32 = 4294967295
                        _ = v12
                    case 676:
                        v4 = (v4 & v4)
                    case 0:
                        var v13 uint32 = 214
                        _ = v13
                    }
                } else {
                    if false {
                        v14 := GenS1{F0: 65535}
                        v4 -= v4
                        _ = v14
                    } else {
                        if true {
                            v15 := GenS1{F0: 1}
                            for fi16 := 0; fi16 < 20; fi16++ {
                                v9 = (v9 + int32(v4))
                                v17 := GenS3{F0: v4}
                                v18 := GenS3{F0: 1}
                                _ = fi16
                                _ = v17
                                _ = v18
                            }
                            fmt.Println(v1, v9, v2)
                            _ = v15
                        } else {
                            if len(v3) > 0 {
                                v3[0] = (fi8 << ((42) & 63))
                            }
                            if (0.0 <= (0.1 - 0.1)) {
                                v19 := GenS4{F0: true}
                                _ = v19
                            }
                        }
                    }
                }
                for fi20 := 0; fi20 < 11; fi20++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi20
                }
            }
            _ = fi8
            _ = v9
        }
    default:
        v21 := GenS5{F0: "foo bar", F1: 891}
        _ = v21
    }
    v3 = append(v3, 9223372036854775807)
    fmt.Println(v4)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v22 := GenS6{F0: true}
    v4 = v4
    for fi23 := 0; fi23 < 5; fi23++ {
        switch "a" {
        case "\n\t":
            v24 := GenS7{F0: 325, F1: 1, F2: (v2 & 0)}
            _ = v24
        default:
            var v25 int32 = 2147483647
            _ = v25
        }
        _ = fi23
    }
    v26 := GenS1{F0: 65535}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
