package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string }
func genFn0(p0 uint8, p1 uint64, p2 string) uint32 {
    var v6 string = ((p2 + "") + (("foo bar" + p2) + p2))
    var v7 int16 = 1
    if (!false) {
        v7 += (-(v7))
        v8 := GenS1{F0: 2}
        switch ("foo bar" + v6) {
        case "a":
            fmt.Println(p0, p1, v6)
        case "hello":
            p1 -= (p1 << ((4294967295) & 63))
        }
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 1
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]int)
    v3 := make(map[int]uint32)
    var v4 int16 = 32767
    v5 := []uint32{65535, 1, 0, 733}
    fmt.Println(v4)
    for fi9 := 0; fi9 < 6; fi9++ {
        for fi10 := 0; fi10 < 7; fi10++ {
            var v11 bool = ((!true) && false)
            _ = fi10
            _ = v11
        }
        fmt.Println(fi9, v4)
        _ = fi9
    }
    v12 := GenS2{F0: "\n\t"}
    switch 32767 {
    case 42:
        v4 += v4
    case 195:
        fmt.Println(v4)
    case 1:
        if (false && (-1 == 127)) {
            v4 = (v4 / (((v4 * (v4 ^ (^(v4))))) | 1))
            if true {
                v4 = (v4 << ((255) & 63))
                if false {
                    var v13 uint32 = 42
                    if false {
                        for fi14 := 0; fi14 < 1; fi14++ {
                            switch fi14 {
                            case 2147483647:
                                v4 += (v4 ^ (v4 + v4))
                            default:
                                v5 = append(v5, v13)
                            }
                            _ = fi14
                        }
                        v13 -= (v13 % ((42) | 1))
                    }
                    _ = v13
                }
                _ = len(v5)
            }
        } else {
            switch -1 {
            case -872:
                _ = len(v5)
            case 9223372036854775807:
                for fi15 := 0; fi15 < 4; fi15++ {
                    v4 = (v4 % ((-827) | 1))
                    _ = fi15
                }
            case 617:
                var v16 int16 = (v4 - (v4 >> ((354) & 63)))
                _ = v16
            }
        }
    }
    v4 = (-(v4))
    var v17 uint64 = 4294967295
    v18 := GenS2{F0: (("hello" + ("foo bar" + (("foo bar" + "foo bar") + ("a" + "hello")))) + ("hello" + ""))}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
