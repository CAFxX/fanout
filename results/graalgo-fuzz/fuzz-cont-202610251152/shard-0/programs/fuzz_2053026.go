package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 bool }
func genFn0(p0 int64) int64 {
    v6 := GenS1{F0: (true && (0 < 65535))}
    v7 := GenS2{F0: 1, F1: 1}
    var v8 int16 = 32767
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return p0
}

func main() {
    var v1 float64 = 0.0
    var v2 string = "foo bar"
    v3 := []string{}
    var v4 float32 = float32(v1)
    var v5 uint = 1
    v5 = (v5 >> ((178) & 63))
    v9 := GenS3{F0: (!(true || (!(65535 > 65535))))}
    if (!(true && false)) {
        v1 = v1
        if false {
            if len(v3) > 0 {
                v3[0] = ("hello" + "hello")
            }
            if true {
                for fi10 := 0; fi10 < 6; fi10++ {
                    if (!false) {
                        _ = len(v3)
                    }
                    fmt.Println(v1)
                    var v11 float64 = v1
                    _ = fi10
                    _ = v11
                }
                v1 += (0.0 / 1.0)
            } else {
                if (!false) {
                    v5 -= v5
                }
            }
        } else {
            var v12 uint8 = 10
            v13 := GenS1{F0: false}
            if false {
                for fi14 := 0; fi14 < 11; fi14++ {
                    var v15 int8 = 1
                    switch fi14 {
                    case -9223372036854775808:
                        switch 212 {
                        case 4294967295:
                            v16 := GenS3{F0: true}
                            _ = v16
                        case 309:
                            var v17 uint8 = (v12 / (((v12 % ((v12) | 1))) | 1))
                            _ = v17
                        }
                    case 128:
                        v5 -= v5
                    case -523:
                        v15 = v15
                    default:
                        v18 := GenS3{F0: true}
                        _ = v18
                    }
                    _ = fi14
                    _ = v15
                }
                switch v2 {
                case "":
                    _ = len(v3)
                case "hello":
                    for fi19 := 0; fi19 < 7; fi19++ {
                        for fi20 := 0; fi20 < 7; fi20++ {
                            v2 = ((v2 + "") + "a")
                            var v21 uint16 = 758
                            var v22 float64 = v1
                            _ = fi20
                            _ = v21
                            _ = v22
                        }
                        _ = fi19
                    }
                }
            } else {
                if len(v3) > 0 {
                    v3[0] = "hello"
                }
            }
            _ = v12
            _ = v13
        }
        for fi23 := 0; fi23 < 9; fi23++ {
            fmt.Println(v4, v2, v5)
            _ = fi23
        }
    }
    v4 += 0.1
    v4 = ((v4 / v4) / (v4 - v4))
    fmt.Println(v4)
    fmt.Println(v5, v1, v2)
    _ = len(v3)
    v4 += (v4 + (v4 * v4))
    if ((1e10 < v1) && (550 >= 238)) {
        for fi24 := 0; fi24 < 20; fi24++ {
            var v25 bool = (!(!false))
            _ = fi24
            _ = v25
        }
    }
    var v26 int64 = genFn0(956)
    if len(v3) > 0 {
        v3[0] = ("a" + "\n\t")
    }
    v27 := GenS3{F0: (((32767 > 42) && ("\n\t" == v2)) || (-128 <= 22))}
    v26 += (-(v26))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
