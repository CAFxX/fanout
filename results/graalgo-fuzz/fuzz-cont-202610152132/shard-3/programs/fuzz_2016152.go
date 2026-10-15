package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 uint8; F1 string; F2 uint64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64; F1 int16 }
func genFn0(p0 int, p1 int8, p2 uint) int32 {
    v4 := GenS1{F0: 907, F1: 461}
    p2 = 255
    for fi5 := 0; fi5 < 2; fi5++ {
        fmt.Println(p2)
        if (true && (!false)) {
            if true {
                var v6 bool = true
                _ = v6
            } else {
                for fi7 := 0; fi7 < 2; fi7++ {
                    var v8 uint16 = 0
                    _ = fi7
                    _ = v8
                }
                for fi9 := 0; fi9 < 4; fi9++ {
                    var v10 int = (fi5 & 1)
                    p2 = 65535
                    _ = fi9
                    _ = v10
                }
            }
            p2 = p2
        }
        v11 := GenS2{F0: 42, F1: "a", F2: 18446744073709551615}
        _ = fi5
        _ = v11
    }
    v12 := []int{42, p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v4
    return -1
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]bool)
    v3 := []string{"foo bar", "foo bar", (("a" + (((((("a" + ("a" + ("hello" + "hello"))) + "hello") + (("foo bar" + "") + "a")) + (("\n\t" + "") + ("\n\t" + "hello"))) + "") + "")) + (("hello" + "a") + "a")), ("a" + "")}
    switch 32767 {
    case 9223372036854775807:
        v13 := GenS3{F0: 1}
        _ = v13
    case -9223372036854775808:
        v14 := GenS1{F0: 0, F1: 670}
        _ = v14
    case -2147483648:
        switch "" {
        case "foo bar":
            if (0 == 0) {
                var v15 float64 = ((1e10 / 0.1) - (1e10 - 3.14159))
                fmt.Println(v15)
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v15
            } else {
                for fi16 := 0; fi16 < 15; fi16++ {
                    if false {
                        var v17 uint32 = 65535
                        _ = v17
                    }
                    _ = fi16
                }
            }
        case "":
            if len(v3) > 0 {
                v3[0] = (("a" + ("foo bar" + ("" + ""))) + (("\n\t" + (("foo bar" + "a") + "foo bar")) + "a"))
            }
        default:
            var v18 uint8 = 42
            _ = v18
        }
    }
    var v19 uint8 = 127
    switch "hello" {
    case "hello":
        v20 := v2[-727]
        _ = v20
    }
    var v21 uint32 = 4294967295
    v19 -= (v19 << ((4294967295) & 63))
    for fi22 := 0; fi22 < 20; fi22++ {
        v23 := GenS2{F0: v19, F1: "hello", F2: 2}
        _ = len(v1)
        fmt.Println(v19)
        _ = fi22
        _ = v23
    }
    var v24 uint = 0
    v25 := GenS4{F0: 654, F1: 32767}
    v19 = 42
    var v26 float64 = -2.25
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    var v27 float32 = 3.14159
    for fi28 := 0; fi28 < 10; fi28++ {
        delete(v2, fi28)
        var v29 uint = v24
        _ = fi28
        _ = v29
    }
    _ = len(v3)
    _ = len(v1)
    v30 := GenS1{F0: 42, F1: 181}
    delete(v1, "foo bar")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v30
}
