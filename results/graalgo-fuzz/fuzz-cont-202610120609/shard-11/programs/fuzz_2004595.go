package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 string }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 uint32; F1 string; F2 string }
type GenS4 struct { F0 uint16 }
func genFn0(p0 int, p1 string) int {
    if (65535 == 65535) {
        if false {
            var v6 int = (p0 >> ((899) & 63))
            var v7 uint32 = 1
            v8 := []int{(p0 >> ((426) & 63)), (p0 ^ v6), p0, (p0 * (v6 * 0))}
            _ = v6
            _ = v7
            _ = v8
        }
        for fi9 := 0; fi9 < 14; fi9++ {
            for fi10 := 0; fi10 < 5; fi10++ {
                p1 = (("" + (p1 + (p1 + p1))) + (p1 + "\n\t"))
                if false {
                    switch fi10 {
                    case 128:
                        switch (("a" + p1) + (p1 + "\n\t")) {
                        case "a":
                            v11 := GenS1{F0: -32768, F1: 42, F2: p1}
                            _ = v11
                        default:
                            var v12 int64 = -9223372036854775808
                            _ = v12
                        }
                    case 32768:
                        var v13 uint32 = 692
                        _ = v13
                    case -747:
                        var v14 float64 = 1.5
                        _ = v14
                    }
                    if (0.0 >= 1e10) {
                        switch p1 {
                        case "a":
                            v15 := GenS1{F0: -1, F1: 18446744073709551615, F2: "foo bar"}
                            _ = v15
                        case "\n\t":
                            v16 := []int{(p0 | p0), (p0 << ((2) & 63))}
                            _ = v16
                        }
                        v17 := GenS1{F0: 42, F1: 18446744073709551615, F2: "\n\t"}
                        _ = v17
                    }
                    v18 := GenS2{F0: 2, F1: 2}
                    _ = v18
                } else {
                    switch (p1 + "a") {
                    case "hello":
                        fmt.Println(fi10, fi9)
                    }
                    v19 := []int{p0, fi10, fi10}
                    v20 := GenS1{F0: -269, F1: 804, F2: p1}
                    _ = v19
                    _ = v20
                }
                _ = fi10
            }
            var v21 uint16 = 924
            _ = fi9
            _ = v21
        }
    }
    var v22 uint64 = 4294967295
    for fi23 := 0; fi23 < 6; fi23++ {
        var v24 float64 = float64((float32(1e10) * float32(1.5)))
        fmt.Println(v24, fi23, fi23)
        _ = fi23
        _ = v24
    }
    v25 := []int{(p0 | p0), 32767}
    _ = p0
    _ = p1
    _ = v22
    _ = v25
    return p0
}

func main() {
    var v1 int = 9223372036854775807
    var v2 int32 = -2147483648
    var v3 bool = (!(!false))
    var v4 int = (-(v1))
    var v5 float32 = 0.0
    v4 -= 0
    var v26 uint64 = 1
    v27 := []int{(-(v1))}
    v26 = 4294967295
    v27 = append(v27, (v4 * -9223372036854775808))
    _ = len(v27)
    var v28 int = 2147483647
    v1 -= (^(v1))
    v29 := GenS1{F0: 42, F1: 18446744073709551615, F2: "hello"}
    if v3 {
        v3 = false
    }
    v30 := GenS3{F0: 65535, F1: ("\n\t" + "hello"), F2: (("hello" + "\n\t") + "\n\t")}
    fmt.Println(v28, v26, v3)
    var v31 uint8 = 255
    var v32 float64 = (3.14159 - (0.1 * 1e10))
    for fi33 := 0; fi33 < 17; fi33++ {
        switch "\n\t" {
        case "\n\t":
            switch genFn0(int(v26), "\n\t") {
            case 128:
                for fi34 := 0; fi34 < 20; fi34++ {
                    fmt.Println(v31, v26)
                    v35 := GenS4{F0: 65535}
                    v28 -= 1
                    _ = fi34
                    _ = v35
                }
            }
        case "a":
            fmt.Println(v5, fi33, v1)
        default:
            v3 = true
        }
        _ = fi33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
