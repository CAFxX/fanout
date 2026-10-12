package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 uint8; F2 int32 }
type GenS4 struct { F0 bool }
func genFn0() string {
    var v4 uint32 = 450
    v4 = (v4 & v4)
    _ = v4
    return (("foo bar" + "\n\t") + "foo bar")
}
func genFn1(p0 int32) int8 {
    if (32767 != 0) {
        var v5 float32 = (1.5 / 0.1)
        if (1 > -526) {
            var v6 uint64 = 0
            v5 -= v5
            _ = v6
        }
        _ = v5
    } else {
        switch (("a" + "foo bar") + ("\n\t" + "\n\t")) {
        case "\n\t":
            v7 := []int{42, 32767}
            _ = v7
        case "foo bar":
            if false {
                var v8 int32 = p0
                for fi9 := 0; fi9 < 17; fi9++ {
                    v10 := GenS1{F0: p0}
                    _ = fi9
                    _ = v10
                }
                v11 := []int{618, -2147483648}
                _ = v11
                _ = v8
            } else {
                p0 = p0
                fmt.Println(p0)
                if false {
                    for fi12 := 0; fi12 < 15; fi12++ {
                        if true {
                            var v13 int16 = 1000
                            switch fi12 {
                            case -721:
                                v14 := GenS2{F0: (v13 & (v13 & v13))}
                                _ = v14
                            default:
                                for fi15 := 0; fi15 < 13; fi15++ {
                                    v16 := []int{2147483647, (fi12 / (((fi12 & fi12)) | 1)), fi12, (-(fi12))}
                                    if (1e10 >= (((1.5 + (-2.25 + 3.14159)) - float64(3.14159)) + 0.0)) {
                                        v13 += (v13 * v13)
                                        var v17 float32 = 1.5
                                        var v18 int32 = (p0 * (p0 << ((18446744073709551615) & 63)))
                                        _ = v17
                                        _ = v18
                                    } else {
                                        var v19 int32 = (p0 + -1)
                                        _ = v19
                                    }
                                    p0 += p0
                                    _ = fi15
                                    _ = v16
                                }
                            }
                            fmt.Println(p0)
                            _ = v13
                        } else {
                            v20 := GenS1{F0: p0}
                            p0 += p0
                            p0 -= p0
                            _ = v20
                        }
                        var v21 int32 = p0
                        var v22 uint = 0
                        _ = fi12
                        _ = v21
                        _ = v22
                    }
                }
            }
        case "a":
            var v23 int64 = 42
            _ = v23
        default:
            p0 -= -1
        }
    }
    var v24 int64 = -69
    v25 := []int{4294967295}
    if (true && ("hello" <= "a")) {
        v24 += (^(v24))
        for fi26 := 0; fi26 < 13; fi26++ {
            fmt.Println(p0)
            v24 = (v24 << ((309) & 63))
            for fi27 := 0; fi27 < 12; fi27++ {
                if len(v25) >= 2 {
                    v25 = v25[1:len(v25)]
                }
                if len(v25) >= 2 {
                    v25 = v25[1:len(v25)]
                }
                fmt.Println(v24, fi27)
                v25 = append(v25, fi27)
                _ = fi27
            }
            _ = fi26
        }
        for fi28 := 0; fi28 < 18; fi28++ {
            p0 = -1
            p0 += 2147483647
            if len(v25) >= 2 {
                v25 = v25[1:len(v25)]
            }
            _ = fi28
        }
    }
    _ = p0
    _ = v24
    _ = v25
    return 34
}

func main() {
    v1 := []int{128, 4294967295, -932}
    var v2 float32 = 3.14159
    var v3 float64 = 0.1
    v3 = v3
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v3, v3)
    _ = len(v1)
    v2 = 3.14159
    v2 = 1.5
    _ = len(v1)
    for fi29 := 0; fi29 < 7; fi29++ {
        var v30 int16 = -32768
        if true {
            fmt.Println(v2, fi29)
            if (!((1.5 + v3) != float64(1e10))) {
                var v31 string = "\n\t"
                v3 = 1e10
                v32 := GenS3{F0: 538, F1: 74, F2: 2147483647}
                _ = v31
                _ = v32
            }
            var v33 int = (fi29 | (fi29 - fi29))
            _ = v33
        }
        _ = fi29
        _ = v30
    }
    for fi34 := 0; fi34 < 17; fi34++ {
        switch (fi34 - fi34) {
        case 42:
            var v35 float32 = v2
            _ = v35
        case 2147483647:
            var v36 int8 = -128
            _ = v36
        default:
            for fi37 := 0; fi37 < 3; fi37++ {
                v38 := GenS4{F0: true}
                var v39 uint32 = 518
                _ = fi37
                _ = v38
                _ = v39
            }
        }
        _ = fi34
    }
    if (true || false) {
        _ = len(v1)
        v40 := GenS4{F0: false}
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v40
    } else {
        if len(v1) > 0 {
            v1[0] = 128
        }
        v41 := GenS4{F0: false}
        _ = v41
    }
    var v42 uint32 = 65535
    var v43 int32 = -200
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v42
    _ = v43
}
