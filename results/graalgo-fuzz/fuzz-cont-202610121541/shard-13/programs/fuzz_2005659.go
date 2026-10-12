package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 uint8 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint8, p1 int64, p2 uint) uint64 {
    p2 = p2
    v5 := GenS1{F0: "", F1: 127, F2: (^(p0))}
    var v6 float32 = 1e10
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 1
}

func main() {
    var v1 int32 = 42
    v2 := []string{}
    var v3 uint32 = 195
    v4 := []uint32{2147483647, (v3 * (v3 / (((v3 << ((2) & 63))) | 1))), (v3 - v3), (v3 ^ (v3 / ((v3) | 1)))}
    v1 = v1
    fmt.Println(v1, v3)
    var v7 uint16 = 0
    var v8 int8 = 1
    var v9 uint = 196
    switch 439 {
    case 1:
        for fi10 := 0; fi10 < 12; fi10++ {
            v2 = append(v2, "foo bar")
            switch "hello" {
            case "foo bar":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case "hello":
                if len(v2) > 0 {
                    v2[0] = "foo bar"
                }
            case "a":
                v1 -= v1
            }
            v3 += (v3 & uint32(v9))
            _ = fi10
        }
    case 768:
        var v11 int64 = 42
        _ = v11
    case 4294967295:
        if (-2.25 > 1e10) {
            if false {
                v12 := GenS2{F0: uint8(v3)}
                for fi13 := 0; fi13 < 20; fi13++ {
                    fmt.Println(v8, v1, v3)
                    _ = fi13
                }
                _ = v12
            }
        } else {
            v14 := GenS2{F0: 128}
            v9 += v9
            _ = v14
        }
    default:
        var v15 uint = 42
        _ = v15
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        switch "a" {
        case "foo bar":
            v3 += (v3 << ((v9) & 63))
        default:
            v7 -= 65535
        }
        _ = fi16
    }
    switch 42 {
    case -1:
        var v17 uint = (v9 | (v9 >> ((v9) & 63)))
        _ = v17
    }
    for fi18 := 0; fi18 < 16; fi18++ {
        v2 = append(v2, "a")
        v9 = 0
        _ = fi18
    }
    v7 = (v7 - v7)
    v1 -= (v1 * (-(v1)))
    var v19 uint16 = (v7 << ((65535) & 63))
    v1 = v1
    v20 := GenS1{F0: ("" + ("\n\t" + "")), F1: (v8 & (-(v8))), F2: 255}
    var v21 uint = 18446744073709551615
    v2 = append(v2, ("hello" + "hello"))
    fmt.Println(v3, v19, v7)
    v7 = v7
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
