package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64 }
func genFn0() bool {
    var v7 int8 = 1
    v8 := []int{2147483647, 42}
    if ((4294967295 < 0) && ((true && (-32768 < -32768)) && false)) {
        v7 = (v7 ^ v7)
        for fi9 := 0; fi9 < 16; fi9++ {
            for fi10 := 0; fi10 < 17; fi10++ {
                if true {
                    if len(v8) > 0 {
                        v8[0] = fi9
                    }
                    fmt.Println(v7, fi10)
                    v7 = (v7 - (v7 / ((0) | 1)))
                }
                fmt.Println(fi10)
                v7 = v7
                _ = fi10
            }
            _ = fi9
        }
    }
    _ = v7
    _ = v8
    return (1.5 != (((-2.25 - 0.0) / 1.0) / float64(1e10)))
}

func main() {
    v1 := make(map[string]string)
    var v2 uint8 = 2
    var v3 int = 2147483647
    var v4 uint64 = 42
    v5 := []string{}
    var v6 string = ("\n\t" + "hello")
    var v11 uint = 65535
    v12 := GenS1{F0: 1}
    v5 = append(v5, "a")
    for fi13 := 0; fi13 < 9; fi13++ {
        switch fi13 {
        case -1:
            v6 = v6
        }
        for fi14 := 0; fi14 < 16; fi14++ {
            fmt.Println(v3, v6, v2)
            if false {
                _ = len(v5)
                if true {
                    fmt.Println(v6, v2, v3)
                    delete(v1, v6)
                    _ = len(v1)
                }
                delete(v1, "")
            }
            v6 = ""
            _ = fi14
        }
        var v15 int16 = 1000
        _ = fi13
        _ = v15
    }
    if len(v5) > 0 {
        v5[0] = ((v6 + "a") + "")
    }
    v16 := GenS2{F0: v4}
    v2 = (v2 | (^(v2)))
    v3 = (v3 >> ((96) & 63))
    var v17 int32 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
