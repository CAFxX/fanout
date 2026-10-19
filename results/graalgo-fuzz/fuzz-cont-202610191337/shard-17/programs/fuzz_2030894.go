package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 int16; F1 uint8 }

func main() {
    var v1 int = 128
    var v2 uint64 = 4294967295
    var v3 uint64 = v2
    var v4 string = "foo bar"
    v5 := []int{}
    v6 := []uint32{65535}
    var v7 uint64 = 1
    var v8 uint8 = 112
    v4 = v4
    for fi9 := 0; fi9 < 3; fi9++ {
        v10 := GenS1{F0: ((v4 + ((v4 + v4) + (("\n\t" + ("" + ((v4 + "a") + ((v4 + ((("\n\t" + v4) + ((v4 + "") + ("" + ""))) + ("hello" + "hello"))) + "hello")))) + v4))) + ""), F1: v8}
        for fi11 := 0; fi11 < 7; fi11++ {
            v12 := GenS1{F0: v4, F1: (^(v8))}
            v13 := GenS1{F0: "a", F1: v8}
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            if ((42 == 0) || (!false)) {
                var v14 float64 = 1.5
                var v15 float32 = ((0.0 / (0.0 + (1.5 / 0.1))) * 3.14159)
                _ = v14
                _ = v15
            }
            _ = len(v6)
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi9
        _ = v10
    }
    var v16 int64 = 42
    fmt.Println(v8)
    v2 = v3
    for fi17 := 0; fi17 < 18; fi17++ {
        v1 = fi17
        switch v4 {
        case "":
            v18 := GenS1{F0: ("hello" + (v4 + (v4 + v4))), F1: v8}
            _ = v18
        default:
            v19 := GenS1{F0: v4, F1: (v8 - (v8 * v8))}
            _ = v19
        }
        _ = fi17
    }
    fmt.Println(v1, v4, v1)
    var v20 int16 = 42
    if false {
        fmt.Println(v7, v3)
    }
    switch (v1 ^ v1) {
    case 1:
        for fi21 := 0; fi21 < 1; fi21++ {
            _ = len(v6)
            v22 := GenS1{F0: ("a" + (v4 + (v4 + ("hello" + "\n\t")))), F1: (^(v8))}
            _ = fi21
            _ = v22
        }
    case 9223372036854775807:
        if false {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            if true {
                var v23 int64 = v16
                v24 := GenS1{F0: "foo bar", F1: v8}
                _ = v23
                _ = v24
            } else {
                for fi25 := 0; fi25 < 19; fi25++ {
                    for fi26 := 0; fi26 < 4; fi26++ {
                        v3 = (v7 ^ v3)
                        var v27 float64 = -2.25
                        _ = fi26
                        _ = v27
                    }
                    _ = fi25
                }
                for fi28 := 0; fi28 < 19; fi28++ {
                    v29 := GenS1{F0: v4, F1: v8}
                    _ = fi28
                    _ = v29
                }
                for fi30 := 0; fi30 < 6; fi30++ {
                    for fi31 := 0; fi31 < 16; fi31++ {
                        var v32 uint64 = (^(v2))
                        var v33 int32 = -764
                        _ = fi31
                        _ = v32
                        _ = v33
                    }
                    fmt.Println(v2, v3, v2)
                    _ = fi30
                }
            }
        }
    case 2147483647:
        var v34 float32 = 0.0
        _ = v34
    default:
        v35 := GenS2{F0: -787, F1: v8}
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
