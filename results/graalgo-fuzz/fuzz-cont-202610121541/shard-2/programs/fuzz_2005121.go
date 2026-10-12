package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 int }

func main() {
    var v1 float32 = 1.5
    var v2 uint8 = 127
    var v3 uint64 = 18446744073709551615
    v4 := []uint32{0, 1, 138, 895}
    var v5 int8 = 0
    v6 := make(map[int]string)
    var v7 uint = 0
    fmt.Println(v3)
    v4 = append(v4, 42)
    switch "" {
    case "a":
        if true {
            for fi8 := 0; fi8 < 17; fi8++ {
                v2 += (^(v2))
                _ = fi8
            }
        } else {
            v9 := GenS1{F0: 42, F1: -9223372036854775808, F2: 42}
            var v10 uint64 = v3
            fmt.Println(v5)
            _ = v10
            _ = v9
        }
    case "":
        v11 := GenS1{F0: 0, F1: -1, F2: 726}
        _ = v11
    case "\n\t":
        if true {
            v12 := GenS1{F0: 966, F1: -338, F2: 331}
            _ = v12
        }
    }
    if true {
        var v13 uint32 = 0
        fmt.Println(v3)
        _ = v13
    } else {
        var v14 uint = v7
        _ = v14
    }
    fmt.Println(v5, v5)
    for fi15 := 0; fi15 < 11; fi15++ {
        for fi16 := 0; fi16 < 17; fi16++ {
            var v17 bool = false
            _ = len(v6)
            v6[(fi15 % (((fi16 % ((2147483647) | 1))) | 1))] = ("hello" + ("" + "foo bar"))
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    if true {
        var v18 int32 = 0
        if (true && (!(!true))) {
            if true {
                if true {
                    var v19 uint32 = 1
                    var v20 uint16 = 2
                    v21 := GenS1{F0: v20, F1: 0, F2: 127}
                    _ = v19
                    _ = v20
                    _ = v21
                } else {
                    var v22 int16 = 934
                    if true {
                        switch 9223372036854775807 {
                        case 1:
                            fmt.Println(v22)
                        case 110:
                            fmt.Println(v3)
                        case -311:
                            v5 = (v5 % (((v5 << ((v7) & 63))) | 1))
                        }
                        for fi23 := 0; fi23 < 2; fi23++ {
                            fmt.Println(v22, v3, v5)
                            _ = fi23
                        }
                        v2 = (^(v2))
                    }
                    _ = v22
                }
                fmt.Println(v1)
            } else {
                v24 := GenS1{F0: 65535, F1: 9223372036854775807, F2: int(v5)}
                _ = v24
            }
        }
        _ = len(v4)
        _ = v18
    } else {
        switch -426 {
        case 4294967295:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        if true {
            v1 = float32((3.14159 / (-2.25 + float64(v1))))
        } else {
            var v25 string = "\n\t"
            _ = v25
        }
    }
    var v26 uint32 = 371
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
