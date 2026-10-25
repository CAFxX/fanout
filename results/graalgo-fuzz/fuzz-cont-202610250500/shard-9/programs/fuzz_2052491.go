package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 bool }
func genFn0(p0 uint16, p1 uint, p2 int8) uint {
    v4 := []int{2147483647}
    for fi5 := 0; fi5 < 3; fi5++ {
        v4 = append(v4, fi5)
        if len(v4) > 0 {
            v4[0] = fi5
        }
        _ = fi5
    }
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 uint32 = 1
        if (!(!(false || (true || true)))) {
            if true {
                var v8 int64 = 700
                fmt.Println(v8, v7, p0)
                _ = v8
            } else {
                var v9 bool = (183 < (p0 << ((42) & 63)))
                _ = v9
            }
            p1 = (p1 % (((p1 ^ (p1 * (^(p1))))) | 1))
        }
        _ = fi6
        _ = v7
    }
    p1 += p1
    if len(v4) > 0 {
        v4[0] = 0
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 2
}
func genFn1(p0 uint16, p1 string) int16 {
    v10 := []int{-9223372036854775808}
    v11 := GenS1{F0: 1, F1: 569, F2: false}
    p0 = p0
    p1 = "a"
    if false {
        var v12 uint16 = p0
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 408
}

func main() {
    var v1 uint32 = 195
    var v2 uint = 122
    var v3 uint64 = 4294967295
    v1 -= 0
    var v13 int8 = 99
    var v14 uint8 = 253
    switch "a" {
    case "foo bar":
        for fi15 := 0; fi15 < 2; fi15++ {
            var v16 string = ("foo bar" + (("foo bar" + "foo bar") + ((("" + "hello") + ("foo bar" + ("foo bar" + ""))) + ("hello" + "hello"))))
            _ = fi15
            _ = v16
        }
    case "a":
        for fi17 := 0; fi17 < 11; fi17++ {
            v1 = (v1 << ((genFn0(255, (v2 + (v2 & v2)), -40)) & 63))
            v13 += 14
            _ = fi17
        }
    case "":
        v2 += v2
    }
    v2 = (v2 >> (((v2 >> ((genFn0(42, v2, v13)) & 63))) & 63))
    var v18 int16 = int16(v13)
    v13 -= 0
    v3 = uint64(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
}
