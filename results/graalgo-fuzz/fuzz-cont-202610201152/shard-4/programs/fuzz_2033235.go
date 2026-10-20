package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 bool; F2 bool }
type GenS3 struct { F0 int }
type GenS4 struct { F0 string; F1 int64; F2 uint64 }

func main() {
    var v1 string = "hello"
    var v2 bool = (!(379 >= 65535))
    var v3 int64 = -1
    v4 := []string{(((v1 + v1) + v1) + v1)}
    v5 := []int{2147483647, 0, -385}
    _ = len(v4)
    for fi6 := 0; fi6 < 12; fi6++ {
        switch v1 {
        case "a":
            if true {
                if v2 {
                    v4 = append(v4, "")
                } else {
                    v1 = "hello"
                    fmt.Println(v2, v3, v3)
                    switch (^(fi6)) {
                    case 32768:
                        for fi7 := 0; fi7 < 19; fi7++ {
                            v8 := GenS1{F0: false}
                            _ = len(v4)
                            _ = fi7
                            _ = v8
                        }
                    case 42:
                        switch ("" + "\n\t") {
                        case "foo bar":
                            v9 := GenS2{F0: v1, F1: v2, F2: true}
                            _ = v9
                        }
                    case 9223372036854775807:
                        v10 := GenS3{F0: fi6}
                        _ = v10
                    }
                }
                v1 = v1
                v3 += -9223372036854775808
            } else {
                v2 = (v1 > "hello")
                v11 := GenS3{F0: (-(fi6))}
                _ = v11
            }
        case "":
            v2 = (287 > 491)
        case "foo bar":
            v4 = append(v4, ("\n\t" + v1))
        default:
            switch v1 {
            case "hello":
                var v12 int8 = 0
                _ = v12
            }
        }
        v2 = ((^(fi6)) > fi6)
        _ = fi6
    }
    if v2 {
        fmt.Println(v3, v2)
        v2 = v2
    } else {
        v2 = (1 != 42)
    }
    v1 = (("a" + (v1 + "")) + (v1 + v1))
    switch "a" {
    case "a":
        for fi13 := 0; fi13 < 3; fi13++ {
            var v14 int8 = -128
            for fi15 := 0; fi15 < 16; fi15++ {
                if true {
                    v1 = ("foo bar" + v1)
                }
                _ = fi15
            }
            _ = fi13
            _ = v14
        }
    case "\n\t":
        v1 = (v1 + v1)
    default:
        for fi16 := 0; fi16 < 16; fi16++ {
            if v2 {
                if (1 > 18446744073709551615) {
                    if false {
                        fmt.Println(v3, v1, v2)
                        var v17 uint16 = 0
                        if false {
                            if false {
                                v2 = v2
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            } else {
                                var v18 bool = v2
                                fmt.Println(v1, fi16, v1)
                                _ = v18
                            }
                            if len(v4) > 0 {
                                v4[0] = (v1 + "\n\t")
                            }
                            v1 = "hello"
                        }
                        _ = v17
                    }
                    v19 := GenS2{F0: (v1 + v1), F1: v2, F2: false}
                    for fi20 := 0; fi20 < 3; fi20++ {
                        var v21 float64 = 0.0
                        _ = fi20
                        _ = v21
                    }
                    _ = v19
                } else {
                    v22 := GenS4{F0: "", F1: -1, F2: 0}
                    _ = v22
                }
            }
            fmt.Println(v3)
            _ = fi16
        }
    }
    if v2 {
        if len(v5) > 0 {
            v5[0] = 445
        }
    }
    v1 = "foo bar"
    var v23 int8 = -53
    var v24 int64 = (v3 ^ v3)
    v24 = -1
    var v25 int64 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
