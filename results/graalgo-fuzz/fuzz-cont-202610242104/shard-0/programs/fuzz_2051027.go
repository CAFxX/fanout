package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint = 65535
    v3 := []uint32{}
    var v4 uint16 = 811
    v4 -= 627
    fmt.Println(v2, v4)
    switch "foo bar" {
    case "a":
        if (!true) {
            fmt.Println(v2, v2)
            if (2 != 2) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                delete(v1, 0)
                if false {
                    v5 := GenS1{F0: 503}
                    v2 += v2
                    _ = v5
                } else {
                    if false {
                        v2 = (v2 * v2)
                        if true {
                            for fi6 := 0; fi6 < 19; fi6++ {
                                v4 -= (v4 ^ 255)
                                _ = len(v3)
                                _ = fi6
                            }
                            v4 -= 32767
                            var v7 uint = (v2 | (v2 ^ (^(v2))))
                            _ = v7
                        } else {
                            var v8 uint8 = 112
                            _ = len(v1)
                            fmt.Println(v4)
                            _ = v8
                        }
                    } else {
                        var v9 uint = 65535
                        delete(v1, 589)
                        for fi10 := 0; fi10 < 5; fi10++ {
                            fmt.Println(v9, v9)
                            v1[-9223372036854775808] = 0
                            _ = fi10
                        }
                        _ = v9
                    }
                    switch 879 {
                    case 9223372036854775807:
                        v4 = 775
                    case 550:
                        fmt.Println(v4, v2)
                    case 32768:
                        switch 0 {
                        case -484:
                            var v11 string = "foo bar"
                            _ = v11
                        default:
                            var v12 int16 = 32767
                            _ = v12
                        }
                    }
                }
            }
        } else {
            var v13 uint16 = v4
            _ = len(v1)
            _ = v13
        }
    case "":
        v2 = (v2 - v2)
    case "\n\t":
        var v14 uint8 = 2
        _ = v14
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        var v16 string = (("\n\t" + (("a" + (("foo bar" + "a") + ("hello" + ("" + "foo bar")))) + "")) + (("hello" + "a") + "\n\t"))
        _ = fi15
        _ = v16
    }
    if len(v3) > 0 {
        v3[0] = 1
    }
    var v17 float32 = (float32(0.0) / 3.14159)
    var v18 uint8 = 27
    v18 += (^(v18))
    switch "" {
    case "foo bar":
        if (((v17 + (v17 * v17)) - 3.14159) < ((v17 / (v17 - (float32((-2.25 / float64(float32((3.14159 - -2.25))))) + v17))) - (float32((float64(float32(((((-2.25 / 3.14159) - (float64((v17 - v17)) - (0.0 / 3.14159))) - 1e10) - 1e10))) + ((1e10 * (((((0.0 * 0.0) + 1.5) * 3.14159) - 3.14159) * 0.1)) - -2.25))) / v17))) {
            for fi19 := 0; fi19 < 19; fi19++ {
                var v20 string = "\n\t"
                for fi21 := 0; fi21 < 10; fi21++ {
                    var v22 float32 = v17
                    v4 = v4
                    _ = fi21
                    _ = v22
                }
                _ = fi19
                _ = v20
            }
            for fi23 := 0; fi23 < 5; fi23++ {
                if (958 == (v4 % ((v4) | 1))) {
                    var v24 float64 = 1e10
                    _ = v24
                }
                var v25 int16 = 1000
                var v26 uint64 = 114
                _ = fi23
                _ = v25
                _ = v26
            }
        }
    case "a":
        _ = len(v3)
    }
    v18 -= 2
    for fi27 := 0; fi27 < 16; fi27++ {
        var v28 string = ""
        _ = fi27
        _ = v28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
