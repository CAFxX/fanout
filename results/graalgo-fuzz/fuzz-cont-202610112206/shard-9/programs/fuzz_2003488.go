package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 uint32 }
func genFn0(p0 int64, p1 uint64) uint32 {
    var v4 uint8 = 2
    v5 := []int{}
    if false {
        p1 = 241
        var v6 int64 = 258
        _ = v6
    }
    v7 := GenS1{F0: 0, F1: ("\n\t" + "hello"), F2: uint32(p1)}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v7
    return 65535
}
func genFn1(p0 uint32, p1 int8) int32 {
    p1 += p1
    var v8 float64 = 0.0
    for fi9 := 0; fi9 < 11; fi9++ {
        var v10 uint64 = 45
        var v11 uint = 1
        v10 = (v10 * v10)
        _ = fi9
        _ = v10
        _ = v11
    }
    fmt.Println(v8)
    for fi12 := 0; fi12 < 12; fi12++ {
        if false {
            var v13 bool = false
            v8 = (v8 + v8)
            _ = v13
        }
        v8 = 0.0
        v14 := []int{fi12, (fi12 | (^(fi12))), (fi12 - fi12), fi12}
        _ = fi12
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v8
    return 35
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int32 = -1
    v3 := []int{4294967295, 42, -9223372036854775808}
    var v15 int16 = -32768
    switch 2147483647 {
    case -997:
        if ("\n\t" != ("" + ("a" + ""))) {
            v15 -= (^(v15))
            var v16 float64 = (((-2.25 * (0.1 - 1e10)) + (3.14159 + ((-2.25 / float64(3.14159)) / 1.0))) * 0.0)
            v15 = int16(v2)
            _ = v16
        }
    }
    if ((0.0 * 0.1) > 3.14159) {
        for fi17 := 0; fi17 < 2; fi17++ {
            v18 := GenS1{F0: 65535, F1: "foo bar", F2: 2147483647}
            var v19 uint16 = 255
            fmt.Println(v19)
            _ = fi17
            _ = v18
            _ = v19
        }
        var v20 uint32 = 694
        _ = v20
    } else {
        fmt.Println(v2)
    }
    if (521 <= 299) {
        v21 := GenS1{F0: 4294967295, F1: ("a" + ("hello" + ("hello" + ("\n\t" + ("" + ("foo bar" + (("\n\t" + "\n\t") + "a"))))))), F2: 42}
        v22 := GenS1{F0: 65535, F1: "", F2: genFn0(-9223372036854775808, 0)}
        _ = v21
        _ = v22
    } else {
        v15 = (v15 % ((0) | 1))
    }
    for fi23 := 0; fi23 < 18; fi23++ {
        for fi24 := 0; fi24 < 4; fi24++ {
            v2 = (v2 >> ((2) & 63))
            v25 := GenS1{F0: 42, F1: ("\n\t" + "a"), F2: 622}
            var v26 uint = 504
            _ = fi24
            _ = v25
            _ = v26
        }
        var v27 uint16 = 0
        var v28 string = "hello"
        _ = fi23
        _ = v27
        _ = v28
    }
    v3 = append(v3, 9223372036854775807)
    if true {
        if false {
            v29 := GenS1{F0: 820, F1: "\n\t", F2: 2147483648}
            _ = v29
        } else {
            var v30 int64 = 9223372036854775807
            _ = v30
        }
    }
    for fi31 := 0; fi31 < 20; fi31++ {
        v32 := v1[65535]
        v32 = (579 != 122)
        switch ("hello" + "\n\t") {
        case "a":
            for fi33 := 0; fi33 < 1; fi33++ {
                var v34 uint8 = 42
                _ = fi33
                _ = v34
            }
        case "\n\t":
            v32 = v32
        default:
            fmt.Println(v2)
        }
        _ = fi31
        _ = v32
    }
    v2 -= (v2 % ((-2147483648) | 1))
    var v35 int64 = 1
    v36 := v1[2147483648]
    v35 -= v35
    if (1 <= -2147483648) {
        if (v36 || (32767 < 42)) {
            for fi37 := 0; fi37 < 12; fi37++ {
                var v38 uint16 = 255
                _ = fi37
                _ = v38
            }
            fmt.Println(v2, v35, v15)
            fmt.Println(v15, v35)
        }
    } else {
        var v39 int32 = genFn1(2, 1)
        var v40 uint = 2
        _ = v39
        _ = v40
    }
    var v41 string = "\n\t"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v41
}
