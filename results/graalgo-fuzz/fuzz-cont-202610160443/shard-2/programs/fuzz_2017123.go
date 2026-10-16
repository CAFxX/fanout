package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 int64; F1 bool; F2 uint8 }
type GenS3 struct { F0 int16; F1 string; F2 string }
type GenS4 struct { F0 uint; F1 bool }
type GenS5 struct { F0 int8; F1 string }
type GenS6 struct { F0 int32; F1 uint64; F2 int64 }
func genFn0(p0 uint32, p1 int16) int16 {
    var v6 bool = true
    var v7 int = -358
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 42
}
func genFn1() uint64 {
    v8 := GenS1{F0: -9223372036854775808, F1: 100}
    v9 := GenS1{F0: 0, F1: 57}
    var v10 uint8 = 42
    _ = v10
    _ = v8
    _ = v9
    return 4294967295
}

func main() {
    var v1 uint16 = 784
    v2 := []uint32{}
    var v3 int64 = 1
    var v4 uint16 = (v1 & v1)
    var v5 uint16 = v4
    switch (("foo bar" + ("foo bar" + ("foo bar" + ("\n\t" + "hello")))) + "hello") {
    case "hello":
        v1 += 0
    case "foo bar":
        if (0.1 <= -2.25) {
            for fi11 := 0; fi11 < 5; fi11++ {
                var v12 float64 = float64((0.1 - float32(0.1)))
                fmt.Println(v3)
                _ = fi11
                _ = v12
            }
            var v13 uint64 = 1
            _ = v13
        }
    case "a":
        if false {
            v14 := GenS2{F0: (v3 % ((v3) | 1)), F1: (true || false), F2: 113}
            fmt.Println(v4)
            _ = v14
        }
    default:
        if len(v2) > 0 {
            v2[0] = 2147483648
        }
    }
    for fi15 := 0; fi15 < 15; fi15++ {
        for fi16 := 0; fi16 < 16; fi16++ {
            var v17 uint64 = 4294967295
            _ = fi16
            _ = v17
        }
        for fi18 := 0; fi18 < 13; fi18++ {
            var v19 uint16 = (v4 % (((v4 + v4)) | 1))
            var v20 uint16 = (^(v5))
            v21 := GenS1{F0: v3, F1: 100}
            _ = fi18
            _ = v19
            _ = v20
            _ = v21
        }
        _ = fi15
    }
    if true {
        for fi22 := 0; fi22 < 18; fi22++ {
            fmt.Println(v5)
            _ = fi22
        }
        var v23 int8 = 1
        _ = v23
    }
    var v24 bool = ((641 >= 9223372036854775807) || true)
    switch ("a" + "\n\t") {
    case "":
        if (v24 || (-9223372036854775808 >= 128)) {
            fmt.Println(v1, v3)
        } else {
            v1 += 42
            for fi25 := 0; fi25 < 20; fi25++ {
                v26 := GenS3{F0: -670, F1: ("foo bar" + ""), F2: ""}
                var v27 uint8 = 2
                for fi28 := 0; fi28 < 18; fi28++ {
                    if v24 {
                        if len(v2) > 0 {
                            v2[0] = 2147483648
                        }
                    } else {
                        var v29 uint8 = v27
                        _ = v29
                    }
                    if (false && (("a" != (("foo bar" + "") + ("foo bar" + "a"))) || (42 >= v3))) {
                        v5 = v1
                    } else {
                        v1 = (^(v1))
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        switch ("hello" + "\n\t") {
                        case "a":
                            switch 1 {
                            case 2147483647:
                                v30 := GenS4{F0: 255, F1: v24}
                                _ = v30
                            }
                        case "foo bar":
                            fmt.Println(fi25, v27, fi28)
                        case "\n\t":
                            v5 += v1
                        default:
                            switch (("hello" + ("foo bar" + "\n\t")) + ("a" + ("foo bar" + ""))) {
                            case "":
                                switch ("\n\t" + (("\n\t" + "foo bar") + ("\n\t" + ("a" + "")))) {
                                case "":
                                    switch (fi28 ^ (^(fi25))) {
                                    case 42:
                                        var v31 int = fi28
                                        if len(v2) > 0 {
                                            v2[0] = 65535
                                        }
                                        _ = len(v2)
                                        _ = v31
                                    default:
                                        if v24 {
                                            v1 += 2
                                            v5 += (^(v1))
                                            v5 = 301
                                        } else {
                                            v32 := GenS1{F0: v3, F1: -128}
                                            v5 -= (v4 - v4)
                                            v5 += (v5 << ((2) & 63))
                                            _ = v32
                                        }
                                        fmt.Println(fi25, v27, fi25)
                                        fmt.Println(v1, v24, v3)
                                    }
                                case "a":
                                    if false {
                                        v33 := GenS5{F0: 42, F1: ""}
                                        var v34 bool = (1 > (v27 ^ v27))
                                        v4 = 509
                                        _ = v33
                                        _ = v34
                                    }
                                }
                            case "a":
                                v2 = append(v2, 65535)
                            }
                        }
                    }
                    _ = fi28
                }
                _ = fi25
                _ = v26
                _ = v27
            }
            v1 -= v1
        }
    default:
        fmt.Println(v1, v24, v3)
    }
    for fi35 := 0; fi35 < 13; fi35++ {
        var v36 uint64 = 18446744073709551615
        if v24 {
            v37 := GenS1{F0: v3, F1: -82}
            var v38 uint16 = (^(v1))
            v39 := GenS6{F0: 477, F1: v36, F2: v3}
            _ = v37
            _ = v38
            _ = v39
        }
        var v40 uint = 631
        _ = fi35
        _ = v36
        _ = v40
    }
    for fi41 := 0; fi41 < 19; fi41++ {
        if len(v2) > 0 {
            v2[0] = 1
        }
        _ = fi41
    }
    var v42 uint8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v42
    _ = v5
}
