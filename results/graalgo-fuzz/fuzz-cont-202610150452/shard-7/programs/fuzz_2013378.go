package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]bool)
    var v4 int = 42
    v4 = (-(v4))
    for fi5 := 0; fi5 < 9; fi5++ {
        for fi6 := 0; fi6 < 6; fi6++ {
            v7 := GenS1{F0: 42}
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    for fi8 := 0; fi8 < 13; fi8++ {
        if true {
            for fi9 := 0; fi9 < 20; fi9++ {
                v10 := GenS1{F0: 32767}
                _ = fi9
                _ = v10
            }
            var v11 float64 = 1.5
            fmt.Println(v11)
            _ = v11
        }
        switch "a" {
        case "foo bar":
            var v12 int16 = 32767
            _ = v12
        case "a":
            var v13 int32 = -1
            _ = v13
        }
        for fi14 := 0; fi14 < 12; fi14++ {
            var v15 int64 = 42
            v16 := []int{(fi8 + 32768), (^(fi8)), v4}
            switch -1 {
            case 639:
                var v17 int64 = v15
                _ = v17
            case 32767:
                v4 += 4294967295
            }
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi8
    }
    switch (("" + ("a" + "\n\t")) + "\n\t") {
    case "\n\t":
        v18 := GenS1{F0: 255}
        _ = v18
    case "a":
        switch v4 {
        case -432:
            v19 := GenS1{F0: 255}
            _ = v19
        default:
            v4 = v4
        }
    case "":
        var v20 int16 = 1000
        _ = v20
    default:
        v4 = 1
    }
    v21 := []int{(v4 & 42)}
    _ = len(v21)
    v22 := GenS1{F0: 250}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
