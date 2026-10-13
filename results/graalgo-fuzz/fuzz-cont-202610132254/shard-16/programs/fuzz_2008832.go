package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 string; F1 int8; F2 string }
type GenS4 struct { F0 int; F1 uint32 }
type GenS5 struct { F0 int32; F1 bool }
type GenS6 struct { F0 uint8; F1 int64; F2 string }
func genFn0() int32 {
    for fi6 := 0; fi6 < 4; fi6++ {
        v7 := GenS1{F0: -9223372036854775808}
        var v8 float32 = -2.25
        v8 = v8
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 float64 = float64(0.1)
    _ = v9
    return 346
}
func genFn1(p0 int32, p1 bool, p2 int32) uint32 {
    v10 := GenS2{F0: 128, F1: false}
    p0 -= p0
    if p1 {
        v11 := GenS1{F0: -1}
        for fi12 := 0; fi12 < 8; fi12++ {
            v13 := GenS1{F0: 9223372036854775807}
            _ = fi12
            _ = v13
        }
        for fi14 := 0; fi14 < 10; fi14++ {
            p2 += (-(p2))
            v15 := []int{fi14, -2147483648, fi14, fi14}
            _ = fi14
            _ = v15
        }
        _ = v11
    } else {
        for fi16 := 0; fi16 < 11; fi16++ {
            switch "foo bar" {
            case "":
                var v17 float32 = 0.0
                _ = v17
            case "a":
                for fi18 := 0; fi18 < 10; fi18++ {
                    if false {
                        fmt.Println(fi18, p0, fi16)
                        v19 := GenS2{F0: 1, F1: (1 == 42)}
                        _ = v19
                    }
                    _ = fi18
                }
            default:
                v20 := []int{(-(fi16))}
                _ = v20
            }
            var v21 float32 = (float32((((-2.25 / 1.0) / 0.1) * ((1.5 * 0.0) + (0.1 - 3.14159)))) - 0.0)
            if p1 {
                p0 = p0
                if true {
                    v22 := GenS3{F0: "a", F1: 95, F2: "foo bar"}
                    _ = v22
                }
            } else {
                fmt.Println(v21, fi16, p2)
                v23 := GenS2{F0: 128, F1: (2 == 50)}
                _ = v23
            }
            _ = fi16
            _ = v21
        }
    }
    v24 := []int{128, -590}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v24
    return 2147483648
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint = 1
    var v3 uint8 = 127
    var v4 float32 = 1.5
    var v5 uint32 = 65535
    var v25 bool = (false || false)
    _ = len(v1)
    v26 := GenS4{F0: 343, F1: v5}
    var v27 string = ""
    for fi28 := 0; fi28 < 18; fi28++ {
        v29 := GenS4{F0: 4294967295, F1: genFn1(-1, false, genFn0())}
        for fi30 := 0; fi30 < 6; fi30++ {
            if (!(-2147483648 < 765)) {
                fmt.Println(v5, v4, v5)
                switch 1 {
                case 2147483647:
                    v31 := []int{(fi30 & int(v2)), -1, (fi30 * (fi28 >> (((v2 * uint(v5))) & 63))), 43}
                    _ = v31
                }
            }
            var v32 string = v27
            _ = fi30
            _ = v32
        }
        _ = fi28
        _ = v29
    }
    var v33 uint8 = 42
    v2 = (v2 & v2)
    v34 := GenS2{F0: (^(v33)), F1: (!v25)}
    v35 := GenS5{F0: 726, F1: v25}
    var v36 string = "a"
    v5 = v5
    var v37 string = v36
    switch "hello" {
    case "foo bar":
        var v38 int8 = 0
        _ = v38
    }
    v39 := []int{}
    switch "\n\t" {
    case "\n\t":
        if (!false) {
            var v40 int8 = -90
            v41 := GenS6{F0: v33, F1: -122, F2: (v37 + "foo bar")}
            _ = v40
            _ = v41
        } else {
            fmt.Println(v33, v33)
            var v42 bool = (!(v25 && v25))
            var v43 int16 = 185
            _ = v42
            _ = v43
        }
    default:
        v3 = v33
    }
    v33 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v39
    _ = v4
    _ = v5
}
