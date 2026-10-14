package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8 }

func main() {
    var v1 uint64 = 2
    var v2 int8 = 100
    var v3 string = ""
    v4 := make(map[uint32]string)
    v5 := []int{-1, 1, 1}
    switch 32768 {
    case 0:
        v2 -= v2
    case 42:
        if true {
            v6 := GenS1{F0: 42}
            _ = v6
        }
    default:
        v1 = v1
    }
    var v7 uint64 = 1
    fmt.Println(v2, v2, v7)
    v8 := GenS1{F0: 4294967295}
    v7 = v7
    v1 = v1
    v9 := GenS1{F0: 18446744073709551615}
    fmt.Println(v1, v2, v3)
    fmt.Println(v7, v7, v3)
    var v10 uint16 = 32767
    if (!true) {
        for fi11 := 0; fi11 < 13; fi11++ {
            v12 := GenS1{F0: 65535}
            for fi13 := 0; fi13 < 15; fi13++ {
                v3 = "\n\t"
                fmt.Println(v1)
                _ = fi13
            }
            if (!(-1 < -9223372036854775808)) {
                var v14 int = 128
                v1 = v7
                for fi15 := 0; fi15 < 13; fi15++ {
                    v1 -= v1
                    delete(v4, 2)
                    var v16 uint64 = v7
                    _ = fi15
                    _ = v16
                }
                _ = v14
            } else {
                var v17 int64 = 42
                var v18 uint = 42
                delete(v4, uint32(v1))
                _ = v17
                _ = v18
            }
            _ = fi11
            _ = v12
        }
        var v19 uint64 = v7
        _ = v19
    } else {
        var v20 bool = (true && ((!((true || ((!(2 <= 42)) && (true || (65535 != 65535)))) && (42 > -627))) && (9223372036854775807 <= 9223372036854775807)))
        _ = v20
    }
    if (169 <= 0) {
        v4[347] = ("foo bar" + v3)
    } else {
        if false {
            switch -1 {
            case 1:
                for fi21 := 0; fi21 < 12; fi21++ {
                    for fi22 := 0; fi22 < 1; fi22++ {
                        if false {
                            v7 = v1
                            v10 = v10
                        }
                        var v23 int64 = 294
                        if true {
                            fmt.Println(fi22, v10, v3)
                            if true {
                                fmt.Println(v3)
                            } else {
                                var v24 string = ""
                                if true {
                                    v7 = 18446744073709551615
                                } else {
                                    switch (-(fi21)) {
                                    case 9223372036854775807:
                                        var v25 int16 = -32768
                                        _ = v25
                                    case 127:
                                        if false {
                                            if (!((((true && (1.5 >= 0.0)) || ((float32(3.14159) + 3.14159) != -2.25)) && false) && (42 >= 0))) {
                                                fmt.Println(v23)
                                            } else {
                                                var v26 uint32 = 2
                                                v27 := GenS2{F0: -1}
                                                if (false && false) {
                                                    if (3.14159 >= 0.0) {
                                                        if ((^(v2)) >= v2) {
                                                            var v28 int64 = v23
                                                            v23 = v23
                                                            _ = v28
                                                        } else {
                                                            var v29 float32 = 1e10
                                                            v29 = 3.14159
                                                            _ = v29
                                                        }
                                                    }
                                                }
                                                _ = v26
                                                _ = v27
                                            }
                                            var v30 uint64 = 1
                                            _ = v30
                                        } else {
                                            if len(v5) > 0 {
                                                v5[0] = (fi22 / ((fi22) | 1))
                                            }
                                            var v31 float64 = (1.5 * (0.0 / (0.1 * (0.0 - 0.1))))
                                            fmt.Println(fi22)
                                            _ = v31
                                        }
                                    }
                                }
                                switch "hello" {
                                case "":
                                    var v32 int32 = int32(fi22)
                                    _ = v32
                                case "foo bar":
                                    var v33 int8 = v2
                                    _ = v33
                                }
                                if (false && true) {
                                    v10 = 32767
                                    fmt.Println(fi22)
                                    var v34 uint32 = 542
                                    _ = v34
                                }
                                _ = v24
                            }
                            v23 = v23
                        }
                        _ = fi22
                        _ = v23
                    }
                    _ = fi21
                }
            case -9223372036854775808:
                var v35 float32 = (1.5 + 3.14159)
                _ = v35
            default:
                fmt.Println(v10, v3, v2)
            }
            v36 := GenS1{F0: 1}
            _ = v36
        }
        v7 = 538
        delete(v4, 383)
    }
    v4[533] = (("foo bar" + v3) + (v3 + ("foo bar" + v3)))
    _ = len(v4)
    _ = len(v5)
    switch -548 {
    case 9223372036854775807:
        var v37 int64 = -1
        _ = v37
    case 4294967295:
        fmt.Println(v1, v1)
    case 767:
        switch -1 {
        case 1:
            v7 += (v1 % (((v1 / (((v7 / ((172) | 1))) | 1))) | 1))
        case 0:
            for fi38 := 0; fi38 < 3; fi38++ {
                fmt.Println(v2)
                v39 := GenS1{F0: 42}
                _ = fi38
                _ = v39
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
