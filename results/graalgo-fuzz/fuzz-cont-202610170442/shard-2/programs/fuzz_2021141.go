package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 uint16; F1 int16; F2 uint }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int32; F1 string; F2 int8 }
func genFn0(p0 uint32, p1 int16) uint16 {
    if true {
        var v7 int16 = p1
        _ = v7
    }
    var v8 uint32 = p0
    var v9 int32 = 505
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 32767
}
func genFn1(p0 uint8, p1 uint32, p2 int8) string {
    p1 -= (^(p1))
    v10 := GenS1{F0: 4294967295}
    p0 -= (p0 | p0)
    for fi11 := 0; fi11 < 6; fi11++ {
        fmt.Println(p0, fi11)
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return "hello"
}

func main() {
    v1 := []uint32{2147483647, 0, 819}
    v2 := []uint32{65535, 65535, 214}
    var v3 int64 = 1
    var v4 int16 = 0
    v5 := make(map[string]int)
    var v6 int = 32767
    var v12 uint16 = 32767
    v4 = v4
    v13 := v5[("foo bar" + "hello")]
    var v14 int64 = v3
    v5["foo bar"] = (v6 % ((v6) | 1))
    if (0.1 != (1.5 - -2.25)) {
        _ = len(v2)
        var v15 int8 = 1
        for fi16 := 0; fi16 < 7; fi16++ {
            v17 := GenS2{F0: (v6 + 127), F1: ("" + ("\n\t" + "\n\t"))}
            _ = fi16
            _ = v17
        }
        _ = v15
    }
    for fi18 := 0; fi18 < 15; fi18++ {
        for fi19 := 0; fi19 < 12; fi19++ {
            v20 := GenS3{F0: v12, F1: (v4 & v4), F2: 925}
            if true {
                var v21 int8 = 42
                if true {
                    v3 += -1
                    v22 := GenS4{F0: ("hello" + ("\n\t" + "a"))}
                    _ = v22
                }
                _ = v21
            }
            var v23 uint16 = 32767
            _ = fi19
            _ = v20
            _ = v23
        }
        delete(v5, "foo bar")
        _ = fi18
    }
    v24 := GenS1{F0: (v13 / ((v13) | 1))}
    v25 := GenS2{F0: (^(v13)), F1: "\n\t"}
    switch 4294967295 {
    case 2147483647:
        v26 := GenS4{F0: "foo bar"}
        _ = v26
    }
    v4 = (v4 & (v4 << ((255) & 63)))
    var v27 uint32 = 1
    if (v3 > v14) {
        switch (^(v6)) {
        case 4294967295:
            v5[("\n\t" + "\n\t")] = (v6 * (v13 + -291))
        }
        for fi28 := 0; fi28 < 7; fi28++ {
            for fi29 := 0; fi29 < 12; fi29++ {
                switch "hello" {
                case "foo bar":
                    v30 := GenS4{F0: "a"}
                    _ = v30
                case "\n\t":
                    var v31 bool = true
                    _ = v31
                default:
                    v32 := GenS3{F0: (v12 << ((107) & 63)), F1: (v4 + (v4 % ((1) | 1))), F2: 808}
                    _ = v32
                }
                _ = fi29
            }
            var v33 float64 = 0.0
            if false {
                v34 := GenS3{F0: genFn0((v27 << ((42) & 63)), (v4 << ((460) & 63))), F1: -32768, F2: 4294967295}
                _ = v34
            }
            _ = fi28
            _ = v33
        }
        if false {
            var v35 uint16 = (v12 - v12)
            var v36 uint32 = (^(v27))
            fmt.Println(v36)
            _ = v35
            _ = v36
        } else {
            if false {
                for fi37 := 0; fi37 < 5; fi37++ {
                    for fi38 := 0; fi38 < 9; fi38++ {
                        v6 = (v13 ^ 4294967295)
                        var v39 int32 = -1
                        var v40 uint64 = 67
                        _ = fi38
                        _ = v39
                        _ = v40
                    }
                    var v41 int16 = (v4 - int16(v14))
                    _ = fi37
                    _ = v41
                }
            }
        }
    } else {
        var v42 bool = false
        v27 -= (v27 + v27)
        _ = v42
    }
    v43 := GenS5{F0: 0, F1: "a", F2: -128}
    var v44 uint32 = v27
    var v45 uint64 = 179
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v24
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
