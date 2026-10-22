package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int; F1 string; F2 uint64 }
type GenS3 struct { F0 int; F1 bool; F2 int32 }
func genFn0(p0 uint8, p1 int, p2 uint16) int32 {
    for fi5 := 0; fi5 < 17; fi5++ {
        for fi6 := 0; fi6 < 16; fi6++ {
            v7 := []int{fi6, fi5, (fi6 | (p1 % (((p1 & (^(fi6)))) | 1)))}
            var v8 bool = false
            switch ("\n\t" + "hello") {
            case "\n\t":
                var v9 float32 = 0.1
                _ = v9
            case "foo bar":
                v8 = (p2 != p2)
            }
            _ = fi6
            _ = v7
            _ = v8
        }
        _ = fi5
    }
    fmt.Println(p0, p1, p2)
    v10 := GenS1{F0: 999, F1: 1, F2: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 440
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{}
    v3 := make(map[int]string)
    v4 := make(map[int]int)
    switch 42 {
    case 114:
        for fi11 := 0; fi11 < 11; fi11++ {
            for fi12 := 0; fi12 < 7; fi12++ {
                _ = len(v2)
                _ = fi12
            }
            var v13 int8 = 127
            if ((fi11 % ((fi11) | 1)) == fi11) {
                var v14 string = "\n\t"
                for fi15 := 0; fi15 < 1; fi15++ {
                    v14 = (v14 + v14)
                    var v16 uint = 741
                    var v17 float32 = (1e10 * (1.5 - -2.25))
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                _ = v14
            }
            _ = fi11
            _ = v13
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v18 uint16 = 42
    if false {
        delete(v3, -327)
        for fi19 := 0; fi19 < 14; fi19++ {
            switch "hello" {
            case "a":
                for fi20 := 0; fi20 < 5; fi20++ {
                    fmt.Println(fi20)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = len(v1)
                    _ = fi20
                }
            default:
                var v21 uint16 = v18
                _ = v21
            }
            _ = fi19
        }
        switch (("foo bar" + "\n\t") + "hello") {
        case "hello":
            fmt.Println(v18)
        case "foo bar":
            var v22 bool = true
            _ = v22
        case "":
            var v23 bool = (!(!true))
            _ = v23
        }
    }
    var v24 float64 = float64(0.1)
    v25 := v4[128]
    var v26 int = 125
    v24 = (v24 * float64(1.5))
    fmt.Println(v25)
    for fi27 := 0; fi27 < 12; fi27++ {
        v28 := GenS1{F0: 42, F1: v18, F2: 100}
        v24 = (((v24 / float64(-2.25)) * (float64(1.5) * (v24 - (3.14159 / (v24 * v24))))) * v24)
        _ = fi27
        _ = v28
    }
    v2 = append(v2, (("foo bar" + "\n\t") + ("" + "")))
    var v29 bool = (-774 == 1)
    if v29 {
        v24 = v24
        v30 := GenS1{F0: 0, F1: 791, F2: 2}
        _ = v30
    } else {
        v29 = (!(!v29))
    }
    if v29 {
        for fi31 := 0; fi31 < 1; fi31++ {
            var v32 int16 = -32768
            switch v25 {
            case 128:
                fmt.Println(v24)
            case 32768:
                var v33 uint = 4294967295
                _ = v33
            case 4294967295:
                switch (fi31 % ((v25) | 1)) {
                case 1:
                    v24 -= v24
                }
            default:
                for fi34 := 0; fi34 < 5; fi34++ {
                    var v35 int64 = 0
                    _ = fi34
                    _ = v35
                }
            }
            v24 = v24
            _ = fi31
            _ = v32
        }
        v36 := GenS1{F0: 2, F1: 2, F2: 2147483648}
        _ = v36
    }
    switch "hello" {
    case "a":
        if (1.5 >= 3.14159) {
            var v37 float64 = (-2.25 + 0.0)
            _ = len(v2)
            for fi38 := 0; fi38 < 1; fi38++ {
                if v29 {
                    var v39 uint8 = 42
                    _ = v39
                } else {
                    var v40 float32 = 1.5
                    fmt.Println(v26, v29)
                    _ = v40
                }
                for fi41 := 0; fi41 < 14; fi41++ {
                    var v42 uint = 65535
                    if (-505 > v26) {
                        v43 := GenS2{F0: v26, F1: "", F2: 18446744073709551615}
                        switch "foo bar" {
                        case "hello":
                            v44 := GenS3{F0: -1, F1: v29, F2: 1}
                            _ = v44
                        case "":
                            v37 += -2.25
                        }
                        if v29 {
                            var v45 float64 = (((v37 / v24) / -2.25) + (v24 / 3.14159))
                            _ = v45
                        } else {
                            var v46 string = "a"
                            delete(v4, (v26 | 4294967295))
                            _ = v46
                        }
                        _ = v43
                    }
                    var v47 bool = v29
                    _ = fi41
                    _ = v42
                    _ = v47
                }
                _ = fi38
            }
            _ = v37
        }
    case "\n\t":
        for fi48 := 0; fi48 < 17; fi48++ {
            var v49 uint64 = 638
            fmt.Println(v24, v49)
            _ = fi48
            _ = v49
        }
    }
    v50 := GenS1{F0: 0, F1: v18, F2: 2}
    for fi51 := 0; fi51 < 18; fi51++ {
        v18 -= (v18 % (((v18 | (v18 << ((42) & 63)))) | 1))
        fmt.Println(v26)
        for fi52 := 0; fi52 < 7; fi52++ {
            if v29 {
                var v53 string = "\n\t"
                v24 = float64((0.1 * 1.5))
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = v53
            } else {
                if (!(629 < 22)) {
                    var v54 float64 = v24
                    _ = v54
                } else {
                    fmt.Println(v18)
                }
                var v55 int64 = 0
                _ = v55
            }
            var v56 string = "hello"
            var v57 float64 = ((1e10 + ((1e10 - v24) - v24)) - v24)
            _ = fi52
            _ = v56
            _ = v57
        }
        _ = fi51
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v4
    _ = v50
}
