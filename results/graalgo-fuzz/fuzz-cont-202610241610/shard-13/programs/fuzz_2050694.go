package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0() bool {
    var v4 uint64 = 641
    var v5 int16 = 42
    v6 := []int{-16, 4294967295}
    var v7 int8 = 127
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return true
}
func genFn1(p0 uint8, p1 uint64) int {
    p0 += 2
    p1 -= 42
    v8 := []int{}
    _ = p0
    _ = p1
    _ = v8
    return 32768
}

func main() {
    var v1 string = ("foo bar" + (((("" + ("hello" + ("foo bar" + ""))) + "foo bar") + "\n\t") + "\n\t"))
    var v2 string = ("\n\t" + v1)
    var v3 uint32 = 42
    v9 := GenS1{F0: ((((("\n\t" + ("" + v1)) + "foo bar") + v2) + "a") + v2)}
    v1 = v1
    if false {
        v10 := []int{genFn1(243, 0), 9223372036854775807, 9223372036854775807, -916}
        var v11 uint32 = v3
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 17; fi12++ {
        v2 = v1
        var v13 float64 = float64(1e10)
        var v14 int = (fi12 | genFn1(127, 912))
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 int8 = 1
    v15 = 1
    v16 := []int{-454, 1, -907}
    if (4294967295 > 18446744073709551615) {
        for fi17 := 0; fi17 < 8; fi17++ {
            for fi18 := 0; fi18 < 3; fi18++ {
                if false {
                    v16 = append(v16, genFn1(42, 162))
                }
                var v19 int32 = 0
                v16 = append(v16, (fi18 - 0))
                _ = fi18
                _ = v19
            }
            _ = fi17
        }
    }
    var v20 uint8 = 97
    var v21 uint64 = 511
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v9
}
