package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint8 }
func genFn0(p0 uint, p1 int, p2 uint64) uint16 {
    if (65535 >= 2) {
        var v4 uint8 = 42
        _ = v4
    } else {
        switch (-(p1)) {
        case -746:
            var v5 int64 = 9223372036854775807
            _ = v5
        default:
            v6 := []int{p1}
            _ = v6
        }
        switch p1 {
        case -1:
            if ("hello" != ((("foo bar" + "a") + "\n\t") + "a")) {
                p1 += p1
            }
        case 128:
            for fi7 := 0; fi7 < 2; fi7++ {
                fmt.Println(p1, fi7, p1)
                p0 = p0
                v8 := []int{fi7, fi7, fi7, (p1 & -9223372036854775808)}
                _ = fi7
                _ = v8
            }
        default:
            var v9 float64 = ((float64((1e10 / 1e10)) + ((float64((0.1 + 1.5)) / 0.1) - 1e10)) - 1e10)
            _ = v9
        }
        for fi10 := 0; fi10 < 11; fi10++ {
            var v11 int = (-(fi10))
            p2 = 340
            _ = fi10
            _ = v11
        }
    }
    fmt.Println(p1)
    var v12 uint32 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 255
}

func main() {
    v1 := make(map[int]string)
    var v2 uint64 = 4294967295
    var v3 int8 = 25
    var v13 bool = false
    delete(v1, 468)
    var v14 uint64 = (v2 - (v2 - v2))
    var v15 uint = 18446744073709551615
    _ = len(v1)
    v16 := GenS1{F0: 736, F1: v13, F2: 2}
    if (393 >= 0) {
        if (v13 || v13) {
            if v13 {
                var v17 string = "hello"
                v17 = "\n\t"
                for fi18 := 0; fi18 < 8; fi18++ {
                    switch 215 {
                    case 32768:
                        if false {
                            for fi19 := 0; fi19 < 11; fi19++ {
                                var v20 uint16 = 2
                                _ = fi19
                                _ = v20
                            }
                            var v21 uint32 = 65535
                            fmt.Println(v17, v13, v2)
                            _ = v21
                        }
                    case -862:
                        v1[331] = ("" + v17)
                    }
                    var v22 uint16 = 32767
                    fmt.Println(v15, v2)
                    _ = fi18
                    _ = v22
                }
                _ = v17
            }
            v14 = (v14 << ((v15) & 63))
        }
    }
    v13 = (526 == 685)
    v23 := GenS1{F0: 65535, F1: ("foo bar" <= "hello"), F2: 128}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v3
}
