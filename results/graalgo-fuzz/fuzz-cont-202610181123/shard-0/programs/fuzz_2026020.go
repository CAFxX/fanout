package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 int8 }
func genFn0(p0 uint8, p1 int, p2 int16) int32 {
    fmt.Println(p1)
    v6 := []int{}
    if (2147483647 > -2147483648) {
        for fi7 := 0; fi7 < 19; fi7++ {
            p1 -= (p1 & (-(fi7)))
            fmt.Println(p0, p2, fi7)
            _ = fi7
        }
        fmt.Println(p1, p0)
    }
    p2 += (p2 / (((p2 * p2)) | 1))
    if ((1 >= 65535) || (!true)) {
        if (false && false) {
            for fi8 := 0; fi8 < 15; fi8++ {
                v9 := GenS1{F0: p2, F1: 2}
                fmt.Println(p1, fi8, p2)
                _ = len(v6)
                _ = fi8
                _ = v9
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 2147483647
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]bool)
    var v3 uint8 = 178
    var v4 bool = false
    var v5 string = "foo bar"
    var v10 float64 = (((0.0 / 1e10) * (0.0 / (0.1 + 0.0))) - 0.1)
    if v4 {
        for fi11 := 0; fi11 < 3; fi11++ {
            var v12 uint64 = 989
            _ = fi11
            _ = v12
        }
        v13 := []int{4294967295, 9223372036854775807, 32768, 42}
        _ = v13
    }
    fmt.Println(v10, v4)
    delete(v1, (((v5 + "foo bar") + "a") + v5))
    fmt.Println(v5)
    v4 = v4
    v14 := GenS2{F0: v4}
    for fi15 := 0; fi15 < 6; fi15++ {
        v5 = ""
        switch fi15 {
        case 929:
            var v16 int8 = -43
            _ = v16
        case 222:
            v5 = (v5 + "a")
        case -1:
            for fi17 := 0; fi17 < 8; fi17++ {
                v1["hello"] = (v5 + v5)
                _ = fi17
            }
        default:
            if ((100 > 42) || (916 <= 4294967295)) {
                switch fi15 {
                case 4294967295:
                    for fi18 := 0; fi18 < 13; fi18++ {
                        v10 += float64(3.14159)
                        var v19 uint = 52
                        switch ("" + "a") {
                        case "hello":
                            var v20 uint = v19
                            _ = v20
                        default:
                            v21 := GenS2{F0: (!true)}
                            _ = v21
                        }
                        fmt.Println(v5, v3)
                        var v22 uint16 = 32767
                        v3 -= 128
                        _ = fi18
                        _ = v19
                        _ = v22
                    }
                case 0:
                    v10 = (0.0 + (v10 + float64(0.0)))
                case 198:
                    v4 = (!true)
                }
            }
        }
        _ = fi15
    }
    var v23 uint16 = 255
    var v24 int32 = 2147483647
    switch v5 {
    case "\n\t":
        var v25 int64 = 889
        _ = v25
    case "hello":
        var v26 uint64 = 0
        _ = v26
    default:
        v24 = v24
    }
    fmt.Println(v10, v10, v23)
    for fi27 := 0; fi27 < 19; fi27++ {
        v28 := GenS2{F0: true}
        v29 := GenS3{F0: -885, F1: v4}
        delete(v1, v5)
        _ = fi27
        _ = v28
        _ = v29
    }
    for fi30 := 0; fi30 < 11; fi30++ {
        v4 = true
        v5 = (((((v5 + ("" + "")) + ("" + (v5 + v5))) + v5) + v5) + ((((v5 + v5) + v5) + v5) + "hello"))
        var v31 int8 = 127
        _ = fi30
        _ = v31
    }
    var v32 float32 = (1.5 / 1.5)
    switch 9223372036854775807 {
    case 32767:
        delete(v1, "hello")
    }
    for fi33 := 0; fi33 < 20; fi33++ {
        for fi34 := 0; fi34 < 1; fi34++ {
            v5 = "a"
            v35 := GenS1{F0: -32768, F1: 360}
            v36 := GenS4{F0: 42}
            _ = fi34
            _ = v35
            _ = v36
        }
        _ = fi33
    }
    v37 := GenS2{F0: (v4 && false)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v32
    _ = v37
    _ = v4
    _ = v5
}
