package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 int64; F2 uint32 }
type GenS3 struct { F0 int8 }

func main() {
    var v1 int32 = 1
    v2 := []string{("\n\t" + ("\n\t" + "hello")), "", "\n\t"}
    var v3 float64 = 1.5
    var v4 bool = true
    var v5 bool = (0 >= 1000)
    v6 := GenS1{F0: 20}
    switch 127 {
    case 128:
        fmt.Println(v4)
    case 42:
        if len(v2) > 0 {
            v2[0] = ("" + (("hello" + (("\n\t" + "a") + ("foo bar" + "foo bar"))) + "a"))
        }
    default:
        v4 = v4
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi7 := 0; fi7 < 4; fi7++ {
        switch fi7 {
        case 399:
            if true {
                v8 := GenS2{F0: (v1 / ((-1) | 1)), F1: 42, F2: 4294967295}
                v9 := GenS2{F0: 188, F1: -9223372036854775808, F2: 1}
                for fi10 := 0; fi10 < 7; fi10++ {
                    v1 += (v1 & -1)
                    v1 = -377
                    _ = fi10
                }
                _ = v8
                _ = v9
            } else {
                v2 = append(v2, "")
            }
        case 4294967295:
            if v5 {
                v11 := GenS2{F0: 2147483647, F1: 9223372036854775807, F2: 358}
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = v11
            } else {
                switch fi7 {
                case -1:
                    var v12 bool = false
                    _ = v12
                default:
                    var v13 int16 = 1
                    _ = v13
                }
                v2 = append(v2, "")
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
        default:
            switch fi7 {
            case 32768:
                v3 -= float64(float32(1e10))
            case 1:
                v3 = (v3 - (v3 * v3))
            }
        }
        v1 += (-(v1))
        _ = fi7
    }
    var v14 int16 = 586
    var v15 float64 = v3
    v2 = append(v2, "foo bar")
    v15 += 1.5
    var v16 uint32 = 65535
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    v17 := GenS3{F0: 100}
    v5 = (false || (-694 != 42))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
