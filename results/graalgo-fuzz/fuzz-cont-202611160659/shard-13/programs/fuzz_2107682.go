package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 string) bool {
    var v7 int64 = 179
    for fi8 := 0; fi8 < 16; fi8++ {
        switch p0 {
        case "hello":
            var v9 int = fi8
            _ = v9
        case "foo bar":
            for fi10 := 0; fi10 < 1; fi10++ {
                p0 = ((("a" + "foo bar") + p0) + "\n\t")
                var v11 int = (fi8 + -9223372036854775808)
                p0 = p0
                switch (v11 << ((987) & 63)) {
                case 32767:
                    var v12 uint16 = 596
                    _ = v12
                case 0:
                    fmt.Println(fi8, fi10)
                }
                _ = fi10
                _ = v11
            }
        case "":
            var v13 bool = (!true)
            _ = v13
        default:
            if true {
                fmt.Println(p0, fi8)
                fmt.Println(v7, v7, p0)
                v14 := []int{49, fi8, fi8}
                _ = v14
            }
        }
        switch (p0 + p0) {
        case "a":
            switch fi8 {
            case 42:
                if (!false) {
                    var v15 float32 = ((((0.1 * ((3.14159 * 0.0) / 1e10)) / 1.0) * -2.25) - 3.14159)
                    var v16 uint = 255
                    v7 += 1
                    _ = v15
                    _ = v16
                } else {
                    if (float32(0.1) != 3.14159) {
                        var v17 int32 = -1
                        var v18 int8 = 1
                        v7 += (v7 >> ((65535) & 63))
                        _ = v17
                        _ = v18
                    } else {
                        if (3.14159 > 1.5) {
                            var v19 int8 = -1
                            v7 = (-(v7))
                            _ = v19
                        }
                        p0 = (p0 + (("" + p0) + (("\n\t" + p0) + p0)))
                        v20 := GenS1{F0: -128}
                        _ = v20
                    }
                    if (42 != 1) {
                        p0 = ""
                        p0 = "hello"
                        v7 -= v7
                    } else {
                        var v21 bool = true
                        _ = v21
                    }
                }
            case -9223372036854775808:
                for fi22 := 0; fi22 < 14; fi22++ {
                    var v23 uint16 = 358
                    _ = fi22
                    _ = v23
                }
            case 1:
                v24 := []int{}
                _ = v24
            default:
                v25 := GenS1{F0: 0}
                _ = v25
            }
        default:
            v7 = (v7 | (v7 + 9223372036854775807))
        }
        _ = fi8
    }
    _ = p0
    _ = v7
    return (!(18446744073709551615 >= 678))
}

func main() {
    var v1 uint64 = 489
    var v2 int8 = 100
    v3 := []int{9223372036854775807}
    var v4 uint64 = (v1 - v1)
    var v5 int = -2147483648
    v6 := make(map[int]string)
    switch 127 {
    case 1:
        for fi26 := 0; fi26 < 15; fi26++ {
            if (748 < 646) {
                var v27 int = -932
                _ = v27
            } else {
                var v28 uint8 = 110
                for fi29 := 0; fi29 < 4; fi29++ {
                    switch "hello" {
                    case "":
                        var v30 bool = (0 != 528)
                        _ = v30
                    default:
                        v3 = append(v3, 0)
                    }
                    fmt.Println(v4, v1, v28)
                    v6[fi29] = "\n\t"
                    _ = fi29
                }
                _ = v28
            }
            v4 += 42
            fmt.Println(fi26, v5)
            _ = fi26
        }
    }
    var v31 uint32 = 4294967295
    v1 = (v4 + (^(v1)))
    v32 := GenS1{F0: 79}
    var v33 uint16 = 32767
    v1 -= v4
    var v34 uint = 65535
    v35 := GenS1{F0: (v2 + -92)}
    _ = len(v6)
    v2 = v2
    v33 = 0
    v34 = v34
    fmt.Println(v31, v1)
    fmt.Println(v34, v33)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
