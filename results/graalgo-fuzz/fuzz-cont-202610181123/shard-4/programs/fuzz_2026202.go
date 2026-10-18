package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 string = ("hello" + "foo bar")
    v2 := []int{374}
    var v3 float64 = 1e10
    var v4 uint = 255
    v5 := []string{v1, (v1 + v1), (v1 + "hello")}
    for fi6 := 0; fi6 < 9; fi6++ {
        v3 = v3
        _ = fi6
    }
    if (!((v4 + v4) < 255)) {
        v4 = v4
        v7 := GenS1{F0: 32767}
        _ = v7
    }
    var v8 string = "a"
    var v9 int64 = -769
    if (1 == 2) {
        v10 := GenS2{F0: 128}
        var v11 bool = true
        var v12 uint = v4
        _ = v10
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 3; fi13++ {
        var v14 int16 = 0
        switch fi13 {
        case -473:
            var v15 uint16 = 32767
            _ = v15
        case 32767:
            if len(v5) > 0 {
                v5[0] = "foo bar"
            }
        default:
            var v16 uint32 = 240
            _ = v16
        }
        fmt.Println(v9, fi13, v4)
        _ = fi13
        _ = v14
    }
    fmt.Println(v1, v4)
    v9 = 42
    var v17 uint16 = 2
    v17 = 255
    for fi18 := 0; fi18 < 18; fi18++ {
        var v19 int8 = 16
        _ = fi18
        _ = v19
    }
    fmt.Println(v9)
    for fi20 := 0; fi20 < 5; fi20++ {
        fmt.Println(v17)
        v8 = ((v1 + "foo bar") + v8)
        v8 = v1
        _ = fi20
    }
    var v21 int16 = -32768
    v8 = (v1 + "a")
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
