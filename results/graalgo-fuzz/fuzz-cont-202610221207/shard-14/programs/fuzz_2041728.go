package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8; F1 uint8; F2 int }
type GenS5 struct { F0 int; F1 int8; F2 string }
func genFn0() int64 {
    if true {
        v4 := GenS1{F0: 65535}
        v5 := GenS1{F0: 951}
        for fi6 := 0; fi6 < 5; fi6++ {
            var v7 uint32 = 1
            if (true || (42 == 2)) {
                v7 = 724
            }
            var v8 uint64 = 279
            _ = fi6
            _ = v7
            _ = v8
        }
        _ = v4
        _ = v5
    }
    switch -630 {
    case -2147483648:
        if false {
            v9 := GenS2{F0: 42}
            _ = v9
        }
    case 32768:
        var v10 int64 = -1
        _ = v10
    case 496:
        var v11 int32 = 42
        _ = v11
    default:
        if true {
            v12 := []int{609, 0, -2147483648}
            if ((((3.14159 * 0.1) - float64((0.0 * 1.5))) >= float64((1e10 - 3.14159))) && (!true)) {
                var v13 float64 = (0.0 + (3.14159 + 3.14159))
                if len(v12) > 0 {
                    v12[0] = -643
                }
                _ = v13
            } else {
                if ((1.5 != 3.14159) || true) {
                    v14 := GenS1{F0: 1}
                    _ = v14
                } else {
                    var v15 int16 = 1000
                    var v16 string = "foo bar"
                    _ = v15
                    _ = v16
                }
                var v17 uint16 = 32767
                _ = v17
            }
            _ = v12
        } else {
            var v18 int = 128
            v19 := GenS2{F0: 32767}
            v20 := []int{v18, (-(v18)), (v18 ^ v18)}
            _ = v18
            _ = v19
            _ = v20
        }
    }
    var v21 float64 = float64(1e10)
    v22 := []int{}
    _ = v21
    _ = v22
    return 1
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]bool)
    v3 := make(map[uint32]string)
    delete(v2, (("" + "") + (("foo bar" + ("" + "a")) + "hello")))
    var v23 int8 = -1
    for fi24 := 0; fi24 < 12; fi24++ {
        v23 = v23
        _ = fi24
    }
    var v25 float64 = -2.25
    fmt.Println(v25)
    if ("foo bar" < (("hello" + "a") + "a")) {
        switch 1 {
        case 0:
            delete(v1, 2147483647)
        default:
            v25 = v25
        }
        v26 := GenS3{F0: -1}
        if false {
            v27 := GenS2{F0: 42}
            v28 := GenS4{F0: 43, F1: 1, F2: -489}
            _ = v27
            _ = v28
        } else {
            var v29 int32 = -1
            var v30 int64 = 0
            _ = v29
            _ = v30
        }
        _ = v26
    }
    var v31 float32 = (3.14159 * -2.25)
    switch ("\n\t" + "a") {
    case "a":
        fmt.Println(v25)
    case "\n\t":
        v32 := GenS3{F0: -478}
        _ = v32
    case "":
        if true {
            if true {
                v33 := []int{}
                switch -5 {
                case 9223372036854775807:
                    v31 += v31
                case -2147483648:
                    v25 = v25
                }
                v33 = append(v33, -9223372036854775808)
                _ = v33
            }
            if (float64(3.14159) <= ((1e10 - float64(((v31 / v31) * (v31 - ((3.14159 + float32((v25 / (-2.25 - ((v25 - float64(float32(v25))) / 1e10))))) - v31))))) / v25)) {
                for fi34 := 0; fi34 < 9; fi34++ {
                    fmt.Println(v23, v25)
                    v31 = ((v31 - v31) / ((v31 + v31) + (0.1 * v31)))
                    _ = fi34
                }
                v23 = -1
            } else {
                var v35 uint8 = 121
                _ = v35
            }
            v36 := GenS4{F0: 0, F1: 127, F2: 4294967295}
            _ = v36
        }
    default:
        if (((-9223372036854775808 > 128) || (!false)) && ((1 >= 852) && (!false))) {
            var v37 int64 = -9223372036854775808
            _ = v37
        } else {
            var v38 int64 = genFn0()
            var v39 int16 = 32767
            v40 := []int{2147483647}
            _ = v38
            _ = v39
            _ = v40
        }
    }
    v25 += (1e10 - 1.5)
    v41 := GenS5{F0: -2147483648, F1: -96, F2: ("" + "hello")}
    for fi42 := 0; fi42 < 10; fi42++ {
        v31 += (-2.25 / -2.25)
        v25 -= v25
        _ = fi42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v3
    _ = v31
    _ = v41
}
