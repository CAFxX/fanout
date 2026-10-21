package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint; F1 uint64 }
func genFn0() int16 {
    var v4 int = 2147483647
    v4 = (v4 + v4)
    fmt.Println(v4)
    switch (v4 * v4) {
    case -9223372036854775808:
        var v5 string = "foo bar"
        _ = v5
    }
    var v6 string = "a"
    _ = v4
    _ = v6
    return -1
}
func genFn1(p0 uint64, p1 string) string {
    for fi7 := 0; fi7 < 13; fi7++ {
        p0 = (p0 / ((p0) | 1))
        v8 := []int{(fi7 % (((-(fi7))) | 1)), fi7}
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 12; fi9++ {
        p0 = (p0 >> ((42) & 63))
        if (1 < 942) {
            v10 := GenS1{F0: p1, F1: "a", F2: ((p1 + (p1 + (((p1 + ((p1 + ("foo bar" + p1)) + "a")) + p1) + p1))) + "hello")}
            var v11 int64 = -1
            _ = v10
            _ = v11
        }
        v12 := GenS1{F0: "hello", F1: "\n\t", F2: p1}
        _ = fi9
        _ = v12
    }
    var v13 float64 = 0.0
    _ = p0
    _ = p1
    _ = v13
    return ""
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 uint64 = 873
    var v3 string = ("foo bar" + ("\n\t" + "\n\t"))
    fmt.Println(v3)
    v2 = 172
    v1 -= -9223372036854775808
    if (560 != -218) {
        if false {
            for fi14 := 0; fi14 < 17; fi14++ {
                v15 := GenS1{F0: (((("" + (v3 + v3)) + v3) + v3) + v3), F1: v3, F2: "\n\t"}
                var v16 uint = 490
                v17 := GenS1{F0: (("hello" + v3) + (v3 + "")), F1: "hello", F2: v3}
                _ = fi14
                _ = v15
                _ = v16
                _ = v17
            }
            var v18 int = -205
            v3 = (v3 + v3)
            _ = v18
        } else {
            v1 += v1
            for fi19 := 0; fi19 < 20; fi19++ {
                var v20 uint8 = 111
                var v21 uint16 = 211
                _ = fi19
                _ = v20
                _ = v21
            }
            v22 := GenS2{F0: 309}
            _ = v22
        }
    }
    var v23 float32 = 3.14159
    v3 = "\n\t"
    fmt.Println(v23, v3)
    v24 := GenS3{F0: 18446744073709551615, F1: v2}
    var v25 uint32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
