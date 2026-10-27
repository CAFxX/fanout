package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 uint }

func main() {
    v1 := []int{813, 79, 9223372036854775807, -914}
    var v2 bool = true
    var v3 string = "hello"
    fmt.Println(v2, v3)
    var v4 int16 = 32767
    switch "" {
    case "hello":
        v5 := GenS1{F0: 1}
        _ = v5
    }
    v3 = v3
    for fi6 := 0; fi6 < 10; fi6++ {
        for fi7 := 0; fi7 < 8; fi7++ {
            var v8 int = fi6
            v1 = append(v1, 32768)
            v4 = (v4 | v4)
            var v9 bool = v2
            _ = fi7
            _ = v8
            _ = v9
        }
        if false {
            v4 = (v4 >> ((255) & 63))
        } else {
            if ((!(2 <= 1)) && false) {
                var v10 bool = (false && (v3 >= "\n\t"))
                if (0 >= 0) {
                    if v2 {
                        v11 := GenS1{F0: 1}
                        _ = v11
                    } else {
                        if v2 {
                            v10 = v10
                            var v12 uint8 = 128
                            if len(v1) > 0 {
                                v1[0] = (fi6 & fi6)
                            }
                            _ = v12
                        }
                    }
                    var v13 bool = v2
                    var v14 uint32 = 2
                    _ = v13
                    _ = v14
                } else {
                    v4 += (^(v4))
                    v2 = (42 >= 42)
                    fmt.Println(v10)
                }
                for fi15 := 0; fi15 < 19; fi15++ {
                    var v16 int = fi6
                    _ = fi15
                    _ = v16
                }
                _ = v10
            } else {
                switch fi6 {
                case 363:
                    for fi17 := 0; fi17 < 10; fi17++ {
                        v18 := GenS1{F0: 65535}
                        _ = fi17
                        _ = v18
                    }
                }
            }
            if false {
                switch (-(fi6)) {
                case -204:
                    if (2 != 1) {
                        if v2 {
                            v4 = v4
                            if len(v1) > 0 {
                                v1[0] = 0
                            }
                        }
                        for fi19 := 0; fi19 < 6; fi19++ {
                            if len(v1) > 0 {
                                v1[0] = (fi19 | fi6)
                            }
                            v20 := GenS2{F0: v3, F1: 357}
                            _ = fi19
                            _ = v20
                        }
                    } else {
                        v2 = (278 <= 255)
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        fmt.Println(v3, v3)
                    }
                case -181:
                    var v21 uint16 = 0
                    _ = v21
                case 9223372036854775807:
                    v2 = ((!(-2147483648 < 126)) && (v2 || ((1.5 - 0.0) > 1.5)))
                default:
                    var v22 uint32 = 2147483648
                    _ = v22
                }
                var v23 uint32 = 2
                _ = v23
            } else {
                var v24 uint64 = 2
                _ = v24
            }
        }
        _ = fi6
    }
    v25 := GenS2{F0: ((v3 + ((v3 + v3) + v3)) + ((v3 + ("a" + v3)) + v3)), F1: 18446744073709551615}
    var v26 uint16 = 42
    if len(v1) > 0 {
        v1[0] = 1
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
