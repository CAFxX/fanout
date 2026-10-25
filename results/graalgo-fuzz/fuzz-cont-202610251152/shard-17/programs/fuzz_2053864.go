package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 int64; F1 int; F2 uint64 }
type GenS3 struct { F0 uint16; F1 int32; F2 uint16 }
func genFn0(p0 uint8) bool {
    p0 = (p0 << ((798) & 63))
    for fi5 := 0; fi5 < 6; fi5++ {
        if false {
            var v6 int64 = 829
            _ = v6
        } else {
            p0 = p0
            if (-9223372036854775808 != (fi5 / ((fi5) | 1))) {
                p0 += (p0 + 2)
            } else {
                if false {
                    v7 := GenS1{F0: 127, F1: true}
                    v8 := GenS1{F0: -1, F1: false}
                    fmt.Println(fi5, p0)
                    _ = v7
                    _ = v8
                } else {
                    v9 := GenS1{F0: -9223372036854775808, F1: true}
                    for fi10 := 0; fi10 < 13; fi10++ {
                        var v11 uint8 = p0
                        switch ("foo bar" + ("" + ("\n\t" + ("\n\t" + (("hello" + "") + (("" + ("\n\t" + "hello")) + "a")))))) {
                        case "foo bar":
                            v12 := []int{}
                            _ = v12
                        case "\n\t":
                            p0 = 255
                        default:
                            var v13 int32 = -773
                            _ = v13
                        }
                        v11 += p0
                        _ = fi10
                        _ = v11
                    }
                    _ = v9
                }
                var v14 string = (("foo bar" + ("" + "")) + ("" + "foo bar"))
                v15 := GenS2{F0: 776, F1: (fi5 >> ((0) & 63)), F2: 19}
                _ = v14
                _ = v15
            }
            fmt.Println(p0, fi5)
        }
        _ = fi5
    }
    _ = p0
    return (255 >= 0)
}

func main() {
    v1 := []string{"foo bar", "hello"}
    var v2 int8 = 0
    var v3 uint64 = 0
    v4 := make(map[string]bool)
    v2 += 1
    var v16 uint16 = 65535
    switch "" {
    case "\n\t":
        v17 := GenS1{F0: -9223372036854775808, F1: (v2 > v2)}
        _ = v17
    case "":
        fmt.Println(v2)
    default:
        v3 -= v3
    }
    if false {
        v18 := GenS3{F0: 460, F1: 0, F2: (v16 - 65535)}
        _ = v18
    } else {
        for fi19 := 0; fi19 < 12; fi19++ {
            v20 := GenS3{F0: (v16 * 0), F1: 0, F2: (^(v16))}
            fmt.Println(fi19, v2, v3)
            if len(v1) > 0 {
                v1[0] = "\n\t"
            }
            _ = fi19
            _ = v20
        }
        v2 += (v2 << ((549) & 63))
        var v21 float64 = (3.14159 - float64(3.14159))
        _ = v21
    }
    v16 -= 32767
    v2 -= (v2 % ((v2) | 1))
    for fi22 := 0; fi22 < 5; fi22++ {
        switch "\n\t" {
        case "":
            if ((false || (524 <= int32(v16))) && true) {
                v23 := v4[""]
                _ = v23
            }
        case "foo bar":
            if false {
                for fi24 := 0; fi24 < 6; fi24++ {
                    if ("hello" >= "foo bar") {
                        if true {
                            v25 := GenS3{F0: v16, F1: 42, F2: v16}
                            _ = v25
                        } else {
                            fmt.Println(v2, v3)
                        }
                        var v26 float64 = 1e10
                        _ = v26
                    } else {
                        var v27 string = "foo bar"
                        var v28 int8 = v2
                        _ = v27
                        _ = v28
                    }
                    _ = fi24
                }
            }
        default:
            v29 := v4[(("\n\t" + "\n\t") + "foo bar")]
            _ = v29
        }
        var v30 uint8 = uint8(v16)
        _ = fi22
        _ = v30
    }
    if (!(!(2147483647 >= 9223372036854775807))) {
        switch ("\n\t" + "hello") {
        case "foo bar":
            var v31 int8 = (v2 << ((65535) & 63))
            _ = v31
        }
        for fi32 := 0; fi32 < 3; fi32++ {
            var v33 uint8 = 2
            _ = fi32
            _ = v33
        }
    } else {
        v34 := GenS3{F0: v16, F1: 917, F2: (v16 >> ((2) & 63))}
        _ = v34
    }
    for fi35 := 0; fi35 < 16; fi35++ {
        v36 := GenS1{F0: fi35, F1: true}
        var v37 int8 = v2
        if (82 != (v2 ^ v37)) {
            var v38 int64 = -157
            _ = v38
        } else {
            switch "\n\t" {
            case "hello":
                _ = len(v1)
            case "":
                var v39 uint16 = 1
                _ = v39
            default:
                if (!false) {
                    var v40 int8 = v37
                    _ = v40
                }
            }
        }
        _ = fi35
        _ = v36
        _ = v37
    }
    v1 = append(v1, ("a" + ("hello" + ("hello" + ""))))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
