package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint8 }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 int; F1 string; F2 int64 }
type GenS4 struct { F0 int32; F1 uint16; F2 bool }
func genFn0(p0 uint16) bool {
    var v4 int = 1
    v5 := GenS1{F0: p0, F1: -513, F2: 1}
    _ = p0
    _ = v4
    _ = v5
    return true
}

func main() {
    v1 := []int{-333, 9223372036854775807}
    v2 := []uint32{}
    v3 := make(map[uint32]string)
    var v6 string = "a"
    var v7 uint16 = 0
    if true {
        v6 = ("" + ("a" + ("foo bar" + "\n\t")))
        v8 := GenS2{F0: 42, F1: uint64(v7)}
        for fi9 := 0; fi9 < 10; fi9++ {
            v6 = "hello"
            v6 = ((v6 + "hello") + "hello")
            _ = fi9
        }
        _ = v8
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        for fi11 := 0; fi11 < 11; fi11++ {
            if false {
                v12 := GenS3{F0: fi11, F1: (("a" + v6) + v6), F2: 0}
                _ = v12
            } else {
                v13 := GenS3{F0: fi10, F1: ("foo bar" + ("\n\t" + "hello")), F2: -9223372036854775808}
                delete(v3, 885)
                var v14 uint16 = 631
                _ = v13
                _ = v14
            }
            _ = fi11
        }
        if true {
            v15 := GenS4{F0: 2147483647, F1: v7, F2: (18446744073709551615 > 18446744073709551615)}
            _ = v15
        }
        v6 = ""
        _ = fi10
    }
    var v16 bool = true
    var v17 int32 = 687
    _ = len(v3)
    v18 := GenS4{F0: v17, F1: (v7 << ((363) & 63)), F2: v16}
    v7 += 65535
    fmt.Println(v6)
    v19 := GenS4{F0: (v17 >> ((627) & 63)), F1: (v7 ^ v7), F2: false}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
