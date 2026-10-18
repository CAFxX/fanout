package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 uint64 }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 int; F1 int }

func main() {
    var v1 float32 = (1e10 + (1e10 / (-2.25 - (((1.5 + 0.1) - (1.5 / 3.14159)) - -2.25))))
    v2 := make(map[string]int)
    var v3 int16 = -32768
    var v4 uint = 2
    var v5 float32 = v1
    var v6 int = 1
    for fi7 := 0; fi7 < 17; fi7++ {
        v5 = 0.1
        v8 := []int{(v6 % (((^(fi7))) | 1))}
        var v9 uint64 = 688
        _ = fi7
        _ = v8
        _ = v9
    }
    fmt.Println(v3, v6, v1)
    var v10 uint8 = 128
    v3 = (v3 | -562)
    switch ("foo bar" + ("foo bar" + "\n\t")) {
    case "a":
        var v11 uint64 = 2
        _ = v11
    }
    v12 := GenS1{F0: 2, F1: 1, F2: 78}
    v13 := GenS1{F0: 42, F1: 42, F2: 398}
    fmt.Println(v3)
    switch ("" + "foo bar") {
    case "":
        var v14 int = 32768
        _ = v14
    case "foo bar":
        v15 := v2[((("\n\t" + "foo bar") + "foo bar") + "a")]
        _ = v15
    case "a":
        v1 -= -2.25
    }
    for fi16 := 0; fi16 < 17; fi16++ {
        var v17 float32 = v1
        delete(v2, ("hello" + "foo bar"))
        v5 = (v5 - 1e10)
        _ = fi16
        _ = v17
    }
    fmt.Println(v6)
    for fi18 := 0; fi18 < 1; fi18++ {
        v5 = v1
        _ = fi18
    }
    switch "\n\t" {
    case "":
        var v19 int = int(v10)
        _ = v19
    case "hello":
        v20 := GenS2{F0: 65535, F1: 42}
        _ = v20
    }
    v21 := GenS3{F0: 4294967295, F1: v6}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
