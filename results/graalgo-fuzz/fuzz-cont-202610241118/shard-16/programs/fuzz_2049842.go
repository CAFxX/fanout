package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint }
func genFn0(p0 uint, p1 int8) string {
    v7 := GenS1{F0: 0, F1: -32768, F2: p0}
    var v8 float64 = -2.25
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return "hello"
}

func main() {
    var v1 uint32 = 1
    var v2 string = (("\n\t" + ("a" + ("" + ("a" + "hello")))) + "hello")
    var v3 uint = 478
    var v4 int64 = -845
    var v5 uint16 = 42
    var v6 int = 42
    v2 = v2
    switch -867 {
    case 706:
        fmt.Println(v1, v2)
    case 161:
        v9 := GenS1{F0: (v4 % ((v4) | 1)), F1: -1, F2: (v3 % ((uint(v5)) | 1))}
        _ = v9
    default:
        if (245 > (^(v1))) {
            v10 := []int{v6}
            v3 += v3
            _ = v10
        } else {
            v11 := []int{v6, 779}
            fmt.Println(v5, v1)
            for fi12 := 0; fi12 < 13; fi12++ {
                fmt.Println(v6, v5)
                _ = fi12
            }
            _ = v11
        }
    }
    if true {
        v1 -= (^(v1))
        fmt.Println(v5, v4, v5)
    } else {
        v3 += v3
    }
    v6 = -2147483648
    v6 += v6
    fmt.Println(v5, v5, v5)
    var v13 int16 = -49
    for fi14 := 0; fi14 < 16; fi14++ {
        for fi15 := 0; fi15 < 16; fi15++ {
            fmt.Println(v5, v3, v2)
            _ = fi15
        }
        for fi16 := 0; fi16 < 5; fi16++ {
            switch fi16 {
            case 128:
                var v17 float32 = (-2.25 - 1e10)
                _ = v17
            case 2147483647:
                if true {
                    if false {
                        if (true && (v2 == v2)) {
                            var v18 string = "\n\t"
                            v13 = 1
                            _ = v18
                        } else {
                            v19 := GenS1{F0: v4, F1: v13, F2: 0}
                            _ = v19
                        }
                        v20 := []int{}
                        _ = v20
                    }
                    fmt.Println(v13)
                } else {
                    var v21 float32 = 1.5
                    v5 -= v5
                    v22 := GenS1{F0: 67, F1: v13, F2: (^(v3))}
                    _ = v21
                    _ = v22
                }
            }
            var v23 float32 = 1.5
            var v24 uint = 65535
            _ = fi16
            _ = v23
            _ = v24
        }
        v25 := []int{0, int(v5), (v6 % (((^(v6))) | 1))}
        _ = fi14
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
