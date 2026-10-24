package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 int32 }
func genFn0(p0 uint16, p1 uint16) uint32 {
    for fi5 := 0; fi5 < 5; fi5++ {
        if (true && (0.0 == 1e10)) {
            var v6 uint64 = 404
            v7 := []int{(fi5 / ((233) | 1)), (fi5 + fi5), (fi5 ^ fi5)}
            _ = v6
            _ = v7
        } else {
            if (p0 == 0) {
                for fi8 := 0; fi8 < 7; fi8++ {
                    var v9 string = ("\n\t" + ("" + "\n\t"))
                    _ = fi8
                    _ = v9
                }
            } else {
                switch "a" {
                case "hello":
                    fmt.Println(fi5)
                case "foo bar":
                    fmt.Println(p0)
                default:
                    v10 := []int{int(p1), -729, 1}
                    _ = v10
                }
                fmt.Println(fi5, p1)
            }
            for fi11 := 0; fi11 < 14; fi11++ {
                v12 := []int{}
                _ = fi11
                _ = v12
            }
        }
        for fi13 := 0; fi13 < 2; fi13++ {
            switch 42 {
            case -2147483648:
                var v14 int8 = -92
                _ = v14
            default:
                var v15 bool = (!false)
                _ = v15
            }
            _ = fi13
        }
        _ = fi5
    }
    switch 89 {
    case 42:
        fmt.Println(p1, p0)
    default:
        switch -479 {
        case 2147483647:
            var v16 uint16 = p1
            _ = v16
        case 4294967295:
            if (-1 > 229) {
                p1 = 1
            } else {
                p1 = p0
            }
        case 32767:
            v17 := GenS1{F0: ((((!true) && (false && (true && (true || false)))) && true) && true), F1: 255}
            _ = v17
        default:
            for fi18 := 0; fi18 < 10; fi18++ {
                p1 = p0
                _ = fi18
            }
        }
    }
    _ = p0
    _ = p1
    return 65535
}

func main() {
    var v1 uint32 = 2
    var v2 float64 = float64(1e10)
    var v3 int16 = 32767
    var v4 uint64 = 0
    v3 += v3
    v19 := GenS2{F0: -185}
    fmt.Println(v4, v2, v4)
    v20 := GenS1{F0: (3.14159 > -2.25), F1: 127}
    var v21 uint32 = v1
    switch 938 {
    case 9223372036854775807:
        v22 := GenS2{F0: -940}
        _ = v22
    case 128:
        v23 := GenS1{F0: ((!true) || false), F1: 127}
        _ = v23
    }
    for fi24 := 0; fi24 < 16; fi24++ {
        v3 = (v3 % ((1000) | 1))
        _ = fi24
    }
    switch 4294967295 {
    case 127:
        fmt.Println(v1, v2, v4)
    case -45:
        var v25 int8 = -75
        _ = v25
    }
    for fi26 := 0; fi26 < 8; fi26++ {
        var v27 uint16 = 255
        _ = fi26
        _ = v27
    }
    fmt.Println(v1, v1, v1)
    for fi28 := 0; fi28 < 2; fi28++ {
        v29 := GenS2{F0: 946}
        _ = fi28
        _ = v29
    }
    v30 := GenS2{F0: 520}
    if (42 <= 83) {
        v4 += (^(v4))
        v4 = v4
        v31 := []int{1, 135, 128}
        _ = v31
    }
    var v32 uint32 = (v21 / (((v21 >> ((65535) & 63))) | 1))
    v4 -= (v4 * v4)
    fmt.Println(v21)
    v33 := GenS2{F0: -2147483648}
    v21 = genFn0(393, 32767)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v4
}
