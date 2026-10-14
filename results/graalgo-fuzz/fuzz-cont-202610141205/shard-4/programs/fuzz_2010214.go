package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 uint16 }
func genFn0(p0 uint, p1 int16, p2 int) uint {
    var v6 float64 = float64(-2.25)
    if true {
        for fi7 := 0; fi7 < 3; fi7++ {
            p0 = p0
            p1 += 0
            for fi8 := 0; fi8 < 3; fi8++ {
                p1 += -32768
                _ = fi8
            }
            _ = fi7
        }
        switch ("foo bar" + "hello") {
        case "foo bar":
            v6 += 1e10
        case "":
            p1 -= -1
        case "a":
            v9 := []int{p2, p2, (p2 * 128), 2147483647}
            _ = v9
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (p0 - 0)
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]uint32)
    var v3 int = 127
    v4 := make(map[int]int)
    v5 := make(map[string]string)
    fmt.Println(v3)
    for fi10 := 0; fi10 < 18; fi10++ {
        switch "\n\t" {
        case "foo bar":
            v3 += 42
        default:
            switch ("\n\t" + "hello") {
            case "foo bar":
                var v11 int8 = 127
                _ = v11
            }
        }
        _ = len(v5)
        _ = fi10
    }
    v12 := GenS1{F0: 4294967295, F1: 127}
    var v13 int64 = 9223372036854775807
    var v14 int16 = 1000
    for fi15 := 0; fi15 < 20; fi15++ {
        var v16 float32 = 3.14159
        _ = fi15
        _ = v16
    }
    var v17 int16 = (-(v14))
    switch (v3 & 127) {
    case 974:
        if (false && (false || (!true))) {
            var v18 uint16 = 2
            for fi19 := 0; fi19 < 17; fi19++ {
                v2[2147483648] = 1
                _ = fi19
            }
            v13 -= v13
            _ = v18
        }
    }
    v17 = v14
    var v20 float64 = 1e10
    var v21 int32 = 855
    var v22 bool = (-879 < v21)
    if v22 {
        v23 := []int{(v3 << ((355) & 63))}
        _ = v23
    }
    if false {
        if v22 {
            var v24 uint8 = 128
            _ = v24
        } else {
            v25 := GenS2{F0: "a"}
            v20 = float64(0.0)
            var v26 bool = (v22 && (1 != v14))
            _ = v25
            _ = v26
        }
        v27 := GenS3{F0: 255, F1: 42}
        _ = v27
    } else {
        v21 = 1
        v17 += v17
        for fi28 := 0; fi28 < 19; fi28++ {
            if true {
                v29 := []int{}
                _ = v29
            }
            var v30 float32 = 0.1
            switch 1 {
            case 42:
                v20 = (v20 - 0.0)
            case 0:
                fmt.Println(v14, v3, v3)
            case -354:
                v17 += v14
            default:
                var v31 string = "a"
                _ = v31
            }
            _ = fi28
            _ = v30
        }
    }
    var v32 int64 = (v13 & -1)
    for fi33 := 0; fi33 < 5; fi33++ {
        if (127 < -128) {
            v34 := []int{9223372036854775807}
            switch "a" {
            case "foo bar":
                if v22 {
                    if v22 {
                        v17 += -868
                        if false {
                            fmt.Println(v3)
                        } else {
                            if (v22 && (0 <= (v21 & (v21 >> ((42) & 63))))) {
                                switch 1 {
                                case 568:
                                    v21 -= int32(v32)
                                case 9223372036854775807:
                                    var v35 int32 = v21
                                    _ = v35
                                }
                            } else {
                                switch fi33 {
                                case 708:
                                    fmt.Println(fi33)
                                }
                            }
                            v14 -= v17
                        }
                    }
                    if ((1 <= 4294967295) || v22) {
                        fmt.Println(v3)
                    }
                    fmt.Println(v14)
                } else {
                    switch (fi33 << ((255) & 63)) {
                    case -2147483648:
                        v22 = v22
                    }
                }
            case "a":
                fmt.Println(v17)
            case "\n\t":
                if true {
                    switch v3 {
                    case -1:
                        var v36 int64 = v13
                        _ = v36
                    default:
                        switch fi33 {
                        case 32767:
                            delete(v2, 785)
                        }
                    }
                    var v37 int8 = 127
                    _ = v37
                }
            }
            v3 = fi33
            _ = v34
        }
        fmt.Println(v3, v17)
        switch 4294967295 {
        case 32768:
            var v38 float32 = (1e10 * (0.1 / 1e10))
            _ = v38
        case 473:
            var v39 int16 = (v17 - v14)
            _ = v39
        }
        _ = fi33
    }
    var v40 int32 = v21
    v22 = v22
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v32
    _ = v4
    _ = v40
    _ = v5
}
