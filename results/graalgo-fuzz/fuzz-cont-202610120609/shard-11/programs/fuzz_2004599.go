package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string }

func main() {
    var v1 uint = 688
    v2 := make(map[uint32]uint32)
    var v3 string = "a"
    v4 := []uint32{2147483647, uint32(v1), 290, 362}
    var v5 int8 = 42
    v6 := []string{v3}
    if (true && false) {
        switch 344 {
        case -513:
            v5 -= (-(v5))
        case 750:
            for fi7 := 0; fi7 < 18; fi7++ {
                switch v3 {
                case "a":
                    var v8 int32 = 0
                    _ = v8
                default:
                    var v9 int8 = v5
                    _ = v9
                }
                for fi10 := 0; fi10 < 9; fi10++ {
                    switch "hello" {
                    case "":
                        v3 = v3
                    case "hello":
                        switch 32768 {
                        case 4294967295:
                            var v11 bool = ((!(0.0 != ((3.14159 + ((3.14159 / (1.5 / 1.0)) / (3.14159 * 3.14159))) * (((3.14159 / 0.1) * (3.14159 - 3.14159)) - (float64(1e10) + 1.5))))) || (127 > 1))
                            _ = v11
                        case 127:
                            v12 := GenS1{F0: 77}
                            _ = v12
                        case -751:
                            v3 = "hello"
                        default:
                            v3 = (v3 + v3)
                        }
                    }
                    _ = fi10
                }
                _ = fi7
            }
        }
        if (1.5 != 3.14159) {
            _ = len(v4)
            var v13 string = ""
            fmt.Println(v13, v1)
            _ = v13
        } else {
            v1 = v1
        }
    }
    var v14 string = v3
    for fi15 := 0; fi15 < 13; fi15++ {
        v16 := GenS2{F0: (v3 + v14)}
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 6; fi17++ {
        var v18 uint8 = 117
        _ = fi17
        _ = v18
    }
    v19 := GenS2{F0: ("\n\t" + (("" + v3) + v3))}
    for fi20 := 0; fi20 < 17; fi20++ {
        switch (("foo bar" + "") + ((v3 + "") + (v3 + (v3 + ((((v14 + "") + "") + ((v3 + (((v3 + (v3 + (v3 + v14))) + v14) + "foo bar")) + "foo bar")) + ((v3 + ("hello" + (v3 + "a"))) + (v3 + "hello"))))))) {
        case "":
            var v21 uint32 = 2
            _ = v21
        case "a":
            var v22 string = "hello"
            _ = v22
        default:
            fmt.Println(v1, fi20)
        }
        var v23 bool = ((fi20 & 32767) != fi20)
        _ = fi20
        _ = v23
    }
    if (1.5 > 1e10) {
        v5 = (v5 << (((v1 - (^(v1)))) & 63))
        fmt.Println(v14, v1, v14)
        var v24 float32 = (-2.25 / -2.25)
        _ = v24
    }
    v14 = ("foo bar" + v3)
    v25 := GenS2{F0: "\n\t"}
    if (!(60 < 1)) {
        v2[210] = 0
        v14 = ""
    } else {
        v3 = ((v3 + "") + (("\n\t" + v14) + "a"))
        v5 -= v5
    }
    if (false || (1 > 127)) {
        v14 = v3
    } else {
        v3 = (((v14 + "foo bar") + v14) + v14)
    }
    v26 := GenS2{F0: v3}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
