package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 uint64, p1 int32) bool {
    v6 := GenS1{F0: 100}
    for fi7 := 0; fi7 < 12; fi7++ {
        for fi8 := 0; fi8 < 2; fi8++ {
            var v9 int32 = p1
            var v10 int = fi7
            if true {
                switch (v10 & (-(fi8))) {
                case -1:
                    var v11 uint64 = (p0 ^ p0)
                    _ = v11
                default:
                    v12 := []int{(fi8 + 127), fi7}
                    _ = v12
                }
                var v13 uint = 255
                _ = v13
            } else {
                v10 = (-(v10))
                fmt.Println(v9, fi7)
            }
            var v14 int = int(p0)
            v9 = 1
            _ = fi8
            _ = v10
            _ = v14
            _ = v9
        }
        v15 := []int{fi7, -2147483648, fi7}
        _ = fi7
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v6
    return true
}

func main() {
    var v1 float64 = 3.14159
    var v2 string = ("hello" + "hello")
    var v3 uint8 = 108
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]uint32)
    v16 := v4["\n\t"]
    var v17 string = v2
    v2 = "foo bar"
    switch v2 {
    case "\n\t":
        var v18 float64 = v1
        _ = v18
    }
    var v19 float32 = float32((v1 - 0.1))
    var v20 int32 = 2147483647
    v21 := v5[(v16 - (v16 << ((255) & 63)))]
    v2 = (v17 + v2)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
