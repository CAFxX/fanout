package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 string; F1 int16; F2 int16 }
type GenS3 struct { F0 int16; F1 string; F2 uint8 }
func genFn0(p0 bool) int64 {
    p0 = p0
    p0 = (719 <= 461)
    p0 = p0
    v7 := []int{32768, 780}
    _ = p0
    _ = v7
    return -1
}
func genFn1(p0 uint32, p1 int32) string {
    var v8 int32 = (p1 ^ (-(p1)))
    p0 = (^(p0))
    var v9 int = 42
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return ("" + ("a" + ("hello" + "foo bar")))
}

func main() {
    var v1 uint16 = 2
    var v2 uint8 = 42
    var v3 float32 = (float32(-2.25) / 0.1)
    var v4 uint64 = 18446744073709551615
    var v5 string = "a"
    var v6 int = 128
    var v10 float32 = v3
    v11 := GenS1{F0: 2, F1: -1}
    v4 += (^(v4))
    if false {
        if (1.5 >= (-2.25 + 0.1)) {
            for fi12 := 0; fi12 < 12; fi12++ {
                v13 := GenS1{F0: 18446744073709551615, F1: 9223372036854775807}
                v14 := []int{(fi12 * v6), v6}
                _ = fi12
                _ = v13
                _ = v14
            }
            for fi15 := 0; fi15 < 7; fi15++ {
                v1 -= 370
                for fi16 := 0; fi16 < 11; fi16++ {
                    if (127 > 100) {
                        var v17 int16 = 42
                        var v18 uint64 = v4
                        _ = v17
                        _ = v18
                    } else {
                        fmt.Println(fi16)
                        v19 := GenS2{F0: ("" + ("hello" + (((v5 + "hello") + ("hello" + (("foo bar" + "") + v5))) + (v5 + v5)))), F1: -32768, F2: -962}
                        _ = v19
                    }
                    fmt.Println(v1)
                    v20 := GenS1{F0: 4294967295, F1: 508}
                    _ = fi16
                    _ = v20
                }
                v4 = (v4 - (v4 * (v4 >> ((18446744073709551615) & 63))))
                _ = fi15
            }
        }
        if (!((v6 % (((v6 | (v6 >> ((65535) & 63)))) | 1)) == (v6 % (((v6 % ((v6) | 1))) | 1)))) {
            v21 := []int{(v6 * v6), v6}
            _ = v21
        }
    }
    v22 := []int{(v6 >> ((450) & 63)), v6, (v6 % ((-1) | 1)), v6}
    switch (v6 >> ((2) & 63)) {
    case 9223372036854775807:
        fmt.Println(v2)
    case -69:
        v23 := GenS3{F0: -32768, F1: v5, F2: v2}
        _ = v23
    case 68:
        v24 := GenS1{F0: 0, F1: (v6 & v6)}
        _ = v24
    default:
        if len(v22) > 0 {
            v22[0] = (v6 >> ((65535) & 63))
        }
    }
    var v25 uint64 = (v4 - v4)
    v6 += (-(v6))
    fmt.Println(v5, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
