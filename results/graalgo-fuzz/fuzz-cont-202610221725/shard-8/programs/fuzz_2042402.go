package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int64; F1 int32; F2 uint64 }
type GenS7 struct { F0 uint8 }
func genFn0(p0 uint32, p1 uint32, p2 int) bool {
    for fi6 := 0; fi6 < 15; fi6++ {
        if true {
            v7 := []int{}
            p2 = fi6
            _ = v7
        } else {
            v8 := GenS1{F0: 1, F1: -1}
            fmt.Println(fi6)
            _ = v8
        }
        _ = fi6
    }
    if (true && (true && false)) {
        v9 := GenS2{F0: 2, F1: 1}
        var v10 uint16 = 0
        for fi11 := 0; fi11 < 2; fi11++ {
            switch "a" {
            case "":
                var v12 uint64 = 18446744073709551615
                _ = v12
            case "foo bar":
                v13 := []int{9223372036854775807}
                _ = v13
            case "a":
                v14 := GenS3{F0: 42}
                _ = v14
            }
            var v15 int = fi11
            _ = fi11
            _ = v15
        }
        _ = v10
        _ = v9
    } else {
        fmt.Println(p2)
        fmt.Println(p2, p0, p0)
    }
    v16 := []int{p2, p2, (p2 + p2), p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return false
}
func genFn1(p0 uint8) string {
    p0 = p0
    v17 := GenS4{F0: (("hello" + ("foo bar" + "a")) + "\n\t"), F1: (p0 * p0)}
    for fi18 := 0; fi18 < 15; fi18++ {
        p0 = p0
        v19 := []int{-909, (fi18 ^ 128)}
        _ = fi18
        _ = v19
    }
    var v20 uint8 = 0
    _ = p0
    _ = v17
    _ = v20
    return ("hello" + "foo bar")
}

func main() {
    v1 := make(map[int]uint32)
    var v2 string = "foo bar"
    var v3 int64 = -1
    var v4 uint = 0
    var v5 bool = false
    var v21 int64 = (v3 + (v3 | (^(v3))))
    var v22 int8 = 100
    for fi23 := 0; fi23 < 19; fi23++ {
        if (false && false) {
            fmt.Println(fi23)
        } else {
            var v24 float64 = (1.5 / (0.1 * float64(0.1)))
            _ = v24
        }
        v2 = ""
        _ = fi23
    }
    var v25 bool = true
    for fi26 := 0; fi26 < 17; fi26++ {
        for fi27 := 0; fi27 < 2; fi27++ {
            var v28 string = v2
            v29 := v1[(fi26 & (^(fi26)))]
            if (1 > 641) {
                var v30 uint8 = 172
                switch v2 {
                case "foo bar":
                    var v31 float64 = 3.14159
                    _ = v31
                }
                _ = v30
            }
            _ = fi27
            _ = v28
            _ = v29
        }
        _ = fi26
    }
    var v32 int64 = v3
    delete(v1, -2147483648)
    var v33 int16 = -32768
    switch "\n\t" {
    case "\n\t":
        for fi34 := 0; fi34 < 3; fi34++ {
            var v35 int32 = -603
            _ = fi34
            _ = v35
        }
    case "":
        fmt.Println(v22, v5)
    case "hello":
        delete(v1, 32767)
    }
    fmt.Println(v2, v32)
    var v36 float32 = 0.0
    for fi37 := 0; fi37 < 10; fi37++ {
        fmt.Println(v2, v25)
        fmt.Println(v3)
        if v5 {
            if (!(-1 >= 42)) {
                var v38 uint8 = 255
                _ = v38
            } else {
                for fi39 := 0; fi39 < 10; fi39++ {
                    v40 := GenS5{F0: 42}
                    _ = fi39
                    _ = v40
                }
                switch "a" {
                case "a":
                    for fi41 := 0; fi41 < 16; fi41++ {
                        v42 := GenS3{F0: 2}
                        _ = fi41
                        _ = v42
                    }
                case "":
                    fmt.Println(v22)
                default:
                    fmt.Println(v4)
                }
            }
            var v43 int16 = -829
            _ = v43
        }
        _ = fi37
    }
    v44 := GenS6{F0: (v21 ^ (v3 - 9223372036854775807)), F1: 0, F2: 18446744073709551615}
    var v45 int = 533
    switch -1 {
    case 32768:
        if v5 {
            fmt.Println(v5, v22, v22)
        } else {
            for fi46 := 0; fi46 < 7; fi46++ {
                var v47 int = -9223372036854775808
                switch v2 {
                case "hello":
                    for fi48 := 0; fi48 < 18; fi48++ {
                        v45 = int(v4)
                        _ = len(v1)
                        v22 -= (v22 - 100)
                        var v49 float64 = ((1.5 * 1e10) + 0.0)
                        if ((fi48 + fi46) == (v47 ^ 32768)) {
                            v25 = (203 != 56)
                            var v50 float32 = (float32(1e10) * (v36 - (1e10 - v36)))
                            fmt.Println(fi46, v21)
                            fmt.Println(v25, v3)
                            _ = v50
                        }
                        _ = fi48
                        _ = v49
                    }
                case "\n\t":
                    switch v2 {
                    case "\n\t":
                        v51 := GenS7{F0: 255}
                        _ = v51
                    case "foo bar":
                        var v52 float64 = 3.14159
                        _ = v52
                    case "":
                        v53 := []int{(v45 / ((v45) | 1))}
                        _ = v53
                    }
                }
                _ = fi46
                _ = v47
            }
            if v5 {
                if v5 {
                    v45 -= 127
                    var v54 bool = false
                    var v55 float32 = float32((0.0 / -2.25))
                    _ = v54
                    _ = v55
                }
            }
        }
    case -825:
        fmt.Println(v36)
    case 9223372036854775807:
        var v56 uint = (v4 & v4)
        _ = v56
    }
    v57 := GenS4{F0: (("foo bar" + "a") + (("hello" + ("foo bar" + v2)) + (v2 + v2))), F1: 127}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v32
    _ = v33
    _ = v36
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v57
}
