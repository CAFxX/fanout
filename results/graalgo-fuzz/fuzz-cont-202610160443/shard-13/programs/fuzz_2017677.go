package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 uint32; F1 int }
type GenS4 struct { F0 int16; F1 string; F2 uint32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint32 }
func genFn0(p0 uint32) int32 {
    for fi4 := 0; fi4 < 5; fi4++ {
        v5 := GenS1{F0: -9223372036854775808}
        for fi6 := 0; fi6 < 12; fi6++ {
            v7 := GenS2{F0: true, F1: -57}
            var v8 int64 = 0
            v9 := GenS3{F0: (p0 + (p0 % (((^(p0))) | 1))), F1: fi6}
            v8 = v8
            var v10 float64 = ((1.5 * float64(float32((-2.25 * 1.5)))) * 1e10)
            v8 -= 91
            _ = fi6
            _ = v10
            _ = v7
            _ = v8
            _ = v9
        }
        v11 := []int{}
        _ = fi4
        _ = v11
        _ = v5
    }
    fmt.Println(p0)
    var v12 int64 = -973
    v13 := []int{9223372036854775807}
    v14 := GenS2{F0: false, F1: 0}
    _ = p0
    _ = v12
    _ = v13
    _ = v14
    return 42
}

func main() {
    var v1 uint8 = 42
    var v2 uint16 = 0
    var v3 float32 = ((3.14159 + -2.25) * 1.5)
    for fi15 := 0; fi15 < 5; fi15++ {
        fmt.Println(v3, fi15)
        v16 := []int{}
        if (!((!((0.1 != 0.0) || ((!(51 <= 1)) && (((fi15 >> ((255) & 63)) == fi15) || false)))) && true)) {
            var v17 bool = ((true && false) || false)
            _ = v17
        }
        _ = fi15
        _ = v16
    }
    v18 := []int{9223372036854775807, -2147483648}
    var v19 int64 = int64(v1)
    v19 += v19
    for fi20 := 0; fi20 < 16; fi20++ {
        v21 := GenS4{F0: -1, F1: "a", F2: 2}
        switch fi20 {
        case 127:
            if len(v18) >= 2 {
                v18 = v18[1:len(v18)]
            }
        case -761:
            var v22 uint32 = 2147483647
            _ = v22
        case -2147483648:
            var v23 string = ((("" + "foo bar") + "\n\t") + "hello")
            _ = v23
        default:
            fmt.Println(v19, v3, v19)
        }
        for fi24 := 0; fi24 < 3; fi24++ {
            switch -2147483648 {
            case 2147483647:
                var v25 float64 = ((1.5 * 0.1) / -2.25)
                _ = v25
            case 32768:
                var v26 uint = 42
                _ = v26
            case 1:
                var v27 int = fi20
                _ = v27
            default:
                v28 := GenS3{F0: 2147483648, F1: (fi24 | (fi20 & (fi24 | (fi24 * (fi20 + fi20)))))}
                if true {
                    v19 = (v19 >> ((2) & 63))
                    var v29 int = (fi20 - 121)
                    var v30 int32 = genFn0(0)
                    _ = v29
                    _ = v30
                } else {
                    var v31 int16 = -32768
                    switch "" {
                    case "":
                        v32 := GenS1{F0: fi20}
                        _ = v32
                    default:
                        var v33 bool = false
                        _ = v33
                    }
                    v18 = append(v18, (fi24 / (((fi24 % ((fi20) | 1))) | 1)))
                    _ = v31
                }
                v34 := GenS3{F0: 2147483648, F1: (fi20 | -718)}
                _ = v28
                _ = v34
            }
            _ = fi24
        }
        _ = fi20
        _ = v21
    }
    var v35 float64 = (1e10 + 3.14159)
    var v36 uint8 = (^(v1))
    for fi37 := 0; fi37 < 2; fi37++ {
        v35 = v35
        fmt.Println(v2)
        var v38 int32 = 42
        _ = fi37
        _ = v38
    }
    v39 := GenS5{F0: -1}
    if len(v18) > 0 {
        v18[0] = 1
    }
    var v40 int = 32768
    v40 = 753
    v35 = 1e10
    var v41 bool = false
    v18 = append(v18, v40)
    v42 := GenS6{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v39
    _ = v40
    _ = v41
    _ = v42
}
