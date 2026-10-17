package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 int8; F1 uint64; F2 uint8 }
type GenS4 struct { F0 string }
func genFn0() int32 {
    switch ("\n\t" + ((("\n\t" + "") + "a") + "hello")) {
    case "foo bar":
        switch ("hello" + "") {
        case "":
            if (!(("a" + (("hello" + "a") + "foo bar")) != "hello")) {
                var v4 float32 = 0.1
                for fi5 := 0; fi5 < 6; fi5++ {
                    v4 = v4
                    v6 := []int{(fi5 / (((^(fi5))) | 1)), (fi5 - fi5), (-(fi5))}
                    _ = fi5
                    _ = v6
                }
                _ = v4
            } else {
                var v7 float64 = (0.1 / 1.5)
                _ = v7
            }
        }
    default:
        var v8 int64 = 1
        _ = v8
    }
    var v9 float64 = 0.1
    _ = v9
    return -2147483648
}

func main() {
    var v1 uint64 = 665
    var v2 uint8 = 255
    var v3 bool = (42 == 2)
    v10 := GenS1{F0: 4294967295}
    if (42 != 0) {
        v11 := []int{-556, 1, 127}
        _ = v11
    }
    var v12 uint8 = (^(v2))
    v13 := []int{32767, -9223372036854775808, 9223372036854775807}
    v14 := GenS1{F0: 896}
    v15 := GenS1{F0: 799}
    var v16 float32 = 0.0
    if (false && (4294967295 != 42)) {
        fmt.Println(v1, v1)
        if true {
            fmt.Println(v2)
            for fi17 := 0; fi17 < 2; fi17++ {
                var v18 float32 = float32(0.1)
                _ = fi17
                _ = v18
            }
            switch "a" {
            case "\n\t":
                switch 1 {
                case 0:
                    v1 -= (v1 << ((997) & 63))
                default:
                    v3 = (!true)
                }
            case "foo bar":
                fmt.Println(v16)
            case "hello":
                v19 := GenS2{F0: 1, F1: genFn0()}
                _ = v19
            default:
                var v20 int64 = 485
                _ = v20
            }
        }
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
    } else {
        var v21 int8 = -1
        v1 = 142
        _ = v21
    }
    v2 += v12
    switch -2147483648 {
    case -9223372036854775808:
        v3 = (-548 >= 32767)
    case 1:
        var v22 int = -844
        _ = v22
    default:
        var v23 bool = false
        _ = v23
    }
    var v24 int32 = 1
    v12 -= v2
    v25 := GenS3{F0: -128, F1: v1, F2: (v12 / (((v2 << ((2) & 63))) | 1))}
    v26 := GenS4{F0: ("hello" + "foo bar")}
    v12 -= v12
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}
