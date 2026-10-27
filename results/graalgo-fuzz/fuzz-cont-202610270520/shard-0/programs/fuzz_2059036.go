package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 int16; F1 bool; F2 uint64 }

func main() {
    var v1 uint8 = 0
    var v2 float64 = 1e10
    var v3 int = -9223372036854775808
    var v4 int64 = 0
    for fi5 := 0; fi5 < 5; fi5++ {
        v3 += fi5
        v1 -= (v1 << ((1) & 63))
        var v6 uint8 = v1
        _ = fi5
        _ = v6
    }
    var v7 float64 = -2.25
    fmt.Println(v3, v3)
    fmt.Println(v3, v7, v4)
    v8 := GenS1{F0: "hello"}
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 bool = false
        v11 := []int{(fi9 & (fi9 | fi9)), -215}
        for fi12 := 0; fi12 < 18; fi12++ {
            v11 = append(v11, (v3 - (-(fi12))))
            if v10 {
                switch (^(v3)) {
                case 42:
                    var v13 int16 = int16(v4)
                    _ = v13
                case 9223372036854775807:
                    switch ("\n\t" + "foo bar") {
                    case "foo bar":
                        if v10 {
                            var v14 uint16 = 32767
                            _ = v14
                        }
                    default:
                        v3 = (fi12 | (-(v3)))
                    }
                case 127:
                    v3 -= 4294967295
                default:
                    if true {
                        var v15 string = ("foo bar" + (("a" + (("" + ("foo bar" + "foo bar")) + "foo bar")) + (((("foo bar" + "") + "foo bar") + "foo bar") + "a")))
                        var v16 bool = ((((v1 | 79) > (^(v1))) || ("foo bar" == "a")) && true)
                        switch (v15 + ("\n\t" + v15)) {
                        case "foo bar":
                            v17 := GenS1{F0: "hello"}
                            _ = v17
                        default:
                            var v18 bool = true
                            _ = v18
                        }
                        _ = v15
                        _ = v16
                    } else {
                        if true {
                            v19 := GenS1{F0: "\n\t"}
                            _ = v19
                        }
                        switch "hello" {
                        case "":
                            v20 := GenS2{F0: "\n\t", F1: 919}
                            _ = v20
                        default:
                            if v10 {
                                if (4294967295 == uint32(v1)) {
                                    fmt.Println(fi12, v10)
                                } else {
                                    switch "\n\t" {
                                    case "a":
                                        v1 += 127
                                    }
                                    var v21 int8 = 100
                                    _ = v21
                                }
                                var v22 uint64 = 2
                                _ = v22
                            }
                        }
                        if false {
                            if len(v11) > 0 {
                                v11[0] = (^(v3))
                            }
                            _ = len(v11)
                        } else {
                            v4 = v4
                            fmt.Println(v10)
                            var v23 uint = 4294967295
                            _ = v23
                        }
                        v1 += (^(v1))
                    }
                }
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                v1 = v1
            }
            v24 := GenS3{F0: 42, F1: (!((!(v2 <= v2)) && v10)), F2: 325}
            fmt.Println(v3, v1)
            _ = fi12
            _ = v24
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    fmt.Println(v3)
    for fi25 := 0; fi25 < 1; fi25++ {
        fmt.Println(v1, v7)
        _ = fi25
    }
    var v26 uint16 = 738
    if false {
        var v27 uint64 = 1
        fmt.Println(v7)
        _ = v27
    }
    var v28 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v26
    _ = v28
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
