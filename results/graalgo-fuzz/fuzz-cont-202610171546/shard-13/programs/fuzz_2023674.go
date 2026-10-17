package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint }
type GenS2 struct { F0 int64; F1 int32 }
func genFn0(p0 int16, p1 string) int8 {
    switch p1 {
    case "foo bar":
        if (251 >= 65535) {
            v6 := GenS1{F0: 2, F1: p1, F2: 2}
            fmt.Println(p0, p1)
            var v7 uint32 = 435
            _ = v6
            _ = v7
        }
    default:
        v8 := GenS1{F0: 42, F1: "hello", F2: 65535}
        _ = v8
    }
    var v9 uint8 = 128
    _ = p0
    _ = p1
    _ = v9
    return -1
}

func main() {
    var v1 int8 = 1
    v2 := make(map[uint32]string)
    var v3 uint64 = 18446744073709551615
    var v4 int = 32767
    v5 := make(map[int]int)
    var v10 uint8 = 255
    v11 := GenS1{F0: (^(v10)), F1: "foo bar", F2: uint(v10)}
    for fi12 := 0; fi12 < 2; fi12++ {
        var v13 uint64 = v3
        if (-1 <= -490) {
            v3 += 4294967295
            v1 = (v1 / ((127) | 1))
            v10 = (v10 >> ((18446744073709551615) & 63))
        } else {
            v14 := v5[fi12]
            v15 := GenS1{F0: v10, F1: "a", F2: 951}
            _ = v14
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    var v16 float64 = 0.0
    var v17 int = (v4 ^ (^(v4)))
    for fi18 := 0; fi18 < 8; fi18++ {
        if (false || (!true)) {
            var v19 int32 = 2147483647
            _ = v19
        }
        var v20 int64 = 365
        v4 += (fi18 - (v17 / (((fi18 | 4294967295)) | 1)))
        _ = fi18
        _ = v20
    }
    fmt.Println(v1)
    v21 := GenS2{F0: 0, F1: -204}
    var v22 uint8 = v10
    var v23 uint = 0
    var v24 int64 = 1
    for fi25 := 0; fi25 < 15; fi25++ {
        v26 := GenS2{F0: v24, F1: -2147483648}
        v27 := v2[2147483648]
        v16 += v16
        _ = fi25
        _ = v26
        _ = v27
    }
    switch (("foo bar" + "a") + ((("foo bar" + (("a" + "hello") + (("" + ("" + "hello")) + "hello"))) + (("\n\t" + (("\n\t" + ("a" + "\n\t")) + "a")) + (((("foo bar" + ("" + "a")) + "foo bar") + "\n\t") + "\n\t"))) + "foo bar")) {
    case "a":
        switch v4 {
        case 42:
            v28 := GenS2{F0: v24, F1: 720}
            _ = v28
        case -212:
            var v29 int = -2147483648
            _ = v29
        case 0:
            fmt.Println(v17)
        }
    case "":
        v10 = (v22 >> ((uint(v22)) & 63))
    case "hello":
        switch -419 {
        case -275:
            v22 -= v22
        }
    default:
        if (!false) {
            if (!true) {
                v30 := GenS1{F0: 2, F1: "a", F2: v23}
                _ = v30
            }
            delete(v5, (^(v4)))
        }
    }
    for fi31 := 0; fi31 < 16; fi31++ {
        v32 := []int{v4}
        switch ("\n\t" + "a") {
        case "\n\t":
            fmt.Println(v23)
        case "hello":
            if false {
                switch (v17 ^ (v4 >> ((uint(v3)) & 63))) {
                case 1:
                    v5[(-(fi31))] = (v17 >> ((4294967295) & 63))
                }
            } else {
                switch "foo bar" {
                case "hello":
                    if false {
                        v33 := GenS1{F0: (v22 * v10), F1: "hello", F2: v23}
                        v34 := GenS2{F0: v24, F1: -1}
                        _ = v33
                        _ = v34
                    } else {
                        var v35 string = "a"
                        v17 = (v4 | (v17 << ((v23) & 63)))
                        v1 = v1
                        _ = v35
                    }
                }
                v32 = append(v32, (-(fi31)))
                var v36 float64 = (1.5 / (float64(1.5) / v16))
                _ = v36
            }
        case "":
            for fi37 := 0; fi37 < 11; fi37++ {
                var v38 int16 = -32768
                _ = fi37
                _ = v38
            }
        }
        v23 -= v23
        _ = fi31
        _ = v32
    }
    v22 -= (^(v22))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
