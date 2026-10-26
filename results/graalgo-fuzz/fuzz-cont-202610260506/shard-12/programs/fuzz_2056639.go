package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 uint64; F1 bool; F2 uint16 }
type GenS3 struct { F0 bool }
func genFn0(p0 bool) uint {
    if false {
        var v4 string = "\n\t"
        var v5 int64 = 826
        v4 = "foo bar"
        _ = v4
        _ = v5
    } else {
        switch "foo bar" {
        case "hello":
            switch 32768 {
            case 127:
                switch ("hello" + "a") {
                case "a":
                    v6 := GenS1{F0: -23, F1: 671}
                    _ = v6
                case "foo bar":
                    v7 := []int{1, -428, 961}
                    _ = v7
                }
            }
        default:
            v8 := GenS1{F0: 42, F1: 33}
            _ = v8
        }
    }
    v9 := GenS1{F0: -2147483648, F1: 38}
    _ = p0
    _ = v9
    return 65535
}

func main() {
    var v1 int16 = 42
    v2 := make(map[int]bool)
    var v3 int = 32767
    v10 := v2[v3]
    delete(v2, v3)
    v11 := GenS2{F0: 1, F1: (2 == 194), F2: 0}
    for fi12 := 0; fi12 < 2; fi12++ {
        var v13 uint64 = 42
        var v14 uint8 = 48
        _ = fi12
        _ = v13
        _ = v14
    }
    v3 = 2147483647
    v15 := GenS1{F0: -565, F1: -901}
    var v16 int32 = 1
    var v17 bool = v10
    if v10 {
        var v18 float64 = (((-2.25 / -2.25) + float64((1.5 + 0.1))) + 1.5)
        _ = v18
    } else {
        v19 := []int{-1, v3, v3}
        _ = v19
    }
    for fi20 := 0; fi20 < 13; fi20++ {
        fmt.Println(v3, v16)
        _ = fi20
    }
    switch ("foo bar" + "hello") {
    case "a":
        for fi21 := 0; fi21 < 5; fi21++ {
            var v22 uint8 = 127
            v23 := []int{-1, v3}
            v24 := v2[(^(fi21))]
            _ = fi21
            _ = v22
            _ = v23
            _ = v24
        }
    case "foo bar":
        var v25 string = "\n\t"
        _ = v25
    case "":
        for fi26 := 0; fi26 < 12; fi26++ {
            var v27 uint32 = 601
            switch v3 {
            case -1:
                var v28 int = -1
                _ = v28
            default:
                delete(v2, (^(v3)))
            }
            _ = fi26
            _ = v27
        }
    default:
        v29 := GenS3{F0: v10}
        _ = v29
    }
    v17 = false
    for fi30 := 0; fi30 < 13; fi30++ {
        v31 := GenS2{F0: 42, F1: v17, F2: 42}
        if ("hello" < "a") {
            var v32 uint32 = 42
            _ = v32
        } else {
            var v33 bool = v10
            v34 := GenS3{F0: (!false)}
            v17 = (547 != 1)
            _ = v33
            _ = v34
        }
        _ = fi30
        _ = v31
    }
    var v35 float32 = 3.14159
    var v36 int8 = 37
    if (1e10 != (-2.25 / ((0.1 - ((float64((((v35 - 0.1) / 1.5) - (v35 * (0.1 - v35)))) * 0.1) / -2.25)) + (-2.25 * 3.14159)))) {
        var v37 uint = 4294967295
        if (!(817 < 42)) {
            var v38 int8 = v36
            v39 := GenS1{F0: v3, F1: v16}
            v40 := GenS2{F0: 18446744073709551615, F1: true, F2: 65535}
            _ = v38
            _ = v39
            _ = v40
        } else {
            v41 := GenS2{F0: 781, F1: (v10 && (false || (((false && v10) || ("a" <= (("a" + "\n\t") + "foo bar"))) || (true || v10)))), F2: 32767}
            _ = v41
        }
        _ = v37
    }
    v36 += (-(v36))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v35
    _ = v36
}
