package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint) int {
    v4 := GenS1{F0: -1, F1: 32767}
    var v5 int8 = 2
    _ = p0
    _ = v4
    _ = v5
    return 128
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]uint32)
    var v3 int8 = 42
    delete(v1, "a")
    v3 = -76
    fmt.Println(v3)
    v3 = v3
    var v6 uint64 = 149
    for fi7 := 0; fi7 < 11; fi7++ {
        v3 += (^(v3))
        _ = fi7
    }
    v8 := v2[701]
    v9 := []int{127, -500, 4294967295}
    if true {
        fmt.Println(v8, v8, v3)
    } else {
        for fi10 := 0; fi10 < 18; fi10++ {
            if (false || true) {
                for fi11 := 0; fi11 < 18; fi11++ {
                    if len(v9) > 0 {
                        v9[0] = int(v3)
                    }
                    if len(v9) > 0 {
                        v9[0] = int(v3)
                    }
                    var v12 uint8 = uint8(v8)
                    _ = fi11
                    _ = v12
                }
            }
            var v13 uint8 = 0
            v14 := GenS1{F0: 0, F1: fi10}
            _ = fi10
            _ = v13
            _ = v14
        }
        if false {
            if true {
                v15 := GenS2{F0: -51}
                _ = v15
            } else {
                v16 := GenS2{F0: 0}
                _ = v16
            }
        }
        fmt.Println(v6, v8)
    }
    for fi17 := 0; fi17 < 18; fi17++ {
        fmt.Println(v8, fi17, v3)
        _ = fi17
    }
    for fi18 := 0; fi18 < 2; fi18++ {
        if (v6 < v6) {
            if ((true || (42 != 0)) || (88 < 4294967295)) {
                v19 := GenS1{F0: -298, F1: (fi18 >> ((4294967295) & 63))}
                _ = v19
            }
            v20 := GenS2{F0: -32768}
            var v21 int8 = v3
            _ = v20
            _ = v21
        } else {
            switch "" {
            case "":
                var v22 uint = 1
                _ = v22
            case "hello":
                var v23 int = genFn0(4294967295)
                _ = v23
            case "\n\t":
                fmt.Println(v8, v6, v6)
            default:
                v24 := GenS1{F0: 674, F1: (fi18 | genFn0(1))}
                _ = v24
            }
            var v25 int8 = (v3 | -1)
            if (1 > 0) {
                v25 = 100
                fmt.Println(fi18, v3, fi18)
            }
            _ = v25
        }
        v26 := GenS3{F0: 20}
        _ = fi18
        _ = v26
    }
    v27 := v1["\n\t"]
    fmt.Println(v27, v27)
    v27 = (v27 && (-674 == 183))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v6
    _ = v8
    _ = v9
}
