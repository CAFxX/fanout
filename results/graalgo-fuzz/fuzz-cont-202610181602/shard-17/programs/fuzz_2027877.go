package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 bool }

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[string]int)
    var v3 int32 = -1
    v4 := []uint32{1, 1, 96, 2147483647}
    var v5 string = ("\n\t" + "a")
    var v6 uint8 = 216
    for fi7 := 0; fi7 < 4; fi7++ {
        if true {
            v5 = (v5 + v5)
            var v8 uint32 = 235
            v3 += (v3 | 1)
            _ = v8
        }
        _ = fi7
    }
    switch "a" {
    case "foo bar":
        v1 += (v1 * (0.0 + v1))
    case "":
        var v9 int8 = 100
        _ = v9
    case "\n\t":
        var v10 bool = false
        _ = v10
    }
    var v11 int8 = 40
    if false {
        var v12 uint = 4294967295
        _ = v12
    } else {
        var v13 uint = 65535
        v11 -= (v11 >> ((v13) & 63))
        _ = v13
    }
    if true {
        fmt.Println(v1)
        var v14 bool = true
        _ = v14
    }
    var v15 int32 = (v3 - v3)
    _ = len(v4)
    var v16 string = v5
    fmt.Println(v11, v15, v1)
    var v17 string = ((v16 + "") + (("a" + "hello") + v16))
    if (v15 <= (v3 + (^(v3)))) {
        v3 += (v15 << ((0) & 63))
        fmt.Println(v1, v17)
    }
    if false {
        if false {
            var v18 uint32 = 0
            fmt.Println(v11, v3, v18)
            fmt.Println(v5, v3, v18)
            _ = v18
        } else {
            for fi19 := 0; fi19 < 10; fi19++ {
                v6 += (v6 & 0)
                if false {
                    switch v17 {
                    case "a":
                        if false {
                            v20 := GenS1{F0: v5, F1: 542, F2: true}
                            _ = v20
                        } else {
                            for fi21 := 0; fi21 < 18; fi21++ {
                                var v22 uint = 2
                                var v23 bool = false
                                v24 := GenS1{F0: v5, F1: 872, F2: v23}
                                _ = fi21
                                _ = v22
                                _ = v23
                                _ = v24
                            }
                            if true {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            } else {
                                v25 := GenS1{F0: "\n\t", F1: 2, F2: (((1.5 >= 1e10) && false) && ((!(-650 <= 42)) || (("hello" < ("foo bar" + ("hello" + v16))) && true)))}
                                for fi26 := 0; fi26 < 6; fi26++ {
                                    v5 = (v17 + "a")
                                    v16 = (("\n\t" + v16) + "a")
                                    v4 = append(v4, 714)
                                    _ = fi26
                                }
                                _ = v25
                            }
                        }
                    case "":
                        var v27 int16 = -395
                        _ = v27
                    }
                } else {
                    for fi28 := 0; fi28 < 10; fi28++ {
                        var v29 int32 = (v3 | (^(v3)))
                        _ = fi28
                        _ = v29
                    }
                    fmt.Println(v11)
                }
                _ = fi19
            }
        }
        v4 = append(v4, 506)
    }
    var v30 int64 = 467
    var v31 uint64 = 826
    var v32 uint = 65535
    if ((v6 - 42) >= v6) {
        for fi33 := 0; fi33 < 7; fi33++ {
            if (-653 >= -1) {
                var v34 uint = 42
                _ = v34
            }
            v32 -= v32
            _ = fi33
        }
        if (true && true) {
            for fi35 := 0; fi35 < 1; fi35++ {
                switch v17 {
                case "a":
                    _ = len(v4)
                case "\n\t":
                    v36 := GenS1{F0: (v5 + "hello"), F1: 2, F2: true}
                    _ = v36
                case "":
                    if len(v4) > 0 {
                        v4[0] = 555
                    }
                default:
                    for fi37 := 0; fi37 < 20; fi37++ {
                        var v38 float32 = (1.5 - float32(((-2.25 + (1.5 * (1e10 / float64(v1)))) / ((0.1 + 1e10) - ((1.5 / 1.0) - 0.0)))))
                        var v39 bool = false
                        _ = fi37
                        _ = v38
                        _ = v39
                    }
                }
                _ = fi35
            }
        }
    } else {
        v16 = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
