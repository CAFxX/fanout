package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 int64; F1 uint; F2 uint }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64; F1 uint16; F2 int64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint; F1 int }
type GenS8 struct { F0 int8 }
func genFn0() uint32 {
    v5 := GenS1{F0: ("foo bar" + ("\n\t" + "hello")), F1: 2}
    for fi6 := 0; fi6 < 8; fi6++ {
        if (((!false) || (!false)) && true) {
            var v7 uint32 = 0
            _ = v7
        }
        var v8 int = (fi6 * 32767)
        _ = fi6
        _ = v8
    }
    if (((("foo bar" + "a") + "a") + "a") < "") {
        var v9 int8 = 0
        if false {
            var v10 uint16 = 1
            v9 -= (v9 << ((1) & 63))
            if true {
                v11 := GenS1{F0: "foo bar", F1: 128}
                _ = v11
            } else {
                var v12 int16 = 32767
                v13 := GenS2{F0: 42, F1: 2, F2: 42}
                var v14 int16 = (v12 ^ (v12 / (((v12 % ((-32768) | 1))) | 1)))
                _ = v12
                _ = v13
                _ = v14
            }
            _ = v10
        }
        _ = v9
    }
    _ = v5
    return 243
}

func main() {
    v1 := make(map[string]int)
    var v2 int16 = 42
    var v3 uint = 0
    var v4 uint32 = 65535
    var v15 string = ("hello" + "a")
    var v16 int32 = 2147483647
    for fi17 := 0; fi17 < 6; fi17++ {
        if (9223372036854775807 >= 9223372036854775807) {
            v18 := GenS3{F0: 0}
            v19 := GenS2{F0: 729, F1: (^(v3)), F2: (v3 + 18446744073709551615)}
            _ = v18
            _ = v19
        } else {
            if true {
                v20 := GenS4{F0: 2, F1: 0, F2: 9223372036854775807}
                v3 -= v3
                var v21 float32 = ((1.5 - 0.0) * -2.25)
                _ = v20
                _ = v21
            }
        }
        _ = fi17
    }
    fmt.Println(v2, v4, v3)
    v22 := GenS5{F0: 127}
    v2 = v2
    var v23 float64 = 1.5
    if (v2 >= -32768) {
        v4 = 942
    } else {
        var v24 uint32 = v4
        if false {
            var v25 uint = (v3 ^ (v3 + v3))
            fmt.Println(v3)
            for fi26 := 0; fi26 < 18; fi26++ {
                v3 += 4294967295
                v3 = 4294967295
                _ = fi26
            }
            _ = v25
        } else {
            for fi27 := 0; fi27 < 19; fi27++ {
                v24 += genFn0()
                _ = fi27
            }
            for fi28 := 0; fi28 < 2; fi28++ {
                for fi29 := 0; fi29 < 2; fi29++ {
                    v15 = "hello"
                    _ = fi29
                }
                v30 := GenS6{F0: (v4 / (((v4 - (v24 & (v4 << ((v3) & 63))))) | 1))}
                v31 := GenS7{F0: (v3 / (((v3 & 172)) | 1)), F1: fi28}
                _ = fi28
                _ = v30
                _ = v31
            }
            for fi32 := 0; fi32 < 14; fi32++ {
                v33 := []int{(fi32 << ((v3) & 63)), -152}
                fmt.Println(v3, v24)
                _ = fi32
                _ = v33
            }
        }
        _ = v24
    }
    v4 += genFn0()
    v23 = v23
    for fi34 := 0; fi34 < 11; fi34++ {
        for fi35 := 0; fi35 < 2; fi35++ {
            if (2 == 42) {
                v15 = "\n\t"
                fmt.Println(v4)
            }
            v4 = (v4 / ((65535) | 1))
            v16 = (v16 - v16)
            v15 = (v15 + ("\n\t" + v15))
            _ = fi35
        }
        fmt.Println(v15, v2, v4)
        _ = fi34
    }
    v23 += (((v23 - float64(float32(0.0))) - (v23 + 0.1)) * 1e10)
    if (267 < uint16(v3)) {
        for fi36 := 0; fi36 < 4; fi36++ {
            switch ("" + ((v15 + "") + (((v15 + v15) + ((v15 + ("foo bar" + v15)) + "a")) + v15))) {
            case "foo bar":
                var v37 uint = (v3 - v3)
                _ = v37
            case "":
                var v38 int8 = 0
                _ = v38
            case "\n\t":
                for fi39 := 0; fi39 < 2; fi39++ {
                    v40 := []int{(fi39 + fi39), -9223372036854775808, 42}
                    var v41 int8 = 100
                    _ = fi39
                    _ = v40
                    _ = v41
                }
            default:
                v42 := GenS2{F0: 9223372036854775807, F1: (v3 ^ 4294967295), F2: v3}
                _ = v42
            }
            _ = fi36
        }
    } else {
        var v43 uint8 = 128
        _ = v43
    }
    v44 := []int{}
    v45 := GenS8{F0: 69}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v44:", len(v44), v44)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v44
    _ = v45
}
