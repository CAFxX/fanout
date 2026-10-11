package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 uint16 }
type GenS2 struct { F0 uint64; F1 uint; F2 bool }
func genFn0(p0 bool, p1 int64, p2 bool) string {
    var v5 uint32 = 2147483648
    fmt.Println(p2)
    var v6 int32 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return "foo bar"
}
func genFn1() uint8 {
    v7 := []int{-1}
    var v8 uint16 = 833
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 int16 = -1
    var v2 float64 = 0.1
    var v3 int32 = -1
    var v4 uint32 = 59
    var v9 float32 = ((((1e10 - 0.0) - 1e10) + 0.1) * 3.14159)
    v4 = 1
    v9 = 0.1
    v1 = v1
    if ("\n\t" >= "foo bar") {
        for fi10 := 0; fi10 < 5; fi10++ {
            for fi11 := 0; fi11 < 19; fi11++ {
                switch ((("a" + "a") + "a") + "a") {
                case "a":
                    v12 := []int{fi11}
                    _ = v12
                case "":
                    v13 := []int{(fi11 + fi11)}
                    if false {
                        v4 += 1
                    }
                    _ = v13
                default:
                    var v14 bool = (!(!((v2 * v2) > 3.14159)))
                    _ = v14
                }
                _ = fi11
            }
            var v15 int8 = int8(v4)
            for fi16 := 0; fi16 < 15; fi16++ {
                v17 := []int{}
                fmt.Println(v4)
                _ = fi16
                _ = v17
            }
            _ = fi10
            _ = v15
        }
    }
    v9 += v9
    for fi18 := 0; fi18 < 12; fi18++ {
        v4 = (^(v4))
        _ = fi18
    }
    for fi19 := 0; fi19 < 11; fi19++ {
        v20 := GenS1{F0: 20, F1: v4, F2: 1}
        v21 := GenS2{F0: 4294967295, F1: 65535, F2: true}
        var v22 int = (fi19 & fi19)
        _ = fi19
        _ = v20
        _ = v21
        _ = v22
    }
    fmt.Println(v1, v9)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
