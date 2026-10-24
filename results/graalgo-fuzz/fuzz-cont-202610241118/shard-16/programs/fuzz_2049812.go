package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 int16; F2 string }
type GenS4 struct { F0 int64 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]int)
    var v3 int = -167
    v4 := GenS1{F0: -415, F1: 57, F2: 127}
    if (-157 <= 1) {
        _ = len(v2)
        v3 -= -224
    }
    v5 := GenS1{F0: -996, F1: 2, F2: 0}
    v3 = v3
    if true {
        var v6 int = (v3 + (v3 % ((32767) | 1)))
        var v7 string = "hello"
        _ = v6
        _ = v7
    }
    v8 := GenS2{F0: 65535}
    fmt.Println(v3)
    v3 -= (v3 + (v3 + (v3 << ((68) & 63))))
    switch "" {
    case "hello":
        v3 -= (-(v3))
    }
    if true {
        v9 := v2[2147483647]
        if (-1 > 100) {
            var v10 uint32 = 4294967295
            _ = v10
        } else {
            fmt.Println(v9, v3)
            for fi11 := 0; fi11 < 20; fi11++ {
                v12 := GenS3{F0: -775, F1: 881, F2: ""}
                fmt.Println(fi11)
                _ = fi11
                _ = v12
            }
        }
        if false {
            if true {
                if (1 >= -48) {
                    if false {
                        v13 := GenS4{F0: -509}
                        v3 -= v9
                        _ = v13
                    }
                }
            } else {
                if (727 == 1) {
                    switch (v3 << ((216) & 63)) {
                    case -2147483648:
                        var v14 int64 = -9223372036854775808
                        _ = v14
                    }
                    if (!(32767 <= -32768)) {
                        v15 := GenS1{F0: 1000, F1: 4294967295, F2: 99}
                        var v16 float64 = 0.1
                        for fi17 := 0; fi17 < 20; fi17++ {
                            v18 := GenS3{F0: 42, F1: 1000, F2: "a"}
                            v16 += (v16 + ((v16 / ((v16 / -2.25) + (v16 / 0.1))) * 1.5))
                            var v19 string = ("" + "")
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                        _ = v15
                        _ = v16
                    }
                }
                if (v9 >= v3) {
                    fmt.Println(v9, v3)
                } else {
                    for fi20 := 0; fi20 < 20; fi20++ {
                        v3 = -552
                        if true {
                            var v21 uint8 = 128
                            _ = v21
                        } else {
                            var v22 int32 = -664
                            _ = v22
                        }
                        _ = fi20
                    }
                }
                v9 = v3
            }
            fmt.Println(v3, v9)
        }
        _ = v9
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
