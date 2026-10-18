package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 int64; F2 int }
func genFn0(p0 int32, p1 int16) string {
    v6 := []int{582, -9223372036854775808}
    var v7 uint32 = 2147483647
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 string = ""
        p0 += (p0 - -2147483648)
        switch (fi8 % ((-2147483648) | 1)) {
        case 127:
            if (true || (v7 <= v7)) {
                fmt.Println(fi8)
            } else {
                var v10 float32 = 0.1
                _ = v10
            }
        case 817:
            switch v9 {
            case "\n\t":
                var v11 int = fi8
                _ = v11
            case "":
                var v12 uint8 = uint8(v7)
                _ = v12
            default:
                fmt.Println(p1, v7, v9)
            }
        default:
            var v13 int64 = 42
            _ = v13
        }
        _ = fi8
        _ = v9
    }
    if (543 > 1) {
        v7 -= v7
        if ("\n\t" != ("hello" + ("hello" + ("hello" + ((("foo bar" + "foo bar") + "a") + "hello"))))) {
            v14 := GenS1{F0: 1}
            _ = v14
        }
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return "hello"
}
func genFn1() int16 {
    v15 := []int{-1}
    var v16 uint8 = 128
    var v17 uint16 = 32767
    switch -619 {
    case 2147483647:
        v18 := GenS1{F0: 54}
        _ = v18
    }
    if true {
        v19 := GenS2{F0: 0, F1: -1, F2: 534}
        var v20 uint8 = v16
        if len(v15) > 0 {
            v15[0] = -835
        }
        _ = v19
        _ = v20
    }
    _ = v15
    _ = v16
    _ = v17
    return 42
}

func main() {
    var v1 float64 = (-2.25 / 3.14159)
    var v2 int16 = -1
    v3 := make(map[int]int)
    v4 := []string{}
    v5 := []string{("a" + "foo bar"), (("\n\t" + "a") + ("" + "foo bar"))}
    var v21 float64 = float64(((3.14159 / 1e10) / 1.5))
    var v22 int32 = -244
    var v23 int16 = (v2 >> ((108) & 63))
    switch 127 {
    case -9223372036854775808:
        v24 := GenS2{F0: v22, F1: 42, F2: 0}
        _ = v24
    default:
        if true {
            var v25 uint = 1
            _ = v25
        } else {
            for fi26 := 0; fi26 < 17; fi26++ {
                switch 9223372036854775807 {
                case 4294967295:
                    if true {
                        if (-128 >= 127) {
                            v27 := GenS1{F0: -128}
                            var v28 int = (fi26 + 1)
                            _ = v27
                            _ = v28
                        } else {
                            var v29 float32 = -2.25
                            if len(v5) > 0 {
                                v5[0] = ("hello" + ("" + ("" + "")))
                            }
                            _ = v29
                        }
                        for fi30 := 0; fi30 < 19; fi30++ {
                            var v31 int32 = (-(v22))
                            _ = fi30
                            _ = v31
                        }
                    }
                case 128:
                    v22 = (v22 >> ((585) & 63))
                case -9223372036854775808:
                    for fi32 := 0; fi32 < 10; fi32++ {
                        v2 -= v23
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        if (920 > 496) {
                            fmt.Println(v22, v2, v1)
                            fmt.Println(v2, v21)
                        }
                        _ = fi32
                    }
                }
                var v33 string = ("" + "hello")
                v22 = (v22 ^ (v22 & -1))
                _ = fi26
                _ = v33
            }
        }
    }
    var v34 int16 = v23
    v35 := v3[2147483647]
    var v36 uint64 = 650
    v35 += 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
