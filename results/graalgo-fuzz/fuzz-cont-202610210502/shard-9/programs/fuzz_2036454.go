package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint; F1 uint }
func genFn0(p0 uint16, p1 int64) int16 {
    var v4 string = ""
    p1 = (p1 >> ((65535) & 63))
    fmt.Println(v4)
    var v5 int64 = (p1 - -9223372036854775808)
    p0 += (p0 / ((42) | 1))
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return -352
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]int)
    var v3 int64 = 9223372036854775807
    fmt.Println(v3)
    var v6 float32 = 0.1
    v7 := []int{}
    switch -2147483648 {
    case 128:
        v3 = (v3 + v3)
    }
    fmt.Println(v3, v6)
    switch "hello" {
    case "\n\t":
        if false {
            v1[2147483647] = "a"
            if (!false) {
                switch (("foo bar" + ("" + ("\n\t" + ("foo bar" + "\n\t")))) + "") {
                case "\n\t":
                    var v8 int = 659
                    _ = v8
                case "":
                    fmt.Println(v3, v3)
                }
                var v9 uint8 = 1
                _ = v9
            } else {
                for fi10 := 0; fi10 < 10; fi10++ {
                    delete(v2, (("a" + "") + "foo bar"))
                    switch ("" + ("foo bar" + "")) {
                    case "a":
                        var v11 int64 = v3
                        _ = v11
                    }
                    _ = fi10
                }
            }
        }
    }
    for fi12 := 0; fi12 < 3; fi12++ {
        fmt.Println(fi12)
        switch "\n\t" {
        case "foo bar":
            v13 := GenS1{F0: "", F1: fi12}
            _ = v13
        case "":
            for fi14 := 0; fi14 < 5; fi14++ {
                if (0 >= v3) {
                    v6 -= (v6 / v6)
                    v6 += v6
                }
                v15 := GenS1{F0: "hello", F1: fi12}
                _ = fi14
                _ = v15
            }
        case "a":
            var v16 float32 = 1.5
            _ = v16
        default:
            var v17 bool = (!(!true))
            _ = v17
        }
        v1[2] = ("foo bar" + ("foo bar" + ((("foo bar" + "hello") + "hello") + "")))
        _ = fi12
    }
    _ = len(v7)
    v3 += v3
    v18 := GenS1{F0: "\n\t", F1: 32768}
    fmt.Println(v6)
    var v19 uint32 = 436
    var v20 int32 = 0
    switch "\n\t" {
    case "foo bar":
        var v21 uint16 = 42
        _ = v21
    default:
        var v22 uint = 255
        _ = v22
    }
    v3 = 0
    if (-32768 != -941) {
        if false {
            v23 := GenS1{F0: "foo bar", F1: 42}
            _ = v23
        }
    } else {
        v6 += (v6 * (float32(0.1) * (0.0 * -2.25)))
    }
    v20 = (v20 << ((736) & 63))
    for fi24 := 0; fi24 < 12; fi24++ {
        for fi25 := 0; fi25 < 12; fi25++ {
            if false {
                if true {
                    var v26 uint = 311
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    v7 = append(v7, 9223372036854775807)
                    _ = v26
                }
                v19 += (^(v19))
                switch fi24 {
                case -2147483648:
                    v19 += 42
                    if (127 == 190) {
                        delete(v1, v19)
                        v3 += (^(v3))
                    }
                    fmt.Println(v6, v6)
                case 32768:
                    v3 = v3
                }
            } else {
                fmt.Println(v19, fi25, v6)
            }
            _ = len(v1)
            _ = fi25
        }
        if false {
            v27 := GenS2{F0: 42, F1: 255}
            _ = v27
        }
        var v28 int8 = 42
        _ = fi24
        _ = v28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v6
    _ = v7
}
