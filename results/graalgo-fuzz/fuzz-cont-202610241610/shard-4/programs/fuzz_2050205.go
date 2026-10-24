package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 int16 }
type GenS2 struct { F0 int32; F1 uint8; F2 int64 }
func genFn0(p0 int64, p1 bool) uint64 {
    for fi4 := 0; fi4 < 7; fi4++ {
        switch "\n\t" {
        case "":
            p0 = 1
        default:
            var v5 uint64 = 286
            _ = v5
        }
        v6 := []int{fi4, 354}
        _ = fi4
        _ = v6
    }
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := GenS1{F0: 9223372036854775807, F1: (fi7 - -505), F2: 1000}
        p1 = p1
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 5; fi9++ {
        if (fi9 < 73) {
            for fi10 := 0; fi10 < 20; fi10++ {
                p1 = (((220 >= 2147483648) || p1) && false)
                switch (("a" + (("a" + "foo bar") + (("foo bar" + "") + "hello"))) + (("a" + "foo bar") + "")) {
                case "foo bar":
                    var v11 string = ("\n\t" + "foo bar")
                    _ = v11
                case "a":
                    v12 := []int{fi9, fi9}
                    _ = v12
                case "\n\t":
                    p0 += p0
                }
                fmt.Println(p0, fi9, p1)
                _ = fi10
            }
            var v13 uint16 = 122
            _ = v13
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    return 594
}
func genFn1() int8 {
    var v14 bool = true
    v15 := []int{-713, 4294967295, 127, 32768}
    v14 = (!(65535 != 1))
    var v16 int32 = -2147483648
    _ = v14
    _ = v15
    _ = v16
    return -128
}

func main() {
    var v1 float64 = 1e10
    var v2 int16 = -400
    var v3 uint = 280
    v2 -= 954
    var v17 uint32 = 1
    var v18 uint32 = 2
    v3 = v3
    if (-2.25 > ((float32(v1) / 0.1) + 0.0)) {
        v19 := []int{1, 2147483647, 32768}
        if len(v19) > 0 {
            v19[0] = 1
        }
        v1 = (v1 * (v1 - float64((0.1 - 0.0))))
        _ = v19
    } else {
        if (-71 == 64) {
            switch "" {
            case "a":
                var v20 int = int(v2)
                _ = v20
            }
            var v21 int16 = (v2 * -32768)
            v22 := GenS2{F0: -424, F1: 2, F2: 0}
            _ = v21
            _ = v22
        }
    }
    v1 = 1.5
    var v23 uint64 = 18446744073709551615
    if (!true) {
        for fi24 := 0; fi24 < 14; fi24++ {
            v23 = (v23 ^ v23)
            var v25 uint64 = (^(v23))
            var v26 uint64 = (v25 - (v23 % (((v23 + v25)) | 1)))
            _ = fi24
            _ = v25
            _ = v26
        }
    }
    if false {
        var v27 int16 = v2
        _ = v27
    }
    if true {
        v28 := []int{-765, 32767, 128, 695}
        v29 := GenS1{F0: 9223372036854775807, F1: 128, F2: (^(v2))}
        _ = v28
        _ = v29
    }
    v30 := []int{-811, 827, 1}
    _ = len(v30)
    v31 := GenS1{F0: -902, F1: -961, F2: (v2 / (((v2 % ((-224) | 1))) | 1))}
    for fi32 := 0; fi32 < 13; fi32++ {
        if false {
            v23 += v23
            v17 = v18
        } else {
            v33 := GenS1{F0: 0, F1: (fi32 | (fi32 << (((v3 ^ (v3 >> ((4294967295) & 63)))) & 63))), F2: 600}
            _ = v33
        }
        v18 = 825
        _ = fi32
    }
    v1 = float64((1e10 / 1.5))
    var v34 int32 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v30
    _ = v31
    _ = v34
}
