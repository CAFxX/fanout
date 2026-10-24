package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 uint16; F1 int32; F2 uint64 }
type GenS4 struct { F0 string; F1 int32 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int16 = 965
    var v3 int16 = -32768
    var v4 int = 32768
    v5 := make(map[uint32]int)
    switch (^(v4)) {
    case -2147483648:
        var v6 int8 = 1
        _ = v6
    }
    if false {
        v7 := GenS1{F0: 205}
        v8 := []int{v4, v4}
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 361}
    fmt.Println(v4)
    v10 := GenS1{F0: 65535}
    v2 += (v2 | 32767)
    var v11 uint16 = 65535
    v2 = (v3 | (v2 | v3))
    if true {
        var v12 uint = 567
        _ = v12
    } else {
        switch ("" + "") {
        case "":
            v13 := GenS1{F0: v11}
            _ = v13
        case "a":
            var v14 uint64 = 2
            _ = v14
        case "foo bar":
            v15 := GenS1{F0: (v11 << ((0) & 63))}
            _ = v15
        default:
            var v16 float32 = ((1e10 / 1.0) * ((3.14159 / 0.1) / -2.25))
            _ = v16
        }
        switch "a" {
        case "foo bar":
            v17 := GenS2{F0: (v2 >> ((0) & 63)), F1: (!false)}
            _ = v17
        }
        v3 = v2
    }
    var v18 uint8 = 127
    v2 = (v2 * (v2 & (v3 * v2)))
    v19 := GenS3{F0: 749, F1: -2147483648, F2: 2}
    fmt.Println(v2)
    switch (("\n\t" + "a") + (("foo bar" + "a") + ("\n\t" + ((("foo bar" + ("hello" + ("foo bar" + "foo bar"))) + "a") + "a")))) {
    case "hello":
        var v20 int16 = v2
        _ = v20
    case "foo bar":
        var v21 int64 = -484
        _ = v21
    case "a":
        var v22 int16 = (v3 ^ v2)
        _ = v22
    default:
        switch (v4 << ((65535) & 63)) {
        case 599:
            v3 = (v2 >> ((42) & 63))
        case 0:
            if true {
                var v23 int16 = 827
                var v24 int16 = v2
                _ = v23
                _ = v24
            }
        case -9223372036854775808:
            v25 := []int{}
            _ = v25
        }
    }
    v26 := GenS4{F0: "\n\t", F1: -1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
