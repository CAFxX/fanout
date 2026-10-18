package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 int16 }

func main() {
    v1 := []int{4294967295}
    v2 := []uint32{42, 65535}
    var v3 int32 = -1
    var v4 uint32 = 406
    v1 = append(v1, 57)
    var v5 bool = (!false)
    var v6 uint16 = 65535
    if true {
        var v7 int64 = 0
        _ = v7
    } else {
        switch "a" {
        case "":
            switch "foo bar" {
            case "a":
                if v5 {
                    v3 = (v3 >> ((18446744073709551615) & 63))
                    switch "hello" {
                    case "\n\t":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                } else {
                    var v8 uint16 = (^(v6))
                    var v9 float32 = 3.14159
                    v10 := GenS1{F0: "hello"}
                    _ = v10
                    _ = v8
                    _ = v9
                }
            }
        }
    }
    var v11 int32 = (v3 + v3)
    for fi12 := 0; fi12 < 18; fi12++ {
        for fi13 := 0; fi13 < 11; fi13++ {
            fmt.Println(fi13, fi12)
            fmt.Println(v4, fi13, v4)
            _ = fi13
        }
        v14 := GenS1{F0: ""}
        _ = fi12
        _ = v14
    }
    v4 = v4
    var v15 float64 = ((0.0 + -2.25) - ((((0.0 * (1.5 + 1.5)) * 1e10) - 1e10) + (float64(3.14159) + 0.0)))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if v5 {
        v15 = 3.14159
        v11 = v11
    } else {
        v16 := GenS2{F0: (v11 - v3), F1: (v4 ^ (^(v4)))}
        v5 = (!(("a" + "hello") != ""))
        switch "hello" {
        case "hello":
            switch 42 {
            case 32768:
                v17 := GenS3{F0: 64}
                _ = v17
            case 938:
                for fi18 := 0; fi18 < 15; fi18++ {
                    v3 += (v11 - v3)
                    var v19 uint8 = 147
                    fmt.Println(v6, v6)
                    _ = fi18
                    _ = v19
                }
            case -2147483648:
                if (-1 != int16(v4)) {
                    var v20 uint = 1
                    _ = v20
                }
            }
        }
        _ = v16
    }
    v3 += v11
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
