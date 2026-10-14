package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]string)
    var v3 int64 = 0
    v4 := make(map[uint32]string)
    var v5 string = "hello"
    switch "a" {
    case "":
        switch -9223372036854775808 {
        case -317:
            var v6 uint = 255
            _ = v6
        case -2147483648:
            if false {
                switch 42 {
                case 9223372036854775807:
                    if (42 <= 370) {
                        var v7 bool = (0 != 1)
                        v7 = v7
                        _ = v7
                    } else {
                        var v8 int8 = -8
                        _ = v8
                    }
                default:
                    if true {
                        delete(v2, 2147483647)
                        v5 = v5
                    }
                }
                v3 += -9223372036854775808
                var v9 string = "foo bar"
                _ = v9
            }
        }
    case "foo bar":
        var v10 float64 = float64(1e10)
        _ = v10
    case "\n\t":
        var v11 uint32 = 462
        _ = v11
    }
    switch 604 {
    case -471:
        v3 = (-(v3))
    case 42:
        if true {
            switch (v5 + (v5 + "\n\t")) {
            case "hello":
                if (443 != 2147483648) {
                    v3 -= (v3 ^ (v3 - v3))
                    v3 = v3
                    var v12 int8 = -128
                    _ = v12
                } else {
                    v4[126] = ((("foo bar" + "a") + (v5 + (v5 + "hello"))) + v5)
                    if true {
                        _ = len(v2)
                        switch v5 {
                        case "a":
                            for fi13 := 0; fi13 < 20; fi13++ {
                                v5 = ((("a" + "") + v5) + "a")
                                var v14 uint64 = 174
                                v5 = v5
                                _ = fi13
                                _ = v14
                            }
                        }
                        v3 -= -1
                    }
                }
            default:
                v3 = v3
            }
            fmt.Println(v5, v3)
        } else {
            for fi15 := 0; fi15 < 16; fi15++ {
                var v16 string = (v5 + v5)
                var v17 uint64 = 2
                _ = fi15
                _ = v16
                _ = v17
            }
        }
    case -2147483648:
        for fi18 := 0; fi18 < 11; fi18++ {
            var v19 bool = (!(false || false))
            fmt.Println(v19)
            _ = fi18
            _ = v19
        }
    }
    var v20 int64 = (-(v3))
    if false {
        var v21 uint16 = 550
        fmt.Println(v20, v5)
        v20 -= v20
        _ = v21
    } else {
        if false {
            if ((2 > 529) && true) {
                v22 := GenS1{F0: true, F1: 963}
                delete(v2, 339)
                _ = v22
            }
            fmt.Println(v3, v3)
        } else {
            for fi23 := 0; fi23 < 1; fi23++ {
                switch ("foo bar" + v5) {
                case "":
                    switch fi23 {
                    case 2147483647:
                        var v24 string = ("a" + (v5 + v5))
                        _ = v24
                    case 9223372036854775807:
                        for fi25 := 0; fi25 < 12; fi25++ {
                            switch (fi23 & 0) {
                            case -1:
                                var v26 uint16 = 255
                                _ = v26
                            case 0:
                                if (3.14159 != (1e10 - 0.0)) {
                                    var v27 int32 = 0
                                    v5 = "\n\t"
                                    var v28 float32 = 1e10
                                    _ = v27
                                    _ = v28
                                }
                            default:
                                v29 := GenS1{F0: (true || ("\n\t" == (v5 + v5))), F1: -2147483648}
                                _ = v29
                            }
                            _ = fi25
                        }
                    case 4294967295:
                        for fi30 := 0; fi30 < 18; fi30++ {
                            v31 := []int{fi23}
                            _ = fi30
                            _ = v31
                        }
                    }
                case "foo bar":
                    for fi32 := 0; fi32 < 14; fi32++ {
                        var v33 int64 = (v20 + v3)
                        fmt.Println(fi23, fi32)
                        var v34 uint16 = 226
                        var v35 float64 = (0.0 / 1.0)
                        var v36 string = (((v5 + ("foo bar" + "hello")) + (v5 + v5)) + ("a" + "a"))
                        v1[795] = "\n\t"
                        _ = len(v4)
                        _ = fi32
                        _ = v33
                        _ = v34
                        _ = v35
                        _ = v36
                    }
                case "a":
                    if (!(v3 == (-(v20)))) {
                        for fi37 := 0; fi37 < 19; fi37++ {
                            v3 = (v20 - (v20 ^ v3))
                            _ = fi37
                        }
                        v5 = v5
                        v38 := GenS1{F0: false, F1: -1}
                        _ = v38
                    } else {
                        for fi39 := 0; fi39 < 13; fi39++ {
                            fmt.Println(fi23)
                            _ = fi39
                        }
                    }
                }
                var v40 int32 = -1
                v5 = "a"
                _ = fi23
                _ = v40
            }
        }
    }
    v3 -= -1
    v41 := []int{42}
    v1[729] = v5
    if false {
        _ = len(v4)
    }
    _ = len(v2)
    v3 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v41
    _ = v5
}
