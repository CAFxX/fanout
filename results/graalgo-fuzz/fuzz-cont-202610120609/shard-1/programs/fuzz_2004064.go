package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint64 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 int64 }
func genFn0(p0 string) bool {
    v7 := GenS1{F0: 152, F1: 65535, F2: 2}
    v8 := []int{-1}
    _ = p0
    _ = v7
    _ = v8
    return true
}
func genFn1() uint16 {
    var v9 bool = (-727 <= 259)
    fmt.Println(v9)
    _ = v9
    return 592
}

func main() {
    v1 := make(map[string]bool)
    v2 := []int{4294967295, 9223372036854775807, 0, -1}
    v3 := make(map[uint32]string)
    v4 := make(map[string]string)
    v5 := []uint32{4294967295, 2147483648, 65535}
    var v6 int32 = 2147483647
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 uint = 2
        switch 202 {
        case 4294967295:
            v12 := v4[(("foo bar" + "a") + "hello")]
            _ = v12
        case 967:
            v13 := v1[""]
            _ = v13
        }
        _ = len(v5)
        _ = fi10
        _ = v11
    }
    v14 := GenS1{F0: 79, F1: 2, F2: 0}
    v6 += v6
    v6 = v6
    for fi15 := 0; fi15 < 20; fi15++ {
        var v16 uint8 = 1
        var v17 string = "hello"
        _ = fi15
        _ = v16
        _ = v17
    }
    if false {
        if false {
            v6 = (-(v6))
        } else {
            v18 := GenS2{F0: 1, F1: 42}
            _ = v18
        }
        v19 := GenS3{F0: -1}
        var v20 bool = false
        _ = v19
        _ = v20
    }
    var v21 uint = 410
    for fi22 := 0; fi22 < 13; fi22++ {
        v23 := GenS1{F0: 1, F1: 2, F2: 0}
        switch ("a" + "\n\t") {
        case "a":
            switch "" {
            case "":
                var v24 int16 = 32767
                _ = v24
            case "a":
                fmt.Println(v6, v21, v6)
            case "\n\t":
                v25 := GenS2{F0: 128, F1: 32767}
                _ = v25
            default:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
        case "foo bar":
            var v26 float64 = 1.5
            _ = v26
        case "\n\t":
            var v27 int = (fi22 & (-(fi22)))
            _ = v27
        }
        _ = fi22
        _ = v23
    }
    if len(v5) > 0 {
        v5[0] = 2147483648
    }
    if true {
        for fi28 := 0; fi28 < 19; fi28++ {
            var v29 float64 = 1.5
            var v30 int32 = (v6 >> ((v21) & 63))
            if (true && ((781 != 18446744073709551615) && true)) {
                var v31 string = ("foo bar" + (((("hello" + (("foo bar" + "") + "\n\t")) + "foo bar") + "foo bar") + ""))
                _ = v31
            }
            _ = fi28
            _ = v29
            _ = v30
        }
    } else {
        for fi32 := 0; fi32 < 4; fi32++ {
            v33 := GenS2{F0: 108, F1: 1000}
            v21 = (v21 / ((v21) | 1))
            delete(v1, ((((("foo bar" + ("foo bar" + "foo bar")) + "hello") + "") + (("hello" + (("foo bar" + "\n\t") + "\n\t")) + ("foo bar" + "a"))) + ((("foo bar" + "hello") + ((("a" + "\n\t") + "") + "hello")) + ((("a" + "foo bar") + "hello") + ("a" + "")))))
            _ = fi32
            _ = v33
        }
        var v34 int16 = 1
        _ = v34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
