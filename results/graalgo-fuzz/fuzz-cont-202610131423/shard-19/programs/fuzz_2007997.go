package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 uint8 = 2
    var v2 float32 = 1.5
    v3 := make(map[string]uint32)
    v4 := []string{"a", ("" + "a"), "foo bar"}
    v3["a"] = 1
    for fi5 := 0; fi5 < 16; fi5++ {
        var v6 uint32 = 430
        var v7 int64 = 42
        _ = fi5
        _ = v6
        _ = v7
    }
    switch "foo bar" {
    case "foo bar":
        var v8 float64 = ((1.5 * (((0.1 - 3.14159) + 0.0) * 3.14159)) / 1.0)
        _ = v8
    case "":
        if false {
            v9 := GenS1{F0: -128}
            _ = v9
        }
    case "hello":
        v1 = 42
    }
    var v10 float32 = -2.25
    if true {
        for fi11 := 0; fi11 < 16; fi11++ {
            for fi12 := 0; fi12 < 5; fi12++ {
                v1 -= 128
                _ = fi12
            }
            var v13 uint32 = 42
            _ = fi11
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 3; fi14++ {
        fmt.Println(v10, v1, fi14)
        v15 := GenS1{F0: 42}
        switch fi14 {
        case 32768:
            var v16 int64 = 767
            _ = v16
        case 143:
            v1 -= v1
        case 4294967295:
            v1 = 127
        }
        _ = fi14
        _ = v15
    }
    var v17 uint8 = v1
    for fi18 := 0; fi18 < 9; fi18++ {
        delete(v3, "foo bar")
        for fi19 := 0; fi19 < 10; fi19++ {
            v20 := GenS1{F0: -95}
            _ = fi19
            _ = v20
        }
        var v21 string = "\n\t"
        _ = fi18
        _ = v21
    }
    v22 := GenS1{F0: -128}
    v23 := GenS1{F0: 100}
    if false {
        if len(v4) > 0 {
            v4[0] = "\n\t"
        }
        fmt.Println(v1, v17, v10)
        v3["hello"] = 55
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
