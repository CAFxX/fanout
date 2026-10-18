package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 int64 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 int32 }

func main() {
    v1 := []uint32{2, 1}
    v2 := make(map[string]int)
    var v3 string = ("foo bar" + "")
    v4 := make(map[int]bool)
    fmt.Println(v3)
    fmt.Println(v3)
    v3 = ("foo bar" + "")
    for fi5 := 0; fi5 < 17; fi5++ {
        v6 := GenS1{F0: 0, F1: 0, F2: 0}
        v2[v3] = 163
        _ = fi5
        _ = v6
    }
    _ = len(v2)
    var v7 string = ("a" + "\n\t")
    var v8 int = 2147483647
    v8 -= v8
    v9 := GenS2{F0: 42, F1: 595}
    delete(v4, v8)
    switch v8 {
    case -410:
        for fi10 := 0; fi10 < 8; fi10++ {
            v7 = ""
            for fi11 := 0; fi11 < 6; fi11++ {
                var v12 uint = 430
                switch (v3 + v7) {
                case "a":
                    v13 := GenS2{F0: -628, F1: 2}
                    _ = v13
                case "":
                    _ = len(v1)
                case "hello":
                    v8 -= (fi11 - (fi11 + (fi11 | v8)))
                }
                fmt.Println(fi10, fi11)
                _ = fi11
                _ = v12
            }
            v14 := GenS3{F0: 2147483647}
            _ = fi10
            _ = v14
        }
    case 32767:
        v8 -= v8
    case 4294967295:
        var v15 uint32 = 948
        _ = v15
    }
    fmt.Println(v7, v7)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
