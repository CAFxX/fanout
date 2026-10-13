package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 uint64; F2 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64; F1 int8; F2 string }
func genFn0(p0 int32) int32 {
    p0 = p0
    for fi7 := 0; fi7 < 16; fi7++ {
        v8 := []int{fi7, fi7}
        v9 := GenS1{F0: 4294967295}
        v10 := GenS1{F0: 42}
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    if (!false) {
        switch "hello" {
        case "foo bar":
            v11 := GenS2{F0: 255, F1: 18446744073709551615, F2: 2}
            _ = v11
        case "a":
            var v12 float32 = 0.1
            _ = v12
        }
        var v13 int16 = -773
        var v14 uint16 = 32767
        _ = v13
        _ = v14
    }
    v15 := GenS1{F0: 0}
    _ = p0
    _ = v15
    return p0
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint64 = 1
    var v3 bool = (-850 >= 87)
    var v4 uint16 = 65535
    v5 := make(map[string]int)
    v6 := make(map[uint32]int)
    v16 := []int{0, 634, 80}
    v17 := GenS3{F0: -2147483648}
    _ = len(v16)
    fmt.Println(v4, v2)
    v18 := GenS4{F0: (v2 - (^(v2))), F1: 127, F2: (("foo bar" + "foo bar") + (("" + "\n\t") + "foo bar"))}
    var v19 string = ""
    _ = len(v6)
    v19 = ""
    if v3 {
        var v20 uint = 446
        _ = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
