package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 int16 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 int8 }
func genFn0(p0 int16, p1 int16, p2 string) int64 {
    v7 := GenS1{F0: -68, F1: -860, F2: (p1 ^ (p0 / ((p0) | 1)))}
    fmt.Println(p2, p0, p2)
    p1 += (p0 ^ (p0 << ((2) & 63)))
    var v8 float32 = 1e10
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 int16) uint32 {
    if false {
        var v9 int8 = -1
        fmt.Println(p0, v9)
        v9 = (v9 * (v9 + (v9 / ((v9) | 1))))
        _ = v9
    } else {
        v10 := GenS2{F0: 2, F1: 42}
        for fi11 := 0; fi11 < 16; fi11++ {
            switch (fi11 / (((fi11 * fi11)) | 1)) {
            case -366:
                for fi12 := 0; fi12 < 16; fi12++ {
                    var v13 float64 = (0.0 * (-2.25 - float64(0.1)))
                    v14 := []int{}
                    v14 = append(v14, (^(fi12)))
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            case 15:
                fmt.Println(p0)
            default:
                v15 := []int{(^(fi11)), fi11}
                _ = v15
            }
            v16 := GenS2{F0: 238, F1: 947}
            _ = fi11
            _ = v16
        }
        _ = v10
    }
    if ("\n\t" < "\n\t") {
        for fi17 := 0; fi17 < 6; fi17++ {
            var v18 uint = 4294967295
            var v19 float64 = (float64(float32(float64(0.0))) - 1e10)
            _ = fi17
            _ = v18
            _ = v19
        }
        var v20 int64 = genFn0((p0 * (p0 | p0)), (p0 - (p0 / ((p0) | 1))), "")
        switch "foo bar" {
        case "foo bar":
            v21 := GenS3{F0: -128, F1: false}
            _ = v21
        case "a":
            fmt.Println(p0, p0)
        case "\n\t":
            if (false || false) {
                if (!(1 != -2147483648)) {
                    if (2 > 32767) {
                        v20 = genFn0(p0, (-(p0)), "\n\t")
                        v22 := GenS2{F0: 255, F1: 65535}
                        _ = v22
                    }
                }
                fmt.Println(v20)
            }
        }
        _ = v20
    } else {
        var v23 int8 = 0
        v24 := GenS3{F0: v23, F1: (!false)}
        var v25 bool = false
        _ = v23
        _ = v24
        _ = v25
    }
    var v26 uint8 = 128
    var v27 int16 = p0
    var v28 int32 = 1
    _ = p0
    _ = v26
    _ = v27
    _ = v28
    return 2
}

func main() {
    var v1 int64 = -1
    var v2 uint8 = 255
    var v3 uint8 = v2
    v4 := []string{("hello" + "hello"), (("\n\t" + "hello") + "foo bar")}
    var v5 string = "foo bar"
    var v6 int = 32767
    v29 := GenS4{F0: 48}
    var v30 int = v6
    var v31 float32 = 0.0
    for fi32 := 0; fi32 < 16; fi32++ {
        var v33 uint64 = 914
        _ = fi32
        _ = v33
    }
    var v34 int = -1
    v30 += v30
    fmt.Println(v5)
    v4 = append(v4, ((v5 + (v5 + v5)) + (((v5 + ((v5 + "") + "foo bar")) + "\n\t") + (((v5 + v5) + (v5 + (("a" + (v5 + "\n\t")) + "a"))) + v5))))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
