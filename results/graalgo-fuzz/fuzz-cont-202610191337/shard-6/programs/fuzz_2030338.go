package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 uint64 = 1
    v2 := make(map[string]string)
    var v3 bool = false
    for fi4 := 0; fi4 < 6; fi4++ {
        switch (("\n\t" + "") + ("hello" + "hello")) {
        case "a":
            for fi5 := 0; fi5 < 10; fi5++ {
                fmt.Println(fi4, v3)
                var v6 int16 = 1000
                _ = fi5
                _ = v6
            }
        default:
            v3 = ("" >= "a")
        }
        v1 = v1
        _ = fi4
    }
    v7 := GenS1{F0: 0}
    v8 := GenS2{F0: 127}
    v9 := GenS1{F0: -800}
    v3 = ((-96 == 100) && (((120 != 22) || v3) && v3))
    v10 := GenS2{F0: int8(v1)}
    if false {
        switch "foo bar" {
        case "\n\t":
            v2["foo bar"] = ""
        }
        v11 := GenS1{F0: 0}
        if v3 {
            fmt.Println(v1, v3)
        }
        _ = v11
    }
    v1 += (^(v1))
    var v12 uint64 = 903
    switch 128 {
    case 32767:
        var v13 uint16 = 91
        _ = v13
    case 0:
        v3 = true
    }
    var v14 int32 = 1
    var v15 uint32 = 2
    v15 = 0
    v3 = true
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
