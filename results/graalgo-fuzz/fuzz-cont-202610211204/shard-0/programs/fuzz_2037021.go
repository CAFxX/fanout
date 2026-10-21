package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 uint; F1 uint16; F2 int8 }
type GenS3 struct { F0 int64 }
func genFn0(p0 bool, p1 int16) bool {
    var v4 int = -197
    v5 := GenS1{F0: 42, F1: -907}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return p0
}

func main() {
    var v1 uint32 = 1
    var v2 bool = (true && (true && true))
    v3 := []int{136, 4294967295}
    var v6 uint64 = 146
    var v7 int16 = 1
    v2 = true
    var v8 int16 = (-(v7))
    for fi9 := 0; fi9 < 9; fi9++ {
        var v10 string = "a"
        switch fi9 {
        case -2147483648:
            v11 := GenS1{F0: 42, F1: (-(fi9))}
            _ = v11
        default:
            v1 = v1
        }
        _ = fi9
        _ = v10
    }
    var v12 string = "a"
    v12 = ("foo bar" + v12)
    v13 := GenS2{F0: 238, F1: 731, F2: 1}
    if v2 {
        v3 = append(v3, 1)
    }
    var v14 int64 = 42
    v6 = v6
    for fi15 := 0; fi15 < 3; fi15++ {
        var v16 uint64 = (v6 % (((v6 % ((42) | 1))) | 1))
        for fi17 := 0; fi17 < 13; fi17++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v14 -= (v14 << ((2) & 63))
            _ = fi17
        }
        switch ((v12 + ((((v12 + v12) + v12) + v12) + "")) + "hello") {
        case "a":
            v8 += 32767
        case "hello":
            var v18 uint32 = 42
            _ = v18
        }
        _ = fi15
        _ = v16
    }
    v19 := GenS1{F0: 0, F1: -287}
    v20 := GenS3{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
