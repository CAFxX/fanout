package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 uint }
type GenS2 struct { F0 int16; F1 int8; F2 bool }
type GenS3 struct { F0 bool; F1 int32; F2 uint32 }
type GenS4 struct { F0 string }
func genFn0() int64 {
    var v5 string = "foo bar"
    v6 := GenS1{F0: 241, F1: 32767, F2: 0}
    v5 = "a"
    if (0 <= 0) {
        v7 := []int{-9223372036854775808, 2147483647, 256}
        v5 = "hello"
        switch (((((v5 + v5) + "\n\t") + (((((v5 + v5) + (v5 + ("" + (v5 + (v5 + (v5 + "\n\t")))))) + v5) + v5) + v5)) + v5) + v5) {
        case "":
            v5 = v5
        case "foo bar":
            v5 = ""
        case "a":
            var v8 float32 = float32((1.5 / (((-2.25 - ((0.1 + (3.14159 - ((0.0 * 1.5) - 1.5))) + float64(((((0.1 / (-2.25 / 3.14159)) * -2.25) + 3.14159) * (0.1 / float32(-2.25)))))) / (0.0 + (1.5 - float64(0.0)))) - 0.1)))
            _ = v8
        default:
            var v9 uint16 = 368
            _ = v9
        }
        _ = v7
    }
    _ = v5
    _ = v6
    return -450
}
func genFn1(p0 int32, p1 int32, p2 int16) int64 {
    p2 = 0
    for fi10 := 0; fi10 < 18; fi10++ {
        var v11 float32 = (float32(float64(-2.25)) - -2.25)
        var v12 int16 = (p2 * p2)
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 uint16 = 0
    v14 := GenS2{F0: (-(p2)), F1: 127, F2: (3.14159 > (0.1 / 1.0))}
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return -1
}

func main() {
    var v1 float64 = 1.5
    var v2 bool = true
    var v3 float64 = (1.5 * 1.5)
    var v4 uint32 = 833
    v15 := []int{42, -385}
    for fi16 := 0; fi16 < 17; fi16++ {
        switch (fi16 / ((fi16) | 1)) {
        case 4294967295:
            var v17 int8 = 114
            _ = v17
        case 2147483647:
            v18 := GenS1{F0: 18446744073709551615, F1: 1, F2: 4294967295}
            _ = v18
        default:
            v19 := GenS1{F0: 453, F1: 415, F2: 42}
            _ = v19
        }
        if (42 == 127) {
            v2 = (v2 && (!v2))
            if (0 < -1) {
                if v2 {
                    var v20 uint32 = 4294967295
                    var v21 uint64 = 208
                    fmt.Println(v3, fi16, v4)
                    _ = v20
                    _ = v21
                }
            }
            if false {
                if (589 != 618) {
                    if (!v2) {
                        var v22 int16 = 1000
                        var v23 float32 = (1e10 * -2.25)
                        var v24 int32 = -2147483648
                        _ = v22
                        _ = v23
                        _ = v24
                    } else {
                        fmt.Println(v4, v4)
                    }
                    var v25 int = 127
                    _ = v25
                }
                var v26 int = fi16
                _ = v26
            } else {
                var v27 int8 = 127
                for fi28 := 0; fi28 < 18; fi28++ {
                    var v29 float32 = 3.14159
                    v15 = append(v15, fi28)
                    _ = fi28
                    _ = v29
                }
                _ = v27
            }
        } else {
            for fi30 := 0; fi30 < 17; fi30++ {
                v31 := GenS3{F0: (!true), F1: -1, F2: v4}
                v32 := GenS4{F0: "\n\t"}
                v33 := GenS2{F0: -947, F1: 19, F2: v2}
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
                _ = fi30
                _ = v31
                _ = v32
                _ = v33
            }
        }
        _ = fi16
    }
    fmt.Println(v3, v3)
    var v34 uint = 42
    var v35 int64 = 0
    if len(v15) > 0 {
        v15[0] = 4294967295
    }
    _ = len(v15)
    var v36 int64 = (v35 / (((v35 + v35)) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
