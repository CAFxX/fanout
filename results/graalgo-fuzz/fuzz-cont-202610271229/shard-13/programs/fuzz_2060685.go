package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8 }
func genFn0(p0 uint16) bool {
    p0 = (p0 << ((42) & 63))
    fmt.Println(p0)
    var v5 int16 = 870
    if false {
        p0 += (p0 - (^(p0)))
        v5 = 32767
    }
    v6 := []int{}
    _ = p0
    _ = v5
    _ = v6
    return true
}

func main() {
    var v1 string = ""
    var v2 int32 = 0
    var v3 float64 = float64((float32(1.5) / 0.1))
    v4 := []int{1, -76, 42, 2147483647}
    v4 = append(v4, -2147483648)
    v7 := GenS1{F0: 32767}
    v4 = append(v4, 42)
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 string = v1
        _ = fi8
        _ = v9
    }
    v3 -= ((v3 + ((v3 / (v3 - v3)) - v3)) + (v3 + ((v3 * v3) + 1e10)))
    for fi10 := 0; fi10 < 6; fi10++ {
        var v11 int8 = 100
        fmt.Println(v2, v2, fi10)
        v12 := GenS1{F0: 2}
        _ = fi10
        _ = v11
        _ = v12
    }
    v3 -= (v3 + 1.5)
    v3 = (((v3 / (1.5 - v3)) / float64(3.14159)) / 1.0)
    if true {
        if false {
            v3 = (((((v3 * v3) / 1.0) + (1e10 * (1e10 * -2.25))) + ((v3 * v3) - ((float64(0.1) / v3) - 0.1))) / v3)
            var v13 int16 = -908
            _ = v13
        } else {
            var v14 int64 = 1
            v4 = append(v4, 2147483647)
            _ = v14
        }
        if (42 < -606) {
            switch 362 {
            case -2147483648:
                if false {
                    fmt.Println(v2, v2, v1)
                } else {
                    for fi15 := 0; fi15 < 14; fi15++ {
                        var v16 float32 = float32((v3 - 0.0))
                        _ = len(v4)
                        _ = fi15
                        _ = v16
                    }
                }
            case 42:
                switch -806 {
                case -2147483648:
                    fmt.Println(v1)
                case 42:
                    v1 = v1
                case 683:
                    var v17 float32 = 3.14159
                    _ = v17
                default:
                    if true {
                        _ = len(v4)
                        fmt.Println(v1)
                        v18 := GenS1{F0: 42}
                        _ = v18
                    }
                }
            case 32767:
                v19 := GenS1{F0: 537}
                _ = v19
            default:
                if true {
                    switch ("a" + (((("a" + v1) + "\n\t") + v1) + v1)) {
                    case "foo bar":
                        switch 1 {
                        case 0:
                            if ((1 != 127) && (true && true)) {
                                v20 := GenS2{F0: 25}
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                var v21 float64 = ((v3 / 1.5) + 0.0)
                                _ = v20
                                _ = v21
                            }
                        case -2147483648:
                            switch "a" {
                            case "\n\t":
                                var v22 uint = 2
                                _ = v22
                            default:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                        case 2147483647:
                            var v23 int = 1
                            _ = v23
                        default:
                            for fi24 := 0; fi24 < 6; fi24++ {
                                v3 = v3
                                _ = fi24
                            }
                        }
                    case "\n\t":
                        fmt.Println(v3, v3, v2)
                    default:
                        v2 = v2
                    }
                }
            }
            fmt.Println(v3, v1, v1)
            switch v1 {
            case "":
                var v25 string = (v1 + "a")
                _ = v25
            case "a":
                v2 = (-(v2))
            }
        }
    }
    v26 := GenS1{F0: 382}
    v1 = ""
    switch 307 {
    case 0:
        v3 -= v3
    case 1:
        var v27 string = "hello"
        _ = v27
    case 128:
        v3 = 3.14159
    }
    fmt.Println(v1)
    for fi28 := 0; fi28 < 13; fi28++ {
        fmt.Println(v3, fi28, v3)
        _ = fi28
    }
    v29 := GenS1{F0: 65535}
    fmt.Println(v3, v3)
    for fi30 := 0; fi30 < 5; fi30++ {
        v1 = v1
        v1 = v1
        var v31 int16 = -28
        _ = fi30
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v4
    _ = v7
}
