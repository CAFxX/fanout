package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 bool; F1 int32 }

func main() {
    var v1 string = (((("\n\t" + "a") + ("hello" + "")) + ((("" + "hello") + "\n\t") + "\n\t")) + "hello")
    var v2 uint64 = 18446744073709551615
    var v3 uint = 18446744073709551615
    var v4 string = "hello"
    v5 := make(map[string]string)
    var v6 uint8 = 2
    v4 = v1
    for fi7 := 0; fi7 < 13; fi7++ {
        for fi8 := 0; fi8 < 17; fi8++ {
            v3 += v3
            switch ("\n\t" + v4) {
            case "a":
                v9 := v5["\n\t"]
                _ = v9
            default:
                var v10 bool = (((0.0 - 0.1) < -2.25) && true)
                _ = v10
            }
            fmt.Println(v6, fi8)
            _ = fi8
        }
        var v11 int16 = 0
        for fi12 := 0; fi12 < 8; fi12++ {
            v1 = ((("foo bar" + ("" + ("a" + (("" + "\n\t") + "\n\t")))) + ((v4 + (v1 + v4)) + v4)) + "\n\t")
            var v13 string = v1
            v14 := []int{fi7, fi12}
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = fi7
        _ = v11
    }
    v15 := GenS1{F0: (^(v2)), F1: 10}
    if false {
        _ = len(v5)
        _ = len(v5)
        v2 = v2
    } else {
        v16 := GenS1{F0: (v2 * v2), F1: v3}
        var v17 float64 = float64((float32(-2.25) * 1.5))
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 20; fi18++ {
        v19 := GenS1{F0: (v2 % (((v2 & v2)) | 1)), F1: v3}
        if true {
            var v20 int32 = 51
            _ = v20
        }
        v6 = uint8(v2)
        _ = fi18
        _ = v19
    }
    var v21 uint8 = 120
    if false {
        v22 := []int{}
        switch 42 {
        case 42:
            v21 -= uint8(v3)
        default:
            var v23 uint16 = 206
            _ = v23
        }
        if true {
            var v24 float64 = (float64((-2.25 + 1.5)) * -2.25)
            _ = v24
        } else {
            var v25 int8 = 42
            v4 = v1
            _ = v25
        }
        _ = v22
    } else {
        v26 := GenS2{F0: false, F1: 1}
        if (-37 <= 1) {
            _ = len(v5)
            fmt.Println(v4, v2)
        }
        _ = v26
    }
    var v27 int64 = -1
    v5[("hello" + "\n\t")] = v1
    v3 -= v3
    v27 -= (^(v27))
    v3 += 495
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
