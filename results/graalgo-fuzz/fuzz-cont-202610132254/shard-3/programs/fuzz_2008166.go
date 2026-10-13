package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 int32; F1 uint8 }

func main() {
    var v1 int16 = 42
    var v2 int64 = 1
    var v3 int16 = v1
    v4 := make(map[uint32]bool)
    v5 := make(map[uint32]string)
    v3 = (-(v3))
    v3 -= int16(v2)
    var v6 uint16 = 65535
    var v7 uint8 = 210
    switch (("a" + ((("a" + "\n\t") + "") + ("a" + "\n\t"))) + "hello") {
    case "foo bar":
        if false {
            var v8 int16 = (v3 / ((-32768) | 1))
            var v9 bool = (24 == v6)
            _ = v8
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        if false {
            v11 := []int{42, (-(fi10)), (^(fi10))}
            v12 := v5[1]
            _ = v11
            _ = v12
        } else {
            var v13 int32 = -1
            var v14 bool = (0.0 <= 3.14159)
            _ = v13
            _ = v14
        }
        v1 = v1
        switch (-(fi10)) {
        case 42:
            for fi15 := 0; fi15 < 1; fi15++ {
                v16 := GenS1{F0: -1, F1: "\n\t"}
                if true {
                    if (-1 >= 636) {
                        v17 := GenS2{F0: (true && false), F1: (v2 % (((v2 / ((v2) | 1))) | 1))}
                        v18 := GenS3{F0: int32(fi15), F1: uint8(v6)}
                        var v19 float32 = 3.14159
                        _ = v17
                        _ = v18
                        _ = v19
                    } else {
                        v6 -= v6
                    }
                    v5[633] = "foo bar"
                    if true {
                        var v20 int = int(v7)
                        _ = v20
                    } else {
                        fmt.Println(v6, v1, v1)
                        delete(v5, 4294967295)
                        v5[435] = "foo bar"
                    }
                }
                v1 -= v1
                _ = fi15
                _ = v16
            }
        case 983:
            v21 := []int{0, (fi10 - (fi10 * fi10))}
            _ = v21
        case 2147483647:
            switch "\n\t" {
            case "foo bar":
                v3 = v1
            }
        default:
            v1 -= 1000
        }
        _ = fi10
    }
    switch int(v3) {
    case 32767:
        if true {
            v1 = 42
            fmt.Println(v3, v6, v6)
            if (((float32(3.14159) - 3.14159) > 0.0) && false) {
                v2 = (v2 / (((v2 + v2)) | 1))
                v1 = (v1 << ((4294967295) & 63))
            } else {
                v7 = (v7 >> ((65535) & 63))
            }
        }
    case 127:
        v22 := []int{}
        _ = v22
    default:
        v3 -= (-(v3))
    }
    var v23 bool = false
    v3 = int16(v6)
    v1 = -1
    switch "" {
    case "\n\t":
        v24 := GenS3{F0: 42, F1: (^(v7))}
        _ = v24
    case "a":
        var v25 int16 = v3
        _ = v25
    }
    for fi26 := 0; fi26 < 3; fi26++ {
        v3 -= (v3 & v3)
        _ = fi26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
