package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int; F1 uint8; F2 int64 }
func genFn0(p0 int64) uint8 {
    var v6 int64 = p0
    v7 := GenS1{F0: 42}
    v8 := GenS1{F0: 516}
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 uint) int8 {
    var v9 float64 = float64(0.0)
    v10 := GenS1{F0: 0}
    _ = p0
    _ = v10
    _ = v9
    return 42
}

func main() {
    v1 := []string{("" + "\n\t"), "hello", "a"}
    var v2 uint16 = 65535
    var v3 float32 = (float32(1e10) * 0.0)
    var v4 int64 = 42
    var v5 float32 = float32((0.0 - 1e10))
    if true {
        if false {
            _ = len(v1)
            v5 = float32(0.1)
        }
    }
    switch 2147483647 {
    case -323:
        v1 = append(v1, "")
    case 32768:
        for fi11 := 0; fi11 < 2; fi11++ {
            fmt.Println(v3, v3, v5)
            if false {
                v12 := GenS1{F0: 4294967295}
                _ = v12
            }
            _ = fi11
        }
    case 32767:
        if false {
            var v13 int = 127
            v14 := GenS2{F0: 2, F1: true}
            _ = v13
            _ = v14
        } else {
            v5 = -2.25
            v4 -= v4
        }
    default:
        for fi15 := 0; fi15 < 5; fi15++ {
            for fi16 := 0; fi16 < 1; fi16++ {
                fmt.Println(fi15, fi16)
                _ = fi16
            }
            _ = fi15
        }
    }
    if false {
        v17 := GenS2{F0: 2, F1: false}
        v3 += float32((0.1 + -2.25))
        v5 -= 0.1
        _ = v17
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        for fi18 := 0; fi18 < 18; fi18++ {
            v1 = append(v1, ("a" + "hello"))
            for fi19 := 0; fi19 < 16; fi19++ {
                v1 = append(v1, "")
                _ = fi19
            }
            _ = fi18
        }
    }
    fmt.Println(v4, v2)
    var v20 uint8 = 0
    for fi21 := 0; fi21 < 17; fi21++ {
        switch ("foo bar" + "hello") {
        case "a":
            if len(v1) > 0 {
                v1[0] = "a"
            }
        case "\n\t":
            v22 := GenS2{F0: 42, F1: ((v4 * int64(v20)) >= int64(v20))}
            _ = v22
        case "hello":
            if (-1 < -938) {
                var v23 uint8 = (v20 - (v20 >> ((18446744073709551615) & 63)))
                var v24 int32 = -2147483648
                _ = v23
                _ = v24
            } else {
                if true {
                    for fi25 := 0; fi25 < 17; fi25++ {
                        v26 := GenS3{F0: genFn0(v4)}
                        _ = fi25
                        _ = v26
                    }
                } else {
                    var v27 int = fi21
                    if true {
                        v28 := GenS2{F0: 189, F1: true}
                        _ = v28
                    } else {
                        v20 = 74
                        v5 += v5
                    }
                    _ = v27
                }
                var v29 uint = 18446744073709551615
                _ = v29
            }
        default:
            v5 = v3
        }
        var v30 int8 = 42
        _ = fi21
        _ = v30
    }
    if (0 > 0) {
        switch "" {
        case "a":
            switch "\n\t" {
            case "":
                var v31 int8 = -1
                _ = v31
            case "\n\t":
                v32 := GenS4{F0: 980, F1: v20, F2: (^(v4))}
                _ = v32
            case "hello":
                v33 := GenS2{F0: 1, F1: ((false && false) && false)}
                _ = v33
            }
        case "":
            v3 += 1e10
        }
        v2 = (v2 | (v2 << ((4294967295) & 63)))
    }
    v2 -= v2
    v4 = (v4 - (v4 << ((997) & 63)))
    v34 := GenS1{F0: 2147483648}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
