package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 int8; F1 uint64 }
func genFn0() int8 {
    var v6 uint64 = 861
    v7 := []int{}
    if false {
        if len(v7) > 0 {
            v7[0] = 9223372036854775807
        }
    }
    v6 += (v6 ^ (v6 << ((0) & 63)))
    _ = v6
    _ = v7
    return 84
}

func main() {
    var v1 uint = 255
    var v2 float32 = 1e10
    var v3 string = "\n\t"
    var v4 int = 127
    v5 := []string{"hello"}
    switch (^(v4)) {
    case -962:
        fmt.Println(v2, v2)
    case 32768:
        var v8 int64 = -9223372036854775808
        _ = v8
    default:
        fmt.Println(v1, v4, v1)
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        fmt.Println(v3, v2)
        if ((495 != int64(fi9)) && true) {
            v3 = ("hello" + (v3 + v3))
            var v10 uint16 = 2
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = v10
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi9
    }
    switch "" {
    case "":
        if false {
            var v11 int32 = -846
            v12 := GenS1{F0: v11, F1: v3}
            _ = v11
            _ = v12
        }
    case "a":
        for fi13 := 0; fi13 < 5; fi13++ {
            v4 += v4
            _ = len(v5)
            _ = fi13
        }
    default:
        v14 := GenS1{F0: -90, F1: v3}
        _ = v14
    }
    v15 := GenS1{F0: 1, F1: (("" + "foo bar") + (v3 + ""))}
    for fi16 := 0; fi16 < 10; fi16++ {
        var v17 string = "\n\t"
        fmt.Println(v4, fi16)
        _ = fi16
        _ = v17
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v4 = (v4 + 4294967295)
    if false {
        var v18 bool = (!(42 >= 1))
        if (v2 <= (((v2 / ((((v2 + v2) / 1.0) - 0.0) + (0.0 - 1e10))) + v2) * v2)) {
            for fi19 := 0; fi19 < 11; fi19++ {
                var v20 int8 = genFn0()
                _ = fi19
                _ = v20
            }
            var v21 int16 = 1
            _ = v21
        } else {
            for fi22 := 0; fi22 < 6; fi22++ {
                var v23 int8 = 104
                var v24 float32 = (1.5 * v2)
                _ = fi22
                _ = v23
                _ = v24
            }
            v25 := GenS1{F0: 1, F1: (v3 + (v3 + "hello"))}
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = v25
        }
        _ = v18
    } else {
        if false {
            v2 = v2
        }
        var v26 int16 = 32767
        _ = v26
    }
    v3 = v3
    fmt.Println(v4, v3, v2)
    for fi27 := 0; fi27 < 17; fi27++ {
        v28 := GenS2{F0: 0, F1: 0}
        _ = fi27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
