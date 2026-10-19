package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 uint }

func main() {
    var v1 int64 = 0
    var v2 int16 = 929
    var v3 int16 = (v2 + v2)
    var v4 uint = 42
    v5 := make(map[string]string)
    var v6 uint = (v4 / ((18446744073709551615) | 1))
    var v7 int16 = v3
    fmt.Println(v4, v6, v3)
    for fi8 := 0; fi8 < 13; fi8++ {
        fmt.Println(v3, v3)
        var v9 int = (fi8 & (fi8 >> ((v4) & 63)))
        _ = fi8
        _ = v9
    }
    if true {
        var v10 int32 = -1
        if ((!(42 >= v6)) || true) {
            var v11 float64 = (3.14159 + 1.5)
            fmt.Println(v11)
            _ = v11
        }
        var v12 int16 = v7
        _ = v10
        _ = v12
    }
    if (0 != 1) {
        var v13 uint32 = 666
        for fi14 := 0; fi14 < 11; fi14++ {
            for fi15 := 0; fi15 < 3; fi15++ {
                v1 = 42
                v3 = -1
                switch ("\n\t" + "") {
                case "hello":
                    v16 := GenS1{F0: 0, F1: v6}
                    v17 := []int{fi15, fi14, -1}
                    _ = v16
                    _ = v17
                case "a":
                    v7 += v2
                }
                _ = fi15
            }
            if (true || true) {
                if false {
                    var v18 int = fi14
                    if false {
                        fmt.Println(fi14, v13, v7)
                        switch "a" {
                        case "\n\t":
                            switch 2147483647 {
                            case 2147483647:
                                if ((-689 >= 2147483647) && ((v1 % ((-251) | 1)) > v1)) {
                                    v19 := GenS1{F0: -2147483648, F1: v6}
                                    v20 := []int{}
                                    _ = v19
                                    _ = v20
                                }
                            case 32767:
                                fmt.Println(v6, v7)
                            }
                        default:
                            v5["hello"] = "foo bar"
                        }
                        if (-2147483648 != 748) {
                            for fi21 := 0; fi21 < 10; fi21++ {
                                var v22 uint64 = 1
                                fmt.Println(v4, v22)
                                _ = fi21
                                _ = v22
                            }
                            var v23 int = v18
                            _ = v23
                        } else {
                            fmt.Println(v3, v18, v13)
                            var v24 int = (v18 & v18)
                            v25 := GenS2{F0: v24, F1: -1}
                            _ = v24
                            _ = v25
                        }
                    } else {
                        fmt.Println(v3, v2, v18)
                        if true {
                            v5["foo bar"] = (("a" + (("\n\t" + "a") + ("foo bar" + "a"))) + ("\n\t" + "a"))
                            for fi26 := 0; fi26 < 13; fi26++ {
                                var v27 uint8 = uint8(v6)
                                v27 -= v27
                                _ = fi26
                                _ = v27
                            }
                        } else {
                            v7 = -51
                        }
                        fmt.Println(v7, v1)
                    }
                    v3 -= (v3 << ((18446744073709551615) & 63))
                    _ = v18
                }
                fmt.Println(v2)
            } else {
                var v28 int64 = int64(v6)
                v7 = v7
                v29 := []int{}
                _ = v28
                _ = v29
            }
            v4 = (v4 ^ (v6 << (((v6 - 345)) & 63)))
            _ = fi14
        }
        v7 -= 117
        _ = v13
    }
    fmt.Println(v2)
    v30 := GenS1{F0: 254, F1: (v6 % ((v4) | 1))}
    var v31 int = 42
    v31 = (v31 ^ -328)
    for fi32 := 0; fi32 < 3; fi32++ {
        if true {
            v33 := GenS2{F0: (v31 << ((v4) & 63)), F1: -32768}
            v34 := GenS2{F0: (v31 << ((4294967295) & 63)), F1: 0}
            v3 -= v7
            _ = v33
            _ = v34
        } else {
            if true {
                for fi35 := 0; fi35 < 2; fi35++ {
                    v36 := GenS3{F0: (v4 | v4)}
                    v7 = (v2 / ((0) | 1))
                    _ = fi35
                    _ = v36
                }
                var v37 uint32 = 831
                _ = v37
            }
        }
        v2 = v7
        _ = fi32
    }
    var v38 float64 = (3.14159 + 0.1)
    var v39 float64 = v38
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
