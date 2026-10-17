package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int8 }
func genFn0(p0 uint, p1 uint64) string {
    for fi7 := 0; fi7 < 18; fi7++ {
        switch "" {
        case "hello":
            var v8 uint16 = 557
            _ = v8
        default:
            if false {
                var v9 uint32 = 65535
                _ = v9
            } else {
                v10 := []int{(fi7 + 4294967295), 1}
                p1 = (p1 >> (((p0 << ((0) & 63))) & 63))
                p0 = (p0 & p0)
                _ = v10
            }
        }
        _ = fi7
    }
    if false {
        v11 := GenS1{F0: p0, F1: 255}
        p1 = (p1 + 4294967295)
        var v12 string = "hello"
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    return "a"
}
func genFn1(p0 uint32, p1 int32, p2 uint8) bool {
    if true {
        if (643 > 2147483647) {
            p0 += p0
            if ((!(18446744073709551615 >= 25)) && false) {
                var v13 int = 32767
                _ = v13
            } else {
                var v14 int16 = -583
                _ = v14
            }
        } else {
            for fi15 := 0; fi15 < 10; fi15++ {
                fmt.Println(p2, fi15)
                var v16 string = ""
                p2 -= p2
                _ = fi15
                _ = v16
            }
            var v17 uint32 = p0
            _ = v17
        }
        v18 := GenS2{F0: 42}
        _ = v18
    }
    if (-1 >= -1) {
        p1 = int32(p0)
        if true {
            v19 := GenS3{F0: uint16(p0)}
            _ = v19
        }
    } else {
        var v20 int = -1
        p0 += (p0 * (p0 ^ (p0 << ((948) & 63))))
        _ = v20
    }
    v21 := GenS4{F0: 42}
    fmt.Println(p1, p2, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    return false
}

func main() {
    var v1 int32 = 0
    var v2 float32 = (3.14159 / 1.5)
    var v3 string = ("" + (("hello" + "hello") + "\n\t"))
    var v4 int16 = -1
    var v5 float64 = 0.1
    v6 := make(map[uint32]string)
    fmt.Println(v2, v5)
    v5 = 0.0
    var v22 uint8 = 255
    v22 -= (v22 + v22)
    var v23 uint16 = 1
    if false {
        switch v3 {
        case "foo bar":
            delete(v6, 971)
        }
    }
    fmt.Println(v23, v1)
    v22 -= 127
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
