package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string; F1 int16 }
type GenS4 struct { F0 string }
func genFn0(p0 uint, p1 int32, p2 int16) uint32 {
    for fi5 := 0; fi5 < 20; fi5++ {
        if false {
            v6 := GenS1{F0: -1, F1: 621}
            p1 = p1
            _ = v6
        } else {
            p2 += p2
            var v7 uint32 = 2147483648
            _ = v7
        }
        if true {
            if true {
                for fi8 := 0; fi8 < 13; fi8++ {
                    p2 += 669
                    v9 := GenS2{F0: p1}
                    _ = fi8
                    _ = v9
                }
            } else {
                var v10 bool = (4294967295 < 2)
                fmt.Println(fi5, p0, v10)
                _ = v10
            }
        }
        p2 -= -77
        _ = fi5
    }
    if false {
        fmt.Println(p1, p1, p2)
    }
    var v11 uint32 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return v11
}

func main() {
    var v1 string = "hello"
    var v2 uint = 383
    var v3 float64 = (1.5 * (0.1 + (3.14159 / (1e10 + (((1.5 * 1.5) - 1.5) - (3.14159 / 1.0))))))
    var v4 string = "foo bar"
    var v12 uint16 = 242
    v13 := []int{2147483647}
    var v14 uint = 18446744073709551615
    v15 := GenS3{F0: v1, F1: 0}
    switch v4 {
    case "hello":
        _ = len(v13)
    case "":
        if (false && (true || true)) {
            var v16 uint16 = 74
            _ = v16
        }
    }
    var v17 int16 = 0
    v18 := GenS4{F0: v4}
    v17 -= (v17 >> (((^(v14))) & 63))
    fmt.Println(v17, v12)
    var v19 uint = v14
    switch 128 {
    case 790:
        switch "a" {
        case "a":
            v17 = 440
        case "\n\t":
            var v20 uint64 = 604
            _ = v20
        }
    default:
        fmt.Println(v4, v3, v19)
    }
    v3 += v3
    v19 += v2
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    var v21 float32 = float32(3.14159)
    fmt.Println(v3, v2)
    v17 -= v17
    for fi22 := 0; fi22 < 4; fi22++ {
        if false {
            for fi23 := 0; fi23 < 19; fi23++ {
                if (fi22 == fi22) {
                    fmt.Println(v14, fi23)
                    var v24 int16 = v17
                    var v25 bool = true
                    _ = v24
                    _ = v25
                } else {
                    var v26 int16 = v17
                    if true {
                        switch (-(fi22)) {
                        case 4294967295:
                            var v27 int16 = (v17 >> ((v2) & 63))
                            _ = v27
                        case -2147483648:
                            if (("a" + "") < (v4 + v4)) {
                                fmt.Println(fi23, v17, v26)
                                var v28 int8 = 42
                                v29 := GenS1{F0: 1, F1: 2}
                                _ = v28
                                _ = v29
                            }
                        case 9223372036854775807:
                            switch (v4 + "foo bar") {
                            case "a":
                                if true {
                                    var v30 int32 = 1
                                    _ = v30
                                }
                            }
                        }
                    } else {
                        v1 = (v4 + "")
                        var v31 int8 = 6
                        if false {
                            v32 := GenS2{F0: -1}
                            _ = v32
                        }
                        _ = v31
                    }
                    _ = v26
                }
                v33 := GenS3{F0: ("foo bar" + ("a" + v4)), F1: v17}
                v17 -= v17
                _ = fi23
                _ = v33
            }
            if (!true) {
                v3 = 0.0
            }
            switch (v4 + v1) {
            case "hello":
                if false {
                    var v34 int32 = 0
                    _ = v34
                }
            case "foo bar":
                v35 := GenS3{F0: "a", F1: v17}
                _ = v35
            }
        } else {
            switch v1 {
            case "\n\t":
                fmt.Println(v4, v21, v3)
            case "a":
                var v36 float64 = (float64(1.5) + ((v3 - 0.1) - 1.5))
                _ = v36
            case "":
                v13 = append(v13, (-(fi22)))
            default:
                var v37 uint64 = 42
                _ = v37
            }
            if true {
                v12 = (^(v12))
                v3 = v3
                var v38 int32 = -2147483648
                _ = v38
            }
        }
        v1 = "a"
        _ = fi22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
