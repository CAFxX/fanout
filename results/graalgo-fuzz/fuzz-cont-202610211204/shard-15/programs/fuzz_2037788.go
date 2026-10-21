package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 int16; F1 uint64; F2 bool }
type GenS3 struct { F0 uint16; F1 uint32; F2 int32 }
type GenS4 struct { F0 int8; F1 string; F2 int32 }
func genFn0(p0 int32, p1 int) uint64 {
    var v4 float32 = (-2.25 - 1.5)
    fmt.Println(v4)
    _ = p0
    _ = p1
    _ = v4
    return 4294967295
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]int)
    v3 := make(map[string]int)
    var v5 uint32 = 1
    for fi6 := 0; fi6 < 9; fi6++ {
        switch "foo bar" {
        case "hello":
            v5 = 0
        case "a":
            for fi7 := 0; fi7 < 17; fi7++ {
                fmt.Println(v5)
                v8 := GenS1{F0: 1, F1: ("hello" + "\n\t")}
                v5 -= v5
                _ = fi7
                _ = v8
            }
        default:
            v9 := v3[(("a" + "\n\t") + "\n\t")]
            _ = v9
        }
        var v10 float64 = 3.14159
        v10 -= float64((0.0 - (0.0 - -2.25)))
        _ = fi6
        _ = v10
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        var v12 float32 = -2.25
        var v13 uint16 = 671
        for fi14 := 0; fi14 < 20; fi14++ {
            switch fi14 {
            case 127:
                v12 = ((v12 + (v12 / -2.25)) / 1e10)
            case 4294967295:
                v15 := GenS1{F0: 1, F1: "hello"}
                _ = v15
            default:
                v5 -= (^(v5))
            }
            _ = fi14
        }
        _ = fi11
        _ = v12
        _ = v13
    }
    delete(v3, "a")
    var v16 int32 = 0
    for fi17 := 0; fi17 < 19; fi17++ {
        var v18 uint8 = 138
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 8; fi19++ {
        v16 = v16
        _ = fi19
    }
    if (!(!false)) {
        switch -950 {
        case 32768:
            v5 = (v5 ^ 2)
        default:
            fmt.Println(v5, v16)
        }
        if false {
            v16 -= v16
            fmt.Println(v5)
            v5 = v5
        } else {
            for fi20 := 0; fi20 < 7; fi20++ {
                v16 -= (v16 + (v16 << ((4294967295) & 63)))
                fmt.Println(fi20, v16, fi20)
                _ = fi20
            }
            fmt.Println(v5)
            fmt.Println(v16, v16)
        }
        for fi21 := 0; fi21 < 8; fi21++ {
            var v22 int64 = 205
            _ = fi21
            _ = v22
        }
    }
    v23 := []int{72, 4294967295}
    v24 := GenS2{F0: -1, F1: 0, F2: ((127 > uint8(v5)) && (0 != -1))}
    for fi25 := 0; fi25 < 6; fi25++ {
        var v26 bool = false
        _ = fi25
        _ = v26
    }
    if false {
        switch -1 {
        case 2147483647:
            if len(v23) >= 2 {
                v23 = v23[1:len(v23)]
            }
        case 127:
            fmt.Println(v5, v16)
        case 33:
            var v27 uint32 = v5
            _ = v27
        default:
            for fi28 := 0; fi28 < 11; fi28++ {
                if (!true) {
                    if false {
                        var v29 int32 = 1
                        switch (("" + "hello") + (("hello" + "foo bar") + "")) {
                        case "\n\t":
                            switch (fi28 / ((fi28) | 1)) {
                            case 134:
                                var v30 int8 = 42
                                _ = v30
                            case 2147483647:
                                if len(v23) > 0 {
                                    v23[0] = fi28
                                }
                            case 1:
                                for fi31 := 0; fi31 < 2; fi31++ {
                                    v16 = v29
                                    var v32 int32 = (v16 ^ (v29 ^ v16))
                                    v33 := GenS3{F0: 255, F1: (^(v5)), F2: int32(fi31)}
                                    _ = fi31
                                    _ = v32
                                    _ = v33
                                }
                            }
                        }
                        if len(v23) >= 2 {
                            v23 = v23[1:len(v23)]
                        }
                        _ = v29
                    }
                } else {
                    v34 := GenS1{F0: 65535, F1: "\n\t"}
                    _ = v34
                }
                v35 := GenS4{F0: 0, F1: "foo bar", F2: -2147483648}
                _ = fi28
                _ = v35
            }
        }
    } else {
        fmt.Println(v5, v16)
    }
    v5 = v5
    switch 32768 {
    case 127:
        var v36 int16 = 42
        _ = v36
    }
    if true {
        var v37 bool = ("foo bar" >= "\n\t")
        _ = v37
    }
    var v38 uint64 = 731
    var v39 uint16 = 255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v38
    _ = v39
    _ = v5
}
