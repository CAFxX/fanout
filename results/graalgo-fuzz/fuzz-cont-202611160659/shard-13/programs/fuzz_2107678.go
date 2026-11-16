package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 uint; F2 int }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]int)
    v3 := make(map[int]int)
    v4 := make(map[uint32]string)
    var v5 uint32 = 540
    var v6 uint8 = 127
    v6 = 202
    if false {
        fmt.Println(v6, v5)
        v7 := GenS1{F0: 1}
        switch "\n\t" {
        case "foo bar":
            v6 += 42
        case "":
            if true {
                v6 = (v6 | v6)
                v5 = (v5 & v5)
            }
        case "hello":
            v8 := v4[v5]
            _ = v8
        default:
            if (false || false) {
                v9 := v1[(v5 * (v5 - (v5 >> ((255) & 63))))]
                var v10 uint8 = v6
                for fi11 := 0; fi11 < 6; fi11++ {
                    if (!true) {
                        var v12 int = -2147483648
                        _ = len(v3)
                        var v13 uint16 = 0
                        _ = v12
                        _ = v13
                    }
                    var v14 bool = (-13 != -37)
                    var v15 int8 = -20
                    _ = fi11
                    _ = v14
                    _ = v15
                }
                _ = v10
                _ = v9
            }
        }
        _ = v7
    } else {
        delete(v2, 127)
        for fi16 := 0; fi16 < 3; fi16++ {
            for fi17 := 0; fi17 < 8; fi17++ {
                v5 = v5
                _ = fi17
            }
            var v18 string = ((((((("a" + "a") + (("hello" + "\n\t") + "a")) + "hello") + "a") + (("" + ("a" + "a")) + "foo bar")) + ("hello" + "\n\t")) + "hello")
            v2[128] = fi16
            _ = fi16
            _ = v18
        }
        if false {
            v19 := GenS1{F0: 42}
            for fi20 := 0; fi20 < 7; fi20++ {
                v5 += 65535
                if (46 != -128) {
                    v6 += 1
                    v5 = v5
                    for fi21 := 0; fi21 < 20; fi21++ {
                        v22 := GenS2{F0: fi20, F1: 4294967295, F2: (fi20 + 128)}
                        v23 := GenS1{F0: -1}
                        v24 := GenS1{F0: 0}
                        var v25 uint16 = 255
                        _ = fi21
                        _ = v22
                        _ = v23
                        _ = v24
                        _ = v25
                    }
                } else {
                    var v26 int8 = 65
                    var v27 float32 = 1.5
                    _ = v26
                    _ = v27
                }
                _ = fi20
            }
            _ = v19
        } else {
            v28 := v1[v5]
            _ = v28
        }
    }
    var v29 int32 = 2147483647
    switch 42 {
    case -582:
        var v30 int32 = (v29 ^ -1)
        _ = v30
    case 0:
        var v31 int = 531
        _ = v31
    default:
        var v32 bool = (true && true)
        _ = v32
    }
    v33 := []int{}
    _ = len(v3)
    for fi34 := 0; fi34 < 11; fi34++ {
        var v35 float32 = 0.1
        fmt.Println(v35, v6, fi34)
        _ = fi34
        _ = v35
    }
    var v36 int16 = -1
    for fi37 := 0; fi37 < 17; fi37++ {
        v5 = 1
        _ = fi37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v33
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
