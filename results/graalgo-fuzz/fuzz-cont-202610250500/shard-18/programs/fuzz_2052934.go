package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 uint32; F2 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8; F1 uint8 }
type GenS5 struct { F0 int8; F1 int; F2 string }
func genFn0(p0 int16) uint {
    v5 := []int{0, 127}
    v6 := GenS1{F0: 22}
    if (1 != 85) {
        p0 += (^(p0))
    } else {
        var v7 uint32 = 2
        var v8 int = -1
        _ = v7
        _ = v8
    }
    p0 += 0
    _ = p0
    _ = v5
    _ = v6
    return 4294967295
}

func main() {
    var v1 int8 = 1
    v2 := []uint32{}
    v3 := make(map[uint32]string)
    v4 := []string{"\n\t", "foo bar", ("a" + "foo bar"), "a"}
    v1 = 1
    fmt.Println(v1)
    for fi9 := 0; fi9 < 19; fi9++ {
        fmt.Println(v1, fi9)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        for fi10 := 0; fi10 < 19; fi10++ {
            v11 := GenS2{F0: 32767, F1: 561, F2: -136}
            var v12 uint32 = 2147483648
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi9
    }
    fmt.Println(v1)
    for fi13 := 0; fi13 < 14; fi13++ {
        v1 += (v1 / ((v1) | 1))
        _ = fi13
    }
    var v14 bool = false
    v15 := GenS3{F0: ("\n\t" + "\n\t")}
    fmt.Println(v1)
    var v16 string = "hello"
    v17 := GenS4{F0: 121, F1: 42}
    v16 = "a"
    v18 := GenS5{F0: (v1 & v1), F1: 92, F2: "hello"}
    fmt.Println(v16, v16, v14)
    for fi19 := 0; fi19 < 4; fi19++ {
        fmt.Println(fi19)
        var v20 int32 = -765
        v16 = v16
        _ = fi19
        _ = v20
    }
    v16 = ((("a" + v16) + v16) + v16)
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
