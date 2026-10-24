package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int8 }
type GenS2 struct { F0 int64; F1 int16 }

func main() {
    var v1 string = "hello"
    var v2 uint16 = 65535
    v3 := make(map[uint32]bool)
    var v4 uint8 = 208
    v4 += (v4 - v4)
    v4 = (v4 >> ((2) & 63))
    v5 := GenS1{F0: -1, F1: "\n\t", F2: -125}
    v6 := GenS1{F0: -210, F1: v1, F2: 45}
    var v7 string = v1
    v3[2147483647] = false
    v8 := GenS2{F0: 9223372036854775807, F1: 940}
    fmt.Println(v2)
    for fi9 := 0; fi9 < 17; fi9++ {
        delete(v3, 2147483648)
        _ = fi9
    }
    switch (((((v7 + v1) + "") + v1) + ((("\n\t" + "hello") + ("hello" + v7)) + "a")) + v1) {
    case "":
        delete(v3, 872)
    case "hello":
        if ((v7 + v1) <= v1) {
            v1 = v7
        } else {
            for fi10 := 0; fi10 < 18; fi10++ {
                switch fi10 {
                case -459:
                    for fi11 := 0; fi11 < 15; fi11++ {
                        if (2147483647 <= -2147483648) {
                            v7 = ((v7 + v7) + v7)
                            var v12 int = fi10
                            fmt.Println(v7, fi10)
                            _ = v12
                        } else {
                            switch v7 {
                            case "a":
                                v2 = v2
                            case "hello":
                                v13 := []int{fi11, (fi10 * (fi11 << ((uint(v2)) & 63)))}
                                _ = v13
                            default:
                                fmt.Println(v2, v1)
                            }
                        }
                        var v14 float32 = 0.0
                        var v15 bool = (!((775 > uint(v2)) && true))
                        _ = fi11
                        _ = v14
                        _ = v15
                    }
                case 4294967295:
                    fmt.Println(fi10, v4)
                default:
                    for fi16 := 0; fi16 < 19; fi16++ {
                        var v17 uint16 = 32767
                        v2 = 32767
                        _ = fi16
                        _ = v17
                    }
                }
                _ = fi10
            }
            var v18 uint = 111
            _ = v18
        }
    default:
        v1 = ("foo bar" + v7)
    }
    switch 187 {
    case 9223372036854775807:
        v1 = v7
    default:
        v1 = "foo bar"
    }
    var v19 float32 = 1.5
    v2 = v2
    for fi20 := 0; fi20 < 2; fi20++ {
        v21 := GenS1{F0: -2147483648, F1: (v7 + (v1 + ((v7 + v1) + v7))), F2: -1}
        _ = fi20
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
