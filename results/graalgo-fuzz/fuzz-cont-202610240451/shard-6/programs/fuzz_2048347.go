package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64; F1 string }
func genFn0(p0 string, p1 uint16, p2 int) uint8 {
    var v4 int64 = -807
    v5 := GenS1{F0: p0, F1: 4294967295, F2: (^(v4))}
    if (0.0 > (float32(1e10) * 0.1)) {
        p2 = (-(p2))
        if ((true || false) || (19 >= 0)) {
            var v6 uint64 = 2
            _ = v6
        }
        v4 -= v4
    } else {
        var v7 string = (("hello" + "a") + p0)
        if (!true) {
            for fi8 := 0; fi8 < 14; fi8++ {
                if (-111 == -1) {
                    fmt.Println(v7, p0, p1)
                    fmt.Println(p0, v7)
                }
                v9 := []int{p2, (fi8 ^ 525), (p2 + fi8)}
                v9 = append(v9, 127)
                _ = fi8
                _ = v9
            }
            for fi10 := 0; fi10 < 10; fi10++ {
                if (2 == 127) {
                    v7 = (v7 + "foo bar")
                    p0 = p0
                    switch (^(p2)) {
                    case 32767:
                        v11 := GenS1{F0: "\n\t", F1: 4294967295, F2: -401}
                        _ = v11
                    case 688:
                        for fi12 := 0; fi12 < 20; fi12++ {
                            var v13 uint16 = (p1 << ((255) & 63))
                            fmt.Println(fi12)
                            p2 -= fi10
                            var v14 string = "hello"
                            if (1 != 255) {
                                fmt.Println(v13)
                                var v15 float64 = (1.5 * ((1e10 / 1.0) + (1.5 + (3.14159 + 1e10))))
                                _ = v15
                            }
                            _ = fi12
                            _ = v13
                            _ = v14
                        }
                    }
                } else {
                    v4 = -9223372036854775808
                }
                v16 := GenS1{F0: "foo bar", F1: 499, F2: v4}
                v17 := []int{}
                _ = fi10
                _ = v16
                _ = v17
            }
        }
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 1
}
func genFn1(p0 uint) int {
    v18 := GenS1{F0: "", F1: p0, F2: 686}
    var v19 int16 = 493
    for fi20 := 0; fi20 < 4; fi20++ {
        v21 := GenS1{F0: "", F1: (p0 << ((p0) & 63)), F2: 9223372036854775807}
        p0 = (p0 << ((2) & 63))
        if true {
            for fi22 := 0; fi22 < 19; fi22++ {
                p0 = (p0 | (p0 * 2))
                p0 = (p0 >> ((p0) & 63))
                var v23 float64 = (0.0 + 3.14159)
                _ = fi22
                _ = v23
            }
        } else {
            if (!(!(!false))) {
                v24 := []int{fi20}
                v19 = (-(v19))
                _ = v24
            } else {
                var v25 bool = (!(0.1 < float32(0.1)))
                fmt.Println(v19, fi20)
                _ = v25
            }
        }
        _ = fi20
        _ = v21
    }
    var v26 string = "\n\t"
    if (!(v26 < v26)) {
        var v27 uint = (p0 & p0)
        _ = v27
    } else {
        for fi28 := 0; fi28 < 6; fi28++ {
            v29 := GenS2{F0: p0}
            fmt.Println(fi28, v19)
            _ = fi28
            _ = v29
        }
    }
    _ = p0
    _ = v18
    _ = v19
    _ = v26
    return 9223372036854775807
}

func main() {
    var v1 uint32 = 173
    v2 := make(map[string]string)
    var v3 bool = (0 > 32767)
    v30 := GenS2{F0: 65535}
    var v31 int32 = 0
    var v32 uint64 = 4294967295
    var v33 uint = 1
    if true {
        v34 := GenS1{F0: "foo bar", F1: v33, F2: 9223372036854775807}
        var v35 int = -547
        var v36 string = "\n\t"
        _ = v34
        _ = v35
        _ = v36
    }
    if (uint(v1) <= (v33 + 465)) {
        _ = len(v2)
        for fi37 := 0; fi37 < 3; fi37++ {
            if ((100 == 403) && v3) {
                v31 = -64
                fmt.Println(v31, v31)
                var v38 int64 = -1
                _ = v38
            } else {
                var v39 int = genFn1((^(v33)))
                v40 := GenS1{F0: "", F1: 42, F2: -1}
                v41 := GenS2{F0: (v33 / ((1) | 1))}
                _ = v39
                _ = v40
                _ = v41
            }
            if ((("\n\t" + "foo bar") + "\n\t") == "\n\t") {
                v1 = (v1 % ((323) | 1))
                var v42 string = ("a" + ("foo bar" + "\n\t"))
                switch 42 {
                case -498:
                    if false {
                        _ = len(v2)
                        v3 = false
                        for fi43 := 0; fi43 < 15; fi43++ {
                            fmt.Println(v42)
                            fmt.Println(v31, v3, fi37)
                            _ = fi43
                        }
                    }
                case 42:
                    var v44 float64 = 0.0
                    _ = v44
                default:
                    v32 += (v32 / ((v32) | 1))
                }
                _ = v42
            }
            fmt.Println(v31)
            _ = fi37
        }
        if ((v32 >> ((42) & 63)) >= v32) {
            var v45 int16 = -32768
            _ = len(v2)
            var v46 uint16 = 2
            _ = v45
            _ = v46
        }
    }
    fmt.Println(v1, v33, v31)
    var v47 int64 = int64(v32)
    v48 := GenS1{F0: "foo bar", F1: v33, F2: int64(v31)}
    v49 := GenS3{F0: (v47 << ((18446744073709551615) & 63)), F1: ("foo bar" + ((("foo bar" + "hello") + "a") + "a"))}
    fmt.Println(v47, v3, v33)
    for fi50 := 0; fi50 < 19; fi50++ {
        delete(v2, ("\n\t" + ("" + "foo bar")))
        _ = fi50
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v47
    _ = v48
    _ = v49
}
