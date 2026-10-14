package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 bool; F2 bool }

func main() {
    v1 := []string{("" + ("foo bar" + ("hello" + "foo bar")))}
    var v2 int16 = 478
    v3 := []uint32{1}
    var v4 uint16 = 32767
    var v5 string = "foo bar"
    var v6 int16 = v2
    var v7 int = -1
    var v8 int16 = v2
    v3 = append(v3, 501)
    v5 = (v5 + "\n\t")
    for fi9 := 0; fi9 < 1; fi9++ {
        var v10 float64 = 0.1
        _ = fi9
        _ = v10
    }
    if true {
        var v11 float32 = float32(float64(1e10))
        _ = v11
    }
    if true {
        var v12 string = (v5 + "a")
        v13 := GenS1{F0: 0}
        for fi14 := 0; fi14 < 8; fi14++ {
            var v15 uint8 = 0
            if (717 >= -576) {
                var v16 uint16 = v4
                _ = v16
            }
            _ = fi14
            _ = v15
        }
        _ = v12
        _ = v13
    } else {
        var v17 int8 = 1
        fmt.Println(v5, v2)
        _ = v17
    }
    v18 := GenS1{F0: 2147483647}
    switch (v5 + v5) {
    case "foo bar":
        v19 := GenS1{F0: 42}
        _ = v19
    }
    v20 := GenS2{F0: false, F1: (355 >= 641), F2: (!true)}
    var v21 uint64 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
