package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool }
func genFn0() int8 {
    var v7 string = "hello"
    v7 = ""
    if true {
        fmt.Println(v7)
        if true {
            var v8 int8 = 100
            var v9 float64 = ((-2.25 / 3.14159) / 3.14159)
            _ = v8
            _ = v9
        } else {
            v10 := []int{}
            if true {
                v10 = append(v10, 587)
            } else {
                fmt.Println(v7)
            }
            var v11 uint = 945
            _ = v10
            _ = v11
        }
    }
    _ = v7
    return 0
}

func main() {
    var v1 string = "a"
    var v2 uint32 = 1
    v3 := make(map[uint32]int)
    var v4 int = -9223372036854775808
    var v5 uint8 = 255
    var v6 int8 = 42
    if (-909 <= 32767) {
        if true {
            v6 += (v6 + (v6 - genFn0()))
            fmt.Println(v4, v2)
            if (false && (false || false)) {
                for fi12 := 0; fi12 < 9; fi12++ {
                    var v13 string = v1
                    _ = fi12
                    _ = v13
                }
                var v14 uint64 = 0
                v4 = (v4 + -2147483648)
                _ = v14
            } else {
                v15 := GenS1{F0: (v4 % ((v4) | 1))}
                _ = v15
            }
        }
    } else {
        _ = len(v3)
        switch "hello" {
        case "a":
            for fi16 := 0; fi16 < 14; fi16++ {
                for fi17 := 0; fi17 < 15; fi17++ {
                    var v18 int16 = -590
                    var v19 bool = (!false)
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                _ = fi16
            }
        case "foo bar":
            if false {
                v4 -= 4294967295
                if false {
                    var v20 int64 = -9223372036854775808
                    v20 += (v20 | v20)
                    switch "\n\t" {
                    case "hello":
                        var v21 string = ((v1 + v1) + v1)
                        _ = v21
                    case "":
                        delete(v3, (v2 & 330))
                    default:
                        v22 := GenS2{F0: true}
                        _ = v22
                    }
                    _ = v20
                }
                var v23 bool = false
                _ = v23
            }
        case "":
            if false {
                fmt.Println(v2)
            } else {
                var v24 int8 = v6
                _ = v24
            }
        default:
            var v25 int32 = 32
            _ = v25
        }
        var v26 int8 = v6
        _ = v26
    }
    var v27 uint16 = 2
    var v28 int32 = 42
    var v29 uint64 = 4294967295
    var v30 int32 = int32(v27)
    for fi31 := 0; fi31 < 8; fi31++ {
        v32 := GenS1{F0: fi31}
        v29 = (v29 + (^(v29)))
        _ = fi31
        _ = v32
    }
    for fi33 := 0; fi33 < 15; fi33++ {
        for fi34 := 0; fi34 < 3; fi34++ {
            v4 = int(v6)
            _ = fi34
        }
        fmt.Println(v27, v4)
        v30 = v30
        _ = fi33
    }
    if false {
        v4 = v4
        if false {
            var v35 uint8 = 127
            v36 := []int{(^(v4)), (-(v4))}
            _ = v35
            _ = v36
        } else {
            if (0.1 < ((-2.25 - 1e10) * 1e10)) {
                var v37 float64 = 0.0
                fmt.Println(v1)
                if ((v29 >= (v29 ^ uint64(v2))) && (v37 != (1e10 - v37))) {
                    v38 := GenS1{F0: 9223372036854775807}
                    _ = v38
                }
                _ = v37
            } else {
                var v39 uint = 2
                for fi40 := 0; fi40 < 3; fi40++ {
                    var v41 int16 = -1
                    v1 = v1
                    v42 := GenS1{F0: (fi40 / (((^(fi40))) | 1))}
                    _ = fi40
                    _ = v41
                    _ = v42
                }
                var v43 int16 = int16(v2)
                _ = v39
                _ = v43
            }
            v44 := v3[(v2 * 244)]
            delete(v3, (v2 / ((v2) | 1)))
            _ = v44
        }
    } else {
        var v45 int16 = 1000
        _ = v45
    }
    for fi46 := 0; fi46 < 16; fi46++ {
        v29 += (v29 ^ v29)
        _ = fi46
    }
    var v47 uint = 642
    fmt.Println(v6, v5)
    var v48 uint64 = v29
    v3[(v2 | (v2 ^ (v2 - (^(v2)))))] = (-(v4))
    v6 -= 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
