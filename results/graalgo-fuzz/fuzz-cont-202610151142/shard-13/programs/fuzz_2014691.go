package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 uint16 = 255
    var v2 int32 = 1
    var v3 int16 = -1
    v1 = v1
    if true {
        v4 := GenS1{F0: 2}
        v5 := GenS1{F0: 827}
        for fi6 := 0; fi6 < 7; fi6++ {
            var v7 int32 = (-(v2))
            v8 := []int{(fi6 + fi6)}
            v9 := GenS1{F0: 18446744073709551615}
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
        _ = v4
        _ = v5
    }
    var v10 uint16 = v1
    if false {
        for fi11 := 0; fi11 < 10; fi11++ {
            var v12 string = ""
            _ = fi11
            _ = v12
        }
    } else {
        v13 := GenS1{F0: 4294967295}
        _ = v13
    }
    var v14 int = -258
    switch "\n\t" {
    case "\n\t":
        v15 := GenS2{F0: 559}
        _ = v15
    case "":
        var v16 int32 = 684
        _ = v16
    case "a":
        v14 -= (-(v14))
    default:
        var v17 int16 = v3
        _ = v17
    }
    fmt.Println(v3, v3)
    fmt.Println(v1)
    fmt.Println(v10, v10)
    var v18 bool = (18446744073709551615 == 0)
    var v19 int8 = 1
    for fi20 := 0; fi20 < 19; fi20++ {
        v21 := []int{32767, v14}
        var v22 uint64 = 4294967295
        var v23 uint = 0
        _ = fi20
        _ = v21
        _ = v22
        _ = v23
    }
    v24 := []int{v14}
    if v18 {
        v14 -= v14
    } else {
        for fi25 := 0; fi25 < 5; fi25++ {
            v2 -= 0
            _ = fi25
        }
    }
    _ = len(v24)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v3
}
