package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 bool }
func genFn0() string {
    v7 := []int{}
    _ = len(v7)
    var v8 int32 = 0
    _ = v7
    _ = v8
    return "foo bar"
}
func genFn1(p0 int, p1 int32, p2 string) uint16 {
    v9 := []int{p0, (p0 | (p0 - 128)), (p0 / ((-2147483648) | 1)), (p0 % ((p0) | 1))}
    var v10 bool = (!(-32768 <= 370))
    p2 = "\n\t"
    fmt.Println(v10, v10, v10)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 0
}

func main() {
    var v1 int16 = 1000
    var v2 int8 = 1
    var v3 int8 = 0
    var v4 int = -9223372036854775808
    var v5 int32 = -930
    var v6 float64 = 3.14159
    v11 := GenS1{F0: 255, F1: -1}
    v12 := GenS2{F0: false}
    var v13 int16 = v1
    if ((((-2.25 + 1e10) * 0.0) * (float32(0.1) + float32((((((-2.25 + v6) / ((1.5 - 1e10) / 3.14159)) - v6) * v6) / float64(3.14159))))) != -2.25) {
        fmt.Println(v6)
    }
    var v14 int8 = (-(v2))
    v3 += v2
    v15 := []int{-2147483648, -640}
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    for fi16 := 0; fi16 < 6; fi16++ {
        v14 += (-(v14))
        for fi17 := 0; fi17 < 3; fi17++ {
            if ((false || false) || (515 < 255)) {
                switch (-(fi17)) {
                case -583:
                    var v18 int8 = (-(v3))
                    _ = v18
                }
            }
            var v19 uint = 0
            v14 = v14
            v20 := GenS2{F0: (32767 < 0)}
            var v21 int32 = v5
            var v22 float64 = v6
            _ = fi17
            _ = v19
            _ = v20
            _ = v21
            _ = v22
        }
        for fi23 := 0; fi23 < 14; fi23++ {
            v5 = v5
            if true {
                var v24 bool = (("a" == "") || (v14 > (v2 % ((-1) | 1))))
                fmt.Println(v2, v24)
                var v25 bool = (!(v1 <= (v1 >> ((42) & 63))))
                v26 := GenS1{F0: 2, F1: 9223372036854775807}
                _ = v24
                _ = v25
                _ = v26
            }
            _ = fi23
        }
        _ = fi16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
