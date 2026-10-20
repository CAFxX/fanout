package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 int }
type GenS2 struct { F0 uint; F1 string }
func genFn0(p0 uint, p1 int64, p2 string) int64 {
    var v6 int8 = 0
    var v7 int64 = (p1 >> ((0) & 63))
    if false {
        var v8 bool = (((0.0 < 0.0) || ((p2 + p2) == "")) && false)
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return v7
}

func main() {
    v1 := []string{"hello"}
    var v2 uint8 = 106
    v3 := make(map[int]uint32)
    var v4 uint16 = 1
    var v5 bool = (false || (18446744073709551615 == 65535))
    for fi9 := 0; fi9 < 11; fi9++ {
        for fi10 := 0; fi10 < 9; fi10++ {
            fmt.Println(fi10, v2)
            switch ("hello" + "\n\t") {
            case "":
                var v11 int32 = -1
                _ = v11
            }
            switch "a" {
            case "foo bar":
                v12 := GenS1{F0: 2147483647, F1: 2147483647, F2: (-(fi9))}
                _ = v12
            default:
                _ = len(v1)
            }
            _ = fi10
        }
        _ = fi9
    }
    switch 4294967295 {
    case 128:
        if false {
            var v13 string = ""
            _ = v13
        }
    case 9223372036854775807:
        v2 -= v2
    case 2147483647:
        v4 = v4
    }
    var v14 uint32 = 1
    v15 := GenS1{F0: 9223372036854775807, F1: 1, F2: 9223372036854775807}
    for fi16 := 0; fi16 < 4; fi16++ {
        var v17 int = fi16
        _ = fi16
        _ = v17
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v18 uint = 18446744073709551615
    fmt.Println(v4, v4, v18)
    if (!((1.5 - -2.25) >= 1.5)) {
        for fi19 := 0; fi19 < 20; fi19++ {
            if (v4 < (v4 - (v4 | 32767))) {
                for fi20 := 0; fi20 < 11; fi20++ {
                    var v21 uint64 = 0
                    _ = fi20
                    _ = v21
                }
                var v22 int8 = 100
                v22 = -1
                _ = v22
            } else {
                var v23 string = "\n\t"
                _ = v23
            }
            if false {
                for fi24 := 0; fi24 < 13; fi24++ {
                    if false {
                        v14 -= (v14 * (v14 + uint32(v4)))
                        switch "\n\t" {
                        case "":
                            v25 := v3[591]
                            _ = v25
                        case "a":
                            var v26 string = ""
                            _ = v26
                        case "hello":
                            v27 := GenS1{F0: int(v4), F1: 1, F2: (^(fi19))}
                            _ = v27
                        }
                    } else {
                        var v28 float32 = 1.5
                        v29 := GenS2{F0: (v18 & 1), F1: ("\n\t" + (("hello" + "foo bar") + "\n\t"))}
                        _ = v28
                        _ = v29
                    }
                    if v5 {
                        v5 = v5
                        v5 = false
                        v30 := v3[(^(fi19))]
                        _ = v30
                    }
                    fmt.Println(v2, v2)
                    _ = fi24
                }
            }
            _ = fi19
        }
    }
    var v31 float32 = 0.1
    v18 = v18
    v2 += v2
    switch 128 {
    case 4294967295:
        v5 = (1 != (v4 >> ((0) & 63)))
    }
    if len(v1) > 0 {
        v1[0] = ("a" + "foo bar")
    }
    v5 = v5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
