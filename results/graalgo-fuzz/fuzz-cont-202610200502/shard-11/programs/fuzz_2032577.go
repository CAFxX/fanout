package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 string; F2 int }
func genFn0() uint64 {
    for fi7 := 0; fi7 < 2; fi7++ {
        v8 := GenS1{F0: "a"}
        if (false && (((3.14159 / (3.14159 * 1.5)) > ((-2.25 / 1e10) / -2.25)) && (true && true))) {
            v9 := []int{(-(fi7))}
            var v10 uint32 = 97
            if len(v9) > 0 {
                v9[0] = fi7
            }
            _ = v10
            _ = v9
        } else {
            v11 := GenS1{F0: ("\n\t" + ("" + "a"))}
            _ = v11
        }
        for fi12 := 0; fi12 < 19; fi12++ {
            var v13 float64 = 1e10
            v13 += (3.14159 - v13)
            fmt.Println(fi7, fi7, fi12)
            _ = fi12
            _ = v13
        }
        _ = fi7
        _ = v8
    }
    for fi14 := 0; fi14 < 2; fi14++ {
        switch ("\n\t" + "\n\t") {
        case "a":
            v15 := GenS2{F0: "a", F1: "foo bar", F2: (fi14 + (-(fi14)))}
            _ = v15
        default:
            fmt.Println(fi14)
        }
        fmt.Println(fi14)
        if true {
            switch "hello" {
            case "foo bar":
                if true {
                    var v16 int32 = 762
                    var v17 float32 = 3.14159
                    _ = v16
                    _ = v17
                }
            }
            var v18 int16 = 323
            for fi19 := 0; fi19 < 6; fi19++ {
                var v20 int64 = 0
                _ = fi19
                _ = v20
            }
            _ = v18
        } else {
            v21 := GenS1{F0: "foo bar"}
            for fi22 := 0; fi22 < 1; fi22++ {
                var v23 string = "foo bar"
                v23 = v23
                v23 = v23
                v24 := GenS2{F0: (v23 + v23), F1: v23, F2: (fi22 * fi14)}
                var v25 float64 = -2.25
                _ = fi22
                _ = v23
                _ = v24
                _ = v25
            }
            _ = v21
        }
        _ = fi14
    }
    var v26 int64 = 0
    _ = v26
    return 901
}

func main() {
    var v1 string = "hello"
    var v2 int8 = -128
    var v3 uint16 = 656
    var v4 uint64 = 0
    var v5 int64 = -348
    var v6 uint = 42
    if (!false) {
        fmt.Println(v3)
    }
    switch "foo bar" {
    case "":
        v2 -= (v2 << ((v6) & 63))
    case "\n\t":
        if true {
            v27 := GenS1{F0: ("hello" + v1)}
            switch 333 {
            case 128:
                for fi28 := 0; fi28 < 14; fi28++ {
                    fmt.Println(v5, v2, v4)
                    v6 += (v6 >> ((v6) & 63))
                    v29 := []int{}
                    _ = fi28
                    _ = v29
                }
            case -1:
                var v30 uint64 = (v4 * 0)
                _ = v30
            }
            _ = v27
        }
    default:
        if (false && (false && true)) {
            var v31 uint16 = v3
            _ = v31
        }
    }
    v1 = v1
    v32 := GenS2{F0: v1, F1: "hello", F2: -282}
    for fi33 := 0; fi33 < 18; fi33++ {
        fmt.Println(v3, v4)
        var v34 uint = v6
        v35 := GenS1{F0: v1}
        _ = fi33
        _ = v34
        _ = v35
    }
    var v36 int16 = 32767
    v37 := []int{4294967295, -496}
    fmt.Println(v2, v6)
    var v38 int8 = -115
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
