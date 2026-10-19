package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 int; F2 uint32 }
func genFn0(p0 uint32, p1 string) int64 {
    p0 += p0
    v6 := GenS1{F0: p0}
    if (!(true || true)) {
        v7 := []int{}
        if true {
            for fi8 := 0; fi8 < 10; fi8++ {
                if false {
                    v9 := GenS1{F0: 65535}
                    if (990 <= 368) {
                        var v10 uint8 = 1
                        _ = v10
                    }
                    v11 := GenS1{F0: (p0 * (p0 * (p0 << ((149) & 63))))}
                    _ = v11
                    _ = v9
                }
                _ = fi8
            }
            var v12 int8 = -1
            _ = v12
        }
        p1 = "\n\t"
        _ = v7
    } else {
        v13 := []int{0, 42, 127}
        v13 = append(v13, 63)
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v6
    return -1
}

func main() {
    var v1 int64 = 42
    var v2 int = -1
    var v3 uint = 0
    v4 := []string{(("a" + "a") + "foo bar"), "hello", "a", ("hello" + "\n\t")}
    var v5 uint32 = 2147483648
    v4 = append(v4, "")
    var v14 float32 = -2.25
    v15 := GenS2{F0: (v1 | (v1 >> (((v3 ^ 2)) & 63))), F1: int(v5), F2: (^(v5))}
    v5 -= (^(v5))
    v1 += (v1 & genFn0(v5, "\n\t"))
    switch "hello" {
    case "\n\t":
        v14 = (v14 * v14)
    case "":
        if (82 > -477) {
            switch (v2 >> (((v3 >> ((v3) & 63))) & 63)) {
            case 2147483647:
                for fi16 := 0; fi16 < 4; fi16++ {
                    _ = len(v4)
                    _ = fi16
                }
            case 127:
                if true {
                    var v17 int8 = 1
                    _ = v17
                } else {
                    switch "a" {
                    case "a":
                        for fi18 := 0; fi18 < 20; fi18++ {
                            v4 = append(v4, "hello")
                            _ = fi18
                        }
                    case "":
                        v14 = v14
                    default:
                        fmt.Println(v2)
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
            default:
                var v19 float32 = 0.1
                _ = v19
            }
            v2 = 9223372036854775807
        }
    case "foo bar":
        fmt.Println(v1, v14, v1)
    }
    fmt.Println(v3, v1, v3)
    var v20 uint32 = (^(v5))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
