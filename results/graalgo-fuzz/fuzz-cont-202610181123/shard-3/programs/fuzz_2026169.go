package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 int8) uint {
    switch 4294967295 {
    case 32768:
        p0 = (p0 / ((p0) | 1))
    case -1:
        v5 := []int{-9223372036854775808, -9223372036854775808, -86}
        _ = v5
    default:
        var v6 int32 = -1
        _ = v6
    }
    fmt.Println(p0)
    p0 += (p0 & p0)
    for fi7 := 0; fi7 < 6; fi7++ {
        p0 += (p0 | (p0 + -82))
        if (-1 > 683) {
            v8 := []int{}
            _ = v8
        }
        _ = fi7
    }
    _ = p0
    return 65535
}

func main() {
    v1 := []uint32{695, 2}
    var v2 float64 = 1e10
    v3 := make(map[int]bool)
    var v4 int64 = -1
    var v9 uint64 = 42
    if ("a" == ("hello" + "\n\t")) {
        var v10 uint = 860
        _ = v10
    } else {
        for fi11 := 0; fi11 < 19; fi11++ {
            var v12 uint16 = 255
            v13 := GenS1{F0: 127}
            switch "foo bar" {
            case "hello":
                for fi14 := 0; fi14 < 17; fi14++ {
                    var v15 uint8 = 0
                    if (!true) {
                        if (false && false) {
                            v2 = -2.25
                        }
                        delete(v3, (fi11 >> ((genFn0(28)) & 63)))
                        var v16 uint16 = (v12 / (((v12 / ((32767) | 1))) | 1))
                        _ = v16
                    }
                    _ = fi14
                    _ = v15
                }
            }
            _ = fi11
            _ = v12
            _ = v13
        }
    }
    fmt.Println(v2, v9)
    for fi17 := 0; fi17 < 10; fi17++ {
        fmt.Println(v4, v4, v9)
        _ = fi17
    }
    delete(v3, 127)
    var v18 string = "a"
    var v19 int8 = -128
    v1 = append(v1, 4294967295)
    if false {
        v20 := v3[-470]
        fmt.Println(v2, v19)
        fmt.Println(v20, v4, v18)
        _ = v20
    }
    fmt.Println(v18, v19)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
