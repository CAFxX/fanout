package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 string }
func genFn0() int8 {
    v5 := []int{-951}
    if false {
        v6 := GenS1{F0: 1}
        var v7 uint64 = 18446744073709551615
        if len(v5) > 0 {
            v5[0] = 128
        }
        _ = v6
        _ = v7
    }
    var v8 int8 = 3
    _ = v5
    _ = v8
    return v8
}

func main() {
    var v1 int16 = -25
    var v2 uint32 = 42
    var v3 int8 = 0
    v4 := make(map[string]uint32)
    var v9 string = ("" + "a")
    v9 = (v9 + v9)
    for fi10 := 0; fi10 < 1; fi10++ {
        var v11 uint32 = (v2 / (((v2 / ((732) | 1))) | 1))
        switch "hello" {
        case "hello":
            var v12 string = v9
            _ = v12
        case "\n\t":
            v11 = v11
        case "foo bar":
            for fi13 := 0; fi13 < 1; fi13++ {
                if (0.1 <= (0.1 / (1e10 - 0.0))) {
                    v14 := GenS1{F0: fi10}
                    _ = v14
                }
                v2 = v11
                var v15 int = (fi10 >> ((515) & 63))
                _ = fi13
                _ = v15
            }
        default:
            switch ((v9 + v9) + ("hello" + v9)) {
            case "a":
                var v16 int8 = v3
                _ = v16
            }
        }
        var v17 uint8 = 223
        _ = fi10
        _ = v11
        _ = v17
    }
    switch 2147483647 {
    case 180:
        for fi18 := 0; fi18 < 7; fi18++ {
            fmt.Println(v2, fi18, v1)
            if (true || ((^(fi18)) >= (fi18 % ((fi18) | 1)))) {
                if (!(!(!(3.14159 >= float64(1.5))))) {
                    var v19 uint32 = v2
                    var v20 int16 = (v1 - (v1 << ((65535) & 63)))
                    var v21 uint = 18446744073709551615
                    _ = v19
                    _ = v20
                    _ = v21
                }
            } else {
                for fi22 := 0; fi22 < 10; fi22++ {
                    v23 := []int{fi22, fi22, 2147483647}
                    fmt.Println(fi22, v1)
                    var v24 float32 = float32(((3.14159 / 0.1) - 0.1))
                    _ = fi22
                    _ = v23
                    _ = v24
                }
            }
            _ = fi18
        }
    case 128:
        v25 := v4[(v9 + v9)]
        _ = v25
    case -478:
        v26 := GenS1{F0: 32768}
        _ = v26
    default:
        v9 = ("a" + "foo bar")
    }
    for fi27 := 0; fi27 < 13; fi27++ {
        var v28 string = v9
        v28 = (v9 + (v9 + v9))
        v28 = ((v28 + v28) + (v9 + v9))
        _ = fi27
        _ = v28
    }
    for fi29 := 0; fi29 < 18; fi29++ {
        if true {
            v30 := []int{(-(fi29)), fi29}
            _ = v30
        } else {
            v3 = v3
            v9 = ("\n\t" + "")
        }
        var v31 int8 = v3
        _ = fi29
        _ = v31
    }
    var v32 uint64 = 4294967295
    v33 := GenS2{F0: 0, F1: (v9 + "\n\t")}
    v34 := []int{267}
    _ = len(v4)
    var v35 float32 = float32(-2.25)
    v36 := GenS2{F0: 42, F1: v9}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v9
}
