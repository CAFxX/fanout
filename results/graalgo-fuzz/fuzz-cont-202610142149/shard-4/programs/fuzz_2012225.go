package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 uint }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]uint32)
    v3 := make(map[int]string)
    var v4 int16 = 32767
    v5 := make(map[uint32]bool)
    v6 := make(map[uint32]string)
    v4 -= v4
    v7 := []int{32767}
    for fi8 := 0; fi8 < 5; fi8++ {
        delete(v1, ("\n\t" + "hello"))
        var v9 int8 = 42
        _ = fi8
        _ = v9
    }
    if ("\n\t" != (("foo bar" + "hello") + "foo bar")) {
        var v10 bool = false
        var v11 string = (("" + "hello") + (("foo bar" + "hello") + (("foo bar" + "hello") + ("foo bar" + ("\n\t" + "foo bar")))))
        _ = v10
        _ = v11
    }
    if false {
        var v12 int16 = (v4 * v4)
        v12 += -32768
        v12 -= v4
        _ = v12
    }
    if ((-2.25 >= 1e10) && false) {
        var v13 uint32 = 4294967295
        for fi14 := 0; fi14 < 11; fi14++ {
            if false {
                v15 := GenS1{F0: (v4 / (((-(v4))) | 1)), F1: -1, F2: 18446744073709551615}
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                var v16 bool = false
                _ = v15
                _ = v16
            }
            _ = len(v6)
            _ = fi14
        }
        var v17 uint = 42
        _ = v13
        _ = v17
    } else {
        v4 += (v4 - -365)
        for fi18 := 0; fi18 < 7; fi18++ {
            v4 -= (^(v4))
            v4 = 55
            _ = fi18
        }
    }
    var v19 bool = (806 != 0)
    v4 -= 949
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
