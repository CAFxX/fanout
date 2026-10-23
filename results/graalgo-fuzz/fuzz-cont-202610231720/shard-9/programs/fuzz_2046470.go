package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 string; F1 uint64; F2 int }
func genFn0(p0 uint16, p1 int8) int {
    var v5 int8 = 99
    var v6 float32 = ((1.5 * 3.14159) + (-2.25 - 0.0))
    var v7 float64 = (1.5 - 0.0)
    if true {
        v7 -= float64(v6)
        v5 -= v5
    } else {
        fmt.Println(v7, v6)
        var v8 int = 95
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 342
}

func main() {
    var v1 int8 = -128
    var v2 string = "hello"
    var v3 float64 = 0.0
    v4 := make(map[int]bool)
    switch genFn0(42, (v1 % ((v1) | 1))) {
    case 0:
        v2 = v2
    case -1:
        v9 := GenS1{F0: 2}
        _ = v9
    default:
        if true {
            v3 = v3
            for fi10 := 0; fi10 < 15; fi10++ {
                for fi11 := 0; fi11 < 1; fi11++ {
                    if (32767 > (fi10 + 0)) {
                        var v12 bool = false
                        v3 += 1.5
                        var v13 int16 = 32767
                        v3 += float64(3.14159)
                        _ = v12
                        _ = v13
                    }
                    v14 := GenS1{F0: 716}
                    _ = fi11
                    _ = v14
                }
                if (!(2147483647 == 2147483648)) {
                    v4[-9223372036854775808] = (!false)
                    fmt.Println(v2, v3, v1)
                } else {
                    for fi15 := 0; fi15 < 10; fi15++ {
                        v3 += v3
                        var v16 uint64 = 4294967295
                        _ = fi15
                        _ = v16
                    }
                }
                v2 = ""
                _ = fi10
            }
        }
    }
    if false {
        v17 := GenS2{F0: 1, F1: -1}
        v18 := GenS1{F0: 349}
        fmt.Println(v1, v2)
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 8; fi19++ {
        if true {
            var v20 uint64 = 1
            var v21 int64 = -1
            _ = v20
            _ = v21
        }
        for fi22 := 0; fi22 < 7; fi22++ {
            switch (fi19 & 4294967295) {
            case 13:
                v23 := GenS1{F0: 65535}
                v24 := GenS2{F0: 1, F1: 9223372036854775807}
                _ = v23
                _ = v24
            }
            _ = fi22
        }
        for fi25 := 0; fi25 < 8; fi25++ {
            if false {
                v26 := GenS3{F0: "\n\t", F1: 4294967295, F2: fi25}
                var v27 uint64 = 66
                if (42 <= -9223372036854775808) {
                    if false {
                        v1 = v1
                        v3 = (v3 - (0.1 / 1e10))
                        v28 := GenS3{F0: ("" + v2), F1: 0, F2: int(v1)}
                        v1 += (v1 - (v1 + (v1 >> ((877) & 63))))
                        v29 := []int{genFn0(91, -79), fi25, -2147483648, fi25}
                        _ = v28
                        _ = v29
                    }
                    var v30 int = 2147483647
                    if false {
                        fmt.Println(v30)
                        v3 += 0.0
                    }
                    var v31 string = (v2 + "foo bar")
                    var v32 int32 = 0
                    _ = v30
                    _ = v31
                    _ = v32
                } else {
                    var v33 int = fi19
                    v34 := GenS2{F0: 1, F1: -471}
                    v35 := []int{(^(fi25)), 0, genFn0(2, -128), (fi25 & fi19)}
                    _ = v33
                    _ = v34
                    _ = v35
                }
                _ = v26
                _ = v27
            } else {
                var v36 uint64 = 18446744073709551615
                var v37 uint16 = 65535
                var v38 int64 = 523
                _ = v36
                _ = v37
                _ = v38
            }
            fmt.Println(v1, fi25, v2)
            v39 := GenS1{F0: 2}
            _ = fi25
            _ = v39
        }
        _ = fi19
    }
    v3 -= (((v3 * 0.0) + ((v3 + v3) - float64(-2.25))) - v3)
    v2 = v2
    if (v1 == (v1 & (-(v1)))) {
        v40 := []int{4294967295, 4294967295}
        fmt.Println(v2, v2, v1)
        _ = v40
    } else {
        v1 += (v1 + (v1 & v1))
        var v41 float64 = 0.1
        _ = v41
    }
    v2 = v2
    v2 = ((v2 + "\n\t") + v2)
    v42 := GenS1{F0: 65535}
    var v43 uint = 18446744073709551615
    if ((v3 * v3) != ((float64(1e10) + ((float64(-2.25) - 0.1) / v3)) * v3)) {
        var v44 string = v2
        var v45 string = ("a" + (("\n\t" + ("" + v44)) + "foo bar"))
        _ = v44
        _ = v45
    }
    fmt.Println(v43)
    if false {
        var v46 bool = false
        _ = v46
    }
    for fi47 := 0; fi47 < 5; fi47++ {
        v3 = 1e10
        _ = fi47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
}
