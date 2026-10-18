package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 uint8) bool {
    var v4 uint32 = 196
    v5 := GenS1{F0: -32768}
    v6 := []int{889, 128, -1, 42}
    var v7 uint = 591
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return (-638 == 0)
}
func genFn1(p0 bool) uint64 {
    v8 := GenS1{F0: 32767}
    fmt.Println(p0)
    v9 := []int{-2147483648, 42, 0}
    for fi10 := 0; fi10 < 16; fi10++ {
        var v11 string = "hello"
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        var v13 int = fi12
        v14 := GenS1{F0: 1}
        _ = fi12
        _ = v13
        _ = v14
    }
    _ = p0
    _ = v8
    _ = v9
    return 4294967295
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{("" + ""), "\n\t", "", ("hello" + ("" + "\n\t"))}
    v3 := make(map[int]int)
    var v15 uint32 = 1
    for fi16 := 0; fi16 < 18; fi16++ {
        var v17 string = (("" + (("a" + ((("" + (("foo bar" + "\n\t") + ("a" + ""))) + ("foo bar" + "hello")) + "a")) + "hello")) + "a")
        v15 = (v15 | (v15 << ((460) & 63)))
        _ = fi16
        _ = v17
    }
    v15 -= 0
    fmt.Println(v15)
    v2 = append(v2, "foo bar")
    fmt.Println(v15)
    for fi18 := 0; fi18 < 6; fi18++ {
        if false {
            v15 -= v15
            var v19 float64 = ((((1e10 * ((3.14159 - 1.5) + 0.1)) / (3.14159 + 1e10)) - (float64(1e10) + 1e10)) * (float64(0.1) / 1.0))
            for fi20 := 0; fi20 < 2; fi20++ {
                v19 = ((v19 / (v19 - v19)) * v19)
                v19 = (v19 - float64((0.1 - 0.0)))
                _ = fi20
            }
            _ = v19
        }
        _ = fi18
    }
    v15 = (v15 / (((v15 + 42)) | 1))
    v15 = v15
    v15 = v15
    var v21 uint16 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v3
}
