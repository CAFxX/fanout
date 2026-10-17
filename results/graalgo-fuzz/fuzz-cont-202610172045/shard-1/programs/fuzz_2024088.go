package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 string }
func genFn0(p0 int8) bool {
    if false {
        var v5 string = "\n\t"
        v6 := []int{367}
        _ = v5
        _ = v6
    } else {
        if (false || (1e10 == float64(3.14159))) {
            v7 := GenS1{F0: 42, F1: 4294967295, F2: "foo bar"}
            p0 = p0
            _ = v7
        } else {
            v8 := GenS1{F0: -32768, F1: 0, F2: "hello"}
            _ = v8
        }
        var v9 int64 = 179
        v9 = v9
        _ = v9
    }
    var v10 uint8 = 255
    switch "\n\t" {
    case "hello":
        if (true && (float32(1.5) != 1.5)) {
            var v11 string = ((("hello" + ("\n\t" + ("a" + "foo bar"))) + "") + "\n\t")
            if false {
                v12 := []int{-492}
                _ = v12
            } else {
                v13 := GenS1{F0: 1, F1: 4294967295, F2: v11}
                var v14 int8 = p0
                v15 := GenS1{F0: 1, F1: 449, F2: v11}
                _ = v13
                _ = v14
                _ = v15
            }
            _ = v11
        } else {
            switch 128 {
            case 9223372036854775807:
                var v16 uint64 = uint64(v10)
                _ = v16
            case 4294967295:
                fmt.Println(v10, v10)
            }
        }
    case "\n\t":
        v17 := []int{}
        _ = v17
    default:
        p0 = 100
    }
    var v18 uint8 = v10
    var v19 int32 = -1
    _ = p0
    _ = v10
    _ = v18
    _ = v19
    return false
}

func main() {
    v1 := []string{("a" + "hello"), ("\n\t" + ("a" + "foo bar")), ""}
    v2 := make(map[uint32]string)
    var v3 string = "\n\t"
    var v4 uint = 599
    if (!false) {
        v20 := GenS1{F0: 1000, F1: 4294967295, F2: v3}
        fmt.Println(v4)
        _ = v20
    } else {
        var v21 uint8 = 127
        v22 := GenS1{F0: 32767, F1: uint32(v4), F2: (("\n\t" + ("a" + (v3 + "\n\t"))) + ("a" + v3))}
        _ = v21
        _ = v22
    }
    if true {
        v23 := GenS1{F0: 32767, F1: 2147483648, F2: v3}
        for fi24 := 0; fi24 < 13; fi24++ {
            fmt.Println(v4)
            var v25 int64 = 1
            for fi26 := 0; fi26 < 17; fi26++ {
                switch 4294967295 {
                case -537:
                    fmt.Println(v25, v3, fi24)
                case 897:
                    v4 = v4
                case 32767:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                default:
                    var v27 string = v3
                    _ = v27
                }
                _ = fi26
            }
            _ = fi24
            _ = v25
        }
        _ = v23
    }
    if true {
        v3 = (v3 + ("foo bar" + ("\n\t" + v3)))
    }
    var v28 string = v3
    var v29 float64 = 0.0
    v30 := v2[1]
    v4 -= v4
    var v31 int32 = 42
    switch 42 {
    case -68:
        if (true || true) {
            v28 = "a"
            v3 = v3
            v32 := GenS1{F0: -32768, F1: 2147483648, F2: (("" + v28) + "a")}
            _ = v32
        } else {
            v2[4294967295] = "foo bar"
            if false {
                if true {
                    v29 = 1.5
                } else {
                    for fi33 := 0; fi33 < 1; fi33++ {
                        fmt.Println(v3)
                        _ = fi33
                    }
                    fmt.Println(v4)
                }
            }
            var v34 int8 = 100
            _ = v34
        }
    case 127:
        var v35 int16 = 42
        _ = v35
    case 42:
        v28 = v3
    default:
        if len(v1) > 0 {
            v1[0] = ((("foo bar" + (v3 + v30)) + "a") + v3)
        }
    }
    v1 = append(v1, v3)
    var v36 float64 = (v29 * v29)
    var v37 int = -9223372036854775808
    fmt.Println(v4, v29)
    var v38 uint8 = 128
    var v39 int8 = -128
    var v40 float32 = (1e10 + 1.5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
}
