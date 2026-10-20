package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 bool; F1 int32; F2 int32 }
func genFn0(p0 int64) int8 {
    if false {
        var v5 uint64 = 2
        var v6 string = "a"
        if (!(!false)) {
            for fi7 := 0; fi7 < 14; fi7++ {
                var v8 bool = (151 == 42)
                var v9 uint8 = 1
                for fi10 := 0; fi10 < 14; fi10++ {
                    if v8 {
                        var v11 int = 9223372036854775807
                        _ = v11
                    }
                    _ = fi10
                }
                _ = fi7
                _ = v8
                _ = v9
            }
            for fi12 := 0; fi12 < 12; fi12++ {
                var v13 uint32 = 2147483648
                _ = fi12
                _ = v13
            }
        } else {
            var v14 string = ("a" + v6)
            v15 := []int{0}
            if len(v15) > 0 {
                v15[0] = 128
            }
            _ = v14
            _ = v15
        }
        _ = v5
        _ = v6
    }
    var v16 uint16 = 0
    v16 = (v16 >> ((516) & 63))
    for fi17 := 0; fi17 < 4; fi17++ {
        var v18 uint32 = 2
        _ = fi17
        _ = v18
    }
    _ = p0
    _ = v16
    return -41
}
func genFn1(p0 int) uint16 {
    v19 := GenS1{F0: 2147483647, F1: 1}
    v20 := []int{32767}
    var v21 string = "\n\t"
    var v22 uint64 = 0
    _ = p0
    _ = v19
    _ = v20
    _ = v21
    _ = v22
    return 32767
}

func main() {
    var v1 int = 34
    var v2 string = "a"
    var v3 string = (v2 + v2)
    v4 := make(map[string]int)
    if (255 < 117) {
        var v23 int8 = -38
        var v24 uint8 = 255
        var v25 uint16 = 1
        _ = v23
        _ = v24
        _ = v25
    }
    v26 := []int{(v1 ^ v1), (-(v1))}
    v4["hello"] = v1
    v2 = ((v2 + "\n\t") + "hello")
    v27 := GenS2{F0: false, F1: 0, F2: 42}
    v2 = ("hello" + v2)
    v2 = (("hello" + v3) + v3)
    for fi28 := 0; fi28 < 13; fi28++ {
        for fi29 := 0; fi29 < 9; fi29++ {
            v2 = v3
            _ = fi29
        }
        v2 = "hello"
        _ = fi28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
