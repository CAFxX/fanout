package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 uint8; F1 int32 }
func genFn0() int {
    var v7 uint16 = 905
    v8 := GenS1{F0: 711, F1: 1}
    var v9 string = ""
    v10 := []int{-1, -9223372036854775808, 32768, 307}
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return -904
}
func genFn1(p0 uint32) string {
    v11 := GenS1{F0: uint64(p0), F1: 720}
    switch "a" {
    case "":
        for fi12 := 0; fi12 < 12; fi12++ {
            if (!true) {
                switch "hello" {
                case "\n\t":
                    v13 := []int{}
                    _ = v13
                case "":
                    var v14 bool = false
                    _ = v14
                }
                for fi15 := 0; fi15 < 14; fi15++ {
                    switch -2147483648 {
                    case 32767:
                        var v16 float64 = 1.5
                        _ = v16
                    case 128:
                        v17 := GenS1{F0: 1, F1: -440}
                        _ = v17
                    }
                    var v18 uint64 = 0
                    var v19 uint32 = (^(p0))
                    v19 = (p0 >> ((2) & 63))
                    _ = fi15
                    _ = v18
                    _ = v19
                }
            } else {
                if false {
                    var v20 uint64 = 18446744073709551615
                    v21 := GenS2{F0: (!false)}
                    _ = v20
                    _ = v21
                } else {
                    var v22 uint8 = 128
                    v23 := []int{genFn0()}
                    _ = v22
                    _ = v23
                }
                fmt.Println(fi12, p0)
            }
            if true {
                var v24 uint = 18446744073709551615
                _ = v24
            } else {
                v25 := []int{}
                _ = v25
            }
            if true {
                for fi26 := 0; fi26 < 6; fi26++ {
                    fmt.Println(fi26, fi26, fi26)
                    fmt.Println(fi12)
                    _ = fi26
                }
            }
            _ = fi12
        }
    case "foo bar":
        if ((float32(0.0) + 0.0) >= float32(float64(0.1))) {
            var v27 string = "foo bar"
            _ = v27
        } else {
            for fi28 := 0; fi28 < 20; fi28++ {
                if (!false) {
                    p0 = (p0 * 224)
                } else {
                    var v29 int = 1
                    _ = v29
                }
                _ = fi28
            }
            var v30 float32 = 3.14159
            _ = v30
        }
    case "a":
        var v31 float64 = (0.1 + -2.25)
        _ = v31
    }
    var v32 int16 = 0
    _ = p0
    _ = v11
    _ = v32
    return ("hello" + "")
}

func main() {
    var v1 uint64 = 459
    v2 := []uint32{1}
    v3 := make(map[int]uint32)
    var v4 uint8 = 42
    var v5 float64 = (float64(0.0) * float64(3.14159))
    var v6 string = ("foo bar" + "\n\t")
    for fi33 := 0; fi33 < 17; fi33++ {
        for fi34 := 0; fi34 < 2; fi34++ {
            v4 -= (v4 | v4)
            _ = fi34
        }
        _ = fi33
    }
    if false {
        if (((3.14159 + 0.0) > 0.0) && true) {
            for fi35 := 0; fi35 < 9; fi35++ {
                var v36 uint32 = 2
                _ = fi35
                _ = v36
            }
        } else {
            var v37 int = -220
            _ = v37
        }
        if (4294967295 == -677) {
            var v38 float32 = (0.1 / ((1.5 * 3.14159) * ((1e10 / 1e10) + ((float32((v5 - (v5 * ((v5 / 1.0) - 0.0)))) / -2.25) - 3.14159))))
            if false {
                v4 = v4
                switch -573 {
                case 32767:
                    for fi39 := 0; fi39 < 5; fi39++ {
                        if true {
                            if true {
                                for fi40 := 0; fi40 < 7; fi40++ {
                                    v6 = (v6 + v6)
                                    switch (v6 + v6) {
                                    case "foo bar":
                                        v41 := GenS1{F0: 4294967295, F1: 97}
                                        _ = v41
                                    }
                                    _ = fi40
                                }
                                v42 := GenS3{F0: v1, F1: fi39}
                                var v43 int16 = 556
                                _ = v42
                                _ = v43
                            }
                            v44 := GenS3{F0: v1, F1: fi39}
                            _ = v44
                        }
                        _ = fi39
                    }
                case -932:
                    var v45 uint64 = 4294967295
                    _ = v45
                case 4294967295:
                    v6 = v6
                }
            }
            if true {
                v4 = v4
                v46 := GenS4{F0: v4, F1: 1}
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = v46
            }
            _ = v38
        }
    }
    v6 = (v6 + v6)
    v5 -= v5
    v5 = (float64((-2.25 + (3.14159 - ((3.14159 / 1.0) + float32(float64(((1.5 / 1e10) + 3.14159))))))) + (v5 / v5))
    fmt.Println(v5, v4)
    var v47 int64 = 9223372036854775807
    fmt.Println(v6, v4, v4)
    var v48 bool = false
    var v49 int16 = 0
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
