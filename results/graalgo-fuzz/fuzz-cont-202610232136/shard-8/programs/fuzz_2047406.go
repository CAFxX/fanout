package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 uint32; F1 uint16; F2 int32 }

func main() {
    var v1 bool = false
    var v2 string = "a"
    v3 := []string{(v2 + v2), "", "\n\t"}
    var v4 int16 = 32767
    var v5 float64 = 1.5
    v4 = (v4 | (v4 + (v4 - v4)))
    var v6 float32 = ((0.1 + (1.5 - 1.5)) + -2.25)
    var v7 string = v2
    v1 = (42 != 1)
    for fi8 := 0; fi8 < 4; fi8++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v9 := GenS1{F0: -1, F1: 227}
        v3 = append(v3, (("foo bar" + (v2 + v7)) + v2))
        _ = fi8
        _ = v9
    }
    var v10 uint16 = 65535
    var v11 float64 = float64(v6)
    fmt.Println(v11, v1, v6)
    v6 = v6
    if (914 != 507) {
        if true {
            v12 := GenS2{F0: 2147483648, F1: 32767, F2: 115}
            if v1 {
                switch 750 {
                case 128:
                    if len(v3) > 0 {
                        v3[0] = "foo bar"
                    }
                case 2147483647:
                    var v13 float32 = v6
                    _ = v13
                case -2147483648:
                    fmt.Println(v10, v10, v2)
                default:
                    if true {
                        for fi14 := 0; fi14 < 5; fi14++ {
                            switch 1 {
                            case -368:
                                var v15 int = 127
                                _ = v15
                            case 127:
                                if v1 {
                                    var v16 int32 = 244
                                    v6 -= ((3.14159 * (((v6 - v6) / v6) - float32((v11 + (float64(v6) - 3.14159))))) / 1e10)
                                    _ = v16
                                } else {
                                    var v17 int32 = 0
                                    _ = v17
                                }
                            }
                            if len(v3) > 0 {
                                v3[0] = v7
                            }
                            if (true && false) {
                                v11 -= v11
                            } else {
                                var v18 uint8 = 127
                                fmt.Println(v7, v10, v7)
                                _ = v18
                            }
                            _ = fi14
                        }
                    } else {
                        switch -805 {
                        case -1:
                            v2 = v7
                        default:
                            switch -1 {
                            case 4294967295:
                                var v19 float64 = 0.0
                                _ = v19
                            default:
                                fmt.Println(v6, v10)
                            }
                        }
                    }
                }
                v3 = append(v3, v7)
            } else {
                fmt.Println(v4, v6)
            }
            v4 -= 1
            _ = v12
        } else {
            switch -9223372036854775808 {
            case 42:
                v20 := GenS2{F0: 65535, F1: (v10 - 42), F2: 42}
                _ = v20
            }
            var v21 uint16 = v10
            _ = v21
        }
    }
    for fi22 := 0; fi22 < 17; fi22++ {
        v2 = v2
        var v23 float32 = -2.25
        _ = fi22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
