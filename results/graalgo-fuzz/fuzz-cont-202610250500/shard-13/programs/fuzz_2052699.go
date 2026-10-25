package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 uint8; F1 int64; F2 int8 }

func main() {
    v1 := []int{}
    var v2 int8 = -1
    v3 := make(map[int]int)
    var v4 int16 = 42
    v5 := make(map[int]string)
    var v6 int16 = 0
    if false {
        switch 127 {
        case 59:
            v2 -= (v2 << ((2) & 63))
        case 232:
            v6 = (v4 % (((^(v6))) | 1))
        case 127:
            if true {
                v6 = (v4 & 0)
                v7 := GenS1{F0: "foo bar", F1: -857}
                v8 := GenS2{F0: "", F1: 2}
                _ = v7
                _ = v8
            }
        }
        if false {
            var v9 bool = true
            v10 := GenS2{F0: ("hello" + "foo bar"), F1: 345}
            _ = v10
            _ = v9
        } else {
            var v11 bool = (false && (-2.25 < 3.14159))
            if v11 {
                for fi12 := 0; fi12 < 13; fi12++ {
                    v13 := v5[fi12]
                    for fi14 := 0; fi14 < 11; fi14++ {
                        v15 := GenS2{F0: ("hello" + v13), F1: 4294967295}
                        var v16 uint8 = 112
                        var v17 int64 = 9223372036854775807
                        _ = fi14
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                    _ = fi12
                    _ = v13
                }
                var v18 bool = (((v4 - (^(v6))) >= (v6 << ((398) & 63))) || v11)
                _ = v18
            } else {
                var v19 int = 0
                v11 = (-2.25 == (0.0 * 1e10))
                fmt.Println(v6, v19)
                _ = v19
            }
            delete(v3, -9223372036854775808)
            _ = v11
        }
    } else {
        switch -2147483648 {
        case 4294967295:
            for fi20 := 0; fi20 < 18; fi20++ {
                _ = len(v1)
                v6 = 131
                v21 := GenS3{F0: 1000, F1: 575}
                _ = fi20
                _ = v21
            }
        case 2147483647:
            var v22 bool = true
            _ = v22
        }
    }
    var v23 string = (("a" + "a") + "hello")
    if true {
        for fi24 := 0; fi24 < 16; fi24++ {
            var v25 uint64 = 0
            switch ("a" + v23) {
            case "":
                v25 = (v25 % ((4294967295) | 1))
            case "hello":
                var v26 uint = 83
                _ = v26
            }
            var v27 uint64 = (v25 * 4294967295)
            _ = fi24
            _ = v25
            _ = v27
        }
    } else {
        var v28 int64 = -1
        var v29 int8 = 0
        _ = v28
        _ = v29
    }
    var v30 uint32 = 4294967295
    v31 := GenS3{F0: (-(v6)), F1: 255}
    switch v23 {
    case "a":
        v1 = append(v1, 32767)
    default:
        for fi32 := 0; fi32 < 8; fi32++ {
            switch int(v30) {
            case 465:
                v6 -= v4
            case 127:
                fmt.Println(fi32)
            case 650:
                if (true || false) {
                    fmt.Println(v2, v23, fi32)
                    switch v23 {
                    case "a":
                        v2 = 127
                    case "\n\t":
                        v33 := GenS1{F0: v23, F1: 0}
                        _ = v33
                    case "hello":
                        for fi34 := 0; fi34 < 20; fi34++ {
                            v35 := v5[32767]
                            _ = fi34
                            _ = v35
                        }
                    default:
                        v2 += -38
                    }
                    if false {
                        v2 = 1
                    } else {
                        if len(v1) > 0 {
                            v1[0] = 9223372036854775807
                        }
                        var v36 uint64 = 1
                        _ = v36
                    }
                } else {
                    v1 = append(v1, (fi32 + fi32))
                    var v37 int32 = -2147483648
                    _ = v37
                }
            default:
                v5[(fi32 & 2147483647)] = v23
            }
            v2 += (v2 * 100)
            _ = fi32
        }
    }
    fmt.Println(v6, v30, v23)
    if true {
        var v38 uint64 = 42
        _ = v38
    }
    v39 := GenS4{F0: 0, F1: -9223372036854775808, F2: v2}
    fmt.Println(v30, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v30
    _ = v31
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
