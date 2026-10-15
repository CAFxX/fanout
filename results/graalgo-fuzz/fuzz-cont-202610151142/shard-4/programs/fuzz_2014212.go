package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 uint64 }
type GenS2 struct { F0 int }

func main() {
    v1 := []uint32{}
    v2 := make(map[string]bool)
    var v3 uint = 18446744073709551615
    v4 := make(map[string]uint32)
    v3 += 2
    v5 := v4[(("" + "") + "hello")]
    v6 := GenS1{F0: -604, F1: ("hello" < "foo bar"), F2: 2}
    var v7 int8 = 127
    fmt.Println(v5, v5)
    var v8 uint16 = 32767
    fmt.Println(v7)
    v5 = v5
    fmt.Println(v7)
    v9 := GenS1{F0: 2147483647, F1: (341 != -9223372036854775808), F2: 2}
    if false {
        for fi10 := 0; fi10 < 8; fi10++ {
            v11 := GenS1{F0: -1, F1: (true || (!("\n\t" < "hello"))), F2: 253}
            v5 = v5
            for fi12 := 0; fi12 < 11; fi12++ {
                fmt.Println(v5)
                if (-1 > -4) {
                    var v13 uint32 = 2
                    _ = v13
                }
                _ = fi12
            }
            _ = fi10
            _ = v11
        }
    }
    delete(v4, (("a" + ("a" + "\n\t")) + (("hello" + (("foo bar" + "\n\t") + (("foo bar" + "foo bar") + "hello"))) + "\n\t")))
    var v14 int = 4294967295
    v15 := GenS2{F0: -938}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
