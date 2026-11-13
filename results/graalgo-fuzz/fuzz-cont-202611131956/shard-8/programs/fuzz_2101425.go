package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int8 }

func main() {
    v1 := []int{1, 845}
    var v2 uint = 42
    v3 := []uint32{65535, 42, 4294967295, 21}
    var v4 bool = true
    var v5 int = 0
    fmt.Println(v2, v4, v4)
    fmt.Println(v4, v4)
    switch -637 {
    case 4294967295:
        v4 = (true || (true || v4))
    case 128:
        var v6 string = ""
        _ = v6
    case -1:
        if v4 {
            if (v4 || (!false)) {
                v7 := GenS1{F0: 0, F1: 128, F2: 1}
                fmt.Println(v4, v2)
                var v8 uint64 = 1
                _ = v7
                _ = v8
            } else {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                switch ((("\n\t" + "foo bar") + "") + "") {
                case "foo bar":
                    v4 = (42 != 111)
                case "a":
                    for fi9 := 0; fi9 < 16; fi9++ {
                        if v4 {
                            var v10 string = "\n\t"
                            _ = v10
                        } else {
                            v11 := GenS1{F0: 0, F1: 85, F2: 127}
                            _ = v11
                        }
                        fmt.Println(v4, v5, fi9)
                        _ = fi9
                    }
                case "":
                    fmt.Println(v5)
                }
                if v4 {
                    for fi12 := 0; fi12 < 8; fi12++ {
                        var v13 float32 = (0.1 / float32(1.5))
                        v2 += 42
                        _ = fi12
                        _ = v13
                    }
                } else {
                    if len(v3) > 0 {
                        v3[0] = 65535
                    }
                    v4 = v4
                    fmt.Println(v4, v5, v2)
                }
            }
        }
    default:
        var v14 uint = 171
        _ = v14
    }
    v5 += (v5 % ((v5) | 1))
    v1 = append(v1, v5)
    for fi15 := 0; fi15 < 15; fi15++ {
        fmt.Println(v4, v2, v2)
        v16 := GenS1{F0: 127, F1: 221, F2: int8(v2)}
        _ = fi15
        _ = v16
    }
    fmt.Println(v4, v5)
    v4 = ("" == "a")
    if v4 {
        v5 -= (v5 + (v5 - (-(v5))))
    }
    for fi17 := 0; fi17 < 10; fi17++ {
        for fi18 := 0; fi18 < 11; fi18++ {
            _ = len(v1)
            if len(v1) > 0 {
                v1[0] = fi17
            }
            _ = fi18
        }
        _ = fi17
    }
    var v19 float64 = -2.25
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
