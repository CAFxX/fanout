package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool; F1 bool }

func main() {
    var v1 bool = true
    var v2 int = 683
    var v3 uint16 = 664
    var v4 uint = 4294967295
    v5 := make(map[string]int)
    var v6 string = (("" + "hello") + "")
    var v7 uint16 = 432
    v1 = (!true)
    v8 := GenS1{F0: 32767, F1: int64(v3)}
    var v9 bool = false
    for fi10 := 0; fi10 < 1; fi10++ {
        v7 = (v3 % (((v3 / (((^(v7))) | 1))) | 1))
        _ = fi10
    }
    switch int(v7) {
    case -684:
        if (1e10 >= (-2.25 + 1e10)) {
            for fi11 := 0; fi11 < 16; fi11++ {
                v7 = (v7 / ((v3) | 1))
                _ = fi11
            }
        } else {
            if true {
                if v9 {
                    v9 = false
                }
            } else {
                v12 := GenS2{F0: (v2 | -9223372036854775808), F1: v6}
                fmt.Println(v1)
                _ = v12
            }
            var v13 int16 = 815
            v14 := GenS3{F0: v9}
            _ = v13
            _ = v14
        }
    default:
        var v15 int64 = 42
        _ = v15
    }
    for fi16 := 0; fi16 < 2; fi16++ {
        v7 += (v7 & 42)
        _ = fi16
    }
    for fi17 := 0; fi17 < 14; fi17++ {
        v18 := GenS4{F0: ((1.5 - 0.1) <= 0.0), F1: false}
        var v19 float64 = 3.14159
        if v9 {
            v1 = (fi17 > fi17)
        } else {
            var v20 uint16 = v3
            _ = v20
        }
        _ = fi17
        _ = v18
        _ = v19
    }
    v6 = ("hello" + v6)
    var v21 int64 = 0
    fmt.Println(v6, v2)
    if v1 {
        v22 := []int{v2}
        var v23 int64 = (-(v21))
        switch ("\n\t" + (("\n\t" + v6) + v6)) {
        case "hello":
            if v9 {
                for fi24 := 0; fi24 < 2; fi24++ {
                    switch "foo bar" {
                    case "\n\t":
                        var v25 int64 = v21
                        _ = v25
                    }
                    for fi26 := 0; fi26 < 8; fi26++ {
                        var v27 int = fi24
                        _ = fi26
                        _ = v27
                    }
                    var v28 int8 = 38
                    _ = fi24
                    _ = v28
                }
            } else {
                var v29 uint8 = 242
                _ = v29
            }
        }
        _ = v22
        _ = v23
    } else {
        fmt.Println(v9, v9)
        v30 := GenS4{F0: v1, F1: v9}
        _ = v30
    }
    var v31 bool = ((v6 + (v6 + v6)) != (v6 + v6))
    v9 = (2 > v3)
    for fi32 := 0; fi32 < 12; fi32++ {
        if v1 {
            var v33 int16 = -1
            for fi34 := 0; fi34 < 12; fi34++ {
                v21 = v21
                _ = len(v5)
                var v35 uint16 = (v7 * (v3 << ((v4) & 63)))
                var v36 bool = (1 != -73)
                v31 = v36
                _ = fi34
                _ = v35
                _ = v36
            }
            v4 = v4
            _ = v33
        }
        switch v2 {
        case 2147483647:
            var v37 int16 = 42
            _ = v37
        case 9223372036854775807:
            if v1 {
                v3 += 32767
            }
        }
        _ = fi32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
