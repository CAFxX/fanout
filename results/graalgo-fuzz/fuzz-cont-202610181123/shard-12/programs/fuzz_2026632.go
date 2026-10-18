package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 uint8; F1 uint }
func genFn0(p0 uint8, p1 int16, p2 uint8) uint8 {
    v7 := []int{}
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 uint64 = 133
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p0
}
func genFn1() int8 {
    var v10 float64 = (float64((-2.25 / 3.14159)) / 1.0)
    switch -9223372036854775808 {
    case 9223372036854775807:
        for fi11 := 0; fi11 < 16; fi11++ {
            v12 := GenS1{F0: "foo bar", F1: 1, F2: 18446744073709551615}
            if false {
                v10 = float64(-2.25)
                var v13 float64 = 0.0
                _ = v13
            } else {
                var v14 int64 = 42
                fmt.Println(v14, v14, v10)
                _ = v14
            }
            fmt.Println(fi11, v10)
            _ = fi11
            _ = v12
        }
    }
    _ = v10
    return 127
}

func main() {
    var v1 float32 = 0.0
    v2 := make(map[string]int)
    var v3 uint = 2
    var v4 float32 = 0.0
    var v5 bool = true
    v6 := make(map[string]string)
    for fi15 := 0; fi15 < 16; fi15++ {
        v16 := []int{fi15}
        v1 = -2.25
        v1 = (3.14159 / (float32((((1.5 + 1.5) + (1e10 + 1e10)) + (0.1 - float64(v1)))) / 1.0))
        _ = fi15
        _ = v16
    }
    v6["\n\t"] = ""
    for fi17 := 0; fi17 < 20; fi17++ {
        if v5 {
            v1 += 0.1
            switch fi17 {
            case 4294967295:
                var v18 uint8 = 2
                _ = v18
            default:
                switch 127 {
                case 0:
                    fmt.Println(v3)
                case 2147483647:
                    v3 = 255
                case 128:
                    v19 := GenS1{F0: (("\n\t" + "a") + "foo bar"), F1: 32767, F2: 1}
                    _ = v19
                default:
                    v20 := v2["foo bar"]
                    _ = v20
                }
            }
            fmt.Println(v4, v1)
        }
        _ = fi17
    }
    if (4294967295 == 949) {
        v21 := GenS2{F0: 1, F1: -123}
        for fi22 := 0; fi22 < 16; fi22++ {
            _ = len(v6)
            v23 := GenS3{F0: 127, F1: (v3 + v3)}
            v3 = 953
            _ = fi22
            _ = v23
        }
        v3 += (v3 ^ v3)
        _ = v21
    } else {
        var v24 bool = true
        var v25 int64 = -869
        _ = v24
        _ = v25
    }
    var v26 uint64 = 555
    var v27 uint32 = 2147483647
    switch -369 {
    case 128:
        v28 := []int{32767, -1, 9223372036854775807}
        _ = v28
    }
    switch 32767 {
    case -1:
        for fi29 := 0; fi29 < 12; fi29++ {
            v27 -= v27
            _ = fi29
        }
    default:
        if (!false) {
            var v30 int16 = 111
            var v31 int16 = v30
            var v32 uint8 = 127
            _ = v30
            _ = v31
            _ = v32
        }
    }
    v5 = true
    switch (((("\n\t" + "a") + "hello") + (("" + "\n\t") + "a")) + ("" + "foo bar")) {
    case "":
        var v33 bool = (!((!(245 <= 42)) || true))
        _ = v33
    case "\n\t":
        v5 = (!false)
    }
    for fi34 := 0; fi34 < 2; fi34++ {
        v5 = v5
        _ = fi34
    }
    var v35 string = "hello"
    v36 := GenS3{F0: 141, F1: uint(v26)}
    v37 := []int{}
    v37 = append(v37, 42)
    fmt.Println(v3)
    var v38 int16 = 1000
    fmt.Println(v26, v5, v38)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
