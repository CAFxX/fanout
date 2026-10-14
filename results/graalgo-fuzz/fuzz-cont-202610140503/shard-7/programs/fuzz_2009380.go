package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 string }
func genFn0(p0 uint, p1 int32) uint8 {
    var v6 int64 = 853
    v7 := GenS1{F0: "hello"}
    if false {
        var v8 uint8 = 127
        _ = v8
    }
    for fi9 := 0; fi9 < 19; fi9++ {
        v10 := GenS1{F0: "hello"}
        var v11 bool = true
        p1 = p1
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 255
}

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = (2 < 42)
    var v3 int8 = -42
    var v4 uint64 = 811
    v5 := make(map[string]string)
    v3 = 0
    var v12 float32 = (0.1 * 0.0)
    if true {
        for fi13 := 0; fi13 < 9; fi13++ {
            v14 := GenS2{F0: v2, F1: "hello"}
            delete(v1, 1)
            _ = fi13
            _ = v14
        }
        v12 = (v12 + (v12 + ((v12 - (v12 - (v12 - (((((1.5 - 1.5) * (v12 * v12)) - v12) - 0.1) / v12)))) + 1.5)))
    }
    var v15 int32 = -2147483648
    v3 -= -128
    v16 := GenS1{F0: "a"}
    if (!((-2.25 >= (1.5 - 0.1)) && (v2 && false))) {
        switch -683 {
        case 262:
            v12 = v12
        case 127:
            delete(v1, 1)
        }
        for fi17 := 0; fi17 < 7; fi17++ {
            var v18 string = "\n\t"
            for fi19 := 0; fi19 < 20; fi19++ {
                var v20 string = ((v18 + v18) + (v18 + ""))
                var v21 int8 = (v3 * 1)
                var v22 bool = (32767 <= 72)
                _ = fi19
                _ = v20
                _ = v21
                _ = v22
            }
            switch ((v18 + "\n\t") + "\n\t") {
            case "a":
                v23 := GenS1{F0: "hello"}
                _ = v23
            default:
                v24 := GenS1{F0: ((v18 + (v18 + "\n\t")) + v18)}
                _ = v24
            }
            _ = fi17
            _ = v18
        }
    } else {
        var v25 uint = 255
        _ = v25
    }
    var v26 uint32 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
