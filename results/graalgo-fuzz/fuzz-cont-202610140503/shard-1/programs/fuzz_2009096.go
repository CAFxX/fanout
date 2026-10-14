package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }

func main() {
    var v1 string = ""
    var v2 int16 = 32767
    var v3 bool = (2147483647 > 1)
    if true {
        for fi4 := 0; fi4 < 4; fi4++ {
            v5 := GenS1{F0: 1, F1: 1}
            for fi6 := 0; fi6 < 3; fi6++ {
                var v7 uint32 = 973
                v8 := []int{fi6, fi6, (fi6 & 127)}
                _ = fi6
                _ = v7
                _ = v8
            }
            _ = fi4
            _ = v5
        }
    } else {
        fmt.Println(v3, v2, v2)
        v1 = (v1 + v1)
    }
    for fi9 := 0; fi9 < 5; fi9++ {
        v10 := []int{fi9, (fi9 * 32767)}
        for fi11 := 0; fi11 < 18; fi11++ {
            var v12 uint32 = 2147483647
            var v13 int = fi11
            if (!(-128 < 100)) {
                if (v3 || v3) {
                    var v14 int16 = 32767
                    var v15 uint32 = v12
                    var v16 float32 = ((-2.25 / 1.0) * ((0.0 / 1e10) + float32(1.5)))
                    _ = v14
                    _ = v15
                    _ = v16
                } else {
                    _ = len(v10)
                    if (40 < (v12 / (((v12 % ((959) | 1))) | 1))) {
                        v13 = fi11
                        v3 = v3
                        if (true && v3) {
                            switch fi11 {
                            case 846:
                                v12 -= 1
                            }
                            var v17 bool = (-578 == 0)
                            _ = v17
                        }
                    }
                }
                var v18 uint = 535
                _ = v18
            } else {
                var v19 int32 = 42
                _ = v19
            }
            _ = fi11
            _ = v12
            _ = v13
        }
        var v20 int = (-(fi9))
        _ = fi9
        _ = v10
        _ = v20
    }
    v2 -= (v2 % ((32767) | 1))
    for fi21 := 0; fi21 < 19; fi21++ {
        var v22 string = "a"
        _ = fi21
        _ = v22
    }
    var v23 string = (("hello" + ((("a" + "foo bar") + v1) + "hello")) + v1)
    switch 1 {
    case 128:
        var v24 int = 107
        _ = v24
    }
    v25 := []int{}
    v3 = false
    v1 = "hello"
    var v26 uint8 = 255
    var v27 bool = ((790 == 113) || (!(!((false && (35 == 1)) || (true || false)))))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v26
    _ = v27
    _ = v3
}
