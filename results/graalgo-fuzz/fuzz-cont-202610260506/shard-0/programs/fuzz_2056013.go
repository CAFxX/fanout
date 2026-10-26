package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 int16 }
func genFn0(p0 string, p1 uint64) int64 {
    v4 := GenS1{F0: 0, F1: p1, F2: 0}
    v5 := GenS1{F0: 550, F1: 0, F2: 42}
    var v6 bool = (-664 >= -1)
    if v6 {
        switch p0 {
        case "hello":
            p0 = "hello"
        case "foo bar":
            for fi7 := 0; fi7 < 7; fi7++ {
                switch "\n\t" {
                case "":
                    p1 = 518
                case "hello":
                    var v8 float64 = 0.1
                    _ = v8
                case "\n\t":
                    var v9 string = p0
                    _ = v9
                default:
                    v10 := []int{32767, 127}
                    _ = v10
                }
                _ = fi7
            }
        case "":
            v11 := []int{2147483647, -347, -157}
            _ = v11
        }
        var v12 int32 = 42
        p0 = (((p0 + p0) + "\n\t") + (p0 + "\n\t"))
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 879
}

func main() {
    var v1 int32 = 42
    var v2 uint32 = 732
    var v3 uint64 = 2
    v13 := []int{32768, 2147483647, 128, 4294967295}
    if false {
        var v14 uint8 = 42
        _ = v14
    }
    switch (("a" + ("foo bar" + "\n\t")) + ("" + "a")) {
    case "a":
        fmt.Println(v2, v3)
    case "foo bar":
        v1 += 1
    }
    var v15 bool = (!(v3 != 490))
    v1 += v1
    var v16 uint = 1
    for fi17 := 0; fi17 < 11; fi17++ {
        var v18 int = fi17
        v1 -= v1
        v1 += -2147483648
        _ = fi17
        _ = v18
    }
    fmt.Println(v3, v2)
    v15 = v15
    var v19 uint32 = (^(v2))
    var v20 int32 = int32(v16)
    if false {
        var v21 uint16 = uint16(v19)
        if v15 {
            var v22 uint64 = 1
            _ = v22
        }
        var v23 uint16 = v21
        _ = v21
        _ = v23
    } else {
        if len(v13) > 0 {
            v13[0] = 9223372036854775807
        }
        v24 := GenS1{F0: 32767, F1: v3, F2: 32767}
        _ = v24
    }
    if (v15 && ((uint8(v3) >= 114) && false)) {
        v3 = (v3 % ((v3) | 1))
        v1 -= v1
        v25 := GenS1{F0: 814, F1: (^(v3)), F2: 0}
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
