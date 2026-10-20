package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64; F1 uint }
func genFn0(p0 string, p1 uint16) int {
    v5 := []int{}
    v6 := GenS1{F0: -612, F1: false, F2: 42}
    if true {
        v7 := GenS2{F0: 255}
        _ = v7
    }
    p1 += (^(p1))
    var v8 uint64 = 42
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v8
    return 592
}

func main() {
    var v1 int16 = -258
    var v2 uint32 = 520
    var v3 uint64 = 1
    var v4 int8 = 1
    v9 := GenS2{F0: 240}
    for fi10 := 0; fi10 < 8; fi10++ {
        v3 -= 0
        v2 = 2147483647
        _ = fi10
    }
    fmt.Println(v2)
    switch "a" {
    case "a":
        if ("foo bar" > ("" + "\n\t")) {
            fmt.Println(v1, v3, v3)
        }
    case "\n\t":
        v11 := GenS3{F0: -9223372036854775808, F1: 42}
        _ = v11
    default:
        v12 := GenS3{F0: 42, F1: 191}
        _ = v12
    }
    var v13 uint32 = (^(v2))
    v14 := []int{}
    var v15 uint8 = 222
    _ = len(v14)
    v2 = 1
    if (true && true) {
        var v16 uint16 = 32767
        _ = v16
    }
    if (2 != v15) {
        switch ("\n\t" + "hello") {
        case "foo bar":
            v13 -= (v13 | uint32(v15))
        case "a":
            if (2147483648 >= 42) {
                var v17 uint16 = 65535
                if len(v14) > 0 {
                    v14[0] = genFn0((("a" + "foo bar") + "foo bar"), (v17 - (v17 % ((v17) | 1))))
                }
                _ = v17
            }
        default:
            for fi18 := 0; fi18 < 8; fi18++ {
                v15 += (v15 >> ((18446744073709551615) & 63))
                if true {
                    v19 := GenS3{F0: -9223372036854775808, F1: 0}
                    v2 -= v2
                    _ = v19
                } else {
                    var v20 string = "\n\t"
                    _ = v20
                }
                var v21 uint8 = (^(v15))
                _ = fi18
                _ = v21
            }
        }
    }
    v22 := GenS3{F0: -1, F1: 255}
    v23 := GenS2{F0: 531}
    if len(v14) > 0 {
        v14[0] = 9223372036854775807
    }
    v4 -= (v4 % ((v4) | 1))
    if false {
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        v3 -= 2
        v2 += 2
    }
    switch genFn0("a", 255) {
    case 32767:
        var v24 float64 = float64(1.5)
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v9
}
