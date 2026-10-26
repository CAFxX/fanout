package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 uint32; F1 string }
func genFn0(p0 uint, p1 uint32, p2 string) int8 {
    var v4 int32 = 42
    if false {
        for fi5 := 0; fi5 < 19; fi5++ {
            if ((854 != 32767) && true) {
                switch "foo bar" {
                case "":
                    v6 := GenS1{F0: (p1 << ((p0) & 63)), F1: 0}
                    _ = v6
                case "a":
                    var v7 uint32 = 42
                    _ = v7
                case "hello":
                    var v8 int64 = 984
                    _ = v8
                default:
                    fmt.Println(v4, fi5)
                }
            }
            _ = fi5
        }
        var v9 bool = true
        p2 = p2
        _ = v9
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 120
}

func main() {
    v1 := []uint32{2, 0, 42}
    var v2 string = "\n\t"
    v3 := []int{128, 128, 713}
    v2 = ((v2 + v2) + v2)
    if ((((65535 == 1) || true) || (!(true && ((0 == -1) || (2 >= 1))))) || (!(53 > 42))) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v2 = "\n\t"
    }
    switch -811 {
    case 734:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        for fi10 := 0; fi10 < 16; fi10++ {
            var v11 uint64 = 390
            var v12 float32 = (3.14159 / ((3.14159 - 1e10) / float32(1.5)))
            if len(v3) > 0 {
                v3[0] = (^(fi10))
            }
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    var v13 bool = true
    v2 = v2
    _ = len(v1)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v14 := GenS2{F0: 65535, F1: "foo bar"}
    v15 := GenS1{F0: 4294967295, F1: 1}
    var v16 int32 = 662
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
