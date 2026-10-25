package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint) uint {
    for fi4 := 0; fi4 < 19; fi4++ {
        var v5 int64 = 0
        v6 := []int{fi4, fi4, 733}
        _ = fi4
        _ = v5
        _ = v6
    }
    v7 := []int{4294967295, 4294967295}
    _ = p0
    _ = v7
    return p0
}
func genFn1(p0 uint, p1 uint16, p2 uint8) string {
    v8 := GenS1{F0: 879}
    var v9 int = 127
    switch "a" {
    case "\n\t":
        for fi10 := 0; fi10 < 19; fi10++ {
            for fi11 := 0; fi11 < 11; fi11++ {
                var v12 uint = (p0 & 109)
                _ = fi11
                _ = v12
            }
            switch -124 {
            case 28:
                for fi13 := 0; fi13 < 8; fi13++ {
                    var v14 uint8 = (^(p2))
                    p2 = (p2 + (p2 & v14))
                    _ = fi13
                    _ = v14
                }
            }
            switch "\n\t" {
            case "":
                var v15 string = "foo bar"
                _ = v15
            }
            _ = fi10
        }
    }
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return "\n\t"
}

func main() {
    v1 := []string{}
    v2 := []uint32{1, 42}
    var v3 int32 = 0
    fmt.Println(v3)
    v16 := GenS2{F0: 42}
    v17 := GenS2{F0: 2147483648}
    v3 -= -535
    var v18 float64 = 0.0
    v19 := GenS1{F0: 1000}
    if false {
        var v20 int32 = v3
        if len(v2) > 0 {
            v2[0] = 40
        }
        _ = len(v1)
        _ = v20
    }
    var v21 float64 = 3.14159
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
}
