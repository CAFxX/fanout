package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 uint; F1 int64 }
func genFn0(p0 bool, p1 uint64, p2 uint64) uint {
    v7 := []int{}
    switch ("a" + "foo bar") {
    case "\n\t":
        if len(v7) > 0 {
            v7[0] = 711
        }
    case "":
        var v8 bool = (!p0)
        _ = v8
    default:
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    }
    v7 = append(v7, 4294967295)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 65535
}

func main() {
    var v1 int64 = 1
    v2 := make(map[string]int)
    v3 := make(map[int]uint32)
    var v4 uint = 42
    v5 := []string{"hello"}
    v6 := make(map[string]string)
    switch -9223372036854775808 {
    case 0:
        v4 += (^(v4))
    default:
        v9 := v3[42]
        _ = v9
    }
    fmt.Println(v1, v1)
    v10 := GenS1{F0: ("hello" + "a"), F1: ("a" + ("hello" + "\n\t")), F2: 783}
    v1 -= (^(v1))
    v11 := GenS2{F0: 32767, F1: -406}
    for fi12 := 0; fi12 < 6; fi12++ {
        v5 = append(v5, ("hello" + ""))
        var v13 float64 = 0.0
        var v14 int8 = -113
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 uint = (v4 >> ((65535) & 63))
    var v16 int8 = 127
    for fi17 := 0; fi17 < 14; fi17++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        switch (("a" + "") + "\n\t") {
        case "\n\t":
            for fi18 := 0; fi18 < 7; fi18++ {
                fmt.Println(fi17, v1, v1)
                switch (("\n\t" + (("a" + (("hello" + "a") + "")) + "hello")) + ("hello" + "")) {
                case "":
                    v16 = (^(v16))
                    v19 := GenS1{F0: ("hello" + ("hello" + "hello")), F1: "", F2: genFn0(((fi18 * (fi18 >> ((v4) & 63))) != (fi17 << ((v4) & 63))), 632, 18446744073709551615)}
                    _ = len(v5)
                    _ = v19
                default:
                    var v20 int32 = 0
                    _ = v20
                }
                v4 += v4
                _ = fi18
            }
        case "":
            v16 = v16
        }
        _ = fi17
    }
    if (true && (!(-1 != 42))) {
        v21 := v2[""]
        _ = v21
    } else {
        if true {
            switch ("hello" + "foo bar") {
            case "foo bar":
                v22 := GenS1{F0: "foo bar", F1: ("foo bar" + "foo bar"), F2: (v15 % (((v4 + v15)) | 1))}
                _ = v22
            default:
                for fi23 := 0; fi23 < 12; fi23++ {
                    var v24 uint8 = 42
                    if false {
                        for fi25 := 0; fi25 < 6; fi25++ {
                            v26 := v3[fi23]
                            _ = fi25
                            _ = v26
                        }
                    }
                    v15 += (v15 >> ((genFn0(false, 0, 0)) & 63))
                    _ = fi23
                    _ = v24
                }
            }
            v27 := GenS3{F0: v15, F1: (v1 | (-(v1)))}
            _ = v27
        }
        if (42 > 18446744073709551615) {
            v28 := GenS1{F0: ("hello" + ("hello" + "a")), F1: "", F2: (^(v15))}
            var v29 uint16 = 0
            _ = v28
            _ = v29
        }
    }
    var v30 int = 9223372036854775807
    var v31 string = ("" + ("hello" + ""))
    var v32 int32 = 192
    v31 = (("a" + ((v31 + v31) + v31)) + v31)
    switch "hello" {
    case "\n\t":
        v33 := GenS3{F0: v4, F1: (^(v1))}
        _ = v33
    case "hello":
        if (1000 == 455) {
            switch (v31 + (("foo bar" + v31) + v31)) {
            case "foo bar":
                var v34 int16 = 1000
                _ = v34
            default:
                v35 := GenS2{F0: uint16(v4), F1: v30}
                _ = v35
            }
            delete(v3, (v30 & (v30 - (v30 / ((128) | 1)))))
        }
    }
    if false {
        v16 = (v16 + v16)
        switch v30 {
        case 9223372036854775807:
            if (!false) {
                var v36 uint64 = 18446744073709551615
                _ = v36
            }
        case 195:
            v5 = append(v5, (("\n\t" + ((v31 + "hello") + v31)) + v31))
        case -440:
            for fi37 := 0; fi37 < 17; fi37++ {
                v1 += int64(v30)
                _ = fi37
            }
        default:
            v30 = (v30 - 167)
        }
        var v38 int8 = v16
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
