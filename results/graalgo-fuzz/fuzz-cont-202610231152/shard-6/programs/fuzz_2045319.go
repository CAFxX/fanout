package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 uint; F2 bool }
func genFn0(p0 uint16, p1 uint8, p2 int32) bool {
    fmt.Println(p0)
    v6 := GenS1{F0: "\n\t"}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (((3.14159 - float64(3.14159)) / (-2.25 * 1.5)) >= 3.14159)
}

func main() {
    var v1 uint8 = 42
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]string)
    v4 := []int{4294967295, 347, -939}
    var v5 int16 = 215
    v1 = 98
    switch 313 {
    case 128:
        if ((375 >= 42) || true) {
            v7 := GenS2{F0: (!(788 == 42))}
            if len(v4) > 0 {
                v4[0] = 32767
            }
            _ = v7
        }
    case 0:
        if false {
            var v8 float64 = 1.5
            fmt.Println(v5)
            delete(v2, 450)
            _ = v8
        } else {
            fmt.Println(v5, v5)
            v2[65535] = 42
            var v9 int8 = -128
            _ = v9
        }
    default:
        if ("hello" < "\n\t") {
            if true {
                if ((v1 & v1) < 2) {
                    var v10 uint32 = 1
                    _ = v10
                }
                if (!(false && true)) {
                    v1 -= 255
                    if false {
                        v1 -= (v1 | 1)
                        var v11 uint64 = 604
                        var v12 float32 = 3.14159
                        _ = v11
                        _ = v12
                    }
                    _ = len(v3)
                }
            } else {
                v13 := GenS3{F0: (v1 << ((65535) & 63)), F1: 167, F2: (true || (-563 > -9223372036854775808))}
                _ = v13
            }
            var v14 float32 = -2.25
            _ = v14
        }
    }
    v1 -= v1
    for fi15 := 0; fi15 < 6; fi15++ {
        fmt.Println(v1, fi15, v5)
        fmt.Println(v1, v5)
        _ = fi15
    }
    fmt.Println(v5, v1)
    v16 := GenS1{F0: "\n\t"}
    switch 128 {
    case -1:
        if (0 == 18446744073709551615) {
            for fi17 := 0; fi17 < 8; fi17++ {
                switch ("\n\t" + "") {
                case "":
                    v1 = v1
                case "hello":
                    var v18 string = "hello"
                    _ = v18
                case "\n\t":
                    var v19 uint32 = 774
                    _ = v19
                }
                _ = fi17
            }
            for fi20 := 0; fi20 < 2; fi20++ {
                v5 = (^(v5))
                fmt.Println(v1, v1, fi20)
                _ = fi20
            }
        } else {
            delete(v2, 2147483648)
            fmt.Println(v1, v5)
        }
    case -272:
        var v21 float32 = 3.14159
        _ = v21
    case 32768:
        var v22 int32 = 0
        _ = v22
    }
    var v23 bool = true
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
