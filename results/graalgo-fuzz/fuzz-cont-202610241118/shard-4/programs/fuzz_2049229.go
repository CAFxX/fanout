package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 uint64; F1 int; F2 uint16 }
func genFn0(p0 bool) int32 {
    for fi6 := 0; fi6 < 10; fi6++ {
        for fi7 := 0; fi7 < 9; fi7++ {
            var v8 int64 = -1
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    fmt.Println(p0)
    _ = p0
    return 0
}
func genFn1(p0 int8, p1 uint8) int64 {
    v9 := GenS1{F0: (false || false), F1: -222}
    for fi10 := 0; fi10 < 4; fi10++ {
        fmt.Println(p1, p0)
        p0 = 60
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = v9
    return 0
}

func main() {
    var v1 uint32 = 0
    var v2 float32 = (3.14159 - 3.14159)
    var v3 bool = true
    v4 := make(map[int]int)
    var v5 uint = 65535
    if ((2 <= 65535) && true) {
        v3 = (!(!("a" > "a")))
        if false {
            var v11 string = ""
            var v12 int64 = -364
            v13 := GenS1{F0: false, F1: -544}
            _ = v11
            _ = v12
            _ = v13
        }
        v14 := []int{}
        _ = v14
    }
    v15 := GenS2{F0: v1, F1: 2}
    switch ("\n\t" + (("" + "foo bar") + ("" + "foo bar"))) {
    case "foo bar":
        switch -599 {
        case 32767:
            var v16 uint16 = 255
            _ = v16
        default:
            v17 := []int{}
            _ = v17
        }
    case "":
        switch -2147483648 {
        case 1:
            v18 := GenS2{F0: 284, F1: 2}
            _ = v18
        case -2147483648:
            var v19 uint = (v5 | v5)
            _ = v19
        default:
            for fi20 := 0; fi20 < 20; fi20++ {
                if v3 {
                    var v21 string = ("a" + ("foo bar" + (("hello" + "") + (("a" + ("" + "foo bar")) + "a"))))
                    v5 = (v5 + (v5 + (v5 << ((uint(v1)) & 63))))
                    _ = v21
                }
                _ = fi20
            }
        }
    case "a":
        v22 := []int{-9223372036854775808, 4294967295}
        _ = v22
    }
    var v23 string = "a"
    v24 := []int{32767, 32767, int(v5)}
    switch "\n\t" {
    case "foo bar":
        for fi25 := 0; fi25 < 20; fi25++ {
            v26 := GenS3{F0: "a", F1: v1}
            _ = fi25
            _ = v26
        }
    case "\n\t":
        v27 := GenS4{F0: 42, F1: -621, F2: 255}
        _ = v27
    case "hello":
        v1 = (v1 ^ 2)
    }
    for fi28 := 0; fi28 < 11; fi28++ {
        for fi29 := 0; fi29 < 4; fi29++ {
            var v30 int = fi29
            v31 := GenS4{F0: 42, F1: v30, F2: 1}
            _ = fi29
            _ = v30
            _ = v31
        }
        _ = fi28
    }
    fmt.Println(v5, v3, v23)
    switch 4294967295 {
    case 9223372036854775807:
        for fi32 := 0; fi32 < 17; fi32++ {
            switch "foo bar" {
            case "a":
                if len(v24) >= 2 {
                    v24 = v24[1:len(v24)]
                }
            case "hello":
                fmt.Println(v5)
            case "\n\t":
                var v33 int8 = 98
                _ = v33
            }
            var v34 float64 = 1e10
            _ = fi32
            _ = v34
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
