package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 string; F2 int }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int32; F1 uint64; F2 bool }
func genFn0() int {
    var v6 int = 39
    var v7 int16 = 0
    _ = v6
    _ = v7
    return (v6 % ((-874) | 1))
}
func genFn1(p0 uint, p1 int, p2 bool) uint16 {
    v8 := GenS1{F0: 42, F1: -614, F2: (^(p1))}
    fmt.Println(p2)
    v9 := []int{p1}
    p0 = (^(p0))
    for fi10 := 0; fi10 < 5; fi10++ {
        p1 = fi10
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 1
}

func main() {
    v1 := make(map[string]int)
    v2 := []string{}
    var v3 int = 2147483647
    var v4 uint64 = 0
    var v5 int8 = -1
    var v11 float64 = (1.5 - 0.0)
    if (!false) {
        var v12 string = ("a" + ("foo bar" + "hello"))
        var v13 uint32 = 1
        for fi14 := 0; fi14 < 20; fi14++ {
            switch ("a" + ((v12 + v12) + "a")) {
            case "a":
                for fi15 := 0; fi15 < 9; fi15++ {
                    fmt.Println(v11, fi15)
                    if false {
                        var v16 float64 = (v11 / float64(3.14159))
                        _ = v16
                    } else {
                        switch v12 {
                        case "a":
                            if true {
                                v17 := GenS2{F0: 794}
                                _ = v17
                            } else {
                                if (!false) {
                                    v11 = 3.14159
                                    v18 := v1[v12]
                                    _ = v18
                                } else {
                                    fmt.Println(fi15, v12, v3)
                                }
                            }
                        case "hello":
                            v12 = "foo bar"
                        case "\n\t":
                            v19 := GenS3{F0: ("\n\t" + v12), F1: ((v12 + "") + v12), F2: (v3 + (fi14 + (fi14 / ((v3) | 1))))}
                            _ = v19
                        }
                        var v20 uint64 = 4294967295
                        _ = v20
                    }
                    switch fi15 {
                    case -410:
                        switch v12 {
                        case "foo bar":
                            var v21 uint = uint(v4)
                            _ = v21
                        case "hello":
                            v3 = genFn0()
                        case "\n\t":
                            var v22 int32 = -235
                            _ = v22
                        default:
                            v23 := v1[v12]
                            v24 := GenS4{F0: false}
                            _ = v23
                            _ = v24
                        }
                    }
                    if (true && (true || (159 <= 10))) {
                        if false {
                            var v25 uint64 = v4
                            delete(v1, v12)
                            var v26 uint64 = v4
                            var v27 uint = uint(v25)
                            _ = v25
                            _ = v26
                            _ = v27
                        } else {
                            var v28 uint = 0
                            _ = v28
                        }
                        if (65535 > 42) {
                            var v29 uint8 = 2
                            _ = v29
                        }
                        v5 -= v5
                    }
                    var v30 int64 = -9223372036854775808
                    _ = fi15
                    _ = v30
                }
            case "hello":
                v4 = (v4 / ((4294967295) | 1))
            case "foo bar":
                v31 := v1[(v12 + v12)]
                _ = v31
            }
            _ = fi14
        }
        _ = v12
        _ = v13
    } else {
        v32 := GenS5{F0: 18446744073709551615}
        var v33 bool = false
        v34 := GenS3{F0: ("a" + (("hello" + "a") + "hello")), F1: "hello", F2: -2147483648}
        _ = v32
        _ = v33
        _ = v34
    }
    var v35 string = ("" + (((("a" + "hello") + "\n\t") + "a") + ""))
    v5 -= v5
    if true {
        v36 := GenS5{F0: 65535}
        var v37 int64 = -497
        _ = v36
        _ = v37
    } else {
        if (v5 < v5) {
            var v38 uint32 = 42
            _ = v38
        } else {
            v1[("foo bar" + (("hello" + v35) + "hello"))] = (v3 >> ((42) & 63))
        }
    }
    var v39 int8 = v5
    v39 = v39
    fmt.Println(v5, v11, v4)
    v40 := GenS6{F0: -1, F1: v4, F2: true}
    v11 = (float64(0.0) + v11)
    v35 = v35
    switch v35 {
    case "foo bar":
        var v41 uint32 = 65535
        _ = v41
    default:
        for fi42 := 0; fi42 < 11; fi42++ {
            var v43 uint32 = 0
            _ = fi42
            _ = v43
        }
    }
    var v44 uint32 = 2
    switch (-(v3)) {
    case 1:
        fmt.Println(v4)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v35
    _ = v39
    _ = v4
    _ = v40
    _ = v44
    _ = v5
}
