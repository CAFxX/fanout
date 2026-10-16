package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 bool }
type GenS2 struct { F0 int; F1 uint32 }
func genFn0(p0 int) uint {
    v5 := []int{(-(p0)), 32768, p0, p0}
    var v6 uint = 65535
    var v7 int8 = 0
    var v8 int64 = 1
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return v6
}

func main() {
    v1 := make(map[int]uint32)
    var v2 bool = (("hello" + "") != ("\n\t" + "hello"))
    var v3 string = (("\n\t" + "hello") + "hello")
    v4 := make(map[string]bool)
    v3 = ""
    var v9 string = ("foo bar" + v3)
    v2 = v2
    switch -515 {
    case 1:
        for fi10 := 0; fi10 < 11; fi10++ {
            var v11 int8 = 95
            v12 := GenS1{F0: fi10, F1: v2, F2: ((!(1 < 2147483647)) || (v2 && (("hello" + v9) != v3)))}
            for fi13 := 0; fi13 < 9; fi13++ {
                var v14 int8 = v11
                var v15 uint8 = 128
                _ = fi13
                _ = v14
                _ = v15
            }
            _ = fi10
            _ = v11
            _ = v12
        }
    case -796:
        for fi16 := 0; fi16 < 18; fi16++ {
            v9 = v9
            v3 = "a"
            var v17 uint32 = 4294967295
            _ = fi16
            _ = v17
        }
    case 9223372036854775807:
        for fi18 := 0; fi18 < 16; fi18++ {
            for fi19 := 0; fi19 < 12; fi19++ {
                var v20 int8 = 0
                switch v9 {
                case "a":
                    var v21 uint = 18446744073709551615
                    _ = v21
                case "hello":
                    fmt.Println(fi18)
                case "":
                    var v22 uint32 = 295
                    _ = v22
                default:
                    var v23 int8 = v20
                    _ = v23
                }
                v24 := GenS1{F0: fi18, F1: (65535 != 535), F2: ((65535 != 692) || (151 <= 42))}
                _ = fi19
                _ = v20
                _ = v24
            }
            var v25 bool = false
            _ = fi18
            _ = v25
        }
    default:
        v26 := GenS1{F0: 0, F1: v2, F2: (171 >= 128)}
        _ = v26
    }
    v27 := GenS1{F0: -1, F1: v2, F2: (!true)}
    var v28 int8 = -23
    var v29 int = -845
    v30 := GenS2{F0: v29, F1: 456}
    var v31 int16 = -1
    v28 = 1
    var v32 string = (v3 + "")
    var v33 uint8 = 2
    if (v2 || v2) {
        var v34 uint16 = 982
        _ = v34
    } else {
        var v35 float32 = (1.5 + 3.14159)
        v36 := []int{}
        _ = v35
        _ = v36
    }
    v29 = (v29 * v29)
    if (v2 || (v2 || (42 >= 0))) {
        v37 := []int{(v29 ^ v29)}
        var v38 uint8 = (^(v33))
        var v39 float32 = 0.0
        _ = v37
        _ = v38
        _ = v39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v9
}
