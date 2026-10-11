package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 int8 }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 uint; F1 uint; F2 int8 }
type GenS4 struct { F0 string; F1 uint16; F2 bool }
type GenS5 struct { F0 uint; F1 uint }
type GenS6 struct { F0 int64; F1 string }
func genFn0(p0 int) uint64 {
    var v6 float32 = (1e10 / 0.1)
    v7 := []int{p0, 1, 0, (^(p0))}
    v8 := GenS1{F0: 18446744073709551615, F1: 255, F2: 100}
    p0 += (p0 ^ 32768)
    for fi9 := 0; fi9 < 6; fi9++ {
        if true {
            v10 := GenS1{F0: 2, F1: 42, F2: 63}
            _ = v10
        } else {
            var v11 int = 4294967295
            var v12 int64 = -132
            var v13 int = (p0 ^ 1)
            _ = v11
            _ = v12
            _ = v13
        }
        fmt.Println(v6)
        v6 = 0.0
        _ = fi9
    }
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 545
}
func genFn1(p0 uint32, p1 int8, p2 string) int64 {
    var v14 string = ""
    if (-915 != -2147483648) {
        fmt.Println(p0, p0, v14)
        if false {
            var v15 float32 = 1e10
            _ = v15
        } else {
            if (true && true) {
                switch v14 {
                case "\n\t":
                    var v16 uint64 = 1
                    _ = v16
                }
                fmt.Println(v14)
            }
        }
    } else {
        v14 = (v14 + p2)
        v17 := []int{32767, 520}
        _ = v17
    }
    if (200 == 127) {
        p0 = (p0 % (((p0 / ((p0) | 1))) | 1))
        for fi18 := 0; fi18 < 15; fi18++ {
            v19 := []int{34}
            if len(v19) > 0 {
                v19[0] = 815
            }
            for fi20 := 0; fi20 < 10; fi20++ {
                v21 := GenS2{F0: false, F1: -1000}
                _ = fi20
                _ = v21
            }
            _ = fi18
            _ = v19
        }
    } else {
        switch "" {
        case "hello":
            p1 = p1
        default:
            v22 := GenS2{F0: false, F1: -9223372036854775808}
            _ = v22
        }
        v23 := GenS2{F0: (3.14159 < (((1e10 - 0.1) - 3.14159) / 0.1)), F1: -544}
        _ = v23
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 1
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]int)
    v3 := make(map[uint32]int)
    var v4 float64 = 1e10
    v5 := []string{("" + "a"), (("foo bar" + "hello") + ("a" + ""))}
    v24 := GenS3{F0: 42, F1: 65535, F2: 42}
    switch ((("\n\t" + (("a" + "foo bar") + "\n\t")) + "hello") + "") {
    case "":
        for fi25 := 0; fi25 < 19; fi25++ {
            if false {
                if true {
                    v4 = v4
                } else {
                    v26 := GenS4{F0: "a", F1: 877, F2: false}
                    var v27 string = "hello"
                    _ = v26
                    _ = v27
                }
            } else {
                v5 = append(v5, ("" + (("" + "\n\t") + "hello")))
                for fi28 := 0; fi28 < 2; fi28++ {
                    v29 := GenS5{F0: 445, F1: 307}
                    var v30 uint = 53
                    _ = fi28
                    _ = v29
                    _ = v30
                }
            }
            _ = fi25
        }
    case "foo bar":
        switch -432 {
        case 4294967295:
            v5 = append(v5, "foo bar")
        case 127:
            v31 := GenS1{F0: 721, F1: 86, F2: -128}
            _ = v31
        }
    }
    v5 = append(v5, ((("foo bar" + "\n\t") + ("\n\t" + "a")) + "foo bar"))
    fmt.Println(v4)
    v4 = (((float64((-2.25 + 1.5)) + v4) - v4) * 1.5)
    var v32 int64 = 0
    for fi33 := 0; fi33 < 12; fi33++ {
        v32 = v32
        var v34 int32 = 2147483647
        var v35 string = "hello"
        _ = fi33
        _ = v34
        _ = v35
    }
    var v36 uint8 = 23
    var v37 uint32 = 0
    for fi38 := 0; fi38 < 7; fi38++ {
        _ = len(v5)
        for fi39 := 0; fi39 < 12; fi39++ {
            if false {
                if (0 > -60) {
                    var v40 int8 = -1
                    _ = v40
                }
            } else {
                var v41 uint64 = 1
                _ = v41
            }
            v42 := GenS6{F0: v32, F1: "\n\t"}
            _ = fi39
            _ = v42
        }
        _ = fi38
    }
    var v43 int32 = -2147483648
    if len(v5) > 0 {
        v5[0] = (("a" + ("foo bar" + ((("" + ("hello" + "foo bar")) + ("foo bar" + "\n\t")) + ("hello" + (("a" + "hello") + "a"))))) + "hello")
    }
    fmt.Println(v32, v32, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v32
    _ = v36
    _ = v37
    _ = v4
    _ = v43
    _ = v5
}
