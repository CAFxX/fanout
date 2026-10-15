package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 int32 }
type GenS2 struct { F0 int; F1 uint32 }
func genFn0(p0 string) int32 {
    v4 := []int{}
    for fi5 := 0; fi5 < 19; fi5++ {
        if len(v4) > 0 {
            v4[0] = (-(fi5))
        }
        if true {
            fmt.Println(p0, fi5)
            switch fi5 {
            case 9223372036854775807:
                p0 = ("\n\t" + "hello")
            case 128:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            default:
                p0 = p0
            }
            switch "foo bar" {
            case "\n\t":
                var v6 bool = (true || true)
                _ = v6
            default:
                if false {
                    fmt.Println(fi5, fi5)
                    p0 = "\n\t"
                }
            }
        }
        _ = fi5
    }
    _ = p0
    _ = v4
    return 1
}
func genFn1(p0 uint32, p1 int32) uint16 {
    fmt.Println(p0, p0)
    p0 = 2
    _ = p0
    _ = p1
    return 32767
}

func main() {
    var v1 uint32 = 2
    var v2 uint16 = 0
    v3 := []uint32{v1, (v1 | v1), (v1 & (v1 & 2147483647))}
    var v7 string = ("hello" + "\n\t")
    var v8 int8 = -1
    v9 := GenS1{F0: ((v1 + 2147483647) != (v1 | v1)), F1: 32767, F2: -2147483648}
    if false {
        for fi10 := 0; fi10 < 7; fi10++ {
            var v11 uint32 = v1
            for fi12 := 0; fi12 < 9; fi12++ {
                if true {
                    var v13 float64 = (((3.14159 - 3.14159) / (1.5 + 1.5)) * float64(0.0))
                    _ = v13
                } else {
                    var v14 bool = false
                    v15 := GenS1{F0: ((2 < 204) && false), F1: fi10, F2: -1}
                    _ = len(v3)
                    _ = v14
                    _ = v15
                }
                v8 = (v8 * (v8 * v8))
                v16 := GenS1{F0: true, F1: 32767, F2: 1}
                _ = len(v3)
                v17 := GenS1{F0: true, F1: (fi12 << ((0) & 63)), F2: -2147483648}
                v8 = (-(v8))
                _ = fi12
                _ = v16
                _ = v17
            }
            _ = fi10
            _ = v11
        }
        fmt.Println(v8, v7)
        var v18 int16 = 0
        _ = v18
    }
    if true {
        var v19 int16 = -1
        fmt.Println(v1)
        var v20 uint8 = 117
        _ = v19
        _ = v20
    } else {
        for fi21 := 0; fi21 < 3; fi21++ {
            var v22 bool = (18446744073709551615 < 4294967295)
            v23 := GenS1{F0: ((!true) && (285 == 0)), F1: -1, F2: 42}
            v22 = true
            _ = fi21
            _ = v22
            _ = v23
        }
    }
    v24 := GenS1{F0: (true || (v7 >= "a")), F1: 4294967295, F2: -2147483648}
    fmt.Println(v2, v8)
    switch v7 {
    case "a":
        fmt.Println(v8)
    default:
        v7 = "\n\t"
    }
    for fi25 := 0; fi25 < 10; fi25++ {
        if (true && true) {
            fmt.Println(fi25, v8, v8)
        }
        _ = fi25
    }
    v26 := GenS1{F0: true, F1: 9223372036854775807, F2: genFn0((v7 + "a"))}
    var v27 uint64 = 18446744073709551615
    fmt.Println(v7)
    _ = len(v3)
    v28 := GenS1{F0: (v7 >= "hello"), F1: -35, F2: 406}
    v29 := GenS1{F0: (("\n\t" <= v7) || false), F1: 32767, F2: genFn0(v7)}
    v30 := GenS2{F0: 32767, F1: (v1 + (v1 << ((255) & 63)))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v7
    _ = v8
    _ = v9
}
