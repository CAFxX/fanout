package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 string }
type GenS2 struct { F0 int8; F1 int16; F2 uint64 }
type GenS3 struct { F0 int32; F1 uint16; F2 bool }
func genFn0(p0 uint16) uint16 {
    if false {
        var v5 uint16 = (p0 + (^(p0)))
        v6 := []int{}
        _ = v5
        _ = v6
    } else {
        var v7 int8 = 0
        fmt.Println(v7, v7)
        _ = v7
    }
    var v8 int = -9223372036854775808
    var v9 int32 = 145
    if (v9 < v9) {
        var v10 bool = false
        _ = v10
    }
    _ = p0
    _ = v8
    _ = v9
    return (p0 >> ((1) & 63))
}

func main() {
    var v1 uint32 = 2147483648
    var v2 uint16 = 65535
    var v3 uint64 = 2
    v4 := []string{("" + ("foo bar" + "foo bar"))}
    switch 2147483647 {
    case -9223372036854775808:
        var v11 uint8 = 2
        _ = v11
    case 1:
        v12 := GenS1{F0: 1, F1: 42, F2: ("hello" + "\n\t")}
        _ = v12
    case -2147483648:
        switch 862 {
        case 4294967295:
            for fi13 := 0; fi13 < 4; fi13++ {
                v2 = v2
                _ = fi13
            }
        default:
            switch 0 {
            case 0:
                var v14 int64 = -9223372036854775808
                _ = v14
            case 546:
                v1 -= 42
            default:
                v2 = uint16(v1)
            }
        }
    }
    switch 0 {
    case 2147483647:
        v2 = v2
    default:
        if len(v4) > 0 {
            v4[0] = ("hello" + (("foo bar" + "hello") + "foo bar"))
        }
    }
    v3 = (^(v3))
    v2 = (v2 - (v2 + (v2 * (v2 + (v2 / (((v2 ^ 42)) | 1))))))
    switch ("foo bar" + "hello") {
    case "":
        v15 := GenS1{F0: 32767, F1: 9223372036854775807, F2: "hello"}
        _ = v15
    case "a":
        var v16 bool = false
        _ = v16
    case "hello":
        if (42 != uint8(v3)) {
            switch "" {
            case "foo bar":
                var v17 uint = 435
                _ = v17
            case "a":
                for fi18 := 0; fi18 < 2; fi18++ {
                    v2 = (^(v2))
                    if (65535 >= (v1 % (((v1 >> ((0) & 63))) | 1))) {
                        for fi19 := 0; fi19 < 14; fi19++ {
                            var v20 float32 = 1.5
                            if (!(!(65535 != (v1 ^ v1)))) {
                                fmt.Println(v2, v2)
                                switch "\n\t" {
                                case "hello":
                                    fmt.Println(v1, fi18, v3)
                                case "\n\t":
                                    var v21 uint64 = (^(v3))
                                    _ = v21
                                }
                                var v22 uint8 = 2
                                _ = v22
                            } else {
                                var v23 uint8 = 127
                                fmt.Println(v20, v3, v3)
                                _ = v23
                            }
                            fmt.Println(fi18, v1, v2)
                            _ = fi19
                            _ = v20
                        }
                        var v24 string = ("foo bar" + (("foo bar" + ("\n\t" + "hello")) + ("\n\t" + "")))
                        _ = v24
                    }
                    var v25 int64 = 694
                    _ = fi18
                    _ = v25
                }
            }
            v26 := GenS1{F0: -1, F1: -1, F2: ("foo bar" + ("foo bar" + (("" + (("\n\t" + "\n\t") + ("foo bar" + "\n\t"))) + (("" + "foo bar") + "a"))))}
            var v27 uint64 = 293
            _ = v26
            _ = v27
        }
    default:
        fmt.Println(v1)
    }
    fmt.Println(v2)
    switch (("" + ("foo bar" + "hello")) + ("" + "\n\t")) {
    case "\n\t":
        v3 = v3
    }
    if false {
        for fi28 := 0; fi28 < 7; fi28++ {
            v2 = (v2 & (v2 * v2))
            _ = fi28
        }
        fmt.Println(v1)
    }
    var v29 int32 = -1
    _ = len(v4)
    for fi30 := 0; fi30 < 2; fi30++ {
        v2 = (v2 % ((uint16(v1)) | 1))
        var v31 int32 = (-(v29))
        var v32 bool = (570 >= (v29 * 1))
        _ = fi30
        _ = v31
        _ = v32
    }
    v33 := GenS2{F0: int8(v1), F1: 1, F2: v3}
    v34 := GenS3{F0: (-(v29)), F1: v2, F2: ((v2 + 65535) == v2)}
    var v35 uint8 = 42
    v36 := GenS3{F0: (-(v29)), F1: (^(v2)), F2: true}
    for fi37 := 0; fi37 < 6; fi37++ {
        var v38 int8 = 50
        _ = fi37
        _ = v38
    }
    if true {
        v35 += v35
        var v39 int16 = 512
        v1 = (v1 | 0)
        _ = v39
    }
    if (int64(v1) == 9223372036854775807) {
        v4 = append(v4, ("\n\t" + "\n\t"))
        var v40 uint16 = v2
        v4 = append(v4, "foo bar")
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
