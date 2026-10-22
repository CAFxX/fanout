package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 int16; F2 uint }
func genFn0(p0 uint16, p1 int) uint32 {
    var v7 bool = (965 == (p0 / ((770) | 1)))
    for fi8 := 0; fi8 < 12; fi8++ {
        if (!(1000 <= 32767)) {
            var v9 uint32 = 437
            for fi10 := 0; fi10 < 7; fi10++ {
                v11 := []int{-9223372036854775808, (p1 + fi10)}
                v12 := GenS1{F0: -518}
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = v9
        }
        if v7 {
            var v13 int8 = 127
            var v14 int64 = 9223372036854775807
            _ = v13
            _ = v14
        }
        _ = fi8
    }
    if v7 {
        var v15 string = (("hello" + "foo bar") + "hello")
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v7
    return 2147483647
}
func genFn1(p0 int64, p1 uint16) uint16 {
    p1 = 26
    p0 -= -1
    _ = p0
    _ = p1
    return p1
}

func main() {
    var v1 uint = 752
    var v2 uint8 = 146
    v3 := make(map[string]bool)
    var v4 uint32 = 1
    var v5 float64 = (-2.25 * -2.25)
    var v6 uint16 = 65535
    fmt.Println(v4, v1)
    v16 := []int{-485}
    var v17 uint32 = uint32(v6)
    switch 128 {
    case -786:
        var v18 uint8 = v2
        _ = v18
    }
    fmt.Println(v1, v4, v4)
    switch ("hello" + "") {
    case "foo bar":
        v19 := GenS2{F0: (!true), F1: 42, F2: v1}
        _ = v19
    case "\n\t":
        for fi20 := 0; fi20 < 14; fi20++ {
            v16 = append(v16, (fi20 % (((-(fi20))) | 1)))
            _ = fi20
        }
    }
    fmt.Println(v5)
    v1 = v1
    if false {
        fmt.Println(v5, v2, v17)
        if true {
            if true {
                switch 2147483647 {
                case 1:
                    var v21 string = ""
                    _ = v21
                }
            }
            var v22 string = "a"
            _ = v22
        }
    } else {
        var v23 string = ("foo bar" + "hello")
        var v24 string = v23
        for fi25 := 0; fi25 < 15; fi25++ {
            v26 := GenS1{F0: 154}
            v27 := GenS2{F0: true, F1: -1, F2: 4294967295}
            v16 = append(v16, 127)
            _ = fi25
            _ = v26
            _ = v27
        }
        _ = v23
        _ = v24
    }
    v28 := GenS1{F0: -2147483648}
    if false {
        switch ("a" + "a") {
        case "hello":
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
        case "foo bar":
            for fi29 := 0; fi29 < 6; fi29++ {
                v1 += (^(v1))
                var v30 int16 = 32767
                var v31 uint8 = 127
                _ = fi29
                _ = v30
                _ = v31
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
