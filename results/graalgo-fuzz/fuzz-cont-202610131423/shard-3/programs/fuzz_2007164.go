package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 uint16 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 string) int64 {
    var v6 uint32 = 1
    var v7 int16 = 32767
    _ = p0
    _ = v6
    _ = v7
    return 527
}
func genFn1(p0 int8, p1 uint8) int16 {
    fmt.Println(p0, p1)
    p1 += (^(p1))
    _ = p0
    _ = p1
    return 32767
}

func main() {
    v1 := make(map[int]int)
    var v2 float32 = (((1e10 + float32(1.5)) - 1e10) + 0.1)
    v3 := make(map[int]uint32)
    var v4 bool = (!false)
    v5 := []string{"a", "\n\t", "a", ("" + "\n\t")}
    v4 = (v4 || true)
    v8 := GenS1{F0: (((1 >= -789) && v4) || v4), F1: 1, F2: 255}
    fmt.Println(v4)
    if true {
        for fi9 := 0; fi9 < 4; fi9++ {
            fmt.Println(fi9)
            for fi10 := 0; fi10 < 19; fi10++ {
                v11 := GenS1{F0: v4, F1: (^(fi9)), F2: 603}
                switch (("\n\t" + ("a" + "\n\t")) + "foo bar") {
                case "\n\t":
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                _ = fi10
                _ = v11
            }
            _ = fi9
        }
        for fi12 := 0; fi12 < 13; fi12++ {
            var v13 int32 = 444
            switch (((("foo bar" + "") + ("" + ("a" + ("\n\t" + "")))) + "") + ("foo bar" + "hello")) {
            case "\n\t":
                v14 := GenS1{F0: v4, F1: 4294967295, F2: 761}
                _ = v14
            default:
                fmt.Println(v4, fi12, v2)
            }
            _ = fi12
            _ = v13
        }
        var v15 int8 = -64
        _ = v15
    }
    v2 -= (0.1 + v2)
    v16 := GenS1{F0: (!v4), F1: 309, F2: 550}
    v17 := GenS2{F0: 766}
    v18 := GenS2{F0: 144}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v5 = append(v5, "\n\t")
    if true {
        _ = len(v3)
    }
    v2 -= ((float32((1.5 + 0.1)) + (v2 + (float32(0.0) + v2))) * 0.0)
    var v19 uint16 = 65535
    switch ("\n\t" + "\n\t") {
    case "a":
        var v20 bool = (255 < 1)
        _ = v20
    case "\n\t":
        fmt.Println(v2, v2, v2)
    case "":
        switch 2147483647 {
        case -645:
            var v21 uint64 = 0
            _ = v21
        case 127:
            v22 := v1[42]
            _ = v22
        }
    default:
        switch 32768 {
        case 128:
            v19 -= 255
        }
    }
    var v23 uint8 = 72
    var v24 uint32 = 2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
