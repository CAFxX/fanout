package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 int32 }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 string; F1 string }
func genFn0(p0 int64, p1 int16, p2 string) uint16 {
    var v5 int32 = -303
    v6 := GenS1{F0: (p2 + ("" + p2)), F1: 0, F2: v5}
    p1 = p1
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 255
}
func genFn1(p0 int32) int {
    v7 := GenS2{F0: 511, F1: 736}
    if (true && (857 <= 255)) {
        fmt.Println(p0)
        v8 := []int{-1, 42, 0}
        var v9 int16 = 32767
        _ = v8
        _ = v9
    } else {
        p0 = (p0 % ((p0) | 1))
        v10 := []int{-1, 128, int(p0), 567}
        _ = v10
    }
    if true {
        var v11 string = "hello"
        _ = v11
    } else {
        for fi12 := 0; fi12 < 5; fi12++ {
            switch "a" {
            case "\n\t":
                var v13 bool = ((("\n\t" + ("foo bar" + (("hello" + ("\n\t" + (("hello" + "foo bar") + "foo bar"))) + (("" + "") + "")))) + (("" + "hello") + ((("a" + "foo bar") + ("" + ("\n\t" + "foo bar"))) + ("hello" + "\n\t")))) <= (("hello" + "a") + ""))
                _ = v13
            }
            _ = fi12
        }
        var v14 float64 = 1e10
        _ = v14
    }
    _ = p0
    _ = v7
    return 4294967295
}

func main() {
    var v1 uint = 65535
    var v2 bool = (false && (0.0 > 0.0))
    var v3 int = -9223372036854775808
    var v4 uint32 = uint32(v1)
    switch (v3 - (v3 >> (((v1 * (v1 ^ 255))) & 63))) {
    case -1:
        var v15 int = v3
        _ = v15
    case 128:
        var v16 uint = 65535
        _ = v16
    case 127:
        v3 += v3
    }
    v2 = (2147483647 >= 648)
    var v17 int8 = 42
    v18 := []int{(v3 << (((v1 + (v1 ^ v1))) & 63)), 1, (^(v3))}
    v19 := GenS3{F0: -1}
    v1 = v1
    v3 -= (v3 << ((v1) & 63))
    v20 := GenS4{F0: "\n\t", F1: ("foo bar" + "\n\t")}
    var v21 int64 = -450
    if len(v18) > 0 {
        v18[0] = (-(v3))
    }
    var v22 int8 = 127
    fmt.Println(v1, v2)
    switch ("" + "foo bar") {
    case "a":
        v1 = (v1 * v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
