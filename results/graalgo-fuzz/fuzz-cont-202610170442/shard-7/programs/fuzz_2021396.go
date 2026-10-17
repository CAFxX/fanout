package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 int16 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 uint32; F1 uint64; F2 int32 }
type GenS4 struct { F0 bool; F1 uint8; F2 uint }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint64; F1 bool; F2 uint32 }
func genFn0(p0 int8, p1 int) int32 {
    v5 := []int{(p1 - p1)}
    var v6 uint8 = 0
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return -1
}
func genFn1(p0 uint32) bool {
    v7 := []int{}
    var v8 int = 4294967295
    if (1.5 <= -2.25) {
        v9 := GenS1{F0: (0 >= -2147483648), F1: 919, F2: -1}
        v10 := GenS2{F0: 2147483647, F1: 255}
        _ = v10
        _ = v9
    } else {
        var v11 int16 = -1
        var v12 uint16 = 918
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        switch ((((("foo bar" + "\n\t") + "hello") + ("" + "foo bar")) + "\n\t") + ("\n\t" + ("a" + "hello"))) {
        case "a":
            for fi14 := 0; fi14 < 10; fi14++ {
                v8 = -2147483648
                var v15 int = -9223372036854775808
                _ = fi14
                _ = v15
            }
        case "hello":
            switch (v8 * (v8 - v8)) {
            case 290:
                var v16 int = fi13
                _ = v16
            case 4294967295:
                fmt.Println(p0)
            case 32768:
                _ = len(v7)
            default:
                var v17 string = ("hello" + ("hello" + "a"))
                _ = v17
            }
        }
        var v18 uint8 = 128
        for fi19 := 0; fi19 < 10; fi19++ {
            if true {
                if (!false) {
                    v8 = (fi19 ^ (fi19 >> ((0) & 63)))
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    var v20 uint = 718
                    _ = v20
                } else {
                    if false {
                        v8 -= (v8 % (((fi19 * (fi19 << ((548) & 63)))) | 1))
                        if len(v7) > 0 {
                            v7[0] = 4294967295
                        }
                        v8 += (fi13 / ((fi19) | 1))
                    } else {
                        v21 := GenS1{F0: (!true), F1: -1, F2: 824}
                        _ = v21
                    }
                }
                if true {
                    v7 = append(v7, (fi13 + (-(fi19))))
                } else {
                    v22 := GenS1{F0: (false && true), F1: 779, F2: -32768}
                    _ = v22
                }
            }
            fmt.Println(v18, v18)
            var v23 string = "a"
            _ = fi19
            _ = v23
        }
        _ = fi13
        _ = v18
    }
    var v24 int32 = genFn0(-118, v8)
    _ = p0
    _ = v24
    _ = v7
    _ = v8
    return (255 <= 357)
}

func main() {
    v1 := make(map[string]string)
    var v2 uint = 18446744073709551615
    v3 := make(map[string]int)
    var v4 float32 = (((1e10 - -2.25) * float32(0.1)) - -2.25)
    if true {
        for fi25 := 0; fi25 < 9; fi25++ {
            v2 += 2
            fmt.Println(v2)
            if (127 >= 127) {
                var v26 int8 = -104
                _ = v26
            } else {
                var v27 string = (("" + (("a" + "") + "\n\t")) + ("" + "foo bar"))
                _ = v27
            }
            _ = fi25
        }
        v3["a"] = 32767
        switch ((("a" + ("a" + "hello")) + "") + "") {
        case "foo bar":
            v28 := GenS3{F0: 4294967295, F1: 105, F2: 2147483647}
            _ = v28
        case "hello":
            fmt.Println(v2)
        default:
            var v29 string = "\n\t"
            _ = v29
        }
    }
    if false {
        var v30 float64 = 1.5
        _ = v30
    }
    v2 -= v2
    v31 := GenS2{F0: 0, F1: 128}
    var v32 int32 = 1
    if true {
        fmt.Println(v2, v2, v32)
    } else {
        v4 = v4
        for fi33 := 0; fi33 < 14; fi33++ {
            v34 := []int{fi33, (fi33 % ((fi33) | 1)), fi33}
            v35 := GenS4{F0: false, F1: 0, F2: v2}
            _ = fi33
            _ = v34
            _ = v35
        }
    }
    for fi36 := 0; fi36 < 7; fi36++ {
        if false {
            for fi37 := 0; fi37 < 6; fi37++ {
                v32 -= v32
                _ = fi37
            }
        } else {
            v38 := GenS2{F0: 600, F1: 2}
            _ = v38
        }
        for fi39 := 0; fi39 < 13; fi39++ {
            v32 = v32
            switch (fi36 | 42) {
            case 4294967295:
                v40 := GenS5{F0: ((("" + "foo bar") + "a") + "foo bar")}
                _ = v40
            case 776:
                fmt.Println(fi36, v2)
            case 128:
                v4 += float32(1.5)
            }
            v2 = 533
            _ = fi39
        }
        _ = fi36
    }
    switch 32767 {
    case 1:
        v1["\n\t"] = ""
    case 42:
        fmt.Println(v2)
    default:
        if true {
            delete(v1, "foo bar")
        }
    }
    if true {
        fmt.Println(v32, v32, v2)
    } else {
        for fi41 := 0; fi41 < 19; fi41++ {
            v3["foo bar"] = -811
            for fi42 := 0; fi42 < 9; fi42++ {
                v43 := v1[""]
                v32 = (v32 ^ v32)
                _ = fi42
                _ = v43
            }
            var v44 uint = (v2 & v2)
            _ = fi41
            _ = v44
        }
    }
    switch 0 {
    case 637:
        if (-946 < int32(v2)) {
            fmt.Println(v2)
        } else {
            for fi45 := 0; fi45 < 14; fi45++ {
                v2 = (v2 ^ v2)
                fmt.Println(v4, v32)
                _ = fi45
            }
        }
    case -2147483648:
        v46 := GenS1{F0: (-1 < 127), F1: -1, F2: -32768}
        _ = v46
    case 4294967295:
        switch 128 {
        case 698:
            v32 += 1
        }
    default:
        v2 += v2
    }
    var v47 int = -2147483648
    fmt.Println(v47, v32, v47)
    v48 := GenS6{F0: 18446744073709551615, F1: true, F2: 42}
    v49 := []int{}
    v47 = v47
    var v50 uint16 = 255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v49:", len(v49), v49)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v50
}
