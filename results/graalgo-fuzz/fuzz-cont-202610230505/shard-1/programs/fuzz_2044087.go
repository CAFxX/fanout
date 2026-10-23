package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 string }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 uint32; F1 uint8; F2 uint16 }
type GenS5 struct { F0 int32; F1 uint64 }
func genFn0(p0 int64, p1 int16, p2 uint16) bool {
    for fi4 := 0; fi4 < 15; fi4++ {
        for fi5 := 0; fi5 < 3; fi5++ {
            p0 = -9223372036854775808
            _ = fi5
        }
        if true {
            for fi6 := 0; fi6 < 11; fi6++ {
                p0 -= -613
                _ = fi6
            }
        }
        _ = fi4
    }
    for fi7 := 0; fi7 < 10; fi7++ {
        fmt.Println(fi7, p0, p2)
        switch fi7 {
        case 997:
            v8 := GenS1{F0: (fi7 << ((65535) & 63)), F1: ("\n\t" + ("a" + ("hello" + ("\n\t" + ("hello" + "foo bar"))))), F2: "foo bar"}
            _ = v8
        case 127:
            fmt.Println(p1, p2, fi7)
        }
        _ = fi7
    }
    v9 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return (0 < 231)
}
func genFn1(p0 string) uint32 {
    v10 := GenS1{F0: 0, F1: (p0 + p0), F2: p0}
    switch 128 {
    case 2147483647:
        v11 := []int{}
        _ = v11
    case -9223372036854775808:
        var v12 int64 = -9223372036854775808
        _ = v12
    case -1:
        v13 := []int{128}
        _ = v13
    }
    _ = p0
    _ = v10
    return 1
}

func main() {
    v1 := make(map[int]bool)
    var v2 float32 = 1.5
    var v3 bool = (!(27 != -97))
    if v3 {
        v3 = (-2.25 > -2.25)
        v3 = (32767 < -221)
    } else {
        v1[32768] = true
        v14 := GenS1{F0: 9223372036854775807, F1: "a", F2: "foo bar"}
        v15 := []int{-2147483648, 1}
        _ = v14
        _ = v15
    }
    if false {
        v16 := []int{-9223372036854775808, 32767, 12}
        for fi17 := 0; fi17 < 19; fi17++ {
            if (false || true) {
                if true {
                    var v18 int = (fi17 & fi17)
                    fmt.Println(fi17, fi17)
                    _ = v18
                } else {
                    v19 := GenS2{F0: 0, F1: "hello"}
                    _ = v19
                }
                v2 = float32(3.14159)
                v20 := GenS3{F0: 42, F1: 797}
                _ = v20
            } else {
                for fi21 := 0; fi21 < 18; fi21++ {
                    var v22 uint64 = 42
                    var v23 int32 = -2147483648
                    if false {
                        var v24 int64 = -712
                        delete(v1, (fi17 - 127))
                        _ = v24
                    } else {
                        var v25 uint = 4294967295
                        fmt.Println(fi17, v23)
                        _ = v25
                    }
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                if v3 {
                    v26 := GenS4{F0: 42, F1: 55, F2: 442}
                    for fi27 := 0; fi27 < 12; fi27++ {
                        v28 := GenS1{F0: fi27, F1: ("foo bar" + ("" + "a")), F2: ""}
                        if (1 > 0) {
                            switch (("a" + "foo bar") + "foo bar") {
                            case "hello":
                                var v29 string = "a"
                                _ = v29
                            }
                        } else {
                            v3 = (!(v2 > v2))
                            var v30 string = "hello"
                            v3 = v3
                            _ = v30
                        }
                        _ = fi27
                        _ = v28
                    }
                    _ = v26
                }
                var v31 uint64 = 2
                _ = v31
            }
            _ = fi17
        }
        v3 = true
        _ = v16
    } else {
        v2 = v2
    }
    v32 := GenS3{F0: 913, F1: 355}
    _ = len(v1)
    v3 = (-88 == 0)
    v33 := GenS5{F0: 42, F1: 4294967295}
    v1[128] = (0 != 231)
    for fi34 := 0; fi34 < 16; fi34++ {
        var v35 string = "hello"
        var v36 float64 = -2.25
        v3 = v3
        _ = fi34
        _ = v35
        _ = v36
    }
    _ = len(v1)
    v37 := GenS1{F0: -843, F1: "a", F2: ("\n\t" + "a")}
    switch 128 {
    case -689:
        v2 -= 0.0
    case -1:
        var v38 string = "\n\t"
        _ = v38
    case 453:
        v2 -= (v2 / (1.5 * (float32(float64(((((1.5 / ((v2 + v2) + v2)) * 1e10) + (1.5 - (v2 - 1e10))) - float32((1e10 / -2.25))))) - (v2 + v2))))
    }
    if (v3 && true) {
        var v39 int64 = 0
        _ = v39
    }
    var v40 float32 = (-2.25 - v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v37
    _ = v40
}
