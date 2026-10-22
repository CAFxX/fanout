package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 int32) int8 {
    p0 -= p0
    switch 0 {
    case 9223372036854775807:
        var v7 int64 = 0
        _ = v7
    }
    p0 = 492
    var v8 uint = 965
    v9 := []int{-9223372036854775808, -1}
    _ = p0
    _ = v8
    _ = v9
    return -1
}
func genFn1() uint64 {
    var v10 bool = (-2.25 != 3.14159)
    switch 127 {
    case 42:
        v10 = (1 < -1)
    default:
        var v11 int32 = 453
        _ = v11
    }
    fmt.Println(v10)
    for fi12 := 0; fi12 < 12; fi12++ {
        for fi13 := 0; fi13 < 15; fi13++ {
            var v14 uint64 = 42
            _ = fi13
            _ = v14
        }
        var v15 int = fi12
        var v16 bool = (18446744073709551615 == 2)
        _ = fi12
        _ = v15
        _ = v16
    }
    _ = v10
    return 2
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint64 = 796
    var v3 int = 32767
    var v4 uint32 = 40
    var v5 uint = 18446744073709551615
    var v6 int32 = 0
    v2 = 696
    switch 127 {
    case 32767:
        v3 = 1
    case 286:
        var v17 uint64 = (v2 >> ((516) & 63))
        _ = v17
    default:
        v18 := []int{(v3 | 2147483647)}
        _ = v18
    }
    switch (("hello" + ("foo bar" + "a")) + "a") {
    case "\n\t":
        var v19 uint64 = v2
        _ = v19
    case "a":
        if false {
            for fi20 := 0; fi20 < 12; fi20++ {
                v6 = (v6 * (v6 % ((v6) | 1)))
                switch v3 {
                case 32767:
                    v4 += (v4 & (v4 - (v4 / ((uint32(v2)) | 1))))
                default:
                    if false {
                        v21 := GenS1{F0: genFn1()}
                        v2 = (v2 + (v2 | v2))
                        for fi22 := 0; fi22 < 18; fi22++ {
                            switch -104 {
                            case 2147483647:
                                var v23 string = ("hello" + ("\n\t" + ("\n\t" + ("a" + ((("hello" + (((("" + "a") + ("" + "\n\t")) + ("" + (((("foo bar" + ((("hello" + ("foo bar" + ("" + ("a" + ("hello" + "foo bar"))))) + "") + "foo bar")) + "") + "foo bar") + ("hello" + "hello")))) + "")) + ((((("a" + "\n\t") + "") + "\n\t") + "hello") + "hello")) + "\n\t")))))
                                _ = v23
                            default:
                                v4 = v4
                            }
                            v4 = v4
                            var v24 int = (fi20 & v3)
                            _ = fi22
                            _ = v24
                        }
                        _ = v21
                    }
                }
                _ = fi20
            }
            v6 -= (v6 + v6)
            for fi25 := 0; fi25 < 20; fi25++ {
                v2 += (v2 & v2)
                _ = fi25
            }
        } else {
            switch "" {
            case "a":
                v1[v3] = v4
            case "":
                var v26 uint64 = uint64(v4)
                _ = v26
            }
        }
    default:
        var v27 int64 = -9223372036854775808
        _ = v27
    }
    v2 += (v2 + 18446744073709551615)
    var v28 uint8 = 48
    var v29 float64 = 3.14159
    v29 = ((float64(((3.14159 + float32(1.5)) - 1e10)) / float64((0.0 + 1.5))) * v29)
    v30 := GenS1{F0: v2}
    if (!true) {
        var v31 uint64 = (v2 % (((^(v2))) | 1))
        _ = v31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
