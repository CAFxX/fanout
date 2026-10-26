package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 int16; F2 uint8 }
type GenS3 struct { F0 int32; F1 uint32 }
func genFn0() uint8 {
    var v5 int32 = -296
    switch ("\n\t" + "foo bar") {
    case "foo bar":
        if true {
            var v6 bool = (!false)
            _ = v6
        }
    case "a":
        for fi7 := 0; fi7 < 13; fi7++ {
            for fi8 := 0; fi8 < 7; fi8++ {
                v9 := []int{(fi8 * (^(fi7))), -9223372036854775808}
                switch "hello" {
                case "\n\t":
                    var v10 string = "hello"
                    _ = v10
                case "a":
                    var v11 bool = false
                    _ = v11
                default:
                    var v12 float64 = -2.25
                    _ = v12
                }
                v13 := GenS1{F0: 32767}
                v14 := GenS1{F0: 65535}
                if false {
                    var v15 int32 = v5
                    _ = v15
                }
                _ = fi8
                _ = v13
                _ = v14
                _ = v9
            }
            _ = fi7
        }
    case "\n\t":
        for fi16 := 0; fi16 < 3; fi16++ {
            v17 := []int{(fi16 * fi16)}
            var v18 int32 = 0
            switch "a" {
            case "foo bar":
                var v19 int32 = (v18 << ((65535) & 63))
                _ = v19
            }
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    var v20 int32 = (v5 & (-(v5)))
    if false {
        fmt.Println(v5, v20)
    } else {
        for fi21 := 0; fi21 < 12; fi21++ {
            var v22 uint8 = 2
            _ = fi21
            _ = v22
        }
    }
    v5 += (v20 + (^(v5)))
    _ = v20
    _ = v5
    return 175
}
func genFn1(p0 int, p1 int, p2 int) int32 {
    v23 := GenS1{F0: 111}
    var v24 float64 = ((3.14159 / 1e10) + 1.5)
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    _ = v24
    return 0
}

func main() {
    v1 := make(map[int]int)
    var v2 int = -384
    var v3 int = (v2 / (((v2 + (v2 + v2))) | 1))
    var v4 int8 = 110
    if (1 == 210) {
        var v25 float32 = (3.14159 / 1.0)
        v3 += (v3 & v3)
        _ = v25
    }
    v26 := GenS2{F0: 42, F1: -1, F2: 42}
    if (-9223372036854775808 < 148) {
        var v27 int = -1
        var v28 uint32 = 177
        _ = len(v1)
        _ = v27
        _ = v28
    }
    v1[v2] = -9223372036854775808
    v29 := []int{(v3 & v2), (v3 * (^(v3))), 1, v3}
    for fi30 := 0; fi30 < 16; fi30++ {
        v31 := GenS3{F0: 42, F1: 0}
        var v32 uint = 2
        _ = fi30
        _ = v31
        _ = v32
    }
    v2 += 42
    var v33 bool = false
    v34 := GenS1{F0: 255}
    var v35 bool = (("hello" + "hello") > "hello")
    v4 = (v4 << ((2) & 63))
    var v36 int64 = -619
    var v37 uint32 = 2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
}
