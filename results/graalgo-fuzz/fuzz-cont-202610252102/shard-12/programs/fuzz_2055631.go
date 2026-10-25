package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 int32 }
func genFn0(p0 uint) bool {
    v6 := GenS1{F0: 0}
    var v7 int16 = -32768
    var v8 float64 = 1.5
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return false
}
func genFn1(p0 int64, p1 int64, p2 bool) int8 {
    switch -2147483648 {
    case -550:
        v9 := GenS2{F0: 4294967295, F1: 62}
        _ = v9
    case 127:
        if (!p2) {
            var v10 uint8 = 169
            var v11 uint16 = 32767
            var v12 uint32 = 796
            _ = v10
            _ = v11
            _ = v12
        }
    case 805:
        for fi13 := 0; fi13 < 18; fi13++ {
            fmt.Println(p2)
            _ = fi13
        }
    }
    if p2 {
        if true {
            p0 += p0
            v14 := GenS3{F0: 255, F1: 1}
            _ = v14
        } else {
            switch "foo bar" {
            case "\n\t":
                var v15 uint16 = 37
                _ = v15
            case "hello":
                for fi16 := 0; fi16 < 15; fi16++ {
                    var v17 int8 = int8(p0)
                    fmt.Println(p0)
                    for fi18 := 0; fi18 < 3; fi18++ {
                        p1 = 42
                        v19 := GenS1{F0: -1}
                        _ = fi18
                        _ = v19
                    }
                    _ = fi16
                    _ = v17
                }
            case "a":
                switch 326 {
                case 32767:
                    var v20 int16 = 1000
                    _ = v20
                case -9223372036854775808:
                    v21 := []int{128, -2147483648, 128}
                    _ = v21
                case 2147483647:
                    var v22 int32 = -1
                    _ = v22
                default:
                    for fi23 := 0; fi23 < 2; fi23++ {
                        fmt.Println(p1, fi23, fi23)
                        _ = fi23
                    }
                }
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 81
}

func main() {
    var v1 int16 = 32767
    v2 := make(map[string]string)
    var v3 bool = (((0.1 * ((-2.25 - (float32(-2.25) - (1.5 - (1e10 - (((1.5 / 1.5) + (3.14159 * (((1.5 + 0.0) * 1e10) * (-2.25 - (3.14159 + -2.25))))) / 1e10))))) / 0.1)) >= (1.5 - (1e10 / float32(float64(1e10))))) || false)
    var v4 int8 = 8
    v5 := []uint32{4294967295}
    var v24 uint8 = 42
    var v25 uint32 = 2147483647
    var v26 string = ("a" + ("hello" + ((("hello" + "a") + "a") + "foo bar")))
    var v27 bool = ((v3 || ((v3 && v3) && v3)) && (642 >= uint(v25)))
    switch ("\n\t" + "") {
    case "foo bar":
        var v28 uint8 = v24
        _ = v28
    case "hello":
        fmt.Println(v27, v27, v1)
    case "":
        var v29 int32 = -1
        _ = v29
    default:
        v26 = v26
    }
    var v30 float32 = 0.1
    v1 -= v1
    v31 := GenS4{F0: 555}
    var v32 bool = ((1e10 < (3.14159 + float64((0.1 * 1.5)))) && (1 >= (v24 % ((227) | 1))))
    switch "hello" {
    case "hello":
        v33 := GenS1{F0: v4}
        _ = v33
    default:
        if true {
            v34 := GenS4{F0: -2147483648}
            var v35 int8 = v4
            var v36 uint16 = 633
            _ = v34
            _ = v35
            _ = v36
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
