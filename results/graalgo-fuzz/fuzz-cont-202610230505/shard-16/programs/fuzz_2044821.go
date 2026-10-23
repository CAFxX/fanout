package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string }
func genFn0(p0 int64, p1 uint8) int64 {
    if false {
        v6 := GenS1{F0: -117, F1: (18446744073709551615 > 4294967295), F2: 872}
        for fi7 := 0; fi7 < 9; fi7++ {
            v8 := GenS1{F0: 1000, F1: (-880 < -32768), F2: 255}
            _ = fi7
            _ = v8
        }
        switch "hello" {
        case "foo bar":
            switch "" {
            case "a":
                p0 = (p0 / (((p0 + 42)) | 1))
            default:
                switch -353 {
                case 9223372036854775807:
                    switch 2147483647 {
                    case 2147483647:
                        var v9 uint8 = 1
                        _ = v9
                    case 42:
                        if (true && false) {
                            fmt.Println(p1, p1)
                        }
                    }
                case -2147483648:
                    for fi10 := 0; fi10 < 4; fi10++ {
                        var v11 float32 = 3.14159
                        v12 := []int{(fi10 + 9223372036854775807), fi10, fi10, (fi10 | -1)}
                        p1 += p1
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                default:
                    p1 += p1
                }
            }
        case "":
            switch "foo bar" {
            case "foo bar":
                p1 -= 146
            }
        case "a":
            p0 += -1
        default:
            var v13 int = 9223372036854775807
            _ = v13
        }
        _ = v6
    } else {
        v14 := []int{910, -2147483648, 127, -1}
        fmt.Println(p0)
        _ = v14
    }
    var v15 string = (("a" + "\n\t") + "foo bar")
    switch v15 {
    case "foo bar":
        p1 = (p1 + (p1 * (p1 + (p1 + (p1 * p1)))))
    case "":
        v16 := GenS1{F0: -731, F1: true, F2: 32767}
        _ = v16
    }
    _ = p0
    _ = p1
    _ = v15
    return p0
}
func genFn1(p0 bool, p1 int32) int16 {
    for fi17 := 0; fi17 < 12; fi17++ {
        for fi18 := 0; fi18 < 9; fi18++ {
            p0 = false
            _ = fi18
        }
        for fi19 := 0; fi19 < 11; fi19++ {
            v20 := GenS2{F0: (("" + ("\n\t" + "")) + "")}
            v21 := []int{(^(fi19)), (fi17 ^ fi17), fi17}
            _ = fi19
            _ = v20
            _ = v21
        }
        _ = fi17
    }
    p0 = (!false)
    var v22 uint = 18446744073709551615
    v23 := []int{-858}
    _ = p0
    _ = p1
    _ = v22
    _ = v23
    return 42
}

func main() {
    var v1 int32 = 2147483647
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]int)
    v4 := make(map[int]int)
    var v5 uint16 = 999
    var v24 float64 = 1.5
    fmt.Println(v5)
    delete(v3, 4294967295)
    if false {
        v25 := GenS3{F0: ""}
        _ = v25
    } else {
        v26 := GenS2{F0: ("\n\t" + "foo bar")}
        fmt.Println(v24)
        _ = v26
    }
    var v27 uint64 = 42
    v24 += (float64(0.1) / (1.5 + (v24 - 1e10)))
    var v28 int8 = -128
    fmt.Println(v1, v28)
    fmt.Println(v5, v27, v24)
    v29 := v4[-1]
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
