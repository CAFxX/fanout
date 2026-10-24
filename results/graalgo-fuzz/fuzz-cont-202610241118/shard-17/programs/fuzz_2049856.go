package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int8, p1 bool, p2 uint) uint8 {
    switch 32767 {
    case 405:
        v6 := GenS1{F0: 42}
        _ = v6
    case 128:
        var v7 uint64 = 4294967295
        _ = v7
    }
    p2 = (^(p2))
    for fi8 := 0; fi8 < 19; fi8++ {
        for fi9 := 0; fi9 < 14; fi9++ {
            var v10 string = ("\n\t" + "a")
            fmt.Println(fi9, p0, fi9)
            _ = fi9
            _ = v10
        }
        if (float64(0.1) < -2.25) {
            for fi11 := 0; fi11 < 17; fi11++ {
                var v12 string = "a"
                switch (fi8 % (((fi8 << (((p2 - p2)) & 63))) | 1)) {
                case 790:
                    v12 = (v12 + "a")
                case 128:
                    v13 := []int{fi11, fi11, fi8}
                    _ = v13
                case 1:
                    var v14 uint8 = 1
                    _ = v14
                default:
                    var v15 int16 = 32767
                    _ = v15
                }
                var v16 uint64 = 0
                _ = fi11
                _ = v12
                _ = v16
            }
            switch "hello" {
            case "hello":
                fmt.Println(p0, fi8)
            case "foo bar":
                var v17 uint8 = 88
                _ = v17
            case "":
                p2 = p2
            }
        }
        var v18 uint = 471
        _ = fi8
        _ = v18
    }
    p1 = ("hello" <= ("hello" + "\n\t"))
    _ = p0
    _ = p1
    _ = p2
    return 2
}

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[uint32]bool)
    var v3 int = 2147483647
    v4 := []uint32{42, 2}
    var v5 int = v3
    v19 := GenS2{F0: 42}
    var v20 bool = true
    fmt.Println(v3)
    for fi21 := 0; fi21 < 1; fi21++ {
        _ = len(v4)
        _ = fi21
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi22 := 0; fi22 < 20; fi22++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi22
    }
    v5 = v5
    fmt.Println(v20, v20, v5)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
