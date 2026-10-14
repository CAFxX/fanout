package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 int8, p1 int8) uint8 {
    var v4 uint32 = 65535
    for fi5 := 0; fi5 < 19; fi5++ {
        switch (fi5 % ((-632) | 1)) {
        case 2147483647:
            fmt.Println(p1, p0)
        case 32768:
            if false {
                p0 -= -128
            } else {
                v4 = (^(v4))
                var v6 float32 = 1e10
                _ = v6
            }
        default:
            for fi7 := 0; fi7 < 10; fi7++ {
                switch "foo bar" {
                case "foo bar":
                    var v8 int16 = -259
                    _ = v8
                case "a":
                    var v9 uint8 = 42
                    _ = v9
                }
                if false {
                    fmt.Println(fi7)
                    v4 -= (v4 * v4)
                    v10 := GenS1{F0: 2147483647}
                    _ = v10
                } else {
                    switch "a" {
                    case "hello":
                        p0 = (p0 * p0)
                    case "a":
                        var v11 int8 = (p1 ^ p1)
                        _ = v11
                    case "foo bar":
                        v12 := []int{fi7, (fi5 >> ((42) & 63))}
                        _ = v12
                    default:
                        if (42 > -886) {
                            if false {
                                var v13 uint32 = (^(v4))
                                v4 = 2147483647
                                v14 := []int{(fi5 * fi7), 15, fi5, (fi7 + 420)}
                                _ = v13
                                _ = v14
                            }
                            p1 += 127
                            fmt.Println(p0)
                            v15 := GenS1{F0: -572}
                            _ = v15
                        } else {
                            p1 = p0
                        }
                    }
                    v16 := []int{(fi5 * (fi5 ^ (fi5 * (^(fi5))))), fi7, (fi5 * fi5), fi7}
                    _ = v16
                }
                _ = fi7
            }
        }
        var v17 bool = true
        v17 = v17
        _ = fi5
        _ = v17
    }
    _ = p0
    _ = p1
    _ = v4
    return 1
}

func main() {
    v1 := make(map[int]bool)
    var v2 float32 = 1.5
    var v3 bool = false
    v18 := GenS1{F0: 42}
    for fi19 := 0; fi19 < 11; fi19++ {
        fmt.Println(v3, v3, fi19)
        _ = fi19
    }
    v20 := GenS1{F0: 0}
    v21 := GenS1{F0: -1}
    switch ("hello" + ("" + "\n\t")) {
    case "\n\t":
        if (!(((65535 > 2147483648) || (!v3)) && false)) {
            var v22 float32 = 0.0
            _ = v22
        }
    }
    if (!false) {
        var v23 uint = 18446744073709551615
        _ = v23
    }
    var v24 int = 809
    var v25 bool = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v3
}
