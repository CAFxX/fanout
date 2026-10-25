package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int64 }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 int; F1 uint8 }
func genFn0(p0 uint16, p1 int, p2 bool) int64 {
    var v6 bool = p2
    fmt.Println(p2, p1, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 942
}
func genFn1(p0 int8) string {
    var v7 float32 = float32(3.14159)
    for fi8 := 0; fi8 < 3; fi8++ {
        fmt.Println(p0)
        _ = fi8
    }
    switch -2147483648 {
    case -632:
        if (-9223372036854775808 > 1) {
            var v9 uint32 = 0
            var v10 float32 = (0.0 * (v7 * 1e10))
            _ = v10
            _ = v9
        }
    case 999:
        for fi11 := 0; fi11 < 18; fi11++ {
            v12 := []int{fi11}
            var v13 uint = 1
            if true {
                for fi14 := 0; fi14 < 14; fi14++ {
                    if true {
                        fmt.Println(v13, v7, v13)
                        v13 = v13
                        if len(v12) >= 2 {
                            v12 = v12[1:len(v12)]
                        }
                        var v15 uint64 = 1
                        v16 := GenS1{F0: 0, F1: (v15 * (v15 >> ((231) & 63))), F2: -9223372036854775808}
                        if (("" < ((("a" + (("a" + "hello") + "")) + "hello") + (("foo bar" + "hello") + "hello"))) || false) {
                            p0 = 127
                        }
                        _ = v15
                        _ = v16
                    }
                    var v17 uint16 = 742
                    switch -9223372036854775808 {
                    case 42:
                        v17 -= v17
                    case -1:
                        var v18 float32 = (3.14159 + v7)
                        _ = v18
                    }
                    v13 -= 2
                    _ = fi14
                    _ = v17
                }
            }
            _ = fi11
            _ = v12
            _ = v13
        }
    }
    _ = p0
    _ = v7
    return ("" + ("" + ""))
}

func main() {
    var v1 uint = 65535
    var v2 uint = 1
    var v3 int8 = 127
    var v4 string = ((("a" + ("hello" + "a")) + "foo bar") + "a")
    v5 := make(map[string]uint32)
    v4 = (v4 + "hello")
    switch v4 {
    case "\n\t":
        switch 99 {
        case 9223372036854775807:
            if true {
                var v19 int64 = 1
                _ = v19
            }
        case 4294967295:
            var v20 uint8 = 2
            _ = v20
        default:
            var v21 float32 = 1.5
            _ = v21
        }
    default:
        for fi22 := 0; fi22 < 13; fi22++ {
            v23 := []int{32767}
            v5[""] = 482
            _ = fi22
            _ = v23
        }
    }
    for fi24 := 0; fi24 < 3; fi24++ {
        fmt.Println(v4, v4)
        var v25 uint16 = 65535
        _ = fi24
        _ = v25
    }
    v26 := []int{128, 127, 4294967295}
    if (!false) {
        var v27 float32 = (0.0 / 0.1)
        if ((v4 + v4) <= v4) {
            var v28 string = v4
            v29 := GenS2{F0: 128, F1: 2147483647}
            _ = len(v26)
            _ = v28
            _ = v29
        }
        _ = v27
    }
    if len(v26) >= 2 {
        v26 = v26[1:len(v26)]
    }
    if ((v3 * v3) > v3) {
        if false {
            delete(v5, "foo bar")
        }
        if ((((true && (255 == 1)) && ((v4 != v4) || false)) && true) && false) {
            var v30 int64 = 9223372036854775807
            _ = v30
        } else {
            for fi31 := 0; fi31 < 6; fi31++ {
                for fi32 := 0; fi32 < 11; fi32++ {
                    fmt.Println(fi32, v1, v1)
                    _ = fi32
                }
                _ = fi31
            }
            var v33 float32 = 0.1
            v1 = 505
            _ = v33
        }
    } else {
        v34 := GenS3{F0: 127, F1: 73}
        _ = v34
    }
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
