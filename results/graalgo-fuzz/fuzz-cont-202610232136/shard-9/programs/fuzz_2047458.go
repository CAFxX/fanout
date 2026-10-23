package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string }
func genFn0(p0 uint32, p1 int32) int {
    p0 += (p0 + (p0 % ((2147483647) | 1)))
    var v4 uint8 = 128
    v5 := GenS1{F0: 856}
    v6 := GenS1{F0: 4294967295}
    v7 := []int{1, -1}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return -144
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]string)
    v3 := []string{"\n\t", "\n\t", ("" + "foo bar")}
    var v8 int16 = 1
    var v9 bool = false
    fmt.Println(v9)
    v9 = ((!false) || (1e10 != (float64((((1e10 - (1e10 + 1.5)) / (3.14159 / ((1.5 - (1.5 - 0.1)) - 1.5))) * 1e10)) - 1e10)))
    v8 = (v8 / (((v8 + (v8 % ((v8) | 1)))) | 1))
    for fi10 := 0; fi10 < 20; fi10++ {
        v9 = (42 <= 2)
        var v11 uint64 = 42
        for fi12 := 0; fi12 < 6; fi12++ {
            v9 = (436 > 464)
            if (0 <= -107) {
                v11 = 2
                var v13 uint64 = v11
                _ = v13
            } else {
                v11 += (v11 * 1)
                if false {
                    v11 = (v11 % ((v11) | 1))
                    v11 += (^(v11))
                    switch "\n\t" {
                    case "a":
                        if (42 > 184) {
                            v14 := GenS1{F0: 123}
                            delete(v2, 42)
                            v9 = ((v8 << ((42) & 63)) != (^(v8)))
                            _ = v14
                        }
                    }
                } else {
                    v15 := GenS1{F0: 18446744073709551615}
                    var v16 string = ("hello" + "\n\t")
                    _ = v15
                    _ = v16
                }
                if v9 {
                    var v17 string = "hello"
                    _ = v17
                }
            }
            _ = fi12
        }
        _ = fi10
        _ = v11
    }
    var v18 float64 = ((3.14159 / -2.25) + 3.14159)
    var v19 uint8 = 42
    if false {
        var v20 int64 = 9223372036854775807
        if v9 {
            for fi21 := 0; fi21 < 4; fi21++ {
                v18 = v18
                _ = fi21
            }
            v22 := GenS2{F0: "a"}
            var v23 int8 = -39
            _ = v22
            _ = v23
        } else {
            if len(v3) > 0 {
                v3[0] = "a"
            }
            v1[758] = "\n\t"
        }
        switch 0 {
        case 4294967295:
            fmt.Println(v19, v20, v9)
        case -680:
            v20 = (v20 / ((-852) | 1))
        default:
            var v24 uint16 = 399
            _ = v24
        }
        _ = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
