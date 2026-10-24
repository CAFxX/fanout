package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int32, p1 string) uint8 {
    v4 := GenS1{F0: ("a" + ("\n\t" + (p1 + (("hello" + "hello") + p1))))}
    var v5 uint32 = 2147483647
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 1
}
func genFn1() int8 {
    var v6 int32 = 1
    switch 32768 {
    case 2147483647:
        v6 = -511
    case 32768:
        if (!(3.14159 <= float32((((3.14159 / float64(3.14159)) - (1.5 / float64((1.5 / 1.5)))) + 1e10)))) {
            if ((!true) || true) {
                v7 := GenS1{F0: "hello"}
                _ = v7
            } else {
                v6 -= v6
                fmt.Println(v6)
            }
        }
    }
    _ = v6
    return 100
}

func main() {
    v1 := []uint32{645, 1}
    v2 := []int{}
    v3 := []int{}
    var v8 string = ("a" + ("hello" + ("a" + "hello")))
    fmt.Println(v8)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v8)
    var v9 float32 = 1.5
    var v10 bool = (v9 >= 0.1)
    switch 32767 {
    case -633:
        var v11 uint = 1
        _ = v11
    case 1:
        v9 -= v9
    default:
        if v10 {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            switch 32767 {
            case 0:
                v8 = v8
            case 397:
                for fi12 := 0; fi12 < 10; fi12++ {
                    var v13 uint16 = 782
                    _ = fi12
                    _ = v13
                }
            case 63:
                switch 127 {
                case -250:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case -1:
                    v9 += v9
                case 1:
                    fmt.Println(v9)
                default:
                    for fi14 := 0; fi14 < 18; fi14++ {
                        var v15 int32 = 0
                        v16 := GenS1{F0: v8}
                        for fi17 := 0; fi17 < 9; fi17++ {
                            if len(v3) > 0 {
                                v3[0] = (fi17 + fi17)
                            }
                            _ = fi17
                        }
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                }
            }
            if (500 > 2) {
                switch 2147483647 {
                case 128:
                    v18 := GenS1{F0: v8}
                    _ = v18
                case 2147483647:
                    switch "\n\t" {
                    case "\n\t":
                        if true {
                            fmt.Println(v10, v8)
                        }
                    case "a":
                        var v19 float32 = (v9 + (v9 + v9))
                        _ = v19
                    default:
                        v20 := GenS1{F0: (("\n\t" + (("\n\t" + ("" + v8)) + "a")) + "hello")}
                        _ = v20
                    }
                case 32767:
                    _ = len(v1)
                }
            } else {
                var v21 int64 = 42
                if len(v3) > 0 {
                    v3[0] = 0
                }
                v22 := GenS2{F0: 1}
                _ = v21
                _ = v22
            }
        }
    }
    fmt.Println(v8)
    v8 = v8
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
