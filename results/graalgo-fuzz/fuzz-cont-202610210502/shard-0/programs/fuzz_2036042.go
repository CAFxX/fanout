package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 string; F1 string; F2 uint8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32 }
func genFn0(p0 int, p1 bool) uint64 {
    v4 := []int{}
    p0 -= p0
    if (42 != 255) {
        switch (^(p0)) {
        case -2147483648:
            if true {
                p0 += p0
                for fi5 := 0; fi5 < 18; fi5++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    p0 -= (^(fi5))
                    _ = fi5
                }
                v4 = append(v4, p0)
            } else {
                p0 += -459
                v6 := GenS1{F0: -398, F1: 1}
                switch ((("hello" + "foo bar") + ("" + "hello")) + "\n\t") {
                case "hello":
                    var v7 int32 = -921
                    _ = v7
                default:
                    var v8 uint64 = 0
                    _ = v8
                }
                _ = v6
            }
        case -1:
            for fi9 := 0; fi9 < 7; fi9++ {
                if false {
                    var v10 int = (^(p0))
                    var v11 string = ("" + "a")
                    v12 := GenS2{F0: "foo bar", F1: "", F2: 42}
                    _ = v10
                    _ = v11
                    _ = v12
                }
                var v13 uint64 = 76
                if ((402 != 2) || p1) {
                    p0 = 4294967295
                    fmt.Println(v13, fi9, v13)
                }
                _ = fi9
                _ = v13
            }
        default:
            p0 = -470
        }
    }
    for fi14 := 0; fi14 < 4; fi14++ {
        if ((!p1) || p1) {
            p1 = p1
            v15 := GenS3{F0: 1}
            for fi16 := 0; fi16 < 4; fi16++ {
                p0 -= (p0 * 127)
                p0 = (p0 & (fi14 | (fi16 / ((fi16) | 1))))
                _ = fi16
            }
            _ = v15
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi14
    }
    var v17 uint = 42
    _ = p0
    _ = p1
    _ = v17
    _ = v4
    return 535
}
func genFn1() uint {
    var v18 string = ("\n\t" + "a")
    switch 32768 {
    case 0:
        switch 128 {
        case 252:
            v18 = (v18 + "")
        case -764:
            var v19 bool = true
            _ = v19
        }
    case 32767:
        var v20 uint32 = 2147483647
        _ = v20
    case 128:
        switch 4294967295 {
        case 2147483647:
            v21 := GenS4{F0: 512}
            _ = v21
        case 743:
            v18 = "a"
        default:
            if (86 < 797) {
                var v22 int8 = 0
                _ = v22
            }
        }
    }
    _ = v18
    return 4294967295
}

func main() {
    var v1 string = "hello"
    v2 := make(map[string]bool)
    var v3 int = 0
    switch (v3 | 32767) {
    case -89:
        v3 = 32767
    case -381:
        if (42 < 2) {
            var v23 int32 = 2147483647
            _ = v23
        } else {
            fmt.Println(v1, v1)
            var v24 int8 = -46
            if (1 == 115) {
                v1 = (("foo bar" + (v1 + (v1 + v1))) + ("a" + "hello"))
            } else {
                if false {
                    if (-2.25 < float64(-2.25)) {
                        if true {
                            var v25 uint16 = 538
                            _ = v25
                        }
                        fmt.Println(v3)
                    } else {
                        v1 = "\n\t"
                        var v26 float32 = ((-2.25 * -2.25) * (1.5 * float32(1e10)))
                        _ = v26
                    }
                    v24 = 1
                    delete(v2, v1)
                } else {
                    v24 += 42
                    _ = len(v2)
                    var v27 int32 = -2147483648
                    _ = v27
                }
                var v28 bool = (true && false)
                v29 := []int{-1, (v3 + (v3 + (v3 & 140)))}
                _ = v28
                _ = v29
            }
            _ = v24
        }
    }
    v30 := []int{}
    v31 := GenS3{F0: 282}
    v1 = "\n\t"
    var v32 int16 = 42
    switch (v3 ^ -609) {
    case 127:
        v32 = v32
    }
    if false {
        var v33 string = v1
        for fi34 := 0; fi34 < 9; fi34++ {
            v1 = ((v1 + "") + (v33 + v1))
            _ = fi34
        }
        _ = v33
    }
    fmt.Println(v1, v1, v3)
    v3 = (v3 / ((32768) | 1))
    _ = len(v2)
    for fi35 := 0; fi35 < 12; fi35++ {
        for fi36 := 0; fi36 < 16; fi36++ {
            v3 += 364
            v3 += (fi35 ^ (^(fi35)))
            v2[v1] = (!true)
            v2[v1] = (v1 != (v1 + "\n\t"))
            _ = fi36
        }
        _ = fi35
    }
    v3 = v3
    v37 := v2[(v1 + ((v1 + ("a" + v1)) + ((("\n\t" + ((("hello" + (v1 + v1)) + v1) + v1)) + v1) + (v1 + "a"))))]
    fmt.Println(v1)
    fmt.Println(v3, v32, v1)
    var v38 float32 = (float32(1e10) - (1.5 / ((1e10 + (float32(0.1) * 3.14159)) + 0.0)))
    v37 = v37
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v32:", v32)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v37
    _ = v38
}
