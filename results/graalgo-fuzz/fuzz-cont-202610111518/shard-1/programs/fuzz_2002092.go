package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 uint8 }
type GenS2 struct { F0 int8; F1 int32; F2 uint8 }
func genFn0(p0 uint8) string {
    fmt.Println(p0)
    for fi5 := 0; fi5 < 15; fi5++ {
        if (("a" + (("" + "hello") + ("hello" + (("" + "\n\t") + "\n\t")))) == "foo bar") {
            v6 := []int{fi5, fi5}
            _ = v6
        } else {
            v7 := GenS1{F0: (fi5 | fi5), F1: (p0 >> ((930) & 63)), F2: p0}
            var v8 float64 = (-2.25 / 1.0)
            _ = v7
            _ = v8
        }
        p0 -= p0
        p0 -= (p0 << ((65535) & 63))
        _ = fi5
    }
    fmt.Println(p0)
    _ = p0
    return ("hello" + "")
}
func genFn1(p0 int8) int16 {
    var v9 int64 = 42
    v10 := GenS1{F0: 128, F1: 229, F2: 61}
    _ = p0
    _ = v10
    _ = v9
    return 42
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint32 = 0
    var v3 int64 = 42
    v4 := make(map[int]string)
    if true {
        if (18446744073709551615 != 367) {
            if (!true) {
                v3 = v3
            }
        }
        v3 += v3
        if true {
            v2 = v2
            v3 += v3
        }
    } else {
        var v11 bool = ((true || true) || true)
        v11 = false
        if false {
            v12 := v1["hello"]
            _ = v12
        }
        _ = v11
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        var v14 uint64 = 1
        var v15 bool = false
        for fi16 := 0; fi16 < 2; fi16++ {
            v15 = ((^(v14)) >= (v14 % (((v14 | (v14 / (((^(v14))) | 1)))) | 1)))
            _ = fi16
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    v3 = v3
    v17 := GenS1{F0: 128, F1: 123, F2: 128}
    fmt.Println(v3, v3)
    v3 = (v3 | v3)
    v18 := GenS1{F0: -853, F1: uint8(v2), F2: 255}
    v19 := GenS2{F0: 100, F1: -1, F2: 255}
    v3 += v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
