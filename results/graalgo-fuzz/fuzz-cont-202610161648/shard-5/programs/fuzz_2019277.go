package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 int8 }

func main() {
    var v1 int = -1
    v2 := make(map[string]bool)
    var v3 uint32 = 4294967295
    if false {
        delete(v2, "a")
    } else {
        if (42 > 100) {
            var v4 uint = 65535
            _ = v4
        }
        v5 := GenS1{F0: 78, F1: (v1 % (((v1 & (v1 | (-(v1))))) | 1))}
        v1 -= -1
        _ = v5
    }
    var v6 int8 = 82
    v6 -= 1
    for fi7 := 0; fi7 < 2; fi7++ {
        switch ("a" + "foo bar") {
        case "foo bar":
            if false {
                var v8 uint = 4294967295
                if (32767 >= v1) {
                    v9 := GenS1{F0: 2, F1: v1}
                    v6 = v6
                    v10 := GenS1{F0: 3, F1: v1}
                    _ = v10
                    _ = v9
                } else {
                    v6 = (v6 >> ((v8) & 63))
                    v11 := GenS1{F0: 41, F1: v1}
                    v12 := []int{v1, 1}
                    _ = v11
                    _ = v12
                }
                fmt.Println(v8, v6)
                _ = v8
            } else {
                var v13 int8 = v6
                switch "" {
                case "":
                    if false {
                        var v14 uint64 = 0
                        _ = v14
                    } else {
                        switch (fi7 + (-(v1))) {
                        case -1:
                            v13 = 100
                        case 163:
                            v15 := []int{v1, fi7}
                            _ = v15
                        case 862:
                            var v16 uint = 1
                            _ = v16
                        }
                    }
                default:
                    var v17 int = fi7
                    _ = v17
                }
                v6 = 100
                _ = v13
            }
        case "hello":
            fmt.Println(v1, v3)
        default:
            v1 = 0
        }
        for fi18 := 0; fi18 < 3; fi18++ {
            v3 -= v3
            v19 := GenS2{F0: v6}
            _ = fi18
            _ = v19
        }
        _ = fi7
    }
    fmt.Println(v1, v3)
    var v20 float32 = (float32(0.0) / (1.5 - float32(1e10)))
    var v21 string = "a"
    v22 := GenS2{F0: v6}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v6
}
