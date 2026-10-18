package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 int8 }
type GenS2 struct { F0 uint64; F1 uint }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint16 = 1
    v3 := make(map[uint32]bool)
    v4 := v3[4294967295]
    v2 += (v2 - v2)
    v4 = v4
    for fi5 := 0; fi5 < 9; fi5++ {
        if (42 <= fi5) {
            v4 = (!(!(0 != (^(v2)))))
        }
        _ = fi5
    }
    if v4 {
        var v6 int16 = 32767
        v7 := []int{42, 32767}
        _ = v6
        _ = v7
    }
    v8 := []int{0, -1, -1}
    v9 := GenS1{F0: 18446744073709551615, F1: 0, F2: 100}
    fmt.Println(v2, v2)
    switch ((("a" + "a") + (("foo bar" + "a") + "hello")) + (("" + "a") + ("hello" + ("\n\t" + "foo bar")))) {
    case "hello":
        for fi10 := 0; fi10 < 10; fi10++ {
            if (9223372036854775807 >= 281) {
                v2 = 32767
                if v4 {
                    _ = len(v3)
                }
                v11 := GenS1{F0: 0, F1: 2147483647, F2: -54}
                _ = v11
            }
            _ = fi10
        }
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        v13 := GenS1{F0: 4294967295, F1: 1, F2: 127}
        _ = fi12
        _ = v13
    }
    var v14 int64 = 9223372036854775807
    var v15 uint32 = 4294967295
    delete(v1, "hello")
    v16 := GenS2{F0: 875, F1: 65535}
    switch "foo bar" {
    case "a":
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
    }
    v17 := GenS2{F0: 42, F1: 1000}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
