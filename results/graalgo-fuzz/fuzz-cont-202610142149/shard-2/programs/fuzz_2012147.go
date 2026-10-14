package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]int)
    v3 := make(map[uint32]bool)
    switch -1 {
    case -960:
        v4 := GenS1{F0: 361}
        _ = v4
    case -2147483648:
        var v5 uint32 = 851
        _ = v5
    }
    switch "foo bar" {
    case "a":
        var v6 uint = 18446744073709551615
        _ = v6
    case "":
        var v7 int64 = -9223372036854775808
        _ = v7
    case "hello":
        v8 := GenS2{F0: -1}
        _ = v8
    }
    switch 4294967295 {
    case 0:
        v9 := GenS1{F0: 557}
        _ = v9
    default:
        var v10 uint8 = 255
        _ = v10
    }
    var v11 string = ""
    if (!((((v11 + v11) + v11) >= v11) && true)) {
        var v12 int64 = 854
        _ = v12
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        v14 := []int{fi13}
        _ = fi13
        _ = v14
    }
    v15 := GenS2{F0: 1000}
    if false {
        switch 578 {
        case 9223372036854775807:
            var v16 string = "foo bar"
            _ = v16
        case 1:
            v11 = "foo bar"
        }
    } else {
        var v17 float32 = float32(0.0)
        var v18 uint = 1
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 11; fi19++ {
        for fi20 := 0; fi20 < 11; fi20++ {
            v11 = "foo bar"
            var v21 uint32 = 0
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    for fi22 := 0; fi22 < 9; fi22++ {
        v11 = "foo bar"
        _ = fi22
    }
    var v23 uint32 = 0
    var v24 uint16 = 42
    v23 -= v23
    var v25 uint64 = uint64(v24)
    for fi26 := 0; fi26 < 19; fi26++ {
        var v27 uint64 = (^(v25))
        for fi28 := 0; fi28 < 12; fi28++ {
            v29 := v1[(v11 + (v11 + ""))]
            delete(v2, (^(fi26)))
            var v30 uint = 65535
            _ = fi28
            _ = v29
            _ = v30
        }
        v25 = v25
        _ = fi26
        _ = v27
    }
    v31 := GenS1{F0: 9223372036854775807}
    v32 := GenS2{F0: 1000}
    v24 = v24
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v31
    _ = v32
}
