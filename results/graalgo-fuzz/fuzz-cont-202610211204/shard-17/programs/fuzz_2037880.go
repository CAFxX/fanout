package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16 }
func genFn0(p0 bool) int8 {
    fmt.Println(p0)
    switch 128 {
    case 9223372036854775807:
        v6 := GenS1{F0: 1, F1: ("foo bar" + "foo bar"), F2: 0}
        _ = v6
    }
    _ = p0
    return -28
}
func genFn1(p0 uint8, p1 int64, p2 uint32) uint64 {
    var v7 uint64 = 42
    v8 := []int{}
    if true {
        var v9 float32 = float32((3.14159 / float64(float32(float64(3.14159)))))
        fmt.Println(v7, p2, p0)
        _ = v9
    } else {
        for fi10 := 0; fi10 < 15; fi10++ {
            for fi11 := 0; fi11 < 16; fi11++ {
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                if (p2 >= 2) {
                    var v12 int32 = -37
                    v8 = append(v8, fi10)
                    _ = v12
                }
                fmt.Println(fi11)
                _ = fi11
            }
            _ = fi10
        }
        var v13 int16 = 652
        _ = v13
    }
    if true {
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (^(v7))
}

func main() {
    var v1 uint32 = 0
    v2 := make(map[int]bool)
    v3 := make(map[int]string)
    v4 := make(map[int]bool)
    var v5 int16 = -1
    v14 := []int{-478, 1, -64, 9223372036854775807}
    for fi15 := 0; fi15 < 11; fi15++ {
        fmt.Println(v5)
        _ = fi15
    }
    v1 += 737
    var v16 float64 = (1e10 + 0.1)
    v17 := GenS1{F0: v5, F1: "\n\t", F2: v5}
    if ((1e10 - (1.5 - 1e10)) > 0.0) {
        if false {
            if false {
                if (!true) {
                    if true {
                        v18 := GenS2{F0: 0}
                        _ = v18
                    } else {
                        v19 := GenS3{F0: 1}
                        _ = v19
                    }
                    var v20 uint64 = 1
                    _ = v20
                } else {
                    var v21 int16 = v5
                    _ = v21
                }
            }
        } else {
            v1 = v1
            for fi22 := 0; fi22 < 2; fi22++ {
                if false {
                    for fi23 := 0; fi23 < 1; fi23++ {
                        var v24 float32 = 0.1
                        if false {
                            switch fi22 {
                            case -9223372036854775808:
                                v1 -= (v1 & (v1 | (^(v1))))
                            case -56:
                                fmt.Println(fi22)
                            default:
                                fmt.Println(v16, fi23, fi23)
                            }
                            var v25 uint32 = v1
                            v26 := GenS1{F0: (^(v5)), F1: "", F2: (v5 >> ((766) & 63))}
                            _ = v25
                            _ = v26
                        }
                        v24 += 3.14159
                        _ = fi23
                        _ = v24
                    }
                    delete(v2, int(v5))
                } else {
                    v27 := GenS2{F0: 1}
                    for fi28 := 0; fi28 < 10; fi28++ {
                        v5 = (v5 >> ((42) & 63))
                        v29 := GenS2{F0: 0}
                        v16 = float64(float32(float64(float32(3.14159))))
                        _ = fi28
                        _ = v29
                    }
                    fmt.Println(v16)
                    _ = v27
                }
                v14 = append(v14, fi22)
                _ = fi22
            }
        }
        v30 := v2[0]
        _ = v30
    }
    fmt.Println(v5, v16)
    if (genFn0(true) <= genFn0(((!true) && true))) {
        if false {
            for fi31 := 0; fi31 < 1; fi31++ {
                switch (("a" + "hello") + ("foo bar" + ("hello" + "foo bar"))) {
                case "\n\t":
                    var v32 float64 = ((v16 / -2.25) - v16)
                    _ = v32
                case "a":
                    fmt.Println(fi31, v16, v5)
                case "foo bar":
                    v5 = -32768
                default:
                    var v33 float64 = ((float64((0.1 - 3.14159)) - v16) - 3.14159)
                    _ = v33
                }
                for fi34 := 0; fi34 < 14; fi34++ {
                    fmt.Println(v16)
                    v5 = v5
                    _ = fi34
                }
                var v35 uint = 42
                _ = fi31
                _ = v35
            }
        }
    } else {
        v36 := v2[127]
        _ = v36
    }
    v37 := GenS1{F0: 32767, F1: ("\n\t" + ("hello" + (("" + "foo bar") + ""))), F2: v5}
    var v38 uint64 = 2
    if len(v14) > 0 {
        v14[0] = 751
    }
    for fi39 := 0; fi39 < 4; fi39++ {
        if false {
            v38 = v38
            v38 = (v38 << ((2) & 63))
        }
        _ = fi39
    }
    v16 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
