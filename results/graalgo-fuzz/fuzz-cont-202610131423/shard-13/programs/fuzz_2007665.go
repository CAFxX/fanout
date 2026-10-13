package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 int16; F2 uint }
func genFn0() int8 {
    if true {
        var v7 int16 = 914
        v8 := []int{9223372036854775807}
        _ = v7
        _ = v8
    } else {
        v9 := GenS1{F0: 68, F1: 2147483647}
        switch 0 {
        case 902:
            for fi10 := 0; fi10 < 4; fi10++ {
                for fi11 := 0; fi11 < 6; fi11++ {
                    fmt.Println(fi10, fi11)
                    _ = fi11
                }
                if true {
                    v12 := []int{(fi10 | 0), fi10, fi10, fi10}
                    var v13 float64 = 0.1
                    if false {
                        var v14 uint8 = 255
                        for fi15 := 0; fi15 < 4; fi15++ {
                            _ = len(v12)
                            _ = fi15
                        }
                        _ = v14
                    }
                    _ = v12
                    _ = v13
                } else {
                    switch fi10 {
                    case 452:
                        var v16 int32 = -851
                        _ = v16
                    case 284:
                        var v17 uint = 409
                        _ = v17
                    }
                }
                _ = fi10
            }
        case 0:
            v18 := []int{}
            _ = v18
        default:
            var v19 string = ("hello" + ("foo bar" + ("foo bar" + "")))
            _ = v19
        }
        v20 := GenS2{F0: 0}
        _ = v20
        _ = v9
    }
    v21 := GenS3{F0: 0, F1: 1000, F2: 676}
    var v22 int8 = 0
    _ = v21
    _ = v22
    return v22
}

func main() {
    var v1 uint16 = 32767
    v2 := make(map[string]string)
    var v3 uint64 = 2
    var v4 float32 = 0.0
    v5 := make(map[string]int)
    var v6 int = -2147483648
    var v23 int64 = 49
    v3 = v3
    switch "a" {
    case "a":
        var v24 uint64 = v3
        _ = v24
    default:
        fmt.Println(v23, v1, v6)
    }
    v5["hello"] = (-(v6))
    v25 := []int{v6}
    fmt.Println(v6, v3, v6)
    if true {
        v26 := v5[("foo bar" + (("hello" + "\n\t") + (("" + ("" + "")) + ("foo bar" + (("" + "") + "")))))]
        v23 += (v23 ^ (v23 & v23))
        _ = v26
    }
    var v27 uint64 = v3
    fmt.Println(v4, v23, v27)
    v28 := GenS1{F0: 255, F1: 768}
    for fi29 := 0; fi29 < 17; fi29++ {
        v4 = v4
        v30 := GenS3{F0: v1, F1: 1, F2: 2}
        _ = fi29
        _ = v30
    }
    switch 42 {
    case 0:
        var v31 uint = 863
        _ = v31
    default:
        for fi32 := 0; fi32 < 13; fi32++ {
            switch (fi32 / ((v6) | 1)) {
            case -876:
                var v33 uint64 = 18446744073709551615
                _ = v33
            default:
                var v34 uint16 = 868
                _ = v34
            }
            for fi35 := 0; fi35 < 14; fi35++ {
                if (0 >= 42) {
                    v6 = (fi32 >> ((1) & 63))
                    v23 += v23
                    v27 = 42
                } else {
                    var v36 uint16 = (v1 | v1)
                    _ = v36
                }
                var v37 int64 = (v23 % ((v23) | 1))
                _ = fi35
                _ = v37
            }
            _ = fi32
        }
    }
    var v38 int32 = 42
    v23 = v23
    if false {
        v27 -= (^(v27))
        for fi39 := 0; fi39 < 4; fi39++ {
            var v40 uint = 0
            _ = fi39
            _ = v40
        }
    }
    v41 := GenS3{F0: (v1 / ((v1) | 1)), F1: 1000, F2: 907}
    for fi42 := 0; fi42 < 19; fi42++ {
        var v43 uint64 = (v3 << ((499) & 63))
        if (42 == 42) {
            var v44 float64 = (float64(v4) - ((float64(v4) - 0.1) / (-2.25 + (-2.25 / float64(-2.25)))))
            _ = v44
        } else {
            v27 += (v3 / ((v27) | 1))
            switch (("foo bar" + "\n\t") + ("\n\t" + (("hello" + "foo bar") + ("a" + "hello")))) {
            case "foo bar":
                var v45 uint32 = 2147483647
                _ = v45
            case "hello":
                fmt.Println(v4, v38, v23)
            case "":
                v6 = fi42
            default:
                v4 = v4
            }
            if ((!(!true)) && false) {
                for fi46 := 0; fi46 < 10; fi46++ {
                    if ((v27 | 42) >= v27) {
                        var v47 bool = ("" < "a")
                        if true {
                            fmt.Println(v27, v43)
                        }
                        v23 = (^(v23))
                        _ = v47
                    }
                    fmt.Println(fi46)
                    v25 = append(v25, 272)
                    fmt.Println(v23)
                    v2[("foo bar" + ("" + "\n\t"))] = ("hello" + "a")
                    _ = fi46
                }
                var v48 int = fi42
                _ = v48
            }
        }
        _ = fi42
        _ = v43
    }
    var v49 int16 = 1000
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v27
    _ = v28
    _ = v3
    _ = v38
    _ = v4
    _ = v41
    _ = v49
    _ = v5
    _ = v6
}
