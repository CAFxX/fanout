package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 uint; F1 int16; F2 int }
type GenS3 struct { F0 int64 }
func genFn0() string {
    var v6 float64 = (((float64((3.14159 - -2.25)) - 0.0) * 0.0) / 1.0)
    switch "\n\t" {
    case "":
        v6 = 3.14159
    case "foo bar":
        for fi7 := 0; fi7 < 3; fi7++ {
            if false {
                fmt.Println(fi7, fi7)
                var v8 uint16 = 481
                _ = v8
            }
            _ = fi7
        }
    }
    fmt.Println(v6)
    _ = v6
    return "a"
}

func main() {
    var v1 uint64 = 0
    v2 := []int{9223372036854775807, 32767}
    v3 := make(map[string]string)
    v4 := make(map[int]uint32)
    var v5 uint8 = 245
    var v9 float32 = (1.5 * (1.5 / 1.0))
    if false {
        var v10 uint16 = 969
        v11 := GenS1{F0: (^(v1)), F1: 128}
        var v12 string = ("" + "\n\t")
        _ = v10
        _ = v11
        _ = v12
    } else {
        v13 := GenS2{F0: 18446744073709551615, F1: 32767, F2: -765}
        _ = v13
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        var v15 int16 = int16(fi14)
        v15 -= (^(v15))
        _ = fi14
        _ = v15
    }
    var v16 string = "hello"
    if false {
        var v17 bool = false
        v18 := GenS1{F0: 1, F1: 42}
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 12; fi19++ {
        _ = len(v3)
        v16 = "a"
        var v20 int64 = -635
        _ = fi19
        _ = v20
    }
    if (!false) {
        v9 += v9
        var v21 uint8 = (^(v5))
        v21 = (v5 << ((427) & 63))
        _ = v21
    } else {
        v9 += (0.0 * 0.1)
        _ = len(v4)
    }
    v2 = append(v2, 42)
    v4[135] = uint32(v5)
    fmt.Println(v9, v9, v1)
    if false {
        switch (("foo bar" + ((("foo bar" + ("" + (v16 + (v16 + ("a" + ((v16 + ((v16 + v16) + v16)) + "\n\t")))))) + v16) + ("" + (v16 + "foo bar")))) + ("" + "a")) {
        case "":
            v5 = v5
        }
    } else {
        v1 = v1
        var v22 float32 = ((v9 - 1.5) * (v9 + ((3.14159 - (1.5 + 0.0)) * -2.25)))
        _ = v22
    }
    var v23 bool = false
    v24 := GenS3{F0: -103}
    for fi25 := 0; fi25 < 2; fi25++ {
        if true {
            v26 := v3[v16]
            _ = v26
        }
        v5 = (v5 - v5)
        _ = fi25
    }
    v27 := GenS3{F0: 1}
    v1 = (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
