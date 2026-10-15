package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 int32; F1 uint64; F2 int8 }
func genFn0(p0 string, p1 int) uint64 {
    v5 := []int{p1, (p1 ^ p1)}
    for fi6 := 0; fi6 < 9; fi6++ {
        v7 := GenS1{F0: 1, F1: fi6}
        fmt.Println(p0, p0, p1)
        p0 = (p0 + "hello")
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 17; fi8++ {
        if ((!(774 != 4294967295)) && (!false)) {
            var v9 uint32 = 932
            _ = v9
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v5
    return 4294967295
}
func genFn1(p0 uint16, p1 uint32) int16 {
    v10 := GenS1{F0: -9223372036854775808, F1: -9223372036854775808}
    p1 += (^(p1))
    v11 := []int{1, 950, -9223372036854775808, 701}
    p1 = (p1 * (p1 >> ((65535) & 63)))
    p1 = (^(p1))
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return -448
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]string)
    v3 := make(map[string]string)
    var v4 float64 = (float64(float32(1e10)) / float64(((0.1 - float32(1.5)) + (1.5 - -2.25))))
    if false {
        var v12 string = ""
        _ = v12
    }
    v4 -= ((v4 / v4) / (v4 - (0.0 - v4)))
    if (!((false || (-9223372036854775808 == 9223372036854775807)) && (!false))) {
        v13 := []int{2147483647, 42, -403}
        if (true || ("hello" > ("hello" + (("\n\t" + ("hello" + "\n\t")) + "foo bar")))) {
            if true {
                v14 := GenS2{F0: -1, F1: 2, F2: 1}
                v4 += 0.1
                _ = v14
            } else {
                var v15 int = 32768
                _ = len(v3)
                var v16 uint = 0
                _ = v15
                _ = v16
            }
            var v17 float64 = ((((v4 + ((((v4 * ((0.0 - ((float64(3.14159) / v4) - 0.1)) + 3.14159)) / (v4 + (v4 * 3.14159))) + v4) + 3.14159)) * v4) - (v4 / ((((((1e10 + v4) + 0.0) + v4) - float64(((-2.25 * float32(v4)) / 1.0))) / v4) / v4))) - float64((3.14159 / (1e10 - ((3.14159 - 0.0) - -2.25)))))
            _ = v17
        }
        switch 1 {
        case 127:
            v4 -= (((v4 + v4) - v4) / float64(1e10))
        }
        _ = v13
    }
    v4 = (1.5 / (v4 + v4))
    switch "foo bar" {
    case "foo bar":
        v18 := []int{}
        _ = v18
    case "hello":
        v4 += v4
    default:
        if true {
            for fi19 := 0; fi19 < 15; fi19++ {
                v20 := GenS1{F0: 9223372036854775807, F1: (-(fi19))}
                _ = fi19
                _ = v20
            }
            switch 42 {
            case 42:
                fmt.Println(v4)
            default:
                v21 := GenS1{F0: -9223372036854775808, F1: 128}
                _ = v21
            }
        }
    }
    v4 += v4
    switch ("\n\t" + "\n\t") {
    case "a":
        var v22 string = ""
        _ = v22
    case "\n\t":
        for fi23 := 0; fi23 < 16; fi23++ {
            switch "" {
            case "foo bar":
                var v24 float32 = 3.14159
                _ = v24
            case "a":
                var v25 bool = (!true)
                _ = v25
            default:
                v26 := v3["foo bar"]
                _ = v26
            }
            var v27 string = ""
            v4 = (3.14159 - v4)
            _ = fi23
            _ = v27
        }
    }
    for fi28 := 0; fi28 < 15; fi28++ {
        v29 := v1["a"]
        _ = fi28
        _ = v29
    }
    v3[""] = ("\n\t" + "")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
