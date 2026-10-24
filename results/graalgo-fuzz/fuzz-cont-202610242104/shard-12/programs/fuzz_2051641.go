package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint32; F1 int }

func main() {
    var v1 int8 = 127
    v2 := make(map[string]bool)
    var v3 uint64 = 4294967295
    v4 := v2["a"]
    for fi5 := 0; fi5 < 19; fi5++ {
        v6 := GenS1{F0: 42, F1: 18446744073709551615}
        for fi7 := 0; fi7 < 16; fi7++ {
            var v8 uint8 = 0
            _ = fi7
            _ = v8
        }
        v9 := GenS1{F0: 1, F1: 18446744073709551615}
        _ = fi5
        _ = v6
        _ = v9
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        var v11 bool = true
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        v13 := GenS2{F0: 877, F1: fi12}
        for fi14 := 0; fi14 < 19; fi14++ {
            var v15 uint16 = 1
            fmt.Println(v1, v3, v3)
            _ = fi14
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    for fi16 := 0; fi16 < 12; fi16++ {
        var v17 int16 = 1000
        fmt.Println(v1, fi16)
        _ = fi16
        _ = v17
    }
    v2["foo bar"] = ((v3 / ((v3) | 1)) > v3)
    if v4 {
        v4 = v4
        for fi18 := 0; fi18 < 4; fi18++ {
            v3 = v3
            v1 -= v1
            _ = fi18
        }
        v1 = -65
    }
    v1 = v1
    var v19 int16 = 42
    fmt.Println(v19)
    switch ("\n\t" + "foo bar") {
    case "hello":
        fmt.Println(v3, v4, v1)
    case "\n\t":
        var v20 int8 = v1
        _ = v20
    }
    fmt.Println(v19, v1, v4)
    v3 = v3
    for fi21 := 0; fi21 < 12; fi21++ {
        var v22 int8 = v1
        v2["a"] = false
        if v4 {
            v23 := v2[("a" + "foo bar")]
            _ = v23
        }
        _ = fi21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
