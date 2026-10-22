package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 int64 }
func genFn0(p0 int16, p1 uint8) int64 {
    var v5 uint16 = 1
    var v6 bool = (!true)
    switch ("foo bar" + ("a" + "\n\t")) {
    case "\n\t":
        for fi7 := 0; fi7 < 18; fi7++ {
            v8 := GenS1{F0: -814, F1: 42, F2: p0}
            _ = fi7
            _ = v8
        }
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return -1
}

func main() {
    var v1 bool = true
    var v2 int64 = 0
    var v3 float32 = -2.25
    var v4 uint = 1
    fmt.Println(v1, v3)
    if v1 {
        switch (("" + ("foo bar" + ("hello" + "a"))) + "") {
        case "a":
            if (-2147483648 < -162) {
                v9 := GenS1{F0: -120, F1: 0, F2: 32767}
                v2 -= 42
                var v10 string = "a"
                _ = v10
                _ = v9
            } else {
                v11 := GenS2{F0: 2}
                v1 = v1
                switch "hello" {
                case "hello":
                    fmt.Println(v3, v1, v2)
                default:
                    for fi12 := 0; fi12 < 14; fi12++ {
                        var v13 float64 = 0.0
                        _ = fi12
                        _ = v13
                    }
                }
                _ = v11
            }
        case "hello":
            var v14 uint = (v4 - v4)
            _ = v14
        case "\n\t":
            v15 := GenS2{F0: 255}
            _ = v15
        }
        for fi16 := 0; fi16 < 8; fi16++ {
            var v17 int32 = 2147483647
            _ = fi16
            _ = v17
        }
    } else {
        var v18 bool = (-2147483648 == 0)
        v19 := []int{2147483647, 9223372036854775807, 4294967295}
        if v18 {
            if len(v19) >= 2 {
                v19 = v19[1:len(v19)]
            }
            for fi20 := 0; fi20 < 10; fi20++ {
                switch fi20 {
                case 128:
                    v19 = append(v19, -383)
                case 32767:
                    var v21 string = "a"
                    _ = v21
                }
                var v22 int = -9223372036854775808
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
                _ = fi20
                _ = v22
            }
        }
        _ = v18
        _ = v19
    }
    switch 42 {
    case 42:
        switch -14 {
        case -2147483648:
            var v23 int8 = 0
            _ = v23
        default:
            fmt.Println(v1, v3, v4)
        }
    case -824:
        v1 = false
    default:
        if (v1 || (1 < 226)) {
            v4 -= 0
            v4 -= 42
        }
    }
    v4 = v4
    switch 42 {
    case 9223372036854775807:
        var v24 float64 = (1.5 + 0.1)
        _ = v24
    case 984:
        fmt.Println(v4)
    default:
        v3 = v3
    }
    v25 := GenS3{F0: (v2 & 9223372036854775807), F1: (-32768 > 32767)}
    var v26 bool = true
    v27 := GenS4{F0: (v2 + 9223372036854775807)}
    if false {
        v1 = (!true)
        var v28 uint16 = 891
        _ = v28
    } else {
        if v1 {
            var v29 float32 = (v3 * 1e10)
            v30 := GenS1{F0: -32768, F1: 41, F2: 32767}
            v31 := []int{}
            _ = v29
            _ = v30
            _ = v31
        } else {
            for fi32 := 0; fi32 < 2; fi32++ {
                for fi33 := 0; fi33 < 17; fi33++ {
                    switch (fi32 * (^(fi32))) {
                    case 1:
                        var v34 uint32 = 882
                        _ = v34
                    case -9223372036854775808:
                        fmt.Println(fi33, v26)
                    default:
                        var v35 int32 = 42
                        _ = v35
                    }
                    _ = fi33
                }
                _ = fi32
            }
            var v36 uint = 42
            _ = v36
        }
        fmt.Println(v2, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
