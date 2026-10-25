package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 int16; F1 int16; F2 int8 }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 uint64; F1 string; F2 uint64 }
func genFn0(p0 uint8) bool {
    if true {
        for fi6 := 0; fi6 < 4; fi6++ {
            p0 = 128
            switch ((("foo bar" + "") + ("foo bar" + "\n\t")) + "hello") {
            case "\n\t":
                p0 -= (p0 << ((2) & 63))
            case "a":
                p0 = p0
            default:
                p0 = p0
            }
            v7 := GenS1{F0: -9223372036854775808, F1: true}
            _ = fi6
            _ = v7
        }
        v8 := []int{-2147483648}
        v8 = append(v8, 1)
        _ = v8
    } else {
        v9 := []int{32768}
        p0 = (p0 * p0)
        for fi10 := 0; fi10 < 15; fi10++ {
            p0 = 42
            for fi11 := 0; fi11 < 4; fi11++ {
                var v12 int16 = 0
                v12 = -271
                _ = fi11
                _ = v12
            }
            _ = fi10
        }
        _ = v9
    }
    p0 = (p0 % ((p0) | 1))
    v13 := GenS1{F0: -9223372036854775808, F1: true}
    switch 0 {
    case 79:
        var v14 uint8 = (^(p0))
        _ = v14
    case 42:
        var v15 float64 = -2.25
        _ = v15
    case 9223372036854775807:
        var v16 int16 = 527
        _ = v16
    default:
        if (337 > 65535) {
            switch -737 {
            case 42:
                v17 := GenS2{F0: 32767, F1: 462, F2: 1}
                _ = v17
            case -506:
                var v18 float64 = 0.0
                _ = v18
            }
            switch ("hello" + ("a" + "a")) {
            case "foo bar":
                v19 := []int{}
                _ = v19
            case "hello":
                if (!(372 == 18446744073709551615)) {
                    var v20 float32 = float32((1.5 * (float64(0.1) + (3.14159 / 0.1))))
                    v20 = 1.5
                    _ = v20
                }
            case "":
                for fi21 := 0; fi21 < 3; fi21++ {
                    if false {
                        v22 := GenS1{F0: -182, F1: false}
                        p0 = p0
                        _ = v22
                    } else {
                        v23 := []int{fi21, fi21}
                        if len(v23) >= 2 {
                            v23 = v23[1:len(v23)]
                        }
                        _ = v23
                    }
                    _ = fi21
                }
            }
            v24 := []int{int(p0), -944, 32768, 4294967295}
            _ = v24
        }
    }
    _ = p0
    _ = v13
    return (false || (-548 <= -298))
}

func main() {
    v1 := make(map[string]bool)
    var v2 string = "a"
    var v3 int8 = 87
    var v4 int32 = -1
    var v5 uint = 42
    for fi25 := 0; fi25 < 16; fi25++ {
        var v26 float64 = 1.5
        _ = fi25
        _ = v26
    }
    switch 166 {
    case 4294967295:
        v27 := []int{988, 0, -355, 32767}
        _ = v27
    default:
        v1[(("a" + ("" + v2)) + v2)] = true
    }
    var v28 uint32 = 4294967295
    for fi29 := 0; fi29 < 13; fi29++ {
        var v30 uint16 = uint16(v28)
        v31 := v1["foo bar"]
        for fi32 := 0; fi32 < 19; fi32++ {
            if v31 {
                fmt.Println(v28, v5)
                v30 -= v30
            } else {
                v5 += (^(v5))
                fmt.Println(v5, v5)
                if true {
                    var v33 uint16 = (v30 / ((42) | 1))
                    v31 = v31
                    _ = len(v1)
                    _ = v33
                } else {
                    v34 := GenS3{F0: (fi32 << ((uint(v30)) & 63)), F1: v3}
                    _ = v34
                }
            }
            _ = fi32
        }
        _ = fi29
        _ = v30
        _ = v31
    }
    if true {
        v35 := v1[("hello" + "foo bar")]
        v36 := []int{4294967295, 9223372036854775807}
        fmt.Println(v35, v3)
        _ = v35
        _ = v36
    }
    v3 = 44
    fmt.Println(v5, v5, v4)
    v3 -= 42
    var v37 string = (v2 + (((((v2 + v2) + v2) + v2) + "a") + v2))
    var v38 float64 = ((1.5 * -2.25) - (1e10 + 0.0))
    for fi39 := 0; fi39 < 8; fi39++ {
        var v40 int64 = 495
        v4 += -2147483648
        v41 := GenS3{F0: fi39, F1: v3}
        _ = fi39
        _ = v40
        _ = v41
    }
    v42 := GenS4{F0: 317, F1: v37, F2: 42}
    switch ((v2 + (v37 + v2)) + (v2 + ("hello" + (((v37 + "hello") + "hello") + ("foo bar" + "a"))))) {
    case "":
        for fi43 := 0; fi43 < 4; fi43++ {
            var v44 int8 = v3
            _ = fi43
            _ = v44
        }
    }
    v3 = v3
    switch 252 {
    case 127:
        switch v2 {
        case "\n\t":
            var v45 float64 = v38
            _ = v45
        }
    case 9223372036854775807:
        v2 = v2
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v42
    _ = v5
}
