package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64; F1 int32 }
func genFn0(p0 string) bool {
    var v6 uint8 = 127
    var v7 uint8 = v6
    var v8 uint32 = 4294967295
    var v9 int64 = -1
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return true
}
func genFn1(p0 uint8, p1 uint64, p2 int64) string {
    var v10 float32 = 3.14159
    v11 := []int{130, -9223372036854775808, 0, 42}
    v12 := GenS1{F0: true, F1: (p1 << ((18446744073709551615) & 63))}
    var v13 uint = 669
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    return "foo bar"
}

func main() {
    v1 := []uint32{4294967295}
    var v2 float64 = 0.1
    var v3 uint64 = 1
    var v4 int64 = 9223372036854775807
    v5 := []string{("foo bar" + ("\n\t" + "hello"))}
    for fi14 := 0; fi14 < 12; fi14++ {
        v4 += (-(v4))
        v4 += (v4 >> ((65535) & 63))
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi14
    }
    var v15 uint16 = 65535
    for fi16 := 0; fi16 < 18; fi16++ {
        if true {
            v17 := GenS1{F0: false, F1: v3}
            _ = v17
        } else {
            fmt.Println(v4, fi16)
        }
        _ = fi16
    }
    v2 -= 0.0
    var v18 float32 = ((0.1 / 1.5) / float32((v2 - float64(1.5))))
    v19 := GenS1{F0: ((0.1 - (v18 + ((v18 + -2.25) - (float32(1e10) + v18)))) == v18), F1: 0}
    fmt.Println(v2, v15)
    var v20 uint32 = 4294967295
    v20 -= uint32(v3)
    for fi21 := 0; fi21 < 10; fi21++ {
        var v22 uint32 = v20
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 10; fi23++ {
        v15 = (v15 >> ((42) & 63))
        _ = fi23
    }
    v2 = (v2 - 3.14159)
    if false {
        if (42 < 1) {
            v24 := GenS2{F0: -1}
            if ((!((uint(v20) == 0) || true)) && true) {
                switch "hello" {
                case "\n\t":
                    fmt.Println(v18, v20)
                case "hello":
                    var v25 int16 = -350
                    _ = v25
                case "foo bar":
                    var v26 bool = true
                    _ = v26
                default:
                    fmt.Println(v15, v15)
                }
                v27 := GenS1{F0: (!true), F1: (v3 << ((255) & 63))}
                v28 := GenS3{F0: v3, F1: 76}
                _ = v27
                _ = v28
            }
            if false {
                var v29 int64 = v4
                _ = v29
            }
            _ = v24
        } else {
            var v30 string = "\n\t"
            var v31 uint8 = 2
            _ = v30
            _ = v31
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
