package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64; F1 uint8 }
func genFn0(p0 uint, p1 int8, p2 uint64) int32 {
    switch "hello" {
    case "foo bar":
        v4 := GenS1{F0: 0, F1: -228, F2: p1}
        _ = v4
    case "hello":
        p1 = p1
    default:
        v5 := []int{679}
        _ = v5
    }
    var v6 int32 = 1
    if true {
        v7 := GenS2{F0: 255}
        p0 = (p0 >> ((65535) & 63))
        v8 := []int{1}
        _ = v7
        _ = v8
    } else {
        p1 = 0
        v6 -= v6
        var v9 uint8 = 10
        _ = v9
    }
    var v10 int = -1
    for fi11 := 0; fi11 < 13; fi11++ {
        p1 = (p1 / ((-6) | 1))
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v6
    return (v6 >> ((p0) & 63))
}
func genFn1(p0 bool, p1 int32) uint32 {
    for fi12 := 0; fi12 < 16; fi12++ {
        var v13 int16 = -32768
        var v14 uint = 18446744073709551615
        v15 := GenS2{F0: 1}
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    switch (("\n\t" + "foo bar") + "") {
    case "":
        var v16 float64 = 1e10
        _ = v16
    case "foo bar":
        var v17 uint8 = 2
        _ = v17
    }
    p1 = (p1 & genFn0(255, 1, 393))
    _ = p0
    _ = p1
    return 0
}

func main() {
    v1 := []int{552, -9223372036854775808, -2147483648}
    var v2 uint = 18446744073709551615
    v3 := []string{"hello"}
    v2 = v2
    for fi18 := 0; fi18 < 17; fi18++ {
        fmt.Println(v2)
        v2 += 135
        _ = fi18
    }
    for fi19 := 0; fi19 < 18; fi19++ {
        if true {
            v20 := GenS1{F0: 32767, F1: 42, F2: -1}
            if true {
                if true {
                    for fi21 := 0; fi21 < 11; fi21++ {
                        var v22 int8 = 116
                        v22 = v22
                        _ = fi21
                        _ = v22
                    }
                } else {
                    for fi23 := 0; fi23 < 5; fi23++ {
                        switch fi19 {
                        case -243:
                            var v24 uint16 = uint16(v2)
                            v2 -= (v2 / (((v2 & 42)) | 1))
                            _ = v24
                        default:
                            v2 = (v2 << (((v2 + v2)) & 63))
                        }
                        v3 = append(v3, "hello")
                        var v25 uint32 = 31
                        _ = fi23
                        _ = v25
                    }
                    fmt.Println(v2, v2)
                    v26 := GenS3{F0: 42}
                    _ = v26
                }
                v2 = v2
            }
            for fi27 := 0; fi27 < 16; fi27++ {
                var v28 int8 = 0
                var v29 uint16 = 482
                v3 = append(v3, "\n\t")
                v2 += v2
                var v30 uint = 2
                _ = fi27
                _ = v28
                _ = v29
                _ = v30
            }
            _ = v20
        }
        _ = fi19
    }
    v1 = append(v1, 572)
    if false {
        v31 := GenS2{F0: 255}
        v32 := GenS4{F0: uint64(v2), F1: 2}
        _ = v31
        _ = v32
    }
    _ = len(v1)
    switch "\n\t" {
    case "":
        var v33 int = 131
        _ = v33
    }
    for fi34 := 0; fi34 < 10; fi34++ {
        var v35 int16 = -32768
        v2 = v2
        var v36 bool = false
        _ = fi34
        _ = v35
        _ = v36
    }
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    if ((3.14159 != float64(1e10)) && (836 >= -9223372036854775808)) {
        _ = len(v3)
    } else {
        v3 = append(v3, "\n\t")
        var v37 uint64 = 1
        _ = v37
    }
    v38 := GenS2{F0: 0}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
}
