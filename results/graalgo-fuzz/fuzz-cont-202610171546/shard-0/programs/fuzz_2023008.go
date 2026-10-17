package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 string) int16 {
    p0 = ("\n\t" + ("\n\t" + p0))
    for fi4 := 0; fi4 < 17; fi4++ {
        var v5 int16 = 1000
        _ = fi4
        _ = v5
    }
    if (!(0 > 42)) {
        var v6 float32 = -2.25
        var v7 float32 = (v6 - v6)
        v7 += 3.14159
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: 943, F1: 679}
    for fi9 := 0; fi9 < 17; fi9++ {
        for fi10 := 0; fi10 < 16; fi10++ {
            var v11 uint = 979
            _ = fi10
            _ = v11
        }
        p0 = (p0 + p0)
        _ = fi9
    }
    _ = p0
    _ = v8
    return 0
}
func genFn1(p0 string, p1 bool, p2 uint8) int8 {
    p0 = "a"
    p2 = 1
    for fi12 := 0; fi12 < 3; fi12++ {
        var v13 string = p0
        v14 := GenS2{F0: (p2 + p2), F1: 783}
        _ = fi12
        _ = v13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    return -1
}

func main() {
    var v1 float32 = ((1.5 / 0.1) * -2.25)
    v2 := []uint32{}
    v3 := make(map[string]uint32)
    var v15 uint8 = 127
    var v16 int16 = -1
    for fi17 := 0; fi17 < 19; fi17++ {
        v18 := GenS1{F0: 255, F1: 618}
        var v19 bool = ((557 != 0) || (!false))
        v19 = (!false)
        _ = fi17
        _ = v18
        _ = v19
    }
    v20 := GenS2{F0: v15, F1: (-(v16))}
    var v21 uint32 = 1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        v1 = v1
        fmt.Println(v21, v15, v15)
        _ = fi22
    }
    v16 = (-(v16))
    v23 := GenS3{F0: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
}
