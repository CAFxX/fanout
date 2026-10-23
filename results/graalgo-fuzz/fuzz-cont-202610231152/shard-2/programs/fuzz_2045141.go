package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int }
func genFn0() int8 {
    var v7 uint8 = 255
    for fi8 := 0; fi8 < 10; fi8++ {
        for fi9 := 0; fi9 < 1; fi9++ {
            var v10 int = fi8
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    for fi11 := 0; fi11 < 17; fi11++ {
        for fi12 := 0; fi12 < 4; fi12++ {
            v13 := []int{fi11}
            v14 := GenS1{F0: 42}
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = fi11
    }
    _ = v7
    return -38
}
func genFn1(p0 uint8, p1 uint16) int32 {
    p1 = (^(p1))
    var v15 int16 = 1000
    switch -433 {
    case 279:
        v15 = (v15 - v15)
    }
    v16 := GenS1{F0: 0}
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    return -2147483648
}

func main() {
    v1 := []string{("\n\t" + "foo bar"), "hello", "a"}
    v2 := make(map[int]int)
    var v3 uint64 = 42
    var v4 string = "hello"
    var v5 uint8 = 127
    v6 := make(map[int]int)
    if len(v1) > 0 {
        v1[0] = ("" + v4)
    }
    switch v4 {
    case "a":
        v5 += (^(v5))
    case "\n\t":
        v17 := GenS2{F0: (^(v3)), F1: 9223372036854775807}
        _ = v17
    case "hello":
        fmt.Println(v4, v5, v5)
    }
    v1 = append(v1, v4)
    for fi18 := 0; fi18 < 6; fi18++ {
        v19 := GenS3{F0: ("a" + (v4 + v4)), F1: 485}
        switch fi18 {
        case -2147483648:
            for fi20 := 0; fi20 < 16; fi20++ {
                var v21 uint64 = v3
                v22 := GenS4{F0: -2147483648}
                v4 = "a"
                v3 = v21
                _ = fi20
                _ = v21
                _ = v22
            }
        case 0:
            v23 := GenS2{F0: 1, F1: 110}
            _ = v23
        }
        _ = fi18
        _ = v19
    }
    v4 = ("a" + v4)
    if (v3 < v3) {
        for fi24 := 0; fi24 < 16; fi24++ {
            fmt.Println(v3)
            if ((!((4294967295 >= 255) && ((v5 - (^(v5))) < 42))) && (1 < v5)) {
                var v25 int = (fi24 >> ((285) & 63))
                var v26 float32 = 1e10
                _ = v25
                _ = v26
            } else {
                var v27 string = v4
                _ = v27
            }
            v5 -= v5
            _ = fi24
        }
    } else {
        var v28 int32 = -973
        v1 = append(v1, ("" + v4))
        var v29 float64 = 1.5
        _ = v28
        _ = v29
    }
    if false {
        if ((42 == -1) || true) {
            if ((!false) && false) {
                for fi30 := 0; fi30 < 8; fi30++ {
                    var v31 float64 = (float64(float32(((1.5 * 0.0) - (1.5 * (float64(float32(((0.1 * (1e10 * 1e10)) * ((-2.25 + (0.1 + 0.0)) * 1e10)))) + 1.5))))) + 0.0)
                    _ = fi30
                    _ = v31
                }
            } else {
                if (160 != 513) {
                    fmt.Println(v5)
                } else {
                    var v32 int16 = 1000
                    var v33 int64 = -153
                    _ = v32
                    _ = v33
                }
            }
            if ((65535 == 42) && (!false)) {
                for fi34 := 0; fi34 < 18; fi34++ {
                    v3 -= 0
                    if false {
                        var v35 int16 = 32767
                        for fi36 := 0; fi36 < 14; fi36++ {
                            if (!true) {
                                _ = len(v1)
                                if len(v1) > 0 {
                                    v1[0] = v4
                                }
                                var v37 uint64 = 0
                                _ = v37
                            }
                            _ = fi36
                        }
                        _ = len(v6)
                        _ = v35
                    }
                    for fi38 := 0; fi38 < 19; fi38++ {
                        v39 := GenS1{F0: fi34}
                        v3 -= v3
                        _ = fi38
                        _ = v39
                    }
                    _ = fi34
                }
                var v40 float64 = 3.14159
                _ = v40
            } else {
                if false {
                    delete(v6, 534)
                    v41 := v6[-2147483648]
                    var v42 int64 = 9223372036854775807
                    _ = v41
                    _ = v42
                }
            }
            for fi43 := 0; fi43 < 5; fi43++ {
                v44 := GenS3{F0: (v4 + (v4 + "foo bar")), F1: 1}
                v45 := GenS5{F0: fi43}
                var v46 float32 = ((1e10 + 0.0) - (3.14159 + 0.0))
                _ = fi43
                _ = v44
                _ = v45
                _ = v46
            }
        }
        var v47 int16 = 0
        _ = v47
    }
    v6[128] = 128
    switch (("foo bar" + "\n\t") + v4) {
    case "hello":
        v4 = ((v4 + "a") + v4)
    case "foo bar":
        for fi48 := 0; fi48 < 1; fi48++ {
            var v49 float64 = -2.25
            _ = fi48
            _ = v49
        }
    }
    switch (v4 + v4) {
    case "":
        v5 -= v5
    case "\n\t":
        for fi50 := 0; fi50 < 14; fi50++ {
            for fi51 := 0; fi51 < 19; fi51++ {
                v3 = 18446744073709551615
                _ = fi51
            }
            _ = fi50
        }
    }
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
