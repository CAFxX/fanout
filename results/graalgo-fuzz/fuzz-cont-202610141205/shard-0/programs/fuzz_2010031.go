package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := []int{-730, 0}
    var v2 float32 = 3.14159
    var v3 int8 = 50
    var v4 uint8 = 112
    switch 32768 {
    case 1:
        if (true || ("" == ("\n\t" + ("foo bar" + "a")))) {
            v5 := GenS1{F0: 963}
            v2 += (0.1 + v2)
            switch "foo bar" {
            case "":
                if true {
                    for fi6 := 0; fi6 < 18; fi6++ {
                        for fi7 := 0; fi7 < 9; fi7++ {
                            fmt.Println(v2, v3, v2)
                            var v8 float64 = ((0.0 / float64(((-2.25 - 3.14159) * v2))) - ((float64(((v2 - (-2.25 * 0.1)) / 0.1)) - (3.14159 / 1.5)) / float64(v2)))
                            _ = fi7
                            _ = v8
                        }
                        _ = fi6
                    }
                    var v9 int = -1
                    _ = v9
                } else {
                    switch "a" {
                    case "a":
                        if false {
                            for fi10 := 0; fi10 < 10; fi10++ {
                                for fi11 := 0; fi11 < 2; fi11++ {
                                    v2 -= 1e10
                                    _ = fi11
                                }
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                var v12 bool = ((false && (true && true)) && false)
                                _ = fi10
                                _ = v12
                            }
                        } else {
                            if len(v1) > 0 {
                                v1[0] = 127
                            }
                        }
                    default:
                        if true {
                            var v13 string = "hello"
                            v3 += v3
                            _ = v13
                        }
                    }
                }
            case "\n\t":
                var v14 int16 = -709
                _ = v14
            default:
                v15 := GenS1{F0: 442}
                _ = v15
            }
            _ = v5
        }
    default:
        var v16 int8 = v3
        _ = v16
    }
    v2 -= v2
    v4 = v4
    if (v4 > (v4 - (v4 % ((v4) | 1)))) {
        for fi17 := 0; fi17 < 15; fi17++ {
            v18 := GenS1{F0: 434}
            _ = fi17
            _ = v18
        }
        v1 = append(v1, 4294967295)
        v1 = append(v1, -9223372036854775808)
    }
    v3 += -128
    fmt.Println(v2)
    if len(v1) > 0 {
        v1[0] = 42
    }
    fmt.Println(v3, v3)
    for fi19 := 0; fi19 < 12; fi19++ {
        var v20 int16 = 32767
        _ = fi19
        _ = v20
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
