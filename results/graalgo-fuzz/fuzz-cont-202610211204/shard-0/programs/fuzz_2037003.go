package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 uint }
type GenS2 struct { F0 bool; F1 int64 }
func genFn0(p0 int32, p1 uint32) bool {
    fmt.Println(p0, p0)
    if false {
        p0 += (^(p0))
        v5 := []int{32767}
        p0 = p0
        _ = v5
    }
    p1 = p1
    v6 := GenS1{F0: 0, F1: 2, F2: 0}
    var v7 uint8 = 2
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return true
}
func genFn1(p0 int64, p1 int32, p2 uint16) uint32 {
    var v8 int16 = -32768
    for fi9 := 0; fi9 < 19; fi9++ {
        for fi10 := 0; fi10 < 10; fi10++ {
            var v11 int64 = p0
            var v12 int16 = v8
            if (v11 < v11) {
                var v13 int32 = p1
                var v14 uint = 1
                _ = v13
                _ = v14
            }
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi9
    }
    switch ("foo bar" + ("a" + "hello")) {
    case "\n\t":
        switch "a" {
        case "hello":
            v8 += -841
        case "\n\t":
            for fi15 := 0; fi15 < 13; fi15++ {
                var v16 int = (fi15 << ((0) & 63))
                for fi17 := 0; fi17 < 11; fi17++ {
                    v18 := []int{fi15}
                    var v19 string = (("foo bar" + "\n\t") + "")
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                _ = fi15
                _ = v16
            }
        case "":
            var v20 float32 = 1e10
            _ = v20
        }
    case "foo bar":
        p0 += p0
    default:
        v8 -= 32767
    }
    v21 := []int{}
    if false {
        for fi22 := 0; fi22 < 11; fi22++ {
            if (!true) {
                switch (fi22 | (-(fi22))) {
                case -233:
                    for fi23 := 0; fi23 < 18; fi23++ {
                        v24 := GenS1{F0: v8, F1: 0, F2: 0}
                        fmt.Println(p1, p0, fi22)
                        v25 := GenS2{F0: (float64(1e10) >= (1.5 + 0.0)), F1: p0}
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                }
                p0 = p0
                fmt.Println(p1)
            } else {
                fmt.Println(p1)
                fmt.Println(p2, p0, p0)
            }
            p0 -= p0
            _ = fi22
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    _ = v8
    return 2147483648
}

func main() {
    var v1 float32 = (0.1 + 1.5)
    var v2 int = 2147483647
    var v3 string = "\n\t"
    var v4 uint8 = 111
    var v26 string = ""
    var v27 float64 = 0.0
    v1 -= v1
    var v28 float32 = v1
    if true {
        var v29 string = "\n\t"
        _ = v29
    } else {
        for fi30 := 0; fi30 < 3; fi30++ {
            var v31 int16 = 42
            v32 := []int{(v2 % ((v2) | 1))}
            v4 -= (v4 & (v4 >> ((255) & 63)))
            _ = fi30
            _ = v31
            _ = v32
        }
    }
    v33 := []int{v2, (v2 % (((v2 * v2)) | 1)), v2}
    var v34 float64 = v27
    v1 = 0.0
    switch (v2 >> ((65535) & 63)) {
    case -1:
        fmt.Println(v27)
    case 9223372036854775807:
        v2 -= v2
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v33
    _ = v34
    _ = v4
}
