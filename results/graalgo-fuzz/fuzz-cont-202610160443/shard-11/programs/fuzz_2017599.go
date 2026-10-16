package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 uint8 }
type GenS2 struct { F0 uint32; F1 uint8; F2 int }
type GenS3 struct { F0 uint }
func genFn0(p0 int32, p1 int) string {
    v7 := []int{(p1 >> ((113) & 63)), (p1 % (((p1 | 32767)) | 1)), p1}
    var v8 string = "a"
    var v9 uint16 = 42
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return v8
}
func genFn1(p0 int16, p1 int64) string {
    v10 := GenS1{F0: p1, F1: 830, F2: 255}
    p1 -= 621
    var v11 float64 = 1e10
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return "foo bar"
}

func main() {
    var v1 string = ("\n\t" + ("hello" + "\n\t"))
    var v2 uint16 = 255
    var v3 uint32 = 4294967295
    var v4 float32 = -2.25
    var v5 int8 = -49
    var v6 uint64 = 4294967295
    switch v1 {
    case "hello":
        switch (v1 + (("a" + v1) + "a")) {
        case "\n\t":
            if (v3 < (v3 >> ((18446744073709551615) & 63))) {
                for fi12 := 0; fi12 < 12; fi12++ {
                    switch (fi12 * (fi12 << ((997) & 63))) {
                    case 4294967295:
                        var v13 uint16 = 255
                        _ = v13
                    case -9223372036854775808:
                        for fi14 := 0; fi14 < 1; fi14++ {
                            v2 = v2
                            if false {
                                v15 := []int{-1}
                                v5 = (-(v5))
                                _ = v15
                            }
                            v16 := []int{261, 2147483647}
                            _ = fi14
                            _ = v16
                        }
                    default:
                        var v17 uint8 = 127
                        _ = v17
                    }
                    var v18 int16 = -379
                    fmt.Println(v18)
                    _ = fi12
                    _ = v18
                }
                var v19 int8 = -128
                _ = v19
            }
        }
    }
    if (-2147483648 > 2147483647) {
        v20 := GenS1{F0: 1, F1: v2, F2: 127}
        _ = v20
    } else {
        var v21 uint32 = (^(v3))
        switch ((v1 + "\n\t") + v1) {
        case "a":
            v22 := GenS2{F0: v3, F1: 2, F2: 87}
            _ = v22
        case "hello":
            if false {
                fmt.Println(v3, v3, v5)
            }
        case "\n\t":
            fmt.Println(v6, v4)
        default:
            for fi23 := 0; fi23 < 4; fi23++ {
                v24 := []int{}
                var v25 int = (fi23 * -414)
                _ = fi23
                _ = v24
                _ = v25
            }
        }
        _ = v21
    }
    v26 := GenS3{F0: 42}
    if (!true) {
        v5 += (v5 - v5)
        for fi27 := 0; fi27 < 2; fi27++ {
            v28 := []int{(-(fi27)), 42}
            _ = fi27
            _ = v28
        }
    }
    var v29 bool = (!false)
    v5 += v5
    var v30 string = "hello"
    v2 -= (v2 >> ((65535) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
