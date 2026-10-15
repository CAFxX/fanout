package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 bool; F1 uint; F2 uint64 }
func genFn0(p0 string) int64 {
    v5 := []int{42, 9223372036854775807}
    switch 0 {
    case 989:
        p0 = (p0 + p0)
    case -361:
        v6 := GenS1{F0: -2147483648, F1: false}
        _ = v6
    default:
        var v7 uint8 = 0
        _ = v7
    }
    _ = p0
    _ = v5
    return 9223372036854775807
}

func main() {
    v1 := []string{"a"}
    var v2 float64 = (0.1 * 1.5)
    var v3 float32 = 1e10
    var v4 uint64 = 199
    var v8 float32 = v3
    var v9 bool = ((!true) || true)
    var v10 uint = 0
    fmt.Println(v3, v2, v9)
    if (1000 != 774) {
        switch "hello" {
        case "":
            for fi11 := 0; fi11 < 16; fi11++ {
                var v12 bool = v9
                v2 = (v2 / 1.5)
                switch "\n\t" {
                case "foo bar":
                    v3 = 1.5
                }
                _ = fi11
                _ = v12
            }
        default:
            fmt.Println(v9, v4, v2)
        }
        var v13 int32 = 42
        _ = v13
    }
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    for fi14 := 0; fi14 < 11; fi14++ {
        if v9 {
            var v15 uint16 = 42
            for fi16 := 0; fi16 < 1; fi16++ {
                switch fi14 {
                case -9223372036854775808:
                    var v17 uint16 = v15
                    _ = v17
                case 4294967295:
                    var v18 uint32 = 2
                    _ = v18
                case 128:
                    switch ("" + ("foo bar" + ("a" + "a"))) {
                    case "a":
                        v10 = v10
                    }
                    v4 -= v4
                    _ = len(v1)
                }
                _ = fi16
            }
            var v19 uint8 = 0
            _ = v15
            _ = v19
        }
        v10 = v10
        v1 = append(v1, "\n\t")
        _ = fi14
    }
    fmt.Println(v3)
    switch -745 {
    case -9223372036854775808:
        for fi20 := 0; fi20 < 18; fi20++ {
            switch ("foo bar" + "a") {
            case "foo bar":
                v2 = v2
            case "hello":
                if v9 {
                    v21 := GenS2{F0: 1, F1: -568}
                    _ = v21
                } else {
                    var v22 int64 = -586
                    _ = v22
                }
            }
            _ = fi20
        }
    case 788:
        v8 += 1.5
    case 42:
        if (!v9) {
            for fi23 := 0; fi23 < 5; fi23++ {
                if len(v1) > 0 {
                    v1[0] = "foo bar"
                }
                for fi24 := 0; fi24 < 18; fi24++ {
                    var v25 int64 = 0
                    v9 = (-1 <= genFn0("\n\t"))
                    v4 = (v4 - (v4 << ((4294967295) & 63)))
                    _ = fi24
                    _ = v25
                }
                _ = fi23
            }
            var v26 string = "a"
            _ = v26
        } else {
            if true {
                v27 := GenS2{F0: 1, F1: -588}
                v10 = (v10 << ((v10) & 63))
                _ = v27
            }
            var v28 uint16 = 255
            _ = v28
        }
    default:
        v29 := GenS1{F0: 42, F1: (0 > 20)}
        _ = v29
    }
    for fi30 := 0; fi30 < 17; fi30++ {
        v4 -= (v4 % (((v4 % ((v4) | 1))) | 1))
        var v31 float64 = v2
        v3 -= 0.1
        _ = fi30
        _ = v31
    }
    v32 := GenS3{F0: false, F1: v10, F2: 0}
    _ = len(v1)
    v1 = append(v1, "a")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v8
    _ = v9
}
