package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 int8; F2 string }
type GenS3 struct { F0 int32 }
func genFn0(p0 bool) int {
    var v7 int64 = -76
    if (2 > 18446744073709551615) {
        p0 = false
    }
    v8 := GenS1{F0: -2147483648}
    _ = p0
    _ = v7
    _ = v8
    return 2147483647
}
func genFn1(p0 string, p1 int, p2 uint32) bool {
    var v9 uint8 = 1
    fmt.Println(p2, p0)
    for fi10 := 0; fi10 < 8; fi10++ {
        switch fi10 {
        case 42:
            v11 := GenS2{F0: 0, F1: 127, F2: p0}
            _ = v11
        }
        fmt.Println(p2, p2)
        var v12 int16 = 1
        _ = fi10
        _ = v12
    }
    p1 = p1
    v9 = (v9 + (v9 & (v9 >> ((65535) & 63))))
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return ((((p2 < p2) && (((p0 + p0) + ("\n\t" + ("hello" + "foo bar"))) <= p0)) && (-69 < 1)) && true)
}

func main() {
    var v1 bool = true
    var v2 int = 2147483647
    v3 := make(map[int]bool)
    v4 := make(map[int]uint32)
    var v5 float64 = 0.0
    var v6 int = (v2 - v2)
    v6 = (v6 & genFn0((v2 >= v6)))
    v13 := GenS3{F0: -232}
    v6 = (-(v6))
    if (float64((3.14159 - (0.1 / (-2.25 - 0.1)))) > float64(((3.14159 + (1e10 - (-2.25 + -2.25))) / 1.0))) {
        v14 := v3[(v6 - v6)]
        var v15 string = "hello"
        delete(v3, -532)
        _ = v14
        _ = v15
    } else {
        switch "foo bar" {
        case "a":
            var v16 int32 = -1
            _ = v16
        case "foo bar":
            var v17 int32 = 42
            _ = v17
        }
        if true {
            switch 32768 {
            case -2147483648:
                if (v1 || ("hello" >= (("\n\t" + "") + "a"))) {
                    for fi18 := 0; fi18 < 3; fi18++ {
                        if (false || v1) {
                            v5 += 0.1
                            v2 = (v6 * (v6 ^ 821))
                            v19 := []int{32767, -71, v2}
                            _ = v19
                        }
                        v2 -= v2
                        var v20 int32 = -1
                        _ = fi18
                        _ = v20
                    }
                    v21 := []int{v2, -919}
                    _ = v21
                }
            default:
                if (415 != 18446744073709551615) {
                    for fi22 := 0; fi22 < 3; fi22++ {
                        var v23 uint = 2
                        _ = fi22
                        _ = v23
                    }
                } else {
                    v3[(-(v2))] = v1
                }
            }
        } else {
            v24 := []int{}
            switch genFn0((!(!(0 == 0)))) {
            case 4294967295:
                var v25 float64 = v5
                _ = v25
            }
            _ = v24
        }
    }
    var v26 bool = true
    v6 -= v2
    for fi27 := 0; fi27 < 1; fi27++ {
        for fi28 := 0; fi28 < 5; fi28++ {
            v5 = 0.1
            var v29 uint16 = 32767
            _ = fi28
            _ = v29
        }
        if v26 {
            var v30 int = 32767
            for fi31 := 0; fi31 < 19; fi31++ {
                v2 = (fi27 & fi27)
                _ = fi31
            }
            _ = v30
        } else {
            for fi32 := 0; fi32 < 16; fi32++ {
                var v33 int = v2
                v2 += v2
                _ = fi32
                _ = v33
            }
        }
        _ = fi27
    }
    if v1 {
        var v34 uint32 = 0
        _ = v34
    } else {
        var v35 uint16 = 111
        _ = v35
    }
    v36 := GenS3{F0: -683}
    v37 := GenS1{F0: -1}
    v38 := v3[(v6 + 32768)]
    v39 := []int{v6, genFn0((128 == 15))}
    var v40 uint8 = 214
    _ = len(v39)
    for fi41 := 0; fi41 < 6; fi41++ {
        var v42 uint8 = (v40 << ((4294967295) & 63))
        var v43 int8 = int8(v6)
        var v44 uint8 = 2
        _ = fi41
        _ = v42
        _ = v43
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v26
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
