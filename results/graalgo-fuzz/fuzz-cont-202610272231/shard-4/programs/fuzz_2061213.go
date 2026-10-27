package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int8; F1 int32; F2 uint16 }

func main() {
    var v1 uint = 2
    v2 := make(map[int]int)
    var v3 uint32 = 2147483648
    switch 2147483647 {
    case -347:
        v4 := []int{int(v3), -1}
        _ = v4
    case 1:
        var v5 float32 = float32(1.5)
        _ = v5
    case 9223372036854775807:
        for fi6 := 0; fi6 < 5; fi6++ {
            v3 = v3
            switch ("foo bar" + "a") {
            case "a":
                v7 := GenS1{F0: 32767}
                _ = v7
            }
            _ = fi6
        }
    }
    switch 880 {
    case -1:
        var v8 string = ("" + (("" + "") + ("hello" + "hello")))
        _ = v8
    case 1:
        var v9 float64 = (1e10 + 1e10)
        _ = v9
    case 127:
        var v10 string = "\n\t"
        _ = v10
    }
    var v11 int32 = -480
    var v12 float64 = (-2.25 * 3.14159)
    if true {
        v13 := GenS1{F0: int16(v3)}
        _ = v13
    } else {
        switch "\n\t" {
        case "":
            switch "foo bar" {
            case "\n\t":
                v14 := GenS1{F0: 42}
                _ = v14
            }
        case "foo bar":
            fmt.Println(v11, v12)
        case "hello":
            var v15 uint = v1
            _ = v15
        default:
            var v16 int16 = 42
            _ = v16
        }
        for fi17 := 0; fi17 < 18; fi17++ {
            for fi18 := 0; fi18 < 12; fi18++ {
                if ((!(true && false)) || (4294967295 < 4294967295)) {
                    var v19 float32 = (1.5 + -2.25)
                    _ = v19
                } else {
                    fmt.Println(fi18, fi17, fi17)
                    var v20 float64 = float64((((-2.25 * 1e10) + 1.5) + 0.0))
                    _ = v20
                }
                _ = fi18
            }
            delete(v2, int(v3))
            fmt.Println(v1, v12, v11)
            _ = fi17
        }
        var v21 uint8 = 2
        _ = v21
    }
    fmt.Println(v12)
    var v22 uint64 = 461
    switch "" {
    case "foo bar":
        var v23 int = -56
        _ = v23
    }
    v24 := GenS1{F0: 220}
    v3 = v3
    for fi25 := 0; fi25 < 14; fi25++ {
        var v26 int8 = 46
        _ = fi25
        _ = v26
    }
    var v27 uint8 = 1
    if (!true) {
        v11 = v11
        for fi28 := 0; fi28 < 3; fi28++ {
            if (42 < -9223372036854775808) {
                v1 -= (v1 / ((uint(v27)) | 1))
                var v29 uint = 1
                _ = v29
            } else {
                v12 -= 1.5
                var v30 int64 = 42
                var v31 uint8 = v27
                _ = v30
                _ = v31
            }
            v11 += int32(fi28)
            _ = fi28
        }
        var v32 uint16 = 1
        _ = v32
    } else {
        var v33 int16 = 42
        v3 -= 2147483648
        _ = v33
    }
    v34 := GenS2{F0: false, F1: -733}
    switch "foo bar" {
    case "a":
        v35 := GenS2{F0: (!("a" <= "hello")), F1: -989}
        _ = v35
    default:
        fmt.Println(v11, v27)
    }
    var v36 float64 = float64(-2.25)
    var v37 string = "\n\t"
    if ((v22 * 178) >= v22) {
        delete(v2, -815)
        v38 := GenS3{F0: 844}
        v39 := GenS4{F0: -128, F1: (v11 * -2147483648), F2: 656}
        _ = v38
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v24
    _ = v27
    _ = v3
    _ = v34
    _ = v36
    _ = v37
}
