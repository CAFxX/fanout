package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 int8; F2 uint8 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 uint16) uint64 {
    p0 = 0
    p0 -= 32767
    var v7 string = (("a" + "") + ((((("a" + (("foo bar" + ("foo bar" + "foo bar")) + "")) + "a") + ((("" + "") + "\n\t") + "a")) + "\n\t") + "foo bar"))
    v8 := []int{}
    if (42 == -1) {
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        if false {
            fmt.Println(p0, v7)
        } else {
            if false {
                var v9 float64 = (1.5 - 1.5)
                fmt.Println(v9)
                v9 = v9
                _ = v9
            }
            var v10 uint = 18446744073709551615
            _ = v10
        }
    } else {
        var v11 string = v7
        v11 = ""
        _ = v11
    }
    _ = p0
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 float32 = 0.1
    var v2 int16 = -32768
    v3 := make(map[int]int)
    v4 := []uint32{101, 2147483648}
    var v5 float64 = 1.5
    var v6 int8 = -87
    v2 += (v2 >> ((2) & 63))
    v2 += (v2 / (((v2 % ((0) | 1))) | 1))
    var v12 int16 = v2
    for fi13 := 0; fi13 < 18; fi13++ {
        var v14 uint32 = 780
        v15 := GenS1{F0: 42}
        var v16 uint64 = 2
        _ = fi13
        _ = v14
        _ = v15
        _ = v16
    }
    v12 = (-(v2))
    v12 = (v2 - v12)
    v17 := GenS1{F0: 0}
    for fi18 := 0; fi18 < 11; fi18++ {
        fmt.Println(v6, v2, v5)
        for fi19 := 0; fi19 < 1; fi19++ {
            v6 -= (v6 ^ (v6 << ((42) & 63)))
            v2 += (v2 | (v12 * 1))
            if false {
                v2 += (^(v12))
                switch fi19 {
                case 4294967295:
                    v20 := GenS1{F0: 42}
                    _ = v20
                case 884:
                    _ = len(v4)
                default:
                    v21 := GenS2{F0: 0, F1: (v6 - 0), F2: 1}
                    _ = v21
                }
            } else {
                v12 -= (-(v2))
                var v22 string = ""
                _ = v22
            }
            _ = fi19
        }
        var v23 uint = 42
        _ = fi18
        _ = v23
    }
    switch 610 {
    case 675:
        var v24 bool = true
        _ = v24
    case 450:
        if ((1 >= 2147483647) && false) {
            if false {
                if (!(true && (127 > 1))) {
                    var v25 uint16 = 2
                    var v26 int = 42
                    _ = v25
                    _ = v26
                } else {
                    for fi27 := 0; fi27 < 17; fi27++ {
                        var v28 int = fi27
                        _ = fi27
                        _ = v28
                    }
                    fmt.Println(v2)
                }
                if (574 < 1) {
                    if (100 >= v6) {
                        fmt.Println(v5)
                        v1 = v1
                        v1 -= v1
                    }
                    if true {
                        fmt.Println(v5, v5)
                        v2 = v2
                        if (-1 < 42) {
                            for fi29 := 0; fi29 < 5; fi29++ {
                                if len(v4) > 0 {
                                    v4[0] = 0
                                }
                                if true {
                                    v5 = ((((v5 - v5) + v5) + (v5 + 0.1)) + v5)
                                } else {
                                    switch "hello" {
                                    case "":
                                        if true {
                                            for fi30 := 0; fi30 < 15; fi30++ {
                                                var v31 uint = 4294967295
                                                fmt.Println(v5)
                                                _ = fi30
                                                _ = v31
                                            }
                                            fmt.Println(v5, fi29, v12)
                                            v12 += v2
                                        }
                                    }
                                }
                                _ = fi29
                            }
                            _ = len(v3)
                            var v32 int16 = v2
                            _ = v32
                        }
                    } else {
                        fmt.Println(v6, v2)
                        v3[770] = 32767
                    }
                } else {
                    switch "" {
                    case "hello":
                        for fi33 := 0; fi33 < 8; fi33++ {
                            switch "hello" {
                            case "foo bar":
                                delete(v3, 1)
                            case "a":
                                _ = len(v3)
                            case "\n\t":
                                v2 = -1
                            }
                            _ = fi33
                        }
                    default:
                        delete(v3, -2147483648)
                    }
                }
                _ = len(v4)
            } else {
                var v34 float32 = 1e10
                v35 := GenS1{F0: 130}
                for fi36 := 0; fi36 < 9; fi36++ {
                    for fi37 := 0; fi37 < 20; fi37++ {
                        if (fi37 >= fi36) {
                            _ = len(v4)
                        }
                        var v38 int8 = (v6 | v6)
                        if ((1 >= 2) && false) {
                            fmt.Println(v34)
                            fmt.Println(v38, v1, v12)
                        }
                        _ = fi37
                        _ = v38
                    }
                    switch (fi36 ^ fi36) {
                    case 127:
                        v39 := GenS3{F0: 2}
                        _ = v39
                    case 422:
                        var v40 int8 = (v6 * v6)
                        _ = v40
                    }
                    _ = fi36
                }
                _ = v34
                _ = v35
            }
            var v41 int16 = 1
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = v41
        }
    case 2147483647:
        v5 = v5
    }
    v42 := GenS2{F0: 87, F1: v6, F2: 127}
    fmt.Println(v6)
    switch 4294967295 {
    case 32767:
        if (!true) {
            v43 := GenS3{F0: 42}
            if ((v5 / v5) < -2.25) {
                fmt.Println(v2, v5)
                _ = len(v3)
            }
            if (!(true && false)) {
                var v44 int32 = 0
                v1 -= 1e10
                v1 = 0.1
                _ = v44
            } else {
                v6 = (v6 - (v6 * 1))
            }
            _ = v43
        } else {
            v2 = (^(v2))
            fmt.Println(v12, v5, v1)
        }
    case 42:
        var v45 uint32 = 429
        _ = v45
    case 41:
        _ = len(v4)
    default:
        v46 := GenS2{F0: 105, F1: v6, F2: 127}
        _ = v46
    }
    v47 := GenS1{F0: 185}
    var v48 uint64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
