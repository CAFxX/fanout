package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 int64 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 int; F1 uint8; F2 uint64 }
func genFn0() bool {
    var v4 string = "\n\t"
    v5 := []int{127, 2147483647}
    switch v4 {
    case "foo bar":
        v5 = append(v5, -1)
    }
    _ = v4
    _ = v5
    return false
}

func main() {
    var v1 bool = false
    v2 := make(map[string]string)
    v3 := make(map[uint32]uint32)
    fmt.Println(v1)
    var v6 int32 = -783
    if (!(1 < -1)) {
        for fi7 := 0; fi7 < 1; fi7++ {
            if v1 {
                for fi8 := 0; fi8 < 20; fi8++ {
                    fmt.Println(v1, v1)
                    fmt.Println(fi8, v1, fi7)
                    _ = fi8
                }
                v1 = v1
            }
            if false {
                for fi9 := 0; fi9 < 16; fi9++ {
                    if v1 {
                        fmt.Println(fi9, v6)
                    }
                    v2[(("\n\t" + "foo bar") + "hello")] = "a"
                    _ = fi9
                }
            }
            v10 := v3[1]
            _ = fi7
            _ = v10
        }
        var v11 int64 = 42
        _ = v11
    }
    var v12 float32 = 0.0
    var v13 uint32 = 4294967295
    for fi14 := 0; fi14 < 2; fi14++ {
        if false {
            v15 := GenS1{F0: 1, F1: (fi14 | fi14), F2: -1}
            var v16 uint16 = 584
            v17 := GenS1{F0: 2, F1: fi14, F2: -9223372036854775808}
            _ = v15
            _ = v16
            _ = v17
        }
        v18 := GenS1{F0: 42, F1: (fi14 >> ((42) & 63)), F2: 9223372036854775807}
        if (0 == 1) {
            v1 = (v1 && (42 == 213))
            for fi19 := 0; fi19 < 5; fi19++ {
                var v20 int32 = (^(v6))
                v21 := GenS1{F0: 42, F1: -318, F2: int64(v6)}
                if true {
                    var v22 string = "foo bar"
                    if v1 {
                        var v23 bool = (42 <= 42)
                        switch "a" {
                        case "a":
                            v24 := GenS1{F0: 734, F1: fi19, F2: -415}
                            _ = v24
                        case "":
                            v6 -= v20
                        default:
                            switch 4294967295 {
                            case 127:
                                v20 += (v20 ^ v6)
                            case 2147483647:
                                v22 = "a"
                            case 658:
                                v25 := GenS1{F0: 1, F1: fi19, F2: -518}
                                _ = v25
                            }
                        }
                        _ = v23
                    }
                    fmt.Println(v12, v6)
                    _ = v22
                }
                _ = fi19
                _ = v20
                _ = v21
            }
            var v26 int64 = -1
            _ = v26
        } else {
            for fi27 := 0; fi27 < 7; fi27++ {
                var v28 uint64 = 2
                switch "" {
                case "a":
                    var v29 uint64 = 622
                    _ = v29
                case "":
                    v30 := GenS1{F0: 4294967295, F1: fi27, F2: 0}
                    _ = v30
                }
                var v31 uint8 = 127
                _ = fi27
                _ = v28
                _ = v31
            }
            switch 3 {
            case 32767:
                v32 := GenS2{F0: -1, F1: ("foo bar" + (("\n\t" + "") + ""))}
                _ = v32
            default:
                var v33 int = 0
                _ = v33
            }
        }
        _ = fi14
        _ = v18
    }
    v12 -= v12
    v34 := GenS3{F0: 9223372036854775807, F1: 61, F2: 4294967295}
    for fi35 := 0; fi35 < 12; fi35++ {
        switch (-(fi35)) {
        case 9223372036854775807:
            fmt.Println(v1, fi35, fi35)
        default:
            fmt.Println(v13)
        }
        var v36 bool = false
        v37 := []int{}
        _ = fi35
        _ = v36
        _ = v37
    }
    v38 := []int{9223372036854775807, -1, 520, -743}
    for fi39 := 0; fi39 < 8; fi39++ {
        v12 -= ((v12 - 0.0) - 1e10)
        _ = fi39
    }
    if v1 {
        for fi40 := 0; fi40 < 12; fi40++ {
            if v1 {
                fmt.Println(v12, v1)
                var v41 string = (("a" + "a") + "foo bar")
                _ = v41
            }
            v12 = 0.0
            var v42 uint16 = 255
            _ = fi40
            _ = v42
        }
        v12 = 1.5
    } else {
        _ = len(v38)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v34
    _ = v38
    _ = v6
}
