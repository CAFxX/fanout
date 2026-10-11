package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
func genFn0(p0 int32) uint32 {
    var v5 uint64 = 689
    fmt.Println(v5)
    v6 := GenS1{F0: 0, F1: 1}
    _ = p0
    _ = v5
    _ = v6
    return 4294967295
}
func genFn1() uint {
    var v7 float32 = 3.14159
    fmt.Println(v7)
    _ = v7
    return 0
}

func main() {
    var v1 uint64 = 1
    v2 := make(map[uint32]bool)
    var v3 string = (((("" + "") + ("a" + "a")) + "a") + "a")
    v4 := make(map[string]int)
    fmt.Println(v1, v1)
    if (true && true) {
        var v8 float64 = float64(0.1)
        switch 559 {
        case 127:
            v3 = "hello"
        case -376:
            var v9 int32 = -1
            _ = v9
        default:
            v3 = v3
        }
        _ = v8
    } else {
        fmt.Println(v1, v1)
        switch v3 {
        case "":
            _ = len(v2)
        }
    }
    switch 127 {
    case -2147483648:
        if true {
            fmt.Println(v1, v3)
            switch ((v3 + "hello") + (v3 + (v3 + "\n\t"))) {
            case "a":
                var v10 uint64 = v1
                _ = v10
            case "foo bar":
                var v11 bool = ((!(((255 == 128) || (!(!(!(true || ((-988 < 932) && (!true))))))) && true)) || false)
                _ = v11
            default:
                for fi12 := 0; fi12 < 5; fi12++ {
                    for fi13 := 0; fi13 < 20; fi13++ {
                        fmt.Println(v3, fi12, v1)
                        _ = fi13
                    }
                    v3 = "hello"
                    _ = fi12
                }
            }
        }
    }
    v14 := []int{}
    v15 := GenS1{F0: 127, F1: 440}
    var v16 uint32 = 423
    fmt.Println(v16, v1, v1)
    v3 = "\n\t"
    v17 := v4["\n\t"]
    _ = len(v14)
    v18 := GenS1{F0: 21, F1: 32767}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
