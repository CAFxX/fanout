package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string; F1 uint32 }

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := make(map[int]uint32)
    v3 := []uint32{}
    var v4 uint = 662
    var v5 uint8 = 127
    v5 = 160
    v3 = append(v3, 1)
    var v6 uint16 = 1
    if true {
        v1 = (v1 << ((v4) & 63))
        for fi7 := 0; fi7 < 7; fi7++ {
            var v8 uint = v4
            v3 = append(v3, 617)
            var v9 uint = v8
            _ = fi7
            _ = v8
            _ = v9
        }
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v4 = (v4 - (v4 - v4))
        v6 = v6
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        for fi11 := 0; fi11 < 19; fi11++ {
            var v12 uint64 = v1
            if true {
                _ = len(v2)
                switch "foo bar" {
                case "":
                    switch "hello" {
                    case "foo bar":
                        var v13 bool = false
                        _ = v13
                    case "hello":
                        switch ("\n\t" + "hello") {
                        case "":
                            v5 = 111
                        case "\n\t":
                            var v14 uint16 = v6
                            _ = v14
                        case "foo bar":
                            delete(v2, (fi11 * 482))
                        default:
                            if ((0 > -1) && true) {
                                fmt.Println(v5)
                            }
                        }
                    default:
                        v15 := GenS1{F0: 319}
                        _ = v15
                    }
                case "a":
                    v16 := GenS1{F0: 1}
                    _ = v16
                }
            } else {
                v17 := GenS2{F0: -27}
                delete(v2, (fi11 ^ (fi10 + (fi10 * fi10))))
                _ = v17
            }
            switch ("hello" + "a") {
            case "hello":
                _ = len(v2)
            default:
                v18 := GenS3{F0: ("foo bar" + ""), F1: 65535}
                _ = v18
            }
            _ = fi11
            _ = v12
        }
        v5 += (v5 & 0)
        fmt.Println(v6, fi10, v4)
        _ = fi10
    }
    v6 = 1
    fmt.Println(v1, v6)
    if (!(!(!(647 <= -300)))) {
        fmt.Println(v4, v4)
        var v19 uint32 = 559
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
