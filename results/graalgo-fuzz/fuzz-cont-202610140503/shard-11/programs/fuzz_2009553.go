package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0() uint16 {
    switch -1 {
    case 182:
        switch (("" + ("a" + "")) + "") {
        case "hello":
            v4 := []int{893, 806, 9223372036854775807}
            _ = v4
        case "":
            if true {
                if true {
                    var v5 uint8 = 128
                    _ = v5
                }
            }
        default:
            var v6 uint32 = 0
            _ = v6
        }
    }
    switch -1 {
    case 0:
        var v7 float64 = -2.25
        _ = v7
    case 1:
        var v8 bool = (false || ((0 != 4294967295) && false))
        _ = v8
    }
    var v9 float64 = (0.1 * 0.1)
    var v10 bool = (469 != 2147483647)
    fmt.Println(v9)
    _ = v10
    _ = v9
    return 6
}

func main() {
    v1 := make(map[int]bool)
    v2 := []string{("foo bar" + ""), "", ("" + ("foo bar" + "hello")), "hello"}
    var v3 int32 = 2147483647
    v3 += -2147483648
    for fi11 := 0; fi11 < 9; fi11++ {
        fmt.Println(fi11, fi11)
        _ = fi11
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v3)
    v2 = append(v2, "a")
    delete(v1, 753)
    v3 += v3
    if false {
        var v12 int = -838
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v12
    }
    var v13 int32 = v3
    if ((false || (-73 > -128)) || (0 > genFn0())) {
        if (!true) {
            switch "" {
            case "foo bar":
                v14 := GenS1{F0: 42}
                _ = v14
            case "hello":
                if false {
                    fmt.Println(v13)
                    if (-1 < 78) {
                        for fi15 := 0; fi15 < 2; fi15++ {
                            if false {
                                v13 = -2147483648
                                var v16 uint8 = 1
                                _ = v16
                            }
                            _ = fi15
                        }
                        if (true && (1 != 2)) {
                            for fi17 := 0; fi17 < 12; fi17++ {
                                fmt.Println(v3, v3)
                                _ = fi17
                            }
                            v18 := GenS1{F0: 705}
                            var v19 int16 = -1
                            _ = v18
                            _ = v19
                        }
                        v1[492] = true
                    } else {
                        v3 = (v3 * -2147483648)
                        switch -9223372036854775808 {
                        case -1:
                            var v20 float64 = (float64(1e10) / 1.5)
                            _ = v20
                        case 127:
                            if len(v2) > 0 {
                                v2[0] = "foo bar"
                            }
                        case -9223372036854775808:
                            for fi21 := 0; fi21 < 19; fi21++ {
                                v3 += -289
                                _ = fi21
                            }
                        }
                        delete(v1, 2147483647)
                    }
                }
            default:
                v22 := GenS1{F0: 128}
                _ = v22
            }
            var v23 float32 = (0.0 / ((0.1 / -2.25) - (1.5 / 1.0)))
            v3 += v13
            _ = v23
        }
        v24 := GenS1{F0: 42}
        v3 -= v3
        _ = v24
    } else {
        v13 += -2147483648
        fmt.Println(v3)
        for fi25 := 0; fi25 < 7; fi25++ {
            if false {
                delete(v1, fi25)
                if true {
                    fmt.Println(fi25, v13)
                    switch fi25 {
                    case -77:
                        delete(v1, 127)
                    default:
                        _ = len(v2)
                    }
                    if (true && false) {
                        fmt.Println(v13, fi25)
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v26 uint64 = 0
                        _ = v26
                    }
                }
                fmt.Println(fi25, v3, fi25)
            }
            v27 := GenS1{F0: fi25}
            _ = fi25
            _ = v27
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v28 uint16 = 621
    v2 = append(v2, ("foo bar" + "a"))
    v2 = append(v2, "")
    v3 = (v3 * (v3 - (^(v13))))
    fmt.Println(v28, v3)
    var v29 int64 = 9223372036854775807
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v2
    _ = v28
    _ = v29
    _ = v3
}
