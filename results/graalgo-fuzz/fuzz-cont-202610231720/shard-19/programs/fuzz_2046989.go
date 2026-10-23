package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 string }

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = 2147483647
    var v3 int32 = (v2 | (v2 - 193))
    var v4 int64 = -150
    var v5 string = "\n\t"
    var v6 uint32 = 42
    v4 = (v4 | 9223372036854775807)
    if false {
        for fi7 := 0; fi7 < 1; fi7++ {
            v6 = v6
            v8 := []int{fi7, 362, 924}
            v4 -= v4
            _ = fi7
            _ = v8
        }
    } else {
        var v9 uint16 = 963
        if true {
            switch -9223372036854775808 {
            case 32767:
                var v10 int16 = 32767
                _ = v10
            case 32768:
                if ((-(v4)) == v4) {
                    v5 = "\n\t"
                    var v11 float64 = -2.25
                    _ = v11
                } else {
                    fmt.Println(v2)
                    if true {
                        fmt.Println(v2, v5)
                    }
                    var v12 float64 = (1e10 + ((-2.25 - float64(0.0)) + 0.0))
                    _ = v12
                }
            case 8:
                var v13 int16 = -996
                _ = v13
            default:
                var v14 bool = false
                _ = v14
            }
            v1["\n\t"] = ((((v5 + v5) + "\n\t") + "hello") == "\n\t")
        } else {
            var v15 uint64 = 83
            if (!(true && true)) {
                if false {
                    v16 := GenS1{F0: (!(!false)), F1: 42}
                    v17 := []int{941, 865}
                    _ = v16
                    _ = v17
                }
                var v18 int = 998
                _ = v18
            } else {
                if (1.5 > (0.0 / 1.0)) {
                    if true {
                        _ = len(v1)
                    }
                    v3 += 42
                }
            }
            v19 := v1["foo bar"]
            _ = v15
            _ = v19
        }
        delete(v1, (v5 + "hello"))
        _ = v9
    }
    switch v5 {
    case "\n\t":
        v20 := GenS2{F0: v5}
        _ = v20
    case "foo bar":
        v2 = (v3 | 1)
    case "a":
        switch 127 {
        case 128:
            v2 += v3
        case 2147483647:
            var v21 uint8 = 17
            _ = v21
        }
    default:
        for fi22 := 0; fi22 < 5; fi22++ {
            var v23 string = v5
            var v24 int32 = v2
            _ = fi22
            _ = v23
            _ = v24
        }
    }
    v6 = (v6 & 1)
    var v25 int8 = -83
    var v26 uint16 = 428
    var v27 int16 = 42
    fmt.Println(v5, v26)
    fmt.Println(v6, v27)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
