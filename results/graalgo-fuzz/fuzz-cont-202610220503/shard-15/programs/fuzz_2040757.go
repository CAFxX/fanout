package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 uint }
func genFn0(p0 int, p1 uint8) uint64 {
    for fi4 := 0; fi4 < 14; fi4++ {
        p0 += (^(fi4))
        _ = fi4
    }
    if (4294967295 == 4294967295) {
        p1 = 178
    }
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 string = "a"
    var v2 string = ("foo bar" + v1)
    v3 := make(map[int]int)
    for fi5 := 0; fi5 < 10; fi5++ {
        fmt.Println(v2, v1, fi5)
        var v6 int = (fi5 + (^(fi5)))
        _ = fi5
        _ = v6
    }
    var v7 int = 589
    v7 = 872
    var v8 float32 = float32(-2.25)
    v1 = ((v2 + v2) + ((v1 + v2) + v1))
    v1 = ("a" + ("" + "a"))
    for fi9 := 0; fi9 < 12; fi9++ {
        for fi10 := 0; fi10 < 15; fi10++ {
            if true {
                var v11 string = ("a" + v2)
                var v12 uint8 = 0
                v12 = 145
                _ = v11
                _ = v12
            }
            fmt.Println(fi10, v1)
            v13 := GenS1{F0: -22, F1: -349}
            _ = fi10
            _ = v13
        }
        _ = fi9
    }
    if false {
        for fi14 := 0; fi14 < 1; fi14++ {
            for fi15 := 0; fi15 < 8; fi15++ {
                switch "\n\t" {
                case "":
                    _ = len(v3)
                case "\n\t":
                    v16 := GenS2{F0: 42}
                    switch "\n\t" {
                    case "hello":
                        fmt.Println(v2)
                    case "\n\t":
                        v17 := []int{v7}
                        _ = v17
                    case "a":
                        v18 := GenS3{F0: 159, F1: 255}
                        _ = v18
                    }
                    _ = v16
                case "a":
                    var v19 float64 = 0.0
                    _ = v19
                default:
                    var v20 int64 = 0
                    _ = v20
                }
                v3[(fi15 >> ((4294967295) & 63))] = fi14
                _ = fi15
            }
            fmt.Println(v8, v7)
            _ = fi14
        }
    }
    v1 = v1
    fmt.Println(v7, v2, v1)
    for fi21 := 0; fi21 < 1; fi21++ {
        switch ("\n\t" + v2) {
        case "":
            fmt.Println(v1, v1)
        case "hello":
            fmt.Println(fi21)
        case "a":
            var v22 float32 = 0.1
            _ = v22
        default:
            v2 = v1
        }
        _ = fi21
    }
    v8 += float32((1.5 / float64(v8)))
    for fi23 := 0; fi23 < 14; fi23++ {
        for fi24 := 0; fi24 < 11; fi24++ {
            var v25 int16 = 1000
            v26 := []int{(fi24 * (^(v7)))}
            _ = fi24
            _ = v25
            _ = v26
        }
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
