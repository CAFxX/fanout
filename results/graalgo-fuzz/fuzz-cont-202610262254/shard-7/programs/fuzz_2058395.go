package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 uint64; F1 int32 }
func genFn0(p0 bool, p1 uint8, p2 bool) int {
    v6 := []int{-662, 32767}
    p1 -= 153
    var v7 float32 = 0.0
    v8 := GenS1{F0: 2147483647, F1: -347}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return -879
}
func genFn1(p0 uint32, p1 int64, p2 uint8) int64 {
    if (((241 < 616) && (0 >= 4294967295)) || (3.14159 > 0.1)) {
        v9 := GenS2{F0: 0, F1: 32767}
        v10 := GenS1{F0: -542, F1: 618}
        _ = v10
        _ = v9
    }
    switch "hello" {
    case "":
        if true {
            v11 := GenS2{F0: 2147483647, F1: 2}
            fmt.Println(p2)
            _ = v11
        } else {
            if true {
                if (false && (true && true)) {
                    var v12 int = 2147483647
                    _ = v12
                }
                var v13 float64 = 0.0
                _ = v13
            }
            v14 := GenS2{F0: -1, F1: 1}
            var v15 uint16 = 318
            _ = v14
            _ = v15
        }
    case "a":
        if false {
            switch 38 {
            case -519:
                for fi16 := 0; fi16 < 9; fi16++ {
                    var v17 int16 = 1
                    _ = fi16
                    _ = v17
                }
            case 1:
                for fi18 := 0; fi18 < 18; fi18++ {
                    var v19 string = "\n\t"
                    _ = fi18
                    _ = v19
                }
            case -2147483648:
                var v20 float64 = -2.25
                _ = v20
            default:
                v21 := []int{1, 32767, -6}
                _ = v21
            }
            var v22 bool = false
            _ = v22
        }
    }
    fmt.Println(p0, p2, p1)
    v23 := GenS1{F0: -963, F1: -1}
    if (!((p1 << ((1) & 63)) == (^(p1)))) {
        var v24 bool = true
        if (genFn0((!v24), 128, true) == -9223372036854775808) {
            var v25 int64 = p1
            _ = v25
        }
        _ = v24
    } else {
        v26 := GenS3{F0: 218, F1: 42}
        for fi27 := 0; fi27 < 10; fi27++ {
            var v28 uint16 = 42
            var v29 int32 = 42
            _ = fi27
            _ = v28
            _ = v29
        }
        var v30 int16 = 0
        _ = v26
        _ = v30
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    return p1
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]int)
    v3 := []string{"", ""}
    var v4 uint8 = 127
    var v5 int16 = 804
    v31 := GenS3{F0: 0, F1: 2147483647}
    v32 := GenS1{F0: -252, F1: -32768}
    v4 -= (v4 / ((224) | 1))
    var v33 string = ("hello" + "")
    if true {
        v3 = append(v3, v33)
        if false {
            for fi34 := 0; fi34 < 2; fi34++ {
                v35 := GenS1{F0: fi34, F1: (v5 + v5)}
                v5 -= (v5 - 1000)
                if (585 != 42) {
                    var v36 int64 = -527
                    var v37 int32 = -2147483648
                    fmt.Println(v4)
                    _ = v36
                    _ = v37
                }
                _ = fi34
                _ = v35
            }
            var v38 int64 = -286
            for fi39 := 0; fi39 < 6; fi39++ {
                var v40 uint32 = 4294967295
                _ = fi39
                _ = v40
            }
            _ = v38
        }
        _ = len(v1)
    }
    switch (v33 + v33) {
    case "hello":
        v41 := GenS3{F0: 0, F1: 2147483647}
        _ = v41
    default:
        var v42 int64 = genFn1(1, -46, (v4 ^ 127))
        _ = v42
    }
    if false {
        var v43 uint = 65535
        var v44 float32 = 0.0
        _ = v43
        _ = v44
    }
    for fi45 := 0; fi45 < 8; fi45++ {
        fmt.Println(v5, v33)
        _ = fi45
    }
    v46 := GenS2{F0: 762, F1: 495}
    var v47 int32 = -2147483648
    v5 += (v5 << ((1) & 63))
    switch v33 {
    case "hello":
        for fi48 := 0; fi48 < 13; fi48++ {
            var v49 int64 = 1
            _ = fi48
            _ = v49
        }
    default:
        for fi50 := 0; fi50 < 19; fi50++ {
            for fi51 := 0; fi51 < 15; fi51++ {
                var v52 int64 = -393
                switch (-(fi50)) {
                case 4294967295:
                    switch v33 {
                    case "hello":
                        var v53 int64 = (v52 >> ((18446744073709551615) & 63))
                        _ = v53
                    }
                }
                _ = fi51
                _ = v52
            }
            _ = fi50
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v46
    _ = v47
    _ = v5
}
