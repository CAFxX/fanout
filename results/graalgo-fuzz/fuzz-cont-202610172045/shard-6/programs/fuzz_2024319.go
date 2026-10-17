package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int32 }
type GenS2 struct { F0 uint; F1 uint16 }

func main() {
    v1 := make(map[int]int)
    var v2 float32 = (-2.25 / (-2.25 * ((3.14159 + ((float32(float64(0.1)) - -2.25) * -2.25)) / (-2.25 * float32(1.5)))))
    var v3 uint16 = 32767
    v4 := []int{9223372036854775807, 700}
    var v5 int = -9223372036854775808
    v6 := []string{"", "foo bar"}
    v3 -= (v3 ^ 0)
    var v7 float64 = 0.1
    var v8 string = "foo bar"
    var v9 bool = (42 != 32767)
    for fi10 := 0; fi10 < 15; fi10++ {
        v11 := GenS1{F0: fi10, F1: 785, F2: -2147483648}
        for fi12 := 0; fi12 < 13; fi12++ {
            v13 := GenS2{F0: 197, F1: (v3 ^ (v3 >> ((65535) & 63)))}
            v3 = (v3 + 2)
            _ = fi12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    var v14 int32 = 191
    switch "a" {
    case "":
        v7 += v7
    default:
        if v9 {
            switch "hello" {
            case "\n\t":
                var v15 uint = 1
                _ = v15
            case "":
                v3 = 1
            case "foo bar":
                for fi16 := 0; fi16 < 15; fi16++ {
                    v17 := GenS1{F0: v5, F1: 42, F2: (v14 - 230)}
                    var v18 string = (v8 + v8)
                    switch "a" {
                    case "\n\t":
                        v1[v5] = (^(v5))
                    default:
                        v14 -= -1
                    }
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            }
        }
    }
    v2 = 1.5
    if (18446744073709551615 > 4294967295) {
        delete(v1, (v5 % (((v5 << ((193) & 63))) | 1)))
        for fi19 := 0; fi19 < 18; fi19++ {
            if (v9 || ((!(float32(v7) != 0.0)) || v9)) {
                switch (v8 + v8) {
                case "hello":
                    fmt.Println(v2, v2, v8)
                case "a":
                    v9 = v9
                case "foo bar":
                    var v20 int16 = int16(v3)
                    _ = v20
                }
                if v9 {
                    if v9 {
                        var v21 int8 = 1
                        fmt.Println(v2)
                        var v22 uint = 503
                        _ = v21
                        _ = v22
                    } else {
                        for fi23 := 0; fi23 < 20; fi23++ {
                            v7 -= v7
                            _ = fi23
                        }
                    }
                    v2 += v2
                }
            }
            var v24 int = (v5 << ((65535) & 63))
            _ = fi19
            _ = v24
        }
    } else {
        for fi25 := 0; fi25 < 19; fi25++ {
            fmt.Println(v9)
            var v26 uint8 = 240
            _ = fi25
            _ = v26
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    v9 = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
