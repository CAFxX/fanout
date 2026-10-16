package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
func genFn0() int64 {
    var v6 int64 = -9223372036854775808
    switch 0 {
    case 4294967295:
        fmt.Println(v6)
    case 474:
        switch 0 {
        case 32767:
            v6 = (v6 + v6)
        case 76:
            var v7 uint16 = 0
            _ = v7
        case 32768:
            var v8 float64 = -2.25
            _ = v8
        }
    case 575:
        var v9 int = -2147483648
        _ = v9
    }
    var v10 int32 = 0
    for fi11 := 0; fi11 < 18; fi11++ {
        if true {
            v12 := []int{fi11, fi11, fi11, -936}
            fmt.Println(v10, v6)
            _ = v12
        }
        _ = fi11
    }
    var v13 float64 = (1e10 + 0.0)
    _ = v10
    _ = v13
    _ = v6
    return v6
}

func main() {
    v1 := make(map[string]bool)
    v2 := []string{("a" + "\n\t")}
    var v3 uint = 755
    var v4 float32 = (1.5 * 1.5)
    var v5 uint8 = 94
    for fi14 := 0; fi14 < 13; fi14++ {
        v4 += (v4 / v4)
        var v15 int8 = 127
        v2 = append(v2, ("foo bar" + "a"))
        _ = fi14
        _ = v15
    }
    if true {
        v4 -= (-2.25 + v4)
    }
    for fi16 := 0; fi16 < 5; fi16++ {
        v4 = v4
        var v17 int = (fi16 / ((309) | 1))
        for fi18 := 0; fi18 < 2; fi18++ {
            var v19 int32 = 1
            v20 := GenS1{F0: 0, F1: 32768}
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = fi16
        _ = v17
    }
    fmt.Println(v4)
    if (255 == v5) {
        fmt.Println(v5, v3)
    } else {
        for fi21 := 0; fi21 < 14; fi21++ {
            fmt.Println(v4, v5, v4)
            v22 := GenS1{F0: 2147483648, F1: (fi21 / ((-9223372036854775808) | 1))}
            _ = fi21
            _ = v22
        }
    }
    var v23 bool = (705 == 255)
    if ((v4 + (v4 - 0.0)) >= v4) {
        switch 0 {
        case 9223372036854775807:
            switch ("foo bar" + ("hello" + "foo bar")) {
            case "\n\t":
                var v24 int32 = 1
                _ = v24
            case "hello":
                v5 += (v5 * uint8(v3))
            }
        case -320:
            v4 = -2.25
        case 42:
            for fi25 := 0; fi25 < 5; fi25++ {
                v26 := v1[""]
                _ = fi25
                _ = v26
            }
        default:
            v2 = append(v2, "foo bar")
        }
        v23 = true
    } else {
        v27 := GenS1{F0: 42, F1: -9223372036854775808}
        _ = v27
    }
    var v28 float32 = v4
    v28 += ((1e10 / 1.5) + v28)
    for fi29 := 0; fi29 < 17; fi29++ {
        v4 = float32(-2.25)
        _ = fi29
    }
    v30 := v1["hello"]
    v2 = append(v2, ("a" + ""))
    v23 = (v23 || v23)
    var v31 int32 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
