package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 uint64 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 bool; F1 uint; F2 uint16 }
func genFn0(p0 uint16, p1 uint8, p2 uint64) string {
    for fi6 := 0; fi6 < 14; fi6++ {
        if false {
            v7 := []int{fi6, (fi6 * 972)}
            var v8 string = ((("a" + ("a" + "a")) + "a") + "a")
            _ = v7
            _ = v8
        }
        if (!(!true)) {
            v9 := []int{fi6, 127}
            _ = v9
        } else {
            if (false || true) {
                var v10 float64 = 0.1
                _ = v10
            } else {
                switch "a" {
                case "foo bar":
                    v11 := GenS1{F0: (^(p2)), F1: 690, F2: 26}
                    _ = v11
                case "a":
                    switch ("foo bar" + "\n\t") {
                    case "hello":
                        if ((3.14159 - (-2.25 / -2.25)) == 1.5) {
                            v12 := GenS2{F0: 32767, F1: "foo bar"}
                            _ = v12
                        }
                    default:
                        if false {
                            v13 := GenS2{F0: 0, F1: "hello"}
                            v14 := GenS1{F0: (p2 ^ (p2 - p2)), F1: 285, F2: (p2 << ((2) & 63))}
                            _ = v13
                            _ = v14
                        } else {
                            v15 := GenS2{F0: 0, F1: ("foo bar" + ("a" + ""))}
                            var v16 uint64 = 106
                            switch "" {
                            case "hello":
                                switch "hello" {
                                case "":
                                    var v17 int32 = int32(v16)
                                    _ = v17
                                case "foo bar":
                                    p0 -= p0
                                default:
                                    fmt.Println(fi6, p0, p2)
                                }
                            default:
                                if true {
                                    v18 := []int{}
                                    _ = v18
                                }
                            }
                            _ = v15
                            _ = v16
                        }
                    }
                }
            }
        }
        for fi19 := 0; fi19 < 9; fi19++ {
            fmt.Println(p0, p0, p2)
            p0 -= p0
            _ = fi19
        }
        _ = fi6
    }
    var v20 uint16 = (p0 - (^(p0)))
    p2 += 766
    v21 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v21
    return "a"
}

func main() {
    var v1 int = -29
    var v2 uint16 = 65535
    var v3 uint32 = 0
    v4 := make(map[int]uint32)
    var v5 uint = 65535
    var v22 int16 = -32768
    v23 := GenS2{F0: 42, F1: "a"}
    v24 := v4[-1]
    v25 := GenS3{F0: (v5 > (v5 + v5)), F1: 2, F2: (v2 & 255)}
    if (true || (100 <= -128)) {
        v5 -= v5
        v2 = v2
    }
    v5 = v5
    v3 -= v3
    delete(v4, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
