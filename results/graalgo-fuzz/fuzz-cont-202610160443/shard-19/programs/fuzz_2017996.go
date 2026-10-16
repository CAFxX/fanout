package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 uint8; F2 int16 }
func genFn0(p0 bool) uint64 {
    var v5 int16 = -32768
    switch ((((("foo bar" + ("" + ("\n\t" + "\n\t"))) + "a") + ("\n\t" + "a")) + "") + "foo bar") {
    case "a":
        v6 := []int{102, 40, -9223372036854775808, 32767}
        _ = v6
    case "":
        v7 := []int{9223372036854775807}
        _ = v7
    case "\n\t":
        v8 := []int{399, 1, -1}
        _ = v8
    default:
        v9 := GenS1{F0: 255}
        _ = v9
    }
    _ = p0
    _ = v5
    return 0
}

func main() {
    var v1 bool = (2 != 774)
    var v2 uint64 = 4294967295
    var v3 uint16 = 65535
    var v4 bool = (!(0 < 2147483647))
    if (((v3 < 32767) && v4) && (3.14159 >= -2.25)) {
        v2 = (v2 % ((18446744073709551615) | 1))
    }
    v10 := []int{-1, 9223372036854775807, 426, 128}
    v10 = append(v10, 953)
    if v4 {
        if (!true) {
            switch "" {
            case "":
                var v11 int32 = 1
                _ = v11
            default:
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
            }
        }
        if false {
            if (v1 && v4) {
                var v12 uint32 = 0
                v2 -= uint64(v3)
                _ = v12
            }
            v4 = ((522 != 0) || false)
        }
        fmt.Println(v3)
    } else {
        fmt.Println(v2)
    }
    v1 = true
    switch "a" {
    case "foo bar":
        v2 -= v2
    case "hello":
        if v1 {
            if v1 {
                switch 42 {
                case 1:
                    fmt.Println(v2, v3, v1)
                case 36:
                    fmt.Println(v1)
                case -2147483648:
                    v10 = append(v10, 303)
                default:
                    v2 = (v2 % (((v2 * (v2 & (v2 % ((genFn0((("\n\t" == ("\n\t" + ("a" + ((("hello" + "a") + (((("" + "hello") + "hello") + "a") + "\n\t")) + ("hello" + (((("a" + "hello") + "hello") + ("hello" + "\n\t")) + ("foo bar" + "foo bar"))))))) && (89 == 42)))) | 1))))) | 1))
                }
            } else {
                var v13 bool = ((-1 >= 32) || (v4 || v1))
                v10 = append(v10, -609)
                fmt.Println(v4)
                _ = v13
            }
            var v14 int = 837
            _ = v14
        }
    }
    v1 = (true || false)
    switch 614 {
    case 4294967295:
        v3 -= v3
    case 32768:
        for fi15 := 0; fi15 < 3; fi15++ {
            fmt.Println(fi15, v1)
            _ = fi15
        }
    case -1:
        fmt.Println(v1)
    }
    v16 := GenS2{F0: uint(v3), F1: 1, F2: -1}
    for fi17 := 0; fi17 < 17; fi17++ {
        var v18 int = fi17
        _ = fi17
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
