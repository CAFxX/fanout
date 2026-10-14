package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16 }

func main() {
    var v1 uint32 = 2147483648
    var v2 uint = 270
    v3 := make(map[uint32]int)
    var v4 uint8 = 201
    v5 := make(map[string]uint32)
    v6 := make(map[int]int)
    switch int(v1) {
    case -2147483648:
        v7 := GenS1{F0: -106}
        _ = v7
    case 128:
        v8 := GenS1{F0: -702}
        _ = v8
    case 2147483647:
        var v9 int64 = 521
        _ = v9
    }
    switch 9223372036854775807 {
    case 696:
        v4 = v4
    case 4294967295:
        v2 -= (^(v2))
    case 2147483647:
        var v10 uint = v2
        _ = v10
    }
    if (true || (-1 > -1)) {
        v1 = v1
        fmt.Println(v1, v4)
        var v11 uint64 = 0
        _ = v11
    } else {
        for fi12 := 0; fi12 < 11; fi12++ {
            v13 := GenS1{F0: 42}
            v4 = (v4 | v4)
            _ = fi12
            _ = v13
        }
        var v14 float32 = (1e10 * ((-2.25 + 1.5) + 1e10))
        _ = v14
    }
    if true {
        v15 := GenS2{F0: -1}
        _ = v15
    }
    var v16 int8 = -128
    v17 := []int{1, 2147483647}
    for fi18 := 0; fi18 < 12; fi18++ {
        v4 += v4
        _ = fi18
    }
    for fi19 := 0; fi19 < 14; fi19++ {
        v20 := GenS1{F0: 358}
        var v21 uint64 = 2
        _ = fi19
        _ = v20
        _ = v21
    }
    switch 42 {
    case 0:
        var v22 string = ""
        _ = v22
    }
    var v23 int64 = 0
    v23 = v23
    if len(v17) > 0 {
        v17[0] = 795
    }
    switch -159 {
    case 32767:
        var v24 uint8 = (v4 ^ 67)
        _ = v24
    case 9223372036854775807:
        var v25 bool = false
        _ = v25
    }
    switch "foo bar" {
    case "hello":
        for fi26 := 0; fi26 < 18; fi26++ {
            v27 := GenS2{F0: 0}
            if false {
                var v28 uint = (v2 - v2)
                _ = v28
            } else {
                v6[-9223372036854775808] = (^(fi26))
                v6[fi26] = fi26
                for fi29 := 0; fi29 < 3; fi29++ {
                    v30 := GenS1{F0: v23}
                    var v31 uint64 = 1
                    _ = fi29
                    _ = v30
                    _ = v31
                }
            }
            for fi32 := 0; fi32 < 3; fi32++ {
                var v33 float32 = 3.14159
                var v34 string = ((("hello" + "foo bar") + ((("foo bar" + "a") + "foo bar") + "\n\t")) + "foo bar")
                _ = fi32
                _ = v33
                _ = v34
            }
            _ = fi26
            _ = v27
        }
    case "":
        var v35 bool = (1e10 <= 0.0)
        _ = v35
    case "a":
        var v36 uint16 = 2
        _ = v36
    }
    v37 := GenS1{F0: 9223372036854775807}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
