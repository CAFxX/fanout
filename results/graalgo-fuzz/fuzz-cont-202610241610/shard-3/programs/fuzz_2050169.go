package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 int32, p1 bool) int {
    v4 := []int{2147483647, -94}
    switch 4294967295 {
    case -898:
        var v5 int8 = 1
        _ = v5
    case 507:
        p1 = (3.14159 == 1e10)
    case 42:
        if true {
            p0 = p0
        } else {
            switch "foo bar" {
            case "\n\t":
                p1 = true
            default:
                switch 2147483647 {
                case 128:
                    fmt.Println(p0)
                case 2147483647:
                    var v6 bool = ((222 > 128) && p1)
                    _ = v6
                }
            }
            if p1 {
                if p1 {
                    if (true || ((true || (2147483647 != 32767)) && true)) {
                        p0 = p0
                    }
                    var v7 int64 = 1
                    _ = v7
                } else {
                    fmt.Println(p1, p0)
                    v8 := GenS1{F0: -1}
                    var v9 bool = (true && p1)
                    _ = v8
                    _ = v9
                }
            }
        }
    default:
        var v10 int8 = 1
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v4
    return -1
}

func main() {
    v1 := make(map[uint32]string)
    var v2 bool = ((0.0 / (0.1 - (-2.25 * ((1.5 / -2.25) * 1.5)))) <= -2.25)
    v3 := []uint32{2147483647, 583, 737, 1}
    switch genFn0(1, v2) {
    case 0:
        var v11 bool = (777 < -245)
        _ = v11
    }
    var v12 int16 = 0
    v2 = (true && (v2 && (0 != 0)))
    var v13 float64 = (1.5 - 1e10)
    delete(v1, 2147483647)
    v2 = v2
    if len(v3) > 0 {
        v3[0] = 1
    }
    v14 := GenS1{F0: 893}
    v15 := GenS1{F0: 42}
    v2 = v2
    for fi16 := 0; fi16 < 2; fi16++ {
        if (4294967295 != 2147483647) {
            v17 := GenS1{F0: 462}
            var v18 uint32 = 2147483647
            _ = v17
            _ = v18
        } else {
            var v19 string = ""
            _ = v19
        }
        if v2 {
            var v20 uint64 = 635
            _ = v20
        }
        _ = fi16
    }
    var v21 string = "a"
    switch 128 {
    case 32768:
        var v22 uint = 255
        _ = v22
    }
    switch "foo bar" {
    case "hello":
        var v23 uint16 = 65535
        _ = v23
    case "foo bar":
        fmt.Println(v13)
    }
    for fi24 := 0; fi24 < 20; fi24++ {
        var v25 int16 = 547
        var v26 float64 = 1e10
        _ = fi24
        _ = v25
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v3
}
