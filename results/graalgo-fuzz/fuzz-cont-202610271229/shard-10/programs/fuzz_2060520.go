package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
func genFn0(p0 uint32) int64 {
    var v4 string = "a"
    v4 = "\n\t"
    if true {
        if false {
            v4 = v4
            for fi5 := 0; fi5 < 3; fi5++ {
                v6 := []int{}
                _ = fi5
                _ = v6
            }
        }
        var v7 int16 = 971
        _ = v7
    }
    var v8 uint8 = 43
    fmt.Println(p0, v8)
    _ = p0
    _ = v4
    _ = v8
    return 1
}

func main() {
    var v1 int = 4294967295
    var v2 uint = 255
    v3 := []int{(v1 % ((v1) | 1))}
    switch (v1 / (((v1 & v1)) | 1)) {
    case -911:
        for fi9 := 0; fi9 < 12; fi9++ {
            v3 = append(v3, (^(fi9)))
            _ = fi9
        }
    case 1:
        if (919 < 1) {
            var v10 float64 = 0.0
            for fi11 := 0; fi11 < 3; fi11++ {
                switch ("\n\t" + "") {
                case "a":
                    v1 -= fi11
                case "\n\t":
                    if false {
                        v1 = (fi11 / ((fi11) | 1))
                        var v12 float64 = v10
                        _ = v12
                    } else {
                        v13 := GenS1{F0: 40, F1: 560}
                        switch fi11 {
                        case 363:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        fmt.Println(v10, v2, v10)
                        _ = v13
                    }
                case "hello":
                    v3 = append(v3, (fi11 - (v1 - -2147483648)))
                default:
                    var v14 int = (fi11 ^ 1)
                    _ = v14
                }
                _ = fi11
            }
            var v15 int64 = 1
            _ = v10
            _ = v15
        }
    case 128:
        var v16 int16 = 1
        _ = v16
    }
    v2 -= 0
    v17 := GenS1{F0: 42, F1: 1}
    v2 -= 431
    for fi18 := 0; fi18 < 6; fi18++ {
        v3 = append(v3, v1)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi18
    }
    v19 := GenS1{F0: 42, F1: 42}
    var v20 float32 = 3.14159
    switch -320 {
    case 9223372036854775807:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
