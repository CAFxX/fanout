package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 uint32 }
func genFn0() int8 {
    var v4 int16 = -1
    if (-2.25 == (3.14159 / 1e10)) {
        v5 := GenS1{F0: 42, F1: 255, F2: 0}
        _ = v5
    } else {
        v6 := []int{32767, 127}
        fmt.Println(v4)
        _ = v6
    }
    for fi7 := 0; fi7 < 2; fi7++ {
        v8 := []int{(-(fi7)), 32767}
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 uint16 = 255
        _ = fi9
        _ = v10
    }
    _ = v4
    return 0
}

func main() {
    var v1 uint16 = 699
    var v2 string = "\n\t"
    var v3 uint64 = 128
    var v11 string = (v2 + v2)
    if false {
        var v12 uint64 = (v3 % ((v3) | 1))
        _ = v12
    }
    if true {
        var v13 float64 = 1.5
        v3 += (v3 + 42)
        fmt.Println(v3)
        _ = v13
    }
    var v14 uint = 680
    v2 = ("hello" + "a")
    for fi15 := 0; fi15 < 2; fi15++ {
        var v16 int16 = -971
        var v17 uint64 = v3
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 uint16 = v1
    v1 = v18
    v2 = ((((v2 + v2) + v2) + "\n\t") + ((v11 + v11) + v11))
    switch 2147483647 {
    case -782:
        v3 = (^(v3))
    case -1:
        v19 := []int{1, -908}
        _ = v19
    case 128:
        var v20 bool = (!true)
        _ = v20
    default:
        v21 := []int{4294967295}
        _ = v21
    }
    var v22 string = ("" + v2)
    var v23 float32 = 0.0
    var v24 uint = v14
    var v25 uint32 = 81
    if (!true) {
        for fi26 := 0; fi26 < 8; fi26++ {
            v2 = "foo bar"
            _ = fi26
        }
        switch (("" + v22) + v11) {
        case "a":
            switch 66 {
            case -9223372036854775808:
                var v27 int16 = 1000
                _ = v27
            case 4294967295:
                if false {
                    var v28 uint64 = v3
                    v2 = (v11 + "a")
                    _ = v28
                } else {
                    v25 = (v25 | (v25 << ((v14) & 63)))
                }
            case 9223372036854775807:
                if true {
                    v23 = 1e10
                    fmt.Println(v23)
                } else {
                    v25 = (v25 | (v25 << ((v14) & 63)))
                    switch "\n\t" {
                    case "hello":
                        fmt.Println(v23, v11)
                    case "\n\t":
                        var v29 int16 = 0
                        _ = v29
                    }
                    v3 -= 1
                }
            }
        case "hello":
            v14 = 42
        }
        switch 419 {
        case 32768:
            fmt.Println(v3, v18)
        case 127:
            for fi30 := 0; fi30 < 20; fi30++ {
                if false {
                    v25 = v25
                    fmt.Println(v1, v1)
                } else {
                    switch v22 {
                    case "a":
                        v31 := []int{int(v18), fi30, fi30}
                        _ = v31
                    }
                    for fi32 := 0; fi32 < 10; fi32++ {
                        v33 := []int{(fi32 | 0), fi32}
                        var v34 uint16 = (v18 << (((v24 % ((v14) | 1))) & 63))
                        v18 = 255
                        if (!true) {
                            var v35 int32 = 2147483647
                            var v36 uint32 = (v25 + uint32(v1))
                            _ = len(v33)
                            _ = v35
                            _ = v36
                        } else {
                            v37 := GenS1{F0: fi32, F1: v1, F2: v25}
                            _ = v37
                        }
                        _ = fi32
                        _ = v33
                        _ = v34
                    }
                }
                _ = fi30
            }
        case 1:
            switch ("hello" + v2) {
            case "hello":
                fmt.Println(v14, v25)
            }
        }
    }
    var v38 int8 = 127
    for fi39 := 0; fi39 < 8; fi39++ {
        v1 -= 2
        _ = fi39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v11
    _ = v14
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v38
}
