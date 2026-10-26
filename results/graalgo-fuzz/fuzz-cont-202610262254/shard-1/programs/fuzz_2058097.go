package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 int16; F1 uint }
func genFn0(p0 uint64, p1 int8) int8 {
    v6 := GenS1{F0: 484, F1: "\n\t"}
    p1 += p1
    _ = p0
    _ = p1
    _ = v6
    return -128
}
func genFn1(p0 bool, p1 uint32) uint64 {
    if p0 {
        p1 += 2147483647
        v7 := GenS2{F0: p0, F1: 42}
        _ = v7
    } else {
        p0 = true
        if p0 {
            p1 += (p1 | (^(p1)))
            if (false && p0) {
                if (!p0) {
                    p0 = p0
                    for fi8 := 0; fi8 < 9; fi8++ {
                        for fi9 := 0; fi9 < 2; fi9++ {
                            var v10 uint32 = (^(p1))
                            p0 = (852 != 1)
                            var v11 string = (("hello" + "\n\t") + "foo bar")
                            var v12 uint32 = p1
                            var v13 int16 = -344
                            _ = fi9
                            _ = v10
                            _ = v11
                            _ = v12
                            _ = v13
                        }
                        _ = fi8
                    }
                    switch 128 {
                    case 291:
                        var v14 uint8 = 168
                        _ = v14
                    case -692:
                        if true {
                            v15 := []int{}
                            p0 = p0
                            _ = v15
                        }
                    case -512:
                        var v16 uint32 = p1
                        _ = v16
                    }
                }
                v17 := []int{}
                _ = v17
            } else {
                var v18 uint8 = 40
                v18 += v18
                for fi19 := 0; fi19 < 15; fi19++ {
                    var v20 int = (fi19 - (fi19 * 128))
                    v21 := GenS1{F0: 238, F1: "a"}
                    p0 = p0
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                _ = v18
            }
        } else {
            if (true && ((-128 > 25) && p0)) {
                fmt.Println(p1)
                fmt.Println(p0, p0)
                var v22 uint32 = (^(p1))
                _ = v22
            }
            for fi23 := 0; fi23 < 20; fi23++ {
                for fi24 := 0; fi24 < 17; fi24++ {
                    var v25 uint32 = 2147483647
                    fmt.Println(v25, p1)
                    _ = fi24
                    _ = v25
                }
                fmt.Println(p0)
                v26 := GenS2{F0: true, F1: 128}
                _ = fi23
                _ = v26
            }
            fmt.Println(p1, p1)
        }
    }
    fmt.Println(p1)
    v27 := GenS3{F0: 32767, F1: 65535}
    var v28 float64 = 3.14159
    _ = p0
    _ = p1
    _ = v27
    _ = v28
    return 42
}

func main() {
    var v1 bool = (32767 > -1)
    v2 := make(map[string]bool)
    v3 := make(map[int]int)
    var v4 int64 = 42
    v5 := make(map[string]string)
    var v29 uint64 = 4294967295
    v29 -= 1
    if v1 {
        v1 = v1
    } else {
        v29 = genFn1((!("\n\t" < ("a" + ("foo bar" + ("a" + "a"))))), 684)
        v30 := GenS1{F0: 2, F1: "a"}
        _ = v30
    }
    for fi31 := 0; fi31 < 17; fi31++ {
        if v1 {
            v32 := GenS1{F0: 1, F1: "a"}
            _ = v32
        } else {
            for fi33 := 0; fi33 < 4; fi33++ {
                fmt.Println(fi33)
                v29 = v29
                v1 = false
                _ = fi33
            }
            v4 = v4
        }
        v34 := GenS1{F0: 318, F1: "\n\t"}
        _ = fi31
        _ = v34
    }
    switch ("foo bar" + ("hello" + "foo bar")) {
    case "foo bar":
        v35 := GenS2{F0: false, F1: 74}
        _ = v35
    case "a":
        v29 += (^(v29))
    case "hello":
        v1 = (!false)
    default:
        var v36 uint64 = (^(v29))
        _ = v36
    }
    switch 32768 {
    case 4294967295:
        switch "" {
        case "":
            v4 = (v4 | v4)
        default:
            var v37 int8 = 22
            _ = v37
        }
    default:
        var v38 int32 = 1
        _ = v38
    }
    v4 -= (v4 + v4)
    var v39 int = 32767
    v40 := []int{(v39 & 0)}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
