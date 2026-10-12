package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint8, p1 int, p2 int) uint32 {
    var v5 uint32 = 657
    switch "hello" {
    case "\n\t":
        v6 := []int{(p2 / (((-(p1))) | 1)), p1, (p2 * p2), p2}
        _ = v6
    case "foo bar":
        var v7 float64 = 3.14159
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (v5 + v5)
}

func main() {
    v1 := []string{"hello"}
    var v2 uint16 = 255
    v3 := make(map[int]bool)
    var v4 uint = 65535
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 string = "\n\t"
        var v10 uint64 = 18446744073709551615
        _ = fi8
        _ = v10
        _ = v9
    }
    fmt.Println(v4)
    v4 = (v4 * 2)
    for fi11 := 0; fi11 < 13; fi11++ {
        for fi12 := 0; fi12 < 9; fi12++ {
            switch (fi11 - (fi12 / (((fi11 % ((fi11) | 1))) | 1))) {
            case 127:
                var v13 uint = 1
                _ = v13
            }
            if false {
                var v14 uint16 = v2
                switch (fi12 / (((fi11 % ((fi11) | 1))) | 1)) {
                case 127:
                    v2 += (v14 >> ((v4) & 63))
                case -917:
                    v14 += v14
                }
                v14 -= (v14 / ((v14) | 1))
                _ = v14
            }
            v2 = v2
            _ = fi12
        }
        if ((("hello" + "") + "hello") < "hello") {
            var v15 int = (fi11 >> (((^(v4))) & 63))
            fmt.Println(fi11)
            var v16 float64 = (1.5 / (3.14159 * 0.1))
            _ = v15
            _ = v16
        }
        _ = fi11
    }
    v17 := GenS1{F0: 123, F1: -1}
    if (!(-32768 < 0)) {
        switch 128 {
        case 127:
            switch -9223372036854775808 {
            case -59:
                for fi18 := 0; fi18 < 18; fi18++ {
                    var v19 uint64 = 4294967295
                    _ = fi18
                    _ = v19
                }
            case 2147483647:
                v4 = v4
            case 0:
                v3[-2147483648] = false
            }
        case 128:
            v4 += v4
        }
        fmt.Println(v2, v4)
        v4 += (v4 / (((^(v4))) | 1))
    }
    var v20 string = ("foo bar" + ("hello" + ""))
    var v21 bool = (-1 < 0)
    switch 2147483647 {
    case 0:
        if v21 {
            switch (v20 + v20) {
            case "hello":
                switch 601 {
                case 379:
                    var v22 uint = 18446744073709551615
                    _ = v22
                case -1:
                    v23 := GenS1{F0: 127, F1: 0}
                    _ = v23
                case 0:
                    v3[127] = ((-2.25 < (1e10 - (((1.5 - float64(1.5)) + 0.1) / (float64(1e10) * ((0.1 - (1e10 - 0.1)) - (0.0 / float64((1e10 - ((((-2.25 - (0.0 / 0.1)) / (0.1 + -2.25)) + 0.0) - 1e10))))))))) && false)
                }
            case "a":
                var v24 int = -774
                _ = v24
            case "foo bar":
                switch (v20 + ("foo bar" + (("" + (("\n\t" + "") + v20)) + (v20 + "foo bar")))) {
                case "a":
                    v20 = "\n\t"
                }
            default:
                for fi25 := 0; fi25 < 20; fi25++ {
                    var v26 uint32 = 65535
                    _ = fi25
                    _ = v26
                }
            }
            switch "a" {
            case "":
                var v27 int = -622
                _ = v27
            case "foo bar":
                switch 1 {
                case -838:
                    var v28 uint16 = v2
                    _ = v28
                case 2147483647:
                    v4 -= 65535
                default:
                    v29 := GenS2{F0: -264}
                    _ = v29
                }
            case "a":
                v4 += uint(v2)
            default:
                var v30 float64 = (3.14159 - 0.0)
                _ = v30
            }
            fmt.Println(v20, v20, v2)
        } else {
            var v31 string = v20
            v32 := GenS1{F0: uint8(v4), F1: 1}
            _ = v31
            _ = v32
        }
    }
    var v33 uint32 = 65535
    for fi34 := 0; fi34 < 18; fi34++ {
        v35 := GenS3{F0: -306}
        _ = fi34
        _ = v35
    }
    _ = len(v3)
    v1 = append(v1, "")
    for fi36 := 0; fi36 < 15; fi36++ {
        if len(v1) > 0 {
            v1[0] = (v20 + v20)
        }
        v2 += 2
        var v37 uint64 = 18446744073709551615
        _ = fi36
        _ = v37
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v33
    _ = v4
}
