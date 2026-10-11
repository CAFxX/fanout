package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 string, p1 int16, p2 uint8) int {
    p1 += 42
    var v5 string = "\n\t"
    var v6 float64 = -2.25
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 32768
}
func genFn1() bool {
    switch -2147483648 {
    case 1:
        var v7 bool = (341 >= 1)
        _ = v7
    case 127:
        var v8 uint16 = 783
        _ = v8
    case -496:
        v9 := GenS1{F0: 128}
        _ = v9
    }
    v10 := GenS1{F0: 42}
    for fi11 := 0; fi11 < 15; fi11++ {
        var v12 uint64 = 127
        _ = fi11
        _ = v12
    }
    _ = v10
    return false
}

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 42
    var v3 bool = true
    v4 := make(map[string]uint32)
    v13 := []int{}
    v14 := GenS1{F0: (v2 + (^(v2)))}
    switch (("foo bar" + "a") + "foo bar") {
    case "foo bar":
        delete(v1, 1)
    }
    for fi15 := 0; fi15 < 13; fi15++ {
        v16 := v4["foo bar"]
        _ = fi15
        _ = v16
    }
    v3 = v3
    fmt.Println(v2)
    v13 = append(v13, 239)
    v17 := GenS1{F0: 11}
    v2 = 128
    fmt.Println(v3)
    switch 42 {
    case 424:
        var v18 bool = true
        _ = v18
    case -809:
        v3 = v3
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
