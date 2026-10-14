package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 int16 }

func main() {
    var v1 int32 = -1
    var v2 float64 = 1.5
    var v3 uint64 = 18446744073709551615
    if false {
        v4 := GenS1{F0: 255}
        switch (("a" + ("a" + "\n\t")) + "") {
        case "":
            v1 += -334
        default:
            var v5 uint = 419
            _ = v5
        }
        _ = v4
    } else {
        v1 = v1
        if (true && false) {
            fmt.Println(v2, v1, v3)
            var v6 int = 32767
            _ = v6
        } else {
            var v7 uint = 18446744073709551615
            _ = v7
        }
    }
    v3 -= (v3 + v3)
    v8 := []int{-514}
    for fi9 := 0; fi9 < 11; fi9++ {
        var v10 int64 = -334
        if true {
            switch 0 {
            case 128:
                fmt.Println(v1, v3, v1)
            default:
                for fi11 := 0; fi11 < 9; fi11++ {
                    switch (fi11 ^ (fi9 << ((685) & 63))) {
                    case 674:
                        switch 128 {
                        case 882:
                            v3 = (v3 / ((v3) | 1))
                        }
                    case 2147483647:
                        var v12 uint64 = 0
                        _ = v12
                    }
                    fmt.Println(fi9, v10, v2)
                    fmt.Println(v10, fi9)
                    _ = fi11
                }
            }
        }
        switch ("foo bar" + "hello") {
        case "foo bar":
            v10 = v10
        case "a":
            var v13 float32 = float32(3.14159)
            _ = v13
        default:
            fmt.Println(v3, v10)
        }
        _ = fi9
        _ = v10
    }
    fmt.Println(v2)
    for fi14 := 0; fi14 < 2; fi14++ {
        for fi15 := 0; fi15 < 19; fi15++ {
            v1 = v1
            var v16 float64 = v2
            _ = fi15
            _ = v16
        }
        _ = fi14
    }
    fmt.Println(v3)
    v3 -= v3
    v17 := GenS2{F0: 51, F1: ("hello" + "a")}
    var v18 uint8 = 177
    v1 -= -2147483648
    for fi19 := 0; fi19 < 6; fi19++ {
        fmt.Println(v1, v2)
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = fi19
    }
    var v20 float64 = v2
    v18 -= (v18 * v18)
    var v21 uint32 = 0
    v22 := GenS3{F0: 32767}
    v2 += 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v8
}
