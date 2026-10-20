package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int8; F1 uint16; F2 int }
type GenS5 struct { F0 int64; F1 string; F2 int64 }
func genFn0(p0 uint32, p1 bool, p2 int64) int64 {
    v7 := []int{}
    p1 = false
    var v8 bool = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (p2 >> ((318) & 63))
}

func main() {
    v1 := []uint32{}
    var v2 uint = 756
    v3 := []int{791, 982}
    v4 := []string{}
    v5 := []string{"\n\t", "hello"}
    v6 := make(map[string]int)
    fmt.Println(v2)
    if (1 <= -1) {
        _ = len(v3)
        switch "\n\t" {
        case "a":
            _ = len(v3)
        case "\n\t":
            var v9 int64 = -662
            _ = v9
        }
    } else {
        var v10 float64 = (0.1 / (-2.25 * 3.14159))
        v11 := GenS1{F0: 836, F1: 9223372036854775807}
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        fmt.Println(v2, fi12)
        switch ("a" + "\n\t") {
        case "hello":
            if false {
                fmt.Println(fi12)
                v2 = v2
            }
        }
        _ = fi12
    }
    fmt.Println(v2)
    var v13 uint32 = 0
    for fi14 := 0; fi14 < 11; fi14++ {
        if false {
            v15 := GenS2{F0: 352, F1: true}
            fmt.Println(fi14, v13, v2)
            v16 := GenS3{F0: 18446744073709551615}
            _ = v15
            _ = v16
        } else {
            v17 := GenS1{F0: 0, F1: -848}
            fmt.Println(v2, v13)
            _ = v17
        }
        _ = fi14
    }
    var v18 uint8 = 42
    var v19 uint64 = 42
    switch "" {
    case "":
        switch "\n\t" {
        case "":
            var v20 uint32 = (v13 % ((v13) | 1))
            _ = v20
        }
    case "\n\t":
        if (!true) {
            fmt.Println(v2, v2, v19)
            v18 -= v18
        } else {
            if false {
                v21 := GenS2{F0: uint16(v2), F1: true}
                var v22 uint8 = v18
                switch 9223372036854775807 {
                case -875:
                    v23 := GenS1{F0: 42, F1: 42}
                    _ = v23
                case -2147483648:
                    var v24 int16 = -1
                    _ = v24
                }
                _ = v21
                _ = v22
            } else {
                var v25 string = "foo bar"
                v26 := GenS4{F0: 42, F1: 2, F2: 42}
                v6[v25] = 127
                _ = v25
                _ = v26
            }
            var v27 uint = uint(v19)
            _ = v27
        }
    }
    delete(v6, "a")
    if false {
        if (0 >= (v2 * v2)) {
            v28 := GenS2{F0: 42, F1: ((!(true || true)) || false)}
            v2 = (v2 >> (((^(v2))) & 63))
            switch ("hello" + ("foo bar" + "")) {
            case "":
                if false {
                    for fi29 := 0; fi29 < 9; fi29++ {
                        if true {
                            var v30 int64 = 1
                            _ = v30
                        } else {
                            v2 += uint(v18)
                        }
                        _ = len(v6)
                        _ = fi29
                    }
                } else {
                    switch 4294967295 {
                    case 1:
                        v13 = (v13 / (((v13 & uint32(v19))) | 1))
                    case -838:
                        var v31 string = "\n\t"
                        _ = v31
                    }
                    v32 := GenS2{F0: 32767, F1: (true || (!(772 >= 470)))}
                    _ = v32
                }
            case "hello":
                var v33 uint64 = (v19 >> (((v2 * 4294967295)) & 63))
                _ = v33
            }
            _ = v28
        }
        switch 0 {
        case 552:
            v34 := GenS3{F0: v19}
            _ = v34
        }
        var v35 uint16 = 65535
        _ = v35
    } else {
        fmt.Println(v19, v13, v19)
        v18 += v18
        _ = len(v4)
    }
    for fi36 := 0; fi36 < 14; fi36++ {
        var v37 float32 = ((1e10 / (1e10 * (1e10 - (1.5 + 1.5)))) + (((float32(0.0) + ((0.0 / 1.5) / (1.5 + float32(((((-2.25 - -2.25) - 0.0) - (-2.25 / 0.1)) / (-2.25 + ((float64(0.0) - ((0.0 - 0.1) - 0.1)) - (0.0 + float64(float32(0.1)))))))))) + (3.14159 / float32(1.5))) - 0.1))
        v2 = (v2 % (((v2 * (v2 >> ((v2) & 63)))) | 1))
        v38 := GenS5{F0: 9223372036854775807, F1: "\n\t", F2: 390}
        _ = fi36
        _ = v37
        _ = v38
    }
    for fi39 := 0; fi39 < 9; fi39++ {
        var v40 bool = (false && ((0 >= 1) && (true && true)))
        switch ("foo bar" + "") {
        case "a":
            var v41 uint16 = 178
            _ = v41
        default:
            if len(v3) > 0 {
                v3[0] = fi39
            }
        }
        var v42 int16 = -32768
        _ = fi39
        _ = v40
        _ = v42
    }
    switch 2147483647 {
    case -38:
        v19 -= 42
    case 32768:
        v19 = 18446744073709551615
    case 106:
        if (true && true) {
            if false {
                v1 = append(v1, (v13 - v13))
            } else {
                v13 += 2
                v43 := GenS3{F0: uint64(v18)}
                var v44 uint = 2
                _ = v43
                _ = v44
            }
            var v45 bool = (true || (42 < 2))
            for fi46 := 0; fi46 < 13; fi46++ {
                var v47 uint8 = 2
                if v45 {
                    fmt.Println(v2, v45)
                } else {
                    fmt.Println(fi46, fi46, v18)
                    switch "\n\t" {
                    case "\n\t":
                        var v48 uint = 439
                        _ = v48
                    case "hello":
                        var v49 uint8 = 42
                        _ = v49
                    case "foo bar":
                        var v50 int64 = genFn0(1, (!(-1 != -9223372036854775808)), genFn0(v13, v45, genFn0((v13 & (v13 & (v13 | 4294967295))), (0 < -1), 1)))
                        _ = v50
                    }
                }
                _ = len(v1)
                _ = fi46
                _ = v47
            }
            _ = v45
        }
    }
    switch 1 {
    case -9223372036854775808:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        var v51 int = 42
        _ = v51
    }
    v52 := GenS1{F0: 255, F1: 9223372036854775807}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v6
}
