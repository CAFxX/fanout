package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool }
func genFn0(p0 int8, p1 int32) string {
    if (false || true) {
        p1 = p1
    } else {
        var v7 int = 1
        for fi8 := 0; fi8 < 3; fi8++ {
            p1 = p1
            _ = fi8
        }
        _ = v7
    }
    v9 := GenS1{F0: 36}
    _ = p0
    _ = p1
    _ = v9
    return ""
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint8 = 114
    var v3 int8 = 63
    var v4 string = ("foo bar" + "a")
    v5 := []int{}
    v6 := make(map[uint32]int)
    switch "foo bar" {
    case "\n\t":
        for fi10 := 0; fi10 < 4; fi10++ {
            if ((1 > 42) || true) {
                for fi11 := 0; fi11 < 4; fi11++ {
                    v12 := GenS2{F0: true}
                    v4 = "foo bar"
                    _ = fi11
                    _ = v12
                }
                v13 := v1[32768]
                var v14 bool = (1 < 477)
                _ = v13
                _ = v14
            }
            _ = fi10
        }
    }
    for fi15 := 0; fi15 < 1; fi15++ {
        v5 = append(v5, (fi15 + (fi15 + fi15)))
        v4 = ("\n\t" + (v4 + ("a" + v4)))
        _ = fi15
    }
    v2 = 1
    var v16 uint8 = v2
    var v17 string = ("\n\t" + "hello")
    var v18 float32 = 0.0
    var v19 bool = false
    v18 -= float32((-2.25 / float64((float32((float64(float32(0.0)) / 3.14159)) / v18))))
    var v20 uint64 = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
