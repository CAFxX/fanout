package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint16 }
type GenS3 struct { F0 int16; F1 uint8 }
type GenS4 struct { F0 bool; F1 int; F2 uint8 }
func genFn0() int16 {
    switch -612 {
    case 567:
        v5 := []int{0}
        _ = v5
    default:
        v6 := GenS1{F0: "hello", F1: false}
        _ = v6
    }
    for fi7 := 0; fi7 < 3; fi7++ {
        v8 := GenS1{F0: "foo bar", F1: false}
        var v9 int64 = 42
        var v10 uint8 = 1
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        if true {
            for fi12 := 0; fi12 < 8; fi12++ {
                var v13 uint32 = 468
                v13 += 2
                _ = fi12
                _ = v13
            }
            var v14 float32 = -2.25
            var v15 int64 = -9223372036854775808
            _ = v14
            _ = v15
        }
        _ = fi11
    }
    v16 := []int{899}
    if (-2.25 < ((1e10 - (1.5 * float64((float32(float64(float32((3.14159 - 1e10)))) - 0.0)))) - 1e10)) {
        var v17 float32 = (1e10 + (3.14159 + 0.0))
        var v18 string = "\n\t"
        _ = v17
        _ = v18
    }
    _ = v16
    return 1000
}

func main() {
    v1 := []uint32{218, 4294967295}
    var v2 float64 = (3.14159 * 1e10)
    v3 := []string{"hello", ""}
    v4 := make(map[int]int)
    for fi19 := 0; fi19 < 17; fi19++ {
        v2 += v2
        v20 := GenS1{F0: "", F1: (!false)}
        fmt.Println(fi19)
        _ = fi19
        _ = v20
    }
    var v21 int8 = -128
    var v22 bool = false
    fmt.Println(v22)
    v2 += float64(3.14159)
    v23 := GenS2{F0: 607, F1: 91, F2: 582}
    switch "" {
    case "a":
        v24 := GenS1{F0: ("" + "a"), F1: false}
        _ = v24
    case "foo bar":
        for fi25 := 0; fi25 < 3; fi25++ {
            var v26 uint32 = 2
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            for fi27 := 0; fi27 < 18; fi27++ {
                if (v2 == v2) {
                    v2 -= (v2 * v2)
                    var v28 uint64 = 0
                    var v29 string = "a"
                    if (0 >= v28) {
                        switch (v29 + "hello") {
                        case "a":
                            v26 = (^(v26))
                        case "":
                            v30 := GenS3{F0: 1, F1: 186}
                            _ = v30
                        case "hello":
                            fmt.Println(fi27, v22, v26)
                        }
                        if true {
                            if len(v3) > 0 {
                                v3[0] = (v29 + v29)
                            }
                        }
                        v31 := GenS4{F0: v22, F1: fi25, F2: 1}
                        _ = v31
                    } else {
                        if (1 <= 1) {
                            var v32 float32 = 0.1
                            v33 := GenS2{F0: (v26 & 4294967295), F1: 171, F2: 32767}
                            fmt.Println(v32, v28, v2)
                            v32 = float32(-2.25)
                            v21 = v21
                            _ = v32
                            _ = v33
                        }
                    }
                    v21 = -128
                    _ = v28
                    _ = v29
                }
                delete(v4, (fi27 % ((9223372036854775807) | 1)))
                v4[fi25] = (fi25 | 32768)
                _ = fi27
            }
            _ = fi25
            _ = v26
        }
    case "hello":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        v21 += (v21 + (-(v21)))
    }
    if false {
        v34 := GenS3{F0: 1000, F1: 255}
        v2 -= -2.25
        _ = v34
    } else {
        v2 -= float64(0.1)
        switch ("foo bar" + "") {
        case "hello":
            var v35 int = 0
            _ = v35
        }
        for fi36 := 0; fi36 < 4; fi36++ {
            var v37 bool = (false || v22)
            switch (("" + ("" + (("foo bar" + "a") + ("a" + (("\n\t" + ((("" + "") + "\n\t") + "foo bar")) + "a"))))) + ("foo bar" + "a")) {
            case "hello":
                var v38 bool = (!(!(1 >= 65535)))
                _ = v38
            default:
                var v39 uint = 86
                _ = v39
            }
            fmt.Println(v21, v21, fi36)
            _ = fi36
            _ = v37
        }
    }
    v22 = (!v22)
    for fi40 := 0; fi40 < 14; fi40++ {
        fmt.Println(v2)
        _ = fi40
    }
    var v41 int32 = -131
    var v42 int16 = 1000
    v2 = v2
    v43 := GenS1{F0: "", F1: ((!v22) || v22)}
    fmt.Println(v42, v2, v21)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
}
