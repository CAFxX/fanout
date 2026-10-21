package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 uint8; F1 int32; F2 string }
func genFn0(p0 int32, p1 int64, p2 int16) uint64 {
    switch "foo bar" {
    case "":
        p2 = p2
    }
    for fi7 := 0; fi7 < 19; fi7++ {
        v8 := []int{}
        _ = fi7
        _ = v8
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 string = ("\n\t" + "hello")
    var v2 int = -9223372036854775808
    var v3 float64 = (float64(3.14159) - 3.14159)
    var v4 uint64 = 994
    var v5 bool = true
    var v6 uint32 = 2147483648
    if true {
        for fi9 := 0; fi9 < 2; fi9++ {
            var v10 float64 = v3
            _ = fi9
            _ = v10
        }
    } else {
        if v5 {
            for fi11 := 0; fi11 < 1; fi11++ {
                v12 := []int{v2, (v2 + (v2 * (v2 + fi11))), (v2 - v2), (v2 / (((-(fi11))) | 1))}
                _ = fi11
                _ = v12
            }
        }
    }
    var v13 bool = (2 != (v4 - uint64(v6)))
    switch v1 {
    case "foo bar":
        var v14 float32 = 3.14159
        _ = v14
    case "hello":
        fmt.Println(v13, v4, v13)
    case "":
        var v15 uint8 = 127
        _ = v15
    default:
        v2 -= (v2 - (v2 & v2))
    }
    for fi16 := 0; fi16 < 7; fi16++ {
        fmt.Println(fi16, v2, v5)
        for fi17 := 0; fi17 < 15; fi17++ {
            var v18 string = v1
            switch v18 {
            case "":
                var v19 string = v18
                _ = v19
            case "a":
                if (v3 > v3) {
                    if ("\n\t" < "a") {
                        v1 = v1
                        v20 := []int{128, (fi16 - v2)}
                        if (true && v5) {
                            fmt.Println(v2)
                            fmt.Println(v5, v1)
                            var v21 uint = 255
                            _ = v21
                        } else {
                            v18 = v1
                            var v22 int = (v2 / ((int(v6)) | 1))
                            _ = v22
                        }
                        v23 := GenS1{F0: (v6 + (v6 ^ (v6 / (((v6 | 879)) | 1)))), F1: 67, F2: v2}
                        _ = v20
                        _ = v23
                    }
                    var v24 int8 = 100
                    _ = v24
                }
            case "hello":
                var v25 float64 = v3
                _ = v25
            }
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    fmt.Println(v6, v3, v1)
    var v26 int = (v2 % ((v2) | 1))
    for fi27 := 0; fi27 < 17; fi27++ {
        v28 := GenS2{F0: 9223372036854775807, F1: 32767}
        v2 -= v2
        if v5 {
            var v29 uint = 1
            v30 := GenS2{F0: -1, F1: 1000}
            var v31 float64 = v3
            _ = v29
            _ = v30
            _ = v31
        } else {
            v6 = v6
            v32 := []int{32767, (-(v2)), (^(v2)), fi27}
            var v33 float32 = ((3.14159 + (1.5 - (0.0 * ((((((-2.25 / float32(((v3 * ((float64((((((0.0 - 0.1) - 0.1) / (float32(((((v3 - (v3 * v3)) / -2.25) * 1e10) - v3)) - 3.14159)) - 0.1) * float32(float64(0.1)))) + (float64(0.0) - v3)) - v3)) / v3))) - ((float32(1.5) / float32((v3 - 1.5))) - 1e10)) + (float32(v3) * float32(v3))) + 0.0) * -2.25) / 0.1)))) * 0.0)
            _ = v32
            _ = v33
        }
        _ = fi27
        _ = v28
    }
    fmt.Println(v1, v5, v5)
    for fi34 := 0; fi34 < 16; fi34++ {
        var v35 int64 = -747
        var v36 float64 = -2.25
        _ = fi34
        _ = v35
        _ = v36
    }
    var v37 string = v1
    v38 := []int{v2, v26}
    for fi39 := 0; fi39 < 11; fi39++ {
        var v40 float64 = 1.5
        var v41 float64 = v3
        v5 = v5
        _ = fi39
        _ = v40
        _ = v41
    }
    var v42 uint = 42
    if len(v38) > 0 {
        v38[0] = v26
    }
    v43 := GenS2{F0: 9223372036854775807, F1: 1000}
    v44 := GenS3{F0: 0, F1: -916, F2: v37}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v26
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
