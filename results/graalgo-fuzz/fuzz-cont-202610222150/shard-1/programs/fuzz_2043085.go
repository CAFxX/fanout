package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 uint16; F1 int64; F2 uint8 }
type GenS4 struct { F0 int64; F1 uint; F2 int32 }
type GenS5 struct { F0 int8 }
func genFn0() int8 {
    var v5 uint = 4294967295
    for fi6 := 0; fi6 < 19; fi6++ {
        v5 += (v5 & 42)
        var v7 uint32 = 2147483648
        _ = fi6
        _ = v7
    }
    fmt.Println(v5)
    var v8 int16 = 32767
    _ = v5
    _ = v8
    return 0
}
func genFn1(p0 bool, p1 uint32, p2 string) uint {
    switch ("a" + p2) {
    case "foo bar":
        fmt.Println(p2, p1, p2)
    case "\n\t":
        fmt.Println(p1, p2)
    case "":
        fmt.Println(p1, p0, p1)
    }
    var v9 int = 4294967295
    for fi10 := 0; fi10 < 4; fi10++ {
        v11 := []int{(v9 ^ 9223372036854775807), (^(v9)), fi10, v9}
        v12 := GenS1{F0: 2}
        fmt.Println(p1, p1)
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 int8 = -1
    p0 = (false && true)
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v9
    return 66
}

func main() {
    v1 := []string{}
    var v2 int32 = 0
    var v3 uint16 = 42
    var v4 bool = (!(0 <= -880))
    if ("hello" >= "a") {
        switch 127 {
        case 4294967295:
            v3 = 753
        case 80:
            if (!v4) {
                v4 = v4
                var v14 bool = v4
                _ = v14
            } else {
                switch "foo bar" {
                case "":
                    switch "" {
                    case "a":
                        switch -268 {
                        case -295:
                            v4 = (-2.25 != 0.1)
                        case 32767:
                            if (v3 == v3) {
                                for fi15 := 0; fi15 < 18; fi15++ {
                                    var v16 int = (fi15 % ((0) | 1))
                                    _ = fi15
                                    _ = v16
                                }
                            } else {
                                fmt.Println(v4, v2, v3)
                                var v17 int = -140
                                v18 := GenS1{F0: 428}
                                _ = v17
                                _ = v18
                            }
                        case -226:
                            v2 = (v2 | -301)
                        }
                    default:
                        v2 -= v2
                    }
                case "a":
                    var v19 string = "hello"
                    _ = v19
                }
                v20 := GenS1{F0: 0}
                var v21 bool = (!true)
                _ = v20
                _ = v21
            }
        }
        var v22 int = 881
        _ = v22
    } else {
        v23 := GenS2{F0: -123, F1: 128}
        _ = v23
    }
    switch 127 {
    case -96:
        for fi24 := 0; fi24 < 4; fi24++ {
            if (v4 || (!(835 <= 0))) {
                v4 = (float64(1.5) < 0.0)
                switch (fi24 << ((455) & 63)) {
                case 2147483647:
                    if (24 == 734) {
                        for fi25 := 0; fi25 < 1; fi25++ {
                            v26 := GenS3{F0: v3, F1: 0, F2: 2}
                            v3 -= (v3 + (v3 % ((v3) | 1)))
                            var v27 float32 = ((3.14159 * float32(0.0)) - 3.14159)
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                    } else {
                        fmt.Println(v4, v2, fi24)
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                case 945:
                    var v28 int16 = 42
                    _ = v28
                case -708:
                    _ = len(v1)
                default:
                    v29 := GenS4{F0: 42, F1: genFn1((true || (true || v4)), 2147483647, "foo bar"), F2: (v2 | int32(fi24))}
                    _ = v29
                }
            }
            v3 = (v3 % ((v3) | 1))
            _ = fi24
        }
    default:
        var v30 uint = 255
        _ = v30
    }
    _ = len(v1)
    fmt.Println(v3)
    var v31 string = "a"
    for fi32 := 0; fi32 < 1; fi32++ {
        if len(v1) > 0 {
            v1[0] = ((((v31 + (v31 + "a")) + v31) + ("foo bar" + "")) + "a")
        }
        _ = fi32
    }
    switch 1 {
    case -368:
        if false {
            for fi33 := 0; fi33 < 18; fi33++ {
                v3 += (v3 & (v3 % ((v3) | 1)))
                _ = fi33
            }
        }
    case 2147483647:
        v34 := GenS4{F0: 0, F1: 42, F2: (v2 & v2)}
        _ = v34
    case 32768:
        for fi35 := 0; fi35 < 16; fi35++ {
            if v4 {
                fmt.Println(v31)
            }
            if v4 {
                v3 += v3
            } else {
                fmt.Println(fi35, v4)
                fmt.Println(v31)
                var v36 uint64 = 42
                _ = v36
            }
            _ = fi35
        }
    default:
        var v37 uint = 1
        _ = v37
    }
    _ = len(v1)
    v38 := GenS5{F0: 0}
    var v39 float32 = (((-2.25 + 1.5) / 0.1) / float32((-2.25 + (0.0 + 1.5))))
    if ((!(1 <= 451)) && (428 == 548)) {
        var v40 uint = 473
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v41 uint8 = 2
        _ = v40
        _ = v41
    }
    v1 = append(v1, v31)
    v2 += -1
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v38
    _ = v39
    _ = v4
}
