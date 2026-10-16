package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int64 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 string; F1 string }
func genFn0(p0 bool, p1 uint16) uint32 {
    v4 := GenS1{F0: -1, F1: -45, F2: -9223372036854775808}
    v5 := []int{0, -2147483648, 128, 870}
    if (true || (791 != 986)) {
        if (false && p0) {
            p0 = ((p0 || p0) && ((p0 || (1 >= 24)) && false))
        } else {
            if p0 {
                if (!p0) {
                    p0 = false
                    var v6 int8 = -9
                    for fi7 := 0; fi7 < 14; fi7++ {
                        fmt.Println(p0)
                        _ = fi7
                    }
                    _ = v6
                }
                fmt.Println(p1, p0)
                if (p0 || p0) {
                    var v8 int32 = -16
                    _ = v8
                }
            } else {
                if len(v5) > 0 {
                    v5[0] = -1
                }
                v5 = append(v5, 42)
                for fi9 := 0; fi9 < 7; fi9++ {
                    var v10 int8 = -1
                    _ = fi9
                    _ = v10
                }
            }
            fmt.Println(p0, p0)
            var v11 uint8 = 70
            _ = v11
        }
    } else {
        if (-2147483648 < -488) {
            v5 = append(v5, 873)
        } else {
            var v12 uint32 = 2147483648
            v13 := GenS2{F0: "", F1: 273}
            _ = v12
            _ = v13
        }
        v14 := GenS1{F0: 42, F1: 26, F2: 42}
        v5 = append(v5, 539)
        _ = v14
    }
    fmt.Println(p1)
    v15 := GenS1{F0: 2147483647, F1: 1, F2: 1}
    _ = p0
    _ = p1
    _ = v15
    _ = v4
    _ = v5
    return 535
}
func genFn1(p0 uint) int32 {
    for fi16 := 0; fi16 < 4; fi16++ {
        var v17 float64 = 1.5
        var v18 int = fi16
        _ = fi16
        _ = v17
        _ = v18
    }
    var v19 float32 = 0.1
    var v20 uint8 = 91
    _ = p0
    _ = v19
    _ = v20
    return -1
}

func main() {
    v1 := make(map[string]int)
    var v2 uint8 = 42
    v3 := make(map[int]bool)
    v21 := GenS3{F0: "a", F1: ("foo bar" + "a")}
    for fi22 := 0; fi22 < 10; fi22++ {
        var v23 int32 = 42
        _ = fi22
        _ = v23
    }
    var v24 uint64 = 42
    var v25 uint8 = v2
    for fi26 := 0; fi26 < 16; fi26++ {
        var v27 int64 = -9223372036854775808
        _ = fi26
        _ = v27
    }
    v25 -= 164
    var v28 int16 = int16(v2)
    v29 := GenS2{F0: ("foo bar" + ("hello" + ("foo bar" + "foo bar"))), F1: 2147483647}
    v30 := []int{127, 815}
    delete(v3, 4294967295)
    if false {
        if false {
            switch 4294967295 {
            case 0:
                v31 := GenS1{F0: 0, F1: -72, F2: 1}
                _ = v31
            case 260:
                switch ("hello" + "foo bar") {
                case "":
                    var v32 uint8 = v25
                    _ = v32
                }
            }
        } else {
            switch 714 {
            case 1:
                v33 := GenS3{F0: "hello", F1: "foo bar"}
                _ = v33
            case 9223372036854775807:
                v34 := GenS2{F0: "hello", F1: 127}
                _ = v34
            }
            if ((263 != 568) && true) {
                fmt.Println(v28, v25, v24)
                var v35 string = "foo bar"
                _ = v35
            } else {
                fmt.Println(v24, v28)
            }
            for fi36 := 0; fi36 < 19; fi36++ {
                v30 = append(v30, (^(fi36)))
                _ = fi36
            }
        }
        fmt.Println(v28, v28, v25)
        var v37 int = 1
        _ = v37
    } else {
        var v38 int8 = -128
        for fi39 := 0; fi39 < 4; fi39++ {
            v30 = append(v30, (fi39 + (fi39 % ((fi39) | 1))))
            v40 := GenS2{F0: "", F1: fi39}
            var v41 int32 = 0
            _ = fi39
            _ = v40
            _ = v41
        }
        var v42 int32 = 42
        _ = v38
        _ = v42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v30:", len(v30), v30)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
