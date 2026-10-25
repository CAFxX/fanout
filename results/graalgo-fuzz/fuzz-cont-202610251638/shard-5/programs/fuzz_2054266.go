package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 int64; F1 int; F2 uint8 }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 float32 = 3.14159
    var v2 bool = (((false || (!true)) || (false && false)) && (9223372036854775807 < 42))
    var v3 bool = v2
    v4 := []int{}
    fmt.Println(v3, v1, v3)
    fmt.Println(v3, v1)
    v5 := GenS1{F0: 2, F1: (("a" + ((("\n\t" + "a") + ("hello" + "foo bar")) + ("\n\t" + "a"))) + "\n\t")}
    v2 = ((((0.1 - float32(0.0)) / 0.1) - (v1 * v1)) == ((v1 / v1) * float32(0.1)))
    var v6 uint8 = 128
    if ((778 != 867) && (933 <= 0)) {
        fmt.Println(v1, v6)
        switch "foo bar" {
        case "foo bar":
            if ((349 <= 42) && (!(int64(v6) < -1))) {
                v7 := GenS1{F0: 666, F1: ("foo bar" + "foo bar")}
                _ = v7
            }
        case "\n\t":
            fmt.Println(v6, v3, v2)
        }
        var v8 float64 = 1.5
        _ = v8
    } else {
        switch -2147483648 {
        case 32768:
            v3 = false
        case -250:
            v2 = v3
        }
        var v9 float32 = v1
        v3 = (!v2)
        _ = v9
    }
    v4 = append(v4, 9223372036854775807)
    var v10 uint16 = 232
    v1 = v1
    if v2 {
        v1 = v1
        _ = len(v4)
    } else {
        v11 := GenS1{F0: (v10 | 65535), F1: ("a" + "\n\t")}
        v3 = (1e10 <= v1)
        _ = v11
    }
    v3 = (32767 <= 0)
    fmt.Println(v10, v6, v10)
    if false {
        switch "\n\t" {
        case "\n\t":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        default:
            _ = len(v4)
        }
        v6 = (v6 | 255)
        switch 9223372036854775807 {
        case -1:
            v6 -= v6
        case 127:
            v12 := GenS2{F0: 537, F1: 128, F2: v6}
            _ = v12
        }
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v13 := GenS3{F0: 2}
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
