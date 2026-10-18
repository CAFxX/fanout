package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint64 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 uint = 255
    var v2 int = 128
    var v3 uint8 = 42
    var v4 float64 = 3.14159
    var v5 int32 = 1
    var v6 string = ""
    var v7 int16 = 1
    v1 = (v1 << ((65535) & 63))
    if (!false) {
        fmt.Println(v7)
    }
    v7 += (v7 % (((v7 - 1)) | 1))
    v8 := GenS1{F0: ((v6 + v6) + v6), F1: 2, F2: 18446744073709551615}
    v9 := GenS2{F0: 42}
    v10 := GenS2{F0: 4294967295}
    v7 += 32767
    switch v2 {
    case -2147483648:
        v11 := []int{v2, 32768, v2}
        _ = v11
    case 1:
        v12 := []int{849, (v2 * -150)}
        _ = v12
    case -1:
        fmt.Println(v5, v7)
    default:
        fmt.Println(v1, v2, v1)
    }
    var v13 uint8 = 42
    v6 = (("foo bar" + v6) + ((v6 + "a") + ((v6 + "foo bar") + v6)))
    v14 := GenS2{F0: 456}
    if true {
        var v15 int32 = v5
        _ = v15
    }
    fmt.Println(v2, v2, v6)
    switch "" {
    case "\n\t":
        if false {
            v16 := GenS1{F0: "a", F1: 4294967295, F2: 4294967295}
            for fi17 := 0; fi17 < 12; fi17++ {
                var v18 int32 = (v5 & 42)
                if (v6 <= "hello") {
                    var v19 int8 = -67
                    v2 += fi17
                    _ = v19
                } else {
                    var v20 uint8 = v3
                    _ = v20
                }
                _ = fi17
                _ = v18
            }
            v1 -= (v1 / ((v1) | 1))
            _ = v16
        } else {
            var v21 int64 = -967
            switch "hello" {
            case "":
                var v22 uint8 = (v13 * 128)
                _ = v22
            }
            if true {
                v3 = (^(v3))
            } else {
                var v23 int32 = -847
                _ = v23
            }
            _ = v21
        }
    case "a":
        var v24 uint32 = 65535
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
