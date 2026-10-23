package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 uint64; F1 int64; F2 int32 }
type GenS3 struct { F0 int64; F1 int }
func genFn0(p0 int16) int8 {
    for fi5 := 0; fi5 < 12; fi5++ {
        p0 -= -464
        _ = fi5
    }
    if false {
        fmt.Println(p0)
        v6 := GenS1{F0: 42, F1: 127}
        var v7 int64 = -9223372036854775808
        _ = v6
        _ = v7
    } else {
        v8 := GenS2{F0: 1, F1: 0, F2: -57}
        switch 0 {
        case 647:
            var v9 int8 = 105
            _ = v9
        case 32767:
            var v10 uint16 = 343
            _ = v10
        default:
            if (!false) {
                var v11 float32 = (1.5 / 1.0)
                v11 = v11
                for fi12 := 0; fi12 < 16; fi12++ {
                    v13 := GenS1{F0: -1, F1: (fi12 | -1)}
                    v11 -= (((v11 - v11) / (v11 * 3.14159)) - 0.0)
                    _ = fi12
                    _ = v13
                }
                _ = v11
            }
        }
        _ = v8
    }
    p0 -= -934
    fmt.Println(p0)
    var v14 uint16 = 255
    _ = p0
    _ = v14
    return 0
}

func main() {
    var v1 float32 = 1.5
    var v2 bool = (255 < 42)
    var v3 int64 = 42
    var v4 float32 = 0.1
    var v15 int8 = genFn0(1000)
    var v16 int16 = -32768
    switch "\n\t" {
    case "a":
        var v17 int = 128
        _ = v17
    }
    for fi18 := 0; fi18 < 12; fi18++ {
        v2 = (734 < 42)
        v1 = v1
        v2 = false
        _ = fi18
    }
    fmt.Println(v1, v1, v2)
    if (42 <= 18446744073709551615) {
        if (v2 || v2) {
            for fi19 := 0; fi19 < 2; fi19++ {
                v20 := []int{(fi19 >> ((42) & 63)), 128}
                for fi21 := 0; fi21 < 10; fi21++ {
                    var v22 uint8 = 255
                    _ = fi21
                    _ = v22
                }
                var v23 int = fi19
                _ = fi19
                _ = v20
                _ = v23
            }
        } else {
            if (!(!v2)) {
                v16 -= (^(v16))
                switch "" {
                case "foo bar":
                    v16 -= (v16 & (v16 >> ((389) & 63)))
                default:
                    var v24 uint16 = 636
                    _ = v24
                }
                v2 = (376 < 32767)
            } else {
                switch ((("a" + "foo bar") + "a") + ("hello" + "\n\t")) {
                case "hello":
                    var v25 float64 = (float64(0.0) + -2.25)
                    _ = v25
                case "":
                    v16 += v16
                case "\n\t":
                    fmt.Println(v16, v15, v4)
                }
                v26 := []int{}
                _ = v26
            }
        }
    }
    switch "" {
    case "a":
        var v27 uint32 = 0
        _ = v27
    default:
        v28 := GenS2{F0: 42, F1: 1, F2: -1}
        _ = v28
    }
    switch 810 {
    case 42:
        for fi29 := 0; fi29 < 13; fi29++ {
            if (v2 && (!(("a" + "a") < "foo bar"))) {
                v30 := GenS1{F0: v15, F1: fi29}
                var v31 int16 = (v16 << ((18446744073709551615) & 63))
                _ = v30
                _ = v31
            }
            v32 := GenS3{F0: v3, F1: 32768}
            _ = fi29
            _ = v32
        }
    case 1:
        var v33 int16 = -383
        _ = v33
    default:
        v34 := []int{}
        _ = v34
    }
    v35 := GenS2{F0: 42, F1: (v3 - v3), F2: 1}
    var v36 uint32 = 4294967295
    var v37 uint = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
}
