package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 string; F1 string }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint8; F1 uint64; F2 int8 }
func genFn0() uint8 {
    var v6 int64 = 1
    v7 := GenS1{F0: 127}
    var v8 uint = 65535
    _ = v6
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 float32 = (1e10 - 1.5)
    v2 := make(map[string]int)
    var v3 int16 = -672
    v4 := make(map[string]string)
    var v5 int16 = (v3 | 315)
    v9 := GenS1{F0: 96}
    if ((-445 != 0) && true) {
        if false {
            for fi10 := 0; fi10 < 6; fi10++ {
                fmt.Println(v3, v5)
                fmt.Println(v5, v5)
                v11 := []int{127, 9223372036854775807}
                _ = fi10
                _ = v11
            }
            for fi12 := 0; fi12 < 3; fi12++ {
                v13 := GenS1{F0: -1}
                fmt.Println(v3, v3)
                _ = fi12
                _ = v13
            }
        }
        var v14 int = -1
        _ = v14
    } else {
        var v15 int32 = -334
        for fi16 := 0; fi16 < 7; fi16++ {
            fmt.Println(v5, v5)
            var v17 int8 = -20
            _ = fi16
            _ = v17
        }
        if false {
            v18 := GenS1{F0: 2}
            _ = v18
        } else {
            var v19 int16 = v3
            v19 += (-(v19))
            switch "a" {
            case "foo bar":
                var v20 int8 = -78
                _ = v20
            case "":
                v21 := v2["hello"]
                _ = v21
            case "a":
                v19 = v19
            default:
                switch 4294967295 {
                case 4294967295:
                    v19 = v3
                case 190:
                    v22 := []int{0, 852}
                    _ = v22
                case 128:
                    v23 := []int{0, -796, 32768}
                    _ = v23
                default:
                    var v24 float32 = ((v1 - (((v1 - (v1 + ((v1 + v1) * ((((3.14159 + v1) + v1) + 0.1) + v1)))) + v1) * v1)) + v1)
                    _ = v24
                }
            }
            _ = v19
        }
        _ = v15
    }
    v25 := GenS1{F0: 42}
    if true {
        for fi26 := 0; fi26 < 19; fi26++ {
            fmt.Println(v1)
            v27 := GenS2{F0: ("a" + "a"), F1: 42}
            _ = fi26
            _ = v27
        }
    } else {
        fmt.Println(v3, v5, v1)
    }
    var v28 int16 = v5
    var v29 uint16 = 255
    fmt.Println(v1)
    if (4294967295 < 18446744073709551615) {
        v3 -= -501
        switch -708 {
        case -9223372036854775808:
            switch -114 {
            case 385:
                v30 := GenS2{F0: "a", F1: 127}
                _ = v30
            case 14:
                v31 := GenS3{F0: ("hello" + ("foo bar" + ("hello" + "hello"))), F1: "a"}
                _ = v31
            }
        case -379:
            v32 := GenS4{F0: -125}
            _ = v32
        default:
            var v33 int32 = 1
            _ = v33
        }
    } else {
        fmt.Println(v28)
        switch "hello" {
        case "foo bar":
            fmt.Println(v28)
        }
    }
    for fi34 := 0; fi34 < 3; fi34++ {
        var v35 string = "a"
        var v36 uint32 = 319
        v28 -= 801
        _ = fi34
        _ = v35
        _ = v36
    }
    v37 := v4[((("\n\t" + "foo bar") + (("hello" + "hello") + "foo bar")) + "foo bar")]
    switch "hello" {
    case "\n\t":
        v38 := GenS5{F0: 42, F1: 936, F2: -128}
        _ = v38
    case "hello":
        _ = len(v2)
    case "":
        var v39 uint64 = 1
        _ = v39
    default:
        fmt.Println(v29, v1, v29)
    }
    var v40 float64 = (1.5 + (-2.25 + 0.1))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v37
    _ = v4
    _ = v40
    _ = v5
    _ = v9
}
