package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 int }
type GenS2 struct { F0 uint32; F1 int16 }
func genFn0(p0 uint32) uint {
    p0 -= (p0 % ((p0) | 1))
    p0 = p0
    v4 := []int{}
    switch 4294967295 {
    case 1:
        var v5 int16 = 246
        _ = v5
    case 4294967295:
        _ = len(v4)
    }
    _ = p0
    _ = v4
    return 255
}
func genFn1() string {
    switch "" {
    case "a":
        var v6 int16 = 1000
        _ = v6
    case "":
        var v7 string = "foo bar"
        _ = v7
    case "hello":
        v8 := GenS1{F0: 549, F1: 32767, F2: 4294967295}
        _ = v8
    default:
        v9 := GenS1{F0: 1, F1: 578, F2: 9223372036854775807}
        _ = v9
    }
    var v10 uint64 = 4294967295
    if (-2.25 >= float32(1e10)) {
        if false {
            var v11 bool = ((false || (-1 <= -1)) && false)
            _ = v11
        } else {
            if false {
                v10 += 42
            } else {
                switch ("hello" + (("\n\t" + "foo bar") + "\n\t")) {
                case "hello":
                    var v12 int16 = int16(v10)
                    _ = v12
                case "\n\t":
                    v13 := []int{}
                    _ = v13
                case "":
                    var v14 int16 = 0
                    _ = v14
                }
                v10 = 4294967295
            }
        }
    } else {
        for fi15 := 0; fi15 < 18; fi15++ {
            v10 = (v10 << ((255) & 63))
            var v16 int8 = 127
            _ = fi15
            _ = v16
        }
        var v17 int64 = -526
        for fi18 := 0; fi18 < 13; fi18++ {
            v10 = (v10 | 899)
            var v19 int = (^(fi18))
            v20 := []int{}
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = v17
    }
    v10 -= 1
    for fi21 := 0; fi21 < 7; fi21++ {
        v10 = 1
        for fi22 := 0; fi22 < 20; fi22++ {
            fmt.Println(v10, fi22)
            var v23 string = (("foo bar" + "") + "foo bar")
            var v24 uint8 = 32
            _ = fi22
            _ = v23
            _ = v24
        }
        _ = fi21
    }
    _ = v10
    return (((("\n\t" + "a") + "a") + ((("a" + "") + "foo bar") + "a")) + "a")
}

func main() {
    var v1 bool = true
    v2 := make(map[string]string)
    var v3 uint8 = 42
    if (((v3 == (v3 * 1)) || (!(0.1 <= (-2.25 + 3.14159)))) && true) {
        var v25 float64 = 1.5
        v25 = float64(0.0)
        _ = v25
    }
    v26 := []int{1, 128, 4294967295}
    if (true && (!v1)) {
        if v1 {
            if v1 {
                var v27 int32 = -233
                _ = v27
            }
            v3 += v3
            var v28 string = "foo bar"
            _ = v28
        } else {
            v3 += v3
        }
        for fi29 := 0; fi29 < 13; fi29++ {
            for fi30 := 0; fi30 < 16; fi30++ {
                var v31 bool = (((2147483647 <= 2147483648) || (v1 || (-123 < -1))) && false)
                v1 = ((!(-1 <= 2147483647)) && ((0 != -105) && (false || v31)))
                _ = fi30
                _ = v31
            }
            v32 := GenS2{F0: 2147483648, F1: -70}
            _ = fi29
            _ = v32
        }
    } else {
        if (v1 && (!(43 != -41))) {
            var v33 float64 = 3.14159
            _ = v33
        } else {
            switch ("" + "") {
            case "a":
                var v34 uint16 = 32767
                _ = v34
            case "\n\t":
                v3 -= 68
            }
            var v35 int32 = 2147483647
            _ = v35
        }
    }
    v3 += (v3 + (v3 & (v3 | v3)))
    for fi36 := 0; fi36 < 12; fi36++ {
        v3 -= v3
        _ = fi36
    }
    for fi37 := 0; fi37 < 10; fi37++ {
        v3 = v3
        _ = fi37
    }
    v1 = (127 >= -2147483648)
    for fi38 := 0; fi38 < 16; fi38++ {
        fmt.Println(v3)
        _ = fi38
    }
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
}
