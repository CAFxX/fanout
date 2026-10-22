package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 int8; F1 string; F2 uint16 }
func genFn0() string {
    v4 := []int{42, 32767}
    v5 := GenS1{F0: "hello", F1: false}
    var v6 uint16 = 255
    _ = v4
    _ = v5
    _ = v6
    return "\n\t"
}
func genFn1(p0 string, p1 int16) uint {
    p1 -= p1
    p1 += p1
    _ = p0
    _ = p1
    return 65535
}

func main() {
    var v1 bool = true
    var v2 int16 = 1
    v3 := make(map[uint32]int)
    v3[788] = -9223372036854775808
    switch 4294967295 {
    case -918:
        var v7 int8 = 127
        _ = v7
    case 42:
        for fi8 := 0; fi8 < 15; fi8++ {
            v2 = 984
            if (v1 && (128 <= 127)) {
                var v9 int16 = (v2 ^ v2)
                var v10 float64 = (3.14159 + 1e10)
                v11 := GenS2{F0: -128, F1: ((("foo bar" + "a") + ("hello" + (("\n\t" + ((("" + "a") + "a") + ("foo bar" + "hello"))) + ("" + "hello")))) + "a"), F2: 570}
                _ = v10
                _ = v11
                _ = v9
            } else {
                var v12 string = ""
                var v13 bool = (127 < 162)
                _ = len(v3)
                _ = v12
                _ = v13
            }
            _ = fi8
        }
    }
    v14 := GenS2{F0: -128, F1: "\n\t", F2: 446}
    for fi15 := 0; fi15 < 12; fi15++ {
        for fi16 := 0; fi16 < 16; fi16++ {
            var v17 uint64 = 329
            v1 = ((false && (v1 && (!(111 == 18446744073709551615)))) && (1 < 42))
            delete(v3, 2147483647)
            v17 = v17
            v2 = (v2 << ((18446744073709551615) & 63))
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    v2 -= v2
    v1 = (3.14159 == -2.25)
    v1 = (false || (v1 && true))
    for fi18 := 0; fi18 < 20; fi18++ {
        fmt.Println(v1, v1, v2)
        v1 = true
        v2 = 0
        _ = fi18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v2
    _ = v3
}
