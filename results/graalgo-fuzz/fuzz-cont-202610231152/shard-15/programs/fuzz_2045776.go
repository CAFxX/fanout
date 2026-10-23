package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0() int64 {
    var v7 string = "a"
    for fi8 := 0; fi8 < 13; fi8++ {
        v7 = (((v7 + "") + v7) + "")
        fmt.Println(v7)
        for fi9 := 0; fi9 < 4; fi9++ {
            fmt.Println(fi8)
            v10 := GenS1{F0: 154}
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    _ = v7
    return 0
}

func main() {
    var v1 int = 1
    var v2 int32 = -1
    var v3 bool = (!(0 >= -9223372036854775808))
    var v4 int64 = 9223372036854775807
    var v5 int8 = 43
    v6 := []string{"foo bar", "hello", ""}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println(v2, v3)
    var v11 string = "hello"
    v12 := GenS1{F0: 9223372036854775807}
    v4 = (v4 * v4)
    if len(v6) > 0 {
        v6[0] = (v11 + v11)
    }
    fmt.Println(v5, v1)
    v13 := GenS1{F0: (v4 ^ 42)}
    v6 = append(v6, v11)
    for fi14 := 0; fi14 < 2; fi14++ {
        v2 += v2
        _ = fi14
    }
    var v15 int16 = 42
    if (155 != 127) {
        v16 := GenS1{F0: v4}
        for fi17 := 0; fi17 < 5; fi17++ {
            fmt.Println(v4, v4, fi17)
            _ = fi17
        }
        fmt.Println(v4, v15)
        _ = v16
    } else {
        v18 := GenS1{F0: genFn0()}
        var v19 int = v1
        _ = v18
        _ = v19
    }
    var v20 uint = 952
    if (v3 && ((!(42 < 4294967295)) && true)) {
        switch "foo bar" {
        case "\n\t":
            var v21 int32 = 645
            _ = v21
        case "foo bar":
            switch v1 {
            case -2147483648:
                if v3 {
                    v22 := GenS1{F0: genFn0()}
                    _ = v22
                } else {
                    var v23 bool = true
                    v15 -= (v15 - (^(v15)))
                    _ = v23
                }
            case 42:
                v20 = v20
            default:
                var v24 uint64 = 2
                _ = v24
            }
        case "":
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        default:
            fmt.Println(v5, v20, v2)
        }
        switch "foo bar" {
        case "foo bar":
            if false {
                if len(v6) > 0 {
                    v6[0] = (v11 + "")
                }
                fmt.Println(v5, v20, v11)
            }
        case "":
            v4 += (v4 / (((v4 - (^(v4)))) | 1))
        case "hello":
            var v25 int = (v1 - 32767)
            _ = v25
        default:
            fmt.Println(v15, v2)
        }
    } else {
        v5 += (v5 << ((340) & 63))
        var v26 bool = ((1 >= 42) && ((v1 ^ (v1 & (v1 | (^(v1))))) >= -849))
        var v27 uint16 = 162
        _ = v26
        _ = v27
    }
    var v28 string = (v11 + "hello")
    fmt.Println(v3, v28, v20)
    if (624 != 42) {
        var v29 int32 = v2
        _ = v29
    } else {
        for fi30 := 0; fi30 < 7; fi30++ {
            for fi31 := 0; fi31 < 14; fi31++ {
                v3 = (42 == 42)
                _ = fi31
            }
            for fi32 := 0; fi32 < 13; fi32++ {
                v20 = (v20 | v20)
                var v33 uint64 = 43
                _ = fi32
                _ = v33
            }
            v34 := GenS1{F0: 1}
            _ = fi30
            _ = v34
        }
        v35 := GenS1{F0: (v4 << ((v20) & 63))}
        for fi36 := 0; fi36 < 13; fi36++ {
            if (v2 >= (v2 >> ((42) & 63))) {
                if ((v15 / ((int16(v4)) | 1)) <= v15) {
                    if v3 {
                        v6 = append(v6, v28)
                        if len(v6) > 0 {
                            v6[0] = (v11 + v28)
                        }
                        fmt.Println(v1)
                    } else {
                        var v37 uint64 = 2
                        _ = v37
                    }
                }
            }
            fmt.Println(v1)
            for fi38 := 0; fi38 < 18; fi38++ {
                v5 -= (-(v5))
                var v39 uint16 = 32767
                fmt.Println(v2, fi36)
                _ = fi38
                _ = v39
            }
            _ = fi36
        }
        _ = v35
    }
    v6 = append(v6, "a")
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
