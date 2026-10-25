package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 int32; F1 string; F2 uint32 }
type GenS3 struct { F0 int64; F1 uint16; F2 string }
func genFn0() uint {
    v6 := []int{42, 9223372036854775807, -956}
    if len(v6) > 0 {
        v6[0] = 2147483647
    }
    var v7 string = "hello"
    _ = v6
    _ = v7
    return 4294967295
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int8 = -1
    v3 := []string{("" + "foo bar"), ("foo bar" + "a"), (("a" + "a") + "")}
    var v4 int16 = 1000
    v5 := make(map[string]string)
    for fi8 := 0; fi8 < 4; fi8++ {
        v4 -= 350
        var v9 int = 231
        if false {
            v10 := GenS1{F0: v2, F1: 144}
            v4 -= v4
            _ = v10
        } else {
            fmt.Println(v9, v9)
            v9 -= (v9 + 1)
            v5["a"] = "foo bar"
        }
        _ = fi8
        _ = v9
    }
    _ = len(v3)
    v11 := GenS1{F0: v2, F1: 1}
    switch -90 {
    case -746:
        v12 := GenS1{F0: v2, F1: 14}
        _ = v12
    case 4294967295:
        var v13 uint = 255
        _ = v13
    }
    v4 = v4
    fmt.Println(v2, v2)
    var v14 uint8 = 177
    switch 860 {
    case -629:
        if true {
            if (18446744073709551615 >= 18446744073709551615) {
                v4 -= v4
            }
            if (true && true) {
                fmt.Println(v14)
            } else {
                var v15 uint32 = 2
                _ = v15
            }
            fmt.Println(v4)
        }
    case 4294967295:
        fmt.Println(v4)
    case 128:
        for fi16 := 0; fi16 < 16; fi16++ {
            if true {
                for fi17 := 0; fi17 < 8; fi17++ {
                    var v18 uint = genFn0()
                    v18 -= v18
                    if true {
                        var v19 int8 = v2
                        v19 += int8(v14)
                        var v20 int8 = 26
                        _ = v19
                        _ = v20
                    }
                    _ = fi17
                    _ = v18
                }
                v2 = 127
            } else {
                fmt.Println(v2)
            }
            for fi21 := 0; fi21 < 14; fi21++ {
                v4 = int16(v2)
                v4 += -32768
                if (((-(fi16)) == (fi16 >> ((968) & 63))) || true) {
                    var v22 int64 = 1
                    var v23 int32 = -170
                    fmt.Println(fi16)
                    if len(v3) > 0 {
                        v3[0] = "hello"
                    }
                    _ = v22
                    _ = v23
                }
                _ = fi21
            }
            _ = fi16
        }
    default:
        v24 := GenS1{F0: v2, F1: 1}
        _ = v24
    }
    v4 = (v4 - -32768)
    v25 := GenS2{F0: 42, F1: "foo bar", F2: 2147483647}
    switch "\n\t" {
    case "\n\t":
        v4 = 0
    case "":
        fmt.Println(v4, v14)
    default:
        fmt.Println(v2, v14, v2)
    }
    v26 := GenS2{F0: 1, F1: "a", F2: 0}
    var v27 bool = false
    for fi28 := 0; fi28 < 19; fi28++ {
        for fi29 := 0; fi29 < 20; fi29++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi29
        }
        v30 := GenS2{F0: 625, F1: ((("\n\t" + ("" + "a")) + (("hello" + "hello") + "foo bar")) + "hello"), F2: 2147483647}
        switch fi28 {
        case 42:
            var v31 uint16 = 42
            _ = v31
        }
        _ = fi28
        _ = v30
    }
    switch 717 {
    case 152:
        for fi32 := 0; fi32 < 9; fi32++ {
            var v33 int = -1
            v34 := v1["foo bar"]
            v4 = (-(v4))
            _ = fi32
            _ = v33
            _ = v34
        }
    case 1:
        v35 := GenS1{F0: v2, F1: (v14 + (v14 - (v14 % (((v14 ^ 255)) | 1))))}
        _ = v35
    }
    if (1.5 >= ((1.5 / -2.25) / float32(-2.25))) {
        v36 := GenS1{F0: (-(v2)), F1: (v14 << ((genFn0()) & 63))}
        _ = v36
    } else {
        if (!(!v27)) {
            v27 = v27
            v2 = v2
        } else {
            v14 = (v14 * v14)
            v37 := GenS3{F0: -1, F1: 42, F2: ""}
            v14 = 42
            _ = v37
        }
        v14 += (^(v14))
    }
    var v38 bool = false
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v38
    _ = v4
    _ = v5
}
