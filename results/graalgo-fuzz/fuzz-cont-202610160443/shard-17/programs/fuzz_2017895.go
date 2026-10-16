package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 string }
func genFn0(p0 int32, p1 uint32, p2 string) string {
    v4 := []int{128}
    _ = len(v4)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return ""
}
func genFn1(p0 int, p1 int) int32 {
    v5 := GenS1{F0: 127, F1: 186}
    var v6 int64 = 1
    fmt.Println(p0)
    var v7 uint8 = 128
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return -40
}

func main() {
    var v1 uint8 = 200
    var v2 string = ("\n\t" + "hello")
    var v3 float32 = ((0.1 * ((-2.25 / 1e10) - -2.25)) / 1.5)
    v8 := []int{}
    var v9 int8 = -1
    fmt.Println(v2, v2, v3)
    if false {
        v10 := GenS1{F0: v1, F1: 2}
        _ = v10
    } else {
        var v11 uint32 = 2147483648
        _ = v11
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        if ((942 < 92) || (v9 != v9)) {
            switch v2 {
            case "foo bar":
                v3 -= float32(1e10)
            case "hello":
                v9 -= (v9 >> ((0) & 63))
            case "a":
                v13 := GenS2{F0: v2}
                _ = v13
            default:
                for fi14 := 0; fi14 < 11; fi14++ {
                    v9 = v9
                    _ = fi14
                }
            }
            var v15 string = v2
            _ = v15
        }
        if true {
            switch (v2 + ("\n\t" + v2)) {
            case "a":
                if (32767 < 32767) {
                    _ = len(v8)
                }
            }
            v8 = append(v8, (fi12 + (^(fi12))))
            for fi16 := 0; fi16 < 13; fi16++ {
                v17 := GenS1{F0: v1, F1: 818}
                _ = fi16
                _ = v17
            }
        } else {
            for fi18 := 0; fi18 < 16; fi18++ {
                v9 = (^(v9))
                if ((-(v9)) != -128) {
                    v9 += -1
                    var v19 int64 = -42
                    v2 = "foo bar"
                    _ = v19
                }
                _ = fi18
            }
        }
        var v20 int16 = 32767
        _ = fi12
        _ = v20
    }
    v21 := GenS2{F0: ("foo bar" + ("foo bar" + "hello"))}
    v22 := GenS1{F0: v1, F1: 34}
    fmt.Println(v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v8
    _ = v9
}
