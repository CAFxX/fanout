package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 int64; F2 bool }
type GenS3 struct { F0 int8; F1 uint64 }
func genFn0(p0 int) uint16 {
    var v7 uint64 = 499
    if (1 == 42) {
        for fi8 := 0; fi8 < 11; fi8++ {
            for fi9 := 0; fi9 < 13; fi9++ {
                var v10 uint8 = 255
                var v11 uint64 = v7
                switch "" {
                case "":
                    p0 -= 128
                case "\n\t":
                    v7 += v7
                case "a":
                    v11 += v11
                }
                _ = fi9
                _ = v10
                _ = v11
            }
            var v12 uint64 = v7
            if true {
                v7 -= v12
                v13 := GenS1{F0: 1}
                v14 := GenS1{F0: 255}
                _ = v13
                _ = v14
            }
            _ = fi8
            _ = v12
        }
        v7 += v7
        if false {
            var v15 bool = ("a" != "hello")
            v7 += v7
            v7 -= v7
            _ = v15
        }
    } else {
        v16 := GenS1{F0: 246}
        switch p0 {
        case -9223372036854775808:
            var v17 string = "a"
            _ = v17
        default:
            v18 := GenS1{F0: 7}
            _ = v18
        }
        _ = v16
    }
    v19 := GenS1{F0: 127}
    var v20 int8 = -1
    for fi21 := 0; fi21 < 1; fi21++ {
        var v22 uint16 = 2
        v7 = 4294967295
        v23 := []int{}
        _ = fi21
        _ = v22
        _ = v23
    }
    _ = p0
    _ = v19
    _ = v20
    _ = v7
    return 65535
}

func main() {
    v1 := make(map[int]bool)
    var v2 bool = false
    v3 := make(map[uint32]bool)
    var v4 int = 9223372036854775807
    var v5 uint8 = 1
    var v6 string = ("" + ("" + ("\n\t" + ("hello" + "hello"))))
    if v2 {
        if (!v2) {
            v2 = (true || (!v2))
        } else {
            switch ("foo bar" + v6) {
            case "hello":
                var v24 uint32 = 1
                _ = v24
            case "\n\t":
                v25 := []int{(v4 | (^(v4)))}
                _ = v25
            case "foo bar":
                var v26 int64 = -395
                _ = v26
            default:
                var v27 bool = (399 >= 42)
                _ = v27
            }
            var v28 uint8 = (v5 & (v5 & v5))
            v29 := []int{(v4 >> ((255) & 63))}
            _ = v28
            _ = v29
        }
        v30 := v1[v4]
        _ = v30
    }
    v5 -= 253
    v4 = v4
    if ((v6 <= "") || (0.1 == (((0.1 / 1.0) / float64(3.14159)) / 1.0))) {
        var v31 string = ((("a" + "hello") + "\n\t") + v6)
        if false {
            v32 := GenS2{F0: 65535, F1: 9223372036854775807, F2: ((1.5 / 3.14159) >= ((((3.14159 / ((0.1 * 3.14159) - (((1.5 / (0.1 - float64((3.14159 + 3.14159)))) * (1.5 / 1.0)) - (3.14159 + (-2.25 / 1.0))))) + 1.5) - (-2.25 + 1.5)) * 0.0))}
            _ = v32
        }
        _ = v31
    }
    for fi33 := 0; fi33 < 18; fi33++ {
        v5 += (v5 << ((836) & 63))
        v6 = v6
        if (v2 || true) {
            var v34 int = v4
            _ = v34
        }
        _ = fi33
    }
    var v35 uint64 = 1
    v36 := []int{(v4 & 0)}
    var v37 uint8 = 135
    switch 305 {
    case 9223372036854775807:
        v38 := v3[1]
        _ = v38
    }
    if (2147483648 <= 65535) {
        v39 := GenS1{F0: v5}
        v2 = true
        var v40 int64 = 0
        _ = v39
        _ = v40
    }
    if (4294967295 >= 18446744073709551615) {
        var v41 float32 = (float32(1e10) / 1.5)
        for fi42 := 0; fi42 < 7; fi42++ {
            for fi43 := 0; fi43 < 7; fi43++ {
                var v44 string = ("\n\t" + ((v6 + "") + v6))
                v45 := GenS1{F0: (v5 >> ((0) & 63))}
                _ = fi43
                _ = v44
                _ = v45
            }
            _ = fi42
        }
        _ = v41
    } else {
        if true {
            delete(v1, -611)
        } else {
            var v46 int8 = 100
            v47 := GenS2{F0: 233, F1: 9223372036854775807, F2: (4294967295 <= (v35 | v35))}
            _ = v46
            _ = v47
        }
        v35 -= v35
        switch ((v6 + v6) + "foo bar") {
        case "hello":
            v48 := GenS1{F0: (v5 | (v5 % (((v37 - (v37 / ((v37) | 1)))) | 1)))}
            _ = v48
        case "foo bar":
            fmt.Println(v4, v5, v35)
        default:
            v37 = (^(v5))
        }
    }
    if ((!(-911 >= int64(v5))) || true) {
        if (0 != 2) {
            v6 = ((v6 + "") + "")
            for fi49 := 0; fi49 < 4; fi49++ {
                v35 += (v35 * 42)
                _ = len(v3)
                v50 := GenS3{F0: 127, F1: (v35 << ((4294967295) & 63))}
                _ = fi49
                _ = v50
            }
        }
        var v51 uint8 = (^(v5))
        _ = v51
    } else {
        v2 = v2
        var v52 int32 = 2147483647
        v53 := GenS3{F0: 42, F1: v35}
        _ = v52
        _ = v53
    }
    v2 = (42 >= 663)
    v4 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
