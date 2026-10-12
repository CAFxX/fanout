package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 string; F2 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32; F1 int8 }
func genFn0(p0 uint8, p1 uint64, p2 int64) bool {
    switch "foo bar" {
    case "\n\t":
        if false {
            v4 := GenS1{F0: 65535}
            v5 := GenS2{F0: (p2 - (p2 ^ 42)), F1: "", F2: 255}
            v6 := []int{0, int(p1), -796}
            _ = v4
            _ = v5
            _ = v6
        }
    }
    v7 := GenS1{F0: 426}
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (!(-2.25 <= -2.25))
}
func genFn1() string {
    v8 := []int{9223372036854775807, 9223372036854775807, 9223372036854775807, 1}
    switch 128 {
    case 4294967295:
        var v9 int32 = -254
        _ = v9
    default:
        var v10 uint = 0
        _ = v10
    }
    var v11 uint = 1
    _ = v11
    _ = v8
    return ("foo bar" + "")
}

func main() {
    var v1 int16 = 428
    var v2 int64 = 42
    var v3 int64 = (v2 >> ((1) & 63))
    v2 = v2
    if true {
        v3 = 349
        for fi12 := 0; fi12 < 10; fi12++ {
            var v13 string = "a"
            fmt.Println(v13, v13, v3)
            _ = fi12
            _ = v13
        }
    }
    v2 = v3
    fmt.Println(v1, v1)
    v3 = (v2 >> ((0) & 63))
    v1 -= v1
    v3 += (v3 * v2)
    v14 := []int{42, int(v1), -2147483648, -9223372036854775808}
    for fi15 := 0; fi15 < 12; fi15++ {
        for fi16 := 0; fi16 < 6; fi16++ {
            if (171 < 0) {
                var v17 string = (("\n\t" + "") + "\n\t")
                _ = v17
            }
            _ = fi16
        }
        _ = fi15
    }
    fmt.Println(v2, v1, v1)
    switch "hello" {
    case "":
        var v18 int8 = -77
        _ = v18
    case "hello":
        if (!(938 >= 215)) {
            for fi19 := 0; fi19 < 4; fi19++ {
                var v20 int64 = -920
                v20 += v2
                _ = fi19
                _ = v20
            }
            if (127 <= -128) {
                v1 -= (v1 * (v1 - (v1 << ((18446744073709551615) & 63))))
            }
        } else {
            if false {
                if false {
                    var v21 int32 = -1
                    _ = v21
                }
                v22 := GenS3{F0: 42}
                _ = v22
            } else {
                for fi23 := 0; fi23 < 8; fi23++ {
                    fmt.Println(fi23, v2)
                    v14 = append(v14, fi23)
                    _ = fi23
                }
                v3 = v3
            }
        }
    }
    if true {
        v24 := GenS4{F0: -448, F1: 0}
        _ = v24
    }
    var v25 uint = 65535
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v2
    _ = v25
    _ = v3
}
