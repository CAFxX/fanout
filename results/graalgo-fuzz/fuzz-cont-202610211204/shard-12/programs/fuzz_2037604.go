package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 int16 }
func genFn0(p0 int64, p1 string, p2 string) int64 {
    if false {
        p1 = "\n\t"
    } else {
        var v5 int16 = 1000
        v6 := GenS1{F0: 1, F1: 2}
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 18; fi7++ {
        if true {
            var v8 int64 = p0
            v9 := []int{581}
            _ = v8
            _ = v9
        } else {
            if true {
                v10 := []int{(-(fi7)), fi7}
                _ = v10
            }
            p2 = (p2 + p2)
            p0 += -1
        }
        v11 := GenS1{F0: 408, F1: 1}
        var v12 float32 = 0.0
        _ = fi7
        _ = v11
        _ = v12
    }
    if (0 >= 18446744073709551615) {
        p2 = (p1 + p1)
        p2 = p1
        for fi13 := 0; fi13 < 2; fi13++ {
            v14 := GenS1{F0: 42, F1: 2}
            var v15 string = "foo bar"
            v16 := []int{}
            _ = fi13
            _ = v14
            _ = v15
            _ = v16
        }
    }
    var v17 int32 = 18
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return p0
}

func main() {
    v1 := []int{877}
    v2 := make(map[uint32]string)
    var v3 int = 145
    var v4 uint64 = 902
    var v18 uint = 1
    v19 := v2[2147483647]
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v18 = 255
    var v20 uint8 = 39
    fmt.Println(v4)
    switch v19 {
    case "":
        v2[42] = v19
    case "a":
        switch (v3 - v3) {
        case 9223372036854775807:
            var v21 float64 = float64(3.14159)
            _ = v21
        case 32767:
            if (!(v19 < (((((v19 + "\n\t") + v19) + "a") + ("a" + "")) + (v19 + v19)))) {
                v22 := GenS1{F0: 214, F1: 296}
                _ = v22
            }
        case -789:
            switch v3 {
            case 2147483647:
                _ = len(v1)
            case 9223372036854775807:
                fmt.Println(v19)
            case -1:
                var v23 uint16 = 32767
                _ = v23
            }
        }
    default:
        for fi24 := 0; fi24 < 4; fi24++ {
            for fi25 := 0; fi25 < 2; fi25++ {
                v26 := GenS2{F0: -32768}
                var v27 bool = true
                _ = fi25
                _ = v26
                _ = v27
            }
            for fi28 := 0; fi28 < 8; fi28++ {
                v1 = append(v1, (-(fi24)))
                _ = fi28
            }
            _ = fi24
        }
    }
    var v29 int64 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v4
}
