package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 uint16 = 32767
    var v2 int16 = -32768
    v3 := []string{"", "a", "foo bar"}
    v4 := make(map[int]uint32)
    var v5 string = "\n\t"
    fmt.Println(v2)
    v1 = (v1 & v1)
    var v6 float64 = (-2.25 / 1.5)
    var v7 int8 = 42
    switch 4294967295 {
    case -2147483648:
        var v8 bool = (float32((0.1 * v6)) >= 0.0)
        _ = v8
    case 2147483647:
        v7 -= v7
    default:
        for fi9 := 0; fi9 < 6; fi9++ {
            var v10 int8 = 0
            _ = fi9
            _ = v10
        }
    }
    v3 = append(v3, v5)
    v6 += 1e10
    var v11 uint16 = (v1 / (((v1 % ((v1) | 1))) | 1))
    v1 -= (v11 >> ((4294967295) & 63))
    if false {
        for fi12 := 0; fi12 < 15; fi12++ {
            if (9223372036854775807 < 0) {
                for fi13 := 0; fi13 < 20; fi13++ {
                    var v14 int = (fi13 % ((fi12) | 1))
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi13
                    _ = v14
                }
                switch v5 {
                case "\n\t":
                    v1 += 0
                case "foo bar":
                    v2 = v2
                default:
                    for fi15 := 0; fi15 < 8; fi15++ {
                        v11 = (v11 & 817)
                        _ = fi15
                    }
                }
            }
            var v16 uint8 = 127
            _ = fi12
            _ = v16
        }
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v17 bool = false
    fmt.Println(v2, v11, v7)
    v11 = v11
    if ((v5 + "foo bar") != v5) {
        var v18 uint8 = 255
        var v19 int16 = -1
        if (1e10 >= float64(((0.0 - float32(v6)) + float32((v6 * v6))))) {
            fmt.Println(v1, v6)
            v20 := GenS1{F0: 9223372036854775807}
            if (((v5 + v5) > (v5 + "\n\t")) || (v6 <= v6)) {
                for fi21 := 0; fi21 < 18; fi21++ {
                    v18 = v18
                    var v22 float32 = 1.5
                    v3 = append(v3, v5)
                    _ = fi21
                    _ = v22
                }
                v23 := GenS1{F0: 1}
                var v24 bool = (!true)
                _ = v23
                _ = v24
            } else {
                if (!true) {
                    delete(v4, -34)
                    if v17 {
                        v4[32768] = 0
                    } else {
                        var v25 float64 = v6
                        _ = v25
                    }
                    fmt.Println(v17)
                }
                var v26 uint8 = v18
                _ = v26
            }
            _ = v20
        }
        _ = v18
        _ = v19
    }
    if false {
        for fi27 := 0; fi27 < 5; fi27++ {
            v6 -= (v6 + v6)
            _ = len(v3)
            v17 = (153 > 255)
            _ = fi27
        }
    }
    for fi28 := 0; fi28 < 9; fi28++ {
        var v29 string = v5
        v6 = v6
        if (664 == (v2 % (((v2 * 750)) | 1))) {
            fmt.Println(v1, v1)
            var v30 int = (-(fi28))
            switch (((v5 + v29) + (("a" + (v5 + ("a" + ""))) + (v29 + v29))) + "\n\t") {
            case "\n\t":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            _ = v30
        } else {
            var v31 string = "foo bar"
            _ = len(v4)
            v6 -= float64(1.5)
            _ = v31
        }
        _ = fi28
        _ = v29
    }
    delete(v4, 127)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
