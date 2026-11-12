package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int32 }

func main() {
    var v1 int32 = -287
    var v2 uint64 = 2
    v3 := make(map[string]bool)
    var v4 uint32 = 1
    v5 := make(map[uint32]int)
    delete(v5, (v4 / (((v4 % ((v4) | 1))) | 1)))
    var v6 int16 = int16(v2)
    if true {
        v1 = 0
        var v7 bool = (9223372036854775807 < 9223372036854775807)
        if v7 {
            if v7 {
                var v8 uint64 = (v2 << ((65535) & 63))
                _ = v8
            } else {
                var v9 uint64 = v2
                delete(v3, "\n\t")
                _ = v9
            }
        } else {
            for fi10 := 0; fi10 < 8; fi10++ {
                for fi11 := 0; fi11 < 13; fi11++ {
                    if v7 {
                        if (v1 >= v1) {
                            var v12 int64 = -834
                            v3["hello"] = true
                            var v13 uint8 = 198
                            v2 -= (v2 / (((v2 % ((330) | 1))) | 1))
                            _ = v12
                            _ = v13
                        }
                        switch 127 {
                        case -9223372036854775808:
                            fmt.Println(v6)
                        case -239:
                            v14 := GenS1{F0: uint32(v2), F1: 128, F2: v1}
                            _ = v14
                        }
                        var v15 float64 = 0.0
                        _ = v15
                    }
                    if ((v6 < (v6 & 777)) && v7) {
                        v1 = (-(v1))
                        v2 -= v2
                    }
                    fmt.Println(v6)
                    _ = fi11
                }
                fmt.Println(v1)
                _ = fi10
            }
            for fi16 := 0; fi16 < 2; fi16++ {
                var v17 int = fi16
                _ = fi16
                _ = v17
            }
        }
        _ = v7
    }
    for fi18 := 0; fi18 < 15; fi18++ {
        var v19 int = 42
        _ = fi18
        _ = v19
    }
    if false {
        v4 -= (v4 | v4)
    }
    var v20 bool = true
    v20 = v20
    v3[(("a" + "") + ("" + "foo bar"))] = (v6 == v6)
    v20 = v20
    var v21 uint = 732
    switch -2147483648 {
    case 2147483647:
        if ((!(100 >= 0)) || (false && (!((3.14159 - 0.0) >= 0.1)))) {
            var v22 float64 = (float64(float32(0.0)) / 1e10)
            v23 := []int{}
            var v24 float64 = -2.25
            _ = v22
            _ = v23
            _ = v24
        }
    case 1:
        var v25 int32 = (-(v1))
        _ = v25
    case -509:
        var v26 float32 = ((float32(1.5) - (1.5 - 1.5)) * (0.0 * (0.1 + ((-2.25 / 1.5) / 1e10))))
        _ = v26
    default:
        for fi27 := 0; fi27 < 15; fi27++ {
            v28 := []int{42, fi27, fi27}
            v20 = ((!(!(v20 && v20))) && v20)
            var v29 uint64 = (^(v2))
            _ = fi27
            _ = v28
            _ = v29
        }
    }
    var v30 int8 = 66
    v31 := GenS1{F0: (v4 << ((v21) & 63)), F1: 0, F2: 1}
    switch (("hello" + "hello") + "a") {
    case "hello":
        fmt.Println(v2, v1)
    case "":
        v3[("a" + "")] = (-925 > 2147483647)
    }
    var v32 int = 1
    if (uint8(v4) < 212) {
        var v33 string = ""
        v32 = v32
        if (true && (v30 >= v30)) {
            var v34 uint = (v21 | (v21 | uint(v4)))
            _ = v34
        }
        _ = v33
    } else {
        var v35 int = (-(v32))
        switch "a" {
        case "a":
            v36 := GenS1{F0: 554, F1: 255, F2: (v1 / ((42) | 1))}
            _ = v36
        case "\n\t":
            var v37 uint64 = (v2 + 2)
            _ = v37
        case "":
            fmt.Println(v2, v35)
        default:
            v38 := GenS1{F0: (v4 ^ v4), F1: 127, F2: v1}
            _ = v38
        }
        _ = len(v3)
        _ = v35
    }
    if true {
        v20 = v20
        switch "" {
        case "\n\t":
            for fi39 := 0; fi39 < 6; fi39++ {
                var v40 int64 = int64(v30)
                _ = fi39
                _ = v40
            }
        }
    } else {
        v1 -= v1
        fmt.Println(v21)
        v3[""] = v20
    }
    if ((-2.25 + 0.0) == 1.5) {
        v41 := GenS1{F0: (v4 - 65535), F1: 171, F2: v1}
        _ = v41
    } else {
        var v42 int64 = 900
        var v43 uint8 = 255
        var v44 uint32 = v4
        _ = v42
        _ = v43
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
