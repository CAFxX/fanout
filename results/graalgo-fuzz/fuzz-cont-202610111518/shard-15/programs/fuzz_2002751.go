package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint32 }
func genFn0(p0 int32) int {
    for fi7 := 0; fi7 < 20; fi7++ {
        v8 := []int{}
        fmt.Println(p0, p0)
        _ = fi7
        _ = v8
    }
    if false {
        switch "" {
        case "foo bar":
            fmt.Println(p0)
        }
    } else {
        var v9 int32 = -1
        for fi10 := 0; fi10 < 2; fi10++ {
            var v11 int64 = 42
            _ = fi10
            _ = v11
        }
        _ = v9
    }
    p0 = (p0 ^ p0)
    v12 := GenS1{F0: 32767}
    _ = p0
    _ = v12
    return -1
}
func genFn1(p0 uint32, p1 int16) int {
    p0 -= (p0 - 42)
    p0 = 2147483647
    _ = p0
    _ = p1
    return -9223372036854775808
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]int)
    var v3 int8 = 1
    v4 := make(map[string]string)
    v5 := []uint32{}
    v6 := []int{0, -2147483648}
    v3 -= 127
    for fi13 := 0; fi13 < 11; fi13++ {
        for fi14 := 0; fi14 < 19; fi14++ {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            v15 := GenS2{F0: 2147483647}
            _ = fi14
            _ = v15
        }
        var v16 uint8 = 1
        _ = fi13
        _ = v16
    }
    fmt.Println(v3)
    var v17 int8 = v3
    if (0 != -824) {
        v17 = 1
    }
    if (false && (v17 != v17)) {
        for fi18 := 0; fi18 < 10; fi18++ {
            v17 += 1
            _ = fi18
        }
        var v19 bool = true
        v20 := GenS3{F0: 128, F1: 255}
        _ = v19
        _ = v20
    }
    var v21 uint64 = 1
    switch "\n\t" {
    case "foo bar":
        var v22 uint16 = 205
        _ = v22
    default:
        v3 += -10
    }
    v23 := GenS3{F0: 209, F1: 1}
    fmt.Println(v17, v21, v17)
    v17 -= (^(v3))
    if false {
        for fi24 := 0; fi24 < 15; fi24++ {
            v3 = (v3 * v17)
            _ = fi24
        }
        switch (("foo bar" + "a") + ("" + ("\n\t" + "hello"))) {
        case "":
            v17 = (-(v17))
        default:
            v21 += (v21 - (^(v21)))
        }
    }
    v3 -= 113
    switch "hello" {
    case "foo bar":
        switch (("\n\t" + ("a" + "foo bar")) + ("foo bar" + ("hello" + "a"))) {
        case "hello":
            v25 := GenS1{F0: 1}
            _ = v25
        case "foo bar":
            if false {
                if false {
                    switch 1 {
                    case 42:
                        for fi26 := 0; fi26 < 20; fi26++ {
                            v17 = (^(v17))
                            v21 = 42
                            if false {
                                var v27 int = -316
                                _ = v27
                            } else {
                                v28 := v4["a"]
                                for fi29 := 0; fi29 < 2; fi29++ {
                                    v3 += 99
                                    if (0 <= 274) {
                                        v28 = ((("a" + "hello") + v28) + v28)
                                        var v30 int8 = (v3 >> ((2) & 63))
                                        _ = v30
                                    }
                                    fmt.Println(v21, v28)
                                    _ = fi29
                                }
                                v3 = (v3 | 1)
                                _ = v28
                            }
                            _ = fi26
                        }
                    case 128:
                        switch int(v17) {
                        case 32768:
                            v31 := GenS3{F0: 128, F1: 189}
                            _ = v31
                        case 128:
                            switch "a" {
                            case "":
                                var v32 int = -900
                                _ = v32
                            case "foo bar":
                                var v33 float32 = (0.1 + -2.25)
                                _ = v33
                            }
                        case 42:
                            _ = len(v5)
                        }
                    case -1:
                        for fi34 := 0; fi34 < 13; fi34++ {
                            if ((true || ((365 != 911) && (false && false))) || ((!true) || (-52 > -32768))) {
                                v21 += (v21 ^ 2)
                                var v35 bool = true
                                _ = v35
                            }
                            for fi36 := 0; fi36 < 6; fi36++ {
                                if false {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    var v37 string = "a"
                                    var v38 uint16 = 65535
                                    _ = v37
                                    _ = v38
                                }
                                _ = fi36
                            }
                            fmt.Println(v17, v21, fi34)
                            _ = fi34
                        }
                    }
                    _ = len(v4)
                    v21 = (v21 * 42)
                }
            }
        default:
            delete(v4, "\n\t")
        }
    case "a":
        v39 := GenS4{F0: 1}
        _ = v39
    case "\n\t":
        v40 := GenS5{F0: 2}
        _ = v40
    }
    fmt.Println(v3, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
