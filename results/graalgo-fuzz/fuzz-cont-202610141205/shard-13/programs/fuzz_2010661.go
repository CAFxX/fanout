package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool }
func genFn0(p0 int64, p1 string) uint64 {
    var v7 int32 = 2147483647
    switch 4294967295 {
    case 1:
        v8 := GenS1{F0: 9223372036854775807}
        _ = v8
    case 32767:
        var v9 float32 = (0.0 / 0.1)
        _ = v9
    case -717:
        fmt.Println(p0, v7, p1)
    default:
        var v10 float32 = (((0.1 * (-2.25 + float32(1.5))) + 1e10) - 0.0)
        _ = v10
    }
    var v11 uint64 = 484
    for fi12 := 0; fi12 < 12; fi12++ {
        fmt.Println(p0, fi12, fi12)
        _ = fi12
    }
    var v13 int8 = 100
    _ = p0
    _ = p1
    _ = v11
    _ = v13
    _ = v7
    return (^(v11))
}

func main() {
    v1 := []string{"hello", "\n\t", (("hello" + "\n\t") + ((("hello" + "") + "hello") + (("foo bar" + "\n\t") + ""))), "a"}
    var v2 int32 = -2147483648
    v3 := []string{}
    var v4 float64 = -2.25
    v5 := make(map[string]uint32)
    v6 := make(map[int]bool)
    if false {
        fmt.Println(v2, v4)
        v2 += -739
    } else {
        var v14 int = 32767
        _ = v14
    }
    var v15 bool = (!(2147483647 <= 42))
    switch 82 {
    case 2147483647:
        v4 += (v4 / v4)
    case 559:
        if true {
            if v15 {
                var v16 string = "foo bar"
                switch (("\n\t" + "") + (v16 + (v16 + "foo bar"))) {
                case "a":
                    fmt.Println(v16, v4, v15)
                case "":
                    v5[(v16 + ("foo bar" + ((("\n\t" + v16) + v16) + ("foo bar" + v16))))] = 212
                case "\n\t":
                    v17 := GenS2{F0: (!false)}
                    _ = v17
                default:
                    for fi18 := 0; fi18 < 10; fi18++ {
                        v19 := v5[v16]
                        var v20 float32 = -2.25
                        for fi21 := 0; fi21 < 15; fi21++ {
                            if v15 {
                                v5[(("" + v16) + ("\n\t" + "hello"))] = 192
                                var v22 float64 = v4
                                _ = v22
                            }
                            _ = fi21
                        }
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                }
                _ = v16
            } else {
                v2 = -2147483648
                if v15 {
                    if (255 >= 1) {
                        var v23 int64 = 42
                        _ = v23
                    }
                    if len(v3) > 0 {
                        v3[0] = ("\n\t" + "")
                    }
                }
                v2 -= -302
            }
        }
    default:
        v24 := GenS1{F0: 32768}
        _ = v24
    }
    fmt.Println(v4, v15)
    var v25 int64 = 42
    v3 = append(v3, "\n\t")
    v26 := v6[-805]
    fmt.Println(v25, v15, v2)
    var v27 uint8 = 128
    if (((142 < 2) || (!(-1 > -116))) && (!true)) {
        switch ("\n\t" + "") {
        case "\n\t":
            v28 := v6[127]
            _ = v28
        }
    } else {
        v25 = (v25 >> ((65535) & 63))
    }
    delete(v6, 32768)
    var v29 int8 = 1
    switch ((("foo bar" + "a") + ("a" + "hello")) + ("foo bar" + "a")) {
    case "a":
        v26 = v26
    case "foo bar":
        fmt.Println(v29)
    default:
        if len(v3) > 0 {
            v3[0] = "hello"
        }
    }
    for fi30 := 0; fi30 < 13; fi30++ {
        v27 = (^(v27))
        for fi31 := 0; fi31 < 17; fi31++ {
            var v32 int32 = (v2 | (v2 ^ -2147483648))
            _ = fi31
            _ = v32
        }
        _ = fi30
    }
    for fi33 := 0; fi33 < 1; fi33++ {
        v2 -= (-(v2))
        var v34 int64 = -1
        for fi35 := 0; fi35 < 16; fi35++ {
            v36 := GenS2{F0: (true || ((false && v26) || v26))}
            switch "a" {
            case "a":
                fmt.Println(v29, v34)
            }
            if true {
                v4 += (v4 - float64(1.5))
                switch 2147483647 {
                case 0:
                    var v37 uint8 = (v27 & (v27 + v27))
                    _ = v37
                case 182:
                    var v38 string = ("\n\t" + "hello")
                    _ = v38
                default:
                    var v39 uint32 = 862
                    _ = v39
                }
            } else {
                v34 += (v25 - (v25 % (((^(v34))) | 1)))
            }
            _ = fi35
            _ = v36
        }
        _ = fi33
        _ = v34
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
