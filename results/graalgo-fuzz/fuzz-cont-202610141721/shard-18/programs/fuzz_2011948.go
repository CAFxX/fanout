package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 uint; F1 uint }
func genFn0() string {
    v4 := GenS1{F0: 2147483647, F1: 18446744073709551615}
    v5 := GenS1{F0: 2147483647, F1: 4294967295}
    for fi6 := 0; fi6 < 13; fi6++ {
        for fi7 := 0; fi7 < 12; fi7++ {
            var v8 int = fi7
            v8 += v8
            v8 = (v8 ^ fi7)
            v8 += fi6
            var v9 uint32 = 2147483647
            _ = fi7
            _ = v8
            _ = v9
        }
        var v10 int16 = 1000
        fmt.Println(fi6)
        _ = fi6
        _ = v10
    }
    for fi11 := 0; fi11 < 19; fi11++ {
        switch 32768 {
        case -1:
            var v12 int64 = 206
            _ = v12
        default:
            switch 127 {
            case 0:
                var v13 uint8 = 22
                _ = v13
            case 9223372036854775807:
                v14 := []int{fi11, (fi11 << ((1) & 63))}
                _ = v14
            case -1:
                v15 := GenS1{F0: 42, F1: 379}
                _ = v15
            default:
                var v16 int32 = 0
                _ = v16
            }
        }
        v17 := GenS1{F0: 0, F1: 2}
        _ = fi11
        _ = v17
    }
    v18 := []int{-1, 42}
    _ = v18
    _ = v4
    _ = v5
    return "\n\t"
}

func main() {
    var v1 float64 = 0.1
    var v2 int32 = -2147483648
    var v3 bool = (!(-24 > -536))
    if false {
        v2 = v2
        var v19 uint64 = 58
        _ = v19
    }
    var v20 bool = v3
    for fi21 := 0; fi21 < 5; fi21++ {
        if true {
            v20 = v20
            var v22 uint8 = 72
            var v23 uint8 = v22
            _ = v22
            _ = v23
        }
        _ = fi21
    }
    var v24 uint32 = 1
    var v25 uint8 = 255
    v2 = (v2 | v2)
    v25 -= 2
    var v26 uint64 = 18446744073709551615
    for fi27 := 0; fi27 < 9; fi27++ {
        v28 := GenS1{F0: -954, F1: 454}
        if ("\n\t" < "\n\t") {
            v1 += float64(1.5)
        } else {
            var v29 int = 1
            _ = v29
        }
        switch 524 {
        case -1:
            v30 := []int{538, fi27}
            _ = v30
        case 472:
            if (47 <= 127) {
                fmt.Println(v3, fi27)
                if v3 {
                    var v31 int16 = 1
                    fmt.Println(v3, v25)
                    if (v26 == 42) {
                        v32 := []int{fi27, fi27, fi27, (^(fi27))}
                        _ = v32
                    }
                    _ = v31
                }
                v33 := GenS2{F0: (v26 << ((482) & 63)), F1: "\n\t"}
                _ = v33
            } else {
                v1 = v1
            }
        case 2147483647:
            v2 = v2
        }
        _ = fi27
        _ = v28
    }
    for fi34 := 0; fi34 < 15; fi34++ {
        v35 := GenS3{F0: 297, F1: 65535}
        var v36 int64 = 650
        _ = fi34
        _ = v35
        _ = v36
    }
    v37 := []int{2147483647, 902, 9223372036854775807}
    var v38 int16 = 32767
    v20 = false
    var v39 string = ""
    v1 = (0.0 * ((0.1 + -2.25) + 0.1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v37
    _ = v38
    _ = v39
}
