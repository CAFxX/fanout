package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 bool; F2 int16 }

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{0, 1}
    var v3 int8 = 100
    var v4 uint64 = 2
    for fi5 := 0; fi5 < 6; fi5++ {
        var v6 int = -2147483648
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 uint16 = 1
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v9 := GenS1{F0: uint16(v4)}
        _ = fi7
        _ = v8
        _ = v9
    }
    delete(v1, 4294967295)
    switch "foo bar" {
    case "":
        v4 -= 18446744073709551615
    case "\n\t":
        v3 -= 1
    default:
        v3 = 42
    }
    if (!(true && false)) {
        fmt.Println(v4, v4)
    }
    fmt.Println(v3, v4)
    for fi10 := 0; fi10 < 1; fi10++ {
        var v11 uint32 = 2
        _ = fi10
        _ = v11
    }
    var v12 int32 = 449
    switch "\n\t" {
    case "hello":
        v12 += (v12 * (-(v12)))
    default:
        v13 := GenS1{F0: 580}
        _ = v13
    }
    switch "hello" {
    case "":
        v2 = append(v2, 0)
    case "\n\t":
        if len(v2) > 0 {
            v2[0] = 0
        }
    default:
        if ("\n\t" == "\n\t") {
            v14 := GenS2{F0: 0, F1: true, F2: -1}
            var v15 bool = (true || ("a" <= (("" + "hello") + "\n\t")))
            v3 = -67
            _ = v14
            _ = v15
        } else {
            v16 := GenS1{F0: 97}
            var v17 uint32 = 2147483647
            _ = v16
            _ = v17
        }
    }
    switch ("hello" + "a") {
    case "":
        var v18 bool = true
        _ = v18
    }
    for fi19 := 0; fi19 < 18; fi19++ {
        delete(v1, 2147483648)
        fmt.Println(fi19, v4)
        if len(v2) > 0 {
            v2[0] = (fi19 >> ((981) & 63))
        }
        _ = fi19
    }
    for fi20 := 0; fi20 < 19; fi20++ {
        for fi21 := 0; fi21 < 13; fi21++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v22 int8 = (v3 % (((v3 * (v3 % ((v3) | 1)))) | 1))
            _ = fi21
            _ = v22
        }
        delete(v1, 2147483648)
        _ = fi20
    }
    v23 := GenS2{F0: 1000, F1: (!true), F2: -1}
    var v24 int32 = (v12 | v12)
    var v25 int = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
