package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0() uint {
    switch 9223372036854775807 {
    case -684:
        var v7 int8 = -68
        _ = v7
    }
    v8 := []int{-755, 42, -275}
    var v9 uint64 = 18446744073709551615
    v8 = append(v8, 743)
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 int64 = 42
    v2 := []uint32{161, 2147483647}
    var v3 int16 = -1
    var v4 bool = (true || (1 == v1))
    var v5 int32 = 2147483647
    var v6 int8 = 42
    switch -9223372036854775808 {
    case 127:
        switch "foo bar" {
        case "\n\t":
            switch "a" {
            case "foo bar":
                var v10 uint16 = 32767
                _ = v10
            case "\n\t":
                v11 := GenS1{F0: v6}
                _ = v11
            }
        }
    case 539:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case 95:
        if (true || v4) {
            if len(v2) > 0 {
                v2[0] = 912
            }
            var v12 float32 = (0.0 - 0.0)
            _ = v12
        }
    }
    v13 := GenS1{F0: (v6 & (^(v6)))}
    v14 := GenS1{F0: v6}
    v15 := GenS1{F0: int8(v1)}
    if (98 < 32767) {
        v5 = 1
        v3 += 1000
    }
    if false {
        v16 := GenS1{F0: v6}
        for fi17 := 0; fi17 < 10; fi17++ {
            v5 = 914
            v5 = 1
            _ = fi17
        }
        _ = v16
    } else {
        if (v4 || (((!true) || (v5 >= v5)) || (((-(v5)) == v5) && (1 > 331)))) {
            v4 = v4
            v18 := GenS1{F0: v6}
            _ = v18
        } else {
            _ = len(v2)
        }
        for fi19 := 0; fi19 < 15; fi19++ {
            fmt.Println(v5)
            var v20 int32 = (v5 / (((v5 % ((-636) | 1))) | 1))
            _ = fi19
            _ = v20
        }
        for fi21 := 0; fi21 < 8; fi21++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if v4 {
                for fi22 := 0; fi22 < 1; fi22++ {
                    if v4 {
                        var v23 float64 = 1.5
                        _ = v23
                    }
                    _ = fi22
                }
            }
            for fi24 := 0; fi24 < 10; fi24++ {
                var v25 int = fi21
                v5 = v5
                _ = fi24
                _ = v25
            }
            _ = fi21
        }
    }
    v5 -= (-(v5))
    for fi26 := 0; fi26 < 14; fi26++ {
        v1 = (v1 >> ((18446744073709551615) & 63))
        v4 = false
        _ = fi26
    }
    switch "foo bar" {
    case "hello":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "\n\t":
        v27 := GenS1{F0: (v6 & v6)}
        _ = v27
    case "a":
        switch "hello" {
        case "a":
            switch "hello" {
            case "a":
                for fi28 := 0; fi28 < 13; fi28++ {
                    fmt.Println(v5)
                    if len(v2) > 0 {
                        v2[0] = 65535
                    }
                    v6 += v6
                    _ = fi28
                }
            }
        case "hello":
            if (654 < 701) {
                v6 -= 127
            }
        default:
            v6 = (v6 | 25)
        }
    }
    v1 += v1
    v4 = v4
    v3 = (v3 / (((v3 ^ 1)) | 1))
    if true {
        v5 -= -1
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
