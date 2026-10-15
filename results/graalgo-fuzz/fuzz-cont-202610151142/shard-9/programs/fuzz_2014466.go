package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 bool; F1 uint8 }
func genFn0(p0 uint16, p1 int8) uint32 {
    var v7 float32 = 1e10
    v8 := GenS1{F0: (("foo bar" + ((("\n\t" + "hello") + "\n\t") + "hello")) + "foo bar"), F1: "a", F2: 127}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := make(map[string]bool)
    var v2 int64 = -790
    var v3 uint = 65535
    v4 := []string{"foo bar", "\n\t"}
    v5 := make(map[int]bool)
    v6 := make(map[string]bool)
    var v9 bool = (0 > 122)
    switch 164 {
    case -446:
        fmt.Println(v2)
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        for fi11 := 0; fi11 < 19; fi11++ {
            v9 = (458 > 32767)
            _ = fi11
        }
        _ = fi10
    }
    v9 = true
    v12 := GenS2{F0: (v3 >= v3)}
    v13 := GenS3{F0: 65535, F1: 133}
    var v14 int = 2147483647
    fmt.Println(v9)
    var v15 float32 = (float32(0.1) - float32(0.0))
    if v9 {
        for fi16 := 0; fi16 < 12; fi16++ {
            v3 = (v3 << ((42) & 63))
            for fi17 := 0; fi17 < 20; fi17++ {
                if v9 {
                    v3 = 595
                    switch "hello" {
                    case "\n\t":
                        if (!true) {
                            v18 := GenS4{F0: true, F1: 1}
                            _ = v18
                        }
                    }
                    _ = len(v4)
                } else {
                    v4 = append(v4, "\n\t")
                }
                _ = fi17
            }
            if (!true) {
                fmt.Println(v2, v14, v15)
            }
            _ = fi16
        }
        v14 += v14
        _ = len(v4)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
