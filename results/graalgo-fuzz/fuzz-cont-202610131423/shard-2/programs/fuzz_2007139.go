package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int16) bool {
    switch 568 {
    case -444:
        var v4 string = ((("hello" + "") + "hello") + "hello")
        _ = v4
    default:
        if true {
            v5 := GenS1{F0: 161}
            if true {
                v6 := GenS1{F0: 0}
                v7 := []int{0, 276}
                _ = v6
                _ = v7
            } else {
                for fi8 := 0; fi8 < 3; fi8++ {
                    for fi9 := 0; fi9 < 9; fi9++ {
                        fmt.Println(fi9)
                        var v10 int = fi9
                        if (0 != 83) {
                            v11 := GenS1{F0: (fi9 | v10)}
                            p0 -= 1000
                            _ = v11
                        }
                        _ = fi9
                        _ = v10
                    }
                    for fi12 := 0; fi12 < 12; fi12++ {
                        v13 := []int{(fi12 << ((18446744073709551615) & 63)), (fi8 << ((18446744073709551615) & 63)), fi12}
                        _ = fi12
                        _ = v13
                    }
                    var v14 uint = 42
                    _ = fi8
                    _ = v14
                }
                var v15 string = ""
                var v16 float64 = (-2.25 / 1.5)
                _ = v15
                _ = v16
            }
            _ = v5
        } else {
            p0 += p0
            fmt.Println(p0)
        }
    }
    if false {
        v17 := []int{128, 128}
        _ = v17
    }
    var v18 float64 = (0.0 - (3.14159 * float64(1.5)))
    v19 := []int{127}
    _ = p0
    _ = v18
    _ = v19
    return (false || false)
}

func main() {
    var v1 string = "hello"
    v2 := []string{}
    var v3 int8 = 127
    v2 = append(v2, (v1 + v1))
    v1 = ("" + v1)
    if true {
        if false {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v20 uint8 = 0
            for fi21 := 0; fi21 < 12; fi21++ {
                var v22 int32 = -1
                for fi23 := 0; fi23 < 17; fi23++ {
                    var v24 int8 = (v3 ^ (-(v3)))
                    var v25 bool = false
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                if len(v2) > 0 {
                    v2[0] = v1
                }
                _ = fi21
                _ = v22
            }
            _ = v20
        }
    }
    var v26 int64 = int64(v3)
    v3 = 0
    fmt.Println(v3)
    v27 := GenS2{F0: 585}
    var v28 uint16 = 919
    v29 := GenS1{F0: -319}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
}
