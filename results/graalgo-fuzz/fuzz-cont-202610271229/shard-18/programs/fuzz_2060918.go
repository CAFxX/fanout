package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 string; F1 int8 }
func genFn0(p0 uint32, p1 uint16, p2 uint) int32 {
    v6 := GenS1{F0: ("a" + ("" + ("foo bar" + "a"))), F1: -2147483648}
    switch -1 {
    case 32768:
        switch -993 {
        case -9223372036854775808:
            v7 := GenS2{F0: 42}
            _ = v7
        case 128:
            v8 := []int{-9223372036854775808, -1}
            _ = v8
        }
    case 42:
        var v9 string = "hello"
        _ = v9
    case -313:
        v10 := []int{}
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 1
}

func main() {
    v1 := []string{"a"}
    v2 := make(map[string]int)
    var v3 int8 = -38
    var v4 uint64 = 4294967295
    var v5 uint64 = 489
    fmt.Println(v3)
    if (!false) {
        delete(v2, "")
        if false {
            var v11 string = ((("" + "\n\t") + "") + "hello")
            var v12 int = 42
            if true {
                fmt.Println(v3)
                fmt.Println(v12)
                v11 = v11
            }
            _ = v11
            _ = v12
        } else {
            v5 = v5
        }
        v5 -= v4
    }
    switch "hello" {
    case "hello":
        var v13 uint = 527
        _ = v13
    }
    for fi14 := 0; fi14 < 17; fi14++ {
        var v15 float32 = float32(3.14159)
        var v16 uint64 = 147
        _ = fi14
        _ = v15
        _ = v16
    }
    if (!(!((!false) || (1.5 >= 0.1)))) {
        for fi17 := 0; fi17 < 14; fi17++ {
            fmt.Println(v4)
            _ = fi17
        }
        switch "foo bar" {
        case "foo bar":
            v18 := GenS3{F0: 255, F1: 0}
            _ = v18
        case "hello":
            for fi19 := 0; fi19 < 18; fi19++ {
                if true {
                    v5 = 99
                    for fi20 := 0; fi20 < 16; fi20++ {
                        v3 = int8(fi19)
                        fmt.Println(v4)
                        _ = fi20
                    }
                    var v21 float32 = -2.25
                    _ = v21
                } else {
                    var v22 bool = true
                    _ = v22
                }
                _ = fi19
            }
        case "\n\t":
            var v23 string = "a"
            _ = v23
        }
    }
    v3 = v3
    if (1 <= 660) {
        v5 += v4
        delete(v2, ("\n\t" + "foo bar"))
    }
    switch (("" + ("a" + "a")) + "a") {
    case "":
        v24 := GenS2{F0: v3}
        _ = v24
    case "foo bar":
        var v25 int8 = 127
        _ = v25
    case "a":
        var v26 uint32 = 531
        _ = v26
    default:
        _ = len(v2)
    }
    for fi27 := 0; fi27 < 3; fi27++ {
        if (!((fi27 % (((-(fi27))) | 1)) <= 32767)) {
            _ = len(v2)
            v5 += (^(v5))
            if (!false) {
                fmt.Println(v4, fi27, v3)
                var v28 int16 = 1000
                _ = v28
            } else {
                if ((!(253 > 317)) && false) {
                    for fi29 := 0; fi29 < 19; fi29++ {
                        v30 := GenS4{F0: ("a" + "hello"), F1: v3}
                        _ = fi29
                        _ = v30
                    }
                    fmt.Println(fi27)
                    v4 = v5
                } else {
                    var v31 int64 = -956
                    _ = v31
                }
                for fi32 := 0; fi32 < 3; fi32++ {
                    var v33 string = "a"
                    v34 := GenS3{F0: 18446744073709551615, F1: 0}
                    _ = fi32
                    _ = v33
                    _ = v34
                }
                if (-545 > 42) {
                    v35 := GenS1{F0: "foo bar", F1: fi27}
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    for fi36 := 0; fi36 < 18; fi36++ {
                        fmt.Println(v4, fi27, fi27)
                        _ = fi36
                    }
                    _ = v35
                }
            }
        }
        switch fi27 {
        case -2147483648:
            var v37 int32 = 1
            _ = v37
        }
        v1 = append(v1, "")
        _ = fi27
    }
    switch "\n\t" {
    case "":
        if (genFn0(4294967295, 724, 18446744073709551615) <= 0) {
            fmt.Println(v4, v5)
            v1 = append(v1, "hello")
            var v38 float32 = (((-2.25 * (0.0 / 1.0)) - (0.0 + float32(((1e10 + 0.1) * 1.5)))) + 3.14159)
            _ = v38
        }
    case "hello":
        fmt.Println(v4, v4, v4)
    default:
        var v39 uint16 = 584
        _ = v39
    }
    v40 := v2["foo bar"]
    fmt.Println(v40, v5, v3)
    var v41 uint64 = v4
    v40 = (v40 * v40)
    v3 += v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
