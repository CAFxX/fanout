package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 uint16; F1 uint16 }

func main() {
    var v1 uint64 = 942
    var v2 bool = (127 <= 255)
    v3 := []uint32{2147483647, 374}
    var v4 float64 = 1e10
    v5 := make(map[string]uint32)
    switch 2147483647 {
    case -649:
        fmt.Println(v1, v1, v4)
    }
    if v2 {
        var v6 int32 = 194
        _ = v6
    } else {
        delete(v5, "hello")
    }
    var v7 uint64 = (^(v1))
    if true {
        if ("a" >= ("" + "\n\t")) {
            v4 -= v4
        }
    }
    if v2 {
        v2 = ((v1 + v1) > v7)
        switch ("" + "\n\t") {
        case "":
            v4 += (1.5 * v4)
        case "hello":
            var v8 int32 = -2147483648
            _ = v8
        case "a":
            v3 = append(v3, 807)
        }
    } else {
        if (!true) {
            var v9 float64 = v4
            v7 = v1
            switch 127 {
            case 1:
                _ = len(v5)
            default:
                if (true && true) {
                    v10 := GenS1{F0: 510, F1: 42}
                    _ = v10
                }
            }
            _ = v9
        }
        for fi11 := 0; fi11 < 12; fi11++ {
            switch "\n\t" {
            case "hello":
                var v12 int32 = 1
                _ = v12
            }
            var v13 uint32 = 0
            _ = fi11
            _ = v13
        }
    }
    var v14 int = 4294967295
    for fi15 := 0; fi15 < 15; fi15++ {
        var v16 int16 = -848
        _ = fi15
        _ = v16
    }
    v4 += (1.5 + float64(float32((float64((((3.14159 * 1.5) - ((1e10 / 1.5) / 1.5)) / (((0.1 - (((((-2.25 - float32(1e10)) / 1.5) / 1.5) * 0.1) + -2.25)) / ((((0.1 + 0.0) * (float32((((v4 * v4) / v4) * 1e10)) / 1.0)) - (1e10 - float32(float64(1e10)))) * (-2.25 / 1e10))) * 1e10))) + ((3.14159 + -2.25) / v4)))))
    if false {
        v2 = v2
    } else {
        var v17 int = v14
        v1 += 658
        _ = v17
    }
    v18 := GenS2{F0: 0, F1: 0}
    for fi19 := 0; fi19 < 2; fi19++ {
        for fi20 := 0; fi20 < 13; fi20++ {
            var v21 int = fi19
            var v22 int16 = 42
            _ = fi20
            _ = v21
            _ = v22
        }
        if (false && false) {
            var v23 bool = (!true)
            _ = v23
        } else {
            v4 = ((v4 / float64(3.14159)) * float64(1.5))
        }
        _ = fi19
    }
    var v24 int64 = 42
    delete(v5, "")
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
