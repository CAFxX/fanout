package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 int32 }
func genFn0(p0 int, p1 uint16, p2 uint8) string {
    v7 := GenS1{F0: "", F1: false, F2: -517}
    switch ("" + ((("hello" + ((("hello" + "a") + "foo bar") + "\n\t")) + ("\n\t" + ("\n\t" + ""))) + "foo bar")) {
    case "\n\t":
        if true {
            fmt.Println(p2, p2, p2)
        }
    default:
        var v8 int32 = -1
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return "\n\t"
}

func main() {
    var v1 float64 = 0.0
    v2 := make(map[uint32]int)
    v3 := []uint32{1}
    v4 := make(map[uint32]string)
    v5 := []uint32{128, 4294967295}
    var v6 float32 = 1e10
    var v9 uint32 = 2
    var v10 float32 = ((1e10 / (float32(0.1) + v6)) + v6)
    for fi11 := 0; fi11 < 1; fi11++ {
        for fi12 := 0; fi12 < 6; fi12++ {
            var v13 int = (fi12 << ((18446744073709551615) & 63))
            _ = fi12
            _ = v13
        }
        switch (((("" + ("foo bar" + "a")) + "a") + "a") + ("foo bar" + "")) {
        case "foo bar":
            if (true && true) {
                var v14 uint8 = 29
                _ = v14
            } else {
                if (0 >= -32768) {
                    _ = len(v3)
                    v15 := v4[v9]
                    _ = v15
                } else {
                    _ = len(v4)
                }
                v3 = append(v3, (^(v9)))
                v1 -= (3.14159 / v1)
            }
        case "a":
            var v16 uint64 = 0
            _ = v16
        default:
            v10 += 3.14159
        }
        _ = fi11
    }
    var v17 int32 = -2147483648
    var v18 uint64 = 564
    for fi19 := 0; fi19 < 5; fi19++ {
        v1 = -2.25
        _ = fi19
    }
    delete(v4, v9)
    v17 = 0
    v17 = -304
    if (!(!("hello" != ("" + "\n\t")))) {
        fmt.Println(v18)
        v1 = 0.1
    } else {
        var v20 uint = 365
        var v21 bool = (-1 != 32768)
        _ = v20
        _ = v21
    }
    v6 += 3.14159
    if ("a" <= "\n\t") {
        v22 := v2[(v9 + 1)]
        v3 = append(v3, (v9 >> ((42) & 63)))
        _ = v22
    } else {
        v10 += float32(1.5)
    }
    switch "hello" {
    case "":
        v23 := GenS1{F0: "a", F1: true, F2: 2147483647}
        _ = v23
    case "\n\t":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
