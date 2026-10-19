package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 int16 }
func genFn0(p0 int8, p1 int16) uint8 {
    switch ("" + "") {
    case "":
        p0 = (p0 << ((65535) & 63))
    case "foo bar":
        fmt.Println(p1, p1)
    case "a":
        p0 = p0
    default:
        if true {
            for fi4 := 0; fi4 < 6; fi4++ {
                var v5 uint16 = 255
                _ = fi4
                _ = v5
            }
        } else {
            for fi6 := 0; fi6 < 20; fi6++ {
                var v7 uint8 = 135
                if (906 < 0) {
                    fmt.Println(p0, p0, p0)
                }
                _ = fi6
                _ = v7
            }
        }
    }
    for fi8 := 0; fi8 < 7; fi8++ {
        fmt.Println(fi8)
        var v9 uint32 = 4294967295
        v10 := []int{}
        _ = fi8
        _ = v10
        _ = v9
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 string = ((("a" + "foo bar") + "a") + "hello")
    var v2 string = (v1 + (v1 + "hello"))
    var v3 float64 = (float64((-2.25 - 0.1)) - -2.25)
    if ((v2 == (v2 + "\n\t")) && false) {
        v11 := GenS1{F0: 4294967295, F1: 42}
        var v12 int = 4294967295
        _ = v11
        _ = v12
    } else {
        for fi13 := 0; fi13 < 13; fi13++ {
            var v14 int16 = 692
            v15 := []int{}
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    v16 := []int{-705, 444, 32767, 4294967295}
    switch 1 {
    case 32767:
        if false {
            fmt.Println(v3, v2, v2)
            v16 = append(v16, 2147483647)
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
        }
    case 792:
        v1 = "a"
    default:
        for fi17 := 0; fi17 < 5; fi17++ {
            v3 += v3
            var v18 int32 = -1
            _ = fi17
            _ = v18
        }
    }
    var v19 uint32 = 822
    if false {
        if false {
            var v20 uint32 = 42
            var v21 uint = 4294967295
            v1 = v2
            _ = v20
            _ = v21
        }
    } else {
        v16 = append(v16, 127)
        var v22 int = 42
        v23 := GenS1{F0: v19, F1: -1}
        _ = v22
        _ = v23
    }
    switch -673 {
    case -1:
        v24 := GenS2{F0: 318}
        _ = v24
    default:
        var v25 uint8 = 1
        _ = v25
    }
    var v26 int16 = 471
    v26 = -1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v26
    _ = v3
}
