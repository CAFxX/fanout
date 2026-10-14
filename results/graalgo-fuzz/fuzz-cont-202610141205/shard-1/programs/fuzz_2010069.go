package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 uint8 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]uint32)
    v3 := []uint32{784, 1, 863}
    var v4 string = ("a" + "a")
    var v5 string = v4
    v5 = v4
    v5 = v4
    switch 32768 {
    case 32767:
        v6 := GenS1{F0: -128, F1: (-819 > 32767)}
        _ = v6
    case 2147483647:
        v7 := GenS2{F0: 2}
        _ = v7
    default:
        for fi8 := 0; fi8 < 12; fi8++ {
            for fi9 := 0; fi9 < 2; fi9++ {
                fmt.Println(fi8, fi8)
                v5 = "foo bar"
                _ = fi9
            }
            v10 := GenS2{F0: 0}
            _ = fi8
            _ = v10
        }
    }
    if (!(3.14159 == 0.0)) {
        v4 = (v5 + v5)
    } else {
        for fi11 := 0; fi11 < 2; fi11++ {
            switch (("hello" + "hello") + "hello") {
            case "foo bar":
                var v12 int16 = -32768
                _ = v12
            case "":
                fmt.Println(fi11)
            case "\n\t":
                delete(v1, v4)
            default:
                if false {
                    _ = len(v1)
                }
            }
            for fi13 := 0; fi13 < 1; fi13++ {
                var v14 string = v4
                if (fi11 >= 2147483647) {
                    if (0 != 0) {
                        v15 := v1[v14]
                        v16 := GenS2{F0: 157}
                        _ = v15
                        _ = v16
                    }
                    v14 = v5
                }
                var v17 uint8 = 106
                _ = fi13
                _ = v14
                _ = v17
            }
            for fi18 := 0; fi18 < 8; fi18++ {
                _ = len(v3)
                var v19 int = fi18
                _ = fi18
                _ = v19
            }
            _ = fi11
        }
        v3 = append(v3, 971)
        var v20 uint = 42
        _ = v20
    }
    var v21 int16 = 489
    fmt.Println(v5, v5)
    var v22 float64 = 0.1
    for fi23 := 0; fi23 < 10; fi23++ {
        fmt.Println(fi23, v21)
        v22 = v22
        _ = fi23
    }
    v21 -= v21
    v22 = v22
    v24 := GenS2{F0: 211}
    v4 = ((("foo bar" + v4) + ((v5 + "hello") + (v4 + ("" + v5)))) + "\n\t")
    for fi25 := 0; fi25 < 16; fi25++ {
        for fi26 := 0; fi26 < 12; fi26++ {
            v3 = append(v3, 404)
            var v27 float64 = v22
            _ = fi26
            _ = v27
        }
        _ = fi25
    }
    if ((253 >= 2) || false) {
        var v28 float64 = v22
        for fi29 := 0; fi29 < 14; fi29++ {
            v30 := GenS1{F0: 127, F1: ((-32768 > (^(v21))) || (42 >= 1))}
            _ = fi29
            _ = v30
        }
        _ = v28
    }
    v5 = ((v5 + ((v4 + v4) + "")) + (("a" + "\n\t") + v5))
    switch ("" + v4) {
    case "a":
        var v31 uint = 255
        _ = v31
    case "foo bar":
        _ = len(v3)
    case "\n\t":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        var v32 int16 = (v21 * (v21 << ((4294967295) & 63)))
        _ = v32
    }
    var v33 int32 = -1
    fmt.Println(v33, v33, v33)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v24
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
