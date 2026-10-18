package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 int16; F1 uint8 }
type GenS4 struct { F0 bool; F1 int16 }
func genFn0(p0 uint8, p1 string) bool {
    var v5 int = -284
    fmt.Println(p0, p1, v5)
    _ = p0
    _ = p1
    _ = v5
    return (-9223372036854775808 > -2147483648)
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]bool)
    var v3 int64 = 1
    var v4 float64 = 1.5
    var v6 int16 = 32767
    if true {
        switch 0 {
        case -9223372036854775808:
            if (-18 <= -46) {
                v7 := v2["\n\t"]
                if true {
                    for fi8 := 0; fi8 < 18; fi8++ {
                        if v7 {
                            v4 = 0.0
                        }
                        switch fi8 {
                        case 1:
                            v7 = ((!(225 != 65535)) && true)
                        }
                        var v9 int = fi8
                        _ = fi8
                        _ = v9
                    }
                    fmt.Println(v6, v3)
                    fmt.Println(v3, v4, v3)
                }
                _ = v7
            } else {
                var v10 uint32 = 4294967295
                _ = v10
            }
        case 1:
            var v11 uint16 = 1
            _ = v11
        default:
            delete(v2, ("foo bar" + "foo bar"))
        }
        fmt.Println(v3, v3)
    } else {
        if (4294967295 < 0) {
            var v12 int = 2147483647
            var v13 uint = 437
            _ = v12
            _ = v13
        }
        var v14 int = 32767
        _ = v14
    }
    v15 := []int{-755, -1, -2147483648}
    switch "foo bar" {
    case "":
        fmt.Println(v3)
    case "a":
        v3 = v3
    }
    var v16 bool = true
    fmt.Println(v3, v6)
    v16 = (!(((1e10 - 1.5) + 0.1) > 3.14159))
    v17 := GenS1{F0: ("\n\t" + "\n\t")}
    v18 := GenS1{F0: ""}
    for fi19 := 0; fi19 < 5; fi19++ {
        var v20 int32 = -2147483648
        _ = len(v15)
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 1; fi21++ {
        for fi22 := 0; fi22 < 13; fi22++ {
            if (!((true && v16) && v16)) {
                var v23 float64 = 1e10
                if v16 {
                    v24 := GenS2{F0: 127, F1: 1}
                    _ = v24
                } else {
                    if ((fi21 > (fi21 - fi22)) && (255 < 1)) {
                        v4 = v23
                        v25 := GenS1{F0: "a"}
                        v26 := GenS1{F0: "hello"}
                        _ = v25
                        _ = v26
                    }
                }
                fmt.Println(v3, v4)
                _ = v23
            }
            _ = fi22
        }
        v6 = v6
        _ = fi21
    }
    for fi27 := 0; fi27 < 19; fi27++ {
        v28 := GenS2{F0: 0, F1: 42}
        _ = fi27
        _ = v28
    }
    if (0 != 42) {
        var v29 int16 = v6
        v30 := GenS2{F0: 1, F1: 2}
        _ = v29
        _ = v30
    } else {
        if len(v15) > 0 {
            v15[0] = 127
        }
        v2[((((((("" + "a") + "foo bar") + ("a" + "hello")) + ("" + "\n\t")) + "\n\t") + "hello") + ("hello" + (("\n\t" + "a") + ("a" + ("foo bar" + "\n\t")))))] = ((0 < -1) && (!v16))
        switch "a" {
        case "foo bar":
            var v31 string = "hello"
            _ = v31
        case "a":
            v6 += 52
        case "\n\t":
            v3 = v3
        }
    }
    v6 = (v6 - (v6 + -790))
    var v32 float32 = 0.1
    v33 := GenS3{F0: 0, F1: 36}
    v34 := v1[2147483648]
    v35 := GenS4{F0: (!(!(v16 && ((true && (4 > 100)) && (9223372036854775807 != v3))))), F1: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v6
}
