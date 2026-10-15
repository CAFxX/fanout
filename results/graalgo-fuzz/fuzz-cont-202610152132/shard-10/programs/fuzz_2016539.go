package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string; F1 int; F2 int64 }

func main() {
    v1 := []int{}
    v2 := []uint32{146, 396}
    var v3 string = "\n\t"
    var v4 int8 = -1
    _ = len(v2)
    var v5 int8 = v4
    v4 -= (v5 / (((^(v4))) | 1))
    v6 := GenS1{F0: 195, F1: (("" + "hello") + v3), F2: 2147483647}
    switch (v3 + ((v3 + ("a" + ("hello" + v3))) + v3)) {
    case "a":
        var v7 int16 = -1
        _ = v7
    case "foo bar":
        if (213 <= 42) {
            _ = len(v2)
            v8 := GenS1{F0: 0, F1: v3, F2: 2147483648}
            _ = v8
        } else {
            if len(v2) > 0 {
                v2[0] = 2
            }
            var v9 uint16 = 969
            v1 = append(v1, 0)
            _ = v9
        }
    case "hello":
        var v10 int8 = (v5 << ((261) & 63))
        _ = v10
    }
    if len(v1) > 0 {
        v1[0] = 0
    }
    if true {
        fmt.Println(v4, v4)
        v5 = v4
    } else {
        if (32767 <= 2) {
            var v11 uint32 = 65535
            v3 = v3
            switch "a" {
            case "a":
                for fi12 := 0; fi12 < 12; fi12++ {
                    _ = len(v2)
                    fmt.Println(v11, v4)
                    v13 := GenS1{F0: 160, F1: "foo bar", F2: (^(v11))}
                    _ = fi12
                    _ = v13
                }
            default:
                v14 := GenS1{F0: 1, F1: "foo bar", F2: v11}
                _ = v14
            }
            _ = v11
        } else {
            var v15 int = 4294967295
            _ = v15
        }
        for fi16 := 0; fi16 < 15; fi16++ {
            v17 := GenS1{F0: 128, F1: "", F2: 0}
            _ = fi16
            _ = v17
        }
        var v18 float32 = -2.25
        _ = v18
    }
    fmt.Println(v4)
    if false {
        for fi19 := 0; fi19 < 13; fi19++ {
            for fi20 := 0; fi20 < 11; fi20++ {
                var v21 string = v3
                _ = fi20
                _ = v21
            }
            switch (fi19 % (((fi19 | (^(fi19)))) | 1)) {
            case -622:
                v5 -= (-(v4))
            case 2147483647:
                fmt.Println(v3, v5)
            case 32768:
                v4 -= v5
            }
            _ = fi19
        }
        _ = len(v2)
    } else {
        if (639 >= -848) {
            var v22 int64 = -456
            switch (v3 + (v3 + v3)) {
            case "foo bar":
                switch 42 {
                case 128:
                    for fi23 := 0; fi23 < 14; fi23++ {
                        v24 := GenS2{F0: 84}
                        for fi25 := 0; fi25 < 17; fi25++ {
                            var v26 uint64 = 4294967295
                            fmt.Println(fi23, v22, v5)
                            _ = fi25
                            _ = v26
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = fi23
                        _ = v24
                    }
                default:
                    var v27 uint = 114
                    _ = v27
                }
            case "":
                v28 := GenS2{F0: 1}
                _ = v28
            case "a":
                fmt.Println(v3, v5)
            default:
                v1 = append(v1, 4294967295)
            }
            fmt.Println(v22)
            _ = v22
        } else {
            var v29 int64 = 1
            v29 -= (v29 << ((907) & 63))
            if len(v2) > 0 {
                v2[0] = 2
            }
            _ = v29
        }
    }
    v1 = append(v1, 2147483647)
    var v30 float64 = 3.14159
    var v31 int64 = 0
    if (!(-855 < -2147483648)) {
        v32 := GenS3{F0: "\n\t", F1: -9223372036854775808, F2: (v31 | v31)}
        _ = v32
    }
    fmt.Println(v4, v5, v4)
    v31 += 42
    v33 := GenS3{F0: v3, F1: 488, F2: v31}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
