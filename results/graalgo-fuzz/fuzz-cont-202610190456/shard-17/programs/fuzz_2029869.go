package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    v1 := []uint32{}
    var v2 float64 = 1.5
    var v3 uint8 = 157
    v4 := []uint32{94, 581, 537}
    var v5 int16 = 1000
    v6 := []uint32{2}
    var v7 bool = ((float64((((((3.14159 * (0.1 + 1.5)) - 3.14159) + 1.5) + (3.14159 + 0.1)) * 0.1)) <= -2.25) || (v5 == (^(v5))))
    if len(v6) > 0 {
        v6[0] = 1
    }
    if v7 {
        fmt.Println(v7, v5)
        switch 42 {
        case 42:
            for fi8 := 0; fi8 < 12; fi8++ {
                var v9 float32 = 1.5
                fmt.Println(v9)
                _ = fi8
                _ = v9
            }
        case 1:
            if len(v4) > 0 {
                v4[0] = 4294967295
            }
        case 32768:
            if false {
                for fi10 := 0; fi10 < 12; fi10++ {
                    switch 128 {
                    case 4294967295:
                        v7 = false
                    case 488:
                        var v11 int64 = -545
                        _ = v11
                    default:
                        switch (("hello" + "\n\t") + "") {
                        case "hello":
                            var v12 float32 = ((1.5 / 1e10) + 3.14159)
                            _ = v12
                        case "":
                            var v13 int16 = (v5 / ((-1) | 1))
                            _ = v13
                        case "\n\t":
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        }
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi10
                }
            } else {
                if (v2 > v2) {
                    _ = len(v6)
                    v1 = append(v1, 42)
                }
                var v14 uint = 464
                var v15 uint16 = 3
                _ = v14
                _ = v15
            }
        }
    }
    var v16 bool = v7
    v2 = v2
    for fi17 := 0; fi17 < 1; fi17++ {
        if len(v1) > 0 {
            v1[0] = 53
        }
        v7 = v7
        v16 = true
        _ = fi17
    }
    var v18 int16 = (v5 - (v5 + (v5 + v5)))
    if ((v16 && (32767 != v18)) && v7) {
        v19 := GenS1{F0: 1}
        _ = v19
    }
    fmt.Println(v18, v18)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
