package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int8) uint32 {
    var v6 uint16 = 32767
    switch "a" {
    case "foo bar":
        if false {
            v7 := []int{}
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            fmt.Println(v6)
            _ = v7
        }
    case "":
        fmt.Println(p0, p0)
    case "\n\t":
        for fi8 := 0; fi8 < 10; fi8++ {
            if (false && (2 >= 0)) {
                v9 := GenS1{F0: 325}
                _ = v9
            }
            _ = fi8
        }
    default:
        for fi10 := 0; fi10 < 4; fi10++ {
            var v11 string = "\n\t"
            _ = fi10
            _ = v11
        }
    }
    p0 -= p0
    v12 := []int{2147483647}
    var v13 string = ("\n\t" + ("foo bar" + ("" + ("" + ""))))
    _ = p0
    _ = v12
    _ = v13
    _ = v6
    return 2147483647
}

func main() {
    var v1 uint = 4294967295
    var v2 uint32 = 2147483648
    v3 := make(map[uint32]uint32)
    var v4 float64 = (1e10 + (float64(0.1) * 3.14159))
    v5 := make(map[uint32]int)
    if (!false) {
        for fi14 := 0; fi14 < 15; fi14++ {
            v4 += 3.14159
            v15 := GenS2{F0: 440}
            _ = fi14
            _ = v15
        }
    }
    v16 := []int{}
    var v17 bool = true
    for fi18 := 0; fi18 < 10; fi18++ {
        for fi19 := 0; fi19 < 14; fi19++ {
            v20 := GenS1{F0: -32768}
            _ = fi19
            _ = v20
        }
        if false {
            v2 = (v2 | v2)
            var v21 string = "hello"
            _ = v21
        } else {
            if (!(!(v17 || (1e10 != v4)))) {
                v1 = v1
                v4 = 3.14159
                for fi22 := 0; fi22 < 14; fi22++ {
                    switch fi18 {
                    case -1:
                        v23 := GenS2{F0: 18446744073709551615}
                        _ = v23
                    case 128:
                        v4 -= -2.25
                    default:
                        var v24 float64 = v4
                        _ = v24
                    }
                    _ = fi22
                }
            }
        }
        fmt.Println(v1, v2)
        _ = fi18
    }
    switch "hello" {
    case "hello":
        v25 := GenS2{F0: 4294967295}
        _ = v25
    case "foo bar":
        v1 = 0
    case "":
        if (true || ((v1 | 2) < (v1 % ((42) | 1)))) {
            fmt.Println(v4, v4, v4)
        }
    }
    if v17 {
        var v26 int8 = 100
        if len(v16) > 0 {
            v16[0] = 513
        }
        fmt.Println(v17)
        _ = v26
    } else {
        _ = len(v5)
        var v27 uint32 = 4294967295
        _ = v27
    }
    v16 = append(v16, -2147483648)
    var v28 uint = v1
    if v17 {
        fmt.Println(v1)
        var v29 int16 = 0
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
