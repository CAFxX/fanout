package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 int8 }
type GenS2 struct { F0 int32; F1 string; F2 int }
func genFn0() int8 {
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 int64 = -1
        v9 := GenS1{F0: 232, F1: fi7, F2: 42}
        for fi10 := 0; fi10 < 12; fi10++ {
            switch (("\n\t" + "a") + "hello") {
            case "hello":
                v8 -= -608
            default:
                var v11 uint16 = 255
                _ = v11
            }
            switch ("" + (((("hello" + ("hello" + "")) + ("\n\t" + ("" + "foo bar"))) + "foo bar") + "\n\t")) {
            case "\n\t":
                v12 := []int{}
                _ = v12
            default:
                var v13 bool = true
                v14 := GenS1{F0: 499, F1: 32767, F2: 1}
                _ = v13
                _ = v14
            }
            v8 = v8
            _ = fi10
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    if true {
        if false {
            var v15 int16 = 1000
            switch 32767 {
            case 551:
                for fi16 := 0; fi16 < 1; fi16++ {
                    v15 -= (v15 + (v15 << ((42) & 63)))
                    v17 := []int{-238}
                    _ = fi16
                    _ = v17
                }
            }
            fmt.Println(v15)
            _ = v15
        }
        if (0 != 42) {
            var v18 string = ""
            for fi19 := 0; fi19 < 4; fi19++ {
                v20 := []int{fi19, fi19}
                v21 := GenS1{F0: 65535, F1: fi19, F2: -41}
                _ = fi19
                _ = v20
                _ = v21
            }
            v18 = v18
            _ = v18
        } else {
            v22 := []int{-2147483648, 1}
            _ = v22
        }
        var v23 int8 = 55
        _ = v23
    } else {
        var v24 bool = (0.1 == float64((1.5 - (0.0 + (1e10 / (0.1 + 0.1))))))
        switch "\n\t" {
        case "\n\t":
            v24 = false
        case "foo bar":
            var v25 uint64 = 24
            _ = v25
        case "hello":
            fmt.Println(v24)
        }
        _ = v24
    }
    for fi26 := 0; fi26 < 17; fi26++ {
        var v27 float32 = (-2.25 + 0.1)
        _ = fi26
        _ = v27
    }
    var v28 uint32 = 1
    _ = v28
    return 100
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{32767}
    var v3 int8 = -118
    v4 := []int{-207}
    v5 := make(map[uint32]int)
    v6 := []string{("hello" + "a"), "a", ""}
    fmt.Println(v3)
    for fi29 := 0; fi29 < 2; fi29++ {
        var v30 uint16 = 924
        var v31 uint64 = 4294967295
        _ = fi29
        _ = v30
        _ = v31
    }
    v3 -= v3
    fmt.Println(v3)
    v32 := v1["a"]
    var v33 uint = 287
    v34 := GenS2{F0: -1, F1: "hello", F2: -13}
    var v35 uint32 = (v32 | v32)
    var v36 uint16 = 2
    for fi37 := 0; fi37 < 18; fi37++ {
        if false {
            switch fi37 {
            case 1:
                fmt.Println(v33)
            }
        } else {
            fmt.Println(v3, fi37)
        }
        delete(v1, "a")
        _ = fi37
    }
    var v38 int16 = -1
    switch 32767 {
    case 671:
        if true {
            v39 := GenS2{F0: -2147483648, F1: "\n\t", F2: -667}
            _ = v39
        }
    default:
        v38 = 1000
    }
    var v40 float64 = 0.1
    var v41 float64 = 3.14159
    v33 = (v33 << ((uint(v35)) & 63))
    v4 = append(v4, -1)
    v42 := v1["\n\t"]
    fmt.Println(v42)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v38
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
