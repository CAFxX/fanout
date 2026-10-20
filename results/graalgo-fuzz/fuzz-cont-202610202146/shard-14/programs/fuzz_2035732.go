package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }

func main() {
    var v1 uint64 = 2
    var v2 int = -1
    var v3 float64 = -2.25
    var v4 int = (v2 | v2)
    var v5 uint32 = 801
    v6 := []int{v2}
    fmt.Println(v3)
    for fi7 := 0; fi7 < 19; fi7++ {
        if false {
            v8 := GenS1{F0: (v1 - (v1 << ((18446744073709551615) & 63))), F1: -958}
            v5 = (v5 / ((v5) | 1))
            _ = v8
        } else {
            for fi9 := 0; fi9 < 5; fi9++ {
                if (v5 < v5) {
                    v1 -= v1
                    _ = len(v6)
                } else {
                    v10 := GenS1{F0: v1, F1: -898}
                    var v11 string = "\n\t"
                    _ = v10
                    _ = v11
                }
                _ = fi9
            }
            for fi12 := 0; fi12 < 17; fi12++ {
                var v13 int8 = -1
                if ((42 != -1) && true) {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    v14 := GenS1{F0: v1, F1: -779}
                    _ = v14
                } else {
                    if false {
                        var v15 uint32 = 1
                        if false {
                            v5 -= 1
                        } else {
                            var v16 float32 = 3.14159
                            _ = v16
                        }
                        _ = v15
                    }
                    v4 -= v2
                }
                if len(v6) > 0 {
                    v6[0] = v4
                }
                fmt.Println(fi12, fi12, v5)
                _ = fi12
                _ = v13
            }
            for fi17 := 0; fi17 < 15; fi17++ {
                fmt.Println(v4)
                fmt.Println(fi17, fi7, v3)
                _ = fi17
            }
        }
        _ = fi7
    }
    if len(v6) > 0 {
        v6[0] = (v4 + (-(v2)))
    }
    fmt.Println(v2)
    v6 = append(v6, (v4 << ((0) & 63)))
    var v18 uint64 = (v1 % ((18446744073709551615) | 1))
    if true {
        for fi19 := 0; fi19 < 1; fi19++ {
            _ = len(v6)
            _ = fi19
        }
    }
    v3 = v3
    if (!(v4 == int(v18))) {
        if (false && true) {
            var v20 uint8 = 49
            _ = v20
        }
    } else {
        if len(v6) > 0 {
            v6[0] = 0
        }
        for fi21 := 0; fi21 < 15; fi21++ {
            fmt.Println(fi21, v4, v5)
            if len(v6) > 0 {
                v6[0] = v2
            }
            _ = fi21
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
