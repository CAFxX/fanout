package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint; F1 bool; F2 int8 }
func genFn0(p0 uint32, p1 uint, p2 uint32) uint {
    v7 := GenS1{F0: 272, F1: (p1 >= 139)}
    v8 := []int{127, 9223372036854775807, 0}
    p1 -= 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 79
}
func genFn1(p0 uint8, p1 string) uint64 {
    p0 += (p0 - 2)
    v9 := GenS1{F0: 65535, F1: false}
    v10 := []int{-473, -9223372036854775808, 32768, 9223372036854775807}
    if ((("foo bar" + ("foo bar" + (p1 + p1))) >= p1) || (!(false || ((18446744073709551615 == 4294967295) && (1e10 < (1e10 + (3.14159 * -2.25))))))) {
        v11 := GenS1{F0: 533, F1: (!(true && true))}
        if true {
            if false {
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                if (true && (32767 != 221)) {
                    for fi12 := 0; fi12 < 2; fi12++ {
                        var v13 uint8 = (p0 * 42)
                        _ = fi12
                        _ = v13
                    }
                    for fi14 := 0; fi14 < 16; fi14++ {
                        switch (fi14 | fi14) {
                        case -315:
                            var v15 uint16 = 65535
                            _ = v15
                        case -9223372036854775808:
                            p1 = p1
                        case 2147483647:
                            v16 := GenS1{F0: 1, F1: (((!true) && (false || (3.14159 != 0.1))) || false)}
                            _ = v16
                        }
                        if (548 > -1) {
                            p0 += p0
                        } else {
                            switch p1 {
                            case "a":
                                var v17 uint8 = (p0 / ((p0) | 1))
                                _ = v17
                            case "foo bar":
                                v18 := GenS2{F0: -71}
                                _ = v18
                            default:
                                fmt.Println(p1)
                            }
                            v19 := GenS1{F0: 230, F1: false}
                            _ = v19
                        }
                        _ = fi14
                    }
                    p1 = ("" + "hello")
                }
                v10 = append(v10, -344)
            }
        } else {
            if len(v10) > 0 {
                v10[0] = -392
            }
            var v20 uint16 = 392
            p1 = p1
            _ = v20
        }
        _ = v11
    }
    v21 := GenS2{F0: -25}
    _ = p0
    _ = p1
    _ = v10
    _ = v21
    _ = v9
    return uint64(p0)
}

func main() {
    var v1 int64 = -1
    v2 := make(map[int]uint32)
    var v3 int64 = (v1 - 1)
    v4 := make(map[int]string)
    var v5 uint32 = 1
    var v6 uint32 = (v5 % ((0) | 1))
    var v22 float64 = ((0.0 * 1e10) + float64(float32(-2.25)))
    delete(v4, 32768)
    fmt.Println(v22)
    v23 := v2[14]
    switch ((("hello" + (("\n\t" + ("" + "foo bar")) + "a")) + ("foo bar" + (("\n\t" + (((("foo bar" + "foo bar") + ((("foo bar" + "foo bar") + "a") + "foo bar")) + "") + "foo bar")) + "\n\t"))) + "foo bar") {
    case "a":
        v24 := GenS3{F0: 255, F1: false, F2: 46}
        _ = v24
    case "foo bar":
        if (2 <= 876) {
            var v25 float64 = (v22 / v22)
            _ = v25
        } else {
            v5 -= (v5 | (v23 | (v6 / (((v6 - (v23 << ((582) & 63)))) | 1))))
            var v26 uint16 = 32767
            _ = v26
        }
    }
    fmt.Println(v6)
    fmt.Println(v6)
    if (true && (true || false)) {
        v27 := GenS2{F0: 100}
        _ = v27
    }
    for fi28 := 0; fi28 < 5; fi28++ {
        v29 := []int{}
        if len(v29) >= 2 {
            v29 = v29[1:len(v29)]
        }
        _ = fi28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
