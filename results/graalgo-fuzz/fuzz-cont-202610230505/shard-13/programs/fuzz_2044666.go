package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
func genFn0(p0 int16) int {
    switch ("\n\t" + ("a" + "\n\t")) {
    case "\n\t":
        if (0 > 1) {
            v6 := GenS1{F0: ("hello" + ((("hello" + "") + ("hello" + "\n\t")) + ("a" + ("foo bar" + (("hello" + "hello") + "a"))))), F1: p0}
            _ = v6
        }
    case "foo bar":
        if false {
            v7 := []int{9223372036854775807}
            v8 := GenS1{F0: "foo bar", F1: (p0 & p0)}
            _ = v7
            _ = v8
        }
    }
    var v9 int32 = 2147483647
    fmt.Println(v9, p0)
    var v10 float32 = -2.25
    _ = p0
    _ = v10
    _ = v9
    return 876
}
func genFn1(p0 int32, p1 int64) uint32 {
    v11 := GenS1{F0: (("foo bar" + "a") + "hello"), F1: 912}
    for fi12 := 0; fi12 < 4; fi12++ {
        p1 -= (p1 >> ((1) & 63))
        if (false && false) {
            var v13 uint = 1
            fmt.Println(v13)
            _ = v13
        } else {
            var v14 int32 = p0
            p0 = v14
            _ = v14
        }
        _ = fi12
    }
    var v15 string = "a"
    v15 = v15
    _ = p0
    _ = p1
    _ = v11
    _ = v15
    return 2147483648
}

func main() {
    v1 := []int{}
    var v2 float32 = 0.0
    v3 := []string{"\n\t", "foo bar"}
    var v4 int8 = 126
    v5 := make(map[int]bool)
    v4 = v4
    v4 += -18
    _ = len(v5)
    switch (("a" + "foo bar") + "") {
    case "\n\t":
        switch "hello" {
        case "a":
            if true {
                fmt.Println(v2, v2)
            } else {
                for fi16 := 0; fi16 < 12; fi16++ {
                    switch fi16 {
                    case -9223372036854775808:
                        var v17 float64 = float64((1.5 * (v2 / v2)))
                        _ = v17
                    case 1:
                        _ = len(v1)
                    case -115:
                        var v18 uint16 = 65535
                        _ = v18
                    }
                    delete(v5, -842)
                    if (true && true) {
                        for fi19 := 0; fi19 < 11; fi19++ {
                            fmt.Println(fi16, v4)
                            _ = len(v5)
                            _ = fi19
                        }
                    } else {
                        if true {
                            var v20 uint32 = 0
                            fmt.Println(v2, v4)
                            _ = len(v3)
                            _ = v20
                        }
                        _ = len(v3)
                        for fi21 := 0; fi21 < 7; fi21++ {
                            var v22 uint32 = 2
                            var v23 uint16 = 32767
                            var v24 uint32 = v22
                            _ = fi21
                            _ = v22
                            _ = v23
                            _ = v24
                        }
                    }
                    _ = fi16
                }
                v25 := GenS1{F0: (("foo bar" + "hello") + "a"), F1: -32768}
                _ = v25
            }
        case "\n\t":
            fmt.Println(v2, v4)
        case "":
            v2 -= 3.14159
        default:
            v2 -= (3.14159 / (v2 * float32(1e10)))
        }
    case "hello":
        for fi26 := 0; fi26 < 4; fi26++ {
            v4 -= (v4 + (v4 ^ (v4 * -128)))
            v2 += 3.14159
            var v27 bool = ((!(!false)) && (366 == -350))
            _ = fi26
            _ = v27
        }
    default:
        v2 += v2
    }
    switch ("\n\t" + "foo bar") {
    case "\n\t":
        delete(v5, genFn0(42))
    default:
        switch -2147483648 {
        case 32767:
            fmt.Println(v4, v4)
        }
    }
    var v28 int8 = v4
    fmt.Println(v4, v28)
    var v29 bool = (!(48 < 128))
    switch -1 {
    case -9223372036854775808:
        fmt.Println(v2, v2, v4)
    case -296:
        if v29 {
            v30 := GenS1{F0: "a", F1: 42}
            var v31 int64 = 0
            _ = v30
            _ = v31
        }
    case 32767:
        delete(v5, 32768)
    default:
        if v29 {
            var v32 uint64 = 0
            var v33 bool = v29
            var v34 string = "a"
            _ = v32
            _ = v33
            _ = v34
        } else {
            if false {
                _ = len(v3)
                v4 = (v4 - (v4 * (v4 ^ v28)))
                fmt.Println(v29, v4, v28)
            }
            if v29 {
                var v35 int = 4294967295
                _ = v35
            }
        }
    }
    if v29 {
        switch ("hello" + "foo bar") {
        case "a":
            _ = len(v1)
        case "\n\t":
            var v36 uint8 = 127
            _ = v36
        default:
            v28 -= (v4 * (v4 * (v28 >> ((2) & 63))))
        }
        var v37 uint32 = genFn1(-225, -1)
        _ = v37
    }
    v29 = (!(-983 <= -2147483648))
    if ((!v29) && true) {
        v29 = false
    }
    var v38 uint8 = 128
    v38 += v38
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v38
    _ = v4
    _ = v5
}
