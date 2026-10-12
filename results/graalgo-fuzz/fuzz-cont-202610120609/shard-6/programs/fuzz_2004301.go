package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32 }
func genFn0() int16 {
    var v5 float64 = 1.5
    var v6 uint8 = 2
    v7 := []int{}
    var v8 uint32 = 2147483647
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return -888
}
func genFn1() string {
    switch "a" {
    case "\n\t":
        if false {
            var v9 uint32 = 2147483648
            var v10 bool = false
            var v11 int = -942
            _ = v10
            _ = v11
            _ = v9
        }
    }
    v12 := []int{}
    _ = v12
    return "foo bar"
}

func main() {
    var v1 int64 = -862
    v2 := make(map[uint32]string)
    var v3 string = "hello"
    v4 := make(map[int]int)
    var v13 int16 = 1000
    if (!true) {
        var v14 int8 = 127
        for fi15 := 0; fi15 < 18; fi15++ {
            var v16 uint8 = 1
            _ = fi15
            _ = v16
        }
        _ = v14
    }
    v3 = ("foo bar" + v3)
    v17 := GenS1{F0: 674}
    var v18 uint16 = 1
    var v19 string = (v3 + ((v3 + v3) + v3))
    for fi20 := 0; fi20 < 11; fi20++ {
        v21 := GenS1{F0: (v18 / ((65535) | 1))}
        v22 := v4[fi20]
        _ = fi20
        _ = v21
        _ = v22
    }
    var v23 uint16 = v18
    v19 = (("foo bar" + "") + ("foo bar" + (v3 + (((v3 + v3) + v19) + "a"))))
    switch 2147483647 {
    case 128:
        v13 -= (^(v13))
    case 782:
        v13 -= -32768
    case 276:
        v24 := []int{322}
        _ = v24
    }
    switch ("hello" + "\n\t") {
    case "":
        if (80 >= 128) {
            v23 = (v23 + (v18 - (v18 - v18)))
            if (-2147483648 > 0) {
                if (18446744073709551615 != 18446744073709551615) {
                    v25 := GenS1{F0: (v18 >> ((516) & 63))}
                    _ = v25
                } else {
                    switch "\n\t" {
                    case "\n\t":
                        fmt.Println(v23)
                    case "foo bar":
                        var v26 int16 = v13
                        _ = v26
                    }
                    v27 := []int{-987, 4294967295, -9223372036854775808}
                    _ = v27
                }
            }
        } else {
            switch 1 {
            case 1:
                v19 = "\n\t"
            case -1:
                fmt.Println(v1, v23)
            }
            var v28 string = (v19 + v19)
            _ = v28
        }
    case "\n\t":
        v29 := GenS2{F0: 1}
        _ = v29
    default:
        if true {
            if (!true) {
                v30 := []int{4294967295, 9223372036854775807}
                _ = v30
            } else {
                var v31 int16 = v13
                if (((((("a" + (v19 + v3)) + (v19 + (("" + "") + ("" + v19)))) + "foo bar") + v3) <= v19) || true) {
                    var v32 int8 = 127
                    _ = v32
                } else {
                    v33 := []int{1, 1, 0}
                    _ = v33
                }
                _ = v31
            }
        }
    }
    for fi34 := 0; fi34 < 20; fi34++ {
        switch ("\n\t" + v3) {
        case "a":
            for fi35 := 0; fi35 < 3; fi35++ {
                fmt.Println(v1)
                var v36 int8 = 4
                v19 = "\n\t"
                _ = fi35
                _ = v36
            }
        case "\n\t":
            v13 += v13
        case "hello":
            v37 := GenS2{F0: 2147483648}
            _ = v37
        }
        var v38 int64 = v1
        _ = fi34
        _ = v38
    }
    for fi39 := 0; fi39 < 13; fi39++ {
        var v40 float32 = (float32(-2.25) * 1.5)
        _ = fi39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
