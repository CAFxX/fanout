package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 bool; F2 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32; F1 bool; F2 bool }

func main() {
    var v1 float64 = float64((3.14159 + 1.5))
    v2 := make(map[uint32]string)
    var v3 int8 = -1
    v4 := []int{42, 9223372036854775807, 127}
    v5 := []int{0}
    v6 := GenS1{F0: 2}
    var v7 uint16 = 42
    if (0 > 42) {
        switch "\n\t" {
        case "a":
            if false {
                for fi8 := 0; fi8 < 19; fi8++ {
                    var v9 uint64 = 18446744073709551615
                    for fi10 := 0; fi10 < 12; fi10++ {
                        var v11 int16 = -46
                        var v12 uint32 = 2147483648
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                    _ = fi8
                    _ = v9
                }
            } else {
                v1 += 3.14159
            }
        }
        v7 += (v7 << ((uint(v7)) & 63))
        switch ("hello" + "foo bar") {
        case "\n\t":
            switch ("foo bar" + "") {
            case "\n\t":
                if false {
                    v1 -= v1
                    v3 += (-(v3))
                    v13 := GenS2{F0: 125, F1: false, F2: "hello"}
                    _ = v13
                }
            }
        case "hello":
            v1 = v1
        }
    }
    if true {
        var v14 int64 = -867
        fmt.Println(v7, v3)
        delete(v2, 183)
        _ = v14
    }
    switch 127 {
    case 2147483647:
        if (v3 != v3) {
            switch -802 {
            case 4294967295:
                for fi15 := 0; fi15 < 10; fi15++ {
                    v3 = 0
                    var v16 float64 = -2.25
                    _ = fi15
                    _ = v16
                }
            case -9223372036854775808:
                switch -930 {
                case 9223372036854775807:
                    for fi17 := 0; fi17 < 12; fi17++ {
                        v3 = v3
                        _ = fi17
                    }
                case 32767:
                    if (1 > 1000) {
                        v18 := GenS1{F0: 243}
                        var v19 int = 562
                        _ = v18
                        _ = v19
                    } else {
                        switch int(v7) {
                        case -442:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        case 301:
                            var v20 string = (("foo bar" + ("hello" + "")) + "")
                            _ = v20
                        case 32768:
                            var v21 int16 = 451
                            _ = v21
                        }
                        v1 = v1
                        var v22 float32 = ((float32(1.5) - ((0.1 / 1.0) * 1.5)) - 3.14159)
                        _ = v22
                    }
                default:
                    var v23 int = -917
                    _ = v23
                }
            }
        } else {
            var v24 int16 = 1
            for fi25 := 0; fi25 < 13; fi25++ {
                v24 = (v24 / (((v24 + v24)) | 1))
                _ = fi25
            }
            _ = v24
        }
    case 0:
        var v26 int16 = 42
        _ = v26
    default:
        for fi27 := 0; fi27 < 19; fi27++ {
            for fi28 := 0; fi28 < 16; fi28++ {
                if false {
                    v1 = v1
                    v1 += (((v1 / ((v1 - (v1 - (v1 - (((v1 + 0.0) * 1e10) - ((-2.25 * (v1 - (v1 * 0.1))) + v1))))) - float64((0.0 - 3.14159)))) - float64(0.0)) * (float64(((float32(v1) / 1.0) - 0.1)) - v1))
                    fmt.Println(v3)
                } else {
                    _ = len(v4)
                    fmt.Println(v3, v7, v7)
                    v3 -= (v3 / (((v3 | (-(v3)))) | 1))
                }
                var v29 string = "\n\t"
                fmt.Println(v1, fi28)
                _ = fi28
                _ = v29
            }
            v1 -= v1
            _ = fi27
        }
    }
    if false {
        v3 = (v3 & v3)
        _ = len(v4)
        v1 = v1
    }
    var v30 uint16 = v7
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if false {
        if true {
            if false {
                switch (("foo bar" + ("\n\t" + ("foo bar" + (("\n\t" + ("\n\t" + ("a" + ("a" + ("foo bar" + ""))))) + "\n\t")))) + "") {
                case "foo bar":
                    if ((!(!(-532 >= 1))) || (128 < 9223372036854775807)) {
                        v30 -= v7
                        for fi31 := 0; fi31 < 6; fi31++ {
                            switch "hello" {
                            case "\n\t":
                                var v32 int8 = -119
                                _ = v32
                            }
                            _ = fi31
                        }
                    }
                case "\n\t":
                    v7 = (v7 % ((v7) | 1))
                default:
                    v7 += v7
                }
                var v33 uint64 = 2
                fmt.Println(v3, v33, v33)
                _ = v33
            }
            var v34 float32 = 0.1
            _ = v34
        } else {
            switch "foo bar" {
            case "\n\t":
                if (-2147483648 == 2147483647) {
                    var v35 uint64 = 0
                    _ = v35
                } else {
                    v30 += (^(v7))
                    v7 += v30
                    v36 := GenS3{F0: v3}
                    _ = v36
                }
            case "foo bar":
                if true {
                    v37 := v2[17]
                    _ = v37
                }
            case "hello":
                var v38 uint = 4294967295
                _ = v38
            }
            for fi39 := 0; fi39 < 2; fi39++ {
                v40 := GenS1{F0: 128}
                if true {
                    if (-113 == -32768) {
                        if true {
                            v30 = 32767
                        }
                    }
                    _ = len(v2)
                    for fi41 := 0; fi41 < 11; fi41++ {
                        if true {
                            var v42 uint32 = 1
                            v43 := GenS2{F0: 0, F1: (255 > (v7 & (^(v7)))), F2: ("hello" + "a")}
                            v30 -= (^(v7))
                            if true {
                                v30 = (v7 + (v7 | (^(v30))))
                                v42 = v42
                            }
                            _ = v42
                            _ = v43
                        }
                        _ = fi41
                    }
                }
                v44 := GenS4{F0: 0, F1: true, F2: false}
                _ = fi39
                _ = v40
                _ = v44
            }
            v3 = (v3 + 42)
        }
        var v45 uint64 = 1
        _ = v45
    }
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
