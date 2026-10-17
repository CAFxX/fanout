package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 int16 }
func genFn0(p0 int8, p1 int64, p2 uint16) uint8 {
    var v7 uint8 = 0
    var v8 int32 = 1
    v9 := GenS1{F0: 65535, F1: false, F2: 1}
    var v10 uint32 = 65535
    for fi11 := 0; fi11 < 18; fi11++ {
        var v12 string = (("" + ("\n\t" + "")) + "\n\t")
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return v7
}

func main() {
    var v1 float64 = 0.0
    var v2 int32 = 2147483647
    v3 := make(map[int]uint32)
    var v4 int8 = 42
    v5 := make(map[int]string)
    var v6 uint8 = 178
    v1 -= float64((float32((v1 / (1.5 / v1))) + ((1e10 * -2.25) * ((-2.25 / 1e10) * (1e10 + -2.25)))))
    v4 = (^(v4))
    v4 = (v4 + 100)
    var v13 string = ""
    var v14 uint32 = 1
    var v15 int8 = -1
    var v16 uint = 72
    fmt.Println(v16, v14, v4)
    for fi17 := 0; fi17 < 8; fi17++ {
        if true {
            v18 := []int{(^(fi17))}
            fmt.Println(v14, v2, v15)
            _ = v18
        }
        _ = fi17
    }
    v19 := v5[2147483647]
    v4 = 42
    for fi20 := 0; fi20 < 3; fi20++ {
        v21 := GenS1{F0: 2, F1: (false || true), F2: int16(fi20)}
        _ = fi20
        _ = v21
    }
    if true {
        v22 := []int{-200, 128, -786}
        v13 = (v19 + ((("foo bar" + ("" + v13)) + ("\n\t" + v19)) + ""))
        _ = v22
    }
    switch 128 {
    case 2147483647:
        switch 98 {
        case 4294967295:
            v15 += 127
        }
    case 42:
        v23 := GenS1{F0: 0, F1: (!((v16 < uint(v14)) || true)), F2: -32768}
        _ = v23
    default:
        for fi24 := 0; fi24 < 19; fi24++ {
            var v25 uint8 = (v6 - 237)
            var v26 string = ((("hello" + ((v13 + (v19 + "\n\t")) + "foo bar")) + ((v19 + v19) + ("foo bar" + ""))) + "\n\t")
            for fi27 := 0; fi27 < 2; fi27++ {
                fmt.Println(v19, fi24, v1)
                v28 := []int{}
                _ = fi27
                _ = v28
            }
            _ = fi24
            _ = v25
            _ = v26
        }
    }
    var v29 bool = (!(v14 < 712))
    for fi30 := 0; fi30 < 6; fi30++ {
        var v31 int = fi30
        _ = fi30
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
