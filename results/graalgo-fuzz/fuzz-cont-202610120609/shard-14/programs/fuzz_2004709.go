package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool; F1 int16; F2 string }

func main() {
    v1 := []string{("\n\t" + ("" + ("foo bar" + "")))}
    v2 := []int{-273, 447, 4294967295, -9223372036854775808}
    v3 := make(map[int]string)
    v4 := []int{128}
    var v5 int64 = 9223372036854775807
    if len(v4) > 0 {
        v4[0] = 527
    }
    fmt.Println(v5)
    if false {
        for fi6 := 0; fi6 < 17; fi6++ {
            v3[fi6] = (("\n\t" + ("\n\t" + ("\n\t" + "a"))) + "")
            _ = fi6
        }
    }
    var v7 float64 = float64(0.0)
    var v8 string = ("" + ("\n\t" + ""))
    for fi9 := 0; fi9 < 1; fi9++ {
        if (2 > 155) {
            var v10 int8 = 0
            v5 = (v5 - -352)
            switch fi9 {
            case -69:
                fmt.Println(v10)
            case 128:
                for fi11 := 0; fi11 < 14; fi11++ {
                    if false {
                        if len(v2) > 0 {
                            v2[0] = fi11
                        }
                        v8 = "foo bar"
                        v10 = v10
                    } else {
                        var v12 int16 = -32768
                        _ = v12
                    }
                    _ = len(v2)
                    _ = fi11
                }
            case -9223372036854775808:
                var v13 uint = 65535
                _ = v13
            default:
                v8 = v8
            }
            _ = v10
        } else {
            v8 = (("foo bar" + v8) + "a")
        }
        _ = fi9
    }
    if true {
        v14 := v3[2147483647]
        _ = v14
    } else {
        var v15 string = v8
        var v16 float32 = (3.14159 * 3.14159)
        _ = v15
        _ = v16
    }
    v5 = 42
    v17 := GenS1{F0: 1000, F1: (v5 & -1)}
    var v18 int = 4294967295
    v7 = ((v7 * v7) + v7)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v19 bool = true
    if v19 {
        v20 := v3[(^(v18))]
        _ = len(v2)
        for fi21 := 0; fi21 < 16; fi21++ {
            fmt.Println(v19, v5, v5)
            for fi22 := 0; fi22 < 19; fi22++ {
                var v23 string = "\n\t"
                _ = len(v4)
                _ = fi22
                _ = v23
            }
            _ = fi21
        }
        _ = v20
    } else {
        var v24 float32 = (1.5 * 0.1)
        v25 := GenS1{F0: 32767, F1: v5}
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v24
        _ = v25
    }
    if v19 {
        v26 := GenS2{F0: 42}
        _ = v26
    }
    v8 = (((v8 + v8) + ((v8 + v8) + v8)) + "a")
    v27 := GenS3{F0: ((-39 <= -301) && v19), F1: 635, F2: "\n\t"}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
