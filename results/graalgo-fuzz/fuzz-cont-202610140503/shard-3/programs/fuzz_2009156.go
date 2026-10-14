package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64 }
func genFn0(p0 uint) string {
    if false {
        v6 := []int{1, 4294967295}
        for fi7 := 0; fi7 < 16; fi7++ {
            for fi8 := 0; fi8 < 19; fi8++ {
                var v9 bool = ((-2.25 <= 0.0) && false)
                _ = fi8
                _ = v9
            }
            _ = fi7
        }
        _ = v6
    }
    p0 -= 0
    _ = p0
    return "hello"
}
func genFn1(p0 uint, p1 uint64) int64 {
    if ("\n\t" > "hello") {
        for fi10 := 0; fi10 < 16; fi10++ {
            var v11 int64 = 42
            p1 -= p1
            v11 = 315
            _ = fi10
            _ = v11
        }
        var v12 float32 = ((1.5 + 1.5) / -2.25)
        v13 := GenS1{F0: (((("a" + "\n\t") + "hello") + ("foo bar" + "a")) + "\n\t"), F1: 58}
        _ = v12
        _ = v13
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    return 9223372036854775807
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []int{-668}
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]uint32)
    var v5 uint32 = 2147483647
    if len(v2) > 0 {
        v2[0] = 32768
    }
    v14 := GenS1{F0: "foo bar", F1: 172}
    if (56 == -128) {
        for fi15 := 0; fi15 < 9; fi15++ {
            for fi16 := 0; fi16 < 1; fi16++ {
                switch "foo bar" {
                case "hello":
                    var v17 uint8 = 127
                    v17 += v17
                    _ = v17
                case "a":
                    var v18 bool = true
                    _ = v18
                }
                _ = fi16
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi15
        }
        for fi19 := 0; fi19 < 16; fi19++ {
            var v20 float64 = 1e10
            var v21 int = (fi19 << ((53) & 63))
            v22 := v4[v5]
            _ = fi19
            _ = v20
            _ = v21
            _ = v22
        }
    } else {
        for fi23 := 0; fi23 < 10; fi23++ {
            var v24 string = ""
            var v25 uint8 = 255
            _ = fi23
            _ = v24
            _ = v25
        }
        v26 := GenS1{F0: "hello", F1: 178}
        _ = v26
    }
    if false {
        v5 -= (v5 | v5)
        v5 -= 2147483647
    }
    var v27 int16 = -389
    var v28 uint32 = v5
    fmt.Println(v28)
    for fi29 := 0; fi29 < 8; fi29++ {
        switch ("" + ("foo bar" + "foo bar")) {
        case "hello":
            var v30 int16 = (v27 >> ((255) & 63))
            _ = v30
        case "a":
            if len(v2) > 0 {
                v2[0] = (fi29 % ((fi29) | 1))
            }
        case "foo bar":
            v31 := v1[fi29]
            _ = v31
        }
        v32 := GenS1{F0: "\n\t", F1: 0}
        v33 := GenS2{F0: true}
        _ = fi29
        _ = v32
        _ = v33
    }
    for fi34 := 0; fi34 < 12; fi34++ {
        fmt.Println(v28, v5, v27)
        _ = fi34
    }
    switch -2147483648 {
    case -2147483648:
        v35 := GenS3{F0: 9223372036854775807}
        _ = v35
    case 598:
        fmt.Println(v27)
    case 4294967295:
        switch -1 {
        case 4294967295:
            fmt.Println(v28)
        case -2147483648:
            delete(v1, 32768)
        case 32767:
            var v36 float32 = 3.14159
            _ = v36
        default:
            var v37 int = -1
            _ = v37
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
