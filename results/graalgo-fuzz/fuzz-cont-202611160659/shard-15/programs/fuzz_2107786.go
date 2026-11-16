package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]string)
    var v3 bool = (((!false) || false) || (-2147483648 == 9223372036854775807))
    v4 := make(map[int]bool)
    var v5 int32 = 0
    switch "\n\t" {
    case "\n\t":
        var v6 uint = 819
        _ = v6
    default:
        v3 = (-675 < -1)
    }
    if (v3 && v3) {
        v7 := []int{0, 127, -908, 922}
        if (!((1e10 - float32(0.1)) >= 0.0)) {
            var v8 bool = (1 != 830)
            v3 = v8
            _ = v8
        } else {
            v3 = (!v3)
            switch 0 {
            case 32767:
                for fi9 := 0; fi9 < 15; fi9++ {
                    v10 := GenS1{F0: 42}
                    for fi11 := 0; fi11 < 8; fi11++ {
                        v12 := v4[fi9]
                        fmt.Println(v5)
                        var v13 int32 = (-(v5))
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                    v3 = (18446744073709551615 > 18446744073709551615)
                    _ = fi9
                    _ = v10
                }
            case -266:
                v3 = (!(!v3))
            case -161:
                if (true || ((!(65535 > 255)) && (-2.25 != float64(-2.25)))) {
                    var v14 float64 = float64(((-2.25 - 1e10) / 0.1))
                    var v15 string = ("hello" + "hello")
                    _ = v14
                    _ = v15
                } else {
                    v16 := GenS1{F0: 42}
                    var v17 int32 = (v5 << ((0) & 63))
                    _ = len(v7)
                    _ = v16
                    _ = v17
                }
            }
        }
        _ = v7
    }
    v18 := v1[65535]
    for fi19 := 0; fi19 < 8; fi19++ {
        v5 = v5
        _ = fi19
    }
    v20 := []int{}
    if (!(65535 != 18446744073709551615)) {
        v18 = v18
        v21 := GenS1{F0: 65535}
        if (!(!(v3 || (248 <= -251)))) {
            v18 = 65535
        }
        _ = v21
    } else {
        _ = len(v2)
    }
    var v22 uint16 = 42
    var v23 uint64 = 18446744073709551615
    _ = len(v20)
    for fi24 := 0; fi24 < 20; fi24++ {
        fmt.Println(v22, v18)
        _ = fi24
    }
    var v25 int32 = (v5 & 0)
    for fi26 := 0; fi26 < 13; fi26++ {
        v23 += 199
        _ = fi26
    }
    v27 := GenS1{F0: (v22 >> ((255) & 63))}
    var v28 float64 = (-2.25 / 3.14159)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v25
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
