package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool; F1 uint16 }
func genFn0(p0 bool, p1 uint) int8 {
    var v5 int32 = 0
    p0 = ("hello" == "a")
    for fi6 := 0; fi6 < 10; fi6++ {
        v7 := []int{4294967295, 720, (-(fi6)), (fi6 & fi6)}
        v8 := GenS1{F0: (^(v5)), F1: 2, F2: 1}
        v9 := GenS1{F0: (-(v5)), F1: 32767, F2: 65535}
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v5
    return 100
}

func main() {
    var v1 uint8 = 151
    var v2 uint8 = (v1 / ((v1) | 1))
    v3 := make(map[string]string)
    var v4 float64 = 0.0
    v4 = v4
    v10 := GenS2{F0: 42}
    v11 := []int{4294967295, 32767}
    v3["foo bar"] = (("foo bar" + "") + ("foo bar" + ""))
    var v12 string = ("hello" + "\n\t")
    v4 += v4
    v2 -= 128
    if false {
        switch 32768 {
        case -283:
            v13 := GenS3{F0: 1}
            _ = v13
        case 136:
            v1 -= 127
        case 485:
            fmt.Println(v1)
        default:
            if false {
                _ = len(v3)
            }
        }
    }
    if (((!(182 <= v2)) || false) && true) {
        v2 -= (v2 << ((42) & 63))
        v2 += 128
    } else {
        var v14 uint8 = v1
        _ = v14
    }
    if false {
        v15 := GenS3{F0: 1}
        for fi16 := 0; fi16 < 4; fi16++ {
            for fi17 := 0; fi17 < 2; fi17++ {
                var v18 float32 = 1e10
                switch ("" + v12) {
                case "foo bar":
                    v11 = append(v11, (fi16 * (fi17 | -66)))
                }
                _ = fi17
                _ = v18
            }
            var v19 float64 = ((float64((3.14159 - 3.14159)) + v4) / 1e10)
            _ = fi16
            _ = v19
        }
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
        _ = v15
    } else {
        v20 := GenS3{F0: 4294967295}
        var v21 int8 = 127
        _ = v20
        _ = v21
    }
    v1 -= (v2 >> ((883) & 63))
    switch 32768 {
    case -2147483648:
        if false {
            var v22 uint64 = 18446744073709551615
            _ = v22
        }
    }
    v23 := v3["foo bar"]
    if true {
        v24 := GenS2{F0: -418}
        v25 := GenS2{F0: 127}
        v2 -= 1
        _ = v24
        _ = v25
    }
    if (9223372036854775807 <= 0) {
        v26 := GenS4{F0: true, F1: 536}
        v1 += v2
        _ = v26
    }
    v4 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
