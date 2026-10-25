package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 uint }
func genFn0(p0 uint64) int32 {
    var v5 uint = 2
    var v6 uint16 = 1
    v7 := []int{0, -2147483648}
    if true {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    }
    for fi8 := 0; fi8 < 17; fi8++ {
        for fi9 := 0; fi9 < 10; fi9++ {
            v5 += (v5 * (v5 | 255))
            if (!(-118 != 100)) {
                if true {
                    _ = len(v7)
                    v10 := GenS1{F0: 2}
                    var v11 int32 = -696
                    _ = v10
                    _ = v11
                } else {
                    v12 := GenS1{F0: 42}
                    if false {
                        v6 = (v6 % ((v6) | 1))
                        if len(v7) > 0 {
                            v7[0] = fi8
                        }
                    }
                    _ = v12
                }
            } else {
                v5 = (^(v5))
            }
            var v13 float64 = 1.5
            var v14 float32 = float32(v13)
            _ = fi9
            _ = v13
            _ = v14
        }
        _ = fi8
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 uint = 255
    var v2 uint = v1
    var v3 uint8 = 2
    var v4 uint = (^(v1))
    var v15 float32 = 3.14159
    v16 := []int{-845, 9223372036854775807, 128}
    v16 = append(v16, 127)
    for fi17 := 0; fi17 < 6; fi17++ {
        v2 += 42
        v4 += v1
        _ = fi17
    }
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    if false {
        fmt.Println(v3, v4)
        v1 += (v1 ^ v2)
    }
    fmt.Println(v3, v15, v3)
    var v18 uint16 = 65535
    for fi19 := 0; fi19 < 9; fi19++ {
        if len(v16) > 0 {
            v16[0] = fi19
        }
        _ = fi19
    }
    _ = len(v16)
    fmt.Println(v18)
    switch 415 {
    case 4294967295:
        if ((("hello" + "a") + ("a" + "foo bar")) >= "\n\t") {
            for fi20 := 0; fi20 < 12; fi20++ {
                v21 := GenS2{F0: (^(v18)), F1: 42}
                _ = fi20
                _ = v21
            }
        }
    default:
        if (false && (1 < 0)) {
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
            switch "" {
            case "hello":
                var v22 string = ("\n\t" + ("hello" + "\n\t"))
                _ = v22
            case "":
                for fi23 := 0; fi23 < 12; fi23++ {
                    if len(v16) >= 2 {
                        v16 = v16[1:len(v16)]
                    }
                    v24 := GenS1{F0: 0}
                    for fi25 := 0; fi25 < 20; fi25++ {
                        v18 -= v18
                        switch 1 {
                        case -413:
                            fmt.Println(v4)
                        default:
                            var v26 float32 = v15
                            _ = v26
                        }
                        _ = fi25
                    }
                    _ = fi23
                    _ = v24
                }
            case "a":
                v16 = append(v16, 42)
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
