package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16; F1 int16; F2 bool }
type GenS3 struct { F0 bool; F1 string }

func main() {
    var v1 uint = 2
    v2 := []uint32{}
    v3 := make(map[uint32]bool)
    v1 = v1
    for fi4 := 0; fi4 < 5; fi4++ {
        if ((false && true) && false) {
            switch (fi4 % ((fi4) | 1)) {
            case 587:
                v1 += (v1 / ((65535) | 1))
            case 421:
                var v5 int16 = 1
                _ = v5
            }
            var v6 int16 = 344
            switch (fi4 | (fi4 << (((^(v1))) & 63))) {
            case 42:
                delete(v3, 45)
            case -9223372036854775808:
                v7 := GenS1{F0: 343}
                _ = v7
            case -67:
                v6 = (v6 * 0)
            default:
                fmt.Println(v1)
            }
            _ = v6
        }
        _ = fi4
    }
    var v8 uint = 247
    for fi9 := 0; fi9 < 18; fi9++ {
        v8 -= (v8 / (((v8 / ((v1) | 1))) | 1))
        fmt.Println(v1)
        _ = fi9
    }
    var v10 int64 = int64(v1)
    v10 = v10
    fmt.Println(v1, v10, v10)
    var v11 uint32 = 0
    v2 = append(v2, (^(v11)))
    v12 := GenS2{F0: 32767, F1: 1, F2: true}
    delete(v3, v11)
    switch ("foo bar" + (("" + "a") + (("a" + ("foo bar" + ("hello" + ""))) + ("\n\t" + "")))) {
    case "foo bar":
        v13 := GenS2{F0: 757, F1: 32767, F2: false}
        _ = v13
    case "a":
        if (false || true) {
            for fi14 := 0; fi14 < 1; fi14++ {
                v15 := GenS3{F0: true, F1: "foo bar"}
                _ = len(v2)
                _ = fi14
                _ = v15
            }
            v2 = append(v2, (v11 << ((v1) & 63)))
            if false {
                v11 -= (v11 | 65535)
                fmt.Println(v8)
                if (!false) {
                    var v16 string = ("" + (("foo bar" + "hello") + ""))
                    _ = v16
                } else {
                    v11 = v11
                }
            } else {
                v17 := GenS3{F0: false, F1: ("" + "")}
                var v18 uint16 = 152
                switch "" {
                case "hello":
                    v19 := GenS1{F0: 540}
                    _ = v19
                case "\n\t":
                    if (2147483647 == 2147483647) {
                        switch 32767 {
                        case -239:
                            for fi20 := 0; fi20 < 19; fi20++ {
                                fmt.Println(v1, v1)
                                _ = fi20
                            }
                        case -832:
                            var v21 float64 = 1.5
                            _ = v21
                        case 1:
                            v8 = 4294967295
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                }
                _ = v17
                _ = v18
            }
        } else {
            if ((false || true) && false) {
                v2 = append(v2, 65535)
                v22 := GenS2{F0: 1, F1: -1, F2: true}
                fmt.Println(v8)
                _ = v22
            }
        }
    case "hello":
        fmt.Println(v1, v8, v1)
    default:
        v1 -= 0
    }
    var v23 int8 = -52
    var v24 uint16 = 728
    if (!false) {
        for fi25 := 0; fi25 < 19; fi25++ {
            var v26 uint32 = v11
            var v27 int16 = 1
            _ = fi25
            _ = v26
            _ = v27
        }
        for fi28 := 0; fi28 < 5; fi28++ {
            var v29 float64 = (0.1 + -2.25)
            _ = fi28
            _ = v29
        }
    } else {
        var v30 int32 = 0
        _ = v30
    }
    v8 -= (v8 / ((v8) | 1))
    for fi31 := 0; fi31 < 18; fi31++ {
        var v32 int8 = (v23 % ((v23) | 1))
        v3[(v11 % ((v11) | 1))] = (true || (!(v11 >= (v11 - (v11 / ((v11) | 1))))))
        _ = fi31
        _ = v32
    }
    v10 += (v10 / ((302) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v8
}
