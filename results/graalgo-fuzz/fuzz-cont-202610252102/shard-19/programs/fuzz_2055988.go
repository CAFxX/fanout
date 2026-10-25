package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 int }

func main() {
    var v1 uint8 = 255
    var v2 float64 = float64(0.0)
    var v3 uint64 = 600
    switch -9223372036854775808 {
    case 2147483647:
        v2 -= (v2 * (0.0 * v2))
    }
    switch ("\n\t" + "\n\t") {
    case "foo bar":
        var v4 uint16 = 0
        _ = v4
    }
    for fi5 := 0; fi5 < 16; fi5++ {
        v6 := []int{}
        if true {
            v7 := GenS1{F0: 0, F1: (fi5 << ((uint(v1)) & 63)), F2: 32768}
            v3 -= v3
            switch fi5 {
            case -407:
                v6 = append(v6, fi5)
            }
            _ = v7
        }
        for fi8 := 0; fi8 < 20; fi8++ {
            var v9 float64 = v2
            _ = fi8
            _ = v9
        }
        _ = fi5
        _ = v6
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 uint64 = v3
        _ = fi10
        _ = v11
    }
    v3 -= uint64(v1)
    for fi12 := 0; fi12 < 12; fi12++ {
        v13 := GenS1{F0: 1, F1: -1, F2: -1}
        var v14 float64 = v2
        _ = fi12
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 4; fi15++ {
        var v16 int16 = 425
        var v17 uint16 = 2
        _ = fi15
        _ = v16
        _ = v17
    }
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
}
