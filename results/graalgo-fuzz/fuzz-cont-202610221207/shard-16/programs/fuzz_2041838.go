package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]string)
    var v3 int = -286
    if (!true) {
        switch "\n\t" {
        case "":
            if (false && ((v3 % ((127) | 1)) != v3)) {
                fmt.Println(v3)
                fmt.Println(v3)
                var v4 string = ("foo bar" + "a")
                _ = v4
            }
        case "foo bar":
            v5 := GenS1{F0: 1}
            _ = v5
        }
        v6 := v2[""]
        _ = v6
    } else {
        fmt.Println(v3)
        fmt.Println(v3)
        v7 := []int{128, (v3 % (((v3 << ((18446744073709551615) & 63))) | 1)), (v3 / ((v3) | 1)), v3}
        _ = v7
    }
    fmt.Println(v3)
    var v8 int16 = 726
    if (v3 >= (v3 << ((18446744073709551615) & 63))) {
        for fi9 := 0; fi9 < 8; fi9++ {
            _ = len(v2)
            fmt.Println(fi9)
            if false {
                var v10 uint64 = 840
                _ = v10
            }
            _ = fi9
        }
        v3 = v3
        var v11 int64 = 9223372036854775807
        _ = v11
    }
    if false {
        v12 := GenS1{F0: 2147483648}
        if (2147483647 == 656) {
            if false {
                v8 -= (v8 & v8)
                if true {
                    v3 -= v3
                    if false {
                        v8 = 1
                    }
                    fmt.Println(v3)
                }
            }
            for fi13 := 0; fi13 < 14; fi13++ {
                v14 := []int{}
                var v15 string = "a"
                _ = fi13
                _ = v14
                _ = v15
            }
        } else {
            var v16 float64 = 3.14159
            var v17 int16 = v8
            for fi18 := 0; fi18 < 16; fi18++ {
                v17 = 0
                for fi19 := 0; fi19 < 7; fi19++ {
                    v16 = ((1e10 * v16) + (float64(3.14159) + v16))
                    v20 := v2[("hello" + "a")]
                    _ = fi19
                    _ = v20
                }
                _ = fi18
            }
            _ = v16
            _ = v17
        }
        _ = v12
    } else {
        v21 := GenS1{F0: 0}
        _ = v21
    }
    for fi22 := 0; fi22 < 15; fi22++ {
        var v23 int16 = (^(v8))
        v3 -= -1
        for fi24 := 0; fi24 < 18; fi24++ {
            switch v3 {
            case 128:
                fmt.Println(fi24, fi24, v3)
            case 264:
                fmt.Println(v23, v3)
            }
            _ = fi24
        }
        _ = fi22
        _ = v23
    }
    for fi25 := 0; fi25 < 9; fi25++ {
        if (!(!false)) {
            switch "hello" {
            case "":
                v26 := GenS1{F0: 42}
                _ = v26
            case "foo bar":
                if (122 > 18446744073709551615) {
                    v27 := GenS2{F0: 744}
                    _ = v27
                }
            case "a":
                var v28 float32 = (1.5 - 0.1)
                _ = v28
            default:
                fmt.Println(v8, fi25)
            }
        }
        var v29 int = 4294967295
        _ = fi25
        _ = v29
    }
    v3 = v3
    for fi30 := 0; fi30 < 17; fi30++ {
        var v31 uint32 = 770
        v3 = v3
        for fi32 := 0; fi32 < 6; fi32++ {
            v3 = 9223372036854775807
            _ = fi32
        }
        _ = fi30
        _ = v31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v8
}
