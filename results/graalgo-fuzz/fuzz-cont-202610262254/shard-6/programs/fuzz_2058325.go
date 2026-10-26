package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 bool }
type GenS2 struct { F0 uint; F1 uint64; F2 uint32 }
func genFn0() uint32 {
    v4 := []int{-107}
    for fi5 := 0; fi5 < 13; fi5++ {
        switch "\n\t" {
        case "":
            var v6 int8 = 127
            _ = v6
        case "a":
            fmt.Println(fi5)
        }
        _ = fi5
    }
    if false {
        var v7 int8 = 1
        v8 := GenS1{F0: ("hello" + ("\n\t" + ("foo bar" + "hello"))), F1: 9223372036854775807, F2: false}
        if (!(2 <= 700)) {
            for fi9 := 0; fi9 < 7; fi9++ {
                fmt.Println(v7, v7)
                _ = fi9
            }
        }
        _ = v7
        _ = v8
    } else {
        if false {
            for fi10 := 0; fi10 < 1; fi10++ {
                var v11 string = "\n\t"
                _ = fi10
                _ = v11
            }
        } else {
            _ = len(v4)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v12 uint16 = 42
            _ = v12
        }
        var v13 int = 32767
        _ = v13
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v14 uint16 = 630
    _ = v14
    _ = v4
    return 65535
}
func genFn1() string {
    v15 := []int{9223372036854775807, -917}
    if ("hello" > "foo bar") {
        var v16 uint32 = genFn0()
        var v17 float64 = -2.25
        var v18 string = ("" + (("hello" + "foo bar") + "hello"))
        _ = v16
        _ = v17
        _ = v18
    } else {
        v15 = append(v15, 128)
        var v19 uint64 = 919
        _ = v19
    }
    _ = v15
    return "hello"
}

func main() {
    var v1 string = "foo bar"
    var v2 int = 32767
    v3 := []uint32{}
    v2 = v2
    for fi20 := 0; fi20 < 4; fi20++ {
        v21 := GenS1{F0: (v1 + ""), F1: fi20, F2: (32767 > 65535)}
        v1 = (v1 + v1)
        _ = fi20
        _ = v21
    }
    v22 := GenS2{F0: 65535, F1: 18446744073709551615, F2: 787}
    fmt.Println(v2, v1)
    _ = len(v3)
    if (0 < 32767) {
        var v23 float32 = 3.14159
        v23 += v23
        for fi24 := 0; fi24 < 12; fi24++ {
            v25 := GenS1{F0: ((v1 + (v1 + "\n\t")) + (v1 + v1)), F1: (^(v2)), F2: (false && false)}
            v26 := GenS2{F0: 231, F1: 4294967295, F2: genFn0()}
            var v27 bool = (850 > 41)
            _ = fi24
            _ = v25
            _ = v26
            _ = v27
        }
        _ = v23
    } else {
        fmt.Println(v2, v2)
        if (-1 < 42) {
            var v28 bool = (false && true)
            var v29 uint32 = 2
            _ = v28
            _ = v29
        }
    }
    var v30 uint = 647
    if (1 >= 937) {
        var v31 float32 = 0.0
        _ = v31
    } else {
        for fi32 := 0; fi32 < 16; fi32++ {
            fmt.Println(v1)
            _ = fi32
        }
        v30 -= 4294967295
    }
    var v33 uint = (v30 * v30)
    var v34 string = (v1 + v1)
    v34 = v1
    for fi35 := 0; fi35 < 17; fi35++ {
        for fi36 := 0; fi36 < 12; fi36++ {
            fmt.Println(fi36)
            _ = fi36
        }
        _ = fi35
    }
    v34 = "foo bar"
    if true {
        for fi37 := 0; fi37 < 16; fi37++ {
            fmt.Println(fi37, v1)
            var v38 int8 = -128
            _ = fi37
            _ = v38
        }
        var v39 int8 = -1
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v30
    _ = v33
    _ = v34
}
