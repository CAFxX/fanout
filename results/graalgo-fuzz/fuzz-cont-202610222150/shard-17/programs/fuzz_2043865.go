package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 int; F1 uint32 }
type GenS5 struct { F0 int32; F1 uint32; F2 uint8 }
type GenS6 struct { F0 uint8; F1 string }
func genFn0(p0 uint32) uint8 {
    var v4 int = 371
    p0 = p0
    var v5 int64 = 251
    if false {
        fmt.Println(v5)
        fmt.Println(v5, v5)
        fmt.Println(v5)
    } else {
        v4 -= v4
        switch (v4 / ((v4) | 1)) {
        case 127:
            v6 := GenS1{F0: ("a" + ("" + ("" + ("foo bar" + "\n\t")))), F1: (v4 >> ((4294967295) & 63))}
            _ = v6
        case 1:
            for fi7 := 0; fi7 < 12; fi7++ {
                var v8 string = "foo bar"
                for fi9 := 0; fi9 < 19; fi9++ {
                    p0 = 119
                    v5 = 1
                    _ = fi9
                }
                _ = fi7
                _ = v8
            }
        case 72:
            switch "hello" {
            case "a":
                v10 := GenS2{F0: -1, F1: "foo bar"}
                _ = v10
            default:
                if (((0.0 - 1.5) / 1e10) <= float32(3.14159)) {
                    for fi11 := 0; fi11 < 17; fi11++ {
                        v12 := []int{(v4 & 127), (v4 & (v4 << ((447) & 63))), fi11, v4}
                        _ = fi11
                        _ = v12
                    }
                    var v13 uint = 4294967295
                    v14 := []int{91, int(v13)}
                    _ = v13
                    _ = v14
                }
            }
        default:
            fmt.Println(p0, v4, v5)
        }
        p0 = 65535
    }
    v15 := GenS3{F0: 118, F1: (v4 & v4)}
    _ = p0
    _ = v15
    _ = v4
    _ = v5
    return 157
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 uint16 = 32767
    var v3 int = 392
    v2 -= v2
    var v16 uint16 = v2
    v17 := GenS2{F0: 82, F1: ("hello" + ("" + "\n\t"))}
    v18 := []int{v3, (v3 * -534), v3, (v3 - -271)}
    var v19 float64 = 0.1
    v20 := GenS4{F0: -504, F1: 65535}
    switch v3 {
    case 2147483647:
        v18 = append(v18, -9223372036854775808)
    case 4294967295:
        v21 := GenS4{F0: int(v2), F1: 726}
        _ = v21
    case -9223372036854775808:
        for fi22 := 0; fi22 < 20; fi22++ {
            v23 := GenS5{F0: -1, F1: 826, F2: 127}
            v24 := GenS4{F0: (v3 >> ((2) & 63)), F1: 1}
            _ = fi22
            _ = v23
            _ = v24
        }
    default:
        fmt.Println(v2)
    }
    var v25 uint8 = 1
    var v26 int = v3
    switch ((("foo bar" + "\n\t") + (("a" + "a") + ((("hello" + "a") + ("\n\t" + "foo bar")) + "a"))) + ("" + "")) {
    case "\n\t":
        if true {
            var v27 int16 = -1
            var v28 int16 = 42
            _ = v27
            _ = v28
        } else {
            fmt.Println(v25, v2, v16)
            v29 := GenS6{F0: v25, F1: "a"}
            _ = v29
        }
    case "a":
        v30 := GenS3{F0: 2, F1: v26}
        _ = v30
    }
    v25 -= genFn0(1)
    var v31 int = 128
    fmt.Println(v31)
    v2 += v16
    fmt.Println(v25)
    fmt.Println(v16, v3, v16)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v3
    _ = v31
}
