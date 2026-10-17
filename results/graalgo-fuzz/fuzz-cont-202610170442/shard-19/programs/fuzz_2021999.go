package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 uint64; F1 int64; F2 string }
type GenS3 struct { F0 int8; F1 int8; F2 int8 }

func main() {
    var v1 int16 = 1000
    var v2 string = ("\n\t" + "\n\t")
    var v3 uint = 255
    var v4 string = v2
    v4 = ("foo bar" + ((v4 + ("a" + (v4 + ((v2 + "foo bar") + "foo bar")))) + "foo bar"))
    for fi5 := 0; fi5 < 2; fi5++ {
        var v6 int32 = 1
        v1 -= -32768
        fmt.Println(v2, v1)
        _ = fi5
        _ = v6
    }
    if (32767 < 4294967295) {
        v1 = (v1 + -1)
    }
    v7 := []int{285}
    fmt.Println(v3)
    fmt.Println(v3)
    var v8 uint32 = 1
    if true {
        var v9 string = (v4 + (v4 + "foo bar"))
        fmt.Println(v2, v9, v4)
        if len(v7) > 0 {
            v7[0] = -2147483648
        }
        _ = v9
    }
    _ = len(v7)
    if true {
        v10 := GenS1{F0: "hello", F1: v1}
        _ = v10
    }
    if true {
        for fi11 := 0; fi11 < 17; fi11++ {
            fmt.Println(v3)
            v4 = (("\n\t" + v2) + ("foo bar" + "a"))
            _ = fi11
        }
    } else {
        for fi12 := 0; fi12 < 6; fi12++ {
            v13 := GenS2{F0: 2, F1: 42, F2: ""}
            _ = fi12
            _ = v13
        }
        if (false && ((v3 > (v3 ^ (v3 & v3))) && (!false))) {
            for fi14 := 0; fi14 < 9; fi14++ {
                for fi15 := 0; fi15 < 6; fi15++ {
                    var v16 uint32 = 864
                    var v17 int16 = (v1 + v1)
                    fmt.Println(v8, v8, v17)
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                if true {
                    if ((v3 ^ (v3 / ((v3) | 1))) < (v3 % (((v3 << (((v3 << ((v3) & 63))) & 63))) | 1))) {
                        v18 := GenS1{F0: v4, F1: (v1 & v1)}
                        if false {
                            if (true && (0 > 2)) {
                                v2 = v2
                                v7 = append(v7, (fi14 - fi14))
                                switch v4 {
                                case "":
                                    v19 := GenS1{F0: (v2 + "foo bar"), F1: v1}
                                    _ = v19
                                case "foo bar":
                                    for fi20 := 0; fi20 < 14; fi20++ {
                                        switch v2 {
                                        case "\n\t":
                                            fmt.Println(fi20, fi20, v2)
                                        case "a":
                                            v1 += (^(v1))
                                        }
                                        if false {
                                            fmt.Println(v1, v1, v4)
                                            v3 += v3
                                            v4 = ("a" + (v4 + ((v4 + ((v4 + "") + "foo bar")) + "hello")))
                                        } else {
                                            var v21 int32 = 1
                                            _ = v21
                                        }
                                        _ = fi20
                                    }
                                default:
                                    var v22 uint16 = 65535
                                    _ = v22
                                }
                            }
                            var v23 uint32 = 4294967295
                            v2 = ""
                            _ = v23
                        }
                        for fi24 := 0; fi24 < 2; fi24++ {
                            v3 += (v3 >> (((v3 - 833)) & 63))
                            _ = fi24
                        }
                        _ = v18
                    }
                }
                _ = fi14
            }
            v8 = (^(v8))
            if len(v7) > 0 {
                v7[0] = -9223372036854775808
            }
        } else {
            for fi25 := 0; fi25 < 14; fi25++ {
                _ = len(v7)
                v4 = ((v2 + ((v4 + "hello") + v4)) + (v4 + v2))
                v26 := GenS3{F0: int8(v1), F1: 0, F2: 42}
                _ = fi25
                _ = v26
            }
            fmt.Println(v8, v4)
            var v27 uint = (v3 - (^(v3)))
            _ = v27
        }
    }
    v28 := GenS1{F0: ((v4 + v4) + v2), F1: (v1 & v1)}
    var v29 int64 = -738
    var v30 int8 = 1
    var v31 int64 = (v29 >> (((v3 ^ v3)) & 63))
    for fi32 := 0; fi32 < 8; fi32++ {
        for fi33 := 0; fi33 < 9; fi33++ {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = fi33
        }
        switch -1 {
        case 2147483647:
            fmt.Println(v1, v1)
        case 42:
            v7 = append(v7, fi32)
        }
        fmt.Println(v4, v3)
        _ = fi32
    }
    var v34 int32 = int32(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v4
    _ = v7
    _ = v8
}
