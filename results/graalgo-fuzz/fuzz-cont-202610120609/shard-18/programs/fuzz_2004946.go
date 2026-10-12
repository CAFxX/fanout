package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint8 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 uint16; F1 int8; F2 uint16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int16; F1 string }
type GenS6 struct { F0 uint16; F1 int64; F2 int }
func genFn0(p0 int64) uint16 {
    if false {
        switch "foo bar" {
        case "":
            if true {
                var v7 int16 = -32768
                _ = v7
            } else {
                if false {
                    var v8 string = "hello"
                    p0 -= (p0 % ((p0) | 1))
                    _ = v8
                } else {
                    var v9 uint64 = 1
                    for fi10 := 0; fi10 < 14; fi10++ {
                        var v11 uint8 = 2
                        _ = fi10
                        _ = v11
                    }
                    switch -481 {
                    case 1:
                        fmt.Println(p0, v9)
                    case 445:
                        for fi12 := 0; fi12 < 5; fi12++ {
                            v9 += 1
                            _ = fi12
                        }
                    }
                    _ = v9
                }
            }
        default:
            switch 0 {
            case 9223372036854775807:
                v13 := []int{-572, 896, 42}
                _ = v13
            case 128:
                v14 := []int{-1, 0, 32767}
                _ = v14
            case 2147483647:
                for fi15 := 0; fi15 < 13; fi15++ {
                    v16 := GenS1{F0: "a", F1: 507, F2: 128}
                    _ = fi15
                    _ = v16
                }
            default:
                v17 := GenS1{F0: "", F1: 2, F2: 128}
                _ = v17
            }
        }
        v18 := []int{}
        if true {
            v19 := GenS1{F0: "\n\t", F1: 255, F2: 128}
            _ = v19
        } else {
            v20 := GenS2{F0: -16, F1: 2147483647}
            _ = v20
        }
        _ = v18
    }
    if (-632 >= -1) {
        v21 := GenS3{F0: 1, F1: -121, F2: 32767}
        _ = v21
    } else {
        var v22 uint = 2
        v22 = v22
        var v23 uint32 = 65535
        _ = v22
        _ = v23
    }
    switch "a" {
    case "\n\t":
        p0 -= p0
    }
    _ = p0
    return 0
}

func main() {
    var v1 float32 = 1e10
    var v2 bool = (true && ((168 == 255) || (1 != 1)))
    v3 := []string{"\n\t"}
    var v4 int16 = 116
    var v5 int32 = -1
    var v6 int64 = -182
    var v24 float64 = (((3.14159 - -2.25) - (1e10 - 0.1)) + ((-2.25 + (((0.1 - -2.25) - float64(v1)) + float64((v1 * (1e10 / float32(-2.25)))))) / -2.25))
    if v2 {
        fmt.Println(v24, v5)
        var v25 uint64 = 42
        v26 := GenS4{F0: 65535}
        _ = v25
        _ = v26
    }
    v4 += v4
    v1 = ((-2.25 * v1) + v1)
    v2 = false
    v27 := GenS5{F0: (v4 / ((v4) | 1)), F1: ((("a" + "a") + ("a" + "foo bar")) + "")}
    var v28 int16 = (v4 * 42)
    var v29 string = (("\n\t" + (("foo bar" + "hello") + "foo bar")) + "hello")
    for fi30 := 0; fi30 < 11; fi30++ {
        switch -416 {
        case -729:
            if true {
                v31 := GenS6{F0: 255, F1: (v6 * v6), F2: 4294967295}
                _ = v31
            }
        case 241:
            v32 := GenS6{F0: 42, F1: (v6 * v6), F2: fi30}
            _ = v32
        case 256:
            var v33 float32 = 3.14159
            _ = v33
        }
        v24 -= v24
        _ = fi30
    }
    var v34 uint = 543
    var v35 uint32 = 2147483648
    if len(v3) > 0 {
        v3[0] = v29
    }
    _ = len(v3)
    v24 = (1.5 * v24)
    var v36 int = -1
    v28 = (^(v4))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
