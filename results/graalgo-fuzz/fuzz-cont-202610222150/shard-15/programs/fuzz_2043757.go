package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 uint16 }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 int8; F1 uint32 }
func genFn0(p0 uint) int16 {
    v6 := GenS1{F0: "\n\t", F1: p0, F2: 42}
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 uint = (p0 % ((65535) | 1))
        var v9 int64 = 9223372036854775807
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 int = -819
    _ = p0
    _ = v10
    _ = v6
    return -32768
}

func main() {
    var v1 string = "\n\t"
    var v2 string = v1
    v3 := []uint32{2147483647, 895, 2147483648, 2}
    var v4 uint32 = 544
    v5 := make(map[uint32]int)
    v2 = ((((v2 + ("" + v2)) + v2) + v2) + "hello")
    switch v2 {
    case "hello":
        for fi11 := 0; fi11 < 2; fi11++ {
            switch v2 {
            case "a":
                var v12 string = "foo bar"
                _ = v12
            default:
                var v13 string = ""
                _ = v13
            }
            if true {
                var v14 int = (fi11 << ((18446744073709551615) & 63))
                _ = v14
            } else {
                var v15 int8 = -40
                _ = v15
            }
            if false {
                fmt.Println(v1, v1, v2)
                for fi16 := 0; fi16 < 2; fi16++ {
                    v4 = v4
                    _ = fi16
                }
                var v17 float32 = (0.0 - 0.0)
                _ = v17
            }
            _ = fi11
        }
    case "foo bar":
        v18 := GenS2{F0: v4, F1: -84}
        _ = v18
    }
    for fi19 := 0; fi19 < 20; fi19++ {
        if (1 <= 122) {
            var v20 uint8 = 255
            if true {
                v4 -= v4
                v2 = ((("foo bar" + v1) + "") + v1)
                v21 := GenS2{F0: (v4 >> ((65535) & 63)), F1: 42}
                _ = v21
            }
            switch "" {
            case "hello":
                fmt.Println(v20)
            case "\n\t":
                var v22 int8 = 1
                _ = v22
            case "a":
                var v23 string = "foo bar"
                _ = v23
            default:
                var v24 float32 = (0.1 - (1e10 + 3.14159))
                _ = v24
            }
            _ = v20
        }
        v2 = ((v1 + v2) + (v2 + v1))
        _ = fi19
    }
    if true {
        v25 := GenS3{F0: 20, F1: false}
        v2 = "foo bar"
        v2 = (v1 + v2)
        _ = v25
    } else {
        var v26 int = 42
        _ = v26
    }
    v27 := GenS4{F0: -1, F1: (v4 << ((65535) & 63))}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1 = v2
    var v28 uint64 = 42
    if (!(!true)) {
        v29 := GenS3{F0: 0, F1: (1 >= 32767)}
        var v30 uint32 = uint32(v28)
        _ = v29
        _ = v30
    } else {
        v2 = ""
        var v31 uint8 = 0
        _ = v31
    }
    var v32 string = v1
    v33 := v5[v4]
    if (((v32 + v1) + ("" + ("\n\t" + (("" + "\n\t") + v32)))) == (((v1 + v2) + ("" + ((v2 + "hello") + v2))) + ((((v1 + "\n\t") + ("hello" + (v1 + v32))) + (("" + "") + v1)) + "hello"))) {
        var v34 float64 = 0.1
        v2 = (v32 + v32)
        switch v32 {
        case "a":
            switch (("a" + (("a" + (v2 + "a")) + v2)) + v2) {
            case "\n\t":
                if (false || (127 == 127)) {
                    if true {
                        var v35 float32 = (1e10 / 3.14159)
                        if ((v34 * v34) >= 3.14159) {
                            v36 := v5[v4]
                            var v37 uint = 4294967295
                            _ = v36
                            _ = v37
                        }
                        v38 := GenS2{F0: v4, F1: -1}
                        _ = v35
                        _ = v38
                    } else {
                        var v39 int64 = 9223372036854775807
                        if len(v3) > 0 {
                            v3[0] = 1
                        }
                        _ = v39
                    }
                    v40 := v5[4294967295]
                    _ = v40
                } else {
                    v34 += v34
                }
            case "hello":
                v41 := GenS4{F0: 42, F1: (v4 / ((365) | 1))}
                _ = v41
            case "a":
                v42 := v5[(v4 ^ (v4 & v4))]
                _ = v42
            }
        case "":
            if true {
                v28 += 892
                v32 = v1
            }
        case "\n\t":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        default:
            v43 := GenS1{F0: v2, F1: 65535, F2: 42}
            _ = v43
        }
        _ = v34
    } else {
        v4 = 2147483647
    }
    var v44 float64 = (0.0 - float64((((1.5 - float32(-2.25)) / (1e10 * 0.1)) + 3.14159)))
    var v45 uint64 = (v28 ^ v28)
    if len(v3) > 0 {
        v3[0] = (v4 ^ (^(v4)))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v44
    _ = v45
    _ = v5
}
