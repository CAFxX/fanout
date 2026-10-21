package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8; F1 int16; F2 bool }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int8 }

func main() {
    var v1 int32 = -2147483648
    var v2 uint = 18446744073709551615
    var v3 int32 = (^(v1))
    if false {
        v4 := GenS1{F0: 4294967295}
        _ = v4
    }
    var v5 string = "\n\t"
    switch 1 {
    case -436:
        if (551 <= -154) {
            v5 = (v5 + v5)
            fmt.Println(v2)
            if true {
                var v6 int16 = 32767
                if true {
                    v2 = 255
                }
                var v7 uint32 = 908
                _ = v6
                _ = v7
            } else {
                if ("foo bar" > (((("\n\t" + v5) + v5) + "") + v5)) {
                    for fi8 := 0; fi8 < 16; fi8++ {
                        fmt.Println(v1, v1, v2)
                        fmt.Println(v2, fi8)
                        v3 = -2147483648
                        _ = fi8
                    }
                    switch "foo bar" {
                    case "foo bar":
                        var v9 int32 = (v1 >> (((v2 | (v2 & v2))) & 63))
                        _ = v9
                    case "a":
                        fmt.Println(v2, v2)
                    }
                } else {
                    v2 -= 130
                    var v10 float64 = 3.14159
                    _ = v10
                }
                fmt.Println(v2)
            }
        } else {
            v1 = 127
        }
    }
    var v11 uint16 = 2
    if false {
        fmt.Println(v3, v1)
        var v12 uint16 = (v11 + (^(v11)))
        _ = v12
    }
    if true {
        v11 += (v11 >> ((v2) & 63))
    } else {
        var v13 int8 = -18
        _ = v13
    }
    if false {
        for fi14 := 0; fi14 < 20; fi14++ {
            for fi15 := 0; fi15 < 14; fi15++ {
                fmt.Println(v3, fi14, v2)
                v16 := GenS1{F0: 18446744073709551615}
                v17 := GenS2{F0: 42}
                _ = fi15
                _ = v16
                _ = v17
            }
            _ = fi14
        }
        if true {
            v2 = v2
            v2 -= (^(v2))
        } else {
            v11 += v11
            for fi18 := 0; fi18 < 1; fi18++ {
                switch v5 {
                case "hello":
                    fmt.Println(v1, v11, v3)
                default:
                    v19 := GenS3{F0: int8(fi18), F1: 446, F2: ((true || (true || (127 >= 0))) && true)}
                    _ = v19
                }
                v20 := GenS4{F0: 2147483647}
                if (!(3.14159 > -2.25)) {
                    var v21 int64 = -273
                    fmt.Println(v1, v2)
                    _ = v21
                }
                _ = fi18
                _ = v20
            }
            var v22 int64 = 9223372036854775807
            _ = v22
        }
        v2 = v2
    } else {
        for fi23 := 0; fi23 < 12; fi23++ {
            v5 = v5
            v24 := GenS5{F0: 42}
            v11 += v11
            _ = fi23
            _ = v24
        }
        v3 -= v1
    }
    v3 -= (-(v3))
    v11 = v11
    v1 = -573
    var v25 uint16 = v11
    if false {
        fmt.Println(v3, v11)
        fmt.Println(v25)
    } else {
        v26 := GenS2{F0: 42}
        _ = v26
    }
    if (0 != 192) {
        v27 := []int{}
        _ = v27
    } else {
        if (v25 > v25) {
            for fi28 := 0; fi28 < 8; fi28++ {
                if true {
                    var v29 int8 = -1
                    v30 := []int{(fi28 | 32768), fi28}
                    _ = len(v30)
                    _ = v29
                    _ = v30
                }
                _ = fi28
            }
        } else {
            v31 := []int{0, 9223372036854775807, 128}
            var v32 int64 = -955
            _ = v31
            _ = v32
        }
        v33 := []int{-455, -83}
        fmt.Println(v25)
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v25
    _ = v3
    _ = v5
}
