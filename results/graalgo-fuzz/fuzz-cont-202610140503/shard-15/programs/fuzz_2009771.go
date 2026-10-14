package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 uint8, p1 int8) uint {
    var v6 int32 = -2147483648
    for fi7 := 0; fi7 < 16; fi7++ {
        v8 := []int{}
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v6
    return 2
}

func main() {
    var v1 uint16 = 371
    v2 := []uint32{}
    var v3 int64 = 0
    v4 := []string{"", "foo bar"}
    v5 := make(map[uint32]uint32)
    if false {
        for fi9 := 0; fi9 < 13; fi9++ {
            var v10 string = "a"
            switch fi9 {
            case 4294967295:
                v10 = v10
            case -1:
                var v11 uint = 575
                _ = v11
            }
            _ = fi9
            _ = v10
        }
        var v12 string = (("\n\t" + "") + "\n\t")
        _ = v12
    } else {
        var v13 int8 = 127
        _ = v13
    }
    v14 := GenS1{F0: ("" + "foo bar")}
    v5[2147483647] = 2147483647
    v15 := GenS1{F0: "\n\t"}
    if ((float32(0.0) / 0.1) >= ((3.14159 + 0.0) + 3.14159)) {
        if true {
            v1 -= v1
            fmt.Println(v1, v3)
        }
        v3 = v3
        v1 = (^(v1))
    }
    switch "a" {
    case "hello":
        v4 = append(v4, "a")
    default:
        v16 := GenS1{F0: "\n\t"}
        _ = v16
    }
    v3 = 42
    var v17 uint16 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
