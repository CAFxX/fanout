package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint; F1 int }
type GenS6 struct { F0 int8; F1 int8; F2 uint8 }

func main() {
    var v1 uint16 = 497
    var v2 float64 = 1.5
    var v3 int16 = 0
    v4 := make(map[uint32]bool)
    for fi5 := 0; fi5 < 10; fi5++ {
        switch "\n\t" {
        case "a":
            if (((float64(((float32(float64(3.14159)) + ((1e10 / (3.14159 + 1e10)) - ((1e10 * 3.14159) / 0.1))) * float32((((0.0 * float64(float32(v2))) / ((v2 + -2.25) / (v2 * 1.5))) * 3.14159)))) - (-2.25 * 1e10)) + v2) < v2) {
                v1 += v1
                fmt.Println(fi5, v2)
            }
        case "\n\t":
            if (65535 < 1) {
                v6 := GenS1{F0: 42, F1: (!false), F2: ((float32(v2) - (0.0 * 3.14159)) >= float32((0.1 - ((float64(3.14159) * v2) - (((1e10 * float64(1.5)) * v2) - -2.25)))))}
                v7 := []int{}
                _ = v6
                _ = v7
            }
        default:
            switch (fi5 + (fi5 - fi5)) {
            case 969:
                v8 := GenS1{F0: 242, F1: (42 < 42), F2: ((false || (false && ((-1 > 42) || false))) && true)}
                _ = v8
            case 32768:
                if false {
                    var v9 uint64 = 1
                    fmt.Println(v9)
                    v1 += 175
                    _ = v9
                } else {
                    var v10 uint16 = v1
                    _ = v10
                }
            case 128:
                fmt.Println(v3)
            default:
                v11 := GenS2{F0: "foo bar"}
                _ = v11
            }
        }
        _ = fi5
    }
    v12 := GenS2{F0: ((("foo bar" + "a") + "") + (("foo bar" + (((("hello" + "") + ("a" + "\n\t")) + "") + "hello")) + ""))}
    if false {
        var v13 int = 4294967295
        for fi14 := 0; fi14 < 3; fi14++ {
            v2 = 0.0
            var v15 int16 = -1
            for fi16 := 0; fi16 < 14; fi16++ {
                var v17 uint16 = 255
                v18 := GenS2{F0: "\n\t"}
                var v19 uint16 = v1
                switch "a" {
                case "":
                    v20 := GenS2{F0: ""}
                    _ = v20
                case "a":
                    fmt.Println(v13)
                case "hello":
                    fmt.Println(v1, fi14, v15)
                }
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
            }
            _ = fi14
            _ = v15
        }
        _ = v13
    } else {
        var v21 bool = true
        v22 := v4[4294967295]
        if false {
            switch -688 {
            case 32767:
                fmt.Println(v21, v2)
            }
        } else {
            if v22 {
                v22 = ((!(-2147483648 != int(v1))) || (820 == 850))
            } else {
                v21 = ((0 > 0) || (128 <= 128))
            }
            switch 127 {
            case -1:
                var v23 string = "foo bar"
                _ = v23
            case 241:
                v24 := GenS1{F0: 42, F1: (387 <= uint(v1)), F2: (!(-758 == 1))}
                _ = v24
            }
            _ = len(v4)
        }
        _ = v21
        _ = v22
    }
    if true {
        var v25 uint16 = v1
        fmt.Println(v1, v3, v3)
        switch 128 {
        case -1:
            var v26 float32 = 3.14159
            _ = v26
        default:
            switch -168 {
            case -7:
                v27 := GenS3{F0: 0}
                _ = v27
            }
        }
        _ = v25
    }
    var v28 uint = 255
    v29 := GenS3{F0: (v28 | (v28 | (v28 >> ((v28) & 63))))}
    v28 = v28
    fmt.Println(v1, v3, v28)
    v28 = (v28 ^ (v28 / (((^(v28))) | 1)))
    if true {
        var v30 int32 = -2147483648
        for fi31 := 0; fi31 < 17; fi31++ {
            v32 := GenS4{F0: v1}
            v33 := []int{(fi31 >> (((v28 & (v28 - v28))) & 63))}
            _ = fi31
            _ = v32
            _ = v33
        }
        _ = v30
    } else {
        v34 := GenS5{F0: v28, F1: 128}
        v28 = (v28 - v28)
        var v35 int64 = 976
        _ = v34
        _ = v35
    }
    var v36 bool = true
    for fi37 := 0; fi37 < 10; fi37++ {
        if (v36 && true) {
            v38 := GenS1{F0: 1, F1: (v2 == (0.0 + (v2 - ((-2.25 * 3.14159) * -2.25)))), F2: (0 > 185)}
            _ = v38
        }
        v36 = false
        delete(v4, 301)
        _ = fi37
    }
    if ((v3 >= (^(v3))) && (v36 && v36)) {
        for fi39 := 0; fi39 < 2; fi39++ {
            v28 = 65535
            _ = fi39
        }
        v40 := GenS6{F0: 42, F1: -1, F2: 67}
        v41 := GenS6{F0: 1, F1: -1, F2: 1}
        _ = v40
        _ = v41
    }
    var v42 uint = (^(v28))
    fmt.Println(v3, v36, v28)
    v36 = (v36 || v36)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    _ = v1
    _ = v12
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v36
    _ = v4
    _ = v42
}
