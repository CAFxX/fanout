package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
func genFn0(p0 uint16, p1 int, p2 bool) uint16 {
    p1 = int(p0)
    var v4 int8 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 65535
}
func genFn1(p0 int8, p1 uint16, p2 uint64) uint {
    fmt.Println(p2)
    var v5 bool = ((p1 / ((uint16(p2)) | 1)) > genFn0((p1 << ((255) & 63)), 32767, false))
    var v6 uint32 = 450
    var v7 string = ("\n\t" + "foo bar")
    v8 := GenS1{F0: -951, F1: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    var v1 string = ("hello" + ("foo bar" + "foo bar"))
    v2 := []uint32{}
    v3 := make(map[string]string)
    v1 = (("\n\t" + (v1 + v1)) + v1)
    for fi9 := 0; fi9 < 4; fi9++ {
        v1 = v1
        _ = fi9
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v1 = v1
    if len(v2) > 0 {
        v2[0] = 42
    }
    delete(v3, (v1 + "foo bar"))
    v1 = v1
    var v10 int16 = 1000
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v2
    _ = v3
}
