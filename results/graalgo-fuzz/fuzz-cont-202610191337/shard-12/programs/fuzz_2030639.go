package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 int16; F2 int16 }
type GenS3 struct { F0 int }
func genFn0() bool {
    switch "\n\t" {
    case "":
        v7 := []int{519, 789, -219, 128}
        _ = v7
    }
    v8 := []int{}
    var v9 int32 = 2147483647
    _ = v8
    _ = v9
    return (false && true)
}

func main() {
    var v1 uint32 = 42
    var v2 int16 = 1000
    v3 := []int{0}
    v4 := make(map[int]uint32)
    var v5 uint8 = 127
    var v6 int = 544
    v6 = 2147483647
    for fi10 := 0; fi10 < 18; fi10++ {
        if false {
            var v11 int = (fi10 - (fi10 + -9223372036854775808))
            var v12 string = "foo bar"
            _ = v11
            _ = v12
        } else {
            fmt.Println(fi10, fi10)
            delete(v4, v6)
        }
        v2 -= 1
        for fi13 := 0; fi13 < 11; fi13++ {
            if false {
                fmt.Println(v5, v1)
            } else {
                var v14 uint8 = v5
                _ = v14
            }
            if (false || (302 < (v6 % (((fi10 - (v6 * (^(fi10))))) | 1)))) {
                if false {
                    v15 := GenS1{F0: -1}
                    switch ("\n\t" + "\n\t") {
                    case "":
                        var v16 int = v6
                        _ = v16
                    case "foo bar":
                        v17 := GenS1{F0: fi10}
                        _ = v17
                    case "\n\t":
                        fmt.Println(fi10)
                        var v18 int = fi10
                        if (!false) {
                            fmt.Println(v2)
                        } else {
                            v19 := GenS2{F0: 4294967295, F1: (v2 | v2), F2: v2}
                            _ = v19
                        }
                        _ = v18
                    }
                    _ = v15
                } else {
                    var v20 float32 = -2.25
                    _ = v20
                }
                if (510 != 4294967295) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                } else {
                    v2 += (v2 - v2)
                    fmt.Println(v2)
                }
                var v21 int = fi10
                _ = v21
            }
            if true {
                var v22 bool = ("foo bar" != ("" + ""))
                delete(v4, (fi10 * 127))
                var v23 int32 = 42
                _ = v22
                _ = v23
            }
            _ = fi13
        }
        _ = fi10
    }
    for fi24 := 0; fi24 < 11; fi24++ {
        var v25 uint8 = 127
        switch v6 {
        case -270:
            var v26 int32 = -2147483648
            _ = v26
        }
        var v27 uint64 = 4294967295
        _ = fi24
        _ = v25
        _ = v27
    }
    var v28 int32 = -2147483648
    if false {
        v5 = v5
        var v29 int64 = 0
        _ = v29
    }
    var v30 uint8 = (v5 >> ((793) & 63))
    if (270 <= -9223372036854775808) {
        _ = len(v4)
        v31 := GenS3{F0: (v6 / ((v6) | 1))}
        var v32 uint = 4294967295
        _ = v31
        _ = v32
    } else {
        _ = len(v3)
    }
    var v33 int = v6
    v3 = append(v3, v33)
    if (v6 < (v33 | (^(v33)))) {
        var v34 int64 = -1
        _ = v34
    }
    v30 -= v30
    v35 := GenS1{F0: v33}
    if (v30 == 42) {
        switch v33 {
        case -582:
            for fi36 := 0; fi36 < 4; fi36++ {
                var v37 bool = ((v28 - v28) == v28)
                v38 := GenS2{F0: 255, F1: v2, F2: (^(v2))}
                v5 = 16
                _ = fi36
                _ = v37
                _ = v38
            }
        case 0:
            for fi39 := 0; fi39 < 12; fi39++ {
                delete(v4, v33)
                var v40 int32 = v28
                var v41 int64 = 399
                _ = fi39
                _ = v40
                _ = v41
            }
        case 2147483647:
            v42 := GenS1{F0: (v33 - v6)}
            _ = v42
        }
        if (-96 <= int8(v2)) {
            v5 -= v30
            v1 += v1
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v33
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
