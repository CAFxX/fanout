package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 int32 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 int64 }
func genFn0(p0 uint16, p1 int16, p2 bool) int32 {
    p0 = p0
    v6 := []int{}
    v6 = append(v6, 813)
    var v7 string = (((("a" + "hello") + ("foo bar" + "a")) + ("a" + "")) + "foo bar")
    var v8 int16 = 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 2147483647
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 bool = ((65535 == 255) || false)
    var v3 string = "a"
    v4 := make(map[uint32]int)
    var v5 uint32 = 2
    v3 = ""
    if v2 {
        var v9 int32 = 42
        for fi10 := 0; fi10 < 4; fi10++ {
            v5 = v5
            for fi11 := 0; fi11 < 1; fi11++ {
                v12 := []int{(fi11 - (fi11 * 2147483647)), (fi11 / ((fi10) | 1))}
                var v13 int8 = 15
                v14 := GenS1{F0: (849 == 1), F1: 901, F2: (v9 + -2147483648)}
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
            }
            var v15 float64 = 3.14159
            _ = fi10
            _ = v15
        }
        switch "hello" {
        case "hello":
            v9 = v9
        default:
            switch 127 {
            case 32767:
                v1[(v5 % (((v5 / (((v5 * 2)) | 1))) | 1))] = 4294967295
            case 816:
                for fi16 := 0; fi16 < 13; fi16++ {
                    var v17 uint8 = 255
                    var v18 float32 = 1e10
                    var v19 int64 = -86
                    _ = fi16
                    _ = v17
                    _ = v18
                    _ = v19
                }
            case 128:
                v20 := []int{}
                _ = v20
            }
        }
        _ = v9
    } else {
        v3 = v3
        v21 := []int{-9223372036854775808, 144, -1}
        v22 := GenS2{F0: 0, F1: 65535}
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 6; fi23++ {
        var v24 uint8 = 32
        var v25 int32 = -2147483648
        _ = fi23
        _ = v24
        _ = v25
    }
    var v26 bool = v2
    var v27 uint16 = 0
    v28 := GenS3{F0: 42}
    var v29 int64 = 1
    v26 = (v2 && v2)
    var v30 int8 = 1
    v27 += (v27 | v27)
    for fi31 := 0; fi31 < 8; fi31++ {
        fmt.Println(v26)
        var v32 uint32 = (v5 | 838)
        _ = fi31
        _ = v32
    }
    fmt.Println(v29, v30, v3)
    switch 196 {
    case -9223372036854775808:
        v33 := v1[v5]
        _ = v33
    case -769:
        v29 -= (^(v29))
    }
    v34 := GenS1{F0: (((v26 && (!(v3 != v3))) || ((^(v27)) == v27)) && (!(2 >= 42))), F1: uint64(v5), F2: 0}
    var v35 float32 = 3.14159
    if v2 {
        switch v3 {
        case "\n\t":
            if v2 {
                v36 := []int{578}
                _ = v36
            }
        case "":
            var v37 float32 = 1.5
            _ = v37
        }
        v29 = 1
    } else {
        var v38 float64 = 3.14159
        _ = v38
    }
    var v39 uint16 = (v27 ^ 255)
    var v40 uint8 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
