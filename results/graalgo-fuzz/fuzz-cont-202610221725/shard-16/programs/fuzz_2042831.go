package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16; F1 int16; F2 string }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 uint16; F1 uint8; F2 int }
type GenS6 struct { F0 int16 }
func genFn0(p0 uint64, p1 string, p2 int32) int32 {
    if true {
        if (false && (0.1 > ((((-2.25 * (float64((-2.25 / 3.14159)) * 0.0)) - ((3.14159 + (0.1 / float64(float32(0.1)))) / -2.25)) - (-2.25 - ((float64((3.14159 + (-2.25 + 3.14159))) * 0.1) - -2.25))) - 0.0))) {
            if (1e10 == 3.14159) {
                switch p1 {
                case "a":
                    fmt.Println(p0, p2, p1)
                case "hello":
                    v6 := GenS1{F0: 36, F1: -50, F2: "a"}
                    _ = v6
                case "foo bar":
                    p1 = (p1 + "a")
                default:
                    p0 += 4294967295
                }
                p2 = 0
            }
            v7 := []int{}
            var v8 int16 = 0
            _ = v7
            _ = v8
        } else {
            v9 := GenS2{F0: (^(p0))}
            for fi10 := 0; fi10 < 19; fi10++ {
                v11 := []int{127, -9223372036854775808, (-(fi10)), fi10}
                _ = fi10
                _ = v11
            }
            for fi12 := 0; fi12 < 13; fi12++ {
                v13 := []int{(fi12 >> ((65535) & 63)), fi12}
                v14 := GenS3{F0: 32767, F1: 32767, F2: p1}
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = v9
        }
        var v15 int16 = 32767
        _ = v15
    } else {
        v16 := GenS4{F0: 0, F1: true}
        fmt.Println(p0, p2, p0)
        v17 := []int{-9223372036854775808, 42, 0}
        _ = v16
        _ = v17
    }
    v18 := []int{99, 581}
    fmt.Println(p0)
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    return 0
}

func main() {
    var v1 string = ("a" + "")
    var v2 bool = true
    var v3 float64 = (0.0 + (float64(1e10) + float64(((3.14159 * 0.1) + 0.1))))
    var v4 float32 = float32(float64(3.14159))
    var v5 float64 = ((1e10 * v3) * (1.5 - v3))
    if v2 {
        if true {
            fmt.Println(v4, v3, v1)
            if (v2 || (v2 && (v1 >= v1))) {
                fmt.Println(v2, v1, v3)
                v19 := GenS4{F0: 42, F1: ((!false) || (v3 <= (v5 * v3)))}
                v20 := GenS2{F0: 1}
                _ = v19
                _ = v20
            } else {
                var v21 uint64 = 1
                _ = v21
            }
        }
        v5 = ((0.1 - v5) - 3.14159)
    } else {
        switch (v1 + ((v1 + v1) + v1)) {
        case "foo bar":
            fmt.Println(v4, v2, v5)
        case "hello":
            for fi22 := 0; fi22 < 7; fi22++ {
                var v23 uint = 60
                v24 := GenS4{F0: 1, F1: v2}
                _ = fi22
                _ = v23
                _ = v24
            }
        }
        var v25 int32 = -1
        v26 := []int{0}
        _ = v25
        _ = v26
    }
    var v27 string = ((v1 + (v1 + "a")) + (v1 + ((((v1 + (((("\n\t" + ("foo bar" + ((v1 + ("\n\t" + "\n\t")) + "a"))) + v1) + "") + v1)) + (v1 + (v1 + v1))) + v1) + v1)))
    v28 := GenS4{F0: 42, F1: (!(!true))}
    v2 = (!v2)
    var v29 uint32 = 610
    if false {
        for fi30 := 0; fi30 < 1; fi30++ {
            for fi31 := 0; fi31 < 14; fi31++ {
                v2 = v2
                v32 := GenS2{F0: 970}
                var v33 uint16 = 32767
                _ = fi31
                _ = v32
                _ = v33
            }
            for fi34 := 0; fi34 < 18; fi34++ {
                v5 += 0.1
                v3 -= v5
                _ = fi34
            }
            _ = fi30
        }
        v2 = v2
        v35 := []int{32768, 128, -255, -654}
        _ = v35
    } else {
        for fi36 := 0; fi36 < 11; fi36++ {
            var v37 int64 = -595
            if v2 {
                fmt.Println(v27, v29)
                fmt.Println(v1, v37, v5)
                if (v37 == 42) {
                    if (-46 != 100) {
                        v38 := []int{}
                        _ = v38
                    }
                    if true {
                        v3 = v5
                    }
                } else {
                    var v39 int64 = (-(v37))
                    for fi40 := 0; fi40 < 20; fi40++ {
                        var v41 int16 = 1
                        _ = fi40
                        _ = v41
                    }
                    _ = v39
                }
            }
            for fi42 := 0; fi42 < 18; fi42++ {
                fmt.Println(v27)
                fmt.Println(v2, v1)
                var v43 uint = 944
                var v44 uint64 = 737
                v45 := GenS5{F0: 1, F1: 127, F2: 42}
                fmt.Println(v2, v37, v5)
                _ = fi42
                _ = v43
                _ = v44
                _ = v45
            }
            _ = fi36
            _ = v37
        }
        v46 := GenS6{F0: 0}
        _ = v46
    }
    fmt.Println(v4)
    v1 = v1
    var v47 int = 32768
    v1 = ""
    var v48 string = v1
    v5 -= v3
    var v49 uint8 = 255
    fmt.Println(v2, v48)
    var v50 uint32 = v29
    var v51 int16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v51
}
