package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint; F1 bool; F2 uint64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 string; F1 uint }
func genFn0(p0 int16, p1 int32) uint16 {
    var v5 uint = 94
    v6 := GenS1{F0: 2147483647}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 32767
}
func genFn1(p0 int8, p1 int64, p2 string) uint8 {
    var v7 float32 = 0.1
    var v8 float32 = 3.14159
    for fi9 := 0; fi9 < 6; fi9++ {
        if ((!true) && ((-2.25 * 3.14159) >= (1.5 - 0.1))) {
            switch ((p2 + "\n\t") + (p2 + "")) {
            case "a":
                if true {
                    p2 = p2
                } else {
                    v10 := []int{fi9, 4294967295, -2147483648}
                    switch fi9 {
                    case -2147483648:
                        p1 = (p1 + 1)
                    }
                    switch ("hello" + "a") {
                    case "foo bar":
                        for fi11 := 0; fi11 < 2; fi11++ {
                            v12 := GenS1{F0: 65535}
                            if (!false) {
                                v8 -= v8
                            }
                            _ = fi11
                            _ = v12
                        }
                    case "hello":
                        p1 += p1
                    }
                    _ = v10
                }
            case "hello":
                var v13 float32 = (1e10 - (float32(1e10) - 1e10))
                _ = v13
            default:
                v14 := GenS2{F0: 0}
                _ = v14
            }
            p1 -= (p1 % (((^(p1))) | 1))
        } else {
            switch fi9 {
            case -2147483648:
                p0 -= p0
            case 128:
                switch p2 {
                case "hello":
                    var v15 uint64 = 1
                    _ = v15
                case "":
                    v7 = float32(3.14159)
                case "foo bar":
                    if true {
                        if false {
                            p0 += p0
                            v16 := GenS3{F0: 1}
                            _ = v16
                        }
                        v8 = v7
                    } else {
                        p0 -= 34
                    }
                default:
                    v17 := GenS4{F0: 1, F1: (65535 <= 32767), F2: 67}
                    _ = v17
                }
            default:
                fmt.Println(p1)
            }
            v8 -= -2.25
            var v18 uint64 = 386
            _ = v18
        }
        for fi19 := 0; fi19 < 5; fi19++ {
            v20 := GenS2{F0: (^(p1))}
            _ = fi19
            _ = v20
        }
        var v21 string = ("hello" + (((p2 + (((((p2 + ("" + p2)) + (("foo bar" + "a") + p2)) + p2) + ((p2 + p2) + ("a" + p2))) + (("a" + "") + p2))) + p2) + p2))
        _ = fi9
        _ = v21
    }
    var v22 float32 = 0.1
    fmt.Println(v8, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    _ = v7
    _ = v8
    return 212
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]int)
    var v3 int64 = 1
    var v4 int32 = 2147483647
    var v23 int64 = (v3 + v3)
    if false {
        v4 -= v4
        v24 := GenS5{F0: -2147483648}
        _ = v24
    }
    v3 = 9223372036854775807
    var v25 int = 4294967295
    if true {
        v3 = v23
        if true {
            for fi26 := 0; fi26 < 18; fi26++ {
                if false {
                    var v27 int = (-(fi26))
                    var v28 uint16 = 1
                    v29 := GenS6{F0: "a", F1: 1}
                    _ = v27
                    _ = v28
                    _ = v29
                }
                _ = fi26
            }
        }
    } else {
        var v30 uint64 = 850
        var v31 string = "hello"
        fmt.Println(v3, v25, v31)
        _ = v30
        _ = v31
    }
    var v32 uint64 = 512
    var v33 uint8 = genFn1(42, (v3 * 8), "")
    v34 := GenS6{F0: "foo bar", F1: 12}
    fmt.Println(v33, v32)
    if false {
        fmt.Println(v32, v32)
    }
    fmt.Println(v4)
    v35 := GenS6{F0: "", F1: 1}
    v36 := []int{v25, v25}
    v37 := GenS3{F0: -1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
}
