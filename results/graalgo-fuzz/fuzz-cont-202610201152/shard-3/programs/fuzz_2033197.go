package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 bool }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 int; F1 string }
func genFn0(p0 uint32, p1 uint, p2 int64) bool {
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 bool = true
        p2 = p2
        p2 -= p2
        _ = fi7
        _ = v8
    }
    switch -9223372036854775808 {
    case 128:
        if true {
            p0 = p0
            for fi9 := 0; fi9 < 19; fi9++ {
                var v10 uint = 264
                if false {
                    p2 = (^(p2))
                    var v11 uint16 = 0
                    _ = v11
                }
                if true {
                    if (true || (!true)) {
                        v10 = p1
                        v12 := []int{2147483647}
                        var v13 bool = false
                        _ = v12
                        _ = v13
                    }
                    v14 := GenS1{F0: 2, F1: 127, F2: ((!false) || (42 != (p0 / ((uint32(p1)) | 1))))}
                    v15 := GenS1{F0: 645, F1: 1, F2: true}
                    _ = v14
                    _ = v15
                }
                _ = fi9
                _ = v10
            }
        }
    case -1:
        fmt.Println(p1)
    case 0:
        v16 := []int{176, 9223372036854775807, 517, 127}
        _ = v16
    }
    v17 := []int{32767}
    fmt.Println(p0)
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return false
}

func main() {
    v1 := []int{4294967295, 32768}
    var v2 float32 = 0.1
    v3 := make(map[string]bool)
    v4 := make(map[string]int)
    var v5 int = -709
    var v6 int16 = 0
    var v18 float64 = (0.0 / (float64(v2) / 1.5))
    v18 -= float64(((float32(v18) - float32(((0.1 * v18) - (3.14159 + v18)))) - 0.0))
    if true {
        for fi19 := 0; fi19 < 19; fi19++ {
            if (70 > -128) {
                for fi20 := 0; fi20 < 2; fi20++ {
                    v21 := GenS1{F0: 0, F1: 42, F2: false}
                    _ = fi20
                    _ = v21
                }
                fmt.Println(v18, v6)
                fmt.Println(v18, v2, v6)
            } else {
                var v22 bool = true
                for fi23 := 0; fi23 < 8; fi23++ {
                    var v24 int32 = 0
                    var v25 int16 = v6
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                var v26 bool = (!v22)
                _ = v22
                _ = v26
            }
            var v27 uint16 = 352
            for fi28 := 0; fi28 < 17; fi28++ {
                v29 := GenS1{F0: v27, F1: 1, F2: ((65535 == 2147483647) && false)}
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi28
                _ = v29
            }
            _ = fi19
            _ = v27
        }
        switch -1 {
        case 4294967295:
            for fi30 := 0; fi30 < 12; fi30++ {
                for fi31 := 0; fi31 < 17; fi31++ {
                    v32 := GenS1{F0: 2, F1: 101, F2: true}
                    v33 := GenS1{F0: 65535, F1: 0, F2: true}
                    _ = fi31
                    _ = v32
                    _ = v33
                }
                v34 := GenS2{F0: 4294967295, F1: 127}
                _ = fi30
                _ = v34
            }
        case 32768:
            v18 -= float64(1e10)
        }
        fmt.Println(v6)
    }
    fmt.Println(v5, v6, v5)
    v18 -= 0.0
    v2 -= v2
    var v35 uint32 = 0
    for fi36 := 0; fi36 < 20; fi36++ {
        var v37 uint32 = (v35 * (^(v35)))
        switch "hello" {
        case "foo bar":
            v38 := GenS3{F0: (fi36 & (fi36 / (((v5 >> ((65535) & 63))) | 1))), F1: "a"}
            _ = v38
        case "\n\t":
            v39 := v4[("" + "a")]
            _ = v39
        case "":
            var v40 int64 = 9223372036854775807
            _ = v40
        }
        _ = fi36
        _ = v37
    }
    fmt.Println(v2, v6, v6)
    v6 = (v6 >> ((0) & 63))
    v3[(("a" + ("hello" + "foo bar")) + "")] = (!false)
    fmt.Println(v6, v18)
    switch v5 {
    case -877:
        v35 = v35
    case -992:
        if len(v1) > 0 {
            v1[0] = v5
        }
    case 828:
        fmt.Println(v35)
    default:
        var v41 int8 = 1
        _ = v41
    }
    var v42 uint16 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
