package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 float64 = ((3.14159 + 0.0) / (float64(0.1) * 3.14159))
    var v2 uint = 255
    v3 := []int{604, 432, 1}
    v4 := []string{}
    var v5 int8 = 100
    var v6 float32 = 0.0
    fmt.Println(v2, v2, v2)
    v1 += ((v1 - 1e10) / v1)
    var v7 uint = (^(v2))
    switch -267 {
    case 128:
        var v8 uint64 = 1
        _ = v8
    case 10:
        fmt.Println(v7, v1)
    default:
        switch ("foo bar" + "a") {
        case "foo bar":
            var v9 int8 = (v5 >> ((v7) & 63))
            _ = v9
        case "":
            var v10 uint8 = 16
            _ = v10
        case "a":
            switch 242 {
            case 0:
                switch -668 {
                case -2147483648:
                    if false {
                        if (2147483647 != -648) {
                            fmt.Println(v5, v6, v6)
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        } else {
                            v2 += v2
                            for fi11 := 0; fi11 < 2; fi11++ {
                                switch "" {
                                case "hello":
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                case "":
                                    var v12 uint64 = 1
                                    _ = v12
                                case "a":
                                    v1 += 1.5
                                default:
                                    v5 += (v5 ^ 1)
                                }
                                for fi13 := 0; fi13 < 12; fi13++ {
                                    if true {
                                        var v14 uint8 = 110
                                        _ = v14
                                    }
                                    _ = fi13
                                }
                                for fi15 := 0; fi15 < 19; fi15++ {
                                    v6 += (v6 + 0.1)
                                    v5 += -1
                                    _ = fi15
                                }
                                _ = fi11
                            }
                        }
                    } else {
                        var v16 bool = (float64(v6) >= v1)
                        v17 := GenS1{F0: 407}
                        _ = v16
                        _ = v17
                    }
                case 0:
                    for fi18 := 0; fi18 < 13; fi18++ {
                        var v19 int = 42
                        if ((v5 << ((v7) & 63)) > v5) {
                            v20 := GenS1{F0: 42}
                            v5 += (^(v5))
                            _ = v20
                        } else {
                            var v21 int = fi18
                            _ = v21
                        }
                        var v22 int = fi18
                        _ = fi18
                        _ = v19
                        _ = v22
                    }
                default:
                    var v23 string = ("hello" + ("\n\t" + ("\n\t" + "hello")))
                    _ = v23
                }
            case 42:
                for fi24 := 0; fi24 < 12; fi24++ {
                    for fi25 := 0; fi25 < 11; fi25++ {
                        fmt.Println(v7, fi25)
                        v26 := GenS1{F0: 42}
                        _ = fi25
                        _ = v26
                    }
                    var v27 int = fi24
                    fmt.Println(fi24, v27)
                    _ = fi24
                    _ = v27
                }
            default:
                if false {
                    if (1 == 2147483648) {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = len(v3)
                        for fi28 := 0; fi28 < 19; fi28++ {
                            if (v6 == v6) {
                                var v29 float32 = (3.14159 / 1e10)
                                _ = v29
                            } else {
                                v2 += (v2 & (v2 % ((2) | 1)))
                                fmt.Println(v5)
                                v2 -= (^(v7))
                            }
                            var v30 int = fi28
                            v7 -= v7
                            _ = fi28
                            _ = v30
                        }
                    }
                } else {
                    v6 += v6
                    switch 536 {
                    case 0:
                        if (!false) {
                            if (42 <= 0) {
                                var v31 uint8 = 2
                                _ = v31
                            }
                            var v32 uint8 = 2
                            _ = v32
                        }
                    case 2147483647:
                        var v33 int16 = -32768
                        _ = v33
                    default:
                        v34 := GenS1{F0: 42}
                        _ = v34
                    }
                }
            }
        default:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    }
    v35 := GenS1{F0: 42}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v36 int16 = 719
    var v37 int16 = v36
    v1 = 1.5
    var v38 uint32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
