package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 uint16 = 65535
    var v2 int16 = -448
    var v3 uint32 = 95
    var v4 int32 = 921
    var v5 uint16 = 65535
    var v6 int = 0
    var v7 int64 = 9223372036854775807
    if true {
        if ((!true) || (false && true)) {
            v2 = (v2 - v2)
            switch "" {
            case "\n\t":
                for fi8 := 0; fi8 < 5; fi8++ {
                    v7 += (^(v7))
                    v9 := []int{v6}
                    var v10 string = (("a" + "") + (("foo bar" + "a") + ((((("hello" + "\n\t") + "") + (("\n\t" + "foo bar") + "foo bar")) + "a") + "hello")))
                    _ = fi8
                    _ = v10
                    _ = v9
                }
            case "hello":
                v4 = 42
            default:
                switch (v6 % (((-(v6))) | 1)) {
                case -2147483648:
                    var v11 uint = 190
                    _ = v11
                case -555:
                    v5 -= v1
                default:
                    var v12 string = ("a" + ("a" + ("\n\t" + (("hello" + "") + "foo bar"))))
                    _ = v12
                }
            }
        } else {
            switch v6 {
            case 2147483647:
                if (147 >= 42) {
                    v4 -= -156
                    if true {
                        if (false || (121 > 107)) {
                            v13 := []int{v6, int(v2), (v6 + 42)}
                            _ = v13
                        } else {
                            fmt.Println(v4, v7, v1)
                        }
                    } else {
                        var v14 int16 = (v2 + v2)
                        fmt.Println(v1, v6, v5)
                        var v15 int32 = (-(v4))
                        _ = v14
                        _ = v15
                    }
                } else {
                    if false {
                        fmt.Println(v3, v6, v1)
                        fmt.Println(v1)
                    } else {
                        switch "foo bar" {
                        case "a":
                            var v16 int64 = v7
                            _ = v16
                        case "\n\t":
                            var v17 uint8 = 0
                            _ = v17
                        default:
                            for fi18 := 0; fi18 < 8; fi18++ {
                                v3 -= 319
                                v4 -= (v4 | (v4 * -1))
                                fmt.Println(fi18)
                                _ = fi18
                            }
                        }
                    }
                }
            case 0:
                var v19 string = ("" + "foo bar")
                _ = v19
            case 32767:
                v20 := []int{109, v6, v6, -1}
                _ = v20
            }
            var v21 int64 = v7
            _ = v21
        }
    }
    if true {
        v5 = (v5 >> ((953) & 63))
    } else {
        for fi22 := 0; fi22 < 1; fi22++ {
            v23 := []int{fi22, 127, (fi22 >> ((4294967295) & 63)), (-(fi22))}
            _ = fi22
            _ = v23
        }
        var v24 int64 = (v7 % ((v7) | 1))
        if (v2 >= (v2 | v2)) {
            v25 := GenS1{F0: 170}
            _ = v25
        } else {
            fmt.Println(v7)
        }
        _ = v24
    }
    fmt.Println(v1, v5, v6)
    v2 -= 1
    v4 += v4
    v3 = v3
    var v26 string = "foo bar"
    fmt.Println(v3, v4)
    if (0.0 < (0.1 * float64(((0.0 * (3.14159 / ((-2.25 / 0.1) + (-2.25 + (0.1 / 1.5))))) * (0.1 - 3.14159))))) {
        v27 := []int{(v6 + v6), v6}
        for fi28 := 0; fi28 < 19; fi28++ {
            if ("hello" != ("foo bar" + "\n\t")) {
                var v29 string = ("" + ((v26 + v26) + v26))
                for fi30 := 0; fi30 < 14; fi30++ {
                    switch v6 {
                    case -9223372036854775808:
                        var v31 uint64 = 2
                        _ = v31
                    case 1:
                        var v32 uint64 = 798
                        _ = v32
                    }
                    var v33 int16 = 510
                    fmt.Println(v33, fi30, fi28)
                    _ = fi30
                    _ = v33
                }
                v7 = (v7 ^ (v7 - int64(v6)))
                _ = v29
            }
            v34 := GenS2{F0: 0}
            var v35 uint32 = (^(v3))
            _ = fi28
            _ = v34
            _ = v35
        }
        var v36 float64 = 0.0
        _ = v27
        _ = v36
    }
    var v37 uint16 = 32767
    var v38 uint32 = (v3 - (^(v3)))
    fmt.Println(v2)
    v39 := []int{603, (^(v6)), v6}
    var v40 uint = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
}
