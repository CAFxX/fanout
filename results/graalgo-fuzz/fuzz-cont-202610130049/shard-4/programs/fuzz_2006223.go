package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 uint16; F1 uint8; F2 int32 }
type GenS3 struct { F0 int; F1 int64; F2 int64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int; F1 uint32 }
func genFn0() string {
    v5 := GenS1{F0: -965, F1: 831}
    switch "" {
    case "hello":
        var v6 uint8 = 128
        _ = v6
    case "a":
        var v7 uint16 = 255
        _ = v7
    default:
        v8 := []int{32767, 2147483647}
        _ = v8
    }
    var v9 uint8 = 2
    _ = v5
    _ = v9
    return ("a" + ("foo bar" + ((("hello" + "a") + "") + "hello")))
}
func genFn1(p0 uint8, p1 int64) string {
    var v10 int = -885
    if true {
        switch v10 {
        case -212:
            v11 := []int{}
            _ = v11
        default:
            for fi12 := 0; fi12 < 18; fi12++ {
                fmt.Println(p0, v10, p1)
                _ = fi12
            }
        }
        if (!(-32768 == -1)) {
            v10 = v10
            p0 += 128
            fmt.Println(p1)
        }
        var v13 float32 = 1e10
        _ = v13
    } else {
        if (-12 != 0) {
            p0 += (p0 % ((p0) | 1))
            fmt.Println(p1, p0, v10)
            v14 := []int{int(p1), 324, v10, v10}
            _ = v14
        }
    }
    fmt.Println(v10, p1, p0)
    if (!(!true)) {
        v15 := []int{9223372036854775807, v10, (v10 & v10), v10}
        for fi16 := 0; fi16 < 15; fi16++ {
            if false {
                switch "hello" {
                case "hello":
                    var v17 string = ("\n\t" + "a")
                    _ = v17
                case "\n\t":
                    v10 = 32767
                case "":
                    _ = len(v15)
                default:
                    p1 = -153
                }
                if ((127 <= 127) || (int32(fi16) == 836)) {
                    if false {
                        for fi18 := 0; fi18 < 8; fi18++ {
                            var v19 uint = 0
                            _ = fi18
                            _ = v19
                        }
                        var v20 uint8 = (p0 & p0)
                        var v21 uint64 = 4294967295
                        _ = v20
                        _ = v21
                    }
                    if len(v15) >= 2 {
                        v15 = v15[1:len(v15)]
                    }
                }
            }
            _ = fi16
        }
        _ = v15
    } else {
        if true {
            switch v10 {
            case -933:
                switch ("hello" + "a") {
                case "hello":
                    fmt.Println(p1, p1)
                default:
                    v22 := GenS2{F0: 187, F1: p0, F2: 74}
                    _ = v22
                }
            case 486:
                p1 = 1
            }
            switch "hello" {
            case "\n\t":
                var v23 uint8 = 36
                _ = v23
            default:
                fmt.Println(p0, p0, p0)
            }
            for fi24 := 0; fi24 < 5; fi24++ {
                p0 = (p0 / ((234) | 1))
                _ = fi24
            }
        } else {
            fmt.Println(v10, p1, p0)
            if (965 <= -2147483648) {
                fmt.Println(p1, v10)
                switch (v10 * v10) {
                case 0:
                    v25 := GenS1{F0: v10, F1: 2147483647}
                    _ = v25
                case 4294967295:
                    if (false && true) {
                        var v26 int16 = 1000
                        var v27 int = (v10 << ((4294967295) & 63))
                        _ = v26
                        _ = v27
                    } else {
                        switch 42 {
                        case 128:
                            v28 := []int{}
                            _ = v28
                        case 2147483647:
                            v29 := []int{v10, (v10 | (v10 * (v10 / ((127) | 1)))), (v10 - (v10 << ((42) & 63))), (v10 | v10)}
                            _ = v29
                        case 0:
                            fmt.Println(p1, p0)
                        default:
                            switch 1 {
                            case -1:
                                p0 = p0
                            case 32768:
                                switch v10 {
                                case 9223372036854775807:
                                    var v30 string = ("\n\t" + "")
                                    _ = v30
                                case -140:
                                    var v31 uint = 18446744073709551615
                                    _ = v31
                                }
                            }
                        }
                        v10 += 9223372036854775807
                    }
                case 32768:
                    if (((0.1 + (0.1 - float32(1.5))) - float32(-2.25)) <= 1.5) {
                        v32 := GenS3{F0: v10, F1: (p1 >> ((65535) & 63)), F2: -1}
                        p0 += p0
                        _ = v32
                    }
                }
            }
            var v33 string = ""
            _ = v33
        }
    }
    _ = p0
    _ = p1
    _ = v10
    return ""
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int = -659
    var v3 uint32 = 1
    var v4 int8 = 127
    v1[v3] = -85
    v34 := GenS1{F0: 967, F1: (v3 % ((v3) | 1))}
    var v35 uint32 = (v3 + (^(v3)))
    for fi36 := 0; fi36 < 5; fi36++ {
        v3 += (^(v35))
        v35 += v3
        _ = fi36
    }
    fmt.Println(v2, v35)
    var v37 uint32 = v3
    v38 := GenS2{F0: 1, F1: 1, F2: 42}
    v39 := GenS4{F0: (v37 << ((2) & 63))}
    for fi40 := 0; fi40 < 2; fi40++ {
        v35 = v37
        for fi41 := 0; fi41 < 7; fi41++ {
            var v42 int64 = 42
            _ = fi41
            _ = v42
        }
        _ = fi40
    }
    fmt.Println(v2)
    v43 := []int{int(v4), -36}
    v44 := GenS5{F0: 32767, F1: (v35 & 42)}
    var v45 int32 = -225
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v43
    _ = v44
    _ = v45
}
