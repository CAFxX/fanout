package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int }
func genFn0() uint64 {
    var v5 int = 4294967295
    v6 := GenS1{F0: -1, F1: 1, F2: 32767}
    var v7 uint16 = 255
    v8 := GenS1{F0: -9223372036854775808, F1: 4294967295, F2: 0}
    var v9 int = v5
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 0
}
func genFn1(p0 uint32, p1 uint16) int {
    var v10 int64 = -9223372036854775808
    p0 -= (p0 | p0)
    _ = p0
    _ = p1
    _ = v10
    return 128
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{"\n\t", ("" + "foo bar"), (("foo bar" + ("foo bar" + (("hello" + (("hello" + ("a" + ("hello" + ("\n\t" + "a")))) + "\n\t")) + "\n\t"))) + ("hello" + (("\n\t" + "\n\t") + "a")))}
    var v3 uint16 = 1
    var v4 bool = ("hello" <= "foo bar")
    v11 := GenS1{F0: 521, F1: 508, F2: v3}
    v3 -= v3
    v12 := v1[65535]
    v12 = v12
    v3 += (v3 - 42)
    v3 -= (v3 + 437)
    var v13 uint32 = 472
    var v14 int32 = 2147483647
    _ = len(v1)
    var v15 uint = 676
    if len(v2) > 0 {
        v2[0] = (("hello" + "hello") + v12)
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v13 = (^(v13))
    if v4 {
        for fi16 := 0; fi16 < 15; fi16++ {
            switch (-(fi16)) {
            case 0:
                v17 := v1[v13]
                _ = v17
            case 727:
                fmt.Println(fi16)
            case 32768:
                v18 := GenS2{F0: 0}
                _ = v18
            }
            v12 = "a"
            _ = fi16
        }
    } else {
        v3 = 858
        if v4 {
            var v19 int32 = 2147483647
            var v20 int = -85
            for fi21 := 0; fi21 < 8; fi21++ {
                switch 127 {
                case -699:
                    v22 := GenS1{F0: 9223372036854775807, F1: (v15 << ((uint(v13)) & 63)), F2: 42}
                    _ = v22
                default:
                    v3 += v3
                }
                if v4 {
                    fmt.Println(v20)
                } else {
                    v14 -= (v14 >> ((255) & 63))
                }
                v12 = "a"
                _ = fi21
            }
            _ = v19
            _ = v20
        } else {
            if v4 {
                for fi23 := 0; fi23 < 13; fi23++ {
                    var v24 float32 = (float32(float64(0.0)) * 1.5)
                    _ = fi23
                    _ = v24
                }
                v14 = 42
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v2 = append(v2, (v12 + (v12 + "a")))
        }
        v25 := GenS3{F0: 231}
        _ = v25
    }
    v12 = ("\n\t" + v12)
    var v26 int32 = 972
    var v27 int16 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
