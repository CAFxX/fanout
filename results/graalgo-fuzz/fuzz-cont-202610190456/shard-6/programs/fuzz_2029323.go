package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string; F1 int32; F2 string }

func main() {
    var v1 int = 42
    v2 := []uint32{4294967295, 4294967295}
    var v3 uint8 = 42
    var v4 float32 = 0.0
    v5 := []string{"\n\t"}
    switch -586 {
    case 9223372036854775807:
        switch v1 {
        case 2147483647:
            if true {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            } else {
                for fi6 := 0; fi6 < 16; fi6++ {
                    fmt.Println(v1, v4)
                    var v7 bool = (!(!(!false)))
                    _ = fi6
                    _ = v7
                }
                fmt.Println(v3)
                fmt.Println(v3, v3, v3)
            }
        case 32768:
            v4 = (float32(1.5) + v4)
        default:
            switch (("foo bar" + (("\n\t" + "") + ("hello" + "a"))) + ("" + ("foo bar" + "a"))) {
            case "\n\t":
                if false {
                    var v8 string = "foo bar"
                    v5 = append(v5, (v8 + "a"))
                    _ = v8
                }
            case "a":
                v9 := GenS1{F0: 128, F1: -621}
                _ = v9
            case "foo bar":
                var v10 int32 = 42
                _ = v10
            default:
                if true {
                    v4 = v4
                    var v11 uint16 = 545
                    if true {
                        fmt.Println(v4, v1, v3)
                    } else {
                        switch "foo bar" {
                        case "a":
                            v12 := GenS1{F0: v3, F1: 42}
                            _ = v12
                        }
                        if (1 < 127) {
                            if true {
                                var v13 int = (v1 ^ (v1 & v1))
                                v14 := GenS1{F0: (v3 % (((^(v3))) | 1)), F1: -1}
                                _ = v13
                                _ = v14
                            }
                        }
                    }
                    _ = v11
                } else {
                    var v15 bool = false
                    var v16 uint32 = 42
                    _ = v15
                    _ = v16
                }
            }
        }
    default:
        v3 = 127
    }
    var v17 uint8 = (v3 << ((154) & 63))
    if false {
        fmt.Println(v1, v1)
        v18 := GenS1{F0: v3, F1: -2147483648}
        v19 := GenS1{F0: (^(v3)), F1: -1}
        _ = v18
        _ = v19
    }
    v17 -= (v17 - (^(v17)))
    v17 = v17
    fmt.Println(v3, v4, v1)
    for fi20 := 0; fi20 < 19; fi20++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v17 = (v3 % ((128) | 1))
        v5 = append(v5, "foo bar")
        _ = fi20
    }
    if false {
        v1 -= -9223372036854775808
    } else {
        v21 := GenS1{F0: v3, F1: 421}
        v1 -= -919
        _ = v21
    }
    for fi22 := 0; fi22 < 20; fi22++ {
        var v23 uint64 = 4294967295
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        for fi24 := 0; fi24 < 3; fi24++ {
            v23 -= (^(v23))
            _ = fi24
        }
        _ = fi22
        _ = v23
    }
    for fi25 := 0; fi25 < 5; fi25++ {
        switch 32767 {
        case 32768:
            var v26 int = v1
            _ = v26
        case -12:
            for fi27 := 0; fi27 < 15; fi27++ {
                if (505 < 2) {
                    v3 = 42
                    v3 = v3
                    v1 -= fi25
                } else {
                    v4 = 1.5
                    fmt.Println(v3, v17, v4)
                    var v28 uint8 = 255
                    _ = v28
                }
                fmt.Println(v3, v4, v1)
                _ = fi27
            }
        default:
            v4 += 0.1
        }
        _ = fi25
    }
    var v29 int32 = 2147483647
    v30 := GenS2{F0: "a"}
    var v31 uint = 0
    for fi32 := 0; fi32 < 19; fi32++ {
        var v33 int8 = 72
        for fi34 := 0; fi34 < 16; fi34++ {
            fmt.Println(v33, v33)
            v35 := GenS3{F0: ("" + (("" + "") + "\n\t")), F1: (v29 | v29), F2: ("foo bar" + (("hello" + "a") + "\n\t"))}
            if (1 >= 1) {
                var v36 uint16 = 255
                _ = v36
            } else {
                v17 += 42
                var v37 int32 = (v29 % ((0) | 1))
                _ = v37
            }
            _ = fi34
            _ = v35
        }
        v4 = v4
        _ = fi32
        _ = v33
    }
    var v38 uint16 = 255
    v2 = append(v2, uint32(v31))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v38
    _ = v4
    _ = v5
}
