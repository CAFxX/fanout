package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 uint32; F1 uint64 }
func genFn0() string {
    var v6 int8 = 0
    v6 = 42
    v6 = (v6 - v6)
    if true {
        if ((!false) || false) {
            var v7 int8 = (v6 >> ((564) & 63))
            _ = v7
        } else {
            var v8 int16 = 32767
            _ = v8
        }
        v6 += v6
        switch "a" {
        case "a":
            v9 := GenS1{F0: ("\n\t" + ("\n\t" + "a")), F1: 625}
            _ = v9
        }
    }
    _ = v6
    return "foo bar"
}

func main() {
    v1 := make(map[string]string)
    var v2 int = 2147483647
    v3 := make(map[int]int)
    v4 := make(map[int]bool)
    var v5 int32 = 801
    v2 += (v2 % ((-216) | 1))
    switch 622 {
    case 32767:
        for fi10 := 0; fi10 < 12; fi10++ {
            var v11 string = "a"
            var v12 int8 = 0
            v5 -= (v5 << ((2) & 63))
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    fmt.Println(v5)
    if (((232 < 0) || true) || (67 < 255)) {
        v13 := GenS2{F0: 42, F1: 66}
        _ = v13
    } else {
        v2 = v2
        var v14 uint32 = 1
        v14 = (v14 ^ (v14 ^ v14))
        _ = v14
    }
    if ((3.14159 / 1.0) > 1.5) {
        var v15 int16 = -1
        v16 := GenS1{F0: (("\n\t" + "foo bar") + "\n\t"), F1: 42}
        _ = v15
        _ = v16
    } else {
        if ((!true) || (!(false && true))) {
            fmt.Println(v2, v5)
            var v17 float64 = 0.1
            _ = v17
        }
    }
    if (false && (0 > 1)) {
        switch "\n\t" {
        case "foo bar":
            if true {
                switch 231 {
                case 128:
                    v2 = -865
                }
                if (2 >= 127) {
                    v5 = v5
                    v18 := []int{v2, (v2 - 9223372036854775807)}
                    _ = v18
                } else {
                    var v19 string = ("foo bar" + "\n\t")
                    var v20 string = v19
                    _ = v19
                    _ = v20
                }
            } else {
                fmt.Println(v5, v2)
                fmt.Println(v2)
            }
        default:
            if true {
                v21 := []int{(v2 & v2), v2, v2}
                v2 = v2
                _ = v21
            } else {
                fmt.Println(v2)
                v2 = -15
            }
        }
        var v22 int = v2
        delete(v4, 32767)
        _ = v22
    }
    v23 := []int{}
    var v24 int = (v2 << ((464) & 63))
    fmt.Println(v24, v5)
    var v25 string = ("hello" + "")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
