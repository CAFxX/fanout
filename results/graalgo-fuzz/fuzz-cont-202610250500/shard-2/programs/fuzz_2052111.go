package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint8; F1 int64; F2 bool }
func genFn0(p0 int8) uint8 {
    var v4 int = -21
    v5 := []int{v4, (v4 + (v4 * (v4 << ((18446744073709551615) & 63)))), (^(v4)), v4}
    var v6 int64 = 42
    for fi7 := 0; fi7 < 2; fi7++ {
        fmt.Println(v6)
        _ = fi7
    }
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 241
}

func main() {
    var v1 uint = 255
    var v2 float64 = (0.1 * float64((3.14159 * -2.25)))
    var v3 int64 = 0
    switch "a" {
    case "\n\t":
        var v8 int8 = 100
        _ = v8
    default:
        switch "\n\t" {
        case "":
            v3 += (v3 << ((v1) & 63))
        }
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        fmt.Println(v2)
        _ = fi9
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        fmt.Println(v2, v2)
        v11 := []int{int(v3)}
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        var v13 uint = v1
        var v14 int64 = v3
        v13 = v13
        _ = fi12
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        if true {
            for fi16 := 0; fi16 < 11; fi16++ {
                v17 := GenS1{F0: -111, F1: ""}
                v3 -= -1
                fmt.Println(v1)
                _ = fi16
                _ = v17
            }
            v2 = v2
        } else {
            v18 := GenS2{F0: 127}
            v1 += (v1 % ((v1) | 1))
            _ = v18
        }
        if ((true || (!true)) || true) {
            if true {
                if false {
                    v19 := GenS3{F0: "\n\t"}
                    _ = v19
                }
                v1 = v1
                v20 := GenS3{F0: "a"}
                _ = v20
            }
            v21 := GenS3{F0: "foo bar"}
            v22 := []int{fi15, (^(fi15)), (fi15 ^ fi15)}
            _ = v21
            _ = v22
        }
        _ = fi15
    }
    fmt.Println(v3, v1, v2)
    v3 += v3
    var v23 uint64 = uint64(v1)
    switch 833 {
    case 448:
        v23 = v23
    case 42:
        fmt.Println(v1, v2, v2)
    }
    if (false && ((^(v23)) == v23)) {
        var v24 bool = (("a" + "a") != "")
        switch 127 {
        case 32767:
            v25 := GenS1{F0: -605, F1: (("" + ("\n\t" + (("hello" + "foo bar") + ("foo bar" + "a")))) + "a")}
            _ = v25
        case 4294967295:
            if (true && v24) {
                v26 := GenS4{F0: (v3 & v3)}
                var v27 float64 = float64((-2.25 + -2.25))
                v27 += float64(-2.25)
                _ = v26
                _ = v27
            }
        }
        v28 := GenS1{F0: -1, F1: "\n\t"}
        _ = v24
        _ = v28
    }
    fmt.Println(v2)
    var v29 float32 = (3.14159 - float32(v2))
    for fi30 := 0; fi30 < 17; fi30++ {
        var v31 int = 32767
        for fi32 := 0; fi32 < 19; fi32++ {
            switch (^(fi32)) {
            case -628:
                v29 += float32(v2)
            case 0:
                switch ("\n\t" + "hello") {
                case "a":
                    v33 := []int{(-(fi30))}
                    _ = v33
                case "":
                    v3 = v3
                case "foo bar":
                    switch -131 {
                    case 626:
                        if false {
                            v34 := GenS5{F0: 2, F1: (^(v3)), F2: true}
                            _ = v34
                        }
                    default:
                        var v35 uint64 = 636
                        _ = v35
                    }
                default:
                    fmt.Println(v23)
                }
            case 128:
                v23 -= 387
            default:
                v1 = 320
            }
            v3 += v3
            _ = fi32
        }
        _ = fi30
        _ = v31
    }
    var v36 int8 = 115
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v36
}
