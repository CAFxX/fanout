package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint32 }

func main() {
    var v1 int8 = -29
    var v2 bool = false
    v3 := make(map[int]bool)
    var v4 uint16 = 32767
    v5 := []int{0, 1}
    var v6 string = "foo bar"
    if (42 != 846) {
        v7 := GenS1{F0: 1}
        _ = v7
    } else {
        switch "\n\t" {
        case "":
            if v2 {
                delete(v3, -203)
                v6 = "hello"
            } else {
                _ = len(v5)
            }
        case "\n\t":
            if len(v5) > 0 {
                v5[0] = 42
            }
        case "foo bar":
            if (9223372036854775807 == 9223372036854775807) {
                switch 32767 {
                case -385:
                    v8 := GenS1{F0: 42}
                    _ = v8
                }
                var v9 int64 = 1
                _ = v9
            }
        }
        var v10 int64 = 0
        fmt.Println(v2, v6, v10)
        _ = v10
    }
    var v11 int32 = -2147483648
    fmt.Println(v6, v1)
    switch 42 {
    case 2147483647:
        v2 = (869 != 18446744073709551615)
    case 32767:
        var v12 int = 9223372036854775807
        _ = v12
    default:
        v13 := GenS2{F0: -635}
        _ = v13
    }
    if v2 {
        v11 = v11
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    } else {
        var v14 int64 = 0
        _ = v14
    }
    delete(v3, -616)
    if v2 {
        if v2 {
            if (v2 || false) {
                _ = len(v3)
            }
        } else {
            for fi15 := 0; fi15 < 20; fi15++ {
                var v16 uint32 = 65535
                for fi17 := 0; fi17 < 14; fi17++ {
                    var v18 int64 = -9223372036854775808
                    var v19 string = (v6 + (("" + "a") + v6))
                    switch fi17 {
                    case 2147483647:
                        if (!(4294967295 == 2)) {
                            switch "hello" {
                            case "foo bar":
                                fmt.Println(v11, v11, v19)
                            case "":
                                fmt.Println(v18)
                            case "a":
                                var v20 uint16 = (v4 / ((65535) | 1))
                                var v21 uint = 2
                                if (!((!true) || v2)) {
                                    if len(v5) > 0 {
                                        v5[0] = 128
                                    }
                                    fmt.Println(fi15)
                                    if (827 >= 42) {
                                        v22 := GenS3{F0: (fi15 / (((fi17 >> ((1) & 63))) | 1))}
                                        _ = v22
                                    } else {
                                        v1 = v1
                                        v23 := v3[fi17]
                                        _ = v23
                                    }
                                } else {
                                    var v24 uint64 = 1
                                    v20 += v4
                                    fmt.Println(v2, fi17, v21)
                                    v25 := GenS1{F0: v16}
                                    _ = v24
                                    _ = v25
                                }
                                v21 += 4294967295
                                _ = v20
                                _ = v21
                            }
                        }
                    case -521:
                        if false {
                            var v26 int32 = v11
                            switch (fi17 / ((42) | 1)) {
                            case 32767:
                                fmt.Println(v1)
                            case -916:
                                v27 := GenS1{F0: (v16 | 1)}
                                _ = v27
                            case 128:
                                v1 = v1
                            default:
                                var v28 int64 = (v18 % ((1) | 1))
                                _ = v28
                            }
                            if true {
                                if v2 {
                                    v5 = append(v5, 2147483647)
                                } else {
                                    var v29 uint16 = v4
                                    var v30 float64 = 3.14159
                                    if v2 {
                                        var v31 bool = (true && ((((v30 + 1e10) * 0.1) / (0.1 / v30)) < ((v30 / (-2.25 + v30)) / float64(1.5))))
                                        var v32 bool = true
                                        if false {
                                            if len(v5) > 0 {
                                                v5[0] = 116
                                            }
                                        } else {
                                            v33 := GenS4{F0: (v16 << ((255) & 63))}
                                            if false {
                                                var v34 float32 = 0.0
                                                _ = v34
                                            } else {
                                                var v35 int16 = 1
                                                var v36 uint = 528
                                                _ = v35
                                                _ = v36
                                            }
                                            _ = v33
                                        }
                                        _ = v31
                                        _ = v32
                                    }
                                    _ = v29
                                    _ = v30
                                }
                            }
                            _ = v26
                        }
                    }
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                _ = fi15
                _ = v16
            }
            v6 = v6
        }
    } else {
        v6 = (v6 + v6)
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        for fi37 := 0; fi37 < 7; fi37++ {
            if (41 <= (v1 - v1)) {
                v38 := v3[0]
                for fi39 := 0; fi39 < 9; fi39++ {
                    _ = len(v3)
                    _ = fi39
                }
                _ = v38
            }
            _ = fi37
        }
    }
    v40 := v3[-428]
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
