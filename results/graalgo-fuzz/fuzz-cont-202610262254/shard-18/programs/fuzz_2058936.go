package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 int16; F1 int32; F2 string }
type GenS4 struct { F0 uint8; F1 uint16; F2 uint32 }
type GenS5 struct { F0 int8 }
func genFn0(p0 string) int32 {
    p0 = (("" + p0) + "\n\t")
    if true {
        var v5 int16 = 305
        _ = v5
    }
    _ = p0
    return 42
}
func genFn1(p0 bool) string {
    fmt.Println(p0)
    v6 := []int{}
    _ = p0
    _ = v6
    return ("hello" + ("hello" + "foo bar"))
}

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 4
    var v3 int32 = -837
    var v4 bool = ((-128 > -90) && (18446744073709551615 < 1))
    switch ("" + "a") {
    case "hello":
        var v7 uint32 = 630
        _ = v7
    case "a":
        v8 := GenS1{F0: 1000}
        _ = v8
    case "foo bar":
        v9 := []int{127, -2147483648, -394, 32768}
        _ = v9
    default:
        _ = len(v1)
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        for fi11 := 0; fi11 < 20; fi11++ {
            fmt.Println(fi10, v3)
            if ((42 > 857) || v4) {
                v12 := GenS1{F0: 0}
                v3 -= v3
                _ = v12
            }
            var v13 uint32 = 474
            _ = fi11
            _ = v13
        }
        _ = fi10
    }
    v14 := GenS2{F0: -32768, F1: 1}
    for fi15 := 0; fi15 < 5; fi15++ {
        v16 := GenS3{F0: 0, F1: v3, F2: ""}
        v17 := GenS4{F0: (v2 ^ (v2 << ((255) & 63))), F1: 32767, F2: 226}
        switch fi15 {
        case -481:
            var v18 float64 = -2.25
            _ = v18
        case 0:
            v4 = v4
        }
        _ = fi15
        _ = v16
        _ = v17
    }
    var v19 uint32 = 2147483647
    v20 := []int{}
    var v21 uint32 = (v19 - v19)
    if v4 {
        var v22 int16 = 32767
        _ = v22
    } else {
        switch "a" {
        case "":
            switch 2147483647 {
            case -9223372036854775808:
                var v23 string = (("" + "a") + "hello")
                _ = v23
            }
        case "hello":
            switch "foo bar" {
            case "hello":
                v24 := GenS3{F0: 324, F1: v3, F2: "a"}
                _ = v24
            default:
                fmt.Println(v4, v2, v3)
            }
        case "a":
            var v25 int64 = -9223372036854775808
            _ = v25
        }
        fmt.Println(v2, v3)
    }
    var v26 float32 = 0.0
    v2 += (v2 * v2)
    for fi27 := 0; fi27 < 19; fi27++ {
        switch "a" {
        case "\n\t":
            if (42 <= 4294967295) {
                v28 := GenS1{F0: -1}
                v26 = (v26 / (-2.25 + (v26 / float32(1e10))))
                if (true && false) {
                    v29 := GenS1{F0: -32768}
                    _ = v29
                } else {
                    v30 := GenS3{F0: 1, F1: (v3 / ((112) | 1)), F2: "a"}
                    _ = len(v1)
                    fmt.Println(v4)
                    _ = v30
                }
                _ = v28
            } else {
                var v31 bool = false
                _ = v31
            }
        case "":
            v32 := GenS5{F0: -29}
            _ = v32
        }
        var v33 uint16 = 65535
        if len(v20) >= 2 {
            v20 = v20[1:len(v20)]
        }
        _ = fi27
        _ = v33
    }
    v34 := v1[32767]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v3
    _ = v34
    _ = v4
}
