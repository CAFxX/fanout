package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 int64; F1 uint16; F2 uint8 }
type GenS3 struct { F0 uint32; F1 int }
func genFn0() uint8 {
    var v6 int = -2147483648
    if false {
        v7 := GenS1{F0: -1, F1: (("hello" + ("foo bar" + "\n\t")) + "foo bar")}
        if (!(!(false || (255 == 128)))) {
            if (0 == 9) {
                v6 -= 42
            }
            var v8 float64 = -2.25
            _ = v8
        }
        fmt.Println(v6)
        _ = v7
    } else {
        var v9 uint16 = 42
        v9 = v9
        v10 := []int{v6, v6, v6}
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: 2147483647, F1: "hello"}
    switch "a" {
    case "foo bar":
        v6 += (-(v6))
    case "hello":
        for fi12 := 0; fi12 < 18; fi12++ {
            switch "a" {
            case "a":
                var v13 uint32 = 893
                _ = v13
            default:
                var v14 int8 = 100
                _ = v14
            }
            _ = fi12
        }
    default:
        fmt.Println(v6)
    }
    var v15 uint = 4294967295
    _ = v11
    _ = v15
    _ = v6
    return 2
}

func main() {
    v1 := []int{32768, 537, 199}
    v2 := make(map[int]uint32)
    var v3 int32 = 1
    v4 := make(map[uint32]uint32)
    v5 := make(map[uint32]string)
    switch "a" {
    case "\n\t":
        for fi16 := 0; fi16 < 13; fi16++ {
            var v17 uint16 = 42
            for fi18 := 0; fi18 < 18; fi18++ {
                var v19 uint32 = 42
                _ = fi18
                _ = v19
            }
            _ = fi16
            _ = v17
        }
    case "a":
        var v20 string = ("hello" + ("foo bar" + ""))
        _ = v20
    case "foo bar":
        var v21 int8 = 127
        _ = v21
    }
    v3 = (v3 * 261)
    if true {
        switch "" {
        case "foo bar":
            fmt.Println(v3)
        }
    }
    v22 := GenS1{F0: 59, F1: "hello"}
    v23 := GenS1{F0: 355, F1: "\n\t"}
    var v24 int = -591
    delete(v2, (^(v24)))
    var v25 bool = false
    fmt.Println(v24)
    var v26 float32 = 3.14159
    v27 := GenS1{F0: (v3 >> ((423) & 63)), F1: "\n\t"}
    var v28 string = ""
    for fi29 := 0; fi29 < 17; fi29++ {
        _ = len(v5)
        v30 := GenS2{F0: 42, F1: 1, F2: 42}
        for fi31 := 0; fi31 < 14; fi31++ {
            fmt.Println(fi29, fi31)
            _ = fi31
        }
        _ = fi29
        _ = v30
    }
    var v32 int64 = 9223372036854775807
    v33 := GenS3{F0: 379, F1: -433}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
